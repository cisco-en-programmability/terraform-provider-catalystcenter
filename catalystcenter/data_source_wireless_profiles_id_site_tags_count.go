package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessProfilesIDSiteTagsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This endpoint retrieves the total count of Site Tags associated with a specific Wireless Profile.This data source
requires the id of the Wireless Profile to be provided as a path parameter.
`,

		ReadContext: dataSourceWirelessProfilesIDSiteTagsCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Wireless profile id
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
							Description: `Count of the requested resource
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

func dataSourceWirelessProfilesIDSiteTagsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheCountOfSiteTagsForAWirelessProfileV1")
		vvID := vID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.RetrieveTheCountOfSiteTagsForAWirelessProfileV1(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheCountOfSiteTagsForAWirelessProfileV1", err,
				"Failure at RetrieveTheCountOfSiteTagsForAWirelessProfileV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheCountOfSiteTagsForAWirelessProfileV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1Item(item *catalystcentersdkgo.ResponseWirelessRetrieveTheCountOfSiteTagsForAWirelessProfileV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
