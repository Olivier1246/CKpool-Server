package main

import (
	"bufio"
	"encoding/json"
	"os"
	"regexp"
	"sort"
	"strings"
)

// parseLogFile lit et analyse le fichier de log CKPool
func parseLogFile(filename string) (*LogData, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data := &LogData{
		Users:  []UserData{},
		Pool:   []PoolStats{},
		Blocks: []BlockData{},
	}

	scanner := bufio.NewScanner(file)
	timestampRegex := regexp.MustCompile(`\[([^\]]+)\]`)
	userRegex := regexp.MustCompile(`User ([^:]+):(.+)`)
	poolRegex := regexp.MustCompile(`Pool:(.+)`)
	blockRegex := regexp.MustCompile(`Block hash changed to (.+)`)

	for scanner.Scan() {
		line := scanner.Text()

		// Ignorer les lignes "Stored local workbase"
		if strings.Contains(line, "Stored local workbase with") {
			continue
		}

		// Extraire le timestamp
		timestampMatch := timestampRegex.FindStringSubmatch(line)
		if len(timestampMatch) < 2 {
			continue
		}
		timestamp := timestampMatch[1]

		// Traiter les données utilisateur
		if userMatch := userRegex.FindStringSubmatch(line); len(userMatch) >= 3 {
			user := parseUserData(userMatch[1], userMatch[2], timestamp)
			if user != nil {
				data.Users = append(data.Users, *user)
			}
		}

		// Traiter les données de pool
		if poolMatch := poolRegex.FindStringSubmatch(line); len(poolMatch) >= 2 {
			pool := parsePoolData(poolMatch[1], timestamp)
			if pool != nil {
				data.Pool = append(data.Pool, *pool)
			}
		}

		// Traiter les changements de hash de bloc
		if blockMatch := blockRegex.FindStringSubmatch(line); len(blockMatch) >= 2 {
			block := BlockData{
				Timestamp: timestamp,
				Hash:      blockMatch[1],
			}
			data.Blocks = append(data.Blocks, block)
		}
	}

	// Trier par timestamp (plus récent en premier)
	sortDataByTimestamp(data)

	return data, scanner.Err()
}

// parseUserData analyse les données JSON d'un utilisateur
func parseUserData(address, jsonData, timestamp string) *UserData {
	var userData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &userData); err != nil {
		return nil
	}

	user := &UserData{
		Timestamp: timestamp,
		Address:   address,
	}

	// Extraire les données avec vérification de type
	if val, ok := userData["hashrate1m"].(string); ok {
		user.Hashrate1m = val
	}
	if val, ok := userData["hashrate5m"].(string); ok {
		user.Hashrate5m = val
	}
	if val, ok := userData["hashrate1hr"].(string); ok {
		user.Hashrate1hr = val
	}
	if val, ok := userData["hashrate1d"].(string); ok {
		user.Hashrate1d = val
	}
	if val, ok := userData["hashrate7d"].(string); ok {
		user.Hashrate7d = val
	}
	if val, ok := userData["lastshare"].(float64); ok {
		user.LastShare = int64(val)
	}
	if val, ok := userData["workers"].(float64); ok {
		user.Workers = int(val)
	}
	if val, ok := userData["shares"].(float64); ok {
		user.Shares = int64(val)
	}
	if val, ok := userData["bestshare"].(float64); ok {
		user.BestShare = int64(val)
	}

	return user
}

// parsePoolData analyse les données JSON du pool
func parsePoolData(jsonData, timestamp string) *PoolStats {
	var poolData map[string]interface{}
	if err := json.Unmarshal([]byte(jsonData), &poolData); err != nil {
		return nil
	}

	pool := &PoolStats{
		Timestamp: timestamp,
	}

	// Déterminer le type de données pool et extraire les valeurs appropriées
	if _, ok := poolData["runtime"]; ok {
		pool.Type = "status"
		extractStatusData(pool, poolData)
	} else if _, ok := poolData["hashrate1m"]; ok {
		pool.Type = "hashrate"
		extractHashrateData(pool, poolData)
	} else if _, ok := poolData["diff"]; ok {
		pool.Type = "mining"
		extractMiningData(pool, poolData)
	}

	return pool
}

// extractStatusData extrait les données de statut du pool
func extractStatusData(pool *PoolStats, data map[string]interface{}) {
	if val, ok := data["runtime"].(float64); ok {
		pool.Runtime = int(val)
	}
	if val, ok := data["lastupdate"].(float64); ok {
		pool.LastUpdate = int64(val)
	}
	if val, ok := data["Users"].(float64); ok {
		pool.Users = int(val)
	}
	if val, ok := data["Workers"].(float64); ok {
		pool.Workers = int(val)
	}
	if val, ok := data["Idle"].(float64); ok {
		pool.Idle = int(val)
	}
	if val, ok := data["Disconnected"].(float64); ok {
		pool.Disconnected = int(val)
	}
}

// extractHashrateData extrait les données de hashrate du pool
func extractHashrateData(pool *PoolStats, data map[string]interface{}) {
	if val, ok := data["hashrate1m"].(string); ok {
		pool.Hashrate1m = val
	}
	if val, ok := data["hashrate5m"].(string); ok {
		pool.Hashrate5m = val
	}
	if val, ok := data["hashrate15m"].(string); ok {
		pool.Hashrate15m = val
	}
	if val, ok := data["hashrate1hr"].(string); ok {
		pool.Hashrate1hr = val
	}
	if val, ok := data["hashrate6hr"].(string); ok {
		pool.Hashrate6hr = val
	}
	if val, ok := data["hashrate1d"].(string); ok {
		pool.Hashrate1d = val
	}
	if val, ok := data["hashrate7d"].(string); ok {
		pool.Hashrate7d = val
	}
}

// extractMiningData extrait les données de mining du pool
func extractMiningData(pool *PoolStats, data map[string]interface{}) {
	if val, ok := data["diff"].(float64); ok {
		pool.Diff = val
	}
	if val, ok := data["accepted"].(float64); ok {
		pool.Accepted = int64(val)
	}
	if val, ok := data["rejected"].(float64); ok {
		pool.Rejected = int64(val)
	}
	if val, ok := data["bestshare"].(float64); ok {
		pool.BestShare = int64(val)
	}
	if val, ok := data["SPS1m"].(float64); ok {
		pool.SPS1m = val
	}
	if val, ok := data["SPS5m"].(float64); ok {
		pool.SPS5m = val
	}
	if val, ok := data["SPS15m"].(float64); ok {
		pool.SPS15m = val
	}
	if val, ok := data["SPS1h"].(float64); ok {
		pool.SPS1h = val
	}
}

// sortDataByTimestamp trie toutes les données par timestamp (plus récent en premier)
func sortDataByTimestamp(data *LogData) {
	sort.Slice(data.Users, func(i, j int) bool {
		return data.Users[i].Timestamp > data.Users[j].Timestamp
	})
	sort.Slice(data.Pool, func(i, j int) bool {
		return data.Pool[i].Timestamp > data.Pool[j].Timestamp
	})
	sort.Slice(data.Blocks, func(i, j int) bool {
		return data.Blocks[i].Timestamp > data.Blocks[j].Timestamp
	})
}
