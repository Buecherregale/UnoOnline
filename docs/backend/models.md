# API Models
### Player
- Id (UUID): Unique player id required for endpoints.
- Name (string): Display name of the player.
### Room
- Id (UUID): Unique room id used in endpoint parameters. 
- Players (Player[]): Current players in the room.
- Owner (Player): Owning player - the one able to start the game. 
 
