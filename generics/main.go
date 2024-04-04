package main

import (
	"context"
	"fmt"
	_case "myapp/generics/case"
	"os"
	"os/signal"
)

func main() {
	_case.SimpleCase()
	fmt.Println("======================")
	_case.CustomNumberTCase()
	fmt.Println("======================")
	_case.BuiltInCase()
	fmt.Println("======================")
	_case.TTypeCase()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer stop()
	<-ctx.Done()
}
