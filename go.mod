module github.com/jtopjian/terraform-provider-sensu

go 1.14

require (
	github.com/hashicorp/go-getter v1.7.0 // indirect
	github.com/hashicorp/go-version v1.6.0
	github.com/hashicorp/terraform-plugin-sdk v1.16.0
	github.com/hashicorp/terraform-plugin-test/v2 v2.1.3 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/rogpeppe/go-internal v1.9.0 // indirect
	github.com/sensu/core/v2 v2.17.0
	github.com/sensu/core/v3 v3.8.0-beta4 // indirect
	github.com/sensu/sensu-go v0.0.0-20221027183945-cd8638430844
	github.com/sensu/sensu-go/types v0.12.0-alpha7
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/net v0.17.0 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)

replace github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.7.0

replace github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.34.0
