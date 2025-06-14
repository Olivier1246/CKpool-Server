package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Parse du template HTML
	tmpl, err := template.New("dashboard").Parse(HTMLTemplate)
	if err != nil {
		log.Fatal("Erreur lors du parsing du template:", err)
	}

	// Handler pour la page principale
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Lire et parser le fichier de log
		data, err := parseLogFile("logs/ckpool.log")
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de la lecture du fichier: %v", err), http.StatusInternalServerError)
			return
		}

		// Générer la page HTML
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de la génération de la page: %v", err), http.StatusInternalServerError)
			return
		}
	})

	// API endpoint pour récupérer les données en JSON
	http.HandleFunc("/api/data", func(w http.ResponseWriter, r *http.Request) {
		data, err := parseLogFile("logs/ckpool.log")
		if err != nil {
			http.Error(w, fmt.Sprintf("Erreur lors de la lecture du fichier: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	})

	fmt.Println("Serveur démarré sur http://0.0.0.0:9000")
	fmt.Println("Fichier de log: logs/ckpool.log")
	fmt.Println("API JSON disponible sur: http://0.0.0.0:9000/api/data")

	log.Fatal(http.ListenAndServe("0.0.0.0:9000", nil))
}
