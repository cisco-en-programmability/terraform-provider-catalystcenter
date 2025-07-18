package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLanAutomationStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on LAN Automation.

- Invoke this API to get the LAN Automation session status.

- Invoke this API to get the LAN Automation session status based on the given Lan Automation session id.
`,

		ReadContext: dataSourceLanAutomationStatusRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. LAN Automation session identifier.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting index of the LAN Automation session. Minimum value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"action": &schema.Schema{
							Description: `State (START/STOP) of the LAN Automation session.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"creation_time": &schema.Schema{
							Description: `LAN Automation session creation time.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovered_device_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_address_in_use_list": &schema.Schema{
										Description: `List of IP address used by the device.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"name": &schema.Schema{
										Description: `Name of the device.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"serial_number": &schema.Schema{
										Description: `Serial number of the device.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"state": &schema.Schema{
										Description: `State of the device (Added to inventory/Deleted from inventory).
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"discovered_device_site_name_hierarchy": &schema.Schema{
							Description: `Discovered device site name.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_devices": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_host_name": &schema.Schema{
										Description: `Hostname of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_management_ipaddress": &schema.Schema{
										Description: `Management IP Address of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_serial_number": &schema.Schema{
										Description: `Serial number of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_site_id": &schema.Schema{
										Description: `Site UUID of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_site_name_hierarchy": &schema.Schema{
										Description: `Site name hierarchy of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_device_discovered": &schema.Schema{
										Description: `Discovery status of the device
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_ipallocated": &schema.Schema{
										Description: `IP Allocation status of the device
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_ipassigned": &schema.Schema{
										Description: `IP Assignment status of the device
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"pnp_device_id": &schema.Schema{
										Description: `PnP device ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"discovery_level": &schema.Schema{
							Description: `Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"discovery_timeout": &schema.Schema{
							Description: `Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080].
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"host_name_file_id": &schema.Schema{
							Description: `File ID for the uploaded file in the nw_orch namespace
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"host_name_prefix": &schema.Schema{
							Description: `Hostname prefix of the discovered devices
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `LAN Automation session id.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_pools": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_pool_name": &schema.Schema{
										Description: `Name of the IP pool.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_pool_role": &schema.Schema{
										Description: `Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"multicast_enabled": &schema.Schema{
							Description: `Shows whether underlay multicast is enabled or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"peer_device_managment_ipaddress": &schema.Schema{
							Description: `Peer seed device management IP address.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"primary_device_interface_names": &schema.Schema{
							Description: `The list of interfaces on primary seed via which the discovered devices are connected.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"primary_device_managment_ipaddress": &schema.Schema{
							Description: `Primary seed device management IP address.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"redistribute_isis_to_bgp": &schema.Schema{
							Description: `Shows whether advertise LAN Automation summary route into BGP is enabled or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of the LAN Automation session along with the number of discovered devices.
`,
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

						"action": &schema.Schema{
							Description: `State (START/STOP) of the LAN Automation session.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"creation_time": &schema.Schema{
							Description: `LAN Automation session creation time.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovered_device_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_address_in_use_list": &schema.Schema{
										Description: `List of IP address used by the device.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"name": &schema.Schema{
										Description: `Name of the device.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"serial_number": &schema.Schema{
										Description: `Serial number of the device.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"state": &schema.Schema{
										Description: `State of the device (Added to inventory/Deleted from inventory).
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"discovered_device_site_name_hierarchy": &schema.Schema{
							Description: `Discovered device site name.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"discovery_devices": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_host_name": &schema.Schema{
										Description: `Hostname of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_management_ipaddress": &schema.Schema{
										Description: `Management IP Address of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_serial_number": &schema.Schema{
										Description: `Serial number of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_site_id": &schema.Schema{
										Description: `Site UUID of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_site_name_hierarchy": &schema.Schema{
										Description: `Site name hierarchy of the device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_device_discovered": &schema.Schema{
										Description: `Discovery status of the device
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_ipallocated": &schema.Schema{
										Description: `IP Allocation status of the device
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"is_ipassigned": &schema.Schema{
										Description: `IP Assignment status of the device
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"pnp_device_id": &schema.Schema{
										Description: `PnP device ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"discovery_level": &schema.Schema{
							Description: `Level below primary seed device upto which the new devices will be LAN Automated by this session, level + seed = tier. Supported range for level is [1-5], default level is 2.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"discovery_timeout": &schema.Schema{
							Description: `Discovery timeout in minutes. Until this time, the stop processing will not be triggered. Any device contacting after the provided discovery timeout will not be processed, and a device reset and reload will be attempted to bring it back to the PnP agent state before process completion. The supported timeout range is in minutes [20-10080].
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"host_name_file_id": &schema.Schema{
							Description: `File ID for the uploaded file in the nw_orch namespace
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"host_name_prefix": &schema.Schema{
							Description: `Hostname prefix of the discovered devices
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `LAN Automation session id.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_pools": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ip_pool_name": &schema.Schema{
										Description: `Name of the IP pool.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip_pool_role": &schema.Schema{
										Description: `Role of the IP pool. Supported roles are: MAIN_POOL and PHYSICAL_LINK_POOL.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"multicast_enabled": &schema.Schema{
							Description: `Shows whether underlay multicast is enabled or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"peer_device_managment_ipaddress": &schema.Schema{
							Description: `Peer seed device management IP address.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"primary_device_interface_names": &schema.Schema{
							Description: `The list of interfaces on primary seed via which the discovered devices are connected.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"primary_device_managment_ipaddress": &schema.Schema{
							Description: `Primary seed device management IP address.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"redistribute_isis_to_bgp": &schema.Schema{
							Description: `Shows whether advertise LAN Automation summary route into BGP is enabled or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of the LAN Automation session along with the number of discovered devices.
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

func dataSourceLanAutomationStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vID, okID := d.GetOk("id")

	method1 := []bool{okOffset, okLimit}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LanAutomationStatus")
		queryParams1 := catalystcentersdkgo.LanAutomationStatusQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.LanAutomation.LanAutomationStatus(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LanAutomationStatus", err,
				"Failure at LanAutomationStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LanAutomationStatus", err,
				"Failure at LanAutomationStatus, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLanAutomationLanAutomationStatusItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomationStatus response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: LanAutomationStatusByID")
		vvID := vID.(string)

		// has_unknown_response: None

		response2, restyResp2, err := client.LanAutomation.LanAutomationStatusByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LanAutomationStatusByID", err,
				"Failure at LanAutomationStatusByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LanAutomationStatusByID", err,
				"Failure at LanAutomationStatusByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenLanAutomationLanAutomationStatusByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomationStatusByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLanAutomationLanAutomationStatusItems(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationStatusResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["discovered_device_site_name_hierarchy"] = item.DiscoveredDeviceSiteNameHierarchy
		respItem["primary_device_managment_ipaddress"] = item.PrimaryDeviceManagmentIPAddress
		respItem["ip_pools"] = flattenLanAutomationLanAutomationStatusItemsIPPools(item.IPPools)
		respItem["primary_device_interface_names"] = item.PrimaryDeviceInterfaceNames
		respItem["status"] = item.Status
		respItem["action"] = item.Action
		respItem["creation_time"] = item.CreationTime
		respItem["multicast_enabled"] = boolPtrToString(item.MulticastEnabled)
		respItem["peer_device_managment_ipaddress"] = item.PeerDeviceManagmentIPAddress
		respItem["discovered_device_list"] = flattenLanAutomationLanAutomationStatusItemsDiscoveredDeviceList(item.DiscoveredDeviceList)
		respItem["redistribute_isis_to_bgp"] = boolPtrToString(item.RedistributeIsisToBgp)
		respItem["discovery_level"] = item.DiscoveryLevel
		respItem["discovery_timeout"] = item.DiscoveryTimeout
		respItem["discovery_devices"] = flattenLanAutomationLanAutomationStatusItemsDiscoveryDevices(item.DiscoveryDevices)
		respItem["host_name_prefix"] = item.HostNamePrefix
		respItem["host_name_file_id"] = item.HostNameFileID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusItemsIPPools(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationStatusResponseIPPools) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["ip_pool_role"] = item.IPPoolRole
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusItemsDiscoveredDeviceList(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationStatusResponseDiscoveredDeviceList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["serial_number"] = item.SerialNumber
		respItem["state"] = item.State
		respItem["ip_address_in_use_list"] = item.IPAddressInUseList
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusItemsDiscoveryDevices(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationStatusResponseDiscoveryDevices) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_serial_number"] = item.DeviceSerialNumber
		respItem["device_host_name"] = item.DeviceHostName
		respItem["device_management_ipaddress"] = item.DeviceManagementIPAddress
		respItem["device_site_id"] = item.DeviceSiteID
		respItem["device_site_name_hierarchy"] = item.DeviceSiteNameHierarchy
		respItem["is_device_discovered"] = boolPtrToString(item.IsDeviceDiscovered)
		respItem["is_ipallocated"] = boolPtrToString(item.IsIPAllocated)
		respItem["is_ipassigned"] = boolPtrToString(item.IsIPAssigned)
		respItem["pnp_device_id"] = item.PnpDeviceID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusByIDItem(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationStatusByIDResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["discovered_device_site_name_hierarchy"] = item.DiscoveredDeviceSiteNameHierarchy
		respItem["primary_device_managment_ipaddress"] = item.PrimaryDeviceManagmentIPAddress
		respItem["ip_pools"] = flattenLanAutomationLanAutomationStatusByIDItemIPPools(item.IPPools)
		respItem["primary_device_interface_names"] = item.PrimaryDeviceInterfaceNames
		respItem["status"] = item.Status
		respItem["action"] = item.Action
		respItem["creation_time"] = item.CreationTime
		respItem["multicast_enabled"] = boolPtrToString(item.MulticastEnabled)
		respItem["peer_device_managment_ipaddress"] = item.PeerDeviceManagmentIPAddress
		respItem["discovered_device_list"] = flattenLanAutomationLanAutomationStatusByIDItemDiscoveredDeviceList(item.DiscoveredDeviceList)
		respItem["redistribute_isis_to_bgp"] = boolPtrToString(item.RedistributeIsisToBgp)
		respItem["discovery_level"] = item.DiscoveryLevel
		respItem["discovery_timeout"] = item.DiscoveryTimeout
		respItem["discovery_devices"] = flattenLanAutomationLanAutomationStatusByIDItemDiscoveryDevices(item.DiscoveryDevices)
		respItem["host_name_prefix"] = item.HostNamePrefix
		respItem["host_name_file_id"] = item.HostNameFileID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusByIDItemIPPools(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationStatusByIDResponseIPPools) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["ip_pool_role"] = item.IPPoolRole
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusByIDItemDiscoveredDeviceList(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveredDeviceList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["serial_number"] = item.SerialNumber
		respItem["state"] = item.State
		respItem["ip_address_in_use_list"] = item.IPAddressInUseList
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationStatusByIDItemDiscoveryDevices(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationStatusByIDResponseDiscoveryDevices) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_serial_number"] = item.DeviceSerialNumber
		respItem["device_host_name"] = item.DeviceHostName
		respItem["device_management_ipaddress"] = item.DeviceManagementIPAddress
		respItem["device_site_id"] = item.DeviceSiteID
		respItem["device_site_name_hierarchy"] = item.DeviceSiteNameHierarchy
		respItem["is_device_discovered"] = boolPtrToString(item.IsDeviceDiscovered)
		respItem["is_ipallocated"] = boolPtrToString(item.IsIPAllocated)
		respItem["is_ipassigned"] = boolPtrToString(item.IsIPAssigned)
		respItem["pnp_device_id"] = item.PnpDeviceID
		respItems = append(respItems, respItem)
	}
	return respItems
}
