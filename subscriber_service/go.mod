module github.com/rezaAmiri123/nov-test/subscriber_service

go 1.18

replace github.com/rezaAmiri123/nov-test/pkg => ../pkg

replace github.com/rezaAmiri123/nov-test/publisher_service => ../publisher_service

require (
	github.com/google/uuid v1.3.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/nats-io/stan.go v0.10.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.13.0
	github.com/rezaAmiri123/microservice/pkg v0.0.0-20221130075717-a145d34a64e9
	github.com/rezaAmiri123/nov-test/pkg v0.0.0-00010101000000-000000000000
	github.com/rezaAmiri123/nov-test/publisher_service v0.0.0-20221222084548-69ec73b7dade
	github.com/rezaAmiri123/test-microservice/pkg v0.0.0-20221130073457-365c542ef312
	github.com/spf13/viper v1.13.0
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/fsnotify/fsnotify v1.5.4 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/jackc/pgx v3.6.2+incompatible // indirect
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/nats-io/nats.go v1.19.0 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pelletier/go-toml v1.9.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/rogpeppe/go-internal v1.8.0 // indirect
	github.com/segmentio/kafka-go v0.4.35 // indirect
	github.com/spf13/afero v1.9.2 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.1 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.8.0 // indirect
	go.uber.org/zap v1.24.0 // indirect
	golang.org/x/crypto v0.0.0-20221010152910-d6f0a8c073c2 // indirect
	golang.org/x/sys v0.0.0-20220926163933-8cfa568d3c25 // indirect
	golang.org/x/text v0.4.0 // indirect
	google.golang.org/grpc v1.51.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
