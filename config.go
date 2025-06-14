package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config structure pour la configuration
type Config struct {
	LogFilePath     string `json:"log_file_path"`
	ServerPort      string `json:"server_port"`
	RefreshInterval int    `json:"refresh_interval"`
}

// LoadConfig charge la configuration depuis le fichier JSON
func LoadConfig(configPath string) (*Config, error) {
	// Valeurs par défaut
	config := &Config{
		LogFilePath:     "/root/logs/ckpool.log",
		ServerPort:      ":9000",
		RefreshInterval: 30,
	}

	// Vérifie si le fichier existe
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		fmt.Printf("Fichier de configuration non trouvé: %s, utilisation des valeurs par défaut\n", configPath)
		return config, nil
	}

	// Lit le fichier
	file, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("erreur lors de l'ouverture du fichier config: %v", err)
	}
	defer file.Close()

	// Décode le JSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		return nil, fmt.Errorf("erreur lors du décodage du fichier config: %v", err)
	}

	return config, nil
}

// CreateDefaultConfig crée un fichier de configuration par défaut
func CreateDefaultConfig(configPath string) error {
	config := &Config{
		LogFilePath:     "logs/ckpool.log",
		ServerPort:      ":8080",
		RefreshInterval: 30,
	}

	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("erreur lors de la création du fichier config: %v", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("erreur lors de l'encodage du fichier config: %v", err)
	}

	fmt.Printf("Fichier de configuration créé: %s\n", configPath)
	return nil
}
