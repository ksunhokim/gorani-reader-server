for f in ../../proto/*.proto; do
	protoc -I="../../proto/" --go_out=plugins=grpc:"../../go/pkg/proto" "$f"
done


