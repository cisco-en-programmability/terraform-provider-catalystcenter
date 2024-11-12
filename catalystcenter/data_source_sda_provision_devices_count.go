package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaProvisionDevicesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns the count of provisioned devices based on query parameters.
`,

		ReadContext: dataSourceSdaProvisionDevicesCountRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId query parameter. ID of the site hierarchy.
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
							Description: `Number of provisioned devices.
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

func dataSourceSdaProvisionDevicesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetProvisionedDevicesCountV1")
		queryParams1 := catalystcentersdkgo.GetProvisionedDevicesCountV1QueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}

		response1, restyResp1, err := client.Sda.GetProvisionedDevicesCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetProvisionedDevicesCountV1", err,
				"Failure at GetProvisionedDevicesCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetProvisionedDevicesCountV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProvisionedDevicesCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetProvisionedDevicesCountV1Item(item *catalystcentersdkgo.ResponseSdaGetProvisionedDevicesCountV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
