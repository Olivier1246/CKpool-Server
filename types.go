package main

// UserData représente les données d'un utilisateur du pool
type UserData struct {
	Timestamp   string `json:"timestamp"`
	Address     string `json:"address"`
	Hashrate1m  string `json:"hashrate1m"`
	Hashrate5m  string `json:"hashrate5m"`
	Hashrate1hr string `json:"hashrate1hr"`
	Hashrate1d  string `json:"hashrate1d"`
	Hashrate7d  string `json:"hashrate7d"`
	LastShare   int64  `json:"lastshare"`
	Workers     int    `json:"workers"`
	Shares      int64  `json:"shares"`
	BestShare   int64  `json:"bestshare"`
}

// PoolStats représente les statistiques du pool
type PoolStats struct {
	Timestamp    string  `json:"timestamp"`
	Type         string  `json:"type"`
	Runtime      int     `json:"runtime,omitempty"`
	LastUpdate   int64   `json:"lastupdate,omitempty"`
	Users        int     `json:"Users,omitempty"`
	Workers      int     `json:"Workers,omitempty"`
	Idle         int     `json:"Idle,omitempty"`
	Disconnected int     `json:"Disconnected,omitempty"`
	Hashrate1m   string  `json:"hashrate1m,omitempty"`
	Hashrate5m   string  `json:"hashrate5m,omitempty"`
	Hashrate15m  string  `json:"hashrate15m,omitempty"`
	Hashrate1hr  string  `json:"hashrate1hr,omitempty"`
	Hashrate6hr  string  `json:"hashrate6hr,omitempty"`
	Hashrate1d   string  `json:"hashrate1d,omitempty"`
	Hashrate7d   string  `json:"hashrate7d,omitempty"`
	Diff         float64 `json:"diff,omitempty"`
	Accepted     int64   `json:"accepted,omitempty"`
	Rejected     int64   `json:"rejected,omitempty"`
	BestShare    int64   `json:"bestshare,omitempty"`
	SPS1m        float64 `json:"SPS1m,omitempty"`
	SPS5m        float64 `json:"SPS5m,omitempty"`
	SPS15m       float64 `json:"SPS15m,omitempty"`
	SPS1h        float64 `json:"SPS1h,omitempty"`
}

// BlockData représente les données d'un bloc
type BlockData struct {
	Timestamp string `json:"timestamp"`
	Hash      string `json:"hash"`
}

// LogData contient toutes les données parsées du fichier de log
type LogData struct {
	Users  []UserData  `json:"users"`
	Pool   []PoolStats `json:"pool"`
	Blocks []BlockData `json:"blocks"`
}
