cd ..

FILE=./api/propanedb.proto
if [ -f "$FILE" ]; then
    echo "$FILE exists."
else 
    echo "$FILE does not exist."
    wget https://raw.githubusercontent.com/elan8/propanedb/master/protos/propanedb.proto -O ./api/propanedb.proto
fi

FILE=./api/test.proto
if [ -f "$FILE" ]; then
    echo "$FILE exists."
else 
    echo "$FILE does not exist."
    wget https://raw.githubusercontent.com/elan8/propanedb/master/protos/test.proto -O ./api/test.proto
fi

docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine  \
/bin/sh -c "protoc  --go_out=:. --go-grpc_out=:. -I.  ./api/propanedb.proto"

docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine  \
/bin/sh -c "protoc  --go_out=:. --go-grpc_out=:. --descriptor_set_out=./pb/test.bin -I.  ./api/test.proto"

# update file ownership
docker run --rm -v $(pwd):$(pwd) -w $(pwd) jevon82/golang-builder-alpine  \
/bin/sh -c "chown -R $(id -u):$(id -g) pb"
