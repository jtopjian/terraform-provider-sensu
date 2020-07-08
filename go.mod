module github.com/jtopjian/terraform-provider-sensu

go 1.12

require (
	github.com/go-resty/resty v1.10.3 // indirect
	github.com/hashicorp/terraform v0.12.0 // indirect
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/robfig/cron v0.0.0-20180505203441-b41be1df6967 // indirect
	github.com/sensu/sensu-go v0.0.0-20200310183930-0c45f08323ae
)

replace github.com/go-resty/resty => gopkg.in/resty.v1 v1.10.3
