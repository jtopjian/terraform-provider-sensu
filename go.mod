module github.com/jtopjian/terraform-provider-sensu

go 1.12

require (
	github.com/go-resty/resty v1.10.3 // indirect
	github.com/gorilla/context v0.0.0-20160226214623-1ea25387ff6f // indirect
	github.com/hashicorp/go-version v1.2.0
	github.com/hashicorp/terraform v0.12.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/robfig/cron v0.0.0-20180505203441-b41be1df6967 // indirect
	github.com/sensu/sensu-go v0.0.0-20200709195451-082f78fa286e
	github.com/sensu/sensu-go/api/core/v2 v2.0.0
	github.com/sensu/sensu-go/api/core/v3 v3.0.0-alpha2.0.20200709195451-082f78fa286e // indirect
	github.com/sensu/sensu-go/types v0.2.2-0.20200709195451-082f78fa286e
	github.com/whyrusleeping/go-logging v0.0.0-20170515211332-0457bb6b88fc // indirect
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.10.3
