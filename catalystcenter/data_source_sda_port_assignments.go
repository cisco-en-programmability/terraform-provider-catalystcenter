package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaPortAssignments() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of port assignments that match the provided query parameters.
`,

		ReadContext: dataSourceSdaPortAssignmentsRead,
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
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return. The maximum number of objects supported in a single request is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"native_vlan_id": &schema.Schema{
				Description: `nativeVlanId query parameter. Native VLAN of the port assignment, this option is only applicable to TRUNKING_DEVICE connectedDeviceType.(VLAN must be between 1 and 4094. In cases value not set when connectedDeviceType is TRUNKING_DEVICE, default value will be '1').
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Network device ID of the port assignment.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"voice_vlan_name": &schema.Schema{
				Description: `voiceVlanName query parameter. Voice VLAN name of the port assignment.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"allowed_vlan_ranges": &schema.Schema{
							Description: `Allowed VLAN of the port assignment, this option is only applicable to TRUNKING_DEVICE connectedDeviceType. (VLAN must be between 1 and 4094 (Ex 100,200,300-400) or 'all'. In cases value not set when connectedDeviceType is TRUNKING_DEVICE, default value will be 'all').
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"authenticate_template_name": &schema.Schema{
							Description: `Authenticate template name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"connected_device_type": &schema.Schema{
							Description: `Connected device type of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"data_vlan_name": &schema.Schema{
							Description: `Data VLAN name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric the device is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_description": &schema.Schema{
							Description: `Interface description of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_name": &schema.Schema{
							Description: `Interface name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"native_vlan_id": &schema.Schema{
							Description: `Native VLAN of the port assignment, this option is only applicable to TRUNKING_DEVICE connectedDeviceType. (VLAN must be between 1 and 4094. In cases value not set when connectedDeviceType is TRUNKING_DEVICE, default value will be 1).
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"network_device_id": &schema.Schema{
							Description: `Network device ID of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"security_group_name": &schema.Schema{
							Description: `Security group name of the port assignment.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"voice_vlan_name": &schema.Schema{
							Description: `Voice VLAN name of the port assignment.
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

func dataSourceSdaPortAssignmentsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vInterfaceName, okInterfaceName := d.GetOk("interface_name")
	vDataVLANName, okDataVLANName := d.GetOk("data_vlan_name")
	vVoiceVLANName, okVoiceVLANName := d.GetOk("voice_vlan_name")
	vNativeVLANID, okNativeVLANID := d.GetOk("native_vlan_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPortAssignments")
		queryParams1 := catalystcentersdkgo.GetPortAssignmentsQueryParams{}

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
		if okNativeVLANID {
			queryParams1.NativeVLANID = vNativeVLANID.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.GetPortAssignments(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPortAssignments", err,
				"Failure at GetPortAssignments, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPortAssignments", err,
				"Failure at GetPortAssignments, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetPortAssignmentsItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortAssignments response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetPortAssignmentsItems(items *[]catalystcentersdkgo.ResponseSdaGetPortAssignmentsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["fabric_id"] = item.FabricID
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["interface_name"] = item.InterfaceName
		respItem["connected_device_type"] = item.ConnectedDeviceType
		respItem["data_vlan_name"] = item.DataVLANName
		respItem["voice_vlan_name"] = item.VoiceVLANName
		respItem["authenticate_template_name"] = item.AuthenticateTemplateName
		respItem["security_group_name"] = item.SecurityGroupName
		respItem["interface_description"] = item.InterfaceDescription
		respItem["native_vlan_id"] = item.NativeVLANID
		respItem["allowed_vlan_ranges"] = item.AllowedVLANRanges
		respItems = append(respItems, respItem)
	}
	return respItems
}
