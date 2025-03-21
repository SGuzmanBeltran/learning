# Exercise 3: Stock Prices Stream with Partitioning

**Objective:** Practice partitioning a topic and consuming specific streams.

1. **Create a topic** `stock-prices` with 3 partitions.

## Producer:
- Simulate real-time stock prices for several companies (e.g., AAPL, GOOGL, TSLA).
- Each price update should go to the appropriate partition.

## Consumers:
- Create consumers to listen only to specific company streams by subscribing to a specific partition.

**Challenge:** Introduce random delays and ensure each consumer processes only relevant messages.

### Review of Exercise 3 🚀✨

Your implementation of Exercise 3 shows a solid understanding of Kafka concepts and effective coding practices. Here’s the breakdown:

### **Code Structure and Logic** ⭐⭐⭐⭐
- **Separation of Concerns**: Great modularization into main, producer, and consumer files.
- **Use of Goroutines**: Effectively handles concurrency with producers and consumers.
- **Partition Management**: Smart use of a custom partitioner for directing messages.
- **Logging**: Clear logging but consider structured logging for better insights.
- **Error Handling**: Good error checks in critical functions.

### **Functionality** ⭐⭐⭐⭐
- **Message Processing**: Correctly filters and processes relevant messages.
- **Topic Creation**: Proper error handling, though consider checking for existing topics.

### **Random Delays** 🕒
- To address the challenge, introduce random sleep durations in the `Consume` function and ensure processing only relevant messages.

### **Suggestions for Improvement** 🔧
- **Config Management**: Consider externalizing configuration settings.
- **Code Documentation**: Add comments for clarity on function purposes.
- **Graceful Shutdown**: Implement mechanisms for clean exits of consumers and producers.

### **Rating** 🌟
**4.5/5**  
You’ve demonstrated strong Kafka skills and effective coding practices. A few minor improvements could enhance the robustness and clarity of your code. Great job overall! Let me know if you have any questions or if you'd like to explore further concepts!