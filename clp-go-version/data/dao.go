package data

import (
	"clp-go-version/entidades"
	"fmt"
	"strings"
)

// DAO é uma estrutura genérica para manipular entidades.
type DAO[E entidades.Entidade] struct {
	Dados []E
}

// NewDAO cria uma nova instância de DAO.
func NewDAO[E entidades.Entidade]() *DAO[E] {
	return &DAO[E]{}
}

// GetDados retorna a lista de entidades armazenadas.
func (d *DAO[E]) GetDados() []E {
	return d.Dados
}

// Adicionar adiciona uma entidade ao DAO.
func (d *DAO[E]) Adicionar(entidade E) {
	d.Dados = append(d.Dados, entidade)
}

// Buscar procura uma entidade pelo ID.
func (d *DAO[E]) Buscar(id int64) *E {
	for i := range d.Dados {
		if d.Dados[i].GetID() == id {
			return &d.Dados[i]
		}
	}
	return nil
}

// Remover remove uma entidade pelo ID.
func (d *DAO[E]) Remover(id int64) {
	filtrados := []E{}
	for _, e := range d.Dados {
		if e.GetID() != id {
			filtrados = append(filtrados, e)
		}
	}
	d.Dados = filtrados
}

// String retorna uma representação textual do DAO.
func (d *DAO[E]) String() string {
	var sb strings.Builder
	for _, e := range d.Dados {
		sb.WriteString(fmt.Sprintf("\n%s", e.String()))
	}
	return sb.String()
}
