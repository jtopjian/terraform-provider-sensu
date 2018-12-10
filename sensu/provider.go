package sensu

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a schema.Provider for Sensu
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENSU_API_URL", ""),
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENSU_USERNAME", ""),
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("SENSU_PASSWORD", ""),
			},

			"namespace": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENSU_NAMESPACE", ""),
			},

			"edition": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("SENSU_EDITION", ""),
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"sensu_asset":        dataSourceAsset(),
			"sensu_check":        dataSourceCheck(),
			"sensu_entity":       dataSourceEntity(),
			"sensu_filter":       dataSourceFilter(),
			"sensu_handler":      dataSourceHandler(),
			"sensu_hook":         dataSourceHook(),
			"sensu_mutator":      dataSourceMutator(),
			"sensu_namespace":    dataSourceNamespace(),
			"sensu_role":         dataSourceRole(),
			"sensu_role_binding": dataSourceRoleBinding(),
			"sensu_user":         dataSourceUser(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"sensu_asset":        resourceAsset(),
			"sensu_check":        resourceCheck(),
			"sensu_filter":       resourceFilter(),
			"sensu_handler":      resourceHandler(),
			"sensu_hook":         resourceHook(),
			"sensu_mutator":      resourceMutator(),
			"sensu_namespace":    resourceNamespace(),
			"sensu_role":         resourceRole(),
			"sensu_role_binding": resourceRoleBinding(),
			"sensu_user":         resourceUser(),
		},

		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		apiUrl:    d.Get("api_url").(string),
		username:  d.Get("username").(string),
		password:  d.Get("password").(string),
		edition:   d.Get("edition").(string),
		namespace: d.Get("namespace").(string),
	}

	if err := config.LoadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
