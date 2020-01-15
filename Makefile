build: ai control persist
	echo built binaries

ai: protos
	go build -o services/ai/ai_binary services/ai/cmd/cmd.go

control: protos
	go build -o services/control/control_binary services/control/cmd/cmd.go

persist: protos
	go build -o services/persist/persist_binary services/persist/cmd/cmd.go

protos: protos/entities.proto
	protoc -I=protos --go_out=plugins=grpc:protos protos/entities.proto 
