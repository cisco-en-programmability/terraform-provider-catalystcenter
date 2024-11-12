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
func resourceImageDistribution() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Distributes a software image on a given device. Software image must be imported successfully into DNA Center before it
can be distributed
`,

		CreateContext: resourceImageDistributionCreate,
		ReadContext:   resourceImageDistributionRead,
		DeleteContext: resourceImageDistributionDelete,
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
						"payload": &schema.Schema{
							Description: `Array of RequestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"image_uuid": &schema.Schema{
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
		},
	}
}

func resourceImageDistributionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestImageDistributionTriggerSoftwareImageDistributionV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.SoftwareImageManagementSwim.TriggerSoftwareImageDistributionV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing TriggerSoftwareImageDistributionV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing TriggerSoftwareImageDistributionV1", err))
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
				"Failure when executing TriggerSoftwareImageDistributionV1", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TriggerSoftwareImageDistributionV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceImageDistributionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceImageDistributionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestImageDistributionTriggerSoftwareImageDistributionV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 {
	request := catalystcentersdkgo.RequestSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1{}
	if v := expandRequestImageDistributionTriggerSoftwareImageDistributionV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestImageDistributionTriggerSoftwareImageDistributionV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 {
	request := []catalystcentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1{}
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
		i := expandRequestImageDistributionTriggerSoftwareImageDistributionV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestImageDistributionTriggerSoftwareImageDistributionV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1 {
	request := catalystcentersdkgo.RequestItemSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuid")))) {
		request.DeviceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_uuid")))) {
		request.ImageUUID = interfaceToString(v)
	}
	return &request
}

func flattenSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1Item(item *catalystcentersdkgo.ResponseSoftwareImageManagementSwimTriggerSoftwareImageDistributionV1Response) []map[string]interface{} {
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
