package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()

	fmt.Println("context: \t", ctx)
	fmt.Println("context error: \t", ctx.Err())
	fmt.Printf("context type: \t%T\n", ctx)
	fmt.Println("context: \t", ctx)
	
	ctx, cancel :=  context.WithCancel(ctx)

	fmt.Println("context: \t", ctx)
	fmt.Println("context error: \t", ctx.Err())
	fmt.Printf("context type: \t%T\n", ctx)
	fmt.Println("context: \t", ctx)

	cancel()

	fmt.Println("context: \t", ctx)
	fmt.Println("context error: \t", ctx.Err())
	fmt.Printf("context type: \t%T\n", ctx)
	fmt.Println("context: \t", ctx)
}
