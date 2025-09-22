import type { Room, StartGameRequest } from "~/util/models";
import { handleApiError } from "~/util/errorUtils";

export const useGameApi = () => {
  const makeGameRequest = async (
    endpoint: string,
    playerId: string,
    method: "POST" = "POST"
  ): Promise<Room | void> => {
    try {
      const requestBody: StartGameRequest = { id: playerId };

      return await $fetch<Room>(endpoint, {
        method,
        body: requestBody,
      });
    } catch (error) {
      handleApiError(error);
    }
  };

  const startRoom = (gameId: string, playerId: string) =>
    makeGameRequest(`/api/rooms/${gameId}`, playerId);

  const startGame = (gameId: string, playerId: string) =>
    makeGameRequest(`/api/rooms/${gameId}/startgame`, playerId);

  const loadRoomData = async (gameId: string): Promise<Room | null> => {
    try {
      return await $fetch<Room>(`/api/rooms/${gameId}`);
    } catch (error) {
      handleApiError(error);
      return null;
    }
  };

  return {
    startRoom,
    startGame,
    loadRoomData,
  };
};
