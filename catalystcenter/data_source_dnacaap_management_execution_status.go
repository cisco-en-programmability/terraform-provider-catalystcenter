package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDnacaapManagementExecutionStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- Retrieves the execution details of a Business API
`,

		ReadContext: dataSourceDnacaapManagementExecutionStatusRead,
		Schema: map[string]*schema.Schema{
			"execution_id": &schema.Schema{
				Description: `executionId path parameter. Execution Id of API
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"bapi_error": &schema.Schema{
							Description: `Returns the error response of the original API  as a string
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"bapi_execution_id": &schema.Schema{
							Description: `Execution Id of the Business API (UUID)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"bapi_key": &schema.Schema{
							Description: `Business API Key (UUID)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"bapi_name": &schema.Schema{
							Description: `Name of the Business API
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"bapi_sync_response": &schema.Schema{
							Description: `Returns the actual response of the original API as a string
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"bapi_sync_response_json": &schema.Schema{
							Description: `Returns the actual response of the original API  as a json
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"end_time": &schema.Schema{
							Description: `Execution End Time of the Business API (Date Time Format)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_time_epoch": &schema.Schema{
							Description: `Execution End Time of the Business API (Epoch Milliseconds)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"runtime_instance_id": &schema.Schema{
							Description: `Pod Id in which the Business API is executed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Description: `Execution Start Time of the Business API (Date Time Format)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"start_time_epoch": &schema.Schema{
							Description: `Execution Start Time of the Business API (Epoch Milliseconds)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Execution status of the Business API
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"time_duration": &schema.Schema{
							Description: `Time taken for Business API Execution (Milliseconds)
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

func dataSourceDnacaapManagementExecutionStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vExecutionID := d.Get("execution_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetBusinessAPIExecutionDetails")
		vvExecutionID := vExecutionID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(vvExecutionID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTaskGetBusinessAPIExecutionDetailsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetBusinessAPIExecutionDetails response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTaskGetBusinessAPIExecutionDetailsItem(item *catalystcentersdkgo.ResponseTaskGetBusinessAPIExecutionDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["bapi_key"] = item.BapiKey
	respItem["bapi_name"] = item.BapiName
	respItem["bapi_execution_id"] = item.BapiExecutionID
	respItem["start_time"] = item.StartTime
	respItem["start_time_epoch"] = item.StartTimeEpoch
	respItem["end_time"] = item.EndTime
	respItem["end_time_epoch"] = item.EndTimeEpoch
	respItem["time_duration"] = item.TimeDuration
	respItem["status"] = item.Status
	respItem["bapi_sync_response"] = item.BapiSyncResponse
	respItem["bapi_sync_response_json"] = flattenTaskGetBusinessAPIExecutionDetailsItemBapiSyncResponseJSON(item.BapiSyncResponseJSON)
	respItem["runtime_instance_id"] = item.RuntimeInstanceID
	respItem["bapi_error"] = item.BapiError
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTaskGetBusinessAPIExecutionDetailsItemBapiSyncResponseJSON(item *catalystcentersdkgo.ResponseTaskGetBusinessAPIExecutionDetailsBapiSyncResponseJSON) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
