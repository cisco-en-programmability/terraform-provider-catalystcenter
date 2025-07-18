package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceApplicationVisibilityNetworkDevicesDisableAppTelemetry() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Application Policy.

- This data source action can be used to disable application telemetry feature on multiple network devices. Request
payload should include the list of network devices where it has to be disabled.
This operation pushes configuration to the network devices, and is only permitted if the provisioning settings do not
mandate a config preview for application telemetry disablement. In cases where such settings are active, attempting to
use this endpoint will result in **422 Unprocessable Content** error.
`,

		CreateContext: resourceApplicationVisibilityNetworkDevicesDisableAppTelemetryCreate,
		ReadContext:   resourceApplicationVisibilityNetworkDevicesDisableAppTelemetryRead,
		DeleteContext: resourceApplicationVisibilityNetworkDevicesDisableAppTelemetryDelete,
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
							Description: `Task Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"url": &schema.Schema{
							Description: `Url`,
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
						"network_device_ids": &schema.Schema{
							Description: `List of network device ids where Application Telemetry has to be disabled
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

func resourceApplicationVisibilityNetworkDevicesDisableAppTelemetryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestApplicationVisibilityNetworkDevicesDisableAppTelemetryDisableApplicationTelemetryFeatureOnMultipleNetworkDevices(ctx, "parameters.0", d)

	response1, restyResp1, err := client.ApplicationPolicy.DisableApplicationTelemetryFeatureOnMultipleNetworkDevices(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DisableApplicationTelemetryFeatureOnMultipleNetworkDevices", err))
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
			if restyResp3 == nil || strings.Contains(restyResp3.String(), "<!doctype html>") {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing DisableApplicationTelemetryFeatureOnMultipleNetworkDevices", err1))
			return diags
		}
	}

	vItem1 := flattenApplicationPolicyDisableApplicationTelemetryFeatureOnMultipleNetworkDevicesItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DisableApplicationTelemetryFeatureOnMultipleNetworkDevices response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceApplicationVisibilityNetworkDevicesDisableAppTelemetryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceApplicationVisibilityNetworkDevicesDisableAppTelemetryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestApplicationVisibilityNetworkDevicesDisableAppTelemetryDisableApplicationTelemetryFeatureOnMultipleNetworkDevices(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyDisableApplicationTelemetryFeatureOnMultipleNetworkDevices {
	request := catalystcentersdkgo.RequestApplicationPolicyDisableApplicationTelemetryFeatureOnMultipleNetworkDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_ids")))) {
		request.NetworkDeviceIDs = interfaceToSliceString(v)
	}
	return &request
}

func flattenApplicationPolicyDisableApplicationTelemetryFeatureOnMultipleNetworkDevicesItem(item *catalystcentersdkgo.ResponseApplicationPolicyDisableApplicationTelemetryFeatureOnMultipleNetworkDevicesResponse) []map[string]interface{} {
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
