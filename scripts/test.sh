cd ./development_test
docker-compose up -d  
sleep 2
go test 
docker-compose down