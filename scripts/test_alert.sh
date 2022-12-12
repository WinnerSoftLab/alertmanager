#!/bin/bash
#set -euo pipefail

curl -H 'Content-Type: application/json' -d '[{"labels":{"alertname":"Test Alert", "severity": "critical", "env": "prod", "host_name": "testmachine", "weight": "10"}}]' http://127.0.0.1:9093/api/v1/alerts
curl -H 'Content-Type: application/json' -d '[{"labels":{"alertname":"Test Alert", "severity": "critical", "env": "prod", "host_name": "testmachine2", "weight": "10"}}]' http://127.0.0.1:9093/api/v1/alerts
curl -H 'Content-Type: application/json' -d '[{"labels":{"alertname":"Test Alert 2", "severity": "critical", "env": "prod", "host_name": "testmachine", "weight": "10"}}]' http://127.0.0.1:9093/api/v1/alerts
