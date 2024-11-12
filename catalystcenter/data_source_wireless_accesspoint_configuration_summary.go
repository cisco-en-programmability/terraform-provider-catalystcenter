package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessAccesspointConfigurationSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Users can query the access point configuration information per device using the ethernet MAC address
`,

		ReadContext: dataSourceWirelessAccesspointConfigurationSummaryRead,
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Description: `key query parameter. The ethernet MAC address of Access point
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"creation_order_index": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"is_being_changed": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ordered_list_oeassoc_name": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"ordered_list_oeindex": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

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

						"auth_entity_class": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"auth_entity_id": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"change_log_list": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"deploy_pending": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": &schema.Schema{
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

						"instance_created_on": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_id": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"instance_origin": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_updated_on": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"instance_version": &schema.Schema{
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"internal_key": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"long_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"url": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"lazy_loaded_entities": &schema.Schema{
							Type:     schema.TypeString, //TEST,
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

						"mesh_dtos": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"primary_controller_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"primary_ip_address": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"radio_dtos": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"creation_order_index": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"is_being_changed": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ordered_list_oeassoc_name": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"ordered_list_oeindex": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"admin_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"antenna_angle": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"antenna_elev_angle": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"antenna_gain": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"antenna_pattern_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"auth_entity_class": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"auth_entity_id": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"change_log_list": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"channel_assignment_mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"channel_number": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"channel_width": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"clean_air_si": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"deploy_pending": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"display_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"if_type": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"if_type_value": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_created_on": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"instance_id": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"instance_origin": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_updated_on": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"instance_version": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"internal_key": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"long_type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"url": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"lazy_loaded_entities": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"mac_address": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_assignment_mode": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"powerlevel": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"radio_band": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"radio_role_assignment": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"slot_id": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
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
					},
				},
			},
		},
	}
}

func dataSourceWirelessAccesspointConfigurationSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vKey := d.Get("key")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointConfigurationV1")
		queryParams1 := catalystcentersdkgo.GetAccessPointConfigurationV1QueryParams{}

		queryParams1.Key = vKey.(string)

		response1, restyResp1, err := client.Wireless.GetAccessPointConfigurationV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointConfigurationV1", err,
				"Failure at GetAccessPointConfigurationV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetAccessPointConfigurationV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointConfigurationV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAccessPointConfigurationV1Item(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["instance_uuid"] = flattenWirelessGetAccessPointConfigurationV1ItemInstanceUUID(item.InstanceUUID)
	respItem["instance_id"] = item.InstanceID
	respItem["auth_entity_id"] = flattenWirelessGetAccessPointConfigurationV1ItemAuthEntityID(item.AuthEntityID)
	respItem["display_name"] = item.DisplayName
	respItem["auth_entity_class"] = flattenWirelessGetAccessPointConfigurationV1ItemAuthEntityClass(item.AuthEntityClass)
	respItem["instance_tenant_id"] = item.InstanceTenantID
	respItem["ordered_list_oeindex"] = item.OrderedListOEIndex
	respItem["ordered_list_oeassoc_name"] = flattenWirelessGetAccessPointConfigurationV1ItemOrderedListOEAssocName(item.OrderedListOEAssocName)
	respItem["creation_order_index"] = item.CreationOrderIndex
	respItem["is_being_changed"] = boolPtrToString(item.IsBeingChanged)
	respItem["deploy_pending"] = item.DeployPending
	respItem["instance_created_on"] = flattenWirelessGetAccessPointConfigurationV1ItemInstanceCreatedOn(item.InstanceCreatedOn)
	respItem["instance_updated_on"] = flattenWirelessGetAccessPointConfigurationV1ItemInstanceUpdatedOn(item.InstanceUpdatedOn)
	respItem["change_log_list"] = flattenWirelessGetAccessPointConfigurationV1ItemChangeLogList(item.ChangeLogList)
	respItem["instance_origin"] = flattenWirelessGetAccessPointConfigurationV1ItemInstanceOrigin(item.InstanceOrigin)
	respItem["lazy_loaded_entities"] = flattenWirelessGetAccessPointConfigurationV1ItemLazyLoadedEntities(item.LazyLoadedEntities)
	respItem["instance_version"] = item.InstanceVersion
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
	respItem["mesh_dtos"] = flattenWirelessGetAccessPointConfigurationV1ItemMeshDTOs(item.MeshDTOs)
	respItem["radio_dtos"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOs(item.RadioDTOs)
	respItem["internal_key"] = flattenWirelessGetAccessPointConfigurationV1ItemInternalKey(item.InternalKey)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetAccessPointConfigurationV1ItemInstanceUUID(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1InstanceUUID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemAuthEntityID(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1AuthEntityID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemAuthEntityClass(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1AuthEntityClass) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemOrderedListOEAssocName(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1OrderedListOEAssocName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemInstanceCreatedOn(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1InstanceCreatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemInstanceUpdatedOn(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1InstanceUpdatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemChangeLogList(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1ChangeLogList) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemInstanceOrigin(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1InstanceOrigin) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemLazyLoadedEntities(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1LazyLoadedEntities) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemMeshDTOs(items *[]catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1MeshDTOs) []interface{} {
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

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOs(items *[]catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_uuid"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInstanceUUID(item.InstanceUUID)
		respItem["instance_id"] = item.InstanceID
		respItem["auth_entity_id"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsAuthEntityID(item.AuthEntityID)
		respItem["display_name"] = item.DisplayName
		respItem["auth_entity_class"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsAuthEntityClass(item.AuthEntityClass)
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["ordered_list_oeindex"] = item.OrderedListOEIndex
		respItem["ordered_list_oeassoc_name"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsOrderedListOEAssocName(item.OrderedListOEAssocName)
		respItem["creation_order_index"] = item.CreationOrderIndex
		respItem["is_being_changed"] = boolPtrToString(item.IsBeingChanged)
		respItem["deploy_pending"] = item.DeployPending
		respItem["instance_created_on"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInstanceCreatedOn(item.InstanceCreatedOn)
		respItem["instance_updated_on"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInstanceUpdatedOn(item.InstanceUpdatedOn)
		respItem["change_log_list"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsChangeLogList(item.ChangeLogList)
		respItem["instance_origin"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInstanceOrigin(item.InstanceOrigin)
		respItem["lazy_loaded_entities"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsLazyLoadedEntities(item.LazyLoadedEntities)
		respItem["instance_version"] = item.InstanceVersion
		respItem["admin_status"] = item.AdminStatus
		respItem["antenna_angle"] = item.AntennaAngle
		respItem["antenna_elev_angle"] = item.AntennaElevAngle
		respItem["antenna_gain"] = item.AntennaGain
		respItem["antenna_pattern_name"] = item.AntennaPatternName
		respItem["channel_assignment_mode"] = item.ChannelAssignmentMode
		respItem["channel_number"] = item.ChannelNumber
		respItem["channel_width"] = item.ChannelWidth
		respItem["clean_air_si"] = item.CleanAirSI
		respItem["if_type"] = item.IfType
		respItem["if_type_value"] = item.IfTypeValue
		respItem["mac_address"] = item.MacAddress
		respItem["power_assignment_mode"] = item.PowerAssignmentMode
		respItem["powerlevel"] = item.Powerlevel
		respItem["radio_band"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsRadioBand(item.RadioBand)
		respItem["radio_role_assignment"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsRadioRoleAssignment(item.RadioRoleAssignment)
		respItem["slot_id"] = item.SlotID
		respItem["internal_key"] = flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInternalKey(item.InternalKey)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInstanceUUID(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceUUID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsAuthEntityID(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsAuthEntityID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsAuthEntityClass(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsAuthEntityClass) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsOrderedListOEAssocName(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsOrderedListOEAssocName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInstanceCreatedOn(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceCreatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInstanceUpdatedOn(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceUpdatedOn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsChangeLogList(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsChangeLogList) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInstanceOrigin(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInstanceOrigin) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsLazyLoadedEntities(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsLazyLoadedEntities) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsRadioBand(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsRadioBand) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsRadioRoleAssignment(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsRadioRoleAssignment) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenWirelessGetAccessPointConfigurationV1ItemRadioDTOsInternalKey(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1RadioDTOsInternalKey) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["id"] = item.ID
	respItem["long_type"] = item.LongType
	respItem["url"] = item.URL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetAccessPointConfigurationV1ItemInternalKey(item *catalystcentersdkgo.ResponseWirelessGetAccessPointConfigurationV1InternalKey) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["id"] = item.ID
	respItem["long_type"] = item.LongType
	respItem["url"] = item.URL

	return []map[string]interface{}{
		respItem,
	}

}
