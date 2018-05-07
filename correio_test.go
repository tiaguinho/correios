package correios

import (
	"reflect"
	"testing"
)

var params = Params{
	CodigoServico: "40010",
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

//Testa a função de calculo de preço e prazo
func TestCalcPrecoPrazo(t *testing.T) {
	results, _ := CalcPrecoPrazo(params)

	expected := &Servico{
		Codigo:                "40010",
		Valor:                 "0,00",
		Prazo:                 "0",
		ValorMaoPropria:       "0,00",
		ValorAvisoRecebimento: "0,00",
		ValorDeclado:          "",
		EntregaDomiciliar:     "",
		EntregaSabado:         "",
		Erro:                  "-4",
		MsgErro:               "Peso excedido.",
	}

	if reflect.DeepEqual(expected, results[0]) == false {
		t.Error("Expected: ", expected, " - Got: ", results[0])
	}
}

//Testa a função de preço
func TestCalcPreco(t *testing.T) {
	results, _ := CalcPreco(params)

	expected := &Servico{
		Codigo:                "40010",
		Valor:                 "0,00",
		Prazo:                 "",
		ValorMaoPropria:       "0,00",
		ValorAvisoRecebimento: "0,00",
		ValorDeclado:          "",
		EntregaDomiciliar:     "",
		EntregaSabado:         "",
		Erro:                  "-4",
		MsgErro:               "Peso excedido.",
	}

	if reflect.DeepEqual(expected, results[0]) == false {
		t.Error("Expected: ", expected, " - Got: ", results[0])
	}
}

//Testa a função de prazo
func TestCalcPrazo(t *testing.T) {
	results, _ := CalcPrazo(params)

	expected := &Servico{
		Codigo:                "40010",
		Valor:                 "",
		Prazo:                 "5",
		ValorMaoPropria:       "",
		ValorAvisoRecebimento: "",
		ValorDeclado:          "",
		EntregaDomiciliar:     "S",
		EntregaSabado:         "N",
		Erro:                  "",
		MsgErro:               "",
	}

	if reflect.DeepEqual(expected, results[0]) == false {
		t.Error("Expected: ", expected, " - Got: ", results[0])
	}
}
