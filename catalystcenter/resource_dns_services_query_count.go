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
func resourceDNSServicesQueryCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Devices.

- Retrieves the total number of DNS Services and offers complex filtering and sorting capabilities. For detailed
information about the usage of the API, please refer to the Open API specification document https://github.com/cisco-en-
programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-DNSServices-1.0.0-resolved.yaml
`,

		CreateContext: resourceDNSServicesQueryCountCreate,
		ReadContext:   resourceDNSServicesQueryCountRead,
		DeleteContext: resourceDNSServicesQueryCountDelete,
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
						"xca_lle_rid": &schema.Schema{
							Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
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

func resourceDNSServicesQueryCountCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vXCaLLERID := resourceItem["xca_lle_rid"]

	request1 := expandRequestDNSServicesQueryCountRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1(ctx, "parameters.0", d)

	headerParams1 := catalystcentersdkgo.RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1HeaderParams{}

	headerParams1.XCaLLERID = vXCaLLERID.(string)

	response1, restyResp1, err := client.Devices.RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters(request1, &headerParams1)

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

	vItem1 := flattenDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting RetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFilters response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceDNSServicesQueryCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceDNSServicesQueryCountDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestDNSServicesQueryCountRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1 {
	request := catalystcentersdkgo.RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_time")))) {
		request.StartTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_time")))) {
		request.EndTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filters")))) {
		request.Filters = expandRequestDNSServicesQueryCountRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1FiltersArray(ctx, key+".filters", d)
	}
	return &request
}

func expandRequestDNSServicesQueryCountRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1FiltersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Filters {
	request := []catalystcentersdkgo.RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Filters{}
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
		i := expandRequestDNSServicesQueryCountRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Filters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestDNSServicesQueryCountRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Filters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Filters {
	request := catalystcentersdkgo.RequestDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Filters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".operator")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".operator")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".operator")))) {
		request.Operator = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToSliceString(v)
	}
	return &request
}

func flattenDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Item(item *catalystcentersdkgo.ResponseDevicesRetrievesTheTotalNumberOfDNSServicesForGivenSetOfComplexFiltersV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
