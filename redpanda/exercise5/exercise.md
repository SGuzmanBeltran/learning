# Exercise 5: Ride-Hailing App Simulation

**Objective:** Build a multi-producer, multi-consumer system with ride matching logic.

1. **Create topics:** `rides` and `drivers`.

## Producers:
- One producer simulates ride requests.
- Another simulates driver availability.

## Consumers:
- Create one consumer to match rides to available drivers.

**Challenge:** Implement a status update system where ride statuses are published to the `rides` topic (e.g., "In Transit", "Completed"). Ensure status updates trigger real-time processing.
