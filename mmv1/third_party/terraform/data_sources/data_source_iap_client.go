package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceGoogleIapClient() *schema.Resource {

	dsSchema := datasourceSchemaFromResourceSchema(ResourceIapClient().Schema)
	addRequiredFieldsToSchema(dsSchema, "brand", "client_id")

	return &schema.Resource{
		Read:   dataSourceGoogleIapClientRead,
		Schema: dsSchema,
	}
}

func dataSourceGoogleIapClientRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	id, err := ReplaceVars(d, config, "{{brand}}/identityAwareProxyClients/{{client_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)
	return resourceIapClientRead(d, meta)
}
