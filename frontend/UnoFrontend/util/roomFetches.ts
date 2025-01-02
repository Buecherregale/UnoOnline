export class RoomFetches {
    baseURL: string;
    UUID: string;

    constructor(baseURL: string, UUID: string) {
        this.baseURL = baseURL;
        this.UUID = UUID;
    }

    async createRoom(playerID: string, baseURL: string) {
        const { data: room, error} = await useFetch('/room', {
                method: 'POST',
                body: {
                    id: playerID,
                },
                baseURL : baseURL,
            }
        )
        if (error) {
            throw error.value
        } else {
            console.log(room)
            console.log(room.value)
        }
    }
}
