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
func resourceIntentNetworkDevicesQueryCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- API to fetch the count of network devices for the given filter query.
`,

		CreateContext: resourceIntentNetworkDevicesQueryCountCreate,
		ReadContext:   resourceIntentNetworkDevicesQueryCountRead,
		DeleteContext: resourceIntentNetworkDevicesQueryCountDelete,
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
							Description: `The reported count.
`,
							Type:     schema.TypeInt,
							Computed: true,
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
						"filter": &schema.Schema{
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
													Description: `The key to filter by
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"operator": &schema.Schema{
													Description: `The operator to use for filtering the values
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"value": &schema.Schema{
													Description: `Value to filter by. For in operator, the value should be a list of values.
`,
													Type:     schema.TypeString, //TEST,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"logical_operator": &schema.Schema{
										Description: `The logical operator to use for combining the filter criteria. If not provided, the default value is AND.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func resourceIntentNetworkDevicesQueryCountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.CountTheNumberOfNetworkDevicesWithFiltersV1(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vItem1 := flattenDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CountTheNumberOfNetworkDevicesWithFiltersV1 response",
			err))
		return diags
	}
	//Analizar verificacion.

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return diags
}
func resourceIntentNetworkDevicesQueryCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIntentNetworkDevicesQueryCountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1 {
	request := catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1{}
	request.Filter = expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1Filter(ctx, key, d)
	return &request
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1Filter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Filter {
	request := catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Filter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".logical_operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".logical_operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".logical_operator")))) {
		request.LogicalOperator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1FilterFiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1FilterFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFilters {
	request := []catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFilters{}
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
		i := expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1FilterFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1FilterFilters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFilters {
	request := catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1FilterFiltersValue(ctx, key+".value.0", d)
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersV1FilterFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFiltersValue {
	var request catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersV1FilterFiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func flattenDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Item(item *catalystcentersdkgo.ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
