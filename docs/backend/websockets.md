# Websocket
Communication of game events after a room has been started is transported via websockets.
The messages are build like this:
```go
struct message {
  Type    string  // the name of the payload type (the tags below + `Payload`)
  Payload any     // the struct instance from below
}
```
The backend has 2 distinct communication channels:
1. Broadcast: Sends the message to all players, the info is not sensitive to the players.
1. Single Player: Sends the message to a single player. May contain things like actual cards the player has.

## Payloads
The payloads send and received by the websocket. 

### Error
**Send to:**  The causing player  
**Send when:** An error occurs - technical or game play based.  
**Structure:**   
- `Code` (int): An error code, roughly based on http errors.
- `Msg` (string): A more specific error message.
### GameStart  
**Send to:** All players  
**Send when:** The player cards have bean dealt at the start of the round.  
**Structure:**   
- `TopCard` (any): The top card on the stack.
### CardPlayed  
**Send to:** All players  
**Send when:** A player has played a card.  
**Structure**:  
- `PlayerId` (UUID): The id of the player that played a card.
- `Name` (string): The display name of the player.
- `Card` (any): The card played.
### PlayerTurn
**Send to:** All players  
**Send when:** The next players turn begins.  
**Structure:**  
- `PlayerId` (UUID): The id of the player, whose turn starts.
- `Name` (string): The display name of the player.
### PlayerWin
**Send to:** All players    
**Send when:** A player wins the game.    
**Structure:**  
- `PlayerId` (UUID): The id of the player that won.  
- `Name` (string): The display name of the player.  

### PlayerDrawsCards  
**Send to:** All players  
**Send when:** A player draws cards, e.g. forced by +2 cards or not having any fitting or initial dealing.   
**Structure:**   
- `PlayerId` (UUID): The id of the player that draws cards.
- `Name` (string): The display name of the player.
- `Amount` (int): The number of cards drawn.

### PlayerSkipped
**Send to:** All players   
**Send when:** A player gets skipped by a skip card.  
**Structure:**   
- *Empty payload*  

### DirectionChanged
**Send to:** All players  
**Send when:** The direction is changed, e.g. by a reverse card.  
**Structure:**   
- `Direction` (int): The current play direction (e.g., clockwise or counterclockwise).  

### PlayerChoseColor
**Send to:** All players  
**Send when:** A player chooses the color after playing a wildcard.  
**Structure:**   
- `PlayerId` (UUID): The id of the player who chose a color.  
- `Name` (string): The display name of the player.  
- `Color` (int): The chosen color.

### AskColor
**Send to:** The player who played the wildcard.  
**Send when:** A wildcard card that requires color selection is played.   
**Structure:**  
- `Options` ([]int): The available colors to choose from.  

### AnswerColor
**Send to:** The server    
**Send when:** A player responds to an `AskColor` event.   
**Structure:**  
- `Chosen` (int): The color chosen by the player.  

### AskCard
**Send to:** The current player  
**Send when:** The players turn starts and a card needs to be played.  
**Structure:**    
- `Options` (any[]): The available cards to choose from.  

### AnswerCard
**Send to:** The server    
**Send when:** A player responds to an `AskCard` event.    
**Structure:**  
- `Card` (any): The card chosen by the player.  

### YouDrawCard
**Send to:** The player who draws    
**Send when:** The player has to draw card(s).  
**Structure:**  
- `Cards` ([]any): The cards drawn by the player.  

### RoomJoin
**Send to:** All players  
**Send when:** A new player joins the room.  
**Structure:**  
- `PlayerId` (UUID): The id of the player joining.
- `Name` (string): The name of the player joining.

### RoomLeft
**Send to:** All players  
**Send when:** A player leaves the room.  
**Structure:**  
- `PlayerId` (UUID): The id of the player leaving.
- `Name` (string): The name of the player leaving.
- `OwnerId` (UUID): The id of the player owning the room. The room owner is changed if the old one is leaving.
- `OwnerName` (UUID): The name of the owner.

### RoomStart
**Send to:** All players  
**Send when:** A room is started via [Start](/docs/backend/restapi.md#start).  
**Structure:**  
- `Players` (\[\][Player](/docs/backend/models.md#player)): The players in the room.

