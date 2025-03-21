package exercise4

import (
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
)

type MockService struct {}

func (m *MockService) GetAllTodos() ([]*Todo, error) {
    return []*Todo{}, nil
}

func (m *MockService) CreateTodo(create CreateTodo) (*Todo, error) {
    return &Todo{}, nil
}

func (m *MockService) DeleteTodo(todo Todo) (*Todo, error) {
    return &Todo{}, nil
}

func TestHttpHandler_GetTodosHandler(t *testing.T) {
	app := fiber.New()
	mockS := &MockService{}
	HttpHandler := CreateNewHttpHandler(
		mockS,
	)
	app.Get("/todos", HttpHandler.GetTodosHandler)
	// Create a mock request with a valid JSON body
    req := httptest.NewRequest("GET", "/todos", strings.NewReader(""))
    req.Header.Set("Content-Type", "application/json")

	// Perform the request using Fiber's testing method
    resp, err := app.Test(req)
    if err != nil {
        t.Fatalf("Test request failed: %v", err)
    }

	if resp.StatusCode != 200 {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

type TodoTests struct {
	name string
	body string
	wantedStatus int
}

func TestHttpHandler_CreateTodoHandler(t *testing.T) {
	app := fiber.New()
	mockS := &MockService{}
	httpHandler := CreateNewHttpHandler(
		mockS,
	)
	app.Post("/todos", httpHandler.CreateTodoHandler)
	tests := []TodoTests{
		{"Normal case", `{"description": "Make bed", "completed": false}`, fiber.StatusCreated},
		{"No description", `{"description": "", "completed": true}`, fiber.StatusBadRequest},
		{"Bad request", `{"descrip": "lol"}`, fiber.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/todos", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")

			// Perform the request using Fiber's testing method
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Test request failed: %v", err)
			}

			if resp.StatusCode != tt.wantedStatus {
				t.Errorf("Expected status %d, got %d", tt.wantedStatus, resp.StatusCode)
			}
		})
	}
}

func TestHttpHandler_DeleteTodoHandler(t *testing.T) {
	app := fiber.New()
	mockS := &MockService{}
	httpHandler := CreateNewHttpHandler(
		mockS,
	)
	app.Delete("/todos", httpHandler.DeleteTodoHandler)
	tests := []TodoTests{
		{"Normal case", `{"id": 11, "description": "Make bed", "completed": false}`, fiber.StatusOK},
		{"Id Zero", `{"id": 0, "description": "", "completed": true}`, fiber.StatusBadRequest},
		{"Bad request, No id", `{"description": "Make bed", "completed": false}`, fiber.StatusBadRequest},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("DELETE", "/todos", strings.NewReader(tt.body))
			req.Header.Set("Content-Type", "application/json")

			// Perform the request using Fiber's testing method
			resp, err := app.Test(req)
			if err != nil {
				t.Fatalf("Test request failed: %v", err)
			}

			if resp.StatusCode != tt.wantedStatus {
				t.Errorf("Expected status %d, got %d", tt.wantedStatus, resp.StatusCode)
			}
		})
	}
}
