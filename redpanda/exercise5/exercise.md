# Exercise 5: Ride-Hailing App Simulation

**Objective:** Build a multi-producer, multi-consumer system with ride matching logic.

1. **Create topics:** `rides` and `drivers`.

## Producers:
- One producer simulates ride requests.
- Another simulates driver availability.

## Consumers:
- Create one consumer to match rides to available drivers.

**Challenge:** Implement a status update system where ride statuses are published to the `rides` topic (e.g., "In Transit", "Completed"). Ensure status updates trigger real-time processing.

## Plan
producers
 - ask for rides
 - announce availbility
 - update rides statuses

ride topic - 2 part - 1 consumer - 2 producers
 - riders
 - update

 drivers topic - 1 consumer - 1 producer

flow
 1. Driver availability
    - Create Driver Producer
    - Create Driver Consumer
    - Change Availability
 2. Ride request
    - Create Rider Producer
    - Create Rides Consumer
    - Create Ride
 3. Match
 4. Update
    - Create Update Producer
    - Create Update Consumer
    - Update Ride State
    - Change Driver Availabitily

## **Rating and Review**

---

1. **Kafka Integration**: **4.5 / 5**  
   - Well-implemented producers and consumers.  
   - Smooth message handling between drivers and riders.  
   - Minor improvements needed for better error handling (e.g., retries).  

2. **Concurrency Handling**: **4.5 / 5**  
   - Correct use of `sync.WaitGroup` and mutexes to prevent race conditions.  
   - Some room for improvement with error recovery in goroutines.  

3. **Code Structure & Readability**: **4 / 5**  
   - The code is well-organized, with clear separation of concerns.  
   - Slight redundancy in state updates (e.g., in `MatchRide`) could be reduced.  

4. **Logic and Matching Efficiency**: **4 / 5**  
   - Matching logic is solid but could benefit from using **maps** for quicker lookups.  
   - Efficient ride creation, but the flow could improve by decoupling more logic.  

5. **Error Handling**: **3 / 5**  
   - Logs errors but does not handle them proactively (e.g., retry mechanism).  
   - Adding retries and more granular error reporting would boost reliability.  

---

### **Final Average Rating**: **4 / 5** ⭐  

---

### **Summary**  
Great job! With minor improvements in **error handling** and **data structure optimizations**, this implementation can reach a 5/5 level. You're on the right track—keep up the excellent work! 🚀