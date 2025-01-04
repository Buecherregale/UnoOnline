import type {Room} from "~/util/models";

export default defineNuxtRouteMiddleware((to, from) => {
    if (import.meta.server) return

    if(from.fullPath.indexOf("/lobby") > -1 && to.fullPath !== from.fullPath) {
        if (!window.confirm('You are in a Lobby are you Sure you want to leave?')) {
            return abortNavigation()
        }
        useState<Room | null>('room').value = null
    }
})