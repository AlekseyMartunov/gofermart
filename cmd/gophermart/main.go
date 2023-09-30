package main

import (
	"context"

	"AlekseyMartunov/internal/app"
)

func main() {
	ctx := context.Background()
	defer ctx.Done()

	app.StartApp(ctx)
}
