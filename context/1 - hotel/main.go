package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Hotel booking canceled due the timeout!")
		return
	case <-time.After(2 * time.Second):
		fmt.Println("Hotel booked!")
	}
}
