# ALL
docker_compose_run: docker_all
	docker-compose up

docker_compose_run_d: docker_all
	docker-compose up -d

docker_all: docker_ai docker_persist docker_control

build_all: control persist

# AI

docker_ai: build_protos
	docker build -t jcfug8/ai_writer:ai_latest services/ai/

# CONTROL

docker_control: control
	docker build -t jcfug8/ai_writer:control_latest services/control/

run_control: control
	services/control/control_binary

control: build_protos ./services/control/cmd/*.go ./services/control/service/*.go
	CGO_ENABLED=0 go build -o services/control/control_binary services/control/cmd/cmd.go

# PERSIST

docker_persist: persist
	docker build -t jcfug8/ai_writer:persist_latest services/persist/

run_persist: persist
	services/persist/persist_binary

persist: build_protos ./services/persist/cmd/*.go ./services/persist/service/*.go
	CGO_ENABLED=0 go build -o services/persist/persist_binary services/persist/cmd/cmd.go

# PROTOS

build_protos: ./protos/entities.proto
	protoc -I=protos --go_out=plugins=grpc:protos protos/entities.proto
	python3 -m grpc_tools.protoc -I=protos --python_out=services/ai/service --grpc_python_out=services/ai/service protos/entities.proto

# CLEAN

clean:
	rm protos/entities.pb.go services/control/control_binary services/ai/ai_binary services/persist/persist_binary
