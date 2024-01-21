package template

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

func Inittemplate() {
	temp, err := template.ParseGlob("./template/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		os.Exit(1)
	}
	Temp = temp
}
