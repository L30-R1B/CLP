package data

import (
	"clp-go-version/entidades"
	"fmt"
	"strings"
)

// DAO é uma estrutura genérica para manipular entidades.
// O DAO (Data Access Object) encapsula operações de armazenamento, busca e remoção
// de dados relacionados a entidades genéricas, reduzindo o acoplamento e centralizando
// a lógica de acesso a dados.
type DAO[E entidades.Entidade] struct {
	Dados []E // Armazena as entidades de tipo genérico.
}

// NewDAO cria uma nova instância de DAO.
// Essa função retorna um ponteiro para um DAO inicializado com um slice vazio.
// A sintaxe genérica [E entidades.Entidade] garante que o tipo `E` satisfaça a interface `Entidade`.
func NewDAO[E entidades.Entidade]() *DAO[E] {
	return &DAO[E]{}
}

// GetDados retorna a lista de entidades armazenadas.
// Essa função expõe os dados armazenados, permitindo que sejam acessados diretamente.
func (d *DAO[E]) GetDados() []E {
	return d.Dados
}

// Adicionar adiciona uma entidade ao DAO.
// O método utiliza `append` para adicionar uma nova entidade ao slice `Dados`.
func (d *DAO[E]) Adicionar(entidade E) {
	d.Dados = append(d.Dados, entidade)
}

// Buscar procura uma entidade pelo ID.
// Percorre o slice `Dados` para encontrar a entidade cujo ID seja igual ao fornecido.
// Retorna um ponteiro para a entidade encontrada ou `nil` caso não seja encontrada.
func (d *DAO[E]) Buscar(id int64) *E {
	for i := range d.Dados {
		// Chama o método `GetID` da interface `Entidade` para comparar os IDs.
		if d.Dados[i].GetID() == id {
			return &d.Dados[i] // Retorna um ponteiro para a entidade encontrada.
		}
	}
	return nil // Retorna nil se a entidade não for encontrada.
}

// Remover remove uma entidade pelo ID.
// Cria um novo slice filtrando as entidades cujo ID seja diferente do fornecido.
// Atualiza o slice `Dados` com os valores filtrados.
func (d *DAO[E]) Remover(id int64) {
	filtrados := []E{} // Slice temporário para armazenar as entidades que serão mantidas.
	for _, e := range d.Dados {
		if e.GetID() != id {
			filtrados = append(filtrados, e) // Adiciona apenas as entidades cujo ID não corresponde.
		}
	}
	d.Dados = filtrados // Atualiza o slice original com o novo slice filtrado.
}

// String retorna uma representação textual do DAO.
// Utiliza `strings.Builder` para concatenar as representações textuais de todas as entidades armazenadas.
// A implementação do método `String` da interface `Entidade` é chamada para cada entidade.
func (d *DAO[E]) String() string {
	var sb strings.Builder
	for _, e := range d.Dados {
		// Chama o método `String` de cada entidade para formatar o texto.
		sb.WriteString(fmt.Sprintf("\n%s", e.String()))
	}
	return sb.String() // Retorna a string concatenada.
}
