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
func resourceNetworkDevicesQueryCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Gets the total number Network Devices based on the provided complex filters and aggregation functions. For detailed
information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceNetworkDevices-1.0.2-resolved.yaml
`,

		CreateContext: resourceNetworkDevicesQueryCountCreate,
		ReadContext:   resourceNetworkDevicesQueryCountRead,
		DeleteContext: resourceNetworkDevicesQueryCountDelete,
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

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
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
						"views": &schema.Schema{
							Description: `Views`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceNetworkDevicesQueryCountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceNetworkDevicesQueryCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceNetworkDevicesQueryCountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1 {
	request := catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".views")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".views")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".views")))) {
		request.Views = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1FiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1FiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters {
	request := []catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters{}
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
		i := expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters {
	request := catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Filters{}
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

func expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes {
	request := []catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes{}
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
		i := expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes {
	request := catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1AggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page {
	request := catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Page{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy {
	request := []catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy{}
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
		i := expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestNetworkDevicesQueryCountGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy {
	request := catalystcentersdkgo.RequestDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1PageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Item(item *catalystcentersdkgo.ResponseDevicesGetsTheTotalNumberNetworkDevicesBasedOnTheProvidedComplexFiltersAndAggregationFunctionsV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
