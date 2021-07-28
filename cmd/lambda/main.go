package main

import (
	"context"
	"fmt"
)

type Handler struct{}

func (h *Handler) Invoke(context.Context, []byte) ([]byte, error) {
	fmt.Printf("Let's go!")
	return nil, nil
}

func main() {
	(&Handler{}).Invoke(context.TODO(), nil)
}
