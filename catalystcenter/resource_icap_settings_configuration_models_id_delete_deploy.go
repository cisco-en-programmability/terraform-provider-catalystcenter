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
func resourceIcapSettingsConfigurationModelsIDDeleteDeploy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sensors.

- Creates a ICAP configuration intent to remove the ICAP RFSTATS or ANOMALY configuration from the device. The task has
not been applied to the device yet. Subsequent preview-approve workflow APIs must be used to complete the preview-
approve process.  The path parameter 'id' can be retrieved from **GET /dna/intent/api/v1/icapSettings** API. For
detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
ICAP_APIs-1.0.0-resolved.yaml
`,

		CreateContext: resourceIcapSettingsConfigurationModelsIDDeleteDeployCreate,
		ReadContext:   resourceIcapSettingsConfigurationModelsIDDeleteDeployRead,
		DeleteContext: resourceIcapSettingsConfigurationModelsIDDeleteDeployDelete,
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
						"id": &schema.Schema{
							Description: `id path parameter. A unique ID of the deployed ICAP object, which can be obtained from **GET /dna/intent/api/v1/icapSettings**
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"object": &schema.Schema{
							Description: `object`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceIcapSettingsConfigurationModelsIDDeleteDeployCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)
	request1 := expandRequestIcapSettingsConfigurationModelsIDDeleteDeployCreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Sensors.CreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice(vvID, request1)

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
			"Failure when executing CreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDevice", err))
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
				"Failure when executing CreatesAICAPConfigurationWorkflowForICAPIntentToRemoveTheICAPConfigurationOnTheDevice", err1))
			return diags
		}
	}

	vItem1 := flattenSensorsCreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceIcapSettingsConfigurationModelsIDDeleteDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIcapSettingsConfigurationModelsIDDeleteDeployDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIcapSettingsConfigurationModelsIDDeleteDeployCreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice {
	var request interface{}
	if v, ok := d.GetOkExists(fixKeyAccess(key)); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key)))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key)))) {
		request = v
	}
	return request.(*catalystcentersdkgo.RequestSensorsCreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDevice)
}

func flattenSensorsCreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceItem(item *catalystcentersdkgo.ResponseSensorsCreatesAiCapConfigurationWorkflowForICapintentToRemoveTheICapConfigurationOnTheDeviceResponse) []map[string]interface{} {
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
