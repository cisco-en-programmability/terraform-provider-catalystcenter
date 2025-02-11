package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceSupervisorCardDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Get supervisor card detail for a given deviceuuid. Response will contain serial no, part no, switch no and slot no.
`,

		ReadContext: dataSourceNetworkDeviceSupervisorCardDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter. instanceuuid of device
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"partno": &schema.Schema{
							Description: `Part number of the supervisor card
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serialno": &schema.Schema{
							Description: `Serial number of the supervisor card
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"slotno": &schema.Schema{
							Description: `Slot number of supervisor card
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"switchno": &schema.Schema{
							Description: `Switch number of the supervisor card
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

func dataSourceNetworkDeviceSupervisorCardDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSupervisorCardDetailV1")
		vvDeviceUUID := vDeviceUUID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Devices.GetSupervisorCardDetailV1(vvDeviceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSupervisorCardDetailV1", err,
				"Failure at GetSupervisorCardDetailV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetSupervisorCardDetailV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSupervisorCardDetailV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetSupervisorCardDetailV1Items(items *[]catalystcentersdkgo.ResponseDevicesGetSupervisorCardDetailV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["serialno"] = item.Serialno
		respItem["partno"] = item.Partno
		respItem["switchno"] = item.Switchno
		respItem["slotno"] = item.Slotno
		respItems = append(respItems, respItem)
	}
	return respItems
}
