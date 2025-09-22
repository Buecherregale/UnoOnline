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

export type CardColor = "red" | "green" | "blue" | "yellow" | "black";

export type CardValue =
  | number
  | "skip"
  | "reverse"
  | "plus2"
  | "wild"
  | "wildcard4";

export interface Card {
  color: CardColor;
  value: CardValue;
  chosen: CardColor | null; // For wild cards, the color chosen by the player
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

// Utility Types
export type PlayerPosition = "top" | "bottom" | "left" | "right";

export interface PlayerWithPosition {
  player: Player;
  position: PlayerPosition;
}

export type uuid = `${string}-${string}-${string}-${string}-${string}`;

export type message = {
  type: string; // the name of the payload type (the tags below + `Payload`)
  payload: any; // the struct instance from below
  message_id: uuid; // unique id to trace messages. Important for answers to messages
  expects_reply: boolean; // if true the server expects the client to reply with the apropriate message with the SAME MessageId
};
