package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAssuranceTasksID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Task.

- returns a task given a specific task id For detailed information about the usage of the API, please refer to the Open
API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceTasks-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceAssuranceTasksIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. unique task id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data": &schema.Schema{
							Description: `Data`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"error_code": &schema.Schema{
							Description: `Error Code`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"failure_reason": &schema.Schema{
							Description: `Failure Reason`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"progress": &schema.Schema{
							Description: `Progress`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"request_type": &schema.Schema{
							Description: `Request Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"result_url": &schema.Schema{
							Description: `Result Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"update_time": &schema.Schema{
							Description: `Update Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceAssuranceTasksIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveASpecificAssuranceTaskByIDV1")
		vvID := vID.(string)

		headerParams1 := catalystcentersdkgo.RetrieveASpecificAssuranceTaskByIDV1HeaderParams{}

		headerParams1.XCaLLERID = vXCaLLERID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Task.RetrieveASpecificAssuranceTaskByIDV1(vvID, &headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveASpecificAssuranceTaskByIDV1", err,
				"Failure at RetrieveASpecificAssuranceTaskByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTaskRetrieveASpecificAssuranceTaskByIDV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveASpecificAssuranceTaskByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTaskRetrieveASpecificAssuranceTaskByIDV1Item(item *catalystcentersdkgo.ResponseTaskRetrieveASpecificAssuranceTaskByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["status"] = item.Status
	respItem["start_time"] = item.StartTime
	respItem["end_time"] = item.EndTime
	respItem["update_time"] = item.UpdateTime
	respItem["progress"] = item.Progress
	respItem["failure_reason"] = item.FailureReason
	respItem["error_code"] = item.ErrorCode
	respItem["request_type"] = item.RequestType
	respItem["data"] = flattenTaskRetrieveASpecificAssuranceTaskByIDV1ItemData(item.Data)
	respItem["result_url"] = item.ResultURL
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTaskRetrieveASpecificAssuranceTaskByIDV1ItemData(item *catalystcentersdkgo.ResponseTaskRetrieveASpecificAssuranceTaskByIDV1ResponseData) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
