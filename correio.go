package correios

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const WEBSERVICE string = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx"

type Params struct {
	CodigoEmpresa    string  `xml:"nCdEmpresa"`
	Senha            string  `xml:"sDsSenha"`
	CodigoServico    string  `xml:"nCdServico"`
	CepOrigem        string  `xml:"sCepOrigem"`
	CepDestino       string  `xml:"sCepDestino"`
	Peso             string  `xml:"nVlPeso"`
	CodigoFormato    int     `xml:"nCdFormato"`
	Comprimento      float64 `xml:"nVlComprimento"`
	Altura           float64 `xml:"nVlAltura"`
	Largura          float64 `xml:"nVlLargura"`
	Diametro         float64 `xml:"nVlDiametro"`
	MaoPropria       string  `xml:"sCdMaoPropria"`
	ValorDeclarado   float64 `xml:"nVlValorDeclarado"`
	AvisoRecebimento string  `xml:"sCdAvisoRecebimento"`
}

type Servico struct {
	Codigo                string `xml:"Codigo"`
	Valor                 string `xml:"Valor"`
	Prazo                 int    `xml:"PrazoEntrega"`
	ValorMaoPropria       string `xml:"ValorMaoPropria"`
	ValorAvisoRecebimento string `xml:"ValorAvisoRecebimento"`
	ValorDeclado          string `xml:"ValorValorDeclado"`
	EntregaDomiciliar     string `xml:"EntregaDomiciliar"`
	EntregaSabado         string `xml:"EntregaSabado"`
	Erro                  int    `xml:"Erro"`
	MsgErro               string `xml:"MsgErro"`
}

type Servicos struct {
	cServico []Servico `xml:"cServico"`
}

//Calcula o preço e o prazo de entrega do item informado na interface
func CalcPrecoPrazo(consulta Params) (Servicos, error) {
	return doRequest("CalcPrecoPrazo", createQuery(consulta))
}

//Calcla o preço da entrega
func CalcPreco(consulta Params) (Servicos, error) {
	return doRequest("CalcPreco", createQuery(consulta))
}

//Calcula somente o prazo de entrega
func CalcPrazo(consulta Params) (Servicos, error) {
	return doRequest("CalcPrazo", createQuery(consulta))
}

//Cria a query de consulta a partir da interface informada
func createQuery(consulta Params) string {
	var query []string
	for key, value := range consulta {
		query[len(query)] = key + "=" + string(value)
	}

	return strings.Join(query, "&")
}

//Faz a request para o webservice dos correios
func doRequest(path, query string) (Servicos, error) {
	resp, err := http.Get(WEBSERVICE + "/" + path + "?" + query)
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	var results Servicos
	err = xml.Unmarshal([]byte(body), &results)

	if err != nil {
		fmt.Println("Error: %v", err)
	}

	return results, err
}
