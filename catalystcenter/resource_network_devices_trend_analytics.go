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
func resourceNetworkDevicesTrendAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Gets the Trend analytics Network device data for the given time range. The data will be grouped based on the given
trend time Interval. The required property for this API is *trendInterval*. For detailed information about the usage of
the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceNetworkDevices-1.0.2-resolved.yaml
`,

		CreateContext: resourceNetworkDevicesTrendAnalyticsCreate,
		ReadContext:   resourceNetworkDevicesTrendAnalyticsRead,
		DeleteContext: resourceNetworkDevicesTrendAnalyticsDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
							Description: `Aggregate Attributes`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aggregate_attributes": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"function": &schema.Schema{
													Description: `Function`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeFloat,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"groups": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aggregate_attributes": &schema.Schema{
													Type:     schema.TypeList,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"function": &schema.Schema{
																Description: `Function`,
																Type:        schema.TypeString,
																ForceNew:    true,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																ForceNew:    true,
																Computed:    true,
															},
															"value": &schema.Schema{
																Description: `Value`,
																Type:        schema.TypeFloat,
																ForceNew:    true,
																Computed:    true,
															},
														},
													},
												},
												"attributes": &schema.Schema{
													Type:     schema.TypeList,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																ForceNew:    true,
																Computed:    true,
															},
															"value": &schema.Schema{
																Description: `Value`,
																Type:        schema.TypeString,
																ForceNew:    true,
																Computed:    true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"timestamp": &schema.Schema{
										Description: `Timestamp`,
										Type:        schema.TypeFloat,
										ForceNew:    true,
										Computed:    true,
									},
								},
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
									"timestamp_order": &schema.Schema{
										Description: `Timestamp Order`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
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
						"trend_interval": &schema.Schema{
							Description: `Trend Interval`,
							Type:        schema.TypeString,
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

func resourceNetworkDevicesTrendAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.GetsTheTrendAnalyticsDataV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing GetsTheTrendAnalyticsDataV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenDevicesGetsTheTrendAnalyticsDataV1Items(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetsTheTrendAnalyticsDataV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceNetworkDevicesTrendAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDevicesTrendAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1 {
	request := catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trend_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trend_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trend_interval")))) {
		request.TrendInterval = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_by")))) {
		request.GroupBy = expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1GroupByArray(ctx, key+".group_by", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1FiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1AggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1Page(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1GroupByArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1GroupBy {
	request := []catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1GroupBy{}
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
		i := expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1GroupBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1GroupBy(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1GroupBy {
	var request catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1GroupBy
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1FiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1Filters {
	request := []catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1Filters{}
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
		i := expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1Filters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1Filters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1Filters {
	request := catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1Filters{}
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

func expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1AggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1AggregateAttributes {
	request := []catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1AggregateAttributes{}
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
		i := expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1AggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1AggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1AggregateAttributes {
	var request catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1AggregateAttributes
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestNetworkDevicesTrendAnalyticsGetsTheTrendAnalyticsDataV1Page(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1Page {
	request := catalystcentersdkgo.RequestDevicesGetsTheTrendAnalyticsDataV1Page{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".timestamp_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".timestamp_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".timestamp_order")))) {
		request.TimestampOrder = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetsTheTrendAnalyticsDataV1Items(items *[]catalystcentersdkgo.ResponseDevicesGetsTheTrendAnalyticsDataV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["attributes"] = flattenDevicesGetsTheTrendAnalyticsDataV1ItemsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesGetsTheTrendAnalyticsDataV1ItemsAggregateAttributes(item.AggregateAttributes)
		respItem["groups"] = flattenDevicesGetsTheTrendAnalyticsDataV1ItemsGroups(item.Groups)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheTrendAnalyticsDataV1ItemsAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseAttributes) []map[string]interface{} {
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

func flattenDevicesGetsTheTrendAnalyticsDataV1ItemsAggregateAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseAggregateAttributes) []map[string]interface{} {
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

func flattenDevicesGetsTheTrendAnalyticsDataV1ItemsGroups(items *[]catalystcentersdkgo.ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenDevicesGetsTheTrendAnalyticsDataV1ItemsGroupsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesGetsTheTrendAnalyticsDataV1ItemsGroupsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetsTheTrendAnalyticsDataV1ItemsGroupsAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroupsAttributes) []map[string]interface{} {
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

func flattenDevicesGetsTheTrendAnalyticsDataV1ItemsGroupsAggregateAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetsTheTrendAnalyticsDataV1ResponseGroupsAggregateAttributes) []map[string]interface{} {
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
