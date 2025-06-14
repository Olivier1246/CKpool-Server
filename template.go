package main

import (
	"html/template"
	"net/http"
)

// HTMLTemplate génère et affiche le template HTML
func HTMLTemplate(w http.ResponseWriter, data LogData) {
	tmpl := `
<!DOCTYPE html>
<html lang="fr">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>CKPool Logs Dashboard</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            margin: 20px;
            background-color: #f5f5f5;
        }
        .container {
            max-width: 1400px;
            margin: 0 auto;
        }
        h1, h2 {
            color: #333;
            text-align: center;
        }
        .section {
            background: white;
            margin: 20px 0;
            padding: 20px;
            border-radius: 8px;
            box-shadow: 0 2px 4px rgba(0,0,0,0.1);
        }
        table {
            width: 100%;
            border-collapse: collapse;
            margin-top: 10px;
        }
        th, td {
            padding: 8px 12px;
            text-align: left;
            border-bottom: 1px solid #ddd;
            font-size: 12px;
        }
        th {
            background-color: #f8f9fa;
            font-weight: bold;
            position: sticky;
            top: 0;
        }
        tr:hover {
            background-color: #f5f5f5;
        }
        .timestamp {
            font-family: monospace;
            color: #666;
        }
        .address {
            font-family: monospace;
            font-size: 10px;
            word-break: break-all;
        }
        .hash {
            font-family: monospace;
            font-size: 10px;
            word-break: break-all;
            color: #007bff;
        }
        .refresh {
            position: fixed;
            top: 20px;
            right: 20px;
            padding: 10px 20px;
            background: #007bff;
            color: white;
            border: none;
            border-radius: 4px;
            cursor: pointer;
        }
        .refresh:hover {
            background: #0056b3;
        }
        .stats {
            display: grid;
            grid-template-columns: repeat(auto-fit, minmax(200px, 1fr));
            gap: 10px;
            margin-bottom: 20px;
        }
        .stat-card {
            background: #e9f4ff;
            padding: 15px;
            border-radius: 4px;
            text-align: center;
        }
        .stat-value {
            font-size: 24px;
            font-weight: bold;
            color: #007bff;
        }
        .stat-label {
            font-size: 12px;
            color: #666;
            margin-top: 5px;
        }
        .tab-container {
            margin-bottom: 20px;
        }
        .tabs {
            display: flex;
            border-bottom: 2px solid #ddd;
        }
        .tab {
            padding: 10px 20px;
            cursor: pointer;
            border: none;
            background: none;
            border-bottom: 2px solid transparent;
        }
        .tab.active {
            border-bottom-color: #007bff;
            color: #007bff;
            font-weight: bold;
        }
        .tab-content {
            display: none;
        }
        .tab-content.active {
            display: block;
        }
        .chart-controls {
            margin-bottom: 20px;
            text-align: center;
        }
        .hash-button {
            margin: 5px;
            padding: 8px 15px;
            background: #f8f9fa;
            border: 1px solid #ddd;
            border-radius: 4px;
            cursor: pointer;
            transition: all 0.3s;
        }
        .hash-button.active {
            background: #007bff;
            color: white;
            border-color: #007bff;
        }
        .hash-button:hover {
            background: #e9ecef;
        }
        .hash-button.active:hover {
            background: #0056b3;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>CKPool Logs Dashboard</h1>
        <button class="refresh" onclick="location.reload()">Actualiser</button>
        
        <div class="stats">
            <div class="stat-card">
                <div class="stat-value">{{len .Users}}</div>
                <div class="stat-label">Entrées Utilisateur</div>
            </div>
            <div class="stat-card">
                <div class="stat-value">{{len .Pool}}</div>
                <div class="stat-label">Entrées Pool</div>
            </div>
            <div class="stat-card">
                <div class="stat-value">{{len .Blocks}}</div>
                <div class="stat-label">Changements de Bloc</div>
            </div>
        </div>

        <div class="tab-container">
            <div class="tabs">
                <button class="tab active" onclick="showTab('charts')">Graphiques Hashrate</button>
                <button class="tab" onclick="showTab('users')">Données Utilisateur</button>
                <button class="tab" onclick="showTab('pool-status')">Statut Pool</button>
                <button class="tab" onclick="showTab('pool-hashrate')">Hashrate Pool</button>
                <button class="tab" onclick="showTab('pool-mining')">Mining Pool</button>
                <button class="tab" onclick="showTab('blocks')">Blocs</button>
            </div>
        </div>

        <div id="charts" class="tab-content active">
            <div class="section">
                <h2>Graphiques Hashrate Utilisateur</h2>
                <div class="chart-controls">
                    <button class="hash-button" onclick="toggleHashrate('1m', this)">Hash 1m (2h)</button>
                    <button class="hash-button" onclick="toggleHashrate('5m', this)">Hash 5m (6h)</button>
                    <button class="hash-button active" onclick="toggleHashrate('1h', this)">Hash 1h (12h)</button>
                    <button class="hash-button" onclick="toggleHashrate('1d', this)">Hash 1d (30j)</button>
                    <button class="hash-button" onclick="toggleHashrate('7d', this)">Hash 7d (6m)</button>
                </div>
                <div style="margin-bottom: 20px;">
                    <canvas id="userHashrateChart" width="400" height="200"></canvas>
                </div>
            </div>
            <div class="section">
                <h2>Graphiques Hashrate Pool</h2>
                <div class="chart-controls">
                    <button class="hash-button" onclick="togglePoolHashrate('1m', this)">Hash 1m (2h)</button>
                    <button class="hash-button" onclick="togglePoolHashrate('5m', this)">Hash 5m (6h)</button>
                    <button class="hash-button active" onclick="togglePoolHashrate('1h', this)">Hash 1h (12h)</button>
                    <button class="hash-button" onclick="togglePoolHashrate('1d', this)">Hash 1d (30j)</button>
                    <button class="hash-button" onclick="togglePoolHashrate('7d', this)">Hash 7d (6m)</button>
                </div>
                <div style="margin-bottom: 20px;">
                    <canvas id="poolHashrateChart" width="400" height="200"></canvas>
                </div>
            </div>
        </div>

        <div id="users" class="tab-content">
            <div class="section">
                <h2>Données Utilisateur</h2>
                <table>
                    <thead>
                        <tr>
                            <th>Timestamp</th>
                            <th>Adresse</th>
                            <th>Hash 1m</th>
                            <th>Hash 5m</th>
                            <th>Hash 1h</th>
                            <th>Hash 1d</th>
                            <th>Hash 7d</th>
                            <th>Workers</th>
                            <th>Shares</th>
                            <th>Best Share</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .Users}}
                        <tr>
                            <td class="timestamp">{{.Timestamp}}</td>
                            <td class="address">{{.Address}}</td>
                            <td>{{.Hashrate1m}}</td>
                            <td>{{.Hashrate5m}}</td>
                            <td>{{.Hashrate1hr}}</td>
                            <td>{{.Hashrate1d}}</td>
                            <td>{{.Hashrate7d}}</td>
                            <td>{{.Workers}}</td>
                            <td>{{.Shares}}</td>
                            <td>{{.BestShare}}</td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>

        <div id="pool-status" class="tab-content">
            <div class="section">
                <h2>Statut du Pool</h2>
                <table>
                    <thead>
                        <tr>
                            <th>Timestamp</th>
                            <th>Runtime</th>
                            <th>Last Update</th>
                            <th>Users</th>
                            <th>Workers</th>
                            <th>Idle</th>
                            <th>Disconnected</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .Pool}}
                        {{if eq .Type "status"}}
                        <tr>
                            <td class="timestamp">{{.Timestamp}}</td>
                            <td>{{.Runtime}}</td>
                            <td>{{.LastUpdate}}</td>
                            <td>{{.Users}}</td>
                            <td>{{.Workers}}</td>
                            <td>{{.Idle}}</td>
                            <td>{{.Disconnected}}</td>
                        </tr>
                        {{end}}
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>

        <div id="pool-hashrate" class="tab-content">
            <div class="section">
                <h2>Hashrate du Pool</h2>
                <table>
                    <thead>
                        <tr>
                            <th>Timestamp</th>
                            <th>Hash 1m</th>
                            <th>Hash 5m</th>
                            <th>Hash 15m</th>
                            <th>Hash 1h</th>
                            <th>Hash 6h</th>
                            <th>Hash 1d</th>
                            <th>Hash 7d</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .Pool}}
                        {{if eq .Type "hashrate"}}
                        <tr>
                            <td class="timestamp">{{.Timestamp}}</td>
                            <td>{{.Hashrate1m}}</td>
                            <td>{{.Hashrate5m}}</td>
                            <td>{{.Hashrate15m}}</td>
                            <td>{{.Hashrate1hr}}</td>
                            <td>{{.Hashrate6hr}}</td>
                            <td>{{.Hashrate1d}}</td>
                            <td>{{.Hashrate7d}}</td>
                        </tr>
                        {{end}}
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>

        <div id="pool-mining" class="tab-content">
            <div class="section">
                <h2>Mining du Pool</h2>
                <table>
                    <thead>
                        <tr>
                            <th>Timestamp</th>
                            <th>Difficulty</th>
                            <th>Accepted</th>
                            <th>Rejected</th>
                            <th>Best Share</th>
                            <th>SPS 1m</th>
                            <th>SPS 5m</th>
                            <th>SPS 15m</th>
                            <th>SPS 1h</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .Pool}}
                        {{if eq .Type "mining"}}
                        <tr>
                            <td class="timestamp">{{.Timestamp}}</td>
                            <td>{{.Diff}}</td>
                            <td>{{.Accepted}}</td>
                            <td>{{.Rejected}}</td>
                            <td>{{.BestShare}}</td>
                            <td>{{.SPS1m}}</td>
                            <td>{{.SPS5m}}</td>
                            <td>{{.SPS15m}}</td>
                            <td>{{.SPS1h}}</td>
                        </tr>
                        {{end}}
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>

        <div id="blocks" class="tab-content">
            <div class="section">
                <h2>Changements de Blocs</h2>
                <table>
                    <thead>
                        <tr>
                            <th>Timestamp</th>
                            <th>Hash du Bloc</th>
                        </tr>
                    </thead>
                    <tbody>
                        {{range .Blocks}}
                        <tr>
                            <td class="timestamp">{{.Timestamp}}</td>
                            <td class="hash">{{.Hash}}</td>
                        </tr>
                        {{end}}
                    </tbody>
                </table>
            </div>
        </div>
    </div>

    <script src="https://cdnjs.cloudflare.com/ajax/libs/Chart.js/3.9.1/chart.min.js"></script>
    <script>
        let userChart = null;
        let poolChart = null;
        let currentUserHashType = '1h';
        let currentPoolHashType = '1h';

        function showTab(tabName) {
            // Cacher tous les contenus
            const contents = document.querySelectorAll('.tab-content');
            contents.forEach(content => content.classList.remove('active'));

            // Désactiver tous les onglets
            const tabs = document.querySelectorAll('.tab');
            tabs.forEach(tab => tab.classList.remove('active'));

            // Afficher le contenu sélectionné
            document.getElementById(tabName).classList.add('active');
            
            // Activer l'onglet cliqué
            event.target.classList.add('active');
        }

        // Fonction pour convertir les hashrates en nombre
        function parseHashrate(hashrate) {
            if (!hashrate) return 0;
            const value = parseFloat(hashrate);
            if (hashrate.includes('T')) return value * 1000000000000;
            if (hashrate.includes('G')) return value * 1000000000;
            if (hashrate.includes('M')) return value * 1000000;
            if (hashrate.includes('K')) return value * 1000;
            return value;
        }

        // Fonction pour formater les hashrates
        function formatHashrate(value) {
            if (value >= 1000000000000) return (value / 1000000000000).toFixed(2) + 'T';
            if (value >= 1000000000) return (value / 1000000000).toFixed(2) + 'G';
            if (value >= 1000000) return (value / 1000000).toFixed(2) + 'M';
            if (value >= 1000) return (value / 1000).toFixed(2) + 'K';
            return value.toString();
        }

        // Fonction pour calculer le nombre de points de données basé sur le type de hash
        function getDataPointsCount(hashType) {
            switch(hashType) {
                case '1m': return 120; // 2 heures (120 minutes)
                case '5m': return 72;  // 6 heures (72 * 5 minutes)
                case '1h': return 12;  // 12 heures
                case '1d': return 30;  // 30 jours
                case '7d': return 26;  // 6 mois (26 semaines)
                default: return 12;
            }
        }

        // Fonction pour obtenir les données de hashrate selon le type
        function getHashrateData(data, hashType, isPool = false) {
            const count = getDataPointsCount(hashType);
            const filteredData = isPool ? 
                data.filter(item => item.type === 'hashrate').slice(0, count) : 
                data.slice(0, count);
            
            let hashrateField;
            switch(hashType) {
                case '1m': hashrateField = 'hashrate1m'; break;
                case '5m': hashrateField = 'hashrate5m'; break;
                case '1h': hashrateField = 'hashrate1hr'; break;
                case '1d': hashrateField = 'hashrate1d'; break;
                case '7d': hashrateField = 'hashrate7d'; break;
                default: hashrateField = 'hashrate1hr';
            }
            
            // Filter data by time window
            const now = new Date();
            let windowMs;
            switch(hashType) {
                case '1m': windowMs = 2 * 60 * 60 * 1000; break; // 2 hours
                case '5m': windowMs = 6 * 60 * 60 * 1000; break; // 6 hours
                case '1h': windowMs = 12 * 60 * 60 * 1000; break; // 12 hours
                case '1d': windowMs = 30 * 24 * 60 * 60 * 1000; break; // 30 days
                case '7d': windowMs = 6 * 30 * 24 * 60 * 60 * 1000; break; // 6 months approx
                default: windowMs = 12 * 60 * 60 * 1000;
            }
            const filteredByTime = filteredData.filter(item => {
                const date = new Date(item.timestamp);
                return (now - date) <= windowMs;
            });

            return {
                labels: filteredByTime.map(item => {
                    const date = new Date(item.timestamp);
                    if (hashType === '1d' || hashType === '7d') {
                        return date.toLocaleDateString();
                    } else {
                        return date.toLocaleTimeString();
                    }
                }).reverse(),
                values: filteredByTime.map(item => parseHashrate(item[hashrateField])).reverse()
            };

        // Fonction pour basculer l'affichage du hashrate utilisateur
        function toggleHashrate(hashType, button) {
            // Mettre à jour les boutons actifs
            document.querySelectorAll('#charts .chart-controls .hash-button').forEach(btn => {
                if (btn.parentElement === button.parentElement) {
                    btn.classList.remove('active');
                }
            });
            button.classList.add('active');
            
            currentUserHashType = hashType;
            updateUserChart();
        }

        // Fonction pour basculer l'affichage du hashrate pool
        function togglePoolHashrate(hashType, button) {
            // Mettre à jour les boutons actifs
            const poolSection = button.closest('.section');
            poolSection.querySelectorAll('.hash-button').forEach(btn => {
                btn.classList.remove('active');
            });
            button.classList.add('active');
            
            currentPoolHashType = hashType;
            updatePoolChart();
        }

        // Fonction pour mettre à jour le graphique utilisateur
        function updateUserChart() {
            if (userChart) {
                userChart.destroy();
            }
            
            const userData = {{.Users}};
            const data = getHashrateData(userData, currentUserHashType);
            const color = getColorForHashType(currentUserHashType);
            
            const userCtx = document.getElementById('userHashrateChart').getContext('2d');
            userChart = new Chart(userCtx, {
                type: 'line',
                data: {
                    labels: data.labels,
                    datasets: [{
                        label: 'Hash ' + currentUserHashType,
                        data: data.values,
                        borderColor: color.border,
                        backgroundColor: color.background,
                        tension: 0.1,
                        fill: true
                    }]
                },
                options: {
                    responsive: true,
                    plugins: {
                        title: {
                            display: true,
                            text: 'Évolution du Hashrate Utilisateur - ' + getTimeRangeLabel(currentUserHashType)
                        },
                        legend: {
                            display: true,
                            position: 'top'
                        }
                    },
                    scales: {
                        x: {
                            display: true,
                            title: {
                                display: true,
                                text: 'Temps'
                            }
                        },
                        y: {
                            display: true,
                            title: {
                                display: true,
                                text: 'Hashrate (H/s)'
                            },
                            ticks: {
                                callback: function(value) {
                                    return formatHashrate(value);
                                }
                            }
                        }
                    }
                }
            });
        }

        // Fonction pour mettre à jour le graphique pool
        function updatePoolChart() {
            if (poolChart) {
                poolChart.destroy();
            }
            
            const poolData = {{.Pool}};
            const data = getHashrateData(poolData, currentPoolHashType, true);
            const color = getColorForHashType(currentPoolHashType);
            
            const poolCtx = document.getElementById('poolHashrateChart').getContext('2d');
            poolChart = new Chart(poolCtx, {
                type: 'line',
                data: {
                    labels: data.labels,
                    datasets: [{
                        label: 'Hash ' + currentPoolHashType,
                        data: data.values,
                        borderColor: color.border,
                        backgroundColor: color.background,
                        tension: 0.1,
                        fill: true
                    }]
                },
                options: {
                    responsive: true,
                    plugins: {
                        title: {
                            display: true,
                            text: 'Évolution du Hashrate Pool - ' + getTimeRangeLabel(currentPoolHashType)
                        },
                        legend: {
                            display: true,
                            position: 'top'
                        }
                    },
                    scales: {
                        x: {
                            display: true,
                            title: {
                                display: true,
                                text: 'Temps'
                            }
                        },
                        y: {
                            display: true,
                            title: {
                                display: true,
                                text: 'Hashrate (H/s)'
                            },
                            ticks: {
                                callback: function(value) {
                                    return formatHashrate(value);
                                }
                            }
                        }
                    }
                }
            });
        }

        // Fonction pour obtenir la couleur selon le type de hash
        function getColorForHashType(hashType) {
            const colors = {
                '1m': { border: 'rgb(255, 99, 132)', background: 'rgba(255, 99, 132, 0.1)' },
                '5m': { border: 'rgb(54, 162, 235)', background: 'rgba(54, 162, 235, 0.1)' },
                '1h': { border: 'rgb(255, 205, 86)', background: 'rgba(255, 205, 86, 0.1)' },
                '1d': { border: 'rgb(75, 192, 192)', background: 'rgba(75, 192, 192, 0.1)' },
                '7d': { border: 'rgb(153, 102, 255)', background: 'rgba(153, 102, 255, 0.1)' }
            };
            return colors[hashType] || colors['1h'];
        }

        // Fonction pour obtenir le libellé de la plage de temps
        function getTimeRangeLabel(hashType) {
            const labels = {
                '1m': 'Sur 2 heures',
                '5m': 'Sur 6 heures',
                '1h': 'Sur 12 heures',
                '1d': 'Sur 30 jours',
                '7d': 'Sur 6 mois'
            };
            return labels[hashType] || 'Sur 12 heures';
        }

        // Initialiser les graphiques au chargement de la page
        document.addEventListener('DOMContentLoaded', function() {
            updateUserChart();
            updatePoolChart();
        });

        // Auto-refresh toutes les 30 secondes
        setInterval(() => {
            location.reload();
        }, 30000);
    </script>
</body>
</html>
`

	// Créer le template et l'exécuter
	t, err := template.New("dashboard").Parse(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}