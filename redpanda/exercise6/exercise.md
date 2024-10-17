# Exercise 6: Fault-Tolerant System

**Objective:** Handle failures and ensure your consumers can recover from disconnections or downtime.

## Task:
- Modify your previous consumers (from Exercises 3-5) to include logic that properly handles rebalancing and reconnects in case of failures.

**Goal:** Test your consumers by intentionally stopping and restarting them to confirm they resume processing without missing messages.

**Challenge:** Implement a "retry mechanism" for failed message processing.
