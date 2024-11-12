package catalystcenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceLanAutomationUpdateDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on LAN Automation.

- Invoke this API to perform a DAY-N update on LAN Automation-related devices. Supported features include Loopback0 IP
update, hostname update, link addition, and link deletion.
`,

		CreateContext: resourceLanAutomationUpdateDeviceCreate,
		ReadContext:   resourceLanAutomationUpdateDeviceRead,
		DeleteContext: resourceLanAutomationUpdateDeviceDelete,
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

						"task_id": &schema.Schema{
							Description: `Task ID 
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `url to check the status of task
`,
							Type:     schema.TypeString,
							Computed: true,
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
						"feature": &schema.Schema{
							Description: `feature query parameter. Feature ID for the update. Supported feature IDs include: LOOPBACK0_IPADDRESS_UPDATE, HOSTNAME_UPDATE, LINK_ADD, and LINK_DELETE. 
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"hostname_update_devices": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_management_ipaddress": &schema.Schema{
										Description: `Device Management IP Address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"new_host_name": &schema.Schema{
										Description: `New hostname for the device
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"link_update": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"destination_device_interface_name": &schema.Schema{
										Description: `Destination Device Interface Name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"destination_device_management_ipaddress": &schema.Schema{
										Description: `Destination Device Management IP Address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ip_pool_name": &schema.Schema{
										Description: `Name of the IP LAN Pool, required for Link Add should be from discovery site of source and destination device.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"source_device_interface_name": &schema.Schema{
										Description: `Source Device Interface Name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"source_device_management_ipaddress": &schema.Schema{
										Description: `Source Device Management IP Address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"loopback_update_device_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_management_ipaddress": &schema.Schema{
										Description: `Device Management IP Address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"new_loopback0_ipaddress": &schema.Schema{
										Description: `New Loopback0 IP Address from LAN Pool of Device Discovery Site(Shared pool should not be used).
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
	}
}

func resourceLanAutomationUpdateDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vFeature := resourceItem["feature"]

	request1 := expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1(ctx, "parameters.0", d)
	queryParams1 := catalystcentersdkgo.LanAutomationDeviceUpdateV1QueryParams{}

	queryParams1.Feature = vFeature.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.LanAutomation.LanAutomationDeviceUpdateV1(request1, &queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing LanAutomationDeviceUpdateV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing LANAutomationDeviceUpdateV1", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing LANAutomationDeviceUpdateV1", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenLanAutomationLanAutomationDeviceUpdateV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting LanAutomationDeviceUpdateV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceLanAutomationUpdateDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceLanAutomationUpdateDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1 {
	request := catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".loopback_update_device_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".loopback_update_device_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".loopback_update_device_list")))) {
		request.LoopbackUpdateDeviceList = expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1LoopbackUpdateDeviceListArray(ctx, key+".loopback_update_device_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".link_update")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".link_update")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".link_update")))) {
		request.LinkUpdate = expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1LinkUpdate(ctx, key+".link_update.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname_update_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname_update_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname_update_devices")))) {
		request.HostnameUpdateDevices = expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1HostnameUpdateDevicesArray(ctx, key+".hostname_update_devices", d)
	}
	return &request
}

func expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1LoopbackUpdateDeviceListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1LoopbackUpdateDeviceList {
	request := []catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1LoopbackUpdateDeviceList{}
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
		i := expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1LoopbackUpdateDeviceList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1LoopbackUpdateDeviceList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1LoopbackUpdateDeviceList {
	request := catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1LoopbackUpdateDeviceList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ipaddress")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_loopback0_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_loopback0_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_loopback0_ipaddress")))) {
		request.NewLoopback0IPAddress = interfaceToString(v)
	}
	return &request
}

func expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1LinkUpdate(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1LinkUpdate {
	request := catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1LinkUpdate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_device_management_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_device_management_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_device_management_ipaddress")))) {
		request.SourceDeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_device_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_device_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_device_interface_name")))) {
		request.SourceDeviceInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".destination_device_management_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".destination_device_management_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".destination_device_management_ipaddress")))) {
		request.DestinationDeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".destination_device_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".destination_device_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".destination_device_interface_name")))) {
		request.DestinationDeviceInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	return &request
}

func expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1HostnameUpdateDevicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1HostnameUpdateDevices {
	request := []catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1HostnameUpdateDevices{}
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
		i := expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1HostnameUpdateDevices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestLanAutomationUpdateDeviceLanAutomationDeviceUpdateV1HostnameUpdateDevices(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1HostnameUpdateDevices {
	request := catalystcentersdkgo.RequestLanAutomationLanAutomationDeviceUpdateV1HostnameUpdateDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ipaddress")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".new_host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".new_host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".new_host_name")))) {
		request.NewHostName = interfaceToString(v)
	}
	return &request
}

func flattenLanAutomationLanAutomationDeviceUpdateV1Item(item *catalystcentersdkgo.ResponseLanAutomationLanAutomationDeviceUpdateV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
