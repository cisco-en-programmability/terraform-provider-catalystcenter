package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReports() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Get list of scheduled report configurations.

- Get scheduled report configuration by reportId
`,

		ReadContext: dataSourceReportsRead,
		Schema: map[string]*schema.Schema{
			"report_id": &schema.Schema{
				Description: `reportId path parameter. reportId of report
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"view_group_id": &schema.Schema{
				Description: `viewGroupId query parameter. viewGroupId of viewgroup for report
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"view_id": &schema.Schema{
				Description: `viewId query parameter. viewId of view for report
`,
				Type:     schema.TypeString,
				Optional: true,
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
							Description: `report name
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"field_group_display_name": &schema.Schema{
													Description: `Field group label/displayname for user
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"field_group_name": &schema.Schema{
													Description: `Field group name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"fields": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"display_name": &schema.Schema{
																Description: `field label/displayname
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"name": &schema.Schema{
																Description: `field name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `filter label/displayname
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `filter name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `filter type
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"value": &schema.Schema{
													Description: `value of filter. data type is based on the filter type.
`,
													Type:     schema.TypeString, //TEST,
													Computed: true,
												},
											},
										},
									},

									"format": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default": &schema.Schema{
													Description: `true, if the format type is considered default
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"format_type": &schema.Schema{
													Description: `format type of report
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `format name of report
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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

			"items": &schema.Schema{
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
							Description: `report name
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
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"field_group_display_name": &schema.Schema{
													Description: `Field group label/displayname for user
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"field_group_name": &schema.Schema{
													Description: `Field group name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"fields": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"display_name": &schema.Schema{
																Description: `field label/displayname
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"name": &schema.Schema{
																Description: `field name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `filter label/displayname
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `filter name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `filter type
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"value": &schema.Schema{
													Description: `value of filter. data type is based on the filter type.
`,
													Type:     schema.TypeString, //TEST,
													Computed: true,
												},
											},
										},
									},

									"format": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default": &schema.Schema{
													Description: `true, if the format type is considered default
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"format_type": &schema.Schema{
													Description: `format type of report
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `format name of report
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
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

func dataSourceReportsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vViewGroupID, okViewGroupID := d.GetOk("view_group_id")
	vViewID, okViewID := d.GetOk("view_id")
	vReportID, okReportID := d.GetOk("report_id")

	method1 := []bool{okViewGroupID, okViewID}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okReportID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetListOfScheduledReportsV1")
		queryParams1 := catalystcentersdkgo.GetListOfScheduledReportsV1QueryParams{}

		if okViewGroupID {
			queryParams1.ViewGroupID = vViewGroupID.(string)
		}
		if okViewID {
			queryParams1.ViewID = vViewID.(string)
		}

		response1, restyResp1, err := client.Reports.GetListOfScheduledReportsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetListOfScheduledReportsV1", err,
				"Failure at GetListOfScheduledReportsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenReportsGetListOfScheduledReportsV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetListOfScheduledReportsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetAScheduledReportV1")
		vvReportID := vReportID.(string)

		response2, restyResp2, err := client.Reports.GetAScheduledReportV1(vvReportID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAScheduledReportV1", err,
				"Failure at GetAScheduledReportV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenReportsGetAScheduledReportV1Item(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAScheduledReportV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenReportsGetListOfScheduledReportsV1Items(items *catalystcentersdkgo.ResponseReportsGetListOfScheduledReportsV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = item.Tags
		respItem["data_category"] = item.DataCategory
		respItem["deliveries"] = flattenReportsGetListOfScheduledReportsV1ItemsDeliveries(item.Deliveries)
		respItem["execution_count"] = item.ExecutionCount
		respItem["executions"] = flattenReportsGetListOfScheduledReportsV1ItemsExecutions(item.Executions)
		respItem["name"] = item.Name
		respItem["report_id"] = item.ReportID
		respItem["report_was_executed"] = boolPtrToString(item.ReportWasExecuted)
		respItem["schedule"] = flattenReportsGetListOfScheduledReportsV1ItemsSchedule(item.Schedule)
		respItem["view"] = flattenReportsGetListOfScheduledReportsV1ItemsView(item.View)
		respItem["view_group_id"] = item.ViewGroupID
		respItem["view_group_version"] = item.ViewGroupVersion
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetListOfScheduledReportsV1ItemsDeliveries(items *[]catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1Deliveries) []interface{} {
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

func flattenReportsGetListOfScheduledReportsV1ItemsExecutions(items *[]catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1Executions) []map[string]interface{} {
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

func flattenReportsGetListOfScheduledReportsV1ItemsSchedule(item *catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1Schedule) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenReportsGetListOfScheduledReportsV1ItemsView(item *catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1View) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["field_groups"] = flattenReportsGetListOfScheduledReportsV1ItemsViewFieldGroups(item.FieldGroups)
	respItem["filters"] = flattenReportsGetListOfScheduledReportsV1ItemsViewFilters(item.Filters)
	respItem["format"] = flattenReportsGetListOfScheduledReportsV1ItemsViewFormat(item.Format)
	respItem["name"] = item.Name
	respItem["view_id"] = item.ViewID
	respItem["description"] = item.Description
	respItem["view_info"] = item.ViewInfo

	return []map[string]interface{}{
		respItem,
	}

}

func flattenReportsGetListOfScheduledReportsV1ItemsViewFieldGroups(items *[]catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1ViewFieldGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["field_group_display_name"] = item.FieldGroupDisplayName
		respItem["field_group_name"] = item.FieldGroupName
		respItem["fields"] = flattenReportsGetListOfScheduledReportsV1ItemsViewFieldGroupsFields(item.Fields)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetListOfScheduledReportsV1ItemsViewFieldGroupsFields(items *[]catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1ViewFieldGroupsFields) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["display_name"] = item.DisplayName
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetListOfScheduledReportsV1ItemsViewFilters(items *[]catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1ViewFilters) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["display_name"] = item.DisplayName
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["value"] = flattenReportsGetListOfScheduledReportsV1ItemsViewFiltersValue(item.Value)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetListOfScheduledReportsV1ItemsViewFiltersValue(item *catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1ViewFiltersValue) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenReportsGetListOfScheduledReportsV1ItemsViewFormat(item *catalystcentersdkgo.ResponseItemReportsGetListOfScheduledReportsV1ViewFormat) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["format_type"] = item.FormatType
	respItem["name"] = item.Name
	respItem["default"] = boolPtrToString(item.Default)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenReportsGetAScheduledReportV1Item(item *catalystcentersdkgo.ResponseReportsGetAScheduledReportV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tags"] = item.Tags
	respItem["data_category"] = item.DataCategory
	respItem["deliveries"] = flattenReportsGetAScheduledReportV1ItemDeliveries(item.Deliveries)
	respItem["execution_count"] = item.ExecutionCount
	respItem["executions"] = flattenReportsGetAScheduledReportV1ItemExecutions(item.Executions)
	respItem["name"] = item.Name
	respItem["report_id"] = item.ReportID
	respItem["report_was_executed"] = boolPtrToString(item.ReportWasExecuted)
	respItem["schedule"] = flattenReportsGetAScheduledReportV1ItemSchedule(item.Schedule)
	respItem["view"] = flattenReportsGetAScheduledReportV1ItemView(item.View)
	respItem["view_group_id"] = item.ViewGroupID
	respItem["view_group_version"] = item.ViewGroupVersion
	return []map[string]interface{}{
		respItem,
	}
}

func flattenReportsGetAScheduledReportV1ItemDeliveries(items *[]catalystcentersdkgo.ResponseReportsGetAScheduledReportV1Deliveries) []interface{} {
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

func flattenReportsGetAScheduledReportV1ItemExecutions(items *[]catalystcentersdkgo.ResponseReportsGetAScheduledReportV1Executions) []map[string]interface{} {
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

func flattenReportsGetAScheduledReportV1ItemSchedule(item *catalystcentersdkgo.ResponseReportsGetAScheduledReportV1Schedule) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenReportsGetAScheduledReportV1ItemView(item *catalystcentersdkgo.ResponseReportsGetAScheduledReportV1View) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["field_groups"] = flattenReportsGetAScheduledReportV1ItemViewFieldGroups(item.FieldGroups)
	respItem["filters"] = flattenReportsGetAScheduledReportV1ItemViewFilters(item.Filters)
	respItem["format"] = flattenReportsGetAScheduledReportV1ItemViewFormat(item.Format)
	respItem["name"] = item.Name
	respItem["view_id"] = item.ViewID
	respItem["description"] = item.Description
	respItem["view_info"] = item.ViewInfo

	return []map[string]interface{}{
		respItem,
	}

}

func flattenReportsGetAScheduledReportV1ItemViewFieldGroups(items *[]catalystcentersdkgo.ResponseReportsGetAScheduledReportV1ViewFieldGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["field_group_display_name"] = item.FieldGroupDisplayName
		respItem["field_group_name"] = item.FieldGroupName
		respItem["fields"] = flattenReportsGetAScheduledReportV1ItemViewFieldGroupsFields(item.Fields)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetAScheduledReportV1ItemViewFieldGroupsFields(items *[]catalystcentersdkgo.ResponseReportsGetAScheduledReportV1ViewFieldGroupsFields) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["display_name"] = item.DisplayName
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetAScheduledReportV1ItemViewFilters(items *[]catalystcentersdkgo.ResponseReportsGetAScheduledReportV1ViewFilters) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["display_name"] = item.DisplayName
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["value"] = flattenReportsGetAScheduledReportV1ItemViewFiltersValue(item.Value)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetAScheduledReportV1ItemViewFiltersValue(item *catalystcentersdkgo.ResponseReportsGetAScheduledReportV1ViewFiltersValue) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenReportsGetAScheduledReportV1ItemViewFormat(item *catalystcentersdkgo.ResponseReportsGetAScheduledReportV1ViewFormat) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["format_type"] = item.FormatType
	respItem["name"] = item.Name
	respItem["default"] = boolPtrToString(item.Default)

	return []map[string]interface{}{
		respItem,
	}

}
