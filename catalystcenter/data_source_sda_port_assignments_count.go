package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaPortAssignmentsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns the count of port assignments that match the provided query parameters.
`,

		ReadContext: dataSourceSdaPortAssignmentsCountRead,
		Schema: map[string]*schema.Schema{
			"data_vlan_name": &schema.Schema{
				Description: `dataVlanName query parameter. Data VLAN name of the port assignment.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the device is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"interface_name": &schema.Schema{
				Description: `interfaceName query parameter. Interface name of the port assignment.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Network device ID of the port assignment.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"voice_vlan_name": &schema.Schema{
				Description: `voiceVlanName query parameter. Voice VLAN name of the port assignment.
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
							Description: `Number of port assignments.
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

func dataSourceSdaPortAssignmentsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vInterfaceName, okInterfaceName := d.GetOk("interface_name")
	vDataVLANName, okDataVLANName := d.GetOk("data_vlan_name")
	vVoiceVLANName, okVoiceVLANName := d.GetOk("voice_vlan_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPortAssignmentCountV1")
		queryParams1 := catalystcentersdkgo.GetPortAssignmentCountV1QueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okInterfaceName {
			queryParams1.InterfaceName = vInterfaceName.(string)
		}
		if okDataVLANName {
			queryParams1.DataVLANName = vDataVLANName.(string)
		}
		if okVoiceVLANName {
			queryParams1.VoiceVLANName = vVoiceVLANName.(string)
		}

		response1, restyResp1, err := client.Sda.GetPortAssignmentCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPortAssignmentCountV1", err,
				"Failure at GetPortAssignmentCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetPortAssignmentCountV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortAssignmentCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetPortAssignmentCountV1Item(item *catalystcentersdkgo.ResponseSdaGetPortAssignmentCountV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
