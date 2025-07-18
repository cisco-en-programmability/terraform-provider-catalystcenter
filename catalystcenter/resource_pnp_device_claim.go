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
func resourcePnpDeviceClaim() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Claims one of more devices with specified workflow
`,

		CreateContext: resourcePnpDeviceClaimCreate,
		ReadContext:   resourcePnpDeviceClaimRead,
		DeleteContext: resourcePnpDeviceClaimDelete,
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
						"authorization_needed": &schema.Schema{
							Description: `Flag to enable/disable PnP device authorization. (true means enable)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"config_file_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"config_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"device_claim_list": &schema.Schema{
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
						"file_service_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"image_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"image_url": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"populate_inventory": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
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

func resourcePnpDeviceClaimCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestPnpDeviceClaimClaimDevice(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.DeviceOnboardingPnp.ClaimDevice(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ClaimDevice", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDeviceOnboardingPnpClaimDeviceItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ClaimDevice response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//REVIEW: '- Analizar como se puede comprobar la ejecucion.'

}
func resourcePnpDeviceClaimRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpDeviceClaimDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestPnpDeviceClaimClaimDevice(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDevice {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_file_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_file_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_file_url")))) {
		request.ConfigFileURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_claim_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_claim_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_claim_list")))) {
		request.DeviceClaimList = expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListArray(ctx, key+".device_claim_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".file_service_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".file_service_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".file_service_id")))) {
		request.FileServiceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_url")))) {
		request.ImageURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".populate_inventory")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".populate_inventory")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".populate_inventory")))) {
		request.PopulateInventory = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authorization_needed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authorization_needed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authorization_needed")))) {
		request.AuthorizationNeeded = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList {
	request := []catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList{}
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
		i := expandRequestPnpDeviceClaimClaimDeviceDeviceClaimList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_list")))) {
		request.ConfigList = expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListArray(ctx, key+".config_list", d)
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

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList {
	request := []catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList{}
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
		i := expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListConfigParametersArray(ctx, key+".config_parameters", d)
	}
	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters {
	request := []catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters{}
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
		i := expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceClaimClaimDeviceDeviceClaimListConfigListConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpClaimDeviceDeviceClaimListConfigListConfigParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func flattenDeviceOnboardingPnpClaimDeviceItem(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpClaimDevice) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["json_array_response"] = flattenDeviceOnboardingPnpClaimDeviceItemJSONArrayResponse(item.JSONArrayResponse)
	respItem["json_response"] = flattenDeviceOnboardingPnpClaimDeviceItemJSONResponse(item.JSONResponse)
	respItem["message"] = item.Message
	respItem["status_code"] = item.StatusCode
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpClaimDeviceItemJSONArrayResponse(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpClaimDeviceJSONArrayResponse) []interface{} {
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

func flattenDeviceOnboardingPnpClaimDeviceItemJSONResponse(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpClaimDeviceJSONResponse) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
