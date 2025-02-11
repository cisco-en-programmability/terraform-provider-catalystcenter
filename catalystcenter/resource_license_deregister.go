package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceLicenseDeregister() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Licenses.

- Deregisters the system with Cisco Smart Software Manager (CSSM)
`,

		CreateContext: resourceLicenseDeregisterCreate,
		ReadContext:   resourceLicenseDeregisterRead,
		DeleteContext: resourceLicenseDeregisterDelete,
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
					Schema: map[string]*schema.Schema{},
				},
			},
		},
	}
}

func resourceLicenseDeregisterCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	// has_unknown_response: None

	response1, restyResp1, err := client.Licenses.SmartLicensingDeregistrationV1()

	vItem1 := flattenLicensesSmartLicensingDeregistrationV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting SmartLicensingDeregistrationV1 response",
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
func resourceLicenseDeregisterRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceLicenseDeregisterDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenLicensesSmartLicensingDeregistrationV1Item(item *catalystcentersdkgo.ResponseLicensesSmartLicensingDeregistrationV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
