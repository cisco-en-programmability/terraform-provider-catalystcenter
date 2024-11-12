package catalystcenter

import (
	"context"

	"errors"

	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkProfilesForSitesSiteAssignmentsBulkCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Assign a network profile for sites to a list of sites. Also assigns the profile to child sites.
`,

		CreateContext: resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateCreate,
		ReadContext:   resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateRead,
		DeleteContext: resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateDelete,
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
							Description: `Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5 
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
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
						"profile_id": &schema.Schema{
							Description: `profileId path parameter. The *id* of the network profile, retrievable from *GET /intent/api/v1/networkProfilesForSites*
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"type": &schema.Schema{
							Type:     schema.TypeString, //TEST,
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

func resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vProfileID := resourceItem["profile_id"]

	vvProfileID := vProfileID.(string)
	request1 := expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSitesV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.SiteDesign.AssignANetworkProfileForSitesToAListOfSitesV1(vvProfileID, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing AssignANetworkProfileForSitesToAListOfSitesV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignANetworkProfileForSitesToAListOfSitesV1", err))
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
				"Failure when executing AssignANetworkProfileForSitesToAListOfSitesV1", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignANetworkProfileForSitesToAListOfSitesV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkProfilesForSitesSiteAssignmentsBulkCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSitesV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1 {
	request := catalystcentersdkgo.RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1{}
	request.Type = expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSitesV1Type(ctx, key, d)
	return &request
}

func expandRequestNetworkProfilesForSitesSiteAssignmentsBulkCreateAssignANetworkProfileForSitesToAListOfSitesV1Type(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Type {
	var request catalystcentersdkgo.RequestSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Type
	request = d.Get(fixKeyAccess(key))
	return &request
}

func flattenSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Item(item *catalystcentersdkgo.ResponseSiteDesignAssignANetworkProfileForSitesToAListOfSitesV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	respItem["task_id"] = item.TaskID
	return []map[string]interface{}{
		respItem,
	}
}
