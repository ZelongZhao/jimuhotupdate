package main

import (
	"github.com/google/wire"
)

func NewServices() (*services, error) {
	panic(wire.Build(
		wire.Struct(new(services), "*"),
		set,
	))
}
