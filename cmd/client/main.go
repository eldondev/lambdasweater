package main

import (
	"context"
	"fmt"
)

type Proxy struct{}

func (h *Handler) Invoke(context.Context, []byte) ([]byte, error) {
	fmt.Printf("Let's go!")
	return nil, nil
}

func main() {
	(&Proxy{}).Invoke(context.TODO(), nil)
}
