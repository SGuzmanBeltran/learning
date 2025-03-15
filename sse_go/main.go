package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"
)

// Message represents the structure of our SSE message
type Message struct {
	ID      int         `json:"id"`
	Event   string      `json:"event"`
	Time    string      `json:"time"`
	Message interface{} `json:"message"`
}

var messageChannel = make(chan Message)
var clients = struct {
	sync.RWMutex
	channels map[chan Message]bool
}{
	channels: make(map[chan Message]bool),
}

// Broadcast sends a message to all connected clients
func broadcast(msg Message) {
	clients.RLock()
	defer clients.RUnlock()

	for clientChan := range clients.channels {
		select {
		case clientChan <- msg:
		default:
			// If channel is blocked, skip this client
		}
	}
}

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: true,
	})

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://127.0.0.1:5500",
		AllowHeaders:     "Origin, Content-Type, Accept, Cache-Control",
		AllowMethods:     "GET, POST",
		ExposeHeaders:    "Content-Type,Content-Length,Cache-Control",
		AllowCredentials: true,
	}))

	app.Get("/sse", func(c *fiber.Ctx) error {
		// Set SSE headers
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Access-Control-Allow-Origin", "*")

		// Create a channel for this client
		clientChan := make(chan Message, 1)

		// Register this client
		clients.Lock()
		clients.channels[clientChan] = true
		clients.Unlock()

		// Cleanup function
		cleanup := func() {
			clients.Lock()
			delete(clients.channels, clientChan)
			clients.Unlock()
			close(clientChan)
		}

		// Set up the stream writer
		c.Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			// Send initial connection message
			fmt.Fprintf(w, "event: connected\ndata: Connection established\n\n")
			err := w.Flush()
			if err != nil {
				cleanup()
				return
			}

			for {
				select {
				case msg, ok := <-clientChan:
					if !ok {
						return
					}

					jsonData, err := json.Marshal(msg)
					if err != nil {
						fmt.Printf("Error marshaling JSON: %v\n", err)
						continue
					}

					// Send the event
					fmt.Fprintf(w, "event: %s\n", msg.Event)
					fmt.Fprintf(w, "data: %s\n\n", jsonData)

					err = w.Flush()
					if err != nil {
						cleanup()
						return
					}
				}
			}
		}))

		return nil
	})

	app.Post("/send", func(c *fiber.Ctx) error {
		var msg Message
		if err := c.BodyParser(&msg); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		msg.Time = time.Now().Format(time.RFC3339)
		broadcast(msg)

		return c.JSON(fiber.Map{
			"status": "Message sent",
		})
	})

	fmt.Println("Server starting on :3000")
	log.Fatal(app.Listen(":3000"))
}
