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
func resourceWirelessAccessPointsProvision() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- This data source action is used to provision access points
`,

		CreateContext: resourceWirelessAccessPointsProvisionCreate,
		ReadContext:   resourceWirelessAccessPointsProvisionRead,
		DeleteContext: resourceWirelessAccessPointsProvisionDelete,
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
							Description: `Task URL
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
						"ap_zone_name": &schema.Schema{
							Description: `AP Zone Name. A custom AP Zone should be passed if no rfProfileName is provided.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"network_devices": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_id": &schema.Schema{
										Description: `Network device ID of access points
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mesh_role": &schema.Schema{
										Description: `Mesh Role (Applicable only when AP is in Bridge Mode)
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"rf_profile_name": &schema.Schema{
							Description: `RF Profile Name. RF Profile is not allowed for custom AP Zones.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `Site ID
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
	}
}

func resourceWirelessAccessPointsProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestWirelessAccessPointsProvisionApProvisionV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.ApProvisionV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ApProvisionV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing APProvisionV1", err))
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
				"Failure when executing APProvisionV1", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenWirelessApProvisionV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ApProvisionV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessAccessPointsProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessAccessPointsProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessAccessPointsProvisionApProvisionV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessApProvisionV1 {
	request := catalystcentersdkgo.RequestWirelessApProvisionV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_devices")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_devices")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_devices")))) {
		request.NetworkDevices = expandRequestWirelessAccessPointsProvisionApProvisionV1NetworkDevicesArray(ctx, key+".network_devices", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_profile_name")))) {
		request.RfProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_zone_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_zone_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_zone_name")))) {
		request.ApZoneName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	return &request
}

func expandRequestWirelessAccessPointsProvisionApProvisionV1NetworkDevicesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestWirelessApProvisionV1NetworkDevices {
	request := []catalystcentersdkgo.RequestWirelessApProvisionV1NetworkDevices{}
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
		i := expandRequestWirelessAccessPointsProvisionApProvisionV1NetworkDevices(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessAccessPointsProvisionApProvisionV1NetworkDevices(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessApProvisionV1NetworkDevices {
	request := catalystcentersdkgo.RequestWirelessApProvisionV1NetworkDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mesh_role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mesh_role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mesh_role")))) {
		request.MeshRole = interfaceToString(v)
	}
	return &request
}

func flattenWirelessApProvisionV1Item(item *catalystcentersdkgo.ResponseWirelessApProvisionV1Response) []map[string]interface{} {
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
