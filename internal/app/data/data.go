package data

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"ports/internal/pkg/parser"
	"ports/internal/pkg/pb"
	"ports/internal/pkg/pb/client"
	"runtime"
	"syscall"

	"github.com/gofiber/fiber/v2"
)

func NewDataApp(portsFile string) DataApp {

	// TODO no time for handle it better :)
	client := client.PortsServiceClient{}
	client.Connect()

	return DataApp{
		Mux: fiber.New(fiber.Config{
			DisableStartupMessage: true,
		}),
		PortsFile: portsFile,
		Address:   ":9090",
		Client:    client,
		Ctx:       context.Background(),
	}

}

type DataApp struct {
	Mux       *fiber.App
	Address   string
	PortsFile string
	Client    client.PortsServiceClient
	Ctx       context.Context
}

// Start starts app
func (app *DataApp) Start() {
	app.SignalListener()
	app.ConfigureRoutes()

	log.Printf("Starting HTTP server on address %s", app.Address)
	if err := app.Mux.Listen(app.Address); err != nil {
		log.Panic(err)
	}

}

// ConfigureRoutes initializes routes for service
func (app DataApp) ConfigureRoutes() {
	app.Mux.Get("/ports", app.GetPortsHandler())
	app.Mux.Post("/ports", app.SavePortHandler())
}

// GetPortsHandler for handling ports data get
func (app DataApp) GetPortsHandler() fiber.Handler {
	// some init could be added here
	return func(c *fiber.Ctx) error {
		return nil
	}
}

// SavePortHandler runs json file scan and
// put each port entry into domain service.
func (app DataApp) SavePortHandler() fiber.Handler {
	// some handler init could be added here if needed
	return func(c *fiber.Ctx) error {
		input, err := os.Open(app.PortsFile)
		if err != nil {
			fmt.Printf("%+v\n", err)

			return app.Err(c, err)
		}
		err = parser.ProcessPortsJSON(input, func(port *pb.Port) error {
			log.Printf("Processing port with code: %+v", port)
			_, err := app.Client.Client.Insert(context.Background(), port)
			return err
		})

		if err != nil {
			log.Println("Got error during processing JSON file", err)
		}

		PrintMemUsage()

		return nil
	}
}

func (app DataApp) Err(c *fiber.Ctx, err error) error {
	c.SendStatus(500)
	return c.JSON(map[string]string{
		"error": err.Error(),
	})
}

// SignalListener for graceful shutdown
func (app DataApp) SignalListener() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		app.Mux.Shutdown()
		// some cleanup can be added here if needed
	}()
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
