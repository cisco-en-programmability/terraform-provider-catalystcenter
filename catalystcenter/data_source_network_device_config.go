package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceConfig() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the config for all devices. This data source has been deprecated and will not be available in a Cisco Catalyst
Center release after Nov 1st 2024 23:59:59 GMT.

- Returns the device config by specified device ID
`,

		ReadContext: dataSourceNetworkDeviceConfigRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"cdp_neighbors": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"health_monitor": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"intf_description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"inventory": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_intf_brief": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address_table": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"running_config": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"snmp": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDeviceConfigRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okNetworkDeviceID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceConfigForAllDevicesV1")

		response1, restyResp1, err := client.Devices.GetDeviceConfigForAllDevicesV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceConfigForAllDevicesV1", err,
				"Failure at GetDeviceConfigForAllDevicesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetDeviceConfigForAllDevicesV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceConfigForAllDevicesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceConfigByIDV1")
		vvNetworkDeviceID := vNetworkDeviceID.(string)

		response2, restyResp2, err := client.Devices.GetDeviceConfigByIDV1(vvNetworkDeviceID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceConfigByIDV1", err,
				"Failure at GetDeviceConfigByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetDeviceConfigByIDV1Item(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceConfigByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceConfigForAllDevicesV1Items(items *[]catalystcentersdkgo.ResponseDevicesGetDeviceConfigForAllDevicesV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetDeviceConfigForAllDevicesV1ItemsAttributeInfo(item.AttributeInfo)
		respItem["cdp_neighbors"] = item.CdpNeighbors
		respItem["health_monitor"] = item.HealthMonitor
		respItem["id"] = item.ID
		respItem["intf_description"] = item.IntfDescription
		respItem["inventory"] = item.Inventory
		respItem["ip_intf_brief"] = item.IPIntfBrief
		respItem["mac_address_table"] = item.MacAddressTable
		respItem["running_config"] = item.RunningConfig
		respItem["snmp"] = item.SNMP
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceConfigForAllDevicesV1ItemsAttributeInfo(item *catalystcentersdkgo.ResponseDevicesGetDeviceConfigForAllDevicesV1ResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceConfigByIDV1Item(item *catalystcentersdkgo.ResponseDevicesGetDeviceConfigByIDV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
