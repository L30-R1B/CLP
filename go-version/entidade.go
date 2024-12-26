package main

import (
	"fmt"
	"time"
)

type Entidade struct {
	ID int64
}

func NewEntidade() Entidade {
	return Entidade{
		ID: time.Now().UnixNano(),
	}
}

func (e Entidade) String() string {
	return fmt.Sprintf("ID: %d\t", e.ID)
}
