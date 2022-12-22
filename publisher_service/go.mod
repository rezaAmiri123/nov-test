module github.com/rezaAmiri123/nov-test/publisher_service

go 1.18

replace github.com/rezaAmiri123/nov-test/pkg => ../pkg

require (
	github.com/opentracing/opentracing-go v1.2.0
	github.com/rezaAmiri123/nov-test/pkg v0.0.0-00010101000000-000000000000
)

require (
	github.com/klauspost/compress v1.15.9 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	github.com/segmentio/kafka-go v0.4.38 // indirect
)
