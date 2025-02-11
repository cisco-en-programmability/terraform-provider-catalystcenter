package catalystcenter

import (
	"context"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceLicenseRegister() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Licenses.

- Registers the system with Cisco Smart Software Manager (CSSM)
`,

		CreateContext: resourceLicenseRegisterCreate,
		ReadContext:   resourceLicenseRegisterRead,
		DeleteContext: resourceLicenseRegisterDelete,
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

						"url": &schema.Schema{
							Description: `URL to track the operation status
`,
							Type:     schema.TypeString,
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
						"smart_account_id": &schema.Schema{
							Description: `The ID of the Smart Account to which the system is registered
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"virtual_account_id": &schema.Schema{
							Description: `The ID of the Virtual Account to which the system is registered
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
	}
}

func resourceLicenseRegisterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestLicenseRegisterSystemLicensingRegistrationV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Licenses.SystemLicensingRegistrationV1(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vItem1 := flattenLicensesSystemLicensingRegistrationV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting SystemLicensingRegistrationV1 response",
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
func resourceLicenseRegisterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceLicenseRegisterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestLicenseRegisterSystemLicensingRegistrationV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestLicensesSystemLicensingRegistrationV1 {
	request := catalystcentersdkgo.RequestLicensesSystemLicensingRegistrationV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smart_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smart_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smart_account_id")))) {
		request.SmartAccountID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_account_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_account_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_account_id")))) {
		request.VirtualAccountID = interfaceToString(v)
	}
	return &request
}

func flattenLicensesSystemLicensingRegistrationV1Item(item *catalystcentersdkgo.ResponseLicensesSystemLicensingRegistrationV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
