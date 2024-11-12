package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicePoe() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns POE details for device.
`,

		ReadContext: dataSourceNetworkDevicePoeRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter. UUID of the device
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"power_allocated": &schema.Schema{
							Description: `Total power available on the switch on all interfaces combined in Watts
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"power_consumed": &schema.Schema{
							Description: `Total power being currently drawn by all interfaces combined in Watts
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"power_remaining": &schema.Schema{
							Description: `Total power remaining in Watts (powerConsumed - powerAllocated)
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

func dataSourceNetworkDevicePoeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: PoeDetailsV1")
		vvDeviceUUID := vDeviceUUID.(string)

		response1, restyResp1, err := client.Devices.PoeDetailsV1(vvDeviceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 PoeDetailsV1", err,
				"Failure at PoeDetailsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesPoeDetailsV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting PoeDetailsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesPoeDetailsV1Item(item *catalystcentersdkgo.ResponseDevicesPoeDetailsV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["power_allocated"] = item.PowerAllocated
	respItem["power_consumed"] = item.PowerConsumed
	respItem["power_remaining"] = item.PowerRemaining
	return []map[string]interface{}{
		respItem,
	}
}
