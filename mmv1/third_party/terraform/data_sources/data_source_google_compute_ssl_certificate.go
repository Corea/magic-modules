package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func DataSourceGoogleComputeSslCertificate() *schema.Resource {
	// Generate datasource schema from resource
	dsSchema := datasourceSchemaFromResourceSchema(ResourceComputeSslCertificate().Schema)

	// Set 'Required' schema elements
	addRequiredFieldsToSchema(dsSchema, "name")

	// Set 'Optional' schema elements
	addOptionalFieldsToSchema(dsSchema, "project")

	return &schema.Resource{
		Read:   dataSourceComputeSslCertificateRead,
		Schema: dsSchema,
	}
}

func dataSourceComputeSslCertificateRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	certificateName := d.Get("name").(string)

	d.SetId(fmt.Sprintf("projects/%s/global/sslCertificates/%s", project, certificateName))

	return resourceComputeSslCertificateRead(d, meta)
}
