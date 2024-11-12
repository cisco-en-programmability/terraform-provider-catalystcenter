package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAnalyticsTasks() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on AI Endpoint Analytics.

- Fetches the details of backend task. Task is typically created by making call to some other API that takes longer time
to execute.
`,

		ReadContext: dataSourceAnalyticsTasksRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Description: `taskId path parameter. Unique identifier for the task.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_info": &schema.Schema{
							Description: `Additional information about the task.
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"created_by": &schema.Schema{
							Description: `Name of the user that created the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"created_on": &schema.Schema{
							Description: `Task creation timestamp in epoch milliseconds.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"errors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"code": &schema.Schema{
										Description: `Error code.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"details": &schema.Schema{
										Description: `Optional details about the error.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"index": &schema.Schema{
										Description: `Index of the input records which had error during processing. In case the input is not an array, or the error is not record specific, this will be -1.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"message": &schema.Schema{
										Description: `Error message.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Unique identifier for the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_on": &schema.Schema{
							Description: `Last update timestamp in epoch milliseconds.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAnalyticsTasksRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vTaskID := d.Get("task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTaskDetailsV1")
		vvTaskID := vTaskID.(string)

		response1, restyResp1, err := client.AIEndpointAnalytics.GetTaskDetailsV1(vvTaskID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTaskDetailsV1", err,
				"Failure at GetTaskDetailsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenAIEndpointAnalyticsGetTaskDetailsV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTaskDetailsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAIEndpointAnalyticsGetTaskDetailsV1Item(item *catalystcentersdkgo.ResponseAIEndpointAnalyticsGetTaskDetailsV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["status"] = item.Status
	respItem["errors"] = flattenAIEndpointAnalyticsGetTaskDetailsV1ItemErrors(item.Errors)
	respItem["additional_info"] = flattenAIEndpointAnalyticsGetTaskDetailsV1ItemAdditionalInfo(item.AdditionalInfo)
	respItem["created_by"] = item.CreatedBy
	respItem["created_on"] = item.CreatedOn
	respItem["last_updated_on"] = item.LastUpdatedOn
	return []map[string]interface{}{
		respItem,
	}
}

func flattenAIEndpointAnalyticsGetTaskDetailsV1ItemErrors(items *[]catalystcentersdkgo.ResponseAIEndpointAnalyticsGetTaskDetailsV1Errors) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["index"] = item.Index
		respItem["code"] = item.Code
		respItem["message"] = item.Message
		respItem["details"] = item.Details
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenAIEndpointAnalyticsGetTaskDetailsV1ItemAdditionalInfo(item *catalystcentersdkgo.ResponseAIEndpointAnalyticsGetTaskDetailsV1AdditionalInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
