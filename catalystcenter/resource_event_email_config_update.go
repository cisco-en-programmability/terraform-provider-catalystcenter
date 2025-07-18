package catalystcenter

import (
	"context"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceEventEmailConfigUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Event Management.

- Update Email Destination
`,

		CreateContext: resourceEventEmailConfigUpdateCreate,
		ReadContext:   resourceEventEmailConfigUpdateRead,
		DeleteContext: resourceEventEmailConfigUpdateDelete,
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

						"status_uri": &schema.Schema{
							Description: `Status Uri`,
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
						"email_config_id": &schema.Schema{
							Description: `Required only for update email configuration
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"from_email": &schema.Schema{
							Description: `From Email`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"primary_smt_p_config": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"host_name": &schema.Schema{
										Description: `Host Name`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Sensitive:   true,
										Computed:    true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"smtp_type": &schema.Schema{
										Description: `smtpType`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"user_name": &schema.Schema{
										Description: `User Name`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"secondary_smt_p_config": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"host_name": &schema.Schema{
										Description: `Host Name`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"password": &schema.Schema{
										Description: `Password`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Sensitive:   true,
										Computed:    true,
									},
									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"smtp_type": &schema.Schema{
										Description: `smtpType`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"user_name": &schema.Schema{
										Description: `User Name`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
								},
							},
						},
						"subject": &schema.Schema{
							Description: `Subject`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"to_email": &schema.Schema{
							Description: `To Email`,
							Type:        schema.TypeString,
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

func resourceEventEmailConfigUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestEventEmailConfigUpdateUpdateEmailDestination(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.EventManagement.UpdateEmailDestination(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing UpdateEmailDestination", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenEventManagementUpdateEmailDestinationItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateEmailDestination response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceEventEmailConfigUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceEventEmailConfigUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestEventEmailConfigUpdateUpdateEmailDestination(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementUpdateEmailDestination {
	request := catalystcentersdkgo.RequestEventManagementUpdateEmailDestination{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".email_config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".email_config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".email_config_id")))) {
		request.EmailConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".primary_smt_p_config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".primary_smt_p_config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".primary_smt_p_config")))) {
		request.PrimarySmtpConfig = expandRequestEventEmailConfigUpdateUpdateEmailDestinationPrimarySmtpConfig(ctx, key+".primary_smt_p_config.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secondary_smt_p_config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secondary_smt_p_config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secondary_smt_p_config")))) {
		request.SecondarySmtpConfig = expandRequestEventEmailConfigUpdateUpdateEmailDestinationSecondarySmtpConfig(ctx, key+".secondary_smt_p_config.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".from_email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".from_email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".from_email")))) {
		request.FromEmail = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".to_email")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".to_email")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".to_email")))) {
		request.ToEmail = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subject")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subject")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subject")))) {
		request.Subject = interfaceToString(v)
	}
	return &request
}

func expandRequestEventEmailConfigUpdateUpdateEmailDestinationPrimarySmtpConfig(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementUpdateEmailDestinationPrimarySmtpConfig {
	request := catalystcentersdkgo.RequestEventManagementUpdateEmailDestinationPrimarySmtpConfig{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smtp_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smtp_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smtp_type")))) {
		request.SmtpType = interfaceToString(v)
	}
	return &request
}

func expandRequestEventEmailConfigUpdateUpdateEmailDestinationSecondarySmtpConfig(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementUpdateEmailDestinationSecondarySmtpConfig {
	request := catalystcentersdkgo.RequestEventManagementUpdateEmailDestinationSecondarySmtpConfig{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".smtp_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".smtp_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".smtp_type")))) {
		request.SmtpType = interfaceToString(v)
	}
	return &request
}

func flattenEventManagementUpdateEmailDestinationItem(item *catalystcentersdkgo.ResponseEventManagementUpdateEmailDestination) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status_uri"] = item.StatusURI
	return []map[string]interface{}{
		respItem,
	}
}
