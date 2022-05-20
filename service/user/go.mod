module user

go 1.15

require (
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/golang/protobuf v1.4.3
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/btree v1.0.0 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro/v3 v3.0.0
	google.golang.org/protobuf v1.25.0
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
