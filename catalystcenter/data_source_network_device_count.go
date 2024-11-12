package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the interface count for the given device

- Returns the count of network devices based on the filter criteria by management IP address, mac address, hostname and
location name
`,

		ReadContext: dataSourceNetworkDeviceCountRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId path parameter. Device ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Description: `hostname query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"location_name": &schema.Schema{
				Description: `locationName query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"mac_address": &schema.Schema{
				Description: `macAddress query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"management_ip_address": &schema.Schema{
				Description: `managementIpAddress query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeInt,
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

func dataSourceNetworkDeviceCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vHostname, okHostname := d.GetOk("hostname")
	vManagementIPAddress, okManagementIPAddress := d.GetOk("management_ip_address")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vLocationName, okLocationName := d.GetOk("location_name")

	method1 := []bool{okDeviceID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okHostname, okManagementIPAddress, okMacAddress, okLocationName}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceInterfaceCountV1")
		vvDeviceID := vDeviceID.(string)

		response1, restyResp1, err := client.Devices.GetDeviceInterfaceCountV1(vvDeviceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceInterfaceCountV1", err,
				"Failure at GetDeviceInterfaceCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetDeviceCountKnowYourNetworkV1")
		queryParams2 := catalystcentersdkgo.GetDeviceCountKnowYourNetworkV1QueryParams{}

		if okHostname {
			queryParams2.Hostname = interfaceToSliceString(vHostname)
		}
		if okManagementIPAddress {
			queryParams2.ManagementIPAddress = interfaceToSliceString(vManagementIPAddress)
		}
		if okMacAddress {
			queryParams2.MacAddress = interfaceToSliceString(vMacAddress)
		}
		if okLocationName {
			queryParams2.LocationName = interfaceToSliceString(vLocationName)
		}

		response2, restyResp2, err := client.Devices.GetDeviceCountKnowYourNetworkV1(&queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceCountKnowYourNetworkV1", err,
				"Failure at GetDeviceCountKnowYourNetworkV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetDeviceCountKnowYourNetworkV1Item(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCountKnowYourNetworkV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceCountKnowYourNetworkV1Item(item *catalystcentersdkgo.ResponseDevicesGetDeviceCountKnowYourNetworkV1) []map[string]interface{} {
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
