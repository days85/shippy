module github.com/days85/shippy/shippy-cli-consignment

go 1.15

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

//replace github.com/days85/shippy/shippy-service-vessel => /Users/nunodias/dev/go/code/src/github.com/days85/shippy/shippy-service-vessel
//replace github.com/days85/shippy/shippy-service-consignment => /Users/nunodias/dev/go/code/src/github.com/days85/shippy/shippy-service-consignment

require (
	github.com/coreos/etcd v3.3.25+incompatible // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/days85/shippy/shippy-service-consignment v0.0.0-20201017235231-53a04e4d33cd
	github.com/days85/shippy/shippy-service-vessel v0.0.0-20201017235231-53a04e4d33cd
	github.com/dustin/go-humanize v1.0.0 // indirect
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-git/go-git/v5 v5.2.0 // indirect
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.2 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/miekg/dns v1.1.33 // indirect
	github.com/nats-io/jwt v1.0.1 // indirect
	github.com/nats-io/nats.go v1.10.0 // indirect
	github.com/nats-io/nkeys v0.2.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.16.0 // indirect
	golang.org/x/crypto v0.0.0-20201016220609-9e8e0b390897 // indirect
	golang.org/x/lint v0.0.0-20200302205851-738671d3881b // indirect
	golang.org/x/net v0.0.0-20201016165138-7b1cca2348c0 // indirect
	golang.org/x/sync v0.0.0-20201008141435-b3e1573b7520 // indirect
	golang.org/x/sys v0.0.0-20201017003518-b09fb700fbb7 // indirect
	golang.org/x/tools v0.0.0-20201015182029-a5d9e455e9c4 // indirect
	google.golang.org/genproto v0.0.0-20201015140912-32ed001d685c // indirect
	google.golang.org/grpc v1.33.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.6 // indirect
)
