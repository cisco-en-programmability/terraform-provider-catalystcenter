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
func resourceAssignToSiteApply() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Assign unprovisioned network devices to a site. Along with that it can also be used to assign unprovisioned network
devices to a different site. If device controllability is enabled, it will be triggered once device assigned to site
successfully. Device Controllability can be enabled/disabled using
**/dna/intent/api/v1/networkDevices/deviceControllability/settings**.
`,

		CreateContext: resourceAssignToSiteApplyCreate,
		ReadContext:   resourceAssignToSiteApplyRead,
		DeleteContext: resourceAssignToSiteApplyDelete,
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
							Description: `Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
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
						"device_ids": &schema.Schema{
							Description: `Unassigned network devices, ranging from a minimum of 1 to a maximum of 100.
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"site_id": &schema.Schema{
							Description: `This must be building Id or floor Id. Access points, Sensors are assigned to floor. Remaining network devices are assigned to building. Site Id can be retrieved using '/intent/api/v1/sites'.
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

func resourceAssignToSiteApplyCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestAssignToSiteApplyAssignNetworkDevicesToASite(ctx, "parameters.0", d)

	response1, restyResp1, err := client.SiteDesign.AssignNetworkDevicesToASite(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ApplicationPolicyIntent", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignNetworkDevicesToASite", err))
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
				"Failure when executing AssignNetworkDevicesToASite", err1))
			return diags
		}
	}

	vItem1 := flattenSiteDesignAssignNetworkDevicesToASiteItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignNetworkDevicesToASite response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags

}
func resourceAssignToSiteApplyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAssignToSiteApplyDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAssignToSiteApplyAssignNetworkDevicesToASite(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignAssignNetworkDevicesToASite {
	request := catalystcentersdkgo.RequestSiteDesignAssignNetworkDevicesToASite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_ids")))) {
		request.DeviceIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	return &request
}

func flattenSiteDesignAssignNetworkDevicesToASiteItem(item *catalystcentersdkgo.ResponseSiteDesignAssignNetworkDevicesToASiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	respItem["task_id"] = item.TaskID
	return []map[string]interface{}{
		respItem,
	}
}
