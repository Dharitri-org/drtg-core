//go:generate protoc -I=. -I=$GOPATH/src -I=$GOPATH/src/github.com/Dharitri-org/protobuf/protobuf  --gogoslick_out=$GOPATH/src alteredAccount.proto
package alteredAccount
