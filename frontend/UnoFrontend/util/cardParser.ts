import type { Card, CardColor, CardValue } from "~/util/models";

/**
 * Maps numeric color codes to CardColor strings
 */
export function mapColor(colorCode: number): CardColor {
  const colorMap: Record<number, CardColor> = {
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
export function mapValue(valueCode: number): CardValue {
  if (valueCode >= 0 && valueCode <= 9) {
    return valueCode;
  }

  const valueMap: Record<number, CardValue> = {
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
    color: mapColor(payload.Color),
    value: mapValue(payload.Value),
    chosen: payload.Chosen !== 0 ? mapColor(payload.Chosen) : null,
  };
}
