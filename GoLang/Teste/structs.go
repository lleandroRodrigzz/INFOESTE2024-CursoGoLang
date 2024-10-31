package main

import "time"

type EnderecoEntrega struct{
	Rua    string `json:"rua"`
	Numero int    `json:"numero"`
	Cidade string `json:"cidade"`
	Estado string `json:"estado"`
	Cep    string `json:"cep"`
}

type DetalhesPagamento struct{
	Bandeira       string `json:"bandeira"`
	NumeroCartao   string `json:"numero_cartao"`
	AutorizacaoID  string `json:"autorizacao_id"`
}

type Itens struct{
	ItemID        string  `json:"item_id"`
	Descricao     string  `json:"descricao"`
	Quantidade    int     `json:"quantidade"`
	PrecoUnitario float32 `json:"preco_unitario"`
}

type Cliente struct{
	ClienteID string `json:"cliente_id"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
}

type Transacao struct{
	TransacaoID			string				`json:"transacao_id"`
	DataHora			time.Time			`json:"data_hora"`
	Cliente				Cliente				`json:"cliente" yaml:"-"`
	Itens				[]Itens				`json:"itens"`
	ValorTotal			float64				`json:"valor_total"`
	MetodoPagamento		string				`json:"metodo_pagamento"`
	StatusPagamento		string				`json:"status_pagamento"`
	DetalhesPagamento	DetalhesPagamento	`json:"detalhes_pagamento"`
	EnderecoEntrega		EnderecoEntrega		`json:"endereco_entrega" yaml:"-"`
}