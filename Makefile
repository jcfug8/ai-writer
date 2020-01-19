# ALL
docker_compose_run: docker_all
	docker-compose up

docker_compose_run_d: docker_all
	docker-compose up -d

docker_all: docker_ai docker_persist docker_control

build_all: ai control persist

# AI

docker_ai: ai
	docker build -t jcfug8/ai_writer:ai_latest -f services/ai/Dockerfile .

run_ai: ai
	services/ai/ai_binary

ai: build_protos ./services/ai/cmd/*.go
	CGO_ENABLED=0 go build -o services/ai/ai_binary services/ai/cmd/cmd.go

# CONTROL

docker_control: control
	docker build -t jcfug8/ai_writer:control_latest -f services/control/Dockerfile .

run_control: control
	services/control/control_binary

control: build_protos ./services/control/cmd/*.go ./services/control/service/*.go
	CGO_ENABLED=0 go build -o services/control/control_binary services/control/cmd/cmd.go

# PERSIST

docker_persist: persist
	docker build -t jcfug8/ai_writer:persist_latest -f services/persist/Dockerfile .

run_persist: persist
	services/persist/persist_binary

persist: build_protos ./services/persist/cmd/*.go ./services/persist/service/*.go
	CGO_ENABLED=0 go build -o services/persist/persist_binary services/persist/cmd/cmd.go

# PROTOS

build_protos: ./protos/entities.proto
	protoc -I=protos --go_out=plugins=grpc:protos protos/entities.proto 

# CLEAN

clean:
	rm protos/entities.pb.go services/control/control_binary services/ai/ai_binary services/persist/persist_binary
