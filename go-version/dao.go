package main

import (
	"fmt"
	"sync"
)

type DAO[T any] struct {
	dados []T
	mutex sync.Mutex
}

func NewDAO[T any]() *DAO[T] {
	return &DAO[T]{dados: []T{}}
}

func (dao *DAO[T]) Adicionar(entidade T) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	dao.dados = append(dao.dados, entidade)
}

func (dao *DAO[T]) Remover(condicao func(T) bool) bool {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	for i, entidade := range dao.dados {
		if condicao(entidade) {
			dao.dados = append(dao.dados[:i], dao.dados[i+1:]...)
			return true
		}
	}
	return false
}

func (dao *DAO[T]) Buscar(condicao func(T) bool) *T {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	for i := range dao.dados {
		fmt.Printf("Verificando: %+v\n", dao.dados[i]) // Log para inspecionar os dados.
		if condicao(dao.dados[i]) {
			return &dao.dados[i]
		}
	}
	return nil
}

func (dao *DAO[T]) String(format func(T) string) string {
	result := ""
	for _, entidade := range dao.dados {
		result += format(entidade) + "\n"
	}
	return result
}
