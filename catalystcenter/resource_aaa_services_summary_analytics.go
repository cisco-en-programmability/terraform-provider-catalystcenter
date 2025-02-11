package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAAAServicesSummaryAnalyticsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(ctx, "parameters.0", d)

	headerParams1 := catalystcentersdkgo.GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1HeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	response1, restyResp1, err := client.Devices.GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters(request1, &headerParams1)

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

	vItem1 := flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFilters response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceAAAServicesSummaryAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAAAServicesSummaryAnalyticsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1 {
	request := catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_by")))) {
		request.GroupBy = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersArray(ctx, key+".filters", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".aggregate_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".aggregate_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".aggregate_attributes")))) {
		request.AggregateAttributes = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributesArray(ctx, key+".aggregate_attributes", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".page")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".page")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".page")))) {
		request.Page = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page(ctx, key+".page.0", d)
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters {
	request := []catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters{}
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
		i := expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters {
	request := catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Filters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".logical_operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".logical_operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".logical_operator")))) {
		request.LogicalOperator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue(ctx, key+".value.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue {
	var request catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1FiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes {
	request := []catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes{}
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
		i := expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes {
	request := catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1AggregateAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page {
	request := catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Page{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".limit")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".limit")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".limit")))) {
		request.Limit = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".offset")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".offset")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".offset")))) {
		request.Offset = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sort_by")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sort_by")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sort_by")))) {
		request.SortBy = expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortByArray(ctx, key+".sort_by", d)
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortByArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy {
	request := []catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy{}
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
		i := expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAAAServicesSummaryAnalyticsGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy {
	request := catalystcentersdkgo.RequestDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1PageSortBy{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".function")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".function")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".function")))) {
		request.Function = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToString(v)
	}
	return &request
}

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Item(item *catalystcentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["groups"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemGroups(item.Groups)
	respItem["attributes"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemAttributes(item.Attributes)
	respItem["aggregate_attributes"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemAggregateAttributes(item.AggregateAttributes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemGroups(items *[]catalystcentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroups) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemGroupsAttributes(item.Attributes)
		respItem["aggregate_attributes"] = flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemGroupsAggregateAttributes(item.AggregateAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemGroupsAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAttributes) []map[string]interface{} {
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

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemGroupsAggregateAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseGroupsAggregateAttributes) []map[string]interface{} {
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

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAttributes) []map[string]interface{} {
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

func flattenDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ItemAggregateAttributes(items *[]catalystcentersdkgo.ResponseDevicesGetSummaryAnalyticsDataOfAAAServicesForGivenSetOfComplexFiltersV1ResponseAggregateAttributes) []map[string]interface{} {
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
