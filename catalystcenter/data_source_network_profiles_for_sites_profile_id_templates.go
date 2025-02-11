package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkProfilesForSitesProfileIDTemplates() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieves a list of CLI templates attached to a network profile based on the network profile ID.
`,

		ReadContext: dataSourceNetworkProfilesForSitesProfileIDTemplatesRead,
		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Description: `profileId path parameter. The id of the network profile, retrievable from GET /intent/api/v1/networkProfilesForSites
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `The id of the template attached to the site profile - /intent/api/v1/templates
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `The name of the template attached to the site profile - /intent/api/v1/templates
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkProfilesForSitesProfileIDTemplatesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vProfileID := d.Get("profile_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveCliTemplatesAttachedToANetworkProfileV1")
		vvProfileID := vProfileID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.NetworkSettings.RetrieveCliTemplatesAttachedToANetworkProfileV1(vvProfileID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveCliTemplatesAttachedToANetworkProfileV1", err,
				"Failure at RetrieveCliTemplatesAttachedToANetworkProfileV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveCliTemplatesAttachedToANetworkProfileV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1Items(items *[]catalystcentersdkgo.ResponseNetworkSettingsRetrieveCliTemplatesAttachedToANetworkProfileV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
