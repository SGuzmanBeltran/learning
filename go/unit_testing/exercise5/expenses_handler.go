package exercise5

import (
	"sync"
	"unit_testing/common"

	"github.com/gofiber/fiber/v2"
)

type Expense struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type TypeExpense struct {
	Expenses []Expense
}

type ExpenseTracker struct {
	ExpensesTypes map[string]TypeExpense
	mu            sync.RWMutex
}

type ExpenseHandler struct {
	tracker *ExpenseTracker
}

func CreateNewExpenseHandler() *ExpenseHandler {
	return &ExpenseHandler{
		tracker: &ExpenseTracker{
			ExpensesTypes: make(map[string]TypeExpense),
		},
	}
}

func (eh *ExpenseHandler) GetExpensesHandler(c *fiber.Ctx) error {
	expenseName := c.Params("name") // Extract "name" from the URL path
	if expenseName == "" {
		return common.HandlerErrorResponse(c, fiber.ErrBadRequest.Code, "Missing expense name")
	}

	eh.tracker.mu.RLock()
	expense, ok := eh.tracker.ExpensesTypes[expenseName]
	eh.tracker.mu.RUnlock()
	if !ok {
		return common.HandlerErrorResponse(c, fiber.ErrBadRequest.Code, "The expense doesnt exist")
	}

	return c.Status(fiber.StatusOK).JSON(expense.Expenses)
}

func (eh *ExpenseHandler) AddExpenseHandler(c *fiber.Ctx) error {
	var newExpense Expense

	// Parse the JSON body into the Expense struct
	if err := c.BodyParser(&newExpense); err != nil {
		return common.HandlerErrorResponse(c, fiber.ErrBadRequest.Code, "Invalid request body")
	}

	if newExpense.Name == "" {
		return common.HandlerErrorResponse(c, fiber.ErrBadRequest.Code, "Expense name is required")
	}

	if newExpense.Value == 0 {
		return common.HandlerErrorResponse(c, fiber.ErrBadRequest.Code, "Expense name is required")
	}

	eh.tracker.mu.Lock()
	defer eh.tracker.mu.Unlock()

	expense, ok := eh.tracker.ExpensesTypes[newExpense.Name]
	if !ok {
		expense = TypeExpense{}
	}
	expense.Expenses = append(expense.Expenses, newExpense)
	eh.tracker.ExpensesTypes[newExpense.Name] = expense

	return c.Status(fiber.StatusCreated).JSON(&common.HandlerResponse{
		Message: "Expense created",
		Status:  fiber.StatusCreated,
	})

}
