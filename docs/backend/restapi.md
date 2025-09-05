# REST-Api Endpoints
### Players
#### Create
**POST** `/players`
- **Description:** Registers a player.  
- **Request Body:**
  - `name` (string): The name of the player. 
- **Response:**
  - `201 - Created` ([Player](models.md#Player)): The newly created player. Contains the id that has to be used for future requests.
### Rooms
#### Create
**POST** `/rooms`
- **Description:** Creates a new room. The owner only has to join the new websocket.
- **Request Body:** 
  - `id` (UUID): The id of the player to own the room.  
- **Response:**
  - `201 - Created`: Room successfully created.
  - **Error Responses:**
    -  `404 - Not found`: Player does not exist.
---
#### Get Info
**GET** `/rooms/{id}`

- **Description:** Retrieves information about the room: players, owner, etc.
- **Path Parameters:**
  - `id` (UUID): Unique identifier of the room.
- **Response:**
  - `200 - OK`([Room](models.md#Room)): Room details.
  - **Error Responses:** 
    - `404 - Not found`: Room does not exist.
---
#### Start
**POST** `/rooms/{id}`
- **Description:** Starts the game for this room. 
- **Path Parameters:**
  - `id` (UUID): Unique identifier of the room.
- **Request Body:**
  - `id` (UUID): The id of the player. 
- **Response:**
  - `200 OK`: Game started.
  - **Error Responses:** 
    - `403 - Forbidden`: Player is not the owner. 
    - `404 - Not found`: Room or player does not exist.
    - `409 - Conflict`: The number of players in the room is not equal to those connected to the websocket.
---
#### Join
**POST** `/rooms/{id}/players`
- **Description:** Lets a player join a room. The player has to additionally join the websocket afterwards. 
- **Path Parameters:**
  - `id` (UUID): Unique identifier of the room.
- **Request Body:**
  - `id` (UUID): The id of the player. 
- **Response:**
  - `200 - OK` ([Room](models.md#Room)): Details of the room.
  - **Error Responses:** 
    - `404 - Not found`: Room does not exist.

---
#### Leave
**DELETE** `/rooms/{id}/players`

- **Description:** Lets a player leave the room. If the last player leaves the room is destroyed. 
- **Path Parameters:**
  - `id` (UUID): Unique identifier of the room.
- **Request Body:**
  - `id` (UUID): The id of the player. 
- **Response:**
  - `200 - OK`: Player removed from room and websocket.
  - **Error Responses:** 
    - `404 - Not found`: Room or player does not exist. 
### Websocket
**ANY** `/ws`
- **Description:** Connects the player to the websocket of the room.
- **URL Parameters:**
  - `roomId` (UUID): The unique room identifier.
  - `playerId` (UUID): The unique player identifier.
- **Response:**
  - `OK`: Connection upgrade to websocket.
  - **Error Responses:**
    - `400 - Bad request`: Invalid url paramters.
    - `404 - Not found`: Room or player does not exist.
    - `403 - Forbidden`: Player did not join the room via (Join)[#Join].

