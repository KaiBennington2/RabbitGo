package main

import (
	"fmt"
	"github.com/KaiBennington2/RabbitGo/internal/infra/app"
	"os"
)

func main() {
	if err := app.StartApplication(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error startup app: \n%s", err)
		os.Exit(1)
	}
}
