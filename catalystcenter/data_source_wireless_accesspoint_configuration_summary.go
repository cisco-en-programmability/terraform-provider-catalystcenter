package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessAccesspointConfigurationSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Users can query access point configuration information for a specific device by using the Ethernet MAC address as a
'key' filter. If no key is specified, all access point details will be retrieved based on the combination of filters
provided.
`,

		ReadContext: dataSourceWirelessAccesspointConfigurationSummaryRead,
		Schema: map[string]*schema.Schema{
			"ap_mode": &schema.Schema{
				Description: `apMode query parameter. AP Mode. Allowed values are Local, Bridge, Monitor, FlexConnect, Sniffer, Rogue Detector, SE-Connect, Flex+Bridge, Sensor.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ap_model": &schema.Schema{
				Description: `apModel query parameter. AP Model
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"key": &schema.Schema{
				Description: `key query parameter. The ethernet MAC address of Access point
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. The default is 500 if not specified. The maximum allowed limit is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"mesh_role": &schema.Schema{
				Description: `meshRole query parameter. Mesh Role. Allowed values are RAP or MAP
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"provisioned": &schema.Schema{
				Description: `provisioned query parameter. Indicate whether AP provisioned or not. Allowed values are True or False
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"wlc_ip_address": &schema.Schema{
				Description: `wlcIpAddress query parameter. WLC IP Address
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"admin_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_height": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"ap_mode": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"eth_mac": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"failover_priority": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"led_brightness_level": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"led_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_address": &schema.Schema{
							Description: `Management Ip Address
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mesh_dtos": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"model": &schema.Schema{
							Description: `AP Model`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"primary_controller_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"primary_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"provisioning_status": &schema.Schema{
							Description: `Provisioning Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"radio_dtos": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"admin_status": &schema.Schema{
										Description: `Admin Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"antenna_angle": &schema.Schema{
										Description: `Antenna Angle`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"antenna_elev_angle": &schema.Schema{
										Description: `Antenna Elev Angle`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"antenna_gain": &schema.Schema{
										Description: `Antenna Gain`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"antenna_pattern_name": &schema.Schema{
										Description: `Antenna Pattern Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel_assignment_mode": &schema.Schema{
										Description: `Channel Assignment Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"channel_number": &schema.Schema{
										Description: `Channel Number`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"channel_width": &schema.Schema{
										Description: `Channel Width`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"clean_air_si": &schema.Schema{
										Description: `Clean Air SI`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"dual_radio_mode": &schema.Schema{
										Description: `Dual Radio Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"if_type": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"if_type_value": &schema.Schema{
										Description: `If Type Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"mac_address": &schema.Schema{
										Description: `Mac Address`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"power_assignment_mode": &schema.Schema{
										Description: `Power Assignment Mode`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"powerlevel": &schema.Schema{
										Description: `Powerlevel`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"radio_band": &schema.Schema{
										Description: `Radio Band`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"radio_role_assignment": &schema.Schema{
										Description: `Radio Role Assignment`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"slot_id": &schema.Schema{
										Description: `Slot Id`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"reachability_status": &schema.Schema{
							Description: `Reachability Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"secondary_controller_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"secondary_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tertiary_controller_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"tertiary_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"wlc_ip_address": &schema.Schema{
							Description: `WLC IP Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessAccesspointConfigurationSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vKey, okKey := d.GetOk("key")
	vWlcIPAddress, okWlcIPAddress := d.GetOk("wlc_ip_address")
	vApMode, okApMode := d.GetOk("ap_mode")
	vApModel, okApModel := d.GetOk("ap_model")
	vMeshRole, okMeshRole := d.GetOk("mesh_role")
	vProvisioned, okProvisioned := d.GetOk("provisioned")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointConfiguration")
		queryParams1 := catalystcentersdkgo.GetAccessPointConfigurationQueryParams{}

		if okKey {
			queryParams1.Key = vKey.(string)
		}
		if okWlcIPAddress {
			queryParams1.WlcIPAddress = vWlcIPAddress.(string)
		}
		if okApMode {
			queryParams1.ApMode = vApMode.(string)
		}
		if okApModel {
			queryParams1.ApModel = vApModel.(string)
		}
		if okMeshRole {
			queryParams1.MeshRole = vMeshRole.(string)
		}
		if okProvisioned {
			queryParams1.Provisioned = vProvisioned.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.GetAccessPointConfiguration(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointConfiguration", err,
				"Failure at GetAccessPointConfiguration, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointConfiguration", err,
				"Failure at GetAccessPointConfiguration, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetAccessPointConfigurationItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointConfiguration response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAccessPointConfigurationItem(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfiguration) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["ap_height"] = item.ApHeight
	respItem["ap_mode"] = item.ApMode
	respItem["ap_name"] = item.ApName
	respItem["eth_mac"] = item.EthMac
	respItem["failover_priority"] = item.FailoverPriority
	respItem["led_brightness_level"] = item.LedBrightnessLevel
	respItem["led_status"] = item.LedStatus
	respItem["location"] = item.Location
	respItem["mac_address"] = item.MacAddress
	respItem["primary_controller_name"] = item.PrimaryControllerName
	respItem["primary_ip_address"] = item.PrimaryIPAddress
	respItem["secondary_controller_name"] = item.SecondaryControllerName
	respItem["secondary_ip_address"] = item.SecondaryIPAddress
	respItem["tertiary_controller_name"] = item.TertiaryControllerName
	respItem["tertiary_ip_address"] = item.TertiaryIPAddress
	respItem["mesh_dtos"] = flattenWirelessGetAccessPointConfigurationItemMeshDTOs(item.MeshDTOs)
	respItem["model"] = item.Model
	respItem["wlc_ip_address"] = item.WlcIPAddress
	respItem["reachability_status"] = item.ReachabilityStatus
	respItem["management_ip_address"] = item.ManagementIPAddress
	respItem["provisioning_status"] = item.ProvisioningStatus
	respItem["radio_dtos"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOs(item.RadioDTOs)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetAccessPointConfigurationItemMeshDTOs(items *[]catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationMeshDTOs) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOs(items *[]catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["if_type"] = item.IfType
		respItem["if_type_value"] = item.IfTypeValue
		respItem["slot_id"] = item.SlotID
		respItem["mac_address"] = item.MacAddress
		respItem["admin_status"] = item.AdminStatus
		respItem["power_assignment_mode"] = item.PowerAssignmentMode
		respItem["powerlevel"] = item.Powerlevel
		respItem["channel_assignment_mode"] = item.ChannelAssignmentMode
		respItem["channel_number"] = item.ChannelNumber
		respItem["channel_width"] = item.ChannelWidth
		respItem["antenna_pattern_name"] = item.AntennaPatternName
		respItem["antenna_angle"] = item.AntennaAngle
		respItem["antenna_elev_angle"] = item.AntennaElevAngle
		respItem["antenna_gain"] = item.AntennaGain
		respItem["radio_role_assignment"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsRadioRoleAssignment(item.RadioRoleAssignment)
		respItem["radio_band"] = flattenWirelessGetAccessPointConfigurationItemRadioDTOsRadioBand(item.RadioBand)
		respItem["clean_air_si"] = item.CleanAirSI
		respItem["dual_radio_mode"] = item.DualRadioMode
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsRadioRoleAssignment(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsRadioRoleAssignment) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationItemRadioDTOsRadioBand(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationRadioDTOsRadioBand) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
