# terraform-provider-sensu

Use Terraform to manage Sensu resources.

### Resources and Data Sources

* A list of supported resources can be found [here](resources).
* A list of supported data sources can be found [here](data_sources).

## Basic Example

To configure this provider, do the following:

With api key:
```hcl
provider "sensu" {
  api_url   = "http://127.0.0.1:8080"
  api_key   = "admin"
  namespace = "default"
}
```

With username/password:
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

* `api_key` - *Optional* - The API KEY to authenticate to the Sensu service. This can also be set
  with the `SENSU_API_KEY` environment variable.

* `username` - *Optional* - The username to connect to the Sensu service as.
  This can also be set with the `SENSU_USERNAME` environment variable.

* `password` - *Optional* - The password to authenticate to the Sensu service
  with. This can also be set with the `SENSU_PASSWORD` environment variable.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

* `edition` - *Optional* - The edition of the Sensu service. This can also
  be set with the `SENSU_EDITION` environment variable.
  
* `trusted_ca_file` - *Optional* - Specify a file to be used as a trusted CA
  certificate. This can also be set with the `SENSU_TRUSTED_CA_FILE` environment
  variable.
  
* `insecure_skip_tls_verify` - *Optional* - Skip TLS verification. This is
  commonly used when communicating with self-signed certificates. This can also
  be set with thet `SENSU_INSECURE_SKIP_TLS_VERIFY` environment variable.
