cd ..
# wget https://raw.githubusercontent.com/elan8/propanedb/master/protos/propanedb.proto -O ./api/propanedb.proto
#cd api
docker run --rm -v $(pwd):$(pwd) -w $(pwd) harbor.jevontech.com/dactory/builder \
/bin/sh -c "protoc  --go_out=:. --go-grpc_out=:. -I.  ./api/propanedb.proto"

# update file ownership
docker run --rm -v $(pwd):$(pwd) -w $(pwd)  harbor.jevontech.com/dactory/builder \
/bin/sh -c "chown -R $(id -u):$(id -g) pb"
