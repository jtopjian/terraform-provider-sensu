module github.com/jtopjian/terraform-provider-sensu

go 1.14

require (
	github.com/gorilla/context v0.0.0-20160226214623-1ea25387ff6f // indirect
	github.com/hashicorp/go-version v1.2.0
	github.com/hashicorp/terraform v0.12.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/sensu/sensu-go v0.0.0-20210726180517-2bbdde848469
	github.com/sensu/sensu-go/api/core/v2 v2.9.0
	github.com/sensu/sensu-go/types v0.7.0
	github.com/whyrusleeping/go-logging v0.0.0-20170515211332-0457bb6b88fc // indirect
	google.golang.org/grpc v1.39.0 // indirect
)

replace google.golang.org/grpc => google.golang.org/grpc v1.29.1
