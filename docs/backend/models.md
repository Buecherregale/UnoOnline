# Data Transfer Objects
Both REST Api and websocket communicate with these dtos, containing the necessary information, which is always available.
### Player
- `Id` (UUID): Unique player id required for endpoints.
- `Name` (string): Display name of the player.
### Room
- `Id` (UUID): Unique room id used in endpoint parameters. 
- `Players` (Player[]): Current players in the room.
- `Owner` (Player): Owning player - the one able to start the game. 
 ### Card
 There is no card dto type. In go cards are denoted as `any` (see [messages.go](/backend/api/ws/messages.go)).  
 The actual implementation of card depends on the card array a game is constructed with.
 ## Uno
 ### UnoCard
 For default uno cards the type looks like [this](/backend/uno/uno.go).  
 - `Color` (Color): The color of the card.
 - `Value` (Value): The (numeric) value.
 - `Chosen` (Color): For wildcards, this field contains the color the player has chosen.
 ### Color
 - `Red`
 - `Blue`
 - `Yellow`
 - `Green`
 - `Black`
 ### Value
 - `0` to `9`
 - `Skip`: Skips the next player.
 - `Reverse`: Changes the direction of play.
 - `Plus2`: Forces the player to draw a card.
 - `Wildcard`: Lets the player choose a color.
 - `Wildcard4`: Lets the player choose a color and forces the next player to draw 4 cards.
