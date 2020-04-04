## 0.8.0 (Unreleased)

BUG FIXES

* Fix typo with `sensu_check.output_metric_format` [GH-19](https://github.com/jtopjian/terraform-provider-sensu/pull/19)

## 0.7.0 (March 29, 2020)

IMPROVEMENTS

* Added `trusted_ca_file` to the provider [GH-18](https://github.com/jtopjian/terraform-provider-sensu/pull/18)
* Added `insecure_skip_tls_verify` to the provider [GH-18](https://github.com/jtopjian/terraform-provider-sensu/pull/18)

## 0.6.0 (March 26, 2020)

IMPROVEMENTS

* Added `disabled` to `sensu_user` [GH-11](https://github.com/jtopjian/terraform-provider-sensu/pull/11)
* Added `labels` to `sensu_check` [GH-10](https://github.com/jtopjian/terraform-provider-sensu/pull/10)
* Added `annotations` to `sensu_check` [GH-10](https://github.com/jtopjian/terraform-provider-sensu/pull/10)

BUG FIXES

* Fixed possible panic with `sensu_check.handlers` [GH-13](https://github.com/jtopjian/terraform-provider-sensu/pull/13)

## 0.5.0 (May 31, 2019)

IMPROVEMENTS

* Support for Terraform v0.12 [GH-7](https://github.com/jtopjian/terraform-provider-sensu/pull/7)
* Support for static binaries [GH-7](https://github.com/jtopjian/terraform-provider-sensu/pull/7)

## 0.4.0 (March 27, 2019)

IMPROVEMENTS

* Added `runtime_assets` to `sensu_handler` [GH-5](https://github.com/jtopjian/terraform-provider-sensu/pull/5)
* Added `runtime_assets` to `sensu_handler` data source [GH-5](https://github.com/jtopjian/terraform-provider-sensu/pull/5)
