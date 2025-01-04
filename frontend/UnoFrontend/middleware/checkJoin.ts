import type {Room} from "~/util/models";
import {getIDFromCookie} from "~/util/getIDFromCookie";

export default defineNuxtRouteMiddleware(async (to, from) => {
    if (import.meta.server) return

    if (to.fullPath.indexOf("/lobby") > -1 ) {
        const roomID = to.fullPath.split("/lobby-")[1]
        if (!roomID) {
            window.alert("Invalid room ID.");
            return navigateTo("/hostOrJoin");
        }
        let room: Room | null = null
        try {
            room = await $fetch(`/api/room/${roomID}`);
        } catch (error) {
            window.alert('no room found with ID: ' + roomID);
            return navigateTo("/hostOrJoin")
        }
        const playerID = getIDFromCookie()
        console.log(room?.players)
        if(room && !(room?.players.some(player => player.id === playerID))) {
            try {
                const responseRoom: Room = await $fetch(`api/room/${roomID}/players`, {
                        method: 'POST',
                        body: {
                            id: playerID,
                        },
                    }
                );
                console.log(responseRoom);
                useState('room',() => responseRoom);
                navigateTo(`/lobby-${roomID}`);
            } catch (error: any) {
                if (error?.response?.status === 409) {
                    return
                }
                console.error('Error communicating with internal API:', error);
                return navigateTo("/hostOrJoin")
            }
        }
    }
})