Correios
========

Integração com webservice dos correios feito em Go.

Instalação
----------

.. code:: golang
	
	go get github.com/tiaguinho/correios

Utilização
----------

Abaixo um exemplo de como utilizar o package dos correios

.. code:: golang

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

