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
func resourceDeployTemplateV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Configuration Templates.

- V2 API to deploy a template.
`,

		CreateContext: resourceDeployTemplateV2Create,
		ReadContext:   resourceDeployTemplateV2Read,
		DeleteContext: resourceDeployTemplateV2Delete,
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
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
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
						"force_push_template": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"is_composite": &schema.Schema{
							Description: `Composite template flag
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"main_template_id": &schema.Schema{
							Description: `Main template UUID of versioned template
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"member_template_deployment_info": &schema.Schema{
							Description: `memberTemplateDeploymentInfo
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"target_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"host_name": &schema.Schema{
										Description: `Hostname of device is required if targetType is MANAGED_DEVICE_HOSTNAME
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of target is required if targetType is MANAGED_DEVICE_UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"params": &schema.Schema{
										Description: `Template params/values to be provisioned
`,
										Type:     schema.TypeString, //TEST,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"resource_params": &schema.Schema{
										Description: `Resource params to be provisioned. Refer to features page for usage details
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"type": &schema.Schema{
										Description: `Target type of device
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"versioned_template_id": &schema.Schema{
										Description: `Versioned templateUUID to be provisioned
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"template_id": &schema.Schema{
							Description: `UUID of template to be provisioned
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

func resourceDeployTemplateV2Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestDeployTemplateV2DeployTemplateV2(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.ConfigurationTemplates.DeployTemplateV2(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing DeployTemplateV2", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeployTemplateV2", err))
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
				"Failure when executing DeployTemplateV2", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenConfigurationTemplatesDeployTemplateV2Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting DeployTemplateV2 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceDeployTemplateV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceDeployTemplateV2Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestDeployTemplateV2DeployTemplateV2(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2 {
	request := catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".force_push_template")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".force_push_template")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".force_push_template")))) {
		request.ForcePushTemplate = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_composite")))) {
		request.IsComposite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".main_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".main_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".main_template_id")))) {
		request.MainTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".member_template_deployment_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".member_template_deployment_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".member_template_deployment_info")))) {
		request.MemberTemplateDeploymentInfo = expandRequestDeployTemplateV2DeployTemplateV2MemberTemplateDeploymentInfoArray(ctx, key+".member_template_deployment_info", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".target_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".target_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".target_info")))) {
		request.TargetInfo = expandRequestDeployTemplateV2DeployTemplateV2TargetInfoArray(ctx, key+".target_info", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	return &request
}

func expandRequestDeployTemplateV2DeployTemplateV2MemberTemplateDeploymentInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo{}
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
		i := expandRequestDeployTemplateV2DeployTemplateV2MemberTemplateDeploymentInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeployTemplateV2DeployTemplateV2MemberTemplateDeploymentInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo {
	var request catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2MemberTemplateDeploymentInfo
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestDeployTemplateV2DeployTemplateV2TargetInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo{}
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
		i := expandRequestDeployTemplateV2DeployTemplateV2TargetInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDeployTemplateV2DeployTemplateV2TargetInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo {
	request := catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".params")))) {
		request.Params = expandRequestDeployTemplateV2DeployTemplateV2TargetInfoParams(ctx, key+".params.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_params")))) {
		request.ResourceParams = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".versioned_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".versioned_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".versioned_template_id")))) {
		request.VersionedTemplateID = interfaceToString(v)
	}
	return &request
}

func expandRequestDeployTemplateV2DeployTemplateV2TargetInfoParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams {
	var request catalystcentersdkgo.RequestConfigurationTemplatesDeployTemplateV2TargetInfoParams
	request = d.Get(fixKeyAccess(key)).(map[string]interface{})
	return &request
}

func flattenConfigurationTemplatesDeployTemplateV2Item(item *catalystcentersdkgo.ResponseConfigurationTemplatesDeployTemplateV2Response) []map[string]interface{} {
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
