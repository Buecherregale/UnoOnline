export const playerFetches = async (name: string, baseURL: string) => {
    const { data: playerID, error} = await useFetch('/player', {
            method: 'POST',
            body: {
                name: name,
            },
            baseURL : baseURL,
            transform:(_playerID: any) => _playerID.id,
        }
    )
    if (error) {
        throw error.value
    } else {
        return playerID
    }
}