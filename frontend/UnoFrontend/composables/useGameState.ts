import type { Player, Room } from "~/util/models";
import { loadPlayerFromCookie } from "~/util/playerCookie";
import { loadRoomFromCookie } from "~/util/roomCookie";
import { validatePlayerSession, validateRoomId } from "~/util/errorUtils";

export const useGameState = () => {
  const roomStore = useRoomStore();
  const playerStore = usePlayerStore();

  const initializePlayer = (): Player => {
    let currentPlayer = playerStore.getPlayer;
    if (!currentPlayer) {
      const tmp = loadPlayerFromCookie();
      validatePlayerSession(tmp);
      currentPlayer = tmp;
      playerStore.setPlayer(currentPlayer);
    }
    return currentPlayer;
  };

  const initializeRoom = (): Room => {
    let room = roomStore.getRoom;
    if (!room) {
      const tmp = loadRoomFromCookie();
      validateRoomId(tmp!.id);
      room = tmp;
      roomStore.updateRoom(room!);
    }
    return room!;
  };

  return {
    initializePlayer,
    initializeRoom,
  };
};
