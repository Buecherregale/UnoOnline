import type {
  Card,
  CardColor,
  CardColorString,
  CardValueString,
} from "~/util/models";

/**
 * Maps numeric color codes to CardColor numbers
 */
export function mapColor(colorCode: CardColor): CardColorString {
  const colorMap: Record<number, CardColorString> = {
    0: "red",
    1: "green",
    2: "blue",
    3: "yellow",
    4: "black",
  };
  return colorMap[colorCode] || "red";
}

/**
 * Maps numeric value codes to CardValue
 */
export function mapValue(valueCode: number): CardValueString {
  if (valueCode >= 0 && valueCode <= 9) {
    return valueCode;
  }

  const valueMap: Record<number, CardValueString> = {
    10: "skip",
    11: "reverse",
    12: "plus2",
    13: "wild",
    14: "wildcard4",
  };
  return valueMap[valueCode] || 0;
}

/**
 * Parses payload to Card object
 */
export function parseCardPayload(payload: any): Card {
  return {
    color: payload.Color,
    value: payload.Value,
    chosen: payload.Chosen !== 0 ? payload.Chosen : null,
  };
}
