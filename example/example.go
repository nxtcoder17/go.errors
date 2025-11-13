package main

import (
	"log/slog"

	"github.com/nxtcoder17/go.errors"
)

func A() error {
	return errors.New("failed to process A").KV("func", "A")
}

func B() error {
	if err := A(); err != nil {
		return errors.New("failed to process B").Wrap(err).KV("func", "B")
	}
	return nil
}

func main() {
	err := B()
	slog.Info(err.Error())
}
