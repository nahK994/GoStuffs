echo "1) start process"
echo "2) kill process"

read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
    docker compose -f docker-compose.yml up -d
    sleep 2
    go run cmd/web/main.go
elif [[ $cmd == 2 ]]; then
    sudo kill -9 $(sudo lsof -t -i:8000)
    docker stop db
    docker rm db
fi
