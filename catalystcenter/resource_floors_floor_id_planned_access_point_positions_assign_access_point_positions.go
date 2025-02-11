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
func resourceFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Assign Planned Access Points to operations ones.
`,

		CreateContext: resourceFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsCreate,
		ReadContext:   resourceFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsRead,
		DeleteContext: resourceFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsDelete,
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
						"floor_id": &schema.Schema{
							Description: `floorId path parameter. Floor Id
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"access_point_id": &schema.Schema{
										Description: `Operational Access Point Id
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"planned_access_point_id": &schema.Schema{
										Description: `Planned Access Point Id
`,
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

func resourceFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vFloorID := resourceItem["floor_id"]

	vvFloorID := vFloorID.(string)
	request1 := expandRequestFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsAssignPlannedAccessPointsToOperationsOnesV2(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.SiteDesign.AssignPlannedAccessPointsToOperationsOnesV2(vvFloorID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vItem1 := flattenSiteDesignAssignPlannedAccessPointsToOperationsOnesV2Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignPlannedAccessPointsToOperationsOnesV2 response",
			err))
		return diags
	}
	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignPlannedAccessPointsToOperationsOnesV2", err))
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
				"Failure when executing AssignPlannedAccessPointsToOperationsOnesV2", err1))
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
func resourceFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsAssignPlannedAccessPointsToOperationsOnesV2(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2 {
	request := catalystcentersdkgo.RequestSiteDesignAssignPlannedAccessPointsToOperationsOnesV2{}
	if v := expandRequestFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsAssignPlannedAccessPointsToOperationsOnesV2ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsAssignPlannedAccessPointsToOperationsOnesV2ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSiteDesignAssignPlannedAccessPointsToOperationsOnesV2 {
	request := []catalystcentersdkgo.RequestItemSiteDesignAssignPlannedAccessPointsToOperationsOnesV2{}
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
		i := expandRequestFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsAssignPlannedAccessPointsToOperationsOnesV2Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestFloorsFloorIDPlannedAccessPointPositionsAssignAccessPointPositionsAssignPlannedAccessPointsToOperationsOnesV2Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSiteDesignAssignPlannedAccessPointsToOperationsOnesV2 {
	request := catalystcentersdkgo.RequestItemSiteDesignAssignPlannedAccessPointsToOperationsOnesV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".planned_access_point_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".planned_access_point_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".planned_access_point_id")))) {
		request.PlannedAccessPointID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".access_point_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".access_point_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".access_point_id")))) {
		request.AccessPointID = interfaceToString(v)
	}
	return &request
}

func flattenSiteDesignAssignPlannedAccessPointsToOperationsOnesV2Item(item *catalystcentersdkgo.ResponseSiteDesignAssignPlannedAccessPointsToOperationsOnesV2Response) []map[string]interface{} {
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
