package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkDevicesInterfacesQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- This data source action returns the Interface Stats for the given Device Id. Please refer to the Feature tab for the
Request Body usage and the API filtering support.
`,

		CreateContext: resourceNetworkDevicesInterfacesQueryCreate,
		ReadContext:   resourceNetworkDevicesInterfacesQueryRead,
		DeleteContext: resourceNetworkDevicesInterfacesQueryDelete,
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
						"device_id": &schema.Schema{
							Description: `deviceId path parameter. Network Device Id
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"end_time": &schema.Schema{
							Description: `UTC epoch timestamp in milliseconds
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Interface Instance Id
`,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"values": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"admin_status": &schema.Schema{
													Description: `The desired state of the interface
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Interface description
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"device_id": &schema.Schema{
													Description: `Device Id
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"duplex_config": &schema.Schema{
													Description: `Interface duplex config status
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"duplex_oper": &schema.Schema{
													Description: `Interface duplex operational status
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"instance_id": &schema.Schema{
													Description: `Interface InstanceId
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"interface_id": &schema.Schema{
													Description: `Interface ifIndex
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"interface_type": &schema.Schema{
													Description: `Physical or Virtual type
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"ipv4_address": &schema.Schema{
													Description: `Interface IPV4 Address
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"ipv6_address_list": &schema.Schema{
													Description: `List of interface IPV6 Address
`,
													Type:     schema.TypeList,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"is_l3_interface": &schema.Schema{
													Description: `Interface is L3 or not
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"is_wan": &schema.Schema{
													Description: `nterface is WAN link or not
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"mac_addr": &schema.Schema{
													Description: `Interface MAC Address
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"media_type": &schema.Schema{
													Description: `Interface media type
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Name of the interface
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"oper_status": &schema.Schema{
													Description: `Interface operational status
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"peer_stack_member": &schema.Schema{
													Description: `Interface peer stack member Id
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"peer_stack_port": &schema.Schema{
													Description: `Interface peer stack member port
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"port_channel_id": &schema.Schema{
													Description: `Interface Port-Channel Id
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"port_mode": &schema.Schema{
													Description: `Interface Port Mode
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"port_type": &schema.Schema{
													Description: `Interface ifType
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"rx_discards": &schema.Schema{
													Description: `Rx Discards in %
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"rx_error": &schema.Schema{
													Description: `Rx Errors in %
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"rx_rate": &schema.Schema{
													Description: `Rx rate in bps
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"rx_utilization": &schema.Schema{
													Description: `Rx Utilization in %
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"speed": &schema.Schema{
													Description: `Speed of the Interface in kbps
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"stack_port_type": &schema.Schema{
													Description: `Interface stack port type. SVL or DAD
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"timestamp": &schema.Schema{
													Description: `Interface stats collected timestamp
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"tx_discards": &schema.Schema{
													Description: `Tx Discards in %
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"tx_error": &schema.Schema{
													Description: `Tx Errors in %
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"tx_rate": &schema.Schema{
													Description: `Tx Rate in bps
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"tx_utilization": &schema.Schema{
													Description: `Tx  Utilization in %
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
												"vlan_id": &schema.Schema{
													Description: `Interface VLAN Id
`,
													Type:     schema.TypeString,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},
						"query": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"fields": &schema.Schema{
										Description: `Required field names, default ALL
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Description: `Name of the field that the filter should be applied to
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"operator": &schema.Schema{
													Description: `Supported operators are eq,in,like
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"value": &schema.Schema{
													Description: `Value of the field
`,
													Type:     schema.TypeString,
													Optional: true,
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
													Description: `Number of records, Max is 1000
`,
													Type:     schema.TypeInt,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"offset": &schema.Schema{
													Description: `Record offset value, default 0
`,
													Type:     schema.TypeFloat,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"order_by": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"name": &schema.Schema{
																Description: `Name of the field used to sort
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"order": &schema.Schema{
																Description: `Possible values asc, des
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
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `UTC epoch timestamp in milliseconds
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDevicesInterfacesQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vDeviceID := resourceItem["device_id"]

	vvDeviceID := vDeviceID.(string)
	request1 := expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfo(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Devices.GetDeviceInterfaceStatsInfo(vvDeviceID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing GetDeviceInterfaceStatsInfo", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	//Analizar verificacion.

	vItems1 := flattenDevicesGetDeviceInterfaceStatsInfoItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetDeviceInterfaceStatsInfo response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceNetworkDevicesInterfacesQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDevicesInterfacesQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfo {
	request := catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".query")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".query")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".query")))) {
		request.Query = expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQuery(ctx, key+".query.0", d)
	}
	return &request
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQuery(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQuery {
	request := catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQuery{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fields")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fields")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fields")))) {
		request.Fields = expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryFieldsArray(ctx, key+".fields", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryFieldsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryFields {
	request := []catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryFields{}
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
		i := expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryFields(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryFields(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryFields {
	var request catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryFields
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryFilters {
	request := []catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryFilters{}
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
		i := expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryFilters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryFilters {
	request := catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryPage(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryPage {
	request := catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order_by")))) {
		request.OrderBy = expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryPageOrderByArray(ctx, key+".order_by", d)
	}
	return &request
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryPageOrderByArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryPageOrderBy {
	request := []catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryPageOrderBy{}
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
		i := expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryPageOrderBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesInterfacesQueryGetDeviceInterfaceStatsInfoQueryPageOrderBy(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryPageOrderBy {
	request := catalystcentersdkgo.RequestDevicesGetDeviceInterfaceStatsInfoQueryPageOrderBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetDeviceInterfaceStatsInfoItems(items *[]catalystcentersdkgo.ResponseDevicesGetDeviceInterfaceStatsInfoResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["values"] = flattenDevicesGetDeviceInterfaceStatsInfoItemsValues(item.Values)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetDeviceInterfaceStatsInfoItemsValues(item *catalystcentersdkgo.ResponseDevicesGetDeviceInterfaceStatsInfoResponseValues) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["device_id"] = item.DeviceID
	respItem["duplex_config"] = item.DuplexConfig
	respItem["duplex_oper"] = item.DuplexOper
	respItem["interface_id"] = item.InterfaceID
	respItem["interface_type"] = item.InterfaceType
	respItem["instance_id"] = item.InstanceID
	respItem["ipv4_address"] = item.IPv4Address
	respItem["ipv6_address_list"] = item.IPv6AddressList
	respItem["is_l3_interface"] = item.IsL3Interface
	respItem["is_wan"] = item.IsWan
	respItem["mac_addr"] = item.MacAddr
	respItem["media_type"] = item.MediaType
	respItem["name"] = item.Name
	respItem["oper_status"] = item.OperStatus
	respItem["peer_stack_member"] = item.PeerStackMember
	respItem["peer_stack_port"] = item.PeerStackPort
	respItem["port_channel_id"] = item.PortChannelID
	respItem["port_mode"] = item.PortMode
	respItem["port_type"] = item.PortType
	respItem["description"] = item.Description
	respItem["rx_discards"] = item.RxDiscards
	respItem["rx_error"] = item.RxError
	respItem["rx_rate"] = item.RxRate
	respItem["rx_utilization"] = item.RxUtilization
	respItem["speed"] = item.Speed
	respItem["stack_port_type"] = item.StackPortType
	respItem["timestamp"] = item.Timestamp
	respItem["tx_discards"] = item.TxDiscards
	respItem["tx_error"] = item.TxError
	respItem["tx_rate"] = item.TxRate
	respItem["tx_utilization"] = item.TxUtilization
	respItem["vlan_id"] = item.VLANID

	return []map[string]interface{}{
		respItem,
	}

}
