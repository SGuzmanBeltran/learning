# Exercise 2: Real-Time Chat Simulation

**Objective:** Simulate a basic real-time chat between two users using multiple producers and one consumer.

1. **Create a topic** named `chat-room`.

## Producers:
- Create two producers simulating **Alice** and **Bob** sending messages to the chat.
- Use delays to make the chat feel natural.

## Consumer:
- Create one consumer that reads all messages from the chat.

**Challenge:** Format messages with `[user - timestamp]: message` and ensure the messages are processed in order.
