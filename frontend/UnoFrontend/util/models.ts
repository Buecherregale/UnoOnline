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

// API Response Types
export interface ApiResponse<T = unknown> {
  data?: T;
  error?: string;
  success: boolean;
}

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

export interface ModalProps {
  isVisible: boolean;
  title: string;
  onClose: () => void;
  onConfirm?: () => void;
}

export interface HostGameModalProps extends ModalProps {
  selectedPlayerCount: number;
  onPlayerCountChange: (count: number) => void;
}

export interface JoinGameModalProps extends ModalProps {
  lobbyId: string;
  onLobbyIdChange: (id: string) => void;
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

export type WebSocketEvent =
  | PlayerJoinedEvent
  | PlayerLeftEvent
  | GameStartedEvent
  | RoomUpdatedEvent;

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
};
