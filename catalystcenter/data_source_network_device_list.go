package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceList() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns list of network devices based on filter criteria such as management IP address, mac address, hostname, etc.
You can use the .* in any value to conduct a wildcard search. For example, to find all hostnames beginning with myhost
in the IP address range 192.25.18.n, issue the following request: GET /dna/intent/api/v1/network-
device?hostname=myhost.*&managementIpAddress=192.25.18..*
If id parameter is provided with comma separated ids, it will return the list of network-devices for the given ids and
ignores the other request parameters. The API returns a paginated response based on 'limit' and 'offset' parameters,
allowing up to 500 records per page. 'limit' specifies the number of records, and 'offset' sets the starting point using
1-based indexing. Use '/dna/intent/api/v1/network-device/count' to get the total record count. For data sets over 500
records, make multiple calls, adjusting 'limit' and 'offset' to retrieve all records incrementally.
`,

		ReadContext: dataSourceNetworkDeviceListRead,
		Schema: map[string]*schema.Schema{
			"associated_wlc_ip": &schema.Schema{
				Description: `associatedWlcIp query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"collection_interval": &schema.Schema{
				Description: `collectionInterval query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"collection_status": &schema.Schema{
				Description: `collectionStatus query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"device_support_level": &schema.Schema{
				Description: `deviceSupportLevel query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"error_code": &schema.Schema{
				Description: `errorCode query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"error_description": &schema.Schema{
				Description: `errorDescription query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"family": &schema.Schema{
				Description: `family query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"hostname": &schema.Schema{
				Description: `hostname query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"id": &schema.Schema{
				Description: `id query parameter. Accepts comma separated ids and return list of network-devices for the given ids. If invalid or not-found ids are provided, null entry will be returned in the list.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"license_name": &schema.Schema{
				Description: `license.name query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"license_status": &schema.Schema{
				Description: `license.status query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"license_type": &schema.Schema{
				Description: `license.type query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Min: 1, Max: 500
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"location": &schema.Schema{
				Description: `location query parameter.`,
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
			"module_equpimenttype": &schema.Schema{
				Description: `module+equpimenttype query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"module_name": &schema.Schema{
				Description: `module+name query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"module_operationstatecode": &schema.Schema{
				Description: `module+operationstatecode query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"module_partnumber": &schema.Schema{
				Description: `module+partnumber query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"module_servicestate": &schema.Schema{
				Description: `module+servicestate query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"module_vendorequipmenttype": &schema.Schema{
				Description: `module+vendorequipmenttype query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"not_synced_for_minutes": &schema.Schema{
				Description: `notSyncedForMinutes query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. offset >= 1 [X gives results from Xth device onwards]
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"platform_id": &schema.Schema{
				Description: `platformId query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"reachability_status": &schema.Schema{
				Description: `reachabilityStatus query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"role": &schema.Schema{
				Description: `role query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"series": &schema.Schema{
				Description: `series query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"software_type": &schema.Schema{
				Description: `softwareType query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"software_version": &schema.Schema{
				Description: `softwareVersion query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"type": &schema.Schema{
				Description: `type query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"up_time": &schema.Schema{
				Description: `upTime query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_ethernet_mac_address": &schema.Schema{
							Description: `AccessPoint Ethernet MacAddress of AP device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ap_manager_interface_ip": &schema.Schema{
							Description: `IP address of WLC on AP manager interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"associated_wlc_ip": &schema.Schema{
							Description: `Associated Wlc Ip address of the AP device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"boot_date_time": &schema.Schema{
							Description: `Device boot time
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
							Description: `Collection status as Synchronizing, Could not synchronize, Not manageable, Managed, Partial Collection Failure, Incomplete, Unreachable, Wrong credential, Reachable, In Progress
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `System description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_support_level": &schema.Schema{
							Description: `Support level of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"dns_resolved_management_address": &schema.Schema{
							Description: `Specifies the resolved ip address of dns name
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
							Description: `Family of device as switch, router, wireless lan controller, accesspoints
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"hostname": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Instance Uuid of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `Instance Tenant Id of the device
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

						"last_device_resync_start_time": &schema.Schema{
							Description: `Start time for last/ongoing sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_managed_resync_reasons": &schema.Schema{
							Description: `Reasons for last successful sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Time in epoch when the network device info last got updated
`,
							Type:     schema.TypeFloat,
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
							Description: `[Deprecated] Location ID that is associated with the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"location_name": &schema.Schema{
							Description: `[Deprecated] Name of the associated location
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mac_address": &schema.Schema{
							Description: `MAC address of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"managed_atleast_once": &schema.Schema{
							Description: `Indicates if device went into Managed state atleast once
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_ip_address": &schema.Schema{
							Description: `IP address of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"management_state": &schema.Schema{
							Description: `Represents the current management state of the network element: managed, unmanaged, under maintenance, and so on.
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

						"pending_sync_requests_count": &schema.Schema{
							Description: `Count of pending sync requests , if any
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"platform_id": &schema.Schema{
							Description: `Platform ID of device
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
							Description: `Device reachability status as Reachable / Unreachable
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reasons_for_device_resync": &schema.Schema{
							Description: `Reason for last/ongoing sync
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"reasons_for_pending_sync_requests": &schema.Schema{
							Description: `Reasons for pending sync requests , if any
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"role": &schema.Schema{
							Description: `Role of device as access, distribution, border router, core
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
							Description: `Serial number of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"series": &schema.Schema{
							Description: `Device series
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

						"software_type": &schema.Schema{
							Description: `Software type on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Description: `Software version on the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sync_requested_by_app": &schema.Schema{
							Description: `Applications which requested for the resync of network device
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
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type of device as switch, router, wireless lan controller, accesspoints
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"up_time": &schema.Schema{
							Description: `Time that shows for how long the device has been up
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"uptime_seconds": &schema.Schema{
							Description: `Uptime in Seconds
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"waas_device_mode": &schema.Schema{
							Description: `WAAS device mode
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

func dataSourceNetworkDeviceListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vHostname, okHostname := d.GetOk("hostname")
	vManagementIPAddress, okManagementIPAddress := d.GetOk("management_ip_address")
	vMacAddress, okMacAddress := d.GetOk("mac_address")
	vLocationName, okLocationName := d.GetOk("location_name")
	vSerialNumber, okSerialNumber := d.GetOk("serial_number")
	vLocation, okLocation := d.GetOk("location")
	vFamily, okFamily := d.GetOk("family")
	vType, okType := d.GetOk("type")
	vSeries, okSeries := d.GetOk("series")
	vCollectionStatus, okCollectionStatus := d.GetOk("collection_status")
	vCollectionInterval, okCollectionInterval := d.GetOk("collection_interval")
	vNotSyncedForMinutes, okNotSyncedForMinutes := d.GetOk("not_synced_for_minutes")
	vErrorCode, okErrorCode := d.GetOk("error_code")
	vErrorDescription, okErrorDescription := d.GetOk("error_description")
	vSoftwareVersion, okSoftwareVersion := d.GetOk("software_version")
	vSoftwareType, okSoftwareType := d.GetOk("software_type")
	vPlatformID, okPlatformID := d.GetOk("platform_id")
	vRole, okRole := d.GetOk("role")
	vReachabilityStatus, okReachabilityStatus := d.GetOk("reachability_status")
	vUpTime, okUpTime := d.GetOk("up_time")
	vAssociatedWlcIP, okAssociatedWlcIP := d.GetOk("associated_wlc_ip")
	vLicensename, okLicensename := d.GetOk("license_name")
	vLicensetypeR, okLicensetypeR := d.GetOk("license_type")
	vLicensestatus, okLicensestatus := d.GetOk("license_status")
	vModulename, okModulename := d.GetOk("module_name")
	vModuleequpimenttypeR, okModuleequpimenttypeR := d.GetOk("module_equpimenttype")
	vModuleservicestate, okModuleservicestate := d.GetOk("module_servicestate")
	vModulevendorequipmenttypeR, okModulevendorequipmenttypeR := d.GetOk("module_vendorequipmenttype")
	vModulepartnumber, okModulepartnumber := d.GetOk("module_partnumber")
	vModuleoperationstatecode, okModuleoperationstatecode := d.GetOk("module_operationstatecode")
	vID, okID := d.GetOk("id")
	vDeviceSupportLevel, okDeviceSupportLevel := d.GetOk("device_support_level")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceList")
		queryParams1 := catalystcentersdkgo.GetDeviceListQueryParams{}

		if okHostname {
			queryParams1.Hostname = interfaceToSliceString(vHostname)
		}
		if okManagementIPAddress {
			queryParams1.ManagementIPAddress = interfaceToSliceString(vManagementIPAddress)
		}
		if okMacAddress {
			queryParams1.MacAddress = interfaceToSliceString(vMacAddress)
		}
		if okLocationName {
			queryParams1.LocationName = interfaceToSliceString(vLocationName)
		}
		if okSerialNumber {
			queryParams1.SerialNumber = interfaceToSliceString(vSerialNumber)
		}
		if okLocation {
			queryParams1.Location = interfaceToSliceString(vLocation)
		}
		if okFamily {
			queryParams1.Family = interfaceToSliceString(vFamily)
		}
		if okType {
			queryParams1.Type = interfaceToSliceString(vType)
		}
		if okSeries {
			queryParams1.Series = interfaceToSliceString(vSeries)
		}
		if okCollectionStatus {
			queryParams1.CollectionStatus = interfaceToSliceString(vCollectionStatus)
		}
		if okCollectionInterval {
			queryParams1.CollectionInterval = interfaceToSliceString(vCollectionInterval)
		}
		if okNotSyncedForMinutes {
			queryParams1.NotSyncedForMinutes = interfaceToSliceString(vNotSyncedForMinutes)
		}
		if okErrorCode {
			queryParams1.ErrorCode = interfaceToSliceString(vErrorCode)
		}
		if okErrorDescription {
			queryParams1.ErrorDescription = interfaceToSliceString(vErrorDescription)
		}
		if okSoftwareVersion {
			queryParams1.SoftwareVersion = interfaceToSliceString(vSoftwareVersion)
		}
		if okSoftwareType {
			queryParams1.SoftwareType = interfaceToSliceString(vSoftwareType)
		}
		if okPlatformID {
			queryParams1.PlatformID = interfaceToSliceString(vPlatformID)
		}
		if okRole {
			queryParams1.Role = interfaceToSliceString(vRole)
		}
		if okReachabilityStatus {
			queryParams1.ReachabilityStatus = interfaceToSliceString(vReachabilityStatus)
		}
		if okUpTime {
			queryParams1.UpTime = interfaceToSliceString(vUpTime)
		}
		if okAssociatedWlcIP {
			queryParams1.AssociatedWlcIP = interfaceToSliceString(vAssociatedWlcIP)
		}
		if okLicensename {
			queryParams1.Licensename = interfaceToSliceString(vLicensename)
		}
		if okLicensetypeR {
			queryParams1.LicensetypeR = interfaceToSliceString(vLicensetypeR)
		}
		if okLicensestatus {
			queryParams1.Licensestatus = interfaceToSliceString(vLicensestatus)
		}
		if okModulename {
			queryParams1.Modulename = interfaceToSliceString(vModulename)
		}
		if okModuleequpimenttypeR {
			queryParams1.ModuleequpimenttypeR = interfaceToSliceString(vModuleequpimenttypeR)
		}
		if okModuleservicestate {
			queryParams1.Moduleservicestate = interfaceToSliceString(vModuleservicestate)
		}
		if okModulevendorequipmenttypeR {
			queryParams1.ModulevendorequipmenttypeR = interfaceToSliceString(vModulevendorequipmenttypeR)
		}
		if okModulepartnumber {
			queryParams1.Modulepartnumber = interfaceToSliceString(vModulepartnumber)
		}
		if okModuleoperationstatecode {
			queryParams1.Moduleoperationstatecode = interfaceToSliceString(vModuleoperationstatecode)
		}
		if okID {
			queryParams1.ID = vID.(string)
		}
		if okDeviceSupportLevel {
			queryParams1.DeviceSupportLevel = vDeviceSupportLevel.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Devices.GetDeviceList(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceList", err,
				"Failure at GetDeviceList, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceList", err,
				"Failure at GetDeviceList, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetDeviceListItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceList response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetDeviceListItems(items *[]catalystcentersdkgo.ResponseDevicesGetDeviceListResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
		respItem["reachability_status"] = item.ReachabilityStatus
		respItem["series"] = item.Series
		respItem["snmp_contact"] = item.SNMPContact
		respItem["snmp_location"] = item.SNMPLocation
		respItem["tag_count"] = item.TagCount
		respItem["tunnel_udp_port"] = item.TunnelUDPPort
		respItem["uptime_seconds"] = item.UptimeSeconds
		respItem["waas_device_mode"] = item.WaasDeviceMode
		respItem["serial_number"] = item.SerialNumber
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["mac_address"] = item.MacAddress
		respItem["up_time"] = item.UpTime
		respItem["device_support_level"] = item.DeviceSupportLevel
		respItem["hostname"] = item.Hostname
		respItem["type"] = item.Type
		respItem["memory_size"] = item.MemorySize
		respItem["family"] = item.Family
		respItem["error_code"] = item.ErrorCode
		respItem["software_type"] = item.SoftwareType
		respItem["software_version"] = item.SoftwareVersion
		respItem["description"] = item.Description
		respItem["role_source"] = item.RoleSource
		respItem["location"] = item.Location
		respItem["role"] = item.Role
		respItem["collection_interval"] = item.CollectionInterval
		respItem["inventory_status_detail"] = item.InventoryStatusDetail
		respItem["ap_ethernet_mac_address"] = item.ApEthernetMacAddress
		respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
		respItem["associated_wlc_ip"] = item.AssociatedWlcIP
		respItem["boot_date_time"] = item.BootDateTime
		respItem["collection_status"] = item.CollectionStatus
		respItem["error_description"] = item.ErrorDescription
		respItem["interface_count"] = item.InterfaceCount
		respItem["last_updated"] = item.LastUpdated
		respItem["line_card_count"] = item.LineCardCount
		respItem["line_card_id"] = item.LineCardID
		respItem["location_name"] = item.LocationName
		respItem["managed_atleast_once"] = boolPtrToString(item.ManagedAtleastOnce)
		respItem["management_ip_address"] = item.ManagementIPAddress
		respItem["platform_id"] = item.PlatformID
		respItem["management_state"] = item.ManagementState
		respItem["pending_sync_requests_count"] = item.PendingSyncRequestsCount
		respItem["reasons_for_device_resync"] = item.ReasonsForDeviceResync
		respItem["reasons_for_pending_sync_requests"] = item.ReasonsForPendingSyncRequests
		respItem["sync_requested_by_app"] = item.SyncRequestedByApp
		respItem["last_managed_resync_reasons"] = item.LastManagedResyncReasons
		respItem["dns_resolved_management_address"] = item.DNSResolvedManagementAddress
		respItem["last_device_resync_start_time"] = item.LastDeviceResyncStartTime
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}
