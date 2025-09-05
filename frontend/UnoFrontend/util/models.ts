/**
 * TypeScript type definitions for API data structures
 */

/** Player entity representing a game participant */
export type Player = {
    id: string    // Unique UUID identifier
    name: string  // Display name
}

/** Room entity representing a game lobby */
export type Room = {
    id: string        // Unique UUID identifier
    players: Player[] // List of all players in room
    owner: Player     // Room creator/host
}