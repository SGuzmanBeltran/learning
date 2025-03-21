# E-commerce Code Refactoring Exercise with Test-Driven Development

## Overview
This exercise provides a poorly designed e-commerce system with multiple SOLID principle violations.
Your task is to refactor the code using SOLID principles while applying Test-Driven Development.

## Requirements
1. Apply SOLID principles to refactor the code:
   - Single Responsibility Principle (SRP): Classes should have only one reason to change
   - Open/Closed Principle (OCP): Software entities should be open for extension but closed for modification
   - Liskov Substitution Principle (LSP): Subtypes must be substitutable for their base types
   - Interface Segregation Principle (ISP): Clients should not depend on interfaces they don't use
   - Dependency Inversion Principle (DIP): High-level modules should depend on abstractions, not low-level modules

2. Follow Test-Driven Development:
   - Write unit tests first, then implement the code to make them pass
   - Start with smaller, focused unit tests for individual components
   - Add integration tests to verify components work together
   - Aim for high test coverage of your refactored code

3. Testing Requirements:
   - Unit tests for each component/class
   - Integration tests for component interactions
   - Mock external dependencies (database, payment gateways, notification services)
   - Test edge cases and error handling
   - Use a testing framework like pytest

## Specific Areas to Improve
1. Break up the monolithic OnlineStore class into smaller, focused classes
2. Create proper abstractions and interfaces
3. Implement dependency injection
4. Apply appropriate design patterns (Factory, Strategy, Observer, etc.)
5. Improve the security of user management
6. Make the payment processing system extensible
7. Enhance error handling throughout the system

## Recommended Approach
1. Start by writing tests for one small component (e.g., product management)
2. Refactor that component to meet SOLID principles
3. Verify your tests pass
4. Repeat for other components
5. Add integration tests as you connect the refactored components

## Expected Outcome
Your refactored code should:
- Follow all SOLID principles
- Have comprehensive unit and integration tests
- Be more maintainable and extensible
- Have the same functionality as the original code