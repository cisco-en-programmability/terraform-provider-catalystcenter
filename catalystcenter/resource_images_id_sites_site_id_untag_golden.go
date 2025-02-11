package catalystcenter

import (
	"context"

	"errors"

	"time"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceImagesIDSitesSiteIDUntagGolden() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Software Image Management (SWIM).

- Untag the golden images specifically designed for a particular device type or supervisor engine module. Conditions for
untagging the golden image:
1) Untagging the golden image can only be done where the golden tagged is applied.
  For example, if golden tagging is applied to a global site, then untagging can only be done on a global site. Even
though the same setting will be inherited on native, attempting to untag will fail.
2) Untagging of SUBPACKAGE and ROMMON image type is not supported.
`,

		CreateContext: resourceImagesIDSitesSiteIDUntagGoldenCreate,
		ReadContext:   resourceImagesIDSitesSiteIDUntagGoldenRead,
		DeleteContext: resourceImagesIDSitesSiteIDUntagGoldenDelete,
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
							Description: `id path parameter. Software image identifier is used for golden tagging or intent to tag it. The value of id can be obtained from the response of the API /dna/intent/api/v1/images?imported=true&isAddonImages=false for the base image and /dna/images/{id}/addonImages where id will be the software image identifier of the base image.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site identifier for tagged image or intent to tag it. The default value is global site id. See [https://developer.cisco.com/docs/dna-center](#!get-site) for siteId
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
							Description: `The product name ordinal is a unique value for each network device product. productNameOrdinal can be obtained from the response of API /dna/intent/api/v1/siteWiseProductNames?siteId=<siteId>
`,
							Type:     schema.TypeFloat,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"supervisor_product_name_ordinal": &schema.Schema{
							Description: `The supervisor engine module ordinal is a unique value for each supervisor module. supervisorProductNameOrdinal can be obtained from the response of API /dna/intent/api/v1/siteWiseProductNames?siteId=<siteId>
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

func resourceImagesIDSitesSiteIDUntagGoldenCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vSiteID := resourceItem["site_id"]

	vvID := vID.(string)
	vvSiteID := vSiteID.(string)
	request1 := expandRequestImagesIDSitesSiteIDUntagGoldenUntaggingGoldenImageV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.SoftwareImageManagementSwim.UntaggingGoldenImageV1(vvID, vvSiteID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vItem1 := flattenSoftwareImageManagementSwimUntaggingGoldenImageV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UntaggingGoldenImageV1 response",
			err))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UntaggingGoldenImageV1", err))
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
				"Failure when executing UntaggingGoldenImageV1", err1))
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
func resourceImagesIDSitesSiteIDUntagGoldenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceImagesIDSitesSiteIDUntagGoldenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestImagesIDSitesSiteIDUntagGoldenUntaggingGoldenImageV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSoftwareImageManagementSwimUntaggingGoldenImageV1 {
	request := catalystcentersdkgo.RequestSoftwareImageManagementSwimUntaggingGoldenImageV1{}
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

func flattenSoftwareImageManagementSwimUntaggingGoldenImageV1Item(item *catalystcentersdkgo.ResponseSoftwareImageManagementSwimUntaggingGoldenImageV1Response) []map[string]interface{} {
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
