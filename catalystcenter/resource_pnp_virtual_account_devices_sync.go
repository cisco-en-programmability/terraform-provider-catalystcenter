package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePnpVirtualAccountDevicesSync() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Synchronizes the device info from the given smart account & virtual account with the PnP database. The response
payload returns a list of synced devices (Deprecated).
`,

		CreateContext: resourcePnpVirtualAccountDevicesSyncCreate,
		ReadContext:   resourcePnpVirtualAccountDevicesSyncRead,
		DeleteContext: resourcePnpVirtualAccountDevicesSyncDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auto_sync_period": &schema.Schema{
							Description: `Auto Sync Period`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"cco_user": &schema.Schema{
							Description: `Cco User`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"expiry": &schema.Schema{
							Description: `Expiry`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"last_sync": &schema.Schema{
							Description: `Last Sync`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address_fqdn": &schema.Schema{
										Description: `Address Fqdn`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"address_ip_v4": &schema.Schema{
										Description: `Address Ip V4`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"cert": &schema.Schema{
										Description: `Cert`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"make_default": &schema.Schema{
										Description: `Make Default`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"profile_id": &schema.Schema{
										Description: `Profile Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"proxy": &schema.Schema{
										Description: `Proxy`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"smart_account_id": &schema.Schema{
							Description: `Smart Account Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sync_result": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"sync_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_sn_list": &schema.Schema{
													Description: `Device Sn List`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"sync_type": &schema.Schema{
													Description: `Sync Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"sync_msg": &schema.Schema{
										Description: `Sync Msg`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"sync_result_str": &schema.Schema{
							Description: `Sync Result Str`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"sync_start_time": &schema.Schema{
							Description: `Sync Start Time`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"sync_status": &schema.Schema{
							Description: `Sync Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"token": &schema.Schema{
							Description: `Token`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"virtual_account_id": &schema.Schema{
							Description: `Virtual Account Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_sync_period": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"cco_user": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"expiry": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"last_sync": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"profile": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"address_fqdn": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"address_ip_v4": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"cert": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"make_default": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"port": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"profile_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
								},
							},
						},
						"smart_account_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"sync_result": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"sync_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_sn_list": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"sync_type": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"sync_msg": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"sync_result_str": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"sync_start_time": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"sync_status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"token": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"virtual_account_id": &schema.Schema{
							Type:     schema.TypeString,
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

func resourcePnpVirtualAccountDevicesSyncCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevices(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.DeviceOnboardingPnp.SyncVirtualAccountDevices(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing SyncVirtualAccountDevices", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDeviceOnboardingPnpSyncVirtualAccountDevicesItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting SyncVirtualAccountDevices response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourcePnpVirtualAccountDevicesSyncRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpVirtualAccountDevicesSyncDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevices(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevices {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_sync_period")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_sync_period")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_sync_period")))) {
		request.AutoSyncPeriod = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cco_user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cco_user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cco_user")))) {
		request.CcoUser = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".expiry")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".expiry")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".expiry")))) {
		request.Expiry = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_sync")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_sync")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_sync")))) {
		request.LastSync = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile")))) {
		request.Profile = expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevicesProfile(ctx, key+".profile.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smart_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smart_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smart_account_id")))) {
		request.SmartAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_result")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_result")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_result")))) {
		request.SyncResult = expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevicesSyncResult(ctx, key+".sync_result.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_result_str")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_result_str")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_result_str")))) {
		request.SyncResultStr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_start_time")))) {
		request.SyncStartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_status")))) {
		request.SyncStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tenant_id")))) {
		request.TenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".token")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".token")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".token")))) {
		request.Token = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_account_id")))) {
		request.VirtualAccountID = interfaceToString(v)
	}
	return &request
}

func expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevicesProfile(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevicesProfile {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevicesProfile{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address_fqdn")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address_fqdn")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address_fqdn")))) {
		request.AddressFqdn = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address_ip_v4")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address_ip_v4")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address_ip_v4")))) {
		request.AddressIPV4 = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cert")))) {
		request.Cert = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".make_default")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".make_default")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".make_default")))) {
		request.MakeDefault = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_id")))) {
		request.ProfileID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy")))) {
		request.Proxy = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevicesSyncResult(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResult {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResult{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_list")))) {
		request.SyncList = expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevicesSyncResultSyncListArray(ctx, key+".sync_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_msg")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_msg")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_msg")))) {
		request.SyncMsg = interfaceToString(v)
	}
	return &request
}

func expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevicesSyncResultSyncListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList {
	request := []catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList{}
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
		i := expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevicesSyncResultSyncList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpVirtualAccountDevicesSyncSyncVirtualAccountDevicesSyncResultSyncList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_sn_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_sn_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_sn_list")))) {
		request.DeviceSnList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sync_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sync_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sync_type")))) {
		request.SyncType = interfaceToString(v)
	}
	return &request
}

func flattenDeviceOnboardingPnpSyncVirtualAccountDevicesItem(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpSyncVirtualAccountDevices) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["virtual_account_id"] = item.VirtualAccountID
	respItem["auto_sync_period"] = item.AutoSyncPeriod
	respItem["sync_result_str"] = item.SyncResultStr
	respItem["profile"] = flattenDeviceOnboardingPnpSyncVirtualAccountDevicesItemProfile(item.Profile)
	respItem["cco_user"] = item.CcoUser
	respItem["sync_result"] = flattenDeviceOnboardingPnpSyncVirtualAccountDevicesItemSyncResult(item.SyncResult)
	respItem["token"] = item.Token
	respItem["sync_start_time"] = item.SyncStartTime
	respItem["last_sync"] = item.LastSync
	respItem["tenant_id"] = item.TenantID
	respItem["smart_account_id"] = item.SmartAccountID
	respItem["expiry"] = item.Expiry
	respItem["sync_status"] = item.SyncStatus
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpSyncVirtualAccountDevicesItemProfile(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesProfile) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["proxy"] = boolPtrToString(item.Proxy)
	respItem["make_default"] = boolPtrToString(item.MakeDefault)
	respItem["port"] = item.Port
	respItem["profile_id"] = item.ProfileID
	respItem["name"] = item.Name
	respItem["address_ip_v4"] = item.AddressIPV4
	respItem["cert"] = item.Cert
	respItem["address_fqdn"] = item.AddressFqdn

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpSyncVirtualAccountDevicesItemSyncResult(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResult) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["sync_list"] = flattenDeviceOnboardingPnpSyncVirtualAccountDevicesItemSyncResultSyncList(item.SyncList)
	respItem["sync_msg"] = item.SyncMsg

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpSyncVirtualAccountDevicesItemSyncResultSyncList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpSyncVirtualAccountDevicesSyncResultSyncList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["sync_type"] = item.SyncType
		respItem["device_sn_list"] = item.DeviceSnList
		respItems = append(respItems, respItem)
	}
	return respItems
}
