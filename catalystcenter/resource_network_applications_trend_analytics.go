package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceNetworkApplicationsTrendAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Applications.

- Retrieves the trend analytics of applications experience data for the specified time range. The data will be grouped
based on the given trend time interval. This data source action facilitates obtaining consolidated insights into the
performance and status of the network applications over the specified start and end time. If startTime and endTime are
not provided, the API defaults to the last 24 hours. **siteId** and **trendInterval** are mandatory. **siteId** must be
a site UUID of a building. For detailed information about the usage of the API, please refer to the Open API
specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-NetworkApplications-1.0.0-resolved.yaml
`,

		CreateContext: resourceNetworkApplicationsTrendAnalyticsCreate,
		ReadContext:   resourceNetworkApplicationsTrendAnalyticsRead,
		DeleteContext: resourceNetworkApplicationsTrendAnalyticsDelete,
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
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
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
										Type:        schema.TypeInt,
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
													Type:        schema.TypeString,
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
																Type:        schema.TypeString,
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
										Type:        schema.TypeInt,
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

									"cursor": &schema.Schema{
										Description: `Cursor`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"limit": &schema.Schema{
										Description: `Limit`,
										Type:        schema.TypeInt,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"time_sort_order": &schema.Schema{
										Description: `Time Sort Order`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"site_ids": &schema.Schema{
							Description: `Site Ids`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
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

func resourceNetworkApplicationsTrendAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications(ctx, "parameters.0", d)

	headerParams1 := catalystcentersdkgo.RetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsHeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	response1, restyResp1, err := client.Applications.RetrievesTheTrendAnalyticsDataRelatedToNetworkApplications(request1, &headerParams1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	//Analizar verificacion.

	vItems1 := flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItems(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetrievesTheTrendAnalyticsDataRelatedToNetworkApplications response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceNetworkApplicationsTrendAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkApplicationsTrendAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications {
	request := catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplications{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trend_interval")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trend_interval")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trend_interval")))) {
		request.TrendInterval = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_by")))) {
		request.GroupBy = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters {
	request := []catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters{}
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
		i := expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters {
	request := catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToIntPtr(v)
	}
	return &request
}

func expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes {
	request := []catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes{}
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
		i := expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes {
	request := catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsAggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkApplicationsTrendAnalyticsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage {
	request := catalystcentersdkgo.RequestApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsPage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".cursor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".cursor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".cursor")))) {
		request.Cursor = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".time_sort_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".time_sort_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".time_sort_order")))) {
		request.TimeSortOrder = interfaceToString(v)
	}
	return &request
}

func flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItems(items *[]catalystcentersdkgo.ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["timestamp"] = item.Timestamp
		respItem["attributes"] = flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsAggregateAttributes(item.AggregateAttributes)
		respItem["groups"] = flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsGroups(item.Groups)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsAttributes(items *[]catalystcentersdkgo.ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAttributes) []map[string]interface{} {
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

func flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsAggregateAttributes(items *[]catalystcentersdkgo.ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseAggregateAttributes) []map[string]interface{} {
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

func flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsGroups(items *[]catalystcentersdkgo.ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsGroupsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsGroupsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsGroupsAttributes(items *[]catalystcentersdkgo.ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAttributes) []map[string]interface{} {
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

func flattenApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsItemsGroupsAggregateAttributes(items *[]catalystcentersdkgo.ResponseApplicationsRetrievesTheTrendAnalyticsDataRelatedToNetworkApplicationsResponseGroupsAggregateAttributes) []map[string]interface{} {
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
