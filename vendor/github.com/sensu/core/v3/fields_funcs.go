package v3

import (
	"strconv"
	"strings"

	corev2 "github.com/sensu/core/v2"
	"github.com/sensu/core/v3/internal/stringutil"
)

// APIKeyFields returns a set of fields that represent that resource.
func APIKeyFields(r Resource) map[string]string {
	resource := r.(*corev2.APIKey)
	fields := map[string]string{
		"api_key.name":     resource.ObjectMeta.Name,
		"api_key.username": resource.Username,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "api_key.labels.")
	return fields
}

// AssetFields returns a set of fields that represent that resource
func AssetFields(r Resource) map[string]string {
	resource := r.(*corev2.Asset)
	fields := map[string]string{
		"asset.name":      resource.ObjectMeta.Name,
		"asset.namespace": resource.ObjectMeta.Namespace,
		"asset.filters":   strings.Join(resource.Filters, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "asset.labels.")
	return fields
}

// CheckConfigFields returns a set of fields that represent that resource
func CheckConfigFields(r Resource) map[string]string {
	resource := r.(*corev2.CheckConfig)
	fields := map[string]string{
		"check.name":           resource.ObjectMeta.Name,
		"check.namespace":      resource.ObjectMeta.Namespace,
		"check.handlers":       strings.Join(resource.Handlers, ","),
		"check.publish":        strconv.FormatBool(resource.Publish),
		"check.round_robin":    strconv.FormatBool(resource.RoundRobin),
		"check.runtime_assets": strings.Join(resource.RuntimeAssets, ","),
		"check.subscriptions":  strings.Join(resource.Subscriptions, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "check.labels.")

	pipelineIDs := []string{}
	for _, pipeline := range resource.Pipelines {
		pipelineIDs = append(pipelineIDs, pipeline.ResourceID())
	}
	fields["check.pipelines"] = strings.Join(pipelineIDs, ",")

	return fields
}

// EntityFields returns a set of fields that represent that resource
func EntityFields(r Resource) map[string]string {
	resource := r.(*corev2.Entity)
	fields := map[string]string{
		"entity.name":          resource.ObjectMeta.Name,
		"entity.namespace":     resource.ObjectMeta.Namespace,
		"entity.deregister":    strconv.FormatBool(resource.Deregister),
		"entity.entity_class":  resource.EntityClass,
		"entity.subscriptions": strings.Join(resource.Subscriptions, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "entity.labels.")
	return fields
}

func isSilenced(e *corev2.Event) string {
	if len(e.Check.Silenced) > 0 {
		return "true"
	}
	return "false"
}

// EventFields returns a set of fields that represent that resource
func EventFields(r Resource) map[string]string {
	resource := r.(*corev2.Event)
	fields := map[string]string{
		"event.name":                 resource.ObjectMeta.Name,
		"event.namespace":            resource.ObjectMeta.Namespace,
		"event.is_silenced":          isSilenced(resource),
		"event.check.is_silenced":    isSilenced(resource),
		"event.check.name":           resource.Check.Name,
		"event.check.handlers":       strings.Join(resource.Check.Handlers, ","),
		"event.check.publish":        strconv.FormatBool(resource.Check.Publish),
		"event.check.round_robin":    strconv.FormatBool(resource.Check.RoundRobin),
		"event.check.runtime_assets": strings.Join(resource.Check.RuntimeAssets, ","),
		"event.check.state":          resource.Check.State,
		"event.check.status":         strconv.Itoa(int(resource.Check.Status)),
		"event.check.subscriptions":  strings.Join(resource.Check.Subscriptions, ","),
		"event.entity.deregister":    strconv.FormatBool(resource.Entity.Deregister),
		"event.entity.name":          resource.Entity.ObjectMeta.Name,
		"event.entity.entity_class":  resource.Entity.EntityClass,
		"event.entity.subscriptions": strings.Join(resource.Entity.Subscriptions, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "event.labels.")
	stringutil.MergeMapWithPrefix(fields, resource.Entity.ObjectMeta.Labels, "event.labels.")
	stringutil.MergeMapWithPrefix(fields, resource.Check.ObjectMeta.Labels, "event.labels.")
	return fields
}

// EventFilterFields returns a set of fields that represent that resource
func EventFilterFields(r Resource) map[string]string {
	resource := r.(*corev2.EventFilter)
	fields := map[string]string{
		"filter.name":           resource.ObjectMeta.Name,
		"filter.namespace":      resource.ObjectMeta.Namespace,
		"filter.action":         resource.Action,
		"filter.runtime_assets": strings.Join(resource.RuntimeAssets, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "filter.labels.")
	return fields
}

// HandlerFields returns a set of fields that represent that resource
func HandlerFields(r Resource) map[string]string {
	resource := r.(*corev2.Handler)
	fields := map[string]string{
		"handler.name":      resource.ObjectMeta.Name,
		"handler.namespace": resource.ObjectMeta.Namespace,
		"handler.filters":   strings.Join(resource.Filters, ","),
		"handler.handlers":  strings.Join(resource.Handlers, ","),
		"handler.mutator":   resource.Mutator,
		"handler.type":      resource.Type,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "handler.labels.")
	return fields
}

// MutatorFields returns a set of fields that represent that resource
func MutatorFields(r Resource) map[string]string {
	resource := r.(*corev2.Mutator)
	fields := map[string]string{
		"mutator.name":           resource.ObjectMeta.Name,
		"mutator.namespace":      resource.ObjectMeta.Namespace,
		"mutator.runtime_assets": strings.Join(resource.RuntimeAssets, ","),
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "mutator.labels.")
	return fields
}

// SilencedFields returns a set of fields that represent that resource
func SilencedFields(r Resource) map[string]string {
	resource := r.(*corev2.Silenced)
	fields := map[string]string{
		"silenced.name":              resource.ObjectMeta.Name,
		"silenced.namespace":         resource.ObjectMeta.Namespace,
		"silenced.check":             resource.Check,
		"silenced.creator":           resource.Creator,
		"silenced.expire_on_resolve": strconv.FormatBool(resource.ExpireOnResolve),
		"silenced.subscription":      resource.Subscription,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "silenced.labels.")
	return fields
}

// HookConfigFields returns a set of fields that represent that resource
func HookConfigFields(r Resource) map[string]string {
	resource := r.(*corev2.HookConfig)
	fields := map[string]string{
		"hook.name":      resource.ObjectMeta.Name,
		"hook.namespace": resource.ObjectMeta.Namespace,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "hook.labels.")
	return fields
}

// PipelineFields returns a set of fields that represent that resource.
func PipelineFields(r Resource) map[string]string {
	resource := r.(*corev2.Pipeline)
	fields := map[string]string{
		"pipeline.name":      resource.ObjectMeta.Name,
		"pipeline.namespace": resource.ObjectMeta.Namespace,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "pipeline.labels.")
	return fields
}

// UserFields returns a set of fields that represent that resource
func UserFields(r Resource) map[string]string {
	resource := r.(*corev2.User)
	return map[string]string{
		"user.username": resource.Username,
		"user.disabled": strconv.FormatBool(resource.Disabled),
		"user.groups":   strings.Join(resource.Groups, ","),
	}
}

// NamespaceFields returns a set of fields that represent that resource
func NamespaceFields(r Resource) map[string]string {
	resource := r.(*Namespace)
	return map[string]string{
		"namespace.name": resource.Metadata.Name,
	}
}

// RoleFields returns a set of fields that represent that resource
func RoleFields(r Resource) map[string]string {
	resource := r.(*corev2.Role)
	fields := map[string]string{
		"role.name":      resource.ObjectMeta.Name,
		"role.namespace": resource.ObjectMeta.Namespace,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "role.labels.")
	return fields
}

// ClusterRoleFields returns a set of fields that represent that resource
func ClusterRoleFields(r Resource) map[string]string {
	resource := r.(*corev2.ClusterRole)
	fields := map[string]string{
		"clusterrole.name": resource.ObjectMeta.Name,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "clusterrole.labels.")
	return fields
}

// ClusterRoleBindingFields returns a set of fields that represent that resource
func ClusterRoleBindingFields(r Resource) map[string]string {
	resource := r.(*corev2.ClusterRoleBinding)
	fields := map[string]string{
		"clusterrolebinding.name":          resource.ObjectMeta.Name,
		"clusterrolebinding.role_ref.name": resource.RoleRef.Name,
		"clusterrolebinding.role_ref.type": resource.RoleRef.Type,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "clusterrolebinding.labels.")
	return fields
}

// RoleBindingFields returns a set of fields that represent that resource
func RoleBindingFields(r Resource) map[string]string {
	resource := r.(*corev2.RoleBinding)
	fields := map[string]string{
		"rolebinding.name":          resource.ObjectMeta.Name,
		"rolebinding.namespace":     resource.ObjectMeta.Namespace,
		"rolebinding.role_ref.name": resource.RoleRef.Name,
		"rolebinding.role_ref.type": resource.RoleRef.Type,
	}
	stringutil.MergeMapWithPrefix(fields, resource.ObjectMeta.Labels, "rolebinding.labels.")
	return fields
}
