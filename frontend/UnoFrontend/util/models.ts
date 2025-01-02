export type Player = {
    id: string
    name: string
}

export type Room = {
    id: string
    players: Player[]
    owner: Player
}