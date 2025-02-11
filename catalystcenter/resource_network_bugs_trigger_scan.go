package catalystcenter

import (
	"context"

	"errors"

	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkBugsTriggerScan() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Compliance.

- Triggers a bugs scan for the supported network devices. The supported devices are switches and routers. If a device is
not supported, the NetworkBugsDevice scanStatus will be Failed with appropriate comments.
`,

		CreateContext: resourceNetworkBugsTriggerScanCreate,
		ReadContext:   resourceNetworkBugsTriggerScanRead,
		DeleteContext: resourceNetworkBugsTriggerScanDelete,
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
						"failed_devices_only": &schema.Schema{
							Description: `failedDevicesOnly query parameter. Used to specify if the scan should run only for the network devices that failed during the previous scan. If not specified, this parameter defaults to false.
`,
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkBugsTriggerScanCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vFailedDevicesOnly := resourceItem["failed_devices_only"]
	vvFailedDevicesOnly := interfaceToBool(vFailedDevicesOnly)

	queryParams1 := catalystcentersdkgo.TriggersABugsScanForTheSupportedNetworkDevicesV1QueryParams{}
	queryParams1.FailedDevicesOnly = vvFailedDevicesOnly
	// has_unknown_response: None

	response1, restyResp1, err := client.Compliance.TriggersABugsScanForTheSupportedNetworkDevicesV1(&queryParams1)

	vItem1 := flattenComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TriggersABugsScanForTheSupportedNetworkDevicesV1 response",
			err))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing TriggersABugsScanForTheSupportedNetworkDevicesV1", err))
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
				"Failure when executing TriggersABugsScanForTheSupportedNetworkDevicesV1", err1))
			return diags
		}
	}

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
func resourceNetworkBugsTriggerScanRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkBugsTriggerScanDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1Item(item *catalystcentersdkgo.ResponseComplianceTriggersABugsScanForTheSupportedNetworkDevicesV1Response) []map[string]interface{} {
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
