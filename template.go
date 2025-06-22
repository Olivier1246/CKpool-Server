package main

import (
	"html/template"
	"net/http"
)

// HTMLTemplate contient le template HTML complet pour le dashboard
const HTMLTemplate = `
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
                <div style="margin-bottom: 20px;">
                    <canvas id="userHashrateChart" width="400" height="200"></canvas>
                </div>
            </div>
            <div class="section">
                <h2>Graphiques Hashrate Pool</h2>
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

        // Fonction pour filtrer les données selon la période
        function filterDataByPeriod(data, hashrateType) {
            const now = new Date();
            let filteredData = [];
            
            switch(hashrateType) {
                case '1m':
                    // Hash 1m : 2 dernières heures (120 points max)
                    const twoHoursAgo = new Date(now.getTime() - 2 * 60 * 60 * 1000);
                    filteredData = data.filter(item => {
                        const itemDate = new Date(item.timestamp);
                        return itemDate >= twoHoursAgo;
                    }).slice(-120);
                    break;
                case '5m':
                    // Hash 5m : 6 dernières heures (72 points max)
                    const sixHoursAgo = new Date(now.getTime() - 6 * 60 * 60 * 1000);
                    filteredData = data.filter(item => {
                        const itemDate = new Date(item.timestamp);
                        return itemDate >= sixHoursAgo;
                    }).slice(-72);
                    break;
                case '1h':
                    // Hash 1h : 12 dernières heures (12 points max)
                    const twelveHoursAgo = new Date(now.getTime() - 12 * 60 * 60 * 1000);
                    filteredData = data.filter(item => {
                        const itemDate = new Date(item.timestamp);
                        return itemDate >= twelveHoursAgo;
                    }).slice(-12);
                    break;
                case '1d':
                    // Hash 1d : 30 derniers jours (30 points max)
                    const thirtyDaysAgo = new Date(now.getTime() - 30 * 24 * 60 * 60 * 1000);
                    filteredData = data.filter(item => {
                        const itemDate = new Date(item.timestamp);
                        return itemDate >= thirtyDaysAgo;
                    }).slice(-30);
                    break;
                case '7d':
                    // Hash 7d : 6 derniers mois (24 points max)
                    const sixMonthsAgo = new Date(now.getTime() - 6 * 30 * 24 * 60 * 60 * 1000);
                    filteredData = data.filter(item => {
                        const itemDate = new Date(item.timestamp);
                        return itemDate >= sixMonthsAgo;
                    }).slice(-24);
                    break;
                default:
                    filteredData = data.slice(-12);
            }
            
            return filteredData.reverse();
        }

        // Initialisation des graphiques
        function initCharts() {
            const userData = {{.Users}};
            const poolData = {{.Pool}};

            // Filtrer les données pour Hash 1h (par défaut)
            const filteredUserData = filterDataByPeriod(userData, '1h');
            const userLabels = filteredUserData.map(item => item.timestamp.split(' ')[1]);
            const userHashrate1h = filteredUserData.map(item => parseHashrate(item.hashrate1hr));

            // Graphique utilisateur - affichage uniquement Hash 1h par défaut
            const userCtx = document.getElementById('userHashrateChart').getContext('2d');
            new Chart(userCtx, {
                type: 'line',
                data: {
                    labels: userLabels,
                    datasets: [
                        {
                            label: 'Hash 1h',
                            data: userHashrate1h,
                            borderColor: 'rgb(255, 205, 86)',
                            backgroundColor: 'rgba(255, 205, 86, 0.1)',
                            tension: 0.1,
                            hidden: false
                        }
                    ]
                },
                options: {
                    responsive: true,
                    plugins: {
                        title: {
                            display: true,
                            text: 'Évolution du Hashrate Utilisateur (12 dernières heures)'
                        },
                        legend: {
                            display: true,
                            position: 'top',
                            onClick: function(e, legendItem) {
                                // Personnaliser le comportement du clic sur la légende
                                const chart = this.chart;
                                const dataset = chart.data.datasets[legendItem.datasetIndex];
                                const hashrateType = legendItem.text.includes('1m') ? '1m' :
                                                   legendItem.text.includes('5m') ? '5m' :
                                                   legendItem.text.includes('1h') ? '1h' :
                                                   legendItem.text.includes('1d') ? '1d' : '7d';
                                
                                // Filtrer et mettre à jour les données
                                const newData = filterDataByPeriod(userData, hashrateType);
                                chart.data.labels = newData.map(item => item.timestamp.split(' ')[1]);
                                
                                // Mettre à jour le titre selon le type
                                let titleText = 'Évolution du Hashrate Utilisateur';
                                switch(hashrateType) {
                                    case '1m': titleText += ' (2 dernières heures)'; break;
                                    case '5m': titleText += ' (6 dernières heures)'; break;
                                    case '1h': titleText += ' (12 dernières heures)'; break;
                                    case '1d': titleText += ' (30 derniers jours)'; break;
                                    case '7d': titleText += ' (6 derniers mois)'; break;
                                }
                                chart.options.plugins.title.text = titleText;
                                
                                // Masquer tous les datasets
                                chart.data.datasets.forEach(ds => ds.hidden = true);
                                
                                // Afficher seulement le dataset sélectionné avec les nouvelles données
                                const fieldName = hashrateType === '1h' ? 'hashrate1hr' : 
                                                hashrateType === '1d' ? 'hashrate1d' :
                                                hashrateType === '7d' ? 'hashrate7d' :
                                                'hashrate' + hashrateType;
                                dataset.data = newData.map(item => parseHashrate(item[fieldName]));
                                dataset.hidden = false;
                                
                                chart.update();
                            }
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

            // Préparation des données pool (hashrate) - Hash 1h par défaut
            const poolHashrateData = poolData.filter(item => item.type === 'hashrate');
            const filteredPoolData = filterDataByPeriod(poolHashrateData, '1h');
            const poolLabels = filteredPoolData.map(item => item.timestamp.split(' ')[1]);
            const poolHashrate1h = filteredPoolData.map(item => parseHashrate(item.hashrate1hr));

            // Graphique pool - affichage uniquement Hash 1h par défaut
            const poolCtx = document.getElementById('poolHashrateChart').getContext('2d');
            new Chart(poolCtx, {
                type: 'line',
                data: {
                    labels: poolLabels,
                    datasets: [
                        {
                            label: 'Hash 1h',
                            data: poolHashrate1h,
                            borderColor: 'rgb(255, 205, 86)',
                            backgroundColor: 'rgba(255, 205, 86, 0.1)',
                            tension: 0.1,
                            hidden: false
                        }
                    ]
                },
                options: {
                    responsive: true,
                    plugins: {
                        title: {
                            display: true,
                            text: 'Évolution du Hashrate Pool (12 dernières heures)'
                        },
                        legend: {
                            display: true,
                            position: 'top',
                            onClick: function(e, legendItem) {
                                // Même logique pour le graphique pool
                                const chart = this.chart;
                                const dataset = chart.data.datasets[legendItem.datasetIndex];
                                const hashrateType = legendItem.text.includes('1m') ? '1m' :
                                                   legendItem.text.includes('5m') ? '5m' :
                                                   legendItem.text.includes('1h') ? '1h' :
                                                   legendItem.text.includes('1d') ? '1d' : '7d';
                                
                                const newData = filterDataByPeriod(poolHashrateData, hashrateType);
                                chart.data.labels = newData.map(item => item.timestamp.split(' ')[1]);
                                
                                let titleText = 'Évolution du Hashrate Pool';
                                switch(hashrateType) {
                                    case '1m': titleText += ' (2 dernières heures)'; break;
                                    case '5m': titleText += ' (6 dernières heures)'; break;
                                    case '1h': titleText += ' (12 dernières heures)'; break;
                                    case '1d': titleText += ' (30 derniers jours)'; break;
                                    case '7d': titleText += ' (6 derniers mois)'; break;
                                }
                                chart.options.plugins.title.text = titleText;
                                
                                chart.data.datasets.forEach(ds => ds.hidden = true);
                                
                                const fieldName = hashrateType === '1h' ? 'hashrate1hr' : 
                                                hashrateType === '1d' ? 'hashrate1d' :
                                                hashrateType === '7d' ? 'hashrate7d' :
                                                'hashrate' + hashrateType;
                                dataset.data = newData.map(item => parseHashrate(item[fieldName]));
                                dataset.hidden = false;
                                
                                chart.update();
                            }
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

            // Ajouter les autres datasets cachés pour permettre la commutation
            const userChart = Chart.getChart('userHashrateChart');
            const poolChart = Chart.getChart('poolHashrateChart');

            // Ajouter tous les datasets pour l'utilisateur
            userChart.data.datasets.push(
                {
                    label: 'Hash 1m',
                    data: [],
                    borderColor: 'rgb(255, 99, 132)',
                    backgroundColor: 'rgba(255, 99, 132, 0.1)',
                    tension: 0.1,
                    hidden: true
                },
                {
                    label: 'Hash 5m',
                    data: [],
                    borderColor: 'rgb(54, 162, 235)',
                    backgroundColor: 'rgba(54, 162, 235, 0.1)',
                    tension: 0.1,
                    hidden: true
                },
                {
                    label: 'Hash 1d',
                    data: [],
                    borderColor: 'rgb(75, 192, 192)',
                    backgroundColor: 'rgba(75, 192, 192, 0.1)',
                    tension: 0.1,
                    hidden: true
                },
                {
                    label: 'Hash 7d',
                    data: [],
                    borderColor: 'rgb(153, 102, 255)',
                    backgroundColor: 'rgba(153, 102, 255, 0.1)',
                    tension: 0.1,
                    hidden: true
                }
            );

            // Ajouter tous les datasets pour le pool
            poolChart.data.datasets.push(
                {
                    label: 'Hash 1m',
                    data: [],
                    borderColor: 'rgb(255, 99, 132)',
                    backgroundColor: 'rgba(255, 99, 132, 0.1)',
                    tension: 0.1,
                    hidden: true
                },
                {
                    label: 'Hash 5m',
                    data: [],
                    borderColor: 'rgb(54, 162, 235)',
                    backgroundColor: 'rgba(54, 162, 235, 0.1)',
                    tension: 0.1,
                    hidden: true
                },
                {
                    label: 'Hash 1d',
                    data: [],
                    borderColor: 'rgb(75, 192, 192)',
                    backgroundColor: 'rgba(75, 192, 192, 0.1)',
                    tension: 0.1,
                    hidden: true
                },
                {
                    label: 'Hash 7d',
                    data: [],
                    borderColor: 'rgb(153, 102, 255)',
                    backgroundColor: 'rgba(153, 102, 255, 0.1)',
                    tension: 0.1,
                    hidden: true
                }
            );

            userChart.update();
            poolChart.update();
        }

        // Initialiser les graphiques au chargement de la page
        document.addEventListener('DOMContentLoaded', function() {
            initCharts();
        });

        // Auto-refresh toutes les 30 secondes
        setInterval(() => {
            location.reload();
        }, 30000);
    </script>
</body>
</html>
`

// HTMLTemplate rend le template HTML avec les données fournies
func HTMLTemplate(w http.ResponseWriter, data LogData) {
	tmpl, err := template.New("dashboard").Parse(htmlTemplateString)
	if err != nil {
		http.Error(w, "Erreur lors du parsing du template: "+err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Erreur lors de l'exécution du template: "+err.Error(), http.StatusInternalServerError)
		return
	}
}
