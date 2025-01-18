# API Documentation

Documentation for the Rest and Websocket API.

## Endpoints

### Room Endpoints

#### Create Room
**POST** `/room`

- **Description:** Creates a new room. The owner only has to join the new websocket.
- **Request Body:** 
  - `id` (UUID): The id of the player to own the room.  
- **Response:**
  - **201 Created**: Room successfully created.
  - **Error Responses:**
    - 404 - Player does not exist.

---

#### Get Room Info
**GET** `/room/{id}`

- **Description:** Retrieves information about the room: players, owner, etc.
- **Path Parameters:**
  - `id` (UUID): Unique identifier of the room.
- **Response:**
  - **200 OK:** Room details.
  - **Error Responses:** 
    - 404 - Room does not exist.

---

#### Start Room
**POST** `/room/{id}`

- **Description:** Starts the game for this room. 
- **Path Parameters:**
  - `id` (UUID): Unique identifier of the room.
- **Request Body:**
  - `id` (UUID): The id of the player. 
- **Response:**
  - **200 OK:** Game started.
  - **Error Responses:** 
    - 403 - Player is not the owner. 
    - 404 - Room does not exist or Player does not exist.
    - 409 - The Number of Players in the Room is not equal to those connected to the Websocket

---

#### Join Room
**POST** `/room/{id}/players`

- **Description:** Lets a player join a room. The player has to additionally join the websocket afterwards. 
- **Path Parameters:**
  - `id` (UUID): Unique identifier of the room.
- **Request Body:**
  - `id` (UUID): The id of the player. 
- **Response:**
  - **200 OK:** Room details.
  - **Error Responses:** 
    - 404 - Room does not exist.

---

#### Leave Room
**DELETE** `/room/{id}/players`

- **Description:** Lets a player leave the room. If the last player leaves the room is destroyed. 
- **Path Parameters:**
  - `id` (UUID): Unique identifier of the room.
- **Request Body:**
  - `id` (UUID): The id of the player. 
- **Response:**
  - **200 OK:** Player removed from room and websocket.
  - **Error Responses:** 
    - 404 - Room does not exist or Player does not exist. 

### Player Endpoints

#### Create Player
**POST** `/player`

- **Description:** Registers a player.  
- **Request Body:**
  - `name` (string): The name of the player. 
- **Response:**
  - **200 OK:** The newly created player. Contains the id that has to be used for future requests.

### Websocket
`/ws`

- **Description:** Join the Websocket of a room as the player. Moves and Gamestate are transferred via this socket. 
- **Query Parameter:** 
  - `roomId` (UUID): The id of the room to join.
  - `playerId` (UUID): The id of the player to join as.
