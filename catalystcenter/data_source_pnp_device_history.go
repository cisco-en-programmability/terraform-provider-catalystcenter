package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpDeviceHistory() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns history for a specific device. Serial number is a required parameter
`,

		ReadContext: dataSourcePnpDeviceHistoryRead,
		Schema: map[string]*schema.Schema{
			"serial_number": &schema.Schema{
				Description: `serialNumber query parameter. Device Serial Number
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"sort": &schema.Schema{
				Description: `sort query parameter. Comma seperated list of fields to sort on
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (des)
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"details": &schema.Schema{
							Description: `Details`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"error_flag": &schema.Schema{
							Description: `Error Flag`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"history_task_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"addn_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Description: `Key`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"time_taken": &schema.Schema{
										Description: `Time Taken`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"work_item_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"command": &schema.Schema{
													Description: `Command`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"output_str": &schema.Schema{
													Description: `Output Str`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},

												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"time_taken": &schema.Schema{
													Description: `Time Taken`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},

						"timestamp": &schema.Schema{
							Description: `Timestamp`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnpDeviceHistoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSerialNumber := d.Get("serial_number")
	vSort, okSort := d.GetOk("sort")
	vSortOrder, okSortOrder := d.GetOk("sort_order")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceHistoryV1")
		queryParams1 := catalystcentersdkgo.GetDeviceHistoryV1QueryParams{}

		queryParams1.SerialNumber = vSerialNumber.(string)

		if okSort {
			queryParams1.Sort = interfaceToSliceString(vSort)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetDeviceHistoryV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceHistoryV1", err,
				"Failure at GetDeviceHistoryV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceOnboardingPnpGetDeviceHistoryV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceHistoryV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetDeviceHistoryV1Items(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpGetDeviceHistoryV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["details"] = item.Details
		respItem["history_task_info"] = flattenDeviceOnboardingPnpGetDeviceHistoryV1ItemsHistoryTaskInfo(item.HistoryTaskInfo)
		respItem["error_flag"] = boolPtrToString(item.ErrorFlag)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceHistoryV1ItemsHistoryTaskInfo(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["type"] = item.Type
	respItem["time_taken"] = item.TimeTaken
	respItem["work_item_list"] = flattenDeviceOnboardingPnpGetDeviceHistoryV1ItemsHistoryTaskInfoWorkItemList(item.WorkItemList)
	respItem["addn_details"] = flattenDeviceOnboardingPnpGetDeviceHistoryV1ItemsHistoryTaskInfoAddnDetails(item.AddnDetails)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpGetDeviceHistoryV1ItemsHistoryTaskInfoWorkItemList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfoWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["start_time"] = item.StartTime
		respItem["end_time"] = item.EndTime
		respItem["time_taken"] = item.TimeTaken
		respItem["output_str"] = item.OutputStr
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpGetDeviceHistoryV1ItemsHistoryTaskInfoAddnDetails(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpGetDeviceHistoryV1ResponseHistoryTaskInfoAddnDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
