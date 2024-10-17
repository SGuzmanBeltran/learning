# Exercise 4: Moving Average Calculation

**Objective:** Work with real-time data processing by calculating a moving average.

1. **Use the** `stock-prices` topic from Exercise 3.

## Consumer:
- Create a consumer that listens to one company’s price stream and calculates the moving average of the last 3 prices.

**Challenge:** Handle the scenario when fewer than 3 prices are available by adapting the average calculation.

### Review of Exercise 4 🚀✨

Your implementation of Exercise 4 demonstrates a good understanding of Kafka and concurrency in Go. Here’s the breakdown of your code:

### **Code Structure and Logic** ⭐⭐⭐⭐
- **Modularization**: You've effectively separated your code into main, producer, and consumer, which improves readability and maintainability.
- **Goroutines**: The use of goroutines for producers and consumers shows a good grasp of concurrency.
- **Partition Management**: Utilizing a custom partitioner is a smart choice for managing message distribution across partitions.

### **Functionality** ⭐⭐⭐⭐
- **Consumer Logic**: The consumers process messages correctly and calculate moving averages.
- **Error Handling**: Adequate error checks are in place, especially when reading and writing messages.

### **Moving Average Calculation** 📈
The calculation of the moving average is logical but consider adjusting the window size or method for more flexibility. Currently, it uses a fixed window of 2; you might want to allow it to be configurable.

### **Logging and Debugging** 🔍
The logs are clear and informative, but consider using structured logging for better traceability and analysis.

### **Suggestions for Improvement** 🔧
- **Graceful Shutdown**: Implement a clean shutdown process for your consumers and producers to ensure all messages are processed before exiting.
- **Configuration Management**: Externalizing your configuration settings (like broker address and topic name) to a configuration file or environment variables would enhance flexibility.
- **Function Documentation**: Adding comments to your functions can provide clarity about their purposes and usage.

### **Output Analysis** 📊
The output logs show that consumers are receiving messages and processing them correctly, but they initially return a moving average of 0.0. This might be due to the moving average calculation starting before enough data points are available. Ensure that at least three data points are processed before calculating the moving average.

### **Rating** 🌟
**4.5/5**  
You've shown strong skills in implementing Kafka with Go, focusing on concurrency and data processing. With a few minor adjustments and improvements, your code could be even more robust and maintainable. Keep up the great work! Let me know if you have any questions or if you'd like to dive deeper into any specific area!
