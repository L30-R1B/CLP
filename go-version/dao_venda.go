package main

import "sync"

type DAOVenda struct {
	instance *DAOVenda
	dao      *DAO[Venda]
	mutex    sync.Mutex
}

func GetVendaInstance() *DAOVenda {
	var once sync.Once
	daoVenda := &DAOVenda{
		dao: NewDAO[Venda](),
	}

	once.Do(func() {
		daoVenda.instance = daoVenda
	})

	return daoVenda.instance
}

func (dao *DAOVenda) Adicionar(venda Venda) {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	dao.dao.Adicionar(venda)
}

func (dao *DAOVenda) BuscarPorID(id int64) *Venda {
	return dao.dao.Buscar(func(v Venda) bool { return v.ID == id })
}

func (dao *DAOVenda) RemoverPorID(id int64) {
	dao.dao.Remover(func(v Venda) bool { return v.ID == id })
}

func (dao *DAOVenda) Listar() []Venda {
	dao.mutex.Lock()
	defer dao.mutex.Unlock()
	return dao.dao.dados
}
