import type { Room, CreateRoomRequest, JoinRoomRequest } from "~/util/models";
import { loadPlayerFromCookie } from "~/util/playerCookie";
import { saveGameStateToCookies } from "~/util/roomCookie";
import { handleApiError, validatePlayerSession, validateRoomId } from "~/util/errorUtils";

export const useRoomActions = () => {
  /**
   * Creates a new room with the specified player count
   */
  const createRoom = async (maxPlayers: number): Promise<Room> => {
    const player = loadPlayerFromCookie();
    validatePlayerSession(player);

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
      handleApiError(error);
    }
  };

  /**
   * Joins an existing room by ID
   */
  const joinRoom = async (roomId: string): Promise<Room> => {
    validateRoomId(roomId);

    const player = loadPlayerFromCookie();
    validatePlayerSession(player);

    try {
      const requestBody: JoinRoomRequest = {
        id: player.id,
      };

      const room: Room = await $fetch<Room>(`api/rooms/${roomId.trim()}/players`, {
        method: "POST",
        body: requestBody,
      });

      return room;
    } catch (error) {
      handleApiError(error);
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
    createRoom,
    joinRoom,
    enterLobby,
  };
};
