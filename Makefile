proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./internal/pkg/pb/ports.proto


build: 
	docker build -f data.Dockerfile -t data .
	docker build -f domain.Dockerfile -t domain .

run: 
	docker run -it -d  --network=host  mongo
	docker run -it --network=host -d data
	docker run -it --network=host -d domain