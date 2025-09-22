import type { Card, Player, message, uuid } from "~/util/models";

export const useGameMessages = () => {
  const createDrawCardMessage = (player: Player, messageId: uuid): message => ({
    type: "PlayerDrawsCardsPayload",
    payload: {
      player,
      amount: 1,
    },
    message_id: messageId,
    expects_reply: true,
  });

  const createPlayCardMessage = (
    player: Player,
    card: Card,
    messageId: uuid
  ): message => ({
    type: "AnswerCardPayload",
    payload: {
      player,
      card,
    },
    message_id: messageId,
    expects_reply: false,
  });

  return {
    createDrawCardMessage,
    createPlayCardMessage,
  };
};
