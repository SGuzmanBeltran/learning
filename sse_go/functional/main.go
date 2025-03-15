package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/valyala/fasthttp"
)

func main() {
	// Fiber instance
	app := fiber.New()

	// CORS for external resources
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Cache-Control",
	}))

	app.Get("/sse", func(c *fiber.Ctx) error {
		c.Set("Content-Type", "text/event-stream")
		c.Set("Cache-Control", "no-cache")
		c.Set("Connection", "keep-alive")
		c.Set("Transfer-Encoding", "chunked")
	
		c.Status(fiber.StatusOK).Context().SetBodyStreamWriter(fasthttp.StreamWriter(func(w *bufio.Writer) {
			fmt.Println("WRITER")
			var i int
			for {
				i++
				eventType := "json_event" // Define your event type
				data := map[string]interface{}{
					"id":    i,
					"event": eventType,
					"time":  time.Now().Format(time.RFC3339),
					"message": fmt.Sprintf("This is message #%d", i),
				}
	
				// Serialize to JSON
				jsonData, err := json.Marshal(data)
				if err != nil {
					fmt.Printf("Error while marshaling JSON: %v\n", err)
					break
				}
	
				// Send event type and JSON data
				fmt.Fprintf(w, "event: %s\n", eventType)
				fmt.Fprintf(w, "data: %s\n\n", jsonData)
				fmt.Println("Sent:", eventType, string(jsonData))
	
				err = w.Flush()
				if err != nil {
					fmt.Printf("Error while flushing: %v. Closing http connection.\n", err)
					break
				}
				time.Sleep(2 * time.Second)
			}
		}))
	
		return nil
	})

	// Start server
	log.Fatal(app.Listen(":3000"))
}
