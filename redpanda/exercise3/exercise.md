# Exercise 3: Stock Prices Stream with Partitioning

**Objective:** Practice partitioning a topic and consuming specific streams.

1. **Create a topic** `stock-prices` with 3 partitions.

## Producer:
- Simulate real-time stock prices for several companies (e.g., AAPL, GOOGL, TSLA).
- Each price update should go to the appropriate partition.

## Consumers:
- Create consumers to listen only to specific company streams by subscribing to a specific partition.

**Challenge:** Introduce random delays and ensure each consumer processes only relevant messages.
