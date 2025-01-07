package entidades

// Totalizavel define um comportamento para calcular o total.
// Essa interface especifica que qualquer tipo que a implemente deve fornecer
// um método chamado `Total`, que retorna o valor total como um float64.
type Totalizavel interface {
	Total() float64 // Método obrigatório que calcula e retorna o total.
}
