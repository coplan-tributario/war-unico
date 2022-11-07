package main

import (
	"bufio"
	"fmt"
	"os"

	sis "war-unico/src/sistemas"
	war "war-unico/src/war-unico"
)

func main() {
	var state string = ""
	var err error

	// Já define por padrão onde deve estar o web.xml e onde será gerado
	var path_web_xml_import string = "../import/web.xml"
	var path_web_xml_export string = "../export/web_new.xml"
	var sistema string = "tributario"
	var servlet string = "incluir"
	var path string = ""
	var municipios []string

	args := os.Args[1:]

	// Check if all parameters have been informed
	for _, arg := range args {
		if arg == "--import" || arg == "--export" || arg == "--sistema" || arg == "--servlet" || arg == "--path" {
			state, err = setState(state, arg)

			if err != nil {
				panic(err)
			}
		} else {
			if state == "--import" {
				path_web_xml_import = arg
			} else if state == "--export" {
				path_web_xml_export = arg
			} else if state == "--sistema" {
				sistema = arg
			} else if state == "--servlet" {
				servlet = arg
			} else if state == "--path" {
				path = arg
			}

			state = ""
		}
	}

	if path != "" {
		// Carrega a lista de municipios através do arquivo que foi passado no path
		municipios, err = loadFileContent(path)

		if err != nil {
			panic(err)
		}
	}

	if path == "" || len(municipios) == 0 {
		switch sistema {
		case "transparencia":
			municipios = sis.RetornaMunicipiosTransparencia()
		case "aplic":
			municipios = sis.RetornaMunicipiosAplic()
		case "central":
			municipios = sis.RetornaMunicipiosCentral()
		case "contabil":
			municipios = sis.RetornaMunicipiosContabil()
		case "planejamento":
			municipios = sis.RetornaMunicipiosPlanejamento()
		default:
			municipios = sis.RetornaMunicipiosTributario()
		}
	}

	// Cria no novo web.xml
	war.ConverteXml(servlet, path_web_xml_import, path_web_xml_export, municipios)
}

func setState(state, newState string) (string, error) {
	if state == "" {
		return newState, nil
	}

	return state, fmt.Errorf("error, have been waiting %s arg, but receive %s", state, newState)
}

func loadFileContent(path string) ([]string, error) {
	var municipios []string

	readFile, err := os.Open(path)

	if err != nil {
		return municipios, err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		municipios = append(municipios, line)
	}

	return municipios, nil
}
