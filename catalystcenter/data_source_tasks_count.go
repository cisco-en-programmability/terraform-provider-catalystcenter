package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTasksCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- Returns the number of tasks that meet the filter criteria
`,

		ReadContext: dataSourceTasksCountRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. This is the epoch millisecond end time upto which task records need to be fetched
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"parent_id": &schema.Schema{
				Description: `parentId query parameter. Fetch tasks that have this parent Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"root_id": &schema.Schema{
				Description: `rootId query parameter. Fetch tasks that have this root Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. This is the epoch millisecond start time from which tasks need to be fetched
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. Fetch tasks that have this status. Available values : PENDING, FAILURE, SUCCESS
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTasksCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vParentID, okParentID := d.GetOk("parent_id")
	vRootID, okRootID := d.GetOk("root_id")
	vStatus, okStatus := d.GetOk("status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTasksCountV1")
		queryParams1 := catalystcentersdkgo.GetTasksCountV1QueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(int)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(int)
		}
		if okParentID {
			queryParams1.ParentID = vParentID.(string)
		}
		if okRootID {
			queryParams1.RootID = vRootID.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}

		response1, restyResp1, err := client.Task.GetTasksCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTasksCountV1", err,
				"Failure at GetTasksCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTaskGetTasksCountV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTasksCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTaskGetTasksCountV1Item(item *catalystcentersdkgo.ResponseTaskGetTasksCountV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
