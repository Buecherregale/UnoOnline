import type {message, Player, Room} from "~/util/models";

export default class WebSocketHelper {
    private socket: WebSocket | null = null;
    private messageQueue: message[] = [];
    private isConnected: boolean = false;
    private onRoomUpdateCallback: ((room: Room) => void) | null = null;

    constructor(playerId: string, roomId: string) {
        let url = `ws://localhost:8080/ws?roomId=${roomId}&playerId=${playerId}`;
        this.socket = new WebSocket(url);

        this.socket.onopen = () => {
            console.log('WebSocket connected');
            this.isConnected = true;
            this.processMessageQueue();
        };

        this.socket.onmessage = (event) => {
            // Nachrichten vom Server verarbeiten
            console.log('Message received:', event.data);
            // check payload type
            const msg: message = JSON.parse(event.data);
            if (msg.type === 'RoomJoinPayload') {
                // update room
                const roomData: Room = JSON.parse(msg.payload);
                if (this.onRoomUpdateCallback) {
                    this.onRoomUpdateCallback(roomData);
                }
            }
        };

        this.socket.onclose = () => {
            console.log('WebSocket closed');
            this.isConnected = false;
        };

        this.socket.onerror = (event) => {
            console.error('WebSocket error:', event);
            this.isConnected = false;
        };
    }

    private processMessageQueue(): void {
        while (this.messageQueue.length > 0 && this.isConnected) {
            const message = this.messageQueue.shift();
            if (message) {
                this.sendMessage(message);
            }
        }
    }

    private sendMessage(message: message): void {
        if (this.socket && this.socket.readyState === WebSocket.OPEN) {
            this.socket.send(JSON.stringify(message));
            console.log('Message sent:', message);
        } else {
            console.error('WebSocket is not connected');
        }
    }

    private queueMessage(message: message): void {
        if (this.isConnected) {
            this.sendMessage(message);
        } else {
            this.messageQueue.push(message);
            console.log('Message queued:', message);
        }
    }

    public playerJoined(player: Player): void {
        const message: message = { type: 'RoomJoinPayload', payload: JSON.stringify(player) };
        this.queueMessage(message);
    }

    public setRoomUpdateCallback(callback: (room: Room) => void): void {
        this.onRoomUpdateCallback = callback;
    }

    public disconnect(): void {
        if (this.socket) {
            this.socket.close();
            this.socket = null;
            this.isConnected = false;
            this.messageQueue = [];
        }
    }
}
