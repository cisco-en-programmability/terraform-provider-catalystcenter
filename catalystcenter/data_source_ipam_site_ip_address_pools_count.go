package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPamSiteIPAddressPoolsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Counts IP address subpools, which reserve address space from a global pool (or global pools).
`,

		ReadContext: dataSourceIPamSiteIPAddressPoolsCountRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId query parameter. The id of the site for which to retrieve IP address subpools. Only subpools whose siteId matches will be counted.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count.
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

func dataSourceIPamSiteIPAddressPoolsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CountsIPAddressSubpoolsV1")
		queryParams1 := catalystcentersdkgo.CountsIPAddressSubpoolsV1QueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.NetworkSettings.CountsIPAddressSubpoolsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CountsIPAddressSubpoolsV1", err,
				"Failure at CountsIPAddressSubpoolsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsCountsIPAddressSubpoolsV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountsIPAddressSubpoolsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsCountsIPAddressSubpoolsV1Item(item *catalystcentersdkgo.ResponseNetworkSettingsCountsIPAddressSubpoolsV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
