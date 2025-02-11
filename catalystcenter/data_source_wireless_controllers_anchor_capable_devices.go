package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersAnchorCapableDevices() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get Anchor capable devices
`,

		ReadContext: dataSourceWirelessControllersAnchorCapableDevicesRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_ip": &schema.Schema{
							Description: `Anchor Controller Ip
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_name": &schema.Schema{
							Description: `Anchor Controller host name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"wireless_mgmt_ip": &schema.Schema{
							Description: `Wireless management Ip Address
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

func dataSourceWirelessControllersAnchorCapableDevicesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAnchorCapableDevicesV1")

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.GetAnchorCapableDevicesV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAnchorCapableDevicesV1", err,
				"Failure at GetAnchorCapableDevicesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetAnchorCapableDevicesV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAnchorCapableDevicesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAnchorCapableDevicesV1Item(item *catalystcentersdkgo.ResponseWirelessGetAnchorCapableDevicesV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_ip"] = item.DeviceIP
	respItem["device_name"] = item.DeviceName
	respItem["wireless_mgmt_ip"] = item.WirelessMgmtIP
	return []map[string]interface{}{
		respItem,
	}
}
