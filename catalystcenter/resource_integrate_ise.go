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
func resourceIntegrateIse() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on System Settings.

- API to accept Cisco ISE server certificate for Cisco ISE server integration. Use ‘Cisco ISE Server Integration Status’
Intent API to check the integration status. This data source action can be used to retry the failed integration.
`,

		CreateContext: resourceIntegrateIseCreate,
		ReadContext:   resourceIntegrateIseRead,
		DeleteContext: resourceIntegrateIseDelete,
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

						"object": &schema.Schema{
							Description: `object`,
							Type:        schema.TypeString,
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
						"id": &schema.Schema{
							Description: `id path parameter. Cisco ISE Server Identifier. Use 'Get Authentication and Policy Servers' intent API to find the identifier.
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"is_cert_accepted_by_user": &schema.Schema{
							Description: `Value true for accept, false for deny. Remove this field and send empty request payload ( {} ) to retry the failed integration
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceIntegrateIseCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vID := resourceItem["id"]

	vvID := vID.(string)
	request1 := expandRequestIntegrateIseAcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, err := client.SystemSettings.AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1(vvID, request1)

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	if err != nil {
		diags = append(diags, diagError(
			"Failure when setting AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1 response",
			err))
		return diags
	}
	if err := d.Set("item", response1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceIntegrateIseRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceIntegrateIseDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestIntegrateIseAcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSystemSettingsAcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1 {
	request := catalystcentersdkgo.RequestSystemSettingsAcceptCiscoIseServerCertificateForCiscoIseServerIntegrationV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_cert_accepted_by_user")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_cert_accepted_by_user")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_cert_accepted_by_user")))) {
		request.IsCertAcceptedByUser = interfaceToBoolPtr(v)
	}
	return &request
}
