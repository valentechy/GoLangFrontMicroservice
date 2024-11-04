package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const backendURL = "http://backend-service"

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
                <h1>The background color is {{.Color}}</h1>
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

	var color string
	_, err = fmt.Fscanf(resp.Body, "%s", &color)
	if err != nil {
		return "", err
	}

	return color, nil
}
