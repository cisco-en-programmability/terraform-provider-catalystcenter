package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceIntentNetworkDevicesQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Returns the list of network devices, determined by the filters. It is possible to filter the network devices based on
various parameters, such as device type, device role, software version, etc.
`,

		CreateContext: resourceIntentNetworkDevicesQueryCreate,
		ReadContext:   resourceIntentNetworkDevicesQueryRead,
		DeleteContext: resourceIntentNetworkDevicesQueryDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"filter": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Description: `The key to filter by
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"operator": &schema.Schema{
													Description: `The operator to use for filtering the values
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"value": &schema.Schema{
													Description: `Value to filter by. For in operator, the value should be a list of values.
`,
													Type:     schema.TypeString, //TEST,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"logical_operator": &schema.Schema{
										Description: `The logical operator to use for combining the filter criteria. If not provided, the default value is AND.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_ethernet_mac_address": &schema.Schema{
										Description: `Ethernet MAC address of the AP network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"ap_manager_interface_ip_address": &schema.Schema{
										Description: `Management IP address of the AP network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"ap_wlc_ip_address": &schema.Schema{
										Description: `Management IP address of the WLC on which AP is associated to
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"boot_time": &schema.Schema{
										Description: `The time at which the network device was last rebooted or powered on represented as epoch in milliseconds
`,
										Type:     schema.TypeFloat,
										ForceNew: true,
										Computed: true,
									},
									"device_support_level": &schema.Schema{
										Description: `The level of support Catalyst Center provides for the network device.
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"dns_resolved_management_ip_address": &schema.Schema{
										Description: `DNS-resolved management IP address of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"error_code": &schema.Schema{
										Description: `Error code indicating the reason for the last resync failure
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"error_description": &schema.Schema{
										Description: `Additional information regarding the reason for resync failure. This is a human-readable error message and should not be expected programmatically.
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"family": &schema.Schema{
										Description: `Product family of the network device. For example, Switches, Routers, etc
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"hostname": &schema.Schema{
										Description: `Hostname of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Unique identifier of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"last_successful_resync_reasons": &schema.Schema{
										Description: `List of reasons for the last successful resync of the device. If multiple resync requests are made before the device can start the resync, all the reasons will be captured. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC
`,
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"mac_address": &schema.Schema{
										Description: `MAC address of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"management_address": &schema.Schema{
										Description: `Management address of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"management_state": &schema.Schema{
										Description: `The status of the network device's manageability. Refer features for more details.
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"pending_resync_request_count": &schema.Schema{
										Description: `Number of pending resync requests for the device
`,
										Type:     schema.TypeInt,
										ForceNew: true,
										Computed: true,
									},
									"pending_resync_request_reasons": &schema.Schema{
										Description: `List of reasons for the pending resync requests. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC
`,
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"platform_ids": &schema.Schema{
										Description: `Platform identifier of the network device
`,
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"reachability_failure_reason": &schema.Schema{
										Description: `Reason for reachability failure. This message that provides more information about the reachability failure.
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"reachability_status": &schema.Schema{
										Description: `Reachability status of the network device. Refer features for more details
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"resync_end_time": &schema.Schema{
										Description: `End time for the last resync represented as epoch in milliseconds
`,
										Type:     schema.TypeFloat,
										ForceNew: true,
										Computed: true,
									},
									"resync_interval_minutes": &schema.Schema{
										Description: `The duration in minutes between the periodic resync attempts for the device
`,
										Type:     schema.TypeInt,
										ForceNew: true,
										Computed: true,
									},
									"resync_interval_source": &schema.Schema{
										Description: `Source of the resync interval. Note: Please refer to PUT /dna/intent/api/v1/networkDevices/resyncIntervalSettings API to update the global resync interval.
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"resync_reasons": &schema.Schema{
										Description: `List of reasons for the ongoing/last resync on the device. If multiple resync requests were made before the resync could start, all the reasons will be captured as an array. Possible values: ADD_DEVICE_SYNC, LINK_UP_DOWN, CONFIG_CHANGE, DEVICE_UPDATED_SYNC, AP_EVENT_BASED_SYNC, APP_REQUESTED_SYNC, PERIODIC_SYNC, UI_SYNC, CUSTOM, UNKNOWN, REFRESH_OBJECTS_FEATURE_BASED_SYNC
`,
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"resync_requested_by_apps": &schema.Schema{
										Description: `List of applications that requested the last/ongoing resync on the device
`,
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"resync_start_time": &schema.Schema{
										Description: `Start time for the last/ongoing resync represented as epoch in milliseconds
`,
										Type:     schema.TypeFloat,
										ForceNew: true,
										Computed: true,
									},
									"role": &schema.Schema{
										Description: `Role assigned to the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"role_source": &schema.Schema{
										Description: `Indicates whether the network device's role was assigned automatically by the software or manually by an administrator.
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"serial_numbers": &schema.Schema{
										Description: `Serial number of the network device. In case of stack device, there will be multiple serial numbers
`,
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"series": &schema.Schema{
										Description: `The model range or series of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"snmp_contact": &schema.Schema{
										Description: `SNMP contact of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"snmp_location": &schema.Schema{
										Description: `SNMP location of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"software_type": &schema.Schema{
										Description: `Type of software running on the network device. For example, IOS-XE, etc.
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"software_version": &schema.Schema{
										Description: `Version of the software running on the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"stack_device": &schema.Schema{
										Description: `Flag indicating if the network device is a stack device
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `Inventory related status of the network device. Refer features for more details
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Description: `Type of the network device. This list of types can be obtained from the API intent/networkDeviceProductNames productName field.
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"user_defined_fields": &schema.Schema{
										Description: `Map of all user defined fields and their values associated with the device. Refer to /dna/intent/api/v1/network-device/user-defined-field API to fetch all the user defined fields.
`,
										Type:     schema.TypeString, //TEST,
										ForceNew: true,
										Computed: true,
									},
									"vendor": &schema.Schema{
										Description: `Vendor of the network device
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"page": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"limit": &schema.Schema{
										Description: `The number of records to show for this page. Min: 1, Max: 500
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"offset": &schema.Schema{
										Description: `The first record to show for this page; the first record is numbered 1.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"sort_by": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `The field to sort by. Default is hostname.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"order": &schema.Schema{
													Description: `The order to sort by.
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"views": &schema.Schema{
							Description: `The specific views being requested. This is an optional parameter which can be passed to get one or more of the network device data. If this is not provided, then it will default to BASIC views. If multiple views are provided, the response will contain the union of the views.
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceIntentNetworkDevicesQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.QueryNetworkDevicesWithFiltersV1(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vItems1 := flattenDevicesQueryNetworkDevicesWithFiltersV1Items(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting QueryNetworkDevicesWithFiltersV1 response",
			err))
		return diags
	}

	//Analizar verificacion.

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return diags
}
func resourceIntentNetworkDevicesQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIntentNetworkDevicesQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1 {
	request := catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1Filter(ctx, key+".filter.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".views")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".views")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".views")))) {
		request.Views = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1Page(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1Filter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1Filter {
	request := catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1Filter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".logical_operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".logical_operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".logical_operator")))) {
		request.LogicalOperator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1FilterFiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1FilterFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFilters {
	request := []catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFilters{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1FilterFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1FilterFilters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFilters {
	request := catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1FilterFiltersValue(ctx, key+".value.0", d)
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1FilterFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFiltersValue {
	var request catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1FilterFiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1Page(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1Page {
	request := catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1Page{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1PageSortBy(ctx, key+".sort_by.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryQueryNetworkDevicesWithFiltersV1PageSortBy(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1PageSortBy {
	request := catalystcentersdkgo.RequestDevicesQueryNetworkDevicesWithFiltersV1PageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesQueryNetworkDevicesWithFiltersV1Items(items *[]catalystcentersdkgo.ResponseDevicesQueryNetworkDevicesWithFiltersV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["management_address"] = item.ManagementAddress
		respItem["dns_resolved_management_ip_address"] = item.DNSResolvedManagementIPAddress
		respItem["hostname"] = item.Hostname
		respItem["mac_address"] = item.MacAddress
		respItem["serial_numbers"] = item.SerialNumbers
		respItem["type"] = item.Type
		respItem["family"] = item.Family
		respItem["series"] = item.Series
		respItem["status"] = item.Status
		respItem["platform_ids"] = item.PlatformIDs
		respItem["software_type"] = item.SoftwareType
		respItem["software_version"] = item.SoftwareVersion
		respItem["vendor"] = item.Vendor
		respItem["stack_device"] = boolPtrToString(item.StackDevice)
		respItem["boot_time"] = item.BootTime
		respItem["role"] = item.Role
		respItem["role_source"] = item.RoleSource
		respItem["ap_ethernet_mac_address"] = item.ApEthernetMacAddress
		respItem["ap_manager_interface_ip_address"] = item.ApManagerInterfaceIPAddress
		respItem["ap_wlc_ip_address"] = item.ApWlcIPAddress
		respItem["device_support_level"] = item.DeviceSupportLevel
		respItem["snmp_location"] = item.SNMPLocation
		respItem["snmp_contact"] = item.SNMPContact
		respItem["reachability_status"] = item.ReachabilityStatus
		respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
		respItem["management_state"] = item.ManagementState
		respItem["last_successful_resync_reasons"] = item.LastSuccessfulResyncReasons
		respItem["resync_start_time"] = item.ResyncStartTime
		respItem["resync_end_time"] = item.ResyncEndTime
		respItem["resync_reasons"] = item.ResyncReasons
		respItem["resync_requested_by_apps"] = item.ResyncRequestedByApps
		respItem["pending_resync_request_count"] = item.PendingResyncRequestCount
		respItem["pending_resync_request_reasons"] = item.PendingResyncRequestReasons
		respItem["resync_interval_source"] = item.ResyncIntervalSource
		respItem["resync_interval_minutes"] = item.ResyncIntervalMinutes
		respItem["error_code"] = item.ErrorCode
		respItem["error_description"] = item.ErrorDescription
		respItem["user_defined_fields"] = flattenDevicesQueryNetworkDevicesWithFiltersV1ItemsUserDefinedFields(item.UserDefinedFields)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesQueryNetworkDevicesWithFiltersV1ItemsUserDefinedFields(item *catalystcentersdkgo.ResponseDevicesQueryNetworkDevicesWithFiltersV1ResponseUserDefinedFields) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
