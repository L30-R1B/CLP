package main

import "sync"

type DAOProduto struct {
	instance *DAOProduto
	dao      *DAO[Produto]
	mutex    sync.Mutex
}

func GetProdutoInstance() *DAOProduto {
	var once sync.Once
	daoProduto := &DAOProduto{
		dao: NewDAO[Produto](),
	}

	once.Do(func() {
		daoProduto.instance = daoProduto
	})

	return daoProduto.instance
}

func (dao *DAOProduto) Adicionar(produto Produto) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	dao.dao.Adicionar(produto)
}

func (dao *DAOProduto) BuscarPorID(id int64) *Produto {
	return dao.dao.Buscar(func(p Produto) bool { return p.ID == id })
}

func (dao *DAOProduto) BuscarPorNome(nome string) *Produto {
	return dao.dao.Buscar(func(p Produto) bool { return p.Nome == nome })
}

func (dao *DAOProduto) RemoverPorID(id int64) {
	dao.dao.Remover(func(p Produto) bool { return p.ID == id })
}

func (dao *DAOProduto) RemoverPorNome(nome string) {
	dao.dao.Remover(func(p Produto) bool { return p.Nome == nome })
}

func (dao *DAOProduto) Listar() []Produto {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	return dao.dao.dados
}
