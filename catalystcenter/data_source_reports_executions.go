package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReportsExecutions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Get details of all executions for a given report
`,

		ReadContext: dataSourceReportsExecutionsRead,
		Schema: map[string]*schema.Schema{
			"report_id": &schema.Schema{
				Description: `reportId path parameter. reportId of report
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"data_category": &schema.Schema{
							Description: `data category of the report
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"deliveries": &schema.Schema{
							Description: `Array of available delivery channels
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"execution_count": &schema.Schema{
							Description: `Total number of report executions
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"executions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"end_time": &schema.Schema{
										Description: `Report execution pipeline end time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"errors": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"execution_id": &schema.Schema{
										Description: `Report execution Id.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"process_status": &schema.Schema{
										Description: `Report execution status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"request_status": &schema.Schema{
										Description: `Report execution acceptance status from scheduler
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"start_time": &schema.Schema{
										Description: `Report execution pipeline start time
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"warnings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"name": &schema.Schema{
							Description: `report dataset name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"report_id": &schema.Schema{
							Description: `report Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"report_was_executed": &schema.Schema{
							Description: `true if atleast one execution has started
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"schedule": &schema.Schema{
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"tags": &schema.Schema{
							Description: `array of tags for report
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"view": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `view description
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"field_groups": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"format": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `view name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"view_id": &schema.Schema{
										Description: `view Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"view_info": &schema.Schema{
										Description: `view filters info
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"view_group_id": &schema.Schema{
							Description: `viewGroupId of the viewgroup for the report
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"view_group_version": &schema.Schema{
							Description: `version of viewgroup for the report
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

func dataSourceReportsExecutionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vReportID := d.Get("report_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllExecutionDetailsForAGivenReportV1")
		vvReportID := vReportID.(string)

		response1, restyResp1, err := client.Reports.GetAllExecutionDetailsForAGivenReportV1(vvReportID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllExecutionDetailsForAGivenReportV1", err,
				"Failure at GetAllExecutionDetailsForAGivenReportV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenReportsGetAllExecutionDetailsForAGivenReportV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllExecutionDetailsForAGivenReportV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenReportsGetAllExecutionDetailsForAGivenReportV1Item(item *catalystcentersdkgo.ResponseReportsGetAllExecutionDetailsForAGivenReportV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tags"] = item.Tags
	respItem["data_category"] = item.DataCategory
	respItem["deliveries"] = flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemDeliveries(item.Deliveries)
	respItem["execution_count"] = item.ExecutionCount
	respItem["executions"] = flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemExecutions(item.Executions)
	respItem["name"] = item.Name
	respItem["report_id"] = item.ReportID
	respItem["report_was_executed"] = boolPtrToString(item.ReportWasExecuted)
	respItem["schedule"] = flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemSchedule(item.Schedule)
	respItem["view"] = flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemView(item.View)
	respItem["view_group_id"] = item.ViewGroupID
	respItem["view_group_version"] = item.ViewGroupVersion
	return []map[string]interface{}{
		respItem,
	}
}

func flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemDeliveries(items *[]catalystcentersdkgo.ResponseReportsGetAllExecutionDetailsForAGivenReportV1Deliveries) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemExecutions(items *[]catalystcentersdkgo.ResponseReportsGetAllExecutionDetailsForAGivenReportV1Executions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["end_time"] = item.EndTime
		respItem["errors"] = item.Errors
		respItem["execution_id"] = item.ExecutionID
		respItem["process_status"] = item.ProcessStatus
		respItem["request_status"] = item.RequestStatus
		respItem["start_time"] = item.StartTime
		respItem["warnings"] = item.Warnings
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemSchedule(item *catalystcentersdkgo.ResponseReportsGetAllExecutionDetailsForAGivenReportV1Schedule) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemView(item *catalystcentersdkgo.ResponseReportsGetAllExecutionDetailsForAGivenReportV1View) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["field_groups"] = flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemViewFieldGroups(item.FieldGroups)
	respItem["filters"] = flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemViewFilters(item.Filters)
	respItem["format"] = flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemViewFormat(item.Format)
	respItem["name"] = item.Name
	respItem["view_id"] = item.ViewID
	respItem["description"] = item.Description
	respItem["view_info"] = item.ViewInfo

	return []map[string]interface{}{
		respItem,
	}

}

func flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemViewFieldGroups(items *[]catalystcentersdkgo.ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFieldGroups) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemViewFilters(items *[]catalystcentersdkgo.ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFilters) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenReportsGetAllExecutionDetailsForAGivenReportV1ItemViewFormat(item *catalystcentersdkgo.ResponseReportsGetAllExecutionDetailsForAGivenReportV1ViewFormat) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
