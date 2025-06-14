package main

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

        // Initialisation des graphiques
        function initCharts() {
            const userData = {{.Users}};
            const poolData = {{.Pool}};

            // Préparation des données utilisateur
            const userLabels = userData.slice(0, 20).map(item => item.timestamp.split(' ')[1]).reverse();
            const userHashrate1m = userData.slice(0, 20).map(item => parseHashrate(item.hashrate1m)).reverse();
            const userHashrate5m = userData.slice(0, 20).map(item => parseHashrate(item.hashrate5m)).reverse();
            const userHashrate1h = userData.slice(0, 20).map(item => parseHashrate(item.hashrate1hr)).reverse();
            const userHashrate1d = userData.slice(0, 20).map(item => parseHashrate(item.hashrate1d)).reverse();
            const userHashrate7d = userData.slice(0, 20).map(item => parseHashrate(item.hashrate7d)).reverse();

            // Graphique utilisateur
            const userCtx = document.getElementById('userHashrateChart').getContext('2d');
            new Chart(userCtx, {
                type: 'line',
                data: {
                    labels: userLabels,
                    datasets: [
                        {
                            label: 'Hash 1m',
                            data: userHashrate1m,
                            borderColor: 'rgb(255, 99, 132)',
                            backgroundColor: 'rgba(255, 99, 132, 0.1)',
                            tension: 0.1
                        },
                        {
                            label: 'Hash 5m',
                            data: userHashrate5m,
                            borderColor: 'rgb(54, 162, 235)',
                            backgroundColor: 'rgba(54, 162, 235, 0.1)',
                            tension: 0.1
                        },
                        {
                            label: 'Hash 1h',
                            data: userHashrate1h,
                            borderColor: 'rgb(255, 205, 86)',
                            backgroundColor: 'rgba(255, 205, 86, 0.1)',
                            tension: 0.1
                        },
                        {
                            label: 'Hash 1d',
                            data: userHashrate1d,
                            borderColor: 'rgb(75, 192, 192)',
                            backgroundColor: 'rgba(75, 192, 192, 0.1)',
                            tension: 0.1
                        },
                        {
                            label: 'Hash 7d',
                            data: userHashrate7d,
                            borderColor: 'rgb(153, 102, 255)',
                            backgroundColor: 'rgba(153, 102, 255, 0.1)',
                            tension: 0.1
                        }
                    ]
                },
                options: {
                    responsive: true,
                    plugins: {
                        title: {
                            display: true,
                            text: 'Évolution du Hashrate Utilisateur'
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

            // Préparation des données pool (hashrate)
            const poolHashrateData = poolData.filter(item => item.type === 'hashrate').slice(0, 20);
            const poolLabels = poolHashrateData.map(item => item.timestamp.split(' ')[1]).reverse();
            const poolHashrate1m = poolHashrateData.map(item => parseHashrate(item.hashrate1m)).reverse();
            const poolHashrate5m = poolHashrateData.map(item => parseHashrate(item.hashrate5m)).reverse();
            const poolHashrate1h = poolHashrateData.map(item => parseHashrate(item.hashrate1hr)).reverse();
            const poolHashrate1d = poolHashrateData.map(item => parseHashrate(item.hashrate1d)).reverse();
            const poolHashrate7d = poolHashrateData.map(item => parseHashrate(item.hashrate7d)).reverse();

            // Graphique pool
            const poolCtx = document.getElementById('poolHashrateChart').getContext('2d');
            new Chart(poolCtx, {
                type: 'line',
                data: {
                    labels: poolLabels,
                    datasets: [
                        {
                            label: 'Hash 1m',
                            data: poolHashrate1m,
                            borderColor: 'rgb(255, 99, 132)',
                            backgroundColor: 'rgba(255, 99, 132, 0.1)',
                            tension: 0.1
                        },
                        {
                            label: 'Hash 5m',
                            data: poolHashrate5m,
                            borderColor: 'rgb(54, 162, 235)',
                            backgroundColor: 'rgba(54, 162, 235, 0.1)',
                            tension: 0.1
                        },
                        {
                            label: 'Hash 1h',
                            data: poolHashrate1h,
                            borderColor: 'rgb(255, 205, 86)',
                            backgroundColor: 'rgba(255, 205, 86, 0.1)',
                            tension: 0.1
                        },
                        {
                            label: 'Hash 1d',
                            data: poolHashrate1d,
                            borderColor: 'rgb(75, 192, 192)',
                            backgroundColor: 'rgba(75, 192, 192, 0.1)',
                            tension: 0.1
                        },
                        {
                            label: 'Hash 7d',
                            data: poolHashrate7d,
                            borderColor: 'rgb(153, 102, 255)',
                            backgroundColor: 'rgba(153, 102, 255, 0.1)',
                            tension: 0.1
                        }
                    ]
                },
                options: {
                    responsive: true,
                    plugins: {
                        title: {
                            display: true,
                            text: 'Évolution du Hashrate Pool'
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
