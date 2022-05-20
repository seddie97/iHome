module getCaptcha

go 1.15

require (
	github.com/afocus/captcha v0.0.0-20191010092841-4bd1f21c8868
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/golang/freetype v0.0.0-20170609003504-e2365dfdc4a0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.5.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/google/btree v1.0.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.5 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lib/pq v1.3.0 // indirect
	github.com/lucas-clemente/quic-go v0.14.1 // indirect
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/registry/consul v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/micro/v3 v3.3.0
	github.com/nats-io/nats-server/v2 v2.1.6 // indirect
	github.com/tmc/grpc-websocket-proxy v0.0.0-20200122045848-3419fae592fc // indirect
	go.uber.org/zap v1.13.0 // indirect
	google.golang.org/protobuf v1.27.1
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
