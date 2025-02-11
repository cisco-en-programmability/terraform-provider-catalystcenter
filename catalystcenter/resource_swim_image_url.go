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
func resourceSwimImageURL() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Fetches a software image from remote file system (using URL for HTTP/FTP) and uploads to DNA Center. Supported image
files extensions are bin, img, tar, smu, pie, aes, iso, ova, tar_gz and qcow2
`,

		CreateContext: resourceSwimImageURLCreate,
		ReadContext:   resourceSwimImageURLRead,
		DeleteContext: resourceSwimImageURLDelete,
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
						"schedule_at": &schema.Schema{
							Description: `scheduleAt query parameter. Epoch Time (The number of milli-seconds since January 1 1970 UTC) at which the distribution should be scheduled (Optional)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"schedule_desc": &schema.Schema{
							Description: `scheduleDesc query parameter. Custom Description (Optional)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"schedule_origin": &schema.Schema{
							Description: `scheduleOrigin query parameter. Originator of this call (Optional)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestSoftwareImageManagementSwimImportSoftwareImageViaURLV1`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"application_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"image_family": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"source_url": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"third_party": &schema.Schema{
										// Type:     schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"vendor": &schema.Schema{
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

func resourceSwimImageURLCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSwimImageURLImportSoftwareImageViaURLV1(ctx, "parameters.0", d)
	queryParams1 := catalystcentersdkgo.ImportSoftwareImageViaURLV1QueryParams{}

	// has_unknown_response: None

	response1, restyResp1, err := client.SoftwareImageManagementSwim.ImportSoftwareImageViaURLV1(request1, &queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ImportSoftwareImageViaURLV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ImportSoftwareImageViaURLV1", err))
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
				"Failure when executing ImportSoftwareImageViaURLV1", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenSoftwareImageManagementSwimImportSoftwareImageViaURLV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportSoftwareImageViaURLV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceSwimImageURLRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSwimImageURLDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSwimImageURLImportSoftwareImageViaURLV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSoftwareImageManagementSwimImportSoftwareImageViaURLV1 {
	request := catalystcentersdkgo.RequestSoftwareImageManagementSwimImportSoftwareImageViaURLV1{}
	if v := expandRequestSwimImageURLImportSoftwareImageViaURLV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestSwimImageURLImportSoftwareImageViaURLV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURLV1 {
	request := []catalystcentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURLV1{}
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
		i := expandRequestSwimImageURLImportSoftwareImageViaURLV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSwimImageURLImportSoftwareImageViaURLV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURLV1 {
	request := catalystcentersdkgo.RequestItemSoftwareImageManagementSwimImportSoftwareImageViaURLV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_type")))) {
		request.ApplicationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_family")))) {
		request.ImageFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".source_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".source_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".source_url")))) {
		request.SourceURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".third_party")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".third_party")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".third_party")))) {
		request.ThirdParty = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vendor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vendor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vendor")))) {
		request.Vendor = interfaceToString(v)
	}
	return &request
}

func flattenSoftwareImageManagementSwimImportSoftwareImageViaURLV1Item(item *catalystcentersdkgo.ResponseSoftwareImageManagementSwimImportSoftwareImageViaURLV1Response) []map[string]interface{} {
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
