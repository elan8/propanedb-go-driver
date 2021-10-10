
# docker build -t propanedb-go-driver/test:latest  -f ./test/Dockerfile .

cd ./test
go test
# docker-compose up 
# sleep 5
#go test
# docker-compose down