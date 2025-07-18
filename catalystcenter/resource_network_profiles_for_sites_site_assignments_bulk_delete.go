package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkProfilesForSitesSiteAssignmentsBulkDelete() *schema.Resource {
	return &schema.Resource{
		Description: `It performs delete operation on Site Design.

- Unassigns a given network profile for sites from multiple sites. The profile must be removed from the containing
building first if this site is a floor.
`,

		CreateContext: resourceNetworkProfilesForSitesSiteAssignmentsBulkDeleteCreate,
		ReadContext:   resourceNetworkProfilesForSitesSiteAssignmentsBulkDeleteRead,
		DeleteContext: resourceNetworkProfilesForSitesSiteAssignmentsBulkDeleteDelete,
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
							Description: `profileId path parameter. The **id** of the network profile, retrievable from **GET /intent/api/v1/networkProfilesForSites**
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"site_id": &schema.Schema{
							Description: `siteId query parameter. The id or ids of the network profile, retrievable from /dna/intent/api/v1/sites.. A list of profile ids can be passed as a queryParameter in two ways:  1. a comma-separated string ( siteId=388a23e9-4739-4be7-a0aa-cc5a95d158dd,2726dc60-3a12-451e-947a-d972ebf58743), or... 2. as separate query parameters with the same name ( siteId=388a23e9-4739-4be7-a0aa-cc5a95d158dd&siteId=2726dc60-3a12-451e-947a-d972ebf58743
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkProfilesForSitesSiteAssignmentsBulkDeleteCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vProfileID := resourceItem["profile_id"]

	vSiteID := resourceItem["site_id"]

	vvProfileID := vProfileID.(string)
	queryParams1 := catalystcentersdkgo.UnassignsANetworkProfileForSitesFromMultipleSitesQueryParams{}

	queryParams1.SiteID = vSiteID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.SiteDesign.UnassignsANetworkProfileForSitesFromMultipleSites(vvProfileID, &queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing UnassignsANetworkProfileForSitesFromMultipleSites", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UnassignsANetworkProfileForSitesFromMultipleSites", err))
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
				"Failure when executing UnassignsANetworkProfileForSitesFromMultipleSites", err1))
			return diags
		}
	}
	vItem1 := flattenSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UnassignsANetworkProfileForSitesFromMultipleSites response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceNetworkProfilesForSitesSiteAssignmentsBulkDeleteRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkProfilesForSitesSiteAssignmentsBulkDeleteDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesItem(item *catalystcentersdkgo.ResponseSiteDesignUnassignsANetworkProfileForSitesFromMultipleSitesResponse) []map[string]interface{} {
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
