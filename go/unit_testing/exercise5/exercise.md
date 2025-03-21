# Exercise
Concurrent Data Structure
Concepts: Race condition detection, benchmarking

Task: Implement a thread-safe cache with Get and Set methods. Write unit tests to ensure thread safety, and create benchmarks to measure performance under different loads.

# Review

Your implementation for the thread-safe cache with Get and Set methods is quite good! Here’s a detailed review and rating of your code, including its strengths and areas for improvement. ⭐️

Review
Strengths:
Thread Safety:

You correctly used sync.RWMutex for read-write locks, ensuring safe concurrent access to the ExpensesTypes map. This is crucial for avoiding race conditions.
Clear Structure:

The organization of your code is logical. Separating the handler and the tracker makes it easier to manage and test.
Error Handling:

You handle various error cases in the HTTP handlers, providing meaningful messages when the input is invalid or when expenses are not found.
Unit Tests:

Your test cases cover different scenarios, including valid and invalid requests. This is great for ensuring the functionality of your handlers.
Benchmarking:

You included benchmark tests for both Get and Add methods, which is essential for understanding the performance of your implementation under load.
Areas for Improvement:
Expense Value Check:

In AddExpenseHandler, you're checking if the expense value is zero to return an error message. Instead, you might want to check if the value is negative, as zero could be a valid expense (e.g., a refund).
go
Copy code
if newExpense.Value < 0 {
    return common.HandlerErrorResponse(c, fiber.ErrBadRequest.Code, "Expense value must be non-negative")
}
Code Duplication:

In both handlers, you're directly accessing the ExpensesTypes map. Consider creating helper methods in ExpenseTracker for AddExpense and GetExpense, which would encapsulate the logic and reduce duplication.
Error Responses Consistency:

Ensure your error responses have consistent structures for easier handling on the client side.
Testing Edge Cases:

You could add more test cases, such as:
Trying to retrieve expenses for a non-existent name.
Adding an expense with a negative value.
Concurrent access tests to ensure that multiple goroutines can safely read/write at the same time.
Use of JSON Tags:

Your Expense struct has JSON tags, but consider making sure your naming conventions are consistent with common practices (e.g., using camelCase for JSON keys).
Overall Rating: ⭐️⭐️⭐️⭐️ (4/5)
Your implementation is solid and shows a good understanding of concurrent programming in Go! With a few improvements, especially in error handling and code reuse, it can be even better. Keep up the great work! 🚀