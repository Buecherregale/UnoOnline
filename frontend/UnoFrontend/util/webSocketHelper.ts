import type { Card, message, Player } from "~/util/models";
import { parseCardPayload } from "~/util/cardParser";

/**
 * WebSocket connection states
 */
export enum WebSocketState {
  CONNECTING = "connecting",
  CONNECTED = "connected",
  DISCONNECTED = "disconnected",
  ERROR = "error",
}

/**
 * WebSocket event handlers interface
 */
export interface WebSocketEventHandlers {
  onPlayerJoined?: (player: Player) => void;
  onPlayerLeft?: (player: Player, newOwner: Player) => void;
  onRoomStarted?: (roomId: string) => void;
  onGameStarted?: (card: Card) => void;
  onDrawCard?: (cards: Card[]) => void;
  onPlayerDrawsCards?: (player: Player, amount: number) => void;
  onError?: (error: Error) => void;
  onStateChange?: (state: WebSocketState) => void;
}

/**
 * WebSocket connection configuration
 */
export interface WebSocketConfig {
  playerId: string;
  roomId: string;
  baseUrl?: string;
  reconnectAttempts?: number;
  reconnectDelay?: number;
}

/**
 * WebSocket Helper - Handles low-level WebSocket operations
 * Responsibilities:
 * - WebSocket connection management
 * - Message queuing and sending
 * - Automatic reconnection logic
 * - Protocol-level error handling
 */
export default class WebSocketHelper {
  private socket: WebSocket | null = null;
  private messageQueue: message[] = [];
  private state: WebSocketState = WebSocketState.DISCONNECTED;
  private config: Required<WebSocketConfig>;
  private eventHandlers: WebSocketEventHandlers = {};
  private reconnectAttempts: number = 0;

  constructor(
    playerId: string,
    roomId: string,
    baseUrl: string = "ws://localhost:8080"
  ) {
    this.config = {
      playerId,
      roomId,
      baseUrl,
      reconnectAttempts: 3,
      reconnectDelay: 1000,
    };

    this.connect();
  }

  /**
   * Establishes WebSocket connection
   */
  private connect(): void {
    const url = `${this.config.baseUrl}/ws?roomId=${this.config.roomId}&playerId=${this.config.playerId}`;

    try {
      this.setState(WebSocketState.CONNECTING);
      this.socket = new WebSocket(url);

      this.socket.onopen = (): void => {
        console.log("WebSocket connected");
        this.setState(WebSocketState.CONNECTED);
        this.reconnectAttempts = 0;
        this.processMessageQueue();
      };

      this.socket.onmessage = (event: MessageEvent): void => {
        this.handleMessage(event);
      };

      this.socket.onclose = (event: CloseEvent): void => {
        console.log("WebSocket closed:", event.code, event.reason);
        this.setState(WebSocketState.DISCONNECTED);
        this.handleReconnection();
      };

      this.socket.onerror = (event: Event): void => {
        console.error("WebSocket error:", event);
        this.setState(WebSocketState.ERROR);
        this.eventHandlers.onError?.(new Error("WebSocket connection failed"));
      };
    } catch (error) {
      this.setState(WebSocketState.ERROR);
      this.eventHandlers.onError?.(error as Error);
    }
  }

