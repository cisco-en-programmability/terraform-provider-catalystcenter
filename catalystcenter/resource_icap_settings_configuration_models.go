package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceIcapSettingsConfigurationModels() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sensors.

- This creates an ICAP configuration intent for preview approval. The intent is not deployed to the device until further
preview-approve APIs are applied. This data source action is the first step in the preview-approve workflow, which
consists of several APIs. Skipping any API in the process is not recommended for a complete preview-approve use case.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
ICAP_APIs-1.0.0-resolved.yaml
`,

		CreateContext: resourceIcapSettingsConfigurationModelsCreate,
		ReadContext:   resourceIcapSettingsConfigurationModelsRead,
		DeleteContext: resourceIcapSettingsConfigurationModelsDelete,
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
						"preview_description": &schema.Schema{
							Description: `previewDescription query parameter. The ICAP intent's preview-deploy description string
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestSensorsCreatesAnICAPConfigurationIntentForPreviewApprove`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"apid": &schema.Schema{
										Description: `Ap Id`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"capture_type": &schema.Schema{
										Description: `Capture Type`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"client_mac": &schema.Schema{
										Description: `Client Mac`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"duration_in_mins": &schema.Schema{
										Description: `Duration In Mins`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"ota_band": &schema.Schema{
										Description: `Ota Band`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"ota_channel": &schema.Schema{
										Description: `Ota Channel`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"ota_channel_width": &schema.Schema{
										Description: `Ota Channel Width`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"slot": &schema.Schema{
										Description: `Slot`,
										Type:        schema.TypeList,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeFloat,
										},
									},
									"wlc_id": &schema.Schema{
										Description: `Wlc Id`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
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

func resourceIcapSettingsConfigurationModelsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestIcapSettingsConfigurationModelsCreatesAnICapConfigurationIntentForPreviewApprove(ctx, "parameters.0", d)
	queryParams1 := catalystcentersdkgo.CreatesAnICapConfigurationIntentForPreviewApproveQueryParams{}

	response1, restyResp1, err := client.Sensors.CreatesAnICapConfigurationIntentForPreviewApprove(request1, &queryParams1)

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
			"Failure when executing CreatesAnICAPConfigurationIntentForPreviewApprove", err))
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
				"Failure when executing CreatesAnICAPConfigurationIntentForPreviewApprove", err1))
			return diags
		}
	}

	vItem1 := flattenSensorsCreatesAnICapConfigurationIntentForPreviewApproveItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreatesAnICapConfigurationIntentForPreviewApprove response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceIcapSettingsConfigurationModelsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIcapSettingsConfigurationModelsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIcapSettingsConfigurationModelsCreatesAnICapConfigurationIntentForPreviewApprove(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreatesAnICapConfigurationIntentForPreviewApprove {
	request := catalystcentersdkgo.RequestSensorsCreatesAnICapConfigurationIntentForPreviewApprove{}
	if v := expandRequestIcapSettingsConfigurationModelsCreatesAnICapConfigurationIntentForPreviewApproveItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestIcapSettingsConfigurationModelsCreatesAnICapConfigurationIntentForPreviewApproveItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSensorsCreatesAnICapConfigurationIntentForPreviewApprove {
	request := []catalystcentersdkgo.RequestItemSensorsCreatesAnICapConfigurationIntentForPreviewApprove{}
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
		i := expandRequestIcapSettingsConfigurationModelsCreatesAnICapConfigurationIntentForPreviewApproveItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestIcapSettingsConfigurationModelsCreatesAnICapConfigurationIntentForPreviewApproveItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSensorsCreatesAnICapConfigurationIntentForPreviewApprove {
	request := catalystcentersdkgo.RequestItemSensorsCreatesAnICapConfigurationIntentForPreviewApprove{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".capture_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".capture_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".capture_type")))) {
		request.CaptureType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".duration_in_mins")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".duration_in_mins")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".duration_in_mins")))) {
		request.DurationInMins = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".client_mac")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".client_mac")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".client_mac")))) {
		request.ClientMac = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlc_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlc_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlc_id")))) {
		request.WlcID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".apid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".apid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".apid")))) {
		request.APID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".slot")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".slot")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".slot")))) {
		if arr, ok := v.([]interface{}); ok {
			slots := make([]float64, len(arr))
			for i, val := range arr {
				slots[i] = interfaceToFloat64(val)
			}
			request.Slot = &slots
		}
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ota_band")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ota_band")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ota_band")))) {
		request.OtaBand = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ota_channel")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ota_channel")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ota_channel")))) {
		request.OtaChannel = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ota_channel_width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ota_channel_width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ota_channel_width")))) {
		request.OtaChannelWidth = interfaceToIntPtr(v)
	}
	return &request
}

func flattenSensorsCreatesAnICapConfigurationIntentForPreviewApproveItem(item *catalystcentersdkgo.ResponseSensorsCreatesAnICapConfigurationIntentForPreviewApproveResponse) []map[string]interface{} {
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
