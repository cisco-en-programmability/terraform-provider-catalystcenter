package catalystcenter

import (
	"context"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePnpDeviceAuthorize() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Cisco DNA Center System.

- Authorizes one of more devices. A device can only be authorized if Authorization is set in Device Settings.
`,

		CreateContext: resourcePnpDeviceAuthorizeCreate,
		ReadContext:   resourcePnpDeviceAuthorizeRead,
		DeleteContext: resourcePnpDeviceAuthorizeDelete,
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

						"json_array_response": &schema.Schema{
							Description: `Json Array Response`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"json_response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"empty": &schema.Schema{
										Description: `Empty`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"message": &schema.Schema{
							Description: `Message`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"status_code": &schema.Schema{
							Description: `Status Code`,
							Type:        schema.TypeFloat,
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
						"device_id_list": &schema.Schema{
							Description: `Device Id List`,
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

func resourcePnpDeviceAuthorizeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestPnpDeviceAuthorizeAuthorizeDeviceV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.DeviceOnboardingPnp.AuthorizeDeviceV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing AuthorizeDeviceV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDeviceOnboardingPnpAuthorizeDeviceV1Item(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AuthorizeDeviceV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

	//REVIEW: '- Analizar como se puede comprobar la ejecucion.'

}
func resourcePnpDeviceAuthorizeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpDeviceAuthorizeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestPnpDeviceAuthorizeAuthorizeDeviceV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpAuthorizeDeviceV1 {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpAuthorizeDeviceV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id_list")))) {
		request.DeviceIDList = interfaceToSliceString(v)
	}
	return &request
}

func flattenDeviceOnboardingPnpAuthorizeDeviceV1Item(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpAuthorizeDeviceV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["json_response"] = flattenDeviceOnboardingPnpAuthorizeDeviceV1ItemJSONResponse(item.JSONResponse)
	respItem["message"] = item.Message
	respItem["status_code"] = item.StatusCode
	respItem["json_array_response"] = item.JSONArrayResponse
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpAuthorizeDeviceV1ItemJSONResponse(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpAuthorizeDeviceV1JSONResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["empty"] = boolPtrToString(item.Empty)

	return []map[string]interface{}{
		respItem,
	}

}
