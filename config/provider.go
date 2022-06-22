package config

import (
	// Note(ezgidemirel): we are importing this to embed provider schema document
	_ "embed"

	tjconfig "github.com/crossplane/terrajet/pkg/config"
	"github.com/crossplane/terrajet/pkg/types/name"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/crossplane-contrib/provider-jet-gcp/config/accessapproval"
	"github.com/crossplane-contrib/provider-jet-gcp/config/bigtable"
	"github.com/crossplane-contrib/provider-jet-gcp/config/cloudfunctions"
	"github.com/crossplane-contrib/provider-jet-gcp/config/cloudiot"
	"github.com/crossplane-contrib/provider-jet-gcp/config/cloudplatform"
	"github.com/crossplane-contrib/provider-jet-gcp/config/compute"
	"github.com/crossplane-contrib/provider-jet-gcp/config/container"
	"github.com/crossplane-contrib/provider-jet-gcp/config/dataflow"
	"github.com/crossplane-contrib/provider-jet-gcp/config/dataproc"
	"github.com/crossplane-contrib/provider-jet-gcp/config/monitoring"
	"github.com/crossplane-contrib/provider-jet-gcp/config/project"
	"github.com/crossplane-contrib/provider-jet-gcp/config/redis"
	"github.com/crossplane-contrib/provider-jet-gcp/config/sql"
	"github.com/crossplane-contrib/provider-jet-gcp/config/storage"
)

const (
	resourcePrefix = "gcp"
	modulePath     = "github.com/crossplane-contrib/provider-jet-gcp"
)

//go:embed schema.json
var providerSchema string

var skipList = []string{
	// Note(turkenh): Following two resources conflicts their singular versions
	// "google_access_context_manager_access_level" and
	// "google_access_context_manager_service_perimeter". Skipping for now.
	"google_access_context_manager_access_levels$",
	"google_access_context_manager_service_perimeters$",
}

// GetProvider returns provider configuration
func GetProvider() *tjconfig.Provider {
	pc := tjconfig.NewProviderWithSchema([]byte(providerSchema), resourcePrefix, modulePath,
		tjconfig.WithDefaultResourceFn(DefaultResource(
			groupOverrides(),
			externalNameConfig(),
		)),
		tjconfig.WithRootGroup("gcp.jet.crossplane.io"),
		tjconfig.WithShortName("gcpjet"),
		// Comment out the following line to generate all resources.
		// tjconfig.WithIncludeList(includeList),
		tjconfig.WithSkipList(skipList))

	for _, configure := range []func(provider *tjconfig.Provider){
		accessapproval.Configure,
		bigtable.Configure,
		cloudfunctions.Configure,
		cloudiot.Configure,
		cloudplatform.Configure,
		container.Configure,
		compute.Configure,
		dataflow.Configure,
		dataproc.Configure,
		redis.Configure,
		monitoring.Configure,
		project.Configure,
		storage.Configure,
		sql.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// DefaultResource returns a DefaultResourceFn that makes sure the original
// DefaultResource call is made with given options here.
func DefaultResource(opts ...tjconfig.ResourceOption) tjconfig.DefaultResourceFn {
	return func(name string, terraformResource *schema.Resource, orgOpts ...tjconfig.ResourceOption) *tjconfig.Resource {
		return tjconfig.DefaultResource(name, terraformResource, append(orgOpts, opts...)...)
	}
}

func init() {
	// GCP specific acronyms

	// Todo(turkenh): move to Terrajet?
	name.AddAcronym("idp", "IdP")
	name.AddAcronym("oauth", "OAuth")
}
