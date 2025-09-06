import type { Room } from './models';

const COOKIE_OPTIONS = {
  maxAge: 60 * 60 * 24, // 24 Stunden
  httpOnly: false,
};

/**
 * saves rooms data to a cookie
 */
export function saveRoomToCookie(room: Room): void {
  try {
    const roomCookie = useCookie<string>('uno-rooms', {
      ...COOKIE_OPTIONS,
      default: () => '',
    });

    roomCookie.value = JSON.stringify(room);
    console.log('Room saved to cookie:', { roomId: room.id });
  } catch (error) {
    console.error('Failed to save rooms to cookie:', error);
  }
}

/**
 * loads rooms data from a cookie
 */
export function getRoomFromCookie(): Room | null {
  try {
    const roomCookie = useCookie<string>('uno-rooms', {
      ...COOKIE_OPTIONS,
      default: () => '',
    });

    if (roomCookie.value && roomCookie.value !== '') {
      return JSON.parse(roomCookie.value) as Room;
    }
    return null;
  } catch (error) {
    console.error('Failed to load rooms from cookie:', error);
    return null;
  }
}

/**
 * saves host status to a cookie
 */
export function saveHostStatusToCookie(isHost: boolean): void {
  try {
    const hostCookie = useCookie<boolean>('uno-is-host', {
      ...COOKIE_OPTIONS,
      default: () => false,
    });

    hostCookie.value = isHost;
    console.log('Host status saved to cookie:', { isHost });
  } catch (error) {
    console.error('Failed to save host status to cookie:', error);
  }
}

/**
 * loads host status from a cookie
 */
export function getHostStatusFromCookie(): boolean {
  try {
    const hostCookie = useCookie<boolean>('uno-is-host', {
      ...COOKIE_OPTIONS,
      default: () => false,
    });

    return hostCookie.value;
  } catch (error) {
    console.error('Failed to load host status from cookie:', error);
    return false;
  }
}

/**
 * deletes game-related cookies
 */
export function clearGameCookies(): void {
  try {
    const roomCookie = useCookie<string>('uno-rooms');
    const hostCookie = useCookie<boolean>('uno-is-host');

    roomCookie.value = '';
    hostCookie.value = false;

    console.log('Game cookies cleared');
  } catch (error) {
    console.error('Failed to clear game cookies:', error);
  }
}

export function saveGameStateToCookies(room: Room, isHost: boolean): void {
  saveRoomToCookie(room);
  saveHostStatusToCookie(isHost);
}
