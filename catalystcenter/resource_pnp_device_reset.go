package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePnpDeviceReset() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Recovers a device from a Workflow Execution Error state
`,

		CreateContext: resourcePnpDeviceResetCreate,
		ReadContext:   resourcePnpDeviceResetRead,
		DeleteContext: resourcePnpDeviceResetDelete,
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

						"json_array_response": &schema.Schema{
							Description: `Json Array Response`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"json_response": &schema.Schema{
							Description: `Json Response`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status_code": &schema.Schema{
							Description: `Status Code`,
							Type:        schema.TypeFloat,
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
						"device_reset_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config_id": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"config_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"value": &schema.Schema{
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
									"device_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"license_level": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"license_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"top_of_stack_serial_number": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"project_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"workflow_id": &schema.Schema{
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

func resourcePnpDeviceResetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestPnpDeviceResetResetDevice(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.DeviceOnboardingPnp.ResetDevice(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ResetDevice", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDeviceOnboardingPnpResetDeviceItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ResetDevice response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//REVIEW: '- Analizar como se puede comprobar la ejecucion.'

}
func resourcePnpDeviceResetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpDeviceResetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestPnpDeviceResetResetDevice(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpResetDevice {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpResetDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_reset_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_reset_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_reset_list")))) {
		request.DeviceResetList = expandRequestPnpDeviceResetResetDeviceDeviceResetListArray(ctx, key+".device_reset_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	return &request
}

func expandRequestPnpDeviceResetResetDeviceDeviceResetListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetList {
	request := []catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetList{}
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
		i := expandRequestPnpDeviceResetResetDeviceDeviceResetList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceResetResetDeviceDeviceResetList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetList {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_list")))) {
		request.ConfigList = expandRequestPnpDeviceResetResetDeviceDeviceResetListConfigListArray(ctx, key+".config_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".top_of_stack_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".top_of_stack_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".top_of_stack_serial_number")))) {
		request.TopOfStackSerialNumber = interfaceToString(v)
	}
	return &request
}

func expandRequestPnpDeviceResetResetDeviceDeviceResetListConfigListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigList {
	request := []catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigList{}
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
		i := expandRequestPnpDeviceResetResetDeviceDeviceResetListConfigList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceResetResetDeviceDeviceResetListConfigList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigList {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceResetResetDeviceDeviceResetListConfigListConfigParametersArray(ctx, key+".config_parameters", d)
	}
	return &request
}

func expandRequestPnpDeviceResetResetDeviceDeviceResetListConfigListConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigListConfigParameters {
	request := []catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigListConfigParameters{}
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
		i := expandRequestPnpDeviceResetResetDeviceDeviceResetListConfigListConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceResetResetDeviceDeviceResetListConfigListConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigListConfigParameters {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpResetDeviceDeviceResetListConfigListConfigParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func flattenDeviceOnboardingPnpResetDeviceItem(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpResetDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["json_array_response"] = flattenDeviceOnboardingPnpResetDeviceItemJSONArrayResponse(item.JSONArrayResponse)
	respItem["json_response"] = flattenDeviceOnboardingPnpResetDeviceItemJSONResponse(item.JSONResponse)
	respItem["message"] = item.Message
	respItem["status_code"] = item.StatusCode
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpResetDeviceItemJSONArrayResponse(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpResetDeviceJSONArrayResponse) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenDeviceOnboardingPnpResetDeviceItemJSONResponse(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpResetDeviceJSONResponse) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
