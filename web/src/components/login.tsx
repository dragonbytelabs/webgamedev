import { css } from "@linaria/core";
import { createSignal } from "solid-js";

const container = css`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 100vh;
  background-color: #f8f9fa;
  font-family: Arial, sans-serif;
`;

const card = css`
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1);
  width: 90%;
  max-width: 400px;
`;

const title = css`
  font-size: 24px;
  font-weight: bold;
  margin-bottom: 20px;
  text-align: center;
  color: #333;
`;

const inputField = css`
  width: 100%;
  padding: 10px;
  margin: 10px 0;
  border-radius: 8px;
  border: 1px solid #ccc;
  font-size: 16px;
`;

const buttonPrimary = css`
  width: 100%;
  padding: 12px;
  margin-top: 20px;
  border-radius: 8px;
  background: #4a1e79;
  color: white;
  border: none;
  font-size: 16px;
  cursor: pointer;
`;

const signupText = css`
  margin-top: 12px;
  font-size: 14px;
  color: #555;
  text-align: center;

  & a {
    color: #4a1e79;
    text-decoration: none;
    font-weight: bold;
  }
`;

export default function LoginWithSocials() {
	const [username, setUsername] = createSignal("");
	const [password, setPassword] = createSignal("");

	const postForm = async () => {
		// Build body the same way <form> would
		const body = new URLSearchParams();
		body.set("username", username());
		body.set("password", password());

		const r = await fetch("/api/login", {
			method: "POST",
			body, // browser sets correct Content-Type header
			credentials: "include",
		});

		if (!r.ok) throw new Error("login failed");
		// If Go redirects (302/303/307), fetch will follow automatically
		// and return the final response; you may instead use `redirect: "manual"`
	};

	return (
		<div class={container}>
			<div class={card}>
				<h1 class={title}>Login</h1>
				<input
					type="text"
					placeholder="Email or Username"
					class={inputField}
					value={username()}
					onInput={(e) => setUsername(e.currentTarget.value)}
				/>
				<input
					type="password"
					placeholder="Password"
					class={inputField}
					value={password()}
					onInput={(e) => setPassword(e.currentTarget.value)}
				/>
				<button type="button" class={buttonPrimary} onClick={postForm}>
					Sign In
				</button>
				<p class={signupText}>
					Don’t Have An Account? <a href="/signup">Sign Up</a>
				</p>
			</div>
		</div>
	);
}
