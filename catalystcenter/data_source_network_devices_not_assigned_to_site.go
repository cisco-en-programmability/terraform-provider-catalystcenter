package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesNotAssignedToSite() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Get network devices that are not assigned to any site.
`,

		ReadContext: dataSourceNetworkDevicesNotAssignedToSiteRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page;The minimum is 1, and the maximum is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_ids": &schema.Schema{
							Description: `Network device Ids.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDevicesNotAssignedToSiteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSiteNotAssignedNetworkDevices")
		queryParams1 := catalystcentersdkgo.GetSiteNotAssignedNetworkDevicesQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.SiteDesign.GetSiteNotAssignedNetworkDevices(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSiteNotAssignedNetworkDevices", err,
				"Failure at GetSiteNotAssignedNetworkDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSiteNotAssignedNetworkDevices", err,
				"Failure at GetSiteNotAssignedNetworkDevices, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetSiteNotAssignedNetworkDevicesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSiteNotAssignedNetworkDevices response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetSiteNotAssignedNetworkDevicesItem(item *catalystcentersdkgo.ResponseSiteDesignGetSiteNotAssignedNetworkDevicesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_ids"] = item.DeviceIDs
	return []map[string]interface{}{
		respItem,
	}
}
