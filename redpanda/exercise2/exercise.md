# Exercise 2: Real-Time Chat Simulation

**Objective:** Simulate a basic real-time chat between two users using multiple producers and one consumer.

1. **Create a topic** named `chat-room`.

## Producers:
- Create two producers simulating **Alice** and **Bob** sending messages to the chat.
- Use delays to make the chat feel natural.

## Consumer:
- Create one consumer that reads all messages from the chat.

**Challenge:** Format messages with `[user - timestamp]: message` and ensure the messages are processed in order.

# Code Review for Exercise 2

## What You Did Well

1. **Goroutines and Sync Handling**  
   - You’ve used goroutines effectively to run the producer and consumer concurrently. The `sync.WaitGroup` ensures that the main function waits for all tasks to complete, preventing premature termination.

2. **Message Struct and Metadata Parsing**  
   - The use of the `Message` struct simplifies message organization and parsing.
   - The function `getMessage` is clear and efficiently extracts metadata and content.

3. **Clear Log Output**  
   - Your logs are easy to follow, making debugging much simpler by showing the sequence of message production and consumption.

4. **Realistic Use Case**  
   - The conversation between "Alice" and "Bob" adds a practical and engaging touch to the example.

---

## Areas for Improvement

1. **Concurrency Issue with Multiple Producers**  
   - Since you are running two producers in separate goroutines, they might interfere with each other. Kafka guarantees message order per partition, but race conditions could occur with multiple producers.  
   **Suggestion**: Add logic to assign unique partitions or use a single producer for all messages. Alternatively, experiment with partition keys.

2. **Error Handling**  
   - While you handle errors when reading and writing messages, it would be helpful to **retry failed operations** or exit gracefully with detailed error messages.  
   **Suggestion**: Implement exponential backoff or retries when operations fail.

3. **Graceful Shutdown of Goroutines**  
   - The consumer runs indefinitely without a termination mechanism.  
   **Suggestion**: Use `context.WithCancel` to allow clean shutdown when the main function exits.

4. **Configuration Improvements**  
   - Hardcoding broker addresses and topics limits flexibility.  
   **Suggestion**: Load configuration from environment variables or files for adaptability.

5. **Use of Time in Test Data**  
   - The timestamps in your test data are far in the future, which feels unrealistic.  
   **Suggestion**: Use `time.Now().Unix()` for more natural timestamps.

## Overall Rating: ⭐️⭐️⭐️⭐️⭐️ (4.8/5)

This is a solid implementation of a Kafka producer-consumer system with concurrent handling and structured messages. With some small tweaks—like graceful shutdown, better error handling, and configuration management—you’ll be well-prepared for more advanced use cases.