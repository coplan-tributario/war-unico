package warunico

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"strings"
)

type WebApp struct {
	XMLName        xml.Name      `xml:"web-app"`
	Text           string        `xml:",chardata"`
	Xmlns          string        `xml:"xmlns,attr"`
	Xsi            string        `xml:"xsi,attr"`
	SchemaLocation string        `xml:"schemaLocation,attr"`
	Version        string        `xml:"version,attr"`
	DisplayName    string        `xml:"display-name"`
	Description    string        `xml:"description"`
	Servlets       []Servlet     `xml:"servlet"`
	ServletsMap    []ServletMap  `xml:"servlet-mapping"`
	SessionConfig  SessionConfig `xml:"session-config"`
}

type Servlet struct {
	XMLName      xml.Name    `xml:"servlet"`
	ServletName  string      `xml:"servlet-name"`
	ServletClass string      `xml:"servlet-class,omitempty"`
	JspFile      string      `xml:"jsp-file,omitempty"`
	InitParams   []InitParam `xml:"init-param"`
}

type InitParam struct {
	XMLName    xml.Name `xml:"init-param"`
	ParamName  string   `xml:"param-name"`
	ParamValue string   `xml:"param-value"`
}

type ServletMap struct {
	XMLName     xml.Name `xml:"servlet-mapping"`
	ServletName string   `xml:"servlet-name"`
	UrlPattern  string   `xml:"url-pattern"`
}

type SessionConfig struct {
	XMLName        xml.Name `xml:"session-config"`
	SessionTimeout string   `xml:"session-timeout"`
}

func ConverteXml(path_web_xml_import string, path_web_xml_export string, municipios []string) {
	// Abre o arquivo XML
	xmlFile, err := os.Open(path_web_xml_import)
	if err != nil {
		panic(err)
	}

	// Mais tarde irá fechar este xml, antes vamos capturar o que precisamos
	defer xmlFile.Close()

	// Realiza a leitura do arquivo
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// Inicializa o XML
	var webApp WebApp

	// Converte o XML na estrutura especificada acima
	xml.Unmarshal(byteValue, &webApp)

	// Percorre os map e adiciona MAIS
	for _, servletMap := range webApp.ServletsMap {
		// Remove o servlet da URL
		urlPattern := strings.Replace(servletMap.UrlPattern, "/servlet/", "/", -1)

		for _, m := range municipios {
			sm := ServletMap{
				ServletName: servletMap.ServletName,
				UrlPattern:  "/" + m + urlPattern,
			}

			webApp.ServletsMap = append(webApp.ServletsMap, sm)
		}
	}

	// Cria dos servlets para redirecionar para a tela de login
	s := Servlet{
		ServletName: "page_redirect",
		JspFile:     "/redirect.html",
	}

	webApp.Servlets = append(webApp.Servlets, s)

	for _, m := range municipios {
		sm := ServletMap{
			ServletName: s.ServletName,
			UrlPattern:  "/" + m + "/",
		}

		webApp.ServletsMap = append(webApp.ServletsMap, sm)
	}

	s = Servlet{
		ServletName: "page_redirect_sem_barra",
		JspFile:     "/redirect_sem_barra.html",
	}

	webApp.Servlets = append(webApp.Servlets, s)

	for _, m := range municipios {
		sm := ServletMap{
			ServletName: s.ServletName,
			UrlPattern:  "/" + m,
		}

		webApp.ServletsMap = append(webApp.ServletsMap, sm)
	}

	file, _ := xml.MarshalIndent(webApp, "", " ")

	_ = ioutil.WriteFile(path_web_xml_export, file, 0644)
}
