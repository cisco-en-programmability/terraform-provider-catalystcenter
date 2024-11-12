package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceEnrichmentDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Enriches a given network device context (device id or device Mac Address or device management IP address) with details
about the device and neighbor topology
`,

		ReadContext: dataSourceDeviceEnrichmentDetailsRead,
		Schema: map[string]*schema.Schema{
			"entity_type": &schema.Schema{
				Description: `entity_type header parameter. Device enrichment details can be fetched based on either Device ID or Device MAC address or Device IP Address. This parameter value must either be device_id/mac_address/ip_address
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_value": &schema.Schema{
				Description: `entity_value header parameter. Contains the actual value for the entity type that has been defined
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_manager_interface_ip": &schema.Schema{
										Description: `IP address of WLC on AP manager interface
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"associated_wlc_ip": &schema.Schema{
										Description: `Associated WLC IP address of the AP device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"boot_date_time": &schema.Schema{
										Description: `Device's last boot UTC timestamp
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"collection_interval": &schema.Schema{
										Description: `Re sync Interval of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"collection_status": &schema.Schema{
										Description: `Device's telemetry data collection status for DNAC
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"error_code": &schema.Schema{
										Description: `Inventory status error code
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"error_description": &schema.Schema{
										Description: `Inventory status description
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"family": &schema.Schema{
										Description: `Device Family
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"hostname": &schema.Schema{
										Description: `Device Hostname
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Device's UUID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance Uuid of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"interface_count": &schema.Schema{
										Description: `Number of interfaces on the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inventory_status_detail": &schema.Schema{
										Description: `Status detail of inventory sync
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_update_time": &schema.Schema{
										Description: `Time in epoch when the network device info last got updated
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"last_updated": &schema.Schema{
										Description: `Time when the network device info last got updated
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"line_card_count": &schema.Schema{
										Description: `Number of linecards on the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"line_card_id": &schema.Schema{
										Description: `IDs of linecards of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"location": &schema.Schema{
										Description: `Device location - Site hierarchy
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"location_name": &schema.Schema{
										Description: `[Deprecated] Name of the associated location
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"mac_address": &schema.Schema{
										Description: `Device MAC address
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"management_ip_address": &schema.Schema{
										Description: `Device Management Ip Address
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"memory_size": &schema.Schema{
										Description: `Processor memory size
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"neighbor_topology": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"links": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `Id of the node
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},

															"label": &schema.Schema{
																Description: `The details of the edge
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"link_status": &schema.Schema{
																Description: `The status of the link (up/down)
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"port_utilization": &schema.Schema{
																Description: `Number of clients
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},

															"source": &schema.Schema{
																Description: `Edge line starting node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"target": &schema.Schema{
																Description: `End node of the edge line
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"nodes": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"clients": &schema.Schema{
																Description: `Number of clients
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},

															"connected_device": &schema.Schema{
																Description: `The connected device to show the connected switch to wlc
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},

															"count": &schema.Schema{
																Description: `The number of group nodes (for ap sepecifically)
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},

															"description": &schema.Schema{
																Description: `Description of the Node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"device_type": &schema.Schema{
																Description: `Device type of the node, like switch, AP, WCL,GateWay
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"fabric_group": &schema.Schema{
																Description: `Fabric device group name
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},

															"family": &schema.Schema{
																Description: `Device Family of the Node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"health_score": &schema.Schema{
																Description: `The total health score of the node
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"id": &schema.Schema{
																Description: `Id of the Node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ip": &schema.Schema{
																Description: `IP Address of the Node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"level": &schema.Schema{
																Description: `The level index to be used by UI widget (starts from 0)
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"name": &schema.Schema{
																Description: `Hostname of the Node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"node_type": &schema.Schema{
																Description: `Type of the Node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"platform_id": &schema.Schema{
																Description: `Type of platform
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"radio_frequency": &schema.Schema{
																Description: `Frequency of wireless radio channel
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},

															"role": &schema.Schema{
																Description: `Role of the Node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"software_version": &schema.Schema{
																Description: `Software Version of the Node
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"user_id": &schema.Schema{
																Description: `User Id of the Node
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"platform_id": &schema.Schema{
										Description: `Device's platform ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"reachability_failure_reason": &schema.Schema{
										Description: `Failure reason for unreachable devices
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"reachability_status": &schema.Schema{
										Description: `Reachability Status of the Device(Reachable/Unreachable)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"role": &schema.Schema{
										Description: `Device role
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"role_source": &schema.Schema{
										Description: `Role source as manual / auto
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"serial_number": &schema.Schema{
										Description: `Device Serial Number
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"series": &schema.Schema{
										Description: `Device Series
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"snmp_contact": &schema.Schema{
										Description: `SNMP contact on device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"snmp_location": &schema.Schema{
										Description: `SNMP location on device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"software_version": &schema.Schema{
										Description: `Device Software Version
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tag_count": &schema.Schema{
										Description: `Number of tags associated with the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tunnel_udp_port": &schema.Schema{
										Description: `Mobility protocol port is stored as tunneludpport for WLC
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `Device Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"up_time": &schema.Schema{
										Description: `Device's uptime
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"waas_device_mode": &schema.Schema{
										Description: `WAAS device mode
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceEnrichmentDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vEntityType := d.Get("entity_type")
	vEntityValue := d.Get("entity_value")
	vPersistbapioutput := d.Get("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceEnrichmentDetailsV1")

		headerParams1 := catalystcentersdkgo.GetDeviceEnrichmentDetailsV1HeaderParams{}

		headerParams1.EntityType = vEntityType.(string)

		headerParams1.EntityValue = vEntityValue.(string)

		headerParams1.Persistbapioutput = vPersistbapioutput.(string)

		response1, restyResp1, err := client.Devices.GetDeviceEnrichmentDetailsV1(&headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceEnrichmentDetailsV1", err,
				"Failure at GetDeviceEnrichmentDetailsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetDeviceEnrichmentDetailsV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceEnrichmentDetailsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceEnrichmentDetailsV1Items(items *catalystcentersdkgo.ResponseDevicesGetDeviceEnrichmentDetailsV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_details"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetails(item.DeviceDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetails(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["family"] = item.Family
	respItem["type"] = item.Type
	respItem["location"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsLocation(item.Location)
	respItem["error_code"] = item.ErrorCode
	respItem["mac_address"] = item.MacAddress
	respItem["role"] = item.Role
	respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
	respItem["associated_wlc_ip"] = item.AssociatedWlcIP
	respItem["boot_date_time"] = item.BootDateTime
	respItem["collection_status"] = item.CollectionStatus
	respItem["interface_count"] = item.InterfaceCount
	respItem["line_card_count"] = item.LineCardCount
	respItem["line_card_id"] = item.LineCardID
	respItem["management_ip_address"] = item.ManagementIPAddress
	respItem["memory_size"] = item.MemorySize
	respItem["platform_id"] = item.PlatformID
	respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
	respItem["reachability_status"] = item.ReachabilityStatus
	respItem["snmp_contact"] = item.SNMPContact
	respItem["snmp_location"] = item.SNMPLocation
	respItem["tunnel_udp_port"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsTunnelUDPPort(item.TunnelUDPPort)
	respItem["waas_device_mode"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsWaasDeviceMode(item.WaasDeviceMode)
	respItem["series"] = item.Series
	respItem["inventory_status_detail"] = item.InventoryStatusDetail
	respItem["collection_interval"] = item.CollectionInterval
	respItem["serial_number"] = item.SerialNumber
	respItem["software_version"] = item.SoftwareVersion
	respItem["role_source"] = item.RoleSource
	respItem["hostname"] = item.Hostname
	respItem["up_time"] = item.UpTime
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["error_description"] = item.ErrorDescription
	respItem["location_name"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsLocationName(item.LocationName)
	respItem["tag_count"] = item.TagCount
	respItem["last_updated"] = item.LastUpdated
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["id"] = item.ID
	respItem["neighbor_topology"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopology(item.NeighborTopology)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsLocation(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsTunnelUDPPort(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsTunnelUDPPort) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsWaasDeviceMode(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsWaasDeviceMode) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsLocationName(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsLocationName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopology(items *[]catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopology) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["nodes"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodes(item.Nodes)
		respItem["links"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyLinks(item.Links)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodes(items *[]catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["role"] = item.Role
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["description"] = item.Description
		respItem["device_type"] = item.DeviceType
		respItem["platform_id"] = item.PlatformID
		respItem["family"] = item.Family
		respItem["ip"] = item.IP
		respItem["software_version"] = item.SoftwareVersion
		respItem["user_id"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesUserID(item.UserID)
		respItem["node_type"] = item.NodeType
		respItem["radio_frequency"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesRadioFrequency(item.RadioFrequency)
		respItem["clients"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesClients(item.Clients)
		respItem["count"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesCount(item.Count)
		respItem["health_score"] = item.HealthScore
		respItem["level"] = item.Level
		respItem["fabric_group"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesFabricGroup(item.FabricGroup)
		respItem["connected_device"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesConnectedDevice(item.ConnectedDevice)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesUserID(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesUserID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesRadioFrequency(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesRadioFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesClients(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesClients) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesCount(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesFabricGroup(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesFabricGroup) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyNodesConnectedDevice(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyNodesConnectedDevice) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyLinks(items *[]catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["source"] = item.Source
		respItem["link_status"] = item.LinkStatus
		respItem["label"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyLinksLabel(item.Label)
		respItem["target"] = item.Target
		respItem["id"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyLinksID(item.ID)
		respItem["port_utilization"] = flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyLinksPortUtilization(item.PortUtilization)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyLinksLabel(items *[]catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksLabel) []interface{} {
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

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyLinksID(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetDeviceEnrichmentDetailsV1ItemsDeviceDetailsNeighborTopologyLinksPortUtilization(item *catalystcentersdkgo.ResponseItemDevicesGetDeviceEnrichmentDetailsV1DeviceDetailsNeighborTopologyLinksPortUtilization) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
