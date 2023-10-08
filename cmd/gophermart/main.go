package main

import (
	"AlekseyMartunov/internal/app"
	"context"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()

	err := app.StartApp(ctx)
	if err != nil {
		panic(err)
	}
}
