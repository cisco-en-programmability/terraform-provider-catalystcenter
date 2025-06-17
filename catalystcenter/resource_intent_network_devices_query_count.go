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
													Description: `Value to filter by. For **in** operator, the value should be a list of values.
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

	request1 := expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFilters(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.CountTheNumberOfNetworkDevicesWithFilters(request1)

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

	vItem1 := flattenDevicesCountTheNumberOfNetworkDevicesWithFiltersItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CountTheNumberOfNetworkDevicesWithFilters response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//Analizar verificacion.

}
func resourceIntentNetworkDevicesQueryCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIntentNetworkDevicesQueryCountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFilters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFilters {
	request := catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFilters{}
	request.Filter = expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersFilter(ctx, key, d)
	return &request
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersFilter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilter {
	request := catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".logical_operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".logical_operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".logical_operator")))) {
		request.LogicalOperator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersFilterFiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersFilterFiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFilters {
	request := []catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFilters{}
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
		i := expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersFilterFilters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersFilterFilters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFilters {
	request := catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFilters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersFilterFiltersValue(ctx, key+".value.0", d)
	}
	return &request
}

func expandRequestIntentNetworkDevicesQueryCountCountTheNumberOfNetworkDevicesWithFiltersFilterFiltersValue(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFiltersValue {
	var request catalystcentersdkgo.RequestDevicesCountTheNumberOfNetworkDevicesWithFiltersFilterFiltersValue
	request = d.Get(fixKeyAccess(key))
	return &request
}

func flattenDevicesCountTheNumberOfNetworkDevicesWithFiltersItem(item *catalystcentersdkgo.ResponseDevicesCountTheNumberOfNetworkDevicesWithFiltersResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
