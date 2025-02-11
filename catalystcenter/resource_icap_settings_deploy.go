package catalystcenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceIcapSettingsDeploy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sensors.

- Deploys the given ICAP intent without preview and approval. The response body contains a task object with a taskId and
a URL for more information about the task. The deployment status of this ICAP intent can be found in the output of the
URL.  For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
ICAP_APIs-1.0.0-resolved.yaml
`,

		CreateContext: resourceIcapSettingsDeployCreate,
		ReadContext:   resourceIcapSettingsDeployRead,
		DeleteContext: resourceIcapSettingsDeployDelete,
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
							Description: `Array of RequestSensorsDeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1`,
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

func resourceIcapSettingsDeployCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	resourceItem := *getResourceItem(d.Get("parameters"))
	vPreviewDescription := resourceItem["preview_description"]
	vvPreviewDescription := interfaceToString(vPreviewDescription)

	request1 := expandRequestIcapSettingsDeployDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1(ctx, "parameters.0", d)
	queryParams1 := catalystcentersdkgo.DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1QueryParams{}
	queryParams1.PreviewDescription = vvPreviewDescription
	// has_unknown_response: None

	response1, restyResp1, err := client.Sensors.DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1(request1, &queryParams1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vItem1 := flattenSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 response",
			err))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1", err))
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
				"Failure when executing DeploysTheGivenICAPConfigurationIntentWithoutPreviewAndApproveV1", err1))
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
func resourceIcapSettingsDeployRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIcapSettingsDeployDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIcapSettingsDeployDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 {
	request := catalystcentersdkgo.RequestSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1{}
	if v := expandRequestIcapSettingsDeployDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestIcapSettingsDeployDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 {
	request := []catalystcentersdkgo.RequestItemSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1{}
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
		i := expandRequestIcapSettingsDeployDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestIcapSettingsDeployDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1 {
	request := catalystcentersdkgo.RequestItemSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1{}
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
		request.Slot = responseInterfaceToSliceFloat64(v)
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

func flattenSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1Item(item *catalystcentersdkgo.ResponseSensorsDeploysTheGivenICapConfigurationIntentWithoutPreviewAndApproveV1Response) []map[string]interface{} {
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
