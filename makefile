cwgo server --type RPC --service user --module github.com/NJUPTzza/zmall/app/user --pass "-use github.com/NJUPTzza/zmall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto

cwgo client --type RPC --service user --module github.com/NJUPTzza/zmall/rpc_gen --I ../idl --idl ../idl/user.proto