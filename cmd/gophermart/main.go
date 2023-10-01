package main

import (
	"context"

	"AlekseyMartunov/internal/app"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()

	err := app.StartApp(ctx)
	if err != nil {
		panic(err)
	}
}
