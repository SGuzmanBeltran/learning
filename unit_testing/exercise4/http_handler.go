package exercise4

import "github.com/gofiber/fiber/v2"

type HandlerResponse struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func HandlerErrorResponse(c *fiber.Ctx, code int, message string) error {
	r := &HandlerResponse{
		Message: message,
		Status:  code,
	}
	return c.Status(r.Status).JSON(r)
}

type CreateTodo struct{
	Id int `json:"id"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}

type Todo struct{
	Id int `json:"id"`
	Description string `json:"description"`
	Completed bool `json:"completed"`
}

type Service interface {
	GetAllTodos() ([]*Todo, error)
	CreateTodo(CreateTodo) (*Todo, error)
	DeleteTodo(Todo) (*Todo, error)
}

type HttpHandler struct {
	service Service
}

func CreateNewHttpHandler(service Service) *HttpHandler {
	return &HttpHandler{
		service: service,
	}
}

func (hh *HttpHandler) GetTodosHandler(c *fiber.Ctx) error {
	todos, err := hh.service.GetAllTodos()

	if err != nil {
		return HandlerErrorResponse(c, fiber.StatusInternalServerError, "Error fetching the todos")
	}

	return c.Status(fiber.StatusOK).JSON(todos)
}

func (hh *HttpHandler) CreateTodoHandler(c *fiber.Ctx) error {
	var todo CreateTodo
	if err := c.BodyParser(&todo); err != nil {
		return HandlerErrorResponse(c, fiber.StatusBadRequest, "Error decoding payload")
	}

	if todo.Description == "" {
		return HandlerErrorResponse(c, fiber.StatusBadRequest, "Description cant be empty")
	}

	newTodo, err := hh.service.CreateTodo(todo)

	if err != nil {
		return HandlerErrorResponse(c, fiber.StatusInternalServerError, "Error creating todo")
	}

	return c.Status(fiber.StatusCreated).JSON(newTodo)
}

func (hh *HttpHandler) DeleteTodoHandler(c *fiber.Ctx) error {
	var todo Todo
	if err := c.BodyParser(&todo); err != nil {
		return HandlerErrorResponse(c, fiber.StatusBadRequest, "Error decoding payload")
	}

	if todo.Id == 0 {
		return HandlerErrorResponse(c, fiber.StatusBadRequest, "Id cant be 0 or negative")
	}

	deletedTodo, err := hh.service.DeleteTodo(todo)

	if err != nil {
		return HandlerErrorResponse(c, fiber.StatusInternalServerError, "Error fetching the todos")
	}

	return c.Status(fiber.StatusOK).JSON(deletedTodo)
}
