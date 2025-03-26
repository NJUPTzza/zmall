cwgo client --type RPC --service user --module github.com/NJUPTzza/zmall/rpc_gen --I ../idl --idl ../idl/user.proto

go mod init github.com/NJUPTzza/zmall/app/user

cwgo server --type RPC --service user --module github.com/NJUPTzza/zmall/app/user --pass "-use github.com/NJUPTzza/zmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto


docker-compose up -d

docker-compose -f docker-compose-infra.yaml up -d