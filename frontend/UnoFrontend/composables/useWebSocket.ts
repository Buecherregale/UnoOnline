import WebSocketHelper, {
  type WebSocketEventHandlers,
  WebSocketState,
} from "~/util/webSocketHelper";

/**
 * Global WebSocket service that persists across page navigation
 * Responsibilities:
 * - Application-level state management
 * - Vue.js integration and reactivity
 * - Cross-page WebSocket persistence
 * - Event handler lifecycle management
 */
class WebSocketService {
  private static instance: WebSocketService | null = null;
  private wsHelper: WebSocketHelper | null = null;
  private currentRoomId: string | null = null;
  private currentPlayerId: string | null = null;
  private eventHandlers: WebSocketEventHandlers = {};

  private constructor() {}

  public static getInstance(): WebSocketService {
    if (!WebSocketService.instance) {
      WebSocketService.instance = new WebSocketService();
    }
    return WebSocketService.instance;
  }

  /**
   * Connects to WebSocket if not already connected
   * Smart connection management - reuses existing connections when possible
   */
  public connect(playerId: string, roomId: string): void {
    // Only create new connection if parameters changed or no connection exists
    if (
      !this.wsHelper ||
      this.currentPlayerId !== playerId ||
      this.currentRoomId !== roomId ||
      !this.wsHelper.isConnected()
    ) {
      // Disconnect existing connection if it exists
      this.disconnect();

      this.currentPlayerId = playerId;
      this.currentRoomId = roomId;
      this.wsHelper = new WebSocketHelper(playerId, roomId);

      // Apply existing event handlers to new connection
      if (Object.keys(this.eventHandlers).length > 0) {
        this.wsHelper.setEventHandlers(this.eventHandlers);
      }

      console.log(`WebSocket service connected for player ${playerId} in room ${roomId}`);
    } else {
      console.log("WebSocket service reusing existing connection");
    }
  }

  /**
   * Sets event handlers for WebSocket events
   */
  public setEventHandlers(handlers: WebSocketEventHandlers): void {
    this.eventHandlers = { ...this.eventHandlers, ...handlers };
    if (this.wsHelper) {
      this.wsHelper.setEventHandlers(this.eventHandlers);
    }
  }

  /**
   * Adds additional event handlers without overwriting existing ones
   */
  public addEventHandlers(handlers: WebSocketEventHandlers): void {
    this.eventHandlers = { ...this.eventHandlers, ...handlers };
    if (this.wsHelper) {
      this.wsHelper.setEventHandlers(this.eventHandlers);
    }
  }

  /**
   * Removes specific event handlers
   */
  public removeEventHandlers(handlerKeys: (keyof WebSocketEventHandlers)[]): void {
    handlerKeys.forEach((key) => {
      delete this.eventHandlers[key];
    });
    if (this.wsHelper) {
      this.wsHelper.setEventHandlers(this.eventHandlers);
    }
  }

  /**
   * Checks if WebSocket is connected
   */
  public isConnected(): boolean {
    return this.wsHelper?.isConnected() ?? false;
  }

  /**
   * Gets current WebSocket state
   */
  public getState(): WebSocketState {
    return this.wsHelper?.getState() ?? WebSocketState.DISCONNECTED;
  }

  /**
   * Gets current room ID
   */
  public getCurrentRoomId(): string | null {
    return this.currentRoomId;
  }

  /**
   * Gets current player ID
   */
  public getCurrentPlayerId(): string | null {
    return this.currentPlayerId;
  }

  /**
   * Sends start game message
   */
  public startGame(): void {
    this.wsHelper?.startGame();
  }

  /**
   * Manually reconnects WebSocket
   */
  public reconnect(): void {
    this.wsHelper?.reconnect();
  }

  /**
   * Disconnects WebSocket and clears resources
   */
  public disconnect(): void {
    if (this.wsHelper) {
      this.wsHelper.disconnect();
      this.wsHelper = null;
    }
    this.currentRoomId = null;
    this.currentPlayerId = null;
    this.eventHandlers = {};
  }

  /**
   * Force disconnect (used when leaving room completely)
   */
  public forceDisconnect(): void {
    this.disconnect();
  }
}

/**
 * Composable for using the global WebSocket service
 */
export const useWebSocket = () => {
  const service = WebSocketService.getInstance();

  return {
    connect: service.connect.bind(service),
    setEventHandlers: service.setEventHandlers.bind(service),
    addEventHandlers: service.addEventHandlers.bind(service),
    removeEventHandlers: service.removeEventHandlers.bind(service),
    isConnected: service.isConnected.bind(service),
    getState: service.getState.bind(service),
    getCurrentRoomId: service.getCurrentRoomId.bind(service),
    getCurrentPlayerId: service.getCurrentPlayerId.bind(service),
    startGame: service.startGame.bind(service),
    reconnect: service.reconnect.bind(service),
    disconnect: service.disconnect.bind(service),
    forceDisconnect: service.forceDisconnect.bind(service),
  };
};
