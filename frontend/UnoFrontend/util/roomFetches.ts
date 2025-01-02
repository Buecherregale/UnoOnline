export class RoomFetches {
    baseURL: string;
    UUID: string;

    constructor(baseURL: string, UUID: string) {
        this.baseURL = baseURL;
        this.UUID = UUID;
    }

    async createRoom(name: string, baseURL: string) {
        /*const { data, error} = await useFetch('/api/player', {
                method: 'POST',
                body: JSON.stringify({
                    name,
                }),
                baseURL : baseURL
            }
        )
        if (error) {
            throw error.value
        } else {
            return data.value
        }*/
        return "TestPlayerUUID";
    }
}
