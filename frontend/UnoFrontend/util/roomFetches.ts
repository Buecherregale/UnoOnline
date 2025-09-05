/**
 * API wrapper class for room-related operations
 */
export class RoomFetches {
    baseURL: string;
    UUID: string;

    /**
     * Initialize RoomFetches with base URL and player UUID
     * @param baseURL - Backend API base URL
     * @param UUID - Player's unique identifier
     */
    constructor(baseURL: string, UUID: string) {
        this.baseURL = baseURL;
        this.UUID = UUID;
    }

    /**
     * Create a new room via API
     * @param playerID - ID of player creating the room
     * @param baseURL - Backend API base URL
     */
    async createRoom(playerID: string, baseURL: string) {
        // Make POST request to create room
        const { data: room, error} = await useFetch('/room', {
                method: 'POST',
                body: {
                    id: playerID,
                },
                baseURL : baseURL,
            }
        )

        // Handle API errors
        if (error) {
            throw error.value
        } else {
            // Debug output for successful creation
            console.log(room)
            console.log(room.value)
        }
    }
}
