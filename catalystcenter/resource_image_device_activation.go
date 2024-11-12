package catalystcenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceImageDeviceActivation() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Activates a software image on a given device. Software image must be present in the device flash
`,

		CreateContext: resourceImageDeviceActivationCreate,
		ReadContext:   resourceImageDeviceActivationRead,
		DeleteContext: resourceImageDeviceActivationDelete,
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
						"client_type": &schema.Schema{
							Description: `Client-Type header parameter. Client-type (Optional)
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"client_url": &schema.Schema{
							Description: `Client-Url header parameter. Client-url (Optional)
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"schedule_validate": &schema.Schema{
							Description: `scheduleValidate query parameter. scheduleValidate, validates data before schedule (Optional)
`,
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"activate_lower_image_version": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"device_upgrade_mode": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"device_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"distribute_if_needed": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"image_uuid_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"smu_image_uuid_list": &schema.Schema{
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
				},
			},
		},
	}
}

func resourceImageDeviceActivationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vClientType := resourceItem["client_type"]

	vClientURL := resourceItem["client_url"]

	request1 := expandRequestImageDeviceActivationTriggerSoftwareImageActivationV1(ctx, "parameters.0", d)

	headerParams1 := catalystcentersdkgo.TriggerSoftwareImageActivationV1HeaderParams{}
	queryParams1 := catalystcentersdkgo.TriggerSoftwareImageActivationV1QueryParams{}

	headerParams1.ClientType = vClientType.(string)

	headerParams1.ClientURL = vClientURL.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.SoftwareImageManagementSwim.TriggerSoftwareImageActivationV1(request1, &headerParams1, &queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing TriggerSoftwareImageActivationV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing TriggerSoftwareImageActivationV1", err))
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
				"Failure when executing TriggerSoftwareImageActivationV1", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenSoftwareImageManagementSwimTriggerSoftwareImageActivationV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TriggerSoftwareImageActivationV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceImageDeviceActivationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceImageDeviceActivationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestImageDeviceActivationTriggerSoftwareImageActivationV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 {
	request := catalystcentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageActivationV1{}
	if v := expandRequestImageDeviceActivationTriggerSoftwareImageActivationV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestImageDeviceActivationTriggerSoftwareImageActivationV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 {
	request := []catalystcentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivationV1{}
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
		i := expandRequestImageDeviceActivationTriggerSoftwareImageActivationV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestImageDeviceActivationTriggerSoftwareImageActivationV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivationV1 {
	request := catalystcentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageActivationV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".activate_lower_image_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".activate_lower_image_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".activate_lower_image_version")))) {
		request.ActivateLowerImageVersion = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_upgrade_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_upgrade_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_upgrade_mode")))) {
		request.DeviceUpgradeMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuid")))) {
		request.DeviceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".distribute_if_needed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".distribute_if_needed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".distribute_if_needed")))) {
		request.DistributeIfNeeded = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_uuid_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_uuid_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_uuid_list")))) {
		request.ImageUUIDList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smu_image_uuid_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smu_image_uuid_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smu_image_uuid_list")))) {
		request.SmuImageUUIDList = interfaceToSliceString(v)
	}
	return &request
}

func flattenSoftwareImageManagementSwimTriggerSoftwareImageActivationV1Item(item *catalystcentersdkgo.ResponseSoftwareImageManagementSwimTriggerSoftwareImageActivationV1Response) []map[string]interface{} {
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
