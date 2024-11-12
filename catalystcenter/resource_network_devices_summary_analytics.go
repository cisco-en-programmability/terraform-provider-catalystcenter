package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkDevicesSummaryAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Gets the summary analytics data related to network devices based on the provided input data. This endpoint helps to
obtain the consolidated insights into the performance and status of the monitored network devices. For detailed
information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceNetworkDevices-1.0.2-resolved.yaml
`,

		CreateContext: resourceNetworkDevicesSummaryAnalyticsCreate,
		ReadContext:   resourceNetworkDevicesSummaryAnalyticsRead,
		DeleteContext: resourceNetworkDevicesSummaryAnalyticsDelete,
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

						"aggregate_attributes": &schema.Schema{
							Description: `Aggregate Attributes`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"attributes": &schema.Schema{
							Description: `Attributes`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"groups": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aggregate_attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"function": &schema.Schema{
													Description: `Function`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name`,
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
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
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
						"aggregate_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"function": &schema.Schema{
										Description: `Function`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"attributes": &schema.Schema{
							Description: `Attributes`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"end_time": &schema.Schema{
							Description: `End Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"filters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"operator": &schema.Schema{
										Description: `Operator`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"group_by": &schema.Schema{
							Description: `Group By`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"page": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"limit": &schema.Schema{
										Description: `Limit`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"offset": &schema.Schema{
										Description: `Offset`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"sort_by": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"order": &schema.Schema{
													Description: `Order`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `Start Time`,
							Type:        schema.TypeInt,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDevicesSummaryAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceNetworkDevicesSummaryAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDevicesSummaryAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1 {
	request := catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_by")))) {
		request.GroupBy = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1FiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Page(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1FiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Filters {
	request := []catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Filters{}
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
		i := expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Filters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Filters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Filters {
	request := catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Filters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes {
	request := []catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes{}
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
		i := expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes {
	request := catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1AggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Page(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Page {
	request := catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Page{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy {
	request := []catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy{}
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
		i := expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesSummaryAnalyticsGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy {
	request := catalystcentersdkgo.RequestDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1PageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Item(item *catalystcentersdkgo.ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attributes"] = flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemAttributes(item.Attributes)
	respItem["aggregate_attributes"] = flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemAggregateAttributes(item.AggregateAttributes)
	respItem["groups"] = flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemGroups(item.Groups)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseAttributes) []interface{} {
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

func flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemAggregateAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseAggregateAttributes) []interface{} {
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

func flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemGroups(items *[]catalystcentersdkgo.ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemGroupsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemGroupsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemGroupsAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroupsAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ItemGroupsAggregateAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetsTheSummaryAnalyticsDataRelatedToNetworkDevicesV1ResponseGroupsAggregateAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["function"] = item.Function
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
