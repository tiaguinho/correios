package correios

import (
	"encoding/xml"
	"fmt"
	"github.com/google/go-querystring/query"
	"io/ioutil"
	"net/http"
)

const WEBSERVICE string = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx"

//Struct com os parametros da requisição
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

//Struct com os dados do retorno de cada serviço
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
	Resultado xml.Name  `xml:"cResultado"`
	Servico   []Servico `xml:"Servicos>cServico"`
}

//Calcula o preço e o prazo de entrega do item informado
func CalcPrecoPrazo(consulta Params) ([]Servico, error) {
	return doRequest("CalcPrecoPrazo", createQuery(consulta))
}

//Calcla o preço da entrega
func CalcPreco(consulta Params) ([]Servico, error) {
	return doRequest("CalcPreco", createQuery(consulta))
}

//Calcula somente o prazo de entrega
func CalcPrazo(consulta Params) ([]Servico, error) {
	return doRequest("CalcPrazo", createQuery(consulta))
}

//Cria a query de consulta a partir da interface informada
func createQuery(consulta Params) string {
	query_string, err := query.Values(consulta)
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	return query_string.Encode()
}

//Faz a request para o webservice dos correios
func doRequest(path, query_string string) ([]Servico, error) {
	resp, err := http.Get(WEBSERVICE + "/" + path + "?" + query_string)
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error: %v", err)
	}

	var results Servicos
	err = xml.Unmarshal(body, &results)

	if err != nil {
		fmt.Println("Error: %v", err)
	}

	return results.Servico, err
}
