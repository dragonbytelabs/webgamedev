# Scenes
Game Structure Overview

Entry Conditions:
Scene: number
inventory: {}
player: Player

Action Cursors - (F)oot/(T)alk/(G)rab/(E)yes:
- Talk (talking head): talk to NPCs 
- Look (eyeball): inspect items → lore/tooltips
- Use/Take (hand): pick up items → satchel/inventory
- Move (foot): walk/leave scene

## S01 — Introduction
Entry Condition: 
SHOW_INTRO = True
Story: S01.[01]


Core Flags:
ALARM_CLICKED: boolean
SNOOZE_TESTED: boolean
HAS_REGISTERED: boolean

Collectables:
- Empty Jar
- Tail warmer
- Inventer's book

### S01.[01] - Player's Bedroom
Entry Condition: 
SHOW_INTRO = True

Description: The player is asleep in bed 

Trigger: 
- Click on Alarm clock - 01.<02>

### S01.<02> - Alarm Clock 
Prompt: “Click the alarming clock to begin.”

Next: S01.<02>.[01] OR S01.<02>.[02]

### S01.<02> --> S01.<02>.[01] - ALARM_CLICKED === True

Flags:
- SNOOZE_TESTED = true

Snooze Gag:
Player (VoiceOver): “The alarming clock is a splendid invention, but adding the snooze lever is pure genius.”

Next: S01.<02>.[03] 

### S01.<02> --> S01.<02>.[02] - ALARM_CLICKED === False 
Description: The player is asleep in bed until alarm clock is clicked

Next: S01.<02>

### S01.<02>.[03] - Jester Enters (Cutscene)

Jester: “Player, wake up, you lazy set of dragon droppings. King Allfire has been waiting to see you.”
Player: “I just want to test the snooze lever one more time. I’ll be down in 9 minutes.”
Jester: “Fine! I’ll tell the King—our Absolute Sovereign, Master of all He Surveys—that you can’t fit him in your busy schedule.”
Player: “Whoa! Hold your jingle bells!”

Next: 01.<02>.[04]

### S01.<02>.[04] - Dress & Jester Exits
Cutscene: Player hurridly gets out of bed and gets dressed.
Player: “Tell the King I’m awake.”
Jester exits.

Next: [S02]

## S02 - Player Prepares to meet King
Entry Condition: HAS_REGISTERED === False
Exit Condition: HAS_REGISTERED === True

Next: S02.[01]

### S02.[01] - Player Monologue
Player (Voice Over): “Maybe it’s about my request to battle in the tournament tomorrow. If I can’t win, Princess Flame will be forced to marry someone else.”

Flags:
SHOW_INTRO = false.

Next: S02.[02]

### S02.[02] - Player Setup (Registration)

Player registration modal with inputs:
- Username
- Email
- Password
- Confirm Password

Next: S02.<03>

### S02.<03> - HAS_REGISTERED 
Check if player is registered. If they are then have them sign in
If not then create a new player with input fields

Next: S02.<03>.[01] OR S02.<03>.[02]

### S02.<03> --> S02.<03>.[01] - HAS_REGISTERED === TRUE  
Create the player in DB then set flag HAS_REGISTERED === TRUE
Close Modal

Next: S02.<03>.[03]

### S02.<03> --> S02.<03>.[02] - HAS_REGISTERED === FALSE

TODO: 
- Show Disclaimer? All unsaved data will be lost. 
- Do we allow the player to continue? 
- How do we handle the game state? 

Close Modal

Next: S02.<03>.[03] 

### S02.<03>.[03] - Player has control starting point is bedroom

Interact (F)oot/(T)alk/(G)rab/(E)yes:
- TGE: Pet Moth in jar
- E: Window
- GE: Tail Warmer
- E: Picture on nightstand of Princess Flame
- GE: Bed
- E: Junk in corner
- E: Alarming clock
- F: Doorway to Hallway
- Invention Book

Collectibles:
- Tail Warmer → INVENTORY.tailWarmer
- Empty Jar → INVENTORY.emptyJar 
- Invention Book → INVENTORY.inventionBook 

Next: [S03] 

### [S03] - Hallway outside bedroom 

Interact (F)oot/(T)alk/(G)rab/(E)yes:
- FE: Library
- GE: Candelabra
- F: Meeting area (DOM table)
- F: Bedroom 

Collectibles:
- Candelabra 

Next: S03.<01>

###  S03.<01> - On the way to meet the King 
Player can choose to Go back to Bedroom or enter library or go to the meeting with the King

Next: S03.<01>.[01] OR S03.<01>.[02] OR S03.<01>.[03]

###  S03.<01> --> S03.<01>.[01] - Go back to Bedroom 
The player chooses to go back to bedroom
Next: S02.<03>.[03]

###  S03.<01> --> S03.<01>.[02] - Pop into the library 
Player enters library. Librarian is not here. Nothing to interact with yet until Player has been requested by King

Next: S03.<01>

###  S03.<01> --> S03.<01>.[03] - Head to meeting place at the DOM table
Player chooses to head straight to the meetin with the KIng

Next: [S04]





