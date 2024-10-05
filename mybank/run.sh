echo "1) Start DB"
echo "2) Drop DB"
echo "3) Start server"
echo "4) Stop server"
echo "5) Stop server and db"

read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
    docker compose -f docker-compose.yml up -d
elif [[ $cmd == 2 ]]; then
    docker compose -f docker-compose.yml down
    docker rm db
    docker pgadmin
elif [[ $cmd == 3 ]]; then
    docker compose -f docker-compose.yml up -d
    go run main.go
elif [[ $cmd == 4 ]]; then
    sudo kill -9 $(sudo lsof -t -i:8080)
elif [[ $cmd == 5 ]]; then
    sudo kill -9 $(sudo lsof -t -i:8080)
    docker compose -f docker-compose.yml down
    docker rm db
    docker rm pgadmin
fi
