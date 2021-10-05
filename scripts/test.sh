cd ./test
docker-compose up 
sleep 5
go test
docker-compose down