package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesResyncIntervalSettingsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Fetch the reysnc interval for the given network device id.
`,

		ReadContext: dataSourceNetworkDevicesResyncIntervalSettingsIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The id of the network device.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interval": &schema.Schema{
							Description: `Resync interval of the device
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

func dataSourceNetworkDevicesResyncIntervalSettingsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetResyncIntervalForTheNetworkDeviceV1")
		vvID := vID.(string)

		response1, restyResp1, err := client.Devices.GetResyncIntervalForTheNetworkDeviceV1(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetResyncIntervalForTheNetworkDeviceV1", err,
				"Failure at GetResyncIntervalForTheNetworkDeviceV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesGetResyncIntervalForTheNetworkDeviceV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetResyncIntervalForTheNetworkDeviceV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetResyncIntervalForTheNetworkDeviceV1Item(item *catalystcentersdkgo.ResponseDevicesGetResyncIntervalForTheNetworkDeviceV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interval"] = item.Interval
	return []map[string]interface{}{
		respItem,
	}
}
