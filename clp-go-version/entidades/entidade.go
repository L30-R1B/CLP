package entidades

// O pacote `entidades` é usado para agrupar tipos e funcionalidades relacionadas a entidades.
// Em Go, pacotes são o principal meio de organizar e reutilizar código.
// Cada arquivo no mesmo diretório, com o mesmo nome de pacote, pertence a esse pacote.

// Entidade define os métodos que uma entidade deve ter.
// Interfaces em Go não especificam como algo é implementado, apenas o que deve ser implementado.
// Isso permite que diferentes tipos satisfaçam uma interface, tornando o código mais flexível e reutilizável.
type Entidade interface {
	// GetID retorna o identificador único da entidade.
	// Em Go, funções de interfaces são declaradas com seus nomes, parâmetros e tipos de retorno.
	// Aqui, GetID deve ser implementado por qualquer tipo que satisfaça a interface Entidade.
	GetID() int64

	// String retorna uma representação textual da entidade.
	// Este método é semelhante ao método `toString()` em outras linguagens, mas em Go ele faz parte
	// da interface `fmt.Stringer` se implementado.
	String() string
}
