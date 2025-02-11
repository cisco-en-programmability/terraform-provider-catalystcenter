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

// resourceAction
func resourceAssuranceIssuesQuery() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Issues.

- Returns all details of each issue along with suggested actions for given set of filters specified in request body. If
there is no start and/or end time, then end time will be defaulted to current time and start time will be defaulted to
24-hours ago from end time. https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-IssuesList-1.0.0-resolved.yaml
`,

		CreateContext: resourceAssuranceIssuesQueryCreate,
		ReadContext:   resourceAssuranceIssuesQueryRead,
		DeleteContext: resourceAssuranceIssuesQueryDelete,
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
						"accept_language": &schema.Schema{
							Description: `Accept-Language header parameter. This header parameter can be used to specify the language in which issue description and suggested actions need to be returned. Available options are 'en' (English), 'ja' (Japanese), 'ko' (Korean), 'zh' (Chinese). If this parameter is not present the issue details are returned in English language.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
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
									"key": &schema.Schema{
										Description: `Key`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"logical_operator": &schema.Schema{
										Description: `Logical Operator`,
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
						"items": &schema.Schema{
							Type:     schema.TypeList,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"additional_attributes": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Description: `Key`,
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
									"category": &schema.Schema{
										Description: `Category`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"device_type": &schema.Schema{
										Description: `Device Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"entity_id": &schema.Schema{
										Description: `Entity Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"entity_type": &schema.Schema{
										Description: `Entity Type`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"first_occurred_time": &schema.Schema{
										Description: `First Occurred Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"is_global": &schema.Schema{
										Description: `Is Global`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										ForceNew: true,
										Computed: true,
									},
									"issue_id": &schema.Schema{
										Description: `Issue Id`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"most_recent_occurred_time": &schema.Schema{
										Description: `Most Recent Occurred Time`,
										Type:        schema.TypeInt,
										ForceNew:    true,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"notes": &schema.Schema{
										Description: `Notes`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"priority": &schema.Schema{
										Description: `Priority`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"severity": &schema.Schema{
										Description: `Severity`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site Hierarchy`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"site_hierarchy_id": &schema.Schema{
										Description: `Site Hierarchy Id`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"site_id": &schema.Schema{
										Description: `Site Id`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"site_name": &schema.Schema{
										Description: `Site Name`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"status": &schema.Schema{
										Description: `Status`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"suggested_actions": &schema.Schema{
										Type:     schema.TypeList,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"message": &schema.Schema{
													Description: `Message`,
													Type:        schema.TypeString,
													ForceNew:    true,
													Computed:    true,
												},
												"steps": &schema.Schema{
													Description: `Steps`,
													Type:        schema.TypeList,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"summary": &schema.Schema{
										Description: `Summary`,
										Type:        schema.TypeString,
										ForceNew:    true,
										Computed:    true,
									},
									"updated_by": &schema.Schema{
										Description: `Updated By`,
										Type:        schema.TypeString, //TEST,
										ForceNew:    true,
										Computed:    true,
									},
									"updated_time": &schema.Schema{
										Description: `Updated Time`,
										Type:        schema.TypeString, //TEST,
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
					},
				},
			},
		},
	}
}

func resourceAssuranceIssuesQueryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vAcceptLanguage := resourceItem["accept_language"]

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1(ctx, "parameters.0", d)

	headerParams1 := catalystcentersdkgo.GetTheDetailsOfIssuesForGivenSetOfFiltersV1HeaderParams{}

	headerParams1.AcceptLanguage = vAcceptLanguage.(string)

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Issues.GetTheDetailsOfIssuesForGivenSetOfFiltersV1(request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing GetTheDetailsOfIssuesForGivenSetOfFiltersV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItems1 := flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Items(response1.Response)
	if err := d.Set("items", vItems1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting GetTheDetailsOfIssuesForGivenSetOfFiltersV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceAssuranceIssuesQueryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAssuranceIssuesQueryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1 {
	request := catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Filters {
	request := []catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Filters{}
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
		i := expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1Filters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1Filters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Filters {
	request := catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Filters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".logical_operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".logical_operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".logical_operator")))) {
		request.LogicalOperator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFilters {
	request := []catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFilters{}
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
		i := expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAssuranceIssuesQueryGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFilters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFilters {
	request := catalystcentersdkgo.RequestIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1FiltersFilters{}
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

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Items(items *[]catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["issue_id"] = item.IssueID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["summary"] = item.Summary
		respItem["priority"] = item.Priority
		respItem["severity"] = item.Severity
		respItem["device_type"] = item.DeviceType
		respItem["category"] = item.Category
		respItem["entity_type"] = item.EntityType
		respItem["entity_id"] = item.EntityID
		respItem["first_occurred_time"] = item.FirstOccurredTime
		respItem["most_recent_occurred_time"] = item.MostRecentOccurredTime
		respItem["status"] = item.Status
		respItem["is_global"] = boolPtrToString(item.IsGlobal)
		respItem["updated_by"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsUpdatedBy(item.UpdatedBy)
		respItem["updated_time"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsUpdatedTime(item.UpdatedTime)
		respItem["notes"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsNotes(item.Notes)
		respItem["site_id"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSiteID(item.SiteID)
		respItem["site_hierarchy_id"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSiteHierarchyID(item.SiteHierarchyID)
		respItem["site_name"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSiteName(item.SiteName)
		respItem["site_hierarchy"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSiteHierarchy(item.SiteHierarchy)
		respItem["suggested_actions"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSuggestedActions(item.SuggestedActions)
		respItem["additional_attributes"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsAdditionalAttributes(item.AdditionalAttributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsUpdatedBy(item *catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseUpdatedBy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsUpdatedTime(item *catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseUpdatedTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsNotes(item *catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseNotes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSiteID(item *catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSiteHierarchyID(item *catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteHierarchyID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSiteName(item *catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSiteHierarchy(item *catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSiteHierarchy) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSuggestedActions(items *[]catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSuggestedActions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItem["steps"] = flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSuggestedActionsSteps(item.Steps)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsSuggestedActionsSteps(items *[]catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseSuggestedActionsSteps) []interface{} {
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

func flattenIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ItemsAdditionalAttributes(items *[]catalystcentersdkgo.ResponseIssuesGetTheDetailsOfIssuesForGivenSetOfFiltersV1ResponseAdditionalAttributes) []map[string]interface{} {
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
