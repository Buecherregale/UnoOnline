import type { message, Player } from "~/util/models";

/**
 * WebSocket connection states
 */
export enum WebSocketState {
  CONNECTING = 'connecting',
  CONNECTED = 'connected',
  DISCONNECTED = 'disconnected',
  ERROR = 'error'
}

/**
 * WebSocket event handlers interface
 */
export interface WebSocketEventHandlers {
  onRoomUpdate?: (player: Player) => void;
  onPlayerJoined?: (player: Player) => void;
  onPlayerLeft?: (playerId: string) => void;
  onGameStarted?: (roomId: string) => void;
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

export default class WebSocketHelper {
  private socket: WebSocket | null = null;
  private messageQueue: message[] = [];
  private state: WebSocketState = WebSocketState.DISCONNECTED;
  private config: Required<WebSocketConfig>;
  private eventHandlers: WebSocketEventHandlers = {};
  private reconnectAttempts: number = 0;

  constructor(playerId: string, roomId: string, baseUrl: string = 'ws://localhost:8080') {
    this.config = {
      playerId,
      roomId,
      baseUrl,
      reconnectAttempts: 3,
      reconnectDelay: 1000
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
        console.log('WebSocket connected');
        this.setState(WebSocketState.CONNECTED);
        this.reconnectAttempts = 0;
        this.processMessageQueue();
      };

      this.socket.onmessage = (event: MessageEvent): void => {
        this.handleMessage(event);
      };

      this.socket.onclose = (event: CloseEvent): void => {
        console.log('WebSocket closed:', event.code, event.reason);
        this.setState(WebSocketState.DISCONNECTED);
        this.handleReconnection();
      };

      this.socket.onerror = (event: Event): void => {
        console.error('WebSocket error:', event);
        this.setState(WebSocketState.ERROR);
        this.eventHandlers.onError?.(new Error('WebSocket connection failed'));
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
      console.log('Message received:', msg);

      let player: Player | null  = null;
      let newOwner: Player | null = null;
      switch (msg.type) {
        case 'RoomJoinPayload':
          player = msg.payload.player as Player;
          this.eventHandlers.onPlayerJoined?.(player);
          break;
        case 'RoomLeftPayload':
          player = msg.payload.player as Player;
          newOwner = msg.payload.newOwner as Player | null;
          this.eventHandlers.onPlayerLeft?.(msg.payload);
          break;
        case 'GameStartedPayload':
          this.eventHandlers.onGameStarted?.(msg.payload);
          break;
        default:
          console.warn('Unknown message type:', msg.type);
      }
    } catch (error) {
      console.error('Error parsing WebSocket message:', error);
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
      console.log(`Attempting to reconnect... (${this.reconnectAttempts}/${this.config.reconnectAttempts})`);

      setTimeout(() => {
        this.connect();
      }, this.config.reconnectDelay * this.reconnectAttempts);
    } else {
      console.error('Max reconnection attempts reached');
      this.eventHandlers.onError?.(new Error('Connection lost and max reconnection attempts reached'));
    }
  }

  /**
   * Sends a message via WebSocket
   */
  private sendMessage(message: message): boolean {
    if (this.socket && this.socket.readyState === WebSocket.OPEN) {
      try {
        this.socket.send(JSON.stringify(message));
        console.log('Message sent:', message);
        return true;
      } catch (error) {
        console.error('Error sending message:', error);
        this.eventHandlers.onError?.(error as Error);
        return false;
      }
    } else {
      console.warn('WebSocket is not connected, message queued');
      return false;
    }
  }

  /**
   * Queues a message for sending when connection is available
   */
  private queueMessage(message: message): void {
    if (this.isConnected()) {
      this.sendMessage(message);
    } else {
      this.messageQueue.push(message);
      console.log('Message queued:', message);
    }
  }

  // Public API methods with explicit return types

  /**
   * Checks if WebSocket is currently connected
   */
  public isConnected(): boolean {
    return this.state === WebSocketState.CONNECTED &&
           this.socket?.readyState === WebSocket.OPEN;
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
   * Legacy method for backward compatibility
   */
  public setRoomUpdateCallback(callback: (player: Player) => void): void {
    this.eventHandlers.onRoomUpdate = callback;
  }

  /**
   * Sends player joined message
   */
  public playerJoined(player: Player): void {
    const message: message = {
      type: 'RoomJoinPayload',
      payload: JSON.stringify(player)
    };
    this.queueMessage(message);
  }

  /**
   * Sends player left message
   */
  public playerLeft(playerId: string): void {
    const message: message = {
      type: 'PlayerLeftPayload',
      payload: playerId
    };
    this.queueMessage(message);
  }

  /**
   * Sends start game message
   */
  public startGame(): void {
    const message: message = {
      type: 'StartGamePayload',
      payload: this.config.roomId
    };
    this.queueMessage(message);
  }

  /**
   * Manually disconnects WebSocket and clears resources
   */
  public disconnect(): void {
    if (this.socket) {
      this.socket.close(1000, 'Manual disconnect');
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
