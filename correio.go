package correios

import (
	"encoding/xml"
	"fmt"
	"github.com/google/go-querystring/query"
	"net/http"
)

// WEBSERVICE é o link do webservice dos correios
const WEBSERVICE string = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx"

// Params é a struct com os parametros da requisição
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

// Servico e a struct com os dados do retorno de cada serviço
type Servico struct {
	Codigo                string `xml:"Codigo"`
	Valor                 string `xml:"Valor"`
	Prazo                 string `xml:"PrazoEntrega"`
	ValorMaoPropria       string `xml:"ValorMaoPropria"`
	ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento"`
	ValorDeclado          string `xml:"ValorValorDeclado"`
	EntregaDomiciliar     string `xml:"EntregaDomiciliar"`
	EntregaSabado         string `xml:"EntregaSabado"`
	Erro                  string `xml:"Erro"`
	MsgErro               string `xml:"MsgErro"`
}

// Servicos guarda a response da consulta
type Servicos struct {
	Resultado xml.Name   `xml:"cResultado"`
	Servico   []*Servico `xml:"Servicos>cServico"`
}

// CalcPrecoPrazo calcula o preço e o prazo de entrega do item informado
func CalcPrecoPrazo(consulta Params) ([]*Servico, error) {
	return doRequest("CalcPrecoPrazo", createQuery(consulta))
}

// CalcPreco calcula o preço da entrega
func CalcPreco(consulta Params) ([]*Servico, error) {
	return doRequest("CalcPreco", createQuery(consulta))
}

// CalcPrazo calcula somente o prazo de entrega
func CalcPrazo(consulta Params) ([]*Servico, error) {
	return doRequest("CalcPrazo", createQuery(consulta))
}

// createQuery cria a query de consulta a partir da interface informada
func createQuery(consulta Params) string {
	qs, _ := query.Values(consulta)

	return qs.Encode()
}

// doRequest faz a request para o webservice dos correios
func doRequest(path, qs string) ([]*Servico, error) {
	resp, err := http.Get(WEBSERVICE + "/" + path + "?" + qs)
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)
	}
	defer resp.Body.Close()

	var results Servicos
	decode := xml.NewDecoder(resp.Body)
	err = decode.Decode(&results)
	if err != nil {
		return nil, fmt.Errorf("Error: %v\n", err)
	}

	return results.Servico, err
}
