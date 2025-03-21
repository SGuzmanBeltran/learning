package exercise5

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

// Mock data
type MockExpense struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type ExpenseTests struct {
	name         string
	body         string
	wantedStatus int
}

func setupApp(handler *ExpenseHandler) *fiber.App {
	app := fiber.New()
	app.Get("/expenses/:name", handler.GetExpensesHandler)
	app.Post("/expenses", handler.AddExpenseHandler)
	return app
}

func TestGetExpensesHandler(t *testing.T) {
	handler := CreateNewExpenseHandler()
	app := setupApp(handler)

	// Add a mock expense
	handler.tracker.ExpensesTypes["Groceries"] = TypeExpense{
		Expenses: []Expense{
			{Name: "Groceries", Value: 100},
		},
	}

	req := httptest.NewRequest("GET", "/expenses/Groceries", nil)
	resp, err := app.Test(req)

	if err != nil {
		t.Fatalf("Test request failed: %v", err)
	}

	assert.Equal(t, fiber.StatusOK, resp.StatusCode)

	var expenses []Expense
	json.NewDecoder(resp.Body).Decode(&expenses)

	assert.Equal(t, 1, len(expenses))
	assert.Equal(t, "Groceries", expenses[0].Name)
	assert.Equal(t, 100, expenses[0].Value)
}

func TestAddExpenseHandler(t *testing.T) {
	handler := CreateNewExpenseHandler()
	app := setupApp(handler)

	tests := []ExpenseTests{
		{"Valid case", `{"name": "Groceries", "value": 50}`, fiber.StatusCreated},
		{"Invalid JSON", `{"name": "Groceries"}`, fiber.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/expenses", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")

			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Test request failed: %v", err)
			}

			assert.Equal(t, tt.wantedStatus, resp.StatusCode)
		})
	}
}

func BenchmarkGetExpensesHandler(b *testing.B) {
	handler := CreateNewExpenseHandler()
	app := setupApp(handler)

	// Add a mock expense
	handler.tracker.ExpensesTypes["Groceries"] = TypeExpense{
		Expenses: []Expense{
			{Name: "Groceries", Value: 100},
		},
	}

	req := httptest.NewRequest("GET", "/expenses/Groceries", nil)

	for i := 0; i < b.N; i++ {
		resp, err := app.Test(req)
		if err != nil {
			b.Fatalf("Test request failed: %v", err)
		}
		if resp.StatusCode != fiber.StatusOK {
			b.Errorf("Expected status 200, got %d", resp.StatusCode)
		}
	}
}

func BenchmarkAddExpenseHandler(b *testing.B) {
	handler := CreateNewExpenseHandler()
	app := setupApp(handler)

	body := `{"name": "Groceries", "value": 50}`
	req := httptest.NewRequest("POST", "/expenses", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	for i := 0; i < b.N; i++ {
		resp, err := app.Test(req)
		if err != nil {
			b.Fatalf("Test request failed: %v", err)
		}
		if resp.StatusCode != fiber.StatusCreated {
			b.Errorf("Expected status 201, got %d", resp.StatusCode)
		}
	}
}
