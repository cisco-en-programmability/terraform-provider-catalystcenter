package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkProfilesForSitesProfileIDTemplatesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieves the count of all CLI templates attached to a network profile by the profile ID.
`,

		ReadContext: dataSourceNetworkProfilesForSitesProfileIDTemplatesCountRead,
		Schema: map[string]*schema.Schema{
			"profile_id": &schema.Schema{
				Description: `profileId path parameter. The id of the network profile, retrievable from GET /intent/api/v1/networkProfilesForSites
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkProfilesForSitesProfileIDTemplatesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vProfileID := d.Get("profile_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1")
		vvProfileID := vProfileID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.NetworkSettings.RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1(vvProfileID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1", err,
				"Failure at RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveCountOfCliTemplatesAttachedToANetworkProfileV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1Item(item *catalystcentersdkgo.ResponseNetworkSettingsRetrieveCountOfCliTemplatesAttachedToANetworkProfileV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
