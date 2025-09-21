import type { Room, CreateRoomRequest, JoinRoomRequest } from "~/util/models";
import { loadPlayerFromCookie } from "~/util/playerCookie";
import { saveGameStateToCookies } from "~/util/roomCookie";
import {
  handleApiError,
  validatePlayerSession,
  validateRoomId,
} from "~/util/errorUtils";
import { usePlayerStore } from "~/stores/player";
import { useRoomStore } from "~/stores/room";

export const useRoomActions = () => {
  /**
   * Creates a new room with the specified player count
   */
  const createRoom = async (maxPlayers: number): Promise<Room> => {
    const playerStore = usePlayerStore();

    let playerId = playerStore.getPlayerId;
    if (!playerId || playerId.trim() === "") {
      const player = loadPlayerFromCookie();
      validatePlayerSession(player);
      playerId = player.id;
    }

    try {
      const requestBody: CreateRoomRequest = {
        id: playerId,
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
    const playerStore = usePlayerStore();

    validateRoomId(roomId);

    let playerId = playerStore.getPlayerId;
    if (!playerId || playerId.trim() === "") {
      const player = loadPlayerFromCookie();
      validatePlayerSession(player);
      playerId = player.id;
    }

    try {
      const requestBody: JoinRoomRequest = {
        id: playerId,
      };

      const room: Room = await $fetch<Room>(
        `api/rooms/${roomId.trim()}/players`,
        {
          method: "POST",
          body: requestBody,
        }
      );

      return room;
    } catch (error) {
      handleApiError(error);
    }
  };

  /**
   * Saves room state and navigates to lobby
   */
  const enterLobby = async (room: Room, isHost: boolean): Promise<void> => {
    const roomStore = useRoomStore();

    saveGameStateToCookies(room, isHost);
    roomStore.setRoom(room, isHost);

    await navigateTo(`/lobby-${room.id}`);
  };

  return {
    createRoom,
    joinRoom,
    enterLobby,
  };
};
