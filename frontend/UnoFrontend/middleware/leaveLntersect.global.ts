export default defineNuxtRouteMiddleware((to, from) => {
    if (import.meta.server) return

    if(from.fullPath.indexOf("/lobby") > -1) {
        if (!window.confirm('You are in a Lobby are you Sure you want to leave?')) {
            return abortNavigation()
        }
    }
})