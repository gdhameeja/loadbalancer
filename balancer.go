package main

import "github.com/gdhameea/backend"

type ServerPool struct {
	backends []*backend.Backend
	current  uint64
}
