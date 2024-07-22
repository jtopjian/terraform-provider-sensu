## 0.15.0 (Unreleased)

## 0.14.0 (July 22, 2024)

IMPROVEMENTS

* Added support for "pipelines" to the `check` resource [GH-63](https://github.com/jtopjian/terraform-provider-sensu/pull/63)

## 0.13.0 (June 21, 2023)

IMPROVEMENTS

* Added support for "secrets" [GH-57](https://github.com/jtopjian/terraform-provider-sensu/pull/57)

## 0.12.2 (February 26, 2023)

OTHER

* Updating dependencies [GH-54](https://github.com/jtopjian/terraform-provider-sensu/pull/54)
* Updating dependencies [GH-55](https://github.com/jtopjian/terraform-provider-sensu/pull/55)

## 0.12.1 (September 12, 2021)

BUG FIXES

* Fixed crash when updating hooks in `sensu_check` [GH-52](https://github.com/jtopjian/terraform-provider-sensu/pull/52)

## 0.12.0 (August 6th, 2021)

IMPROVEMENTS

* Added `sensu_apikey` [GH-48](https://github.com/jtopjian/terraform-provider-sensu/pull/48)

## 0.11.0 (May 30, 2021)

NOTES

* The `edition` argument was removed. You will want to make sure you remove this from all of your configurations, if it ever existed.

IMPROVEMENTS

* Headers can be defined at top level or in each build of `sensu_asset` [GH-44](https://github.com/jtopjian/terraform-provider-sensu/pull/44)
* Removed `edition` [GH-47](https://github.com/jtopjian/terraform-provider-sensu/pull/47)
* Support API key authentication [GH-46](https://github.com/jtopjian/terraform-provider-sensu/pull/46)

## 0.10.6 (October 31, 2020)

BUG FIXES

* `sensu_check.env_vars` can now be updated [GH-42](https://github.com/jtopjian/terraform-provider-sensu/pull/42)

## 0.10.5 (September 5, 2020)

No changes. This release was only created to be published to the Terraform Registry.

## 0.10.4 (July 14, 2020)

BUG FIXES

* Don't create an empty `sensu_check.proxy_requests` when one isn't specified in Terraform [GH-38](https://github.com/jtopjian/terraform-provider-sensu/pull/38)

## 0.10.3 (July 13, 2020)

BUG FIXES

* Fixed nil pointer bug in `sensu_check.proxy_requests` [GH-37](https://github.com/jtopjian/terraform-provider-sensu/pull/37)

## 0.10.2 (July 13, 2020)

BUG FIXES

* Added `sensu_bonsai_asset.build` and deprecated `sensu_bonsai_asset.builds` for naming consistency [GH-36](https://github.com/jtopjian/terraform-provider-sensu/pull/36)


## 0.10.1 (July 11, 2020)

BUG FIXES

* Fixed missing update logic preventing asset builds from being updated [GH-35](https://github.com/jtopjian/terraform-provider-sensu/pull/35)

## 0.10.0 (July 10, 2020)

IMPROVEMENTS

* Added `sensu_entity` resource [GH-32](https://github.com/jtopjian/terraform-provider-sensu/pull/32)
* Added `proxy_requests` to the `sensu_check` resource [GH-32](https://github.com/jtopjian/terraform-provider-sensu/pull/32)
* Added `proxy_requests` to the `sensu_check` data source [GH-33](https://github.com/jtopjian/terraform-provider-sensu/pull/33)
* Added `sensu_bonsai_asset` data source [GH-30](https://github.com/jtopjian/terraform-provider-sensu/pull/30)

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
