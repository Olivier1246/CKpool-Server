package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

var config *Config

func main() {
	// Définit les flags de ligne de commande
	configPath := flag.String("config", "config.json", "Chemin vers le fichier de configuration")
	createConfig := flag.Bool("create-config", false, "Crée un fichier de configuration par défaut")
	flag.Parse()

	// Crée un fichier de config par défaut si demandé
	if *createConfig {
		if err := CreateDefaultConfig(*configPath); err != nil {
			log.Fatalf("Erreur lors de la création du fichier config: %v", err)
		}
		return
	}

	// Charge la configuration
	var err error
	config, err = LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("Erreur lors du chargement de la configuration: %v", err)
	}

	// Vérifie si le fichier log existe
	if _, err := os.Stat(config.LogFilePath); os.IsNotExist(err) {
		fmt.Printf("ATTENTION: Le fichier log n'existe pas: %s\n", config.LogFilePath)
		fmt.Printf("Assurez-vous que le chemin est correct dans le fichier de configuration.\n")
	}

	fmt.Printf("Configuration chargée:\n")
	fmt.Printf("- Fichier log: %s\n", config.LogFilePath)
	fmt.Printf("- Port serveur: %s\n", config.ServerPort)
	fmt.Printf("- Intervalle de rafraîchissement: %d secondes\n", config.RefreshInterval)

	// Configure les routes HTTP
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/api/data", handleAPIData)

	fmt.Printf("Serveur démarré sur http://localhost%s\n", config.ServerPort)
	log.Fatal(http.ListenAndServe(config.ServerPort, nil))
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	// Utilise config.LogFilePath au lieu du chemin codé en dur
	data, err := parseLogFile(config.LogFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du fichier: %v", err), http.StatusInternalServerError)
		return
	}

	// Parse et exécute le template HTML
	tmpl, err := template.New("dashboard").Parse(HTMLTemplate)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors du parsing du template: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'exécution du template: %v", err), http.StatusInternalServerError)
		return
	}
}

func handleAPIData(w http.ResponseWriter, r *http.Request) {
	// Utilise config.LogFilePath au lieu du chemin codé en dur
	data, err := parseLogFile(config.LogFilePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de la lecture du fichier: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("Erreur lors de l'encodage JSON: %v", err), http.StatusInternalServerError)
		return
	}
}
