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
func resourceGlobalCredentialUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Discovery.

- Update global credential for network devices in site(s)
`,

		CreateContext: resourceGlobalCredentialUpdateCreate,
		ReadContext:   resourceGlobalCredentialUpdateRead,
		DeleteContext: resourceGlobalCredentialUpdateDelete,
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
						"global_credential_id": &schema.Schema{
							Description: `globalCredentialId path parameter. Global credential Uuid
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"site_uuids": &schema.Schema{
							Description: `List of siteUuids where credential is to be updated
`,
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
	}
}

func resourceGlobalCredentialUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vGlobalCredentialID := resourceItem["global_credential_id"]

	vvGlobalCredentialID := vGlobalCredentialID.(string)
	request1 := expandRequestGlobalCredentialUpdateUpdateGlobalCredentialsV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Discovery.UpdateGlobalCredentialsV1(vvGlobalCredentialID, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing UpdateGlobalCredentialsV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UpdateGlobalCredentialsV1", err))
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
				"Failure when executing UpdateGlobalCredentialsV1", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenDiscoveryUpdateGlobalCredentialsV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateGlobalCredentialsV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceGlobalCredentialUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceGlobalCredentialUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestGlobalCredentialUpdateUpdateGlobalCredentialsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV1 {
	request := catalystcentersdkgo.RequestDiscoveryUpdateGlobalCredentialsV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_uuids")))) {
		request.SiteUUIDs = interfaceToSliceString(v)
	}
	return &request
}

func flattenDiscoveryUpdateGlobalCredentialsV1Item(item *catalystcentersdkgo.ResponseDiscoveryUpdateGlobalCredentialsV1Response) []map[string]interface{} {
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
