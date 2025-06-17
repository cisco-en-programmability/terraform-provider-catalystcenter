package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplatesTemplateIDNetworkProfilesForSites() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Retrieves the list of network profiles that a CLI template is currently attached to by the template ID.
`,

		ReadContext: dataSourceTemplatesTemplateIDNetworkProfilesForSitesRead,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Description: `templateId path parameter. The **id** of the template, retrievable from **GET /intent/api/v1/templates**
`,
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceTemplatesTemplateIDNetworkProfilesForSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vTemplateID := d.Get("template_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheNetworkProfilesAttachedToACLITemplate")
		vvTemplateID := vTemplateID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.ConfigurationTemplates.RetrieveTheNetworkProfilesAttachedToACLITemplate(vvTemplateID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheNetworkProfilesAttachedToACLITemplate", err,
				"Failure at RetrieveTheNetworkProfilesAttachedToACLITemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheNetworkProfilesAttachedToACLITemplate", err,
				"Failure at RetrieveTheNetworkProfilesAttachedToACLITemplate, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
