# sensu_role

Get information about a Sensu Role.

For full documentation on Sensu Roles bindings, see [here](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-bindings).

## Basic Example

```hcl
data "sensu_role_binding" "role_binding_1" {
  name = "role_binding_1"
}
```

## Argument Reference

* `name` - *Required* - The name of the Sensu role.

* `namespace` - *Optional* - The namespace to manage resources in. This can
  also be set with the `SENSU_NAMESPACE` environment variable. If not set,
  this defaults to `default`.

## Attribute Reference

* `binding_type` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-binding-and-cluster-role-binding-specification) for more information.

* `role` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-binding-and-cluster-role-binding-specification) for more information.

* `users` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-binding-and-cluster-role-binding-specification) for more information.

* `groups` - See the [Sensu rbac reference](https://docs.sensu.io/sensu-go/latest/operations/control-access/rbac/#role-binding-and-cluster-role-binding-specification) for more information.
