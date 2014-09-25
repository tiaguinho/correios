Correios
========

Integração com webservice dos correios feito em Go.

[![Build Status](https://travis-ci.org/tiaguinho/correios.svg?branch=master)](https://travis-ci.org/tiaguinho/correios)

[![GoDoc](https://godoc.org/github.com/tiaguinho/correios?status.png)](https://godoc.org/github.com/tiaguinho/correios)

Para entender melhor o funcionamento da integração com os correios, quais campos são obrigatórios e como eles devem ser preenchidos acesse o link a baixo:

<http://www.correios.com.br/para-voce/correios-de-a-a-z/pdf/calculador-remoto-de-precos-e-prazos/manual-de-implementacao-do-calculo-remoto-de-precos-e-prazos>

## Instalação ##

```bash
go get github.com/tiaguinho/correios
```
Struct da consulta
------------------

A estrutura abaixo mostra os campos que devem ser preenchidos para executar a requisição para o webservice dos correios.

```go
type Params struct {
	CodigoEmpresa    string  `url:"nCdEmpresa"`
	Senha            string  `url:"sDsSenha"`
	CodigoServico    string  `url:"nCdServico"`
	CepOrigem        string  `url:"sCepOrigem"`
	CepDestino       string  `url:"sCepDestino"`
	Peso             string  `url:"nVlPeso"`
	CodigoFormato    int     `url:"nCdFormato"`
	Comprimento      float64 `url:"nVlComprimento"`
	Altura           float64 `url:"nVlAltura"`
	Largura          float64 `url:"nVlLargura"`
	Diametro         float64 `url:"nVlDiametro"`
	MaoPropria       string  `url:"sCdMaoPropria"`
	ValorDeclarado   float64 `url:"nVlValorDeclarado"`
	AvisoRecebimento string  `url:"sCdAvisoRecebimento"`
}
```

## Utilização ##

Abaixo um exemplo de como utilizar o package dos correios

```go
package main

import (
	"fmt"
	"github.com/tiaguinho/correios"
)

func main() {
	params := correios.Params{
		CodigoServico: "40010,40045,40215,40290,41106",
		CepOrigem:     "05311900",
		CepDestino:    "86600280",
		Peso:          "300",
		CodigoFormato: 1,
		Comprimento:   20,
		Altura:        20,
		Largura:       20,
		Diametro:      0,
		MaoPropria:    "N",
	}

	results, _ := correios.CalcPrecoPrazo(params)
	for _, result := range results {
		fmt.Println("Código => ", result.Codigo)
		fmt.Println("Valor => ", result.Valor)
		fmt.Println("Prazo => ", result.Prazo)
		fmt.Println("Valor Mão Própria => ", result.ValorMaoPropria)
		fmt.Println("Valor Aviso Recebimento => ", result.ValorAvisoRecebimento)
		fmt.Println("Valor Declado => ", result.ValorDeclado)
		fmt.Println("Entrega Domiciliar => ", result.EntregaDomiciliar)
		fmt.Println("Entrega Sábado => ", result.EntregaSabado)
		fmt.Println("Erro => ", result.Erro)
		fmt.Println("MsgErro => ", result.MsgErro)

		fmt.Println("...")
	}
}
```
