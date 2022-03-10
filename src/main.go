package main

import (
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
	var municipios []string

	args := os.Args[1:]

	// Check if all parameters have been informed
	for _, arg := range args {
		if arg == "--import" || arg == "--export" || arg == "--sistema" {
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
			}

			state = ""
		}
	}

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

	// Cria no novo web.xml
	war.ConverteXml(path_web_xml_import, path_web_xml_export, municipios)
}

func setState(state, newState string) (string, error) {
	if state == "" {
		return newState, nil
	}

	return state, fmt.Errorf("error, have been waiting %s arg, but receive %s", state, newState)
}
