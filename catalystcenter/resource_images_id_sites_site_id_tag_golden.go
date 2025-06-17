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
func resourceImagesIDSitesSiteIDTagGolden() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Creates golden image tagging specifically for a particular device type or supervisor engine module. Conditions for
tagging the golden image:
1) The golden tagging of SMU, PISRT_SMU, APSP, and APDP image type depends on the golden tagged applied on the base
image. If any discrepancies are identified in the request payload, the golden tagging process will fail. For example:


    a) If the base image is tagged with Device Role: ACCESS, then add-ons can only be done ACCESS role only and the same
is applied if any device tag is used. Any other request will fail.

    b) If the base image is tagged at global or any site level then add-ons also need to be tagged at site level.

2) Tagging of SUBPACKAGE and ROMMON image type is not supported.
`,

		CreateContext: resourceImagesIDSitesSiteIDTagGoldenCreate,
		ReadContext:   resourceImagesIDSitesSiteIDTagGoldenRead,
		DeleteContext: resourceImagesIDSitesSiteIDTagGoldenDelete,
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
							Description: `The UUID of the task
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `The path to the API endpoint to GET for information on the task
`,
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
						"id": &schema.Schema{
							Description: `id path parameter. Software image identifier is used for golden tagging or intent to tag it. The value of **id** can be obtained from the response of the API **/dna/intent/api/v1/images?imported=true&isAddonImages=false** for the base image and **/dna/images/{id}/addonImages** where **id** will be the software image identifier of the base image.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site identifier for tagged image or intent to tag it. The default value is global site id. See [https://developer.cisco.com/docs/dna-center](#!get-site) for **siteId**
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"device_roles": &schema.Schema{
							Description: `Device Roles. Available value will be [ CORE, DISTRIBUTION, UNKNOWN, ACCESS, BORDER ROUTER ]
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"device_tags": &schema.Schema{
							Description: `Device tags can be fetched fom API https://developer.cisco.com/docs/dna-center/#!get-tag
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"product_name_ordinal": &schema.Schema{
							Description: `The product name ordinal is a unique value for each network device product. **productNameOrdinal** can be obtained from the response of API **/dna/intent/api/v1/siteWiseProductNames?siteId=<siteId>**
`,
							Type:     schema.TypeFloat,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"supervisor_product_name_ordinal": &schema.Schema{
							Description: `The supervisor engine module ordinal is a unique value for each supervisor module. **supervisorProductNameOrdinal** can be obtained from the response of API **/dna/intent/api/v1/siteWiseProductNames?siteId=<siteId>**
`,
							Type:     schema.TypeFloat,
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

func resourceImagesIDSitesSiteIDTagGoldenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]
	vSiteID := resourceItem["site_id"]

	vvID := vID.(string)
	vvSiteID := vSiteID.(string)
	request1 := expandRequestImagesIDSitesSiteIDTagGoldenTaggingGoldenImage(ctx, "parameters.0", d)

	response1, restyResp1, err := client.SoftwareImageManagementSwim.TaggingGoldenImage(vvID, vvSiteID, request1)

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
			"Failure when executing TaggingGoldenImage", err))
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
				"Failure when executing TaggingGoldenImage", err1))
			return diags
		}
	}

	vItem1 := flattenSoftwareImageManagementSwimTaggingGoldenImageItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting TaggingGoldenImage response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceImagesIDSitesSiteIDTagGoldenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceImagesIDSitesSiteIDTagGoldenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestImagesIDSitesSiteIDTagGoldenTaggingGoldenImage(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSoftwareImageManagementSwimTaggingGoldenImage {
	request := catalystcentersdkgo.RequestSoftwareImageManagementSwimTaggingGoldenImage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_name_ordinal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_name_ordinal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_name_ordinal")))) {
		request.ProductNameOrdinal = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".supervisor_product_name_ordinal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".supervisor_product_name_ordinal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".supervisor_product_name_ordinal")))) {
		request.SupervisorProductNameOrdinal = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_roles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_roles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_roles")))) {
		request.DeviceRoles = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_tags")))) {
		request.DeviceTags = interfaceToSliceString(v)
	}
	return &request
}

func flattenSoftwareImageManagementSwimTaggingGoldenImageItem(item *catalystcentersdkgo.ResponseSoftwareImageManagementSwimTaggingGoldenImageResponse) []map[string]interface{} {
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
