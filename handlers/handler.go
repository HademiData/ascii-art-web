package handlers

import (
	asciigenerator "ascii-art-web/ascii-generator"
	"fmt"
	"net/http"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not Found", http.StatusNotFound)
		return

	}
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "500 Internal Server Error \ntemplate Error", http.StatusInternalServerError)
		fmt.Println(err)
		return

	}
	tmpl.Execute(w, nil)

}

func AsciiHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "400: Bad Request", http.StatusBadRequest)
		return
	}

	text := r.FormValue("inputString")
	banner := r.FormValue("banner")

	if text == "" {
		http.Error(w, "500: Empty text", http.StatusBadRequest)
		fmt.Println(" Empty text field...")
		return

	}

	charMap, err := asciigenerator.ParseBanner(banner)
	if err != nil {
		http.Error(w, "500 Internal Error:", http.StatusInternalServerError)
		fmt.Println("Error Parsing Banner ...")
		return

	}

	result := asciigenerator.PrintBannertoArt(text, charMap)

	templ, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "internal error:", http.StatusInternalServerError)
		fmt.Println("Error Parsinng template.. ")
		return

	}

	data := struct {
		Result      string
		InputString string
	}{
		Result:      result,
		InputString: text,
	}

	templ.Execute(w, data)

}
