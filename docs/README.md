# terraform-provider-sensu

Use Terraform to manage Sensu resources.

### Resources and Data Sources

* A list of supported resources can be found [here](resources).
* A list of supported data sources can be found [here](data_sources).

## Basic Example

To configure this provider, do the following:

```hcl
provider "sensu" {
  api_url   = "http://127.0.0.1:8080"
  username  = "admin"
  password  = "password"
  namespace = "default"
}
```

## Configuration Reference

The following arguments are supported:

* `api_url` - *Required* - The URL to the Sensu service. This can also be set
  with the `SENSU_API_URL` environment variable.

* `username` - *Required* - The username to connect to the Sensu service as.
  This can also be set with the `SENSU_USERNAME` environment variable.

* `password` - *Required* - The password to authenticate to the Sensu service
  with. This can also be set with the `SENSU_PASSWORD` environment variable.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `edition` - *Optional* - The edition of the Sensu service. This can also
  be set with the `SENSU_EDITION` environment variable.
