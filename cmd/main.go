package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/math-88/medic-online/internal/app"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	signal := <-sigChan
	log.Printf("shutting down bikes server with signal: %s", signal)
}
