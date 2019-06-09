#
# __  __       _         __ _ _      
#|  \/  | __ _| | _____ / _(_) | ___ 
#| |\/| |/ _` | |/ / _ \ |_| | |/ _ \
#| |  | | (_| |   <  __/  _| | |  __/
#|_|  |_|\__,_|_|\_\___|_| |_|_|\___|
#

protoc=/usr/bin/protoc
go=/usr/bin/go

generate-stub:
	@$(protoc) -I /usr/local/include \
		-I $${PWD}/v1beta1 \
		-I $${GOPATH}/src \
		-I $${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--go_out=plugins=grpc:v1beta1 \
		$${PWD}/v1beta1/api.proto

generate-swagger:
	@$(protoc) -I /usr/local/include \
		-I $${PWD}/v1beta1 \
		-I $${GOPATH}/src \
		-I $${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--swagger_out=logtostderr=true:v1beta1 \
		$${PWD}/v1beta1/api.proto

generate-reverse-proxy:
	@$(protoc) -I /usr/local/include \
		-I $${PWD}/v1beta1 \
		-I $${GOPATH}/src \
		-I $${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--grpc-gateway_out=logtostderr=true:v1beta1 \
		$${PWD}/v1beta1/api.proto
