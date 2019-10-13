module github.com/yulpa/yulmails

go 1.12

require (
	cloud.google.com/go v0.40.0 // indirect
	github.com/golang/mock v1.3.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/pprof v0.0.0-20190515194954-54271f7e092f // indirect
	github.com/grpc-ecosystem/grpc-gateway v1.9.0 // indirect
	github.com/kr/pty v1.1.4 // indirect
	github.com/lib/pq v1.1.1 // indirect
	github.com/rogpeppe/fastuuid v1.1.0 // indirect
	gitlab.com/tortuemat/yulmails/cmd v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/tortuemat/yulmails/services/entrypoint v0.0.0-00010101000000-000000000000 // indirect
	go.opencensus.io v0.22.0 // indirect
	golang.org/x/exp v0.0.0-20190510132918-efd6b22b2522 // indirect
	golang.org/x/image v0.0.0-20190523035834-f03afa92d3ff // indirect
	golang.org/x/mobile v0.0.0-20190607214518-6fa95d984e88 // indirect
	golang.org/x/mod v0.1.0 // indirect
	golang.org/x/net v0.0.0-20190607181551-461777fb6f67 // indirect
	golang.org/x/sys v0.0.0-20190609082536-301114b31cce // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	golang.org/x/tools v0.0.0-20190608022120-eacb66d2a7c3 // indirect
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/genproto v0.0.0-20190605220351-eb0b1bdb6ae6 // indirect
	google.golang.org/grpc v1.21.1 // indirect
	honnef.co/go/tools v0.0.0-20190607181801-497c8f037f5a // indirect
)

replace (
	gitlab.com/tortuemat/yulmails/cmd => ./cmd
	gitlab.com/tortuemat/yulmails/services/conservation => ./services/conservation
	gitlab.com/tortuemat/yulmails/services/entrypoint => ./services/entrypoint
)
