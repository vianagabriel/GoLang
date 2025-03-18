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

	reservaHotel(ctx)
}

func reservaHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Reserva do hotel cancelada, Tempo Excedido")
		return
	case <-time.After(2 * time.Second):
		fmt.Println("Reserva concluida com sucesso!")
		return
	}
}
