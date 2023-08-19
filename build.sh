rm -Rf generated build
mkdir generated

cd proto
protoc --go_out=../generated --go_opt=paths=source_relative \
    --go-grpc_out=../generated --go-grpc_opt=paths=source_relative \
    *.proto

cd ..
echo "Building binary.....[+]"
go build -mod=mod -o build/ .
echo "Completed Successfully."