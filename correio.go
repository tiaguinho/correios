package correios

const WEBSERVICE string = "http://ws.correios.com.br/calculador/CalcPrecoPrazo.asmx"

var (
	CodigoEmpresa string
	Senha         string
)

type Consulta struct {
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

func CalcPrecoPrazo() {

}

func CalcPreco() {

}

func CalcPrazo() {

}
