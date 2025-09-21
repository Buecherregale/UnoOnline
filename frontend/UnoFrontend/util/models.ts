/**
 * TypeScript type definitions for API data structures
 */

/** Player entity representing a game participant */
export type Player = {
  id: string; // Unique UUID identifier
  name: string; // Display name
};

/** Room entity representing a game lobby */
export type Room = {
  id: string; // Unique UUID identifier
  players: Player[]; // List of all players in room
  owner: Player; // Room creator/host
};

// Component Props Interfaces
export interface PlayerCardProps {
  player: Player;
  isCurrentPlayer?: boolean;
  cardCount?: number;
}

export interface GameBoardProps {
  players: Player[];
  currentPlayerId: string;
  roomId: string;
}

// Event Types
export interface PlayerJoinedEvent {
  type: "player-joined";
  player: Player;
  room: Room;
}

export interface PlayerLeftEvent {
  type: "player-left";
  playerId: string;
  room: Room;
}

export interface GameStartedEvent {
  type: "game-started";
  roomId: string;
}

export interface RoomUpdatedEvent {
  type: "room-updated";
  room: Room;
}
// Form Data Types
export interface CreatePlayerRequest {
  name: string;
}

export interface CreateRoomRequest {
  id: string;
  maxPlayers?: number;
}

export interface JoinRoomRequest {
  id: string;
}

export interface LeaveRoomRequest {
  id: string;
}

export interface StartGameRequest {
  id: string;
}

// Navigation Types
export interface NavigationState {
  player: Player | null;
  room: Room | null;
  isHost: boolean;
}

// Utility Types
export type PlayerPosition = "top" | "bottom" | "left" | "right";

export interface PlayerWithPosition {
  player: Player;
  position: PlayerPosition;
}

export type message = {
  type: string; // the name of the payload type (the tags below + `Payload`)
  payload: any; // the struct instance from below
  MessageId: string; // unique id to trace messages. Important for answers to messages
  ExpectsAnswer: boolean // if true the server expects the client to reply with the apropriate message with the SAME MessageId
};
