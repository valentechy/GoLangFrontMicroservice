package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const backendURL = "http://localhost:8081"

func main() {
	http.HandleFunc("/primario", handleColor("primario"))
	http.HandleFunc("/secundario", handleColor("secundario"))

	fmt.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleColor(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		color, err := fetchColor(path)
		if err != nil {
			println(err)
			println("Entro aquí")
			http.Error(w, "Error fetching color", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.New("color").Parse(`
            <!DOCTYPE html>
            <html>
            <head>
                <title>Color Page</title>
            </head>
            <body style="background-color: {{.Color}};">
                <h1>El color de fondo es {{.Color}}</h1>
            </body>
            </html>
        `))

		tmpl.Execute(w, struct{ Color string }{Color: color})
	}
}

func fetchColor(path string) (string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s", backendURL, path))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Color string `json:"color"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	return result.Color, nil
}
