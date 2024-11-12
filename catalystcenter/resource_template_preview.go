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
func resourceTemplatePreview() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Configuration Templates.

- API to preview a template.
`,

		CreateContext: resourceTemplatePreviewCreate,
		ReadContext:   resourceTemplatePreviewRead,
		DeleteContext: resourceTemplatePreviewDelete,
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

						"cli_preview": &schema.Schema{
							Description: `Generated template preview
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_id": &schema.Schema{
							Description: `UUID of device
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"validation_errors": &schema.Schema{
							Description: `Validation error in template content if any
`,
							Type:     schema.TypeString, //TEST,
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
						"device_id": &schema.Schema{
							Description: `UUID of device to get template preview
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"params": &schema.Schema{
							Description: `Params to render preview
`,
							Type:     schema.TypeString, //TEST,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"resource_params": &schema.Schema{
							Description: `Resource params to render preview
`,
							Type:     schema.TypeString, //TEST,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `UUID of template to get template preview
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

func resourceTemplatePreviewCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestTemplatePreviewPreviewTemplateV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.ConfigurationTemplates.PreviewTemplateV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing PreviewTemplateV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenConfigurationTemplatesPreviewTemplateV1Item(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting PreviewTemplateV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceTemplatePreviewRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceTemplatePreviewDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestTemplatePreviewPreviewTemplateV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesPreviewTemplateV1 {
	request := catalystcentersdkgo.RequestConfigurationTemplatesPreviewTemplateV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".params")))) {
		request.Params = expandRequestTemplatePreviewPreviewTemplateV1Params(ctx, key+".params.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".resource_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".resource_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".resource_params")))) {
		request.ResourceParams = expandRequestTemplatePreviewPreviewTemplateV1ResourceParams(ctx, key+".resource_params.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	return &request
}

func expandRequestTemplatePreviewPreviewTemplateV1Params(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesPreviewTemplateV1Params {
	var request catalystcentersdkgo.RequestConfigurationTemplatesPreviewTemplateV1Params
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestTemplatePreviewPreviewTemplateV1ResourceParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesPreviewTemplateV1ResourceParams {
	var request catalystcentersdkgo.RequestConfigurationTemplatesPreviewTemplateV1ResourceParams
	request = d.Get(fixKeyAccess(key))
	return &request
}

func flattenConfigurationTemplatesPreviewTemplateV1Item(item *catalystcentersdkgo.ResponseConfigurationTemplatesPreviewTemplateV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cli_preview"] = item.CliPreview
	respItem["device_id"] = item.DeviceID
	respItem["template_id"] = item.TemplateID
	respItem["validation_errors"] = flattenConfigurationTemplatesPreviewTemplateV1ItemValidationErrors(item.ValidationErrors)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesPreviewTemplateV1ItemValidationErrors(item *catalystcentersdkgo.ResponseConfigurationTemplatesPreviewTemplateV1ValidationErrors) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
