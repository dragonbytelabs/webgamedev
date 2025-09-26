package dbx

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/dragonbytelabs/webgamedev/db"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

type DB struct {
	SQL *sql.DB
	DBX *sqlx.DB
}

func (d *DB) Close() error { return d.SQL.Close() }

// OpenSQLite opens SQLite with sane defaults for server use.
func OpenSQLite(path string) (*DB, error) {
	// modernc driver name is "sqlite"
	// Busy timeout + WAL via pragmas in first migration; we can also pass flags in DSN.
	// For example: file:app.db?_pragma=busy_timeout(5000)
	dsn := fmt.Sprintf(
		"file:%s?_pragma=journal_mode(WAL)&_pragma=busy_timeout(5000)&_pragma=foreign_keys(ON)",
		path,
	)
	sqlDB, err := sql.Open("sqlite", dsn)
	if err != nil {
		return nil, err
	}
	// Connection pool tuning — SQLite is single file; keep modest limits.
	sqlDB.SetMaxOpenConns(1)
	sqlDB.SetMaxIdleConns(1)
	sqlDB.SetConnMaxLifetime(0)

	dbx := sqlx.NewDb(sqlDB, "sqlite")
	return &DB{SQL: sqlDB, DBX: dbx}, nil
}

// ApplyMigrations runs embedded DDL in lexical order (001_*.sql, 002_*.sql, ...).
func (d *DB) ApplyMigrations(ctx context.Context) error {
	dir := "migrations"
	entries, err := fs.ReadDir(db.MigrationsFS, dir)
	if err != nil {
		return err
	}
	fmt.Printf("entries: %v\n", entries)

	// Ensure deterministic order
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(name, ".sql") {
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)

	tx, err := d.SQL.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	for i, name := range names {
		path := filepath.Join(dir, name)
		b, readErr := db.MigrationsFS.ReadFile(path)
		if readErr != nil {
			return fmt.Errorf("read migration %s: %w", name, readErr)
		}

		log.Printf("applying migration %d/%d: %s (%d bytes)", i+1, len(names), name, len(b))

		if _, execErr := tx.ExecContext(ctx, string(b)); execErr != nil {
			// Surface the *real* DB error so you can see the cause
			return fmt.Errorf("migration %s failed: %w", name, execErr)
		}

		log.Printf("applied migration: %s", name)
	}

	if err != nil {
		return err
	}
	return tx.Commit()
}

// Helpers to read query text
func MustQuery(name string) string {
	path := filepath.Join("db/queries", name)
	b, err := db.QueriesFS.ReadFile(path)
	if err != nil {
		log.Fatalf("query %s not found: %v", name, err)
	}
	return string(b)
}

// Example data types
type User struct {
	ID           int64      `db:"id" json:"id"`
	Email        string     `db:"email" json:"email"`
	PasswordHash string     `db:"password_hash" json:"-"`
	DisplayName  *string    `db:"display_name" json:"display_name,omitempty"`
	CreatedAt    time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt    *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}

// CreateUser (sqlx + named params)
func (d *DB) CreateUser(ctx context.Context, email, passwordHash string, display *string) (*User, error) {
	q := MustQuery("create_user.sql")
	args := map[string]any{
		"email":         email,
		"password_hash": passwordHash,
		"display_name":  display,
	}
	var u User
	// RETURNING works on SQLite 3.35+ (modernc bundles recent SQLite)
	if err := d.DBX.GetContext(ctx, &u, q, args); err != nil {
		return nil, err
	}
	return &u, nil
}

// GetUserByEmail (sqlx + named params)
func (d *DB) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	q := MustQuery("get_user_by_email.sql")
	var u User
	if err := d.DBX.GetContext(ctx, &u, q, map[string]any{"email": email}); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
