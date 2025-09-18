import type { Room, CreateRoomRequest, JoinRoomRequest } from "~/util/models";
import { loadPlayerFromCookie } from "~/util/playerCookie";
import { saveGameStateToCookies } from "~/util/roomCookie";

export const useRoomActions = () => {
  const isLoading = ref<boolean>(false);

  /**
   * Creates a new room with the specified player count
   */
  const createRoom = async (maxPlayers: number): Promise<Room> => {
    const player = loadPlayerFromCookie();

    if (!player?.id) {
      throw createError({
        statusCode: 400,
        message: "Player not found in session",
      });
    }

    isLoading.value = true;

    try {
      const requestBody: CreateRoomRequest = {
        id: player.id,
        maxPlayers,
      };

      const room: Room = await $fetch<Room>("/api/rooms", {
        method: "POST",
        body: requestBody,
      });

      return room;
    } catch (error) {
      console.error("Error creating room:", error);
      throw createError({
        statusCode: 500,
        message: "Failed to create room",
      });
    } finally {
      isLoading.value = false;
    }
  };

  /**
   * Joins an existing room by ID
   */
  const joinRoom = async (roomId: string): Promise<Room> => {
    const player = loadPlayerFromCookie();

    if (!player?.id) {
      throw createError({
        statusCode: 400,
        message: "Player not found in session",
      });
    }

    isLoading.value = true;

    try {
      const requestBody: JoinRoomRequest = {
        id: player.id,
      };

      const room: Room = await $fetch<Room>(`api/rooms/${roomId}/players`, {
        method: "POST",
        body: requestBody,
      });

      return room;
    } catch (error) {
      console.error("Error joining room:", error);
      throw createError({
        statusCode: 500,
        message: "Failed to join room",
      });
    } finally {
      isLoading.value = false;
    }
  };

  /**
   * Saves room state and navigates to lobby
   */
  const enterLobby = async (room: Room, isHost: boolean): Promise<void> => {
    saveGameStateToCookies(room, isHost);

    useState<Room>("rooms", () => room);
    useState<boolean>("isHost", () => isHost);

    await navigateTo(`/lobby-${room.id}`);
  };

  return {
    isLoading: readonly(isLoading),
    createRoom,
    joinRoom,
    enterLobby,
  };
};
