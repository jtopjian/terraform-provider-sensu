# sensu_apikey

Manages a Sensu API Key.

For full documentation on Sensu API Key, see [here](https://docs.sensu.io/sensu-go/latest/operations/control-access/apikeys/).

## Basic Example

```hcl
resource "sensu_apikey" "apikey_1" {
  username = "foo"
}
```

## Argument Reference

* `username` - *Required* - The name of the user for whom the api key will be generated

## Attribute Reference

The resource has no computed fields.