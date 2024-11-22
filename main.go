package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const backendURL = "http://backend-service.default.svc.cluster.local:80"

func main() {
	http.HandleFunc("/primario", handleColor("primario"))
	http.HandleFunc("/secundario", handleColor("secundario"))

	fmt.Println("Starting server on :8080")
	fmt.Println("Backend: ", backendURL)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleColor(path string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Entro aqui 1")
		color, err := fetchColor(path)
		fmt.Println("Entro aqui 2")
		if err != nil {
			println(err)
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
	fmt.Printf("Llamando: %s/%s\n", backendURL, path)
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
