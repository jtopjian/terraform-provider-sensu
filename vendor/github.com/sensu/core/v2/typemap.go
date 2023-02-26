package v2

// automatically generated file, do not edit!

import (
	"github.com/sensu/sensu-api-tools"
)

const apiGroup = "core/v2"

var typeMap = map[string]interface{}{
	"api_key":                &APIKey{},
	"adhoc_request":          &AdhocRequest{},
	"any":                    &Any{},
	"asset":                  &Asset{},
	"asset_build":            &AssetBuild{},
	"asset_list":             &AssetList{},
	"auth_provider_claims":   &AuthProviderClaims{},
	"check":                  &Check{},
	"check_config":           &CheckConfig{},
	"check_history":          &CheckHistory{},
	"check_request":          &CheckRequest{},
	"claims":                 &Claims{},
	"cluster_health":         &ClusterHealth{},
	"cluster_role":           &ClusterRole{},
	"cluster_role_binding":   &ClusterRoleBinding{},
	"deregistration":         &Deregistration{},
	"entity":                 &Entity{},
	"event":                  &Event{},
	"event_filter":           &EventFilter{},
	"extension":              &Extension{},
	"handler":                &Handler{},
	"handler_socket":         &HandlerSocket{},
	"health_response":        &HealthResponse{},
	"hook":                   &Hook{},
	"hook_config":            &HookConfig{},
	"hook_list":              &HookList{},
	"keepalive_record":       &KeepaliveRecord{},
	"metric_point":           &MetricPoint{},
	"metric_tag":             &MetricTag{},
	"metric_threshold":       &MetricThreshold{},
	"metric_threshold_rule":  &MetricThresholdRule{},
	"metric_threshold_tag":   &MetricThresholdTag{},
	"metrics":                &Metrics{},
	"mutator":                &Mutator{},
	"namespace":              &Namespace{},
	"network":                &Network{},
	"network_interface":      &NetworkInterface{},
	"object_meta":            &ObjectMeta{},
	"pipeline":               &Pipeline{},
	"pipeline_workflow":      &PipelineWorkflow{},
	"postgres_health":        &PostgresHealth{},
	"process":                &Process{},
	"proxy_requests":         &ProxyRequests{},
	"resource_reference":     &ResourceReference{},
	"role":                   &Role{},
	"role_binding":           &RoleBinding{},
	"role_ref":               &RoleRef{},
	"rule":                   &Rule{},
	"secret":                 &Secret{},
	"silenced":               &Silenced{},
	"subject":                &Subject{},
	"system":                 &System{},
	"tls_options":            &TLSOptions{},
	"tessen_config":          &TessenConfig{},
	"time_window_days":       &TimeWindowDays{},
	"time_window_repeated":   &TimeWindowRepeated{},
	"time_window_time_range": &TimeWindowTimeRange{},
	"time_window_when":       &TimeWindowWhen{},
	"tokens":                 &Tokens{},
	"type_meta":              &TypeMeta{},
	"user":                   &User{},
	"version":                &Version{},
}

func init() {
	for typeAlias, typ := range typeMap {
		if _, ok := typ.(Resource); ok {
			apitools.RegisterType(apiGroup, typ, apitools.WithAlias(typeAlias))
		}
	}
}
