module github.com/liyinda/ingress-admin/backend

go 1.13

replace github.com/coreos/go-systemd/journal => /root/Projects/go/src/github.com/coreos/go-systemd/journal

require (
	github.com/boj/redistore v0.0.0-20180917114910-cd5dcc76aeff // indirect
	github.com/coreos/etcd v3.3.18+incompatible // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gin-contrib/cors v1.3.0
	github.com/gin-gonic/contrib v0.0.0-20191209060500-d6e26eeaa607
	github.com/gin-gonic/gin v1.5.0
	github.com/go-ini/ini v1.52.0
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/google/uuid v1.1.1 // indirect
	github.com/gorilla/sessions v1.2.0 // indirect
	github.com/liyinda/viewdns v0.0.0-20191119085041-d6adbafaaee7
	github.com/tidwall/gjson v1.6.0
	go.etcd.io/etcd v3.3.18+incompatible
	go.uber.org/zap v1.14.0 // indirect
	google.golang.org/grpc v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.2.2
)
