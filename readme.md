### SETUP ENV
```shell script
cd _docker
docker-compose up -d
```

### BUILD
```shell script
go build -o app
MYSQL_CONNECTION="root:garcolkey@tcp(127.0.0.1:3308)/garcol_food_delivery?charset=utf8mb4&parseTime=True&loc=Local" ./app
```