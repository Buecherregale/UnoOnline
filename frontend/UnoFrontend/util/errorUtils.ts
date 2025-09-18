/**
 * Creates a consistent error message based on status code
 */
export function createApiErrorMessage(
  statusCode: number,
  originalMessage?: string
): string {
  const errorMessages: Record<number, string> = {
    400: "Ungültige Anfrage. Bitte überprüfen Sie Ihre Eingaben.",
    401: "Sie sind nicht angemeldet. Bitte melden Sie sich erneut an.",
    403: "Sie haben keine Berechtigung für diese Aktion.",
    404: "Die angeforderte Ressource wurde nicht gefunden.",
    409: "Konflikt: Die Aktion konnte nicht ausgeführt werden.",
    422: "Die übermittelten Daten sind ungültig.",
    500: "Serverfehler. Bitte versuchen Sie es später erneut.",
    502: "Der Server ist temporär nicht erreichbar.",
    503: "Der Service ist temporär nicht verfügbar.",
  };

  return (
    errorMessages[statusCode] ||
    originalMessage ||
    "Ein unerwarteter Fehler ist aufgetreten."
  );
}

/**
 * Handles API errors consistently and throws a proper error
 */
export function handleApiError(error: any): never {
  let statusCode = 500;
  let message = "Ein unerwarteter Fehler ist aufgetreten.";

  // Parse different error formats from $fetch
  if (error?.response) {
    statusCode = error.response.status;
    message = error.response.data?.message || error.response.statusText;
  } else if (error?.statusCode) {
    statusCode = error.statusCode;
    message = error.message;
  } else if (error?.data) {
    statusCode = error.data.statusCode || 500;
    message = error.data.message || message;
  } else if (error?.message) {
    message = error.message;
  }

  const userMessage = createApiErrorMessage(statusCode, message);
  console.error("API Error:", { statusCode, originalMessage: message, error });

  throw new Error(userMessage);
}

/**
 * Validates player session and throws error if invalid
 */
export function validatePlayerSession(
  player: any
): asserts player is { id: string; name: string } {
  if (!player?.id) {
    throw new Error(
      "Spieler nicht in der Sitzung gefunden. Bitte melden Sie sich erneut an."
    );
  }
}

/**
 * Validates room ID input
 */
export function validateRoomId(roomId: string): void {
  if (!roomId || roomId.trim().length === 0) {
    throw new Error("Bitte geben Sie eine gültige Lobby-ID ein.");
  }

  if (roomId.trim().length < 3) {
    throw new Error("Die Lobby-ID muss mindestens 3 Zeichen lang sein.");
  }
}