  /**
   * Handles incoming WebSocket messages
   */
  private handleMessage(event: MessageEvent): void {
    try {
      const msg: message = JSON.parse(event.data);
      console.log("Message received:", msg);

      let player: Player | null = null;
      let newOwner: Player | null = null;
      let amount: number = 0;
      switch (msg.type) {
        case "RoomJoinPayload":
          player = msg.payload.player as Player;
          this.eventHandlers.onPlayerJoined?.(player);
          break;
        case "RoomLeftPayload":
          player = msg.payload.player as Player;
          newOwner = msg.payload.owner as Player;
          this.eventHandlers.onPlayerLeft?.(player, newOwner);
          break;
        case "RoomStartPayload":
          //validation ob returned player gleiche wie in room
          this.eventHandlers.onRoomStarted?.(msg.payload);
          break;
        case "GameStartPayload":
          let Card = parseCardPayload(msg.payload);
          console.log("Game started:", Card);
          this.eventHandlers.onGameStarted?.(Card);
          break;
        case "YouDrawCardPayload":
          let Cards = msg.payload.cards as any[];
          let parsedCards: Card[] = [];
          for (let i = 0; i < Cards.length; i++) {
            let card = parseCardPayload(Cards[i]);
            parsedCards.push(card);
          }
          this.eventHandlers.onDrawCard?.(parsedCards);
          break;
        case "PlayerDrawsCardsPayload":
          player = msg.payload.player as Player;
          amount = msg.payload.amount as number;
          this.eventHandlers.onPlayerDrawsCards?.(player, amount);
          break;
        default:
          console.warn("Unknown message type:", msg.type);
      }
    } catch (error) {
      console.error("Error parsing WebSocket message:", error);
      this.eventHandlers.onError?.(error as Error);
    }
  }

  /**
   * Sets the WebSocket state and notifies handlers
   */
  private setState(newState: WebSocketState): void {
    if (this.state !== newState) {
      this.state = newState;
      this.eventHandlers.onStateChange?.(newState);
    }
  }

  /**
   * Processes queued messages when connection is established
   */
  private processMessageQueue(): void {
    while (this.messageQueue.length > 0 && this.isConnected()) {
      const message = this.messageQueue.shift();
      if (message) {
        this.sendMessage(message);
      }
    }
  }

  /**
   * Handles automatic reconnection logic
   */
  private handleReconnection(): void {
    if (this.reconnectAttempts < this.config.reconnectAttempts) {
      this.reconnectAttempts++;
      console.log(
        `Attempting to reconnect... (${this.reconnectAttempts}/${this.config.reconnectAttempts})`
      );

      setTimeout(() => {
        this.connect();
      }, this.config.reconnectDelay * this.reconnectAttempts);
    } else {
      console.error("Max reconnection attempts reached");
      this.eventHandlers.onError?.(
        new Error("Connection lost and max reconnection attempts reached")
      );
    }
  }

  /**
   * Sends a message via WebSocket
   */
  private sendMessage(message: message): void {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      try {
        this.socket.send(JSON.stringify(message));
        console.log("Message sent:", message);
      } catch (error) {
        console.error("Error sending message:", error);
      }
    } else {
      console.warn("WebSocket is not connected, message queued");
    }
  }

  /**
   * Queues a message for sending when connection is available
   */
  public queueMessage(message: message): void {
    if (this.isConnected()) {
      this.sendMessage(message);
    } else {
      console.log("Message queued:", message);
      this.messageQueue.push(message);
    }
  }

  /**
   * Checks if WebSocket is currently connected
   */
  public isConnected(): boolean {
    return (
      this.state === WebSocketState.CONNECTED &&
      this.socket?.readyState === WebSocket.OPEN
    );
  }

  /**
   * Gets current WebSocket state
   */
  public getState(): WebSocketState {
    return this.state;
  }

  /**
   * Sets event handlers for WebSocket events
   */
  public setEventHandlers(handlers: WebSocketEventHandlers): void {
    this.eventHandlers = { ...this.eventHandlers, ...handlers };
  }

  /**
   * Manually disconnects WebSocket and clears resources
   */
  public disconnect(): void {
    if (this.socket) {
      this.socket.close(1000, "Manual disconnect");
      this.socket = null;
    }
    this.setState(WebSocketState.DISCONNECTED);
    this.messageQueue = [];
    this.eventHandlers = {};
  }

  /**
   * Manually triggers reconnection
   */
  public reconnect(): void {
    this.disconnect();
    this.reconnectAttempts = 0;
    this.connect();
  }
}
