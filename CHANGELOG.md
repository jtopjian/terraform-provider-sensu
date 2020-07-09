## 0.10.0 (Unreleased)

IMPROVEMENTS

* Added `sensu_entity` resource [GH-32](https://github.com/jtopjian/terraform-provider-sensu/pull/32)
* Added `proxy_requests` to the `sensu_check` resource [GH-32](https://github.com/jtopjian/terraform-provider-sensu/pull/32)

BUG FIXES

* Removed `keepalive_timeout` from the `sensu_entity` data source [GH-32](https://github.com/jtopjian/terraform-provider-sensu/pull/32)


## 0.9.0 (July 3, 2020)

IMPROVEMENTS

* Added `sensu_cluster_role` resource [GH-29](https://github.com/jtopjian/terraform-provider-sensu/pull/29)
* Added `sensu_cluster_role` data source [GH-29](https://github.com/jtopjian/terraform-provider-sensu/pull/29)
* Added `sensu_cluster_role_binding` resource [GH-29](https://github.com/jtopjian/terraform-provider-sensu/pull/29)
* Added `sensu_cluster_role_binding` data source [GH-29](https://github.com/jtopjian/terraform-provider-sensu/pull/29)
* Added multi-build support for the Asset resource and data source [GH-31](https://github.com/jtopjian/terraform-provider-sensu/pull/31)

## 0.8.0 (June 4, 2020)

IMPROVEMENTS

* Added `sensu_filter.runtime_assets` [GH-20](https://github.com/jtopjian/terraform-provider-sensu/pull/20)
* Added `sensu_asset.headers` [GH-21](https://github.com/jtopjian/terraform-provider-sensu/pull/21)
* Enabled `sensu_asset`s to be truly deleted instead of just deleted in the Terrafom State [GH-22](https://github.com/jtopjian/terraform-provider-sensu/pull/22)
* Added `sensu_silenced` resource [GH-25](https://github.com/jtopjian/terraform-provider-sensu/pull/25)
* Added `sensu_silenced` data source [GH-25](https://github.com/jtopjian/terraform-provider-sensu/pull/25)

BUG FIXES

* Fix typo with `sensu_check.output_metric_format` [GH-19](https://github.com/jtopjian/terraform-provider-sensu/pull/19)
* Fix empty `when` blocks from being created in `sensu_filter` resource [GH-26](https://github.com/jtopjian/terraform-provider-sensu/pull/26)

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
