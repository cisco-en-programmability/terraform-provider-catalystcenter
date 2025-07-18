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
func resourceEventWebhookCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Event Management.

- Create Webhook Destination
`,

		CreateContext: resourceEventWebhookCreateCreate,
		ReadContext:   resourceEventWebhookCreateRead,
		DeleteContext: resourceEventWebhookCreateDelete,
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

						"api_status": &schema.Schema{
							Description: `Api Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"error_message": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"errors": &schema.Schema{
										Description: `Errors`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"status_message": &schema.Schema{
							Description: `Status Message`,
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
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"headers": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_value": &schema.Schema{
										Description: `Default Value`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"encrypt": &schema.Schema{
										Description: `Encrypt`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"name": &schema.Schema{
										Description: `Name`,
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
						"is_proxy_route": &schema.Schema{
							Description: `Is Proxy Route`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"method": &schema.Schema{
							Description: `Method`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"trust_cert": &schema.Schema{
							Description: `Trust Cert`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"url": &schema.Schema{
							Description: `Url`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"webhook_id": &schema.Schema{
							Description: `Required only for update webhook configuration
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

func resourceEventWebhookCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestEventWebhookCreateCreateWebhookDestination(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.EventManagement.CreateWebhookDestination(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing CreateWebhookDestination", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenEventManagementCreateWebhookDestinationItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreateWebhookDestination response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceEventWebhookCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceEventWebhookCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestEventWebhookCreateCreateWebhookDestination(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementCreateWebhookDestination {
	request := catalystcentersdkgo.RequestEventManagementCreateWebhookDestination{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".webhook_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".webhook_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".webhook_id")))) {
		request.WebhookID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".method")))) {
		request.Method = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".trust_cert")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".trust_cert")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".trust_cert")))) {
		request.TrustCert = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".headers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".headers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".headers")))) {
		request.Headers = expandRequestEventWebhookCreateCreateWebhookDestinationHeadersArray(ctx, key+".headers", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_proxy_route")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_proxy_route")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_proxy_route")))) {
		request.IsProxyRoute = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestEventWebhookCreateCreateWebhookDestinationHeadersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestEventManagementCreateWebhookDestinationHeaders {
	request := []catalystcentersdkgo.RequestEventManagementCreateWebhookDestinationHeaders{}
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
		i := expandRequestEventWebhookCreateCreateWebhookDestinationHeaders(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestEventWebhookCreateCreateWebhookDestinationHeaders(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementCreateWebhookDestinationHeaders {
	request := catalystcentersdkgo.RequestEventManagementCreateWebhookDestinationHeaders{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".encrypt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".encrypt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".encrypt")))) {
		request.Encrypt = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenEventManagementCreateWebhookDestinationItem(item *catalystcentersdkgo.ResponseEventManagementCreateWebhookDestination) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["error_message"] = flattenEventManagementCreateWebhookDestinationItemErrorMessage(item.ErrorMessage)
	respItem["api_status"] = item.APIStatus
	respItem["status_message"] = item.StatusMessage
	return []map[string]interface{}{
		respItem,
	}
}

func flattenEventManagementCreateWebhookDestinationItemErrorMessage(item *catalystcentersdkgo.ResponseEventManagementCreateWebhookDestinationErrorMessage) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["errors"] = item.Errors

	return []map[string]interface{}{
		respItem,
	}

}
