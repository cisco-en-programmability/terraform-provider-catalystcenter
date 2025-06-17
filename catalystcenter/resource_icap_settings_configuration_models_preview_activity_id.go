package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceIcapSettingsConfigurationModelsPreviewActivityID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Sensors.

- Discard the ICAP configuration intent by activity ID, which was returned in TaskResponse's property "taskId" at the
beginning of the preview-approve workflow.  Discarding the intent can only be applied to intent activities that have not
been deployed.
Note that ICAP type FULL, ONBOARDING, OTA, and SPECTRUM for the scheduled-disabled task cannot be discarded or cancelled
because they have already deployed.  The feature can only be disabled by sending in a direct-deploy DELETE with API
/dna/intent/api/v1/icapSettings/deploy/deployedId/{icapDeployedId} For detailed information about the usage of the API,
please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml
`,

		CreateContext: resourceIcapSettingsConfigurationModelsPreviewActivityIDCreate,
		ReadContext:   resourceIcapSettingsConfigurationModelsPreviewActivityIDRead,
		DeleteContext: resourceIcapSettingsConfigurationModelsPreviewActivityIDDelete,
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
						"preview_activity_id": &schema.Schema{
							Description: `previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceIcapSettingsConfigurationModelsPreviewActivityIDCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vPreviewActivityID := resourceItem["preview_activity_id"]

	vvPreviewActivityID := vPreviewActivityID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sensors.DiscardsTheICapConfigurationIntentByActivityID(vvPreviewActivityID)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenSensorsDiscardsTheICapConfigurationIntentByActivityIDItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DiscardsTheICapConfigurationIntentByActivityID response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DiscardsTheICAPConfigurationIntentByActivityID", err))
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
				"Failure when executing DiscardsTheICAPConfigurationIntentByActivityID", err1))
			return diags
		}
	}

	return diags
}
func resourceIcapSettingsConfigurationModelsPreviewActivityIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIcapSettingsConfigurationModelsPreviewActivityIDDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenSensorsDiscardsTheICapConfigurationIntentByActivityIDItem(item *catalystcentersdkgo.ResponseSensorsDiscardsTheICapConfigurationIntentByActivityIDResponse) []map[string]interface{} {
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
