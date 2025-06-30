echo "1) Start server"
echo "2) Kill port"

get_port() {
    read -p "port: " port
    echo $port
}


read -p "Type: " cmd
if [[ $cmd == 1 ]]; then
    go run main.go
elif [[ $cmd == 2 ]]; then
    port=$(get_port)
    sudo kill -9 `sudo lsof -t -i:$port`
fi