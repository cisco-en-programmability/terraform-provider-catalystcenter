package catalystcenter

import (
	"context"
	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceWirelessProvisionAccessPoint() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- Access Point Provision and ReProvision
`,

		CreateContext: resourceWirelessProvisionAccessPointCreate,
		ReadContext:   resourceWirelessProvisionAccessPointRead,
		DeleteContext: resourceWirelessProvisionAccessPointDelete,
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

						"execution_id": &schema.Schema{
							Description: `Execution Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"execution_status_url": &schema.Schema{
							Description: `Execution Status URL
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"message": &schema.Schema{
							Description: `Message
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
						"payload": &schema.Schema{
							Description: `Array of RequestWirelessAPProvisionConnectivity`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom_ap_group_name": &schema.Schema{
										Description: `Custom AP group name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"custom_flex_group_name": &schema.Schema{
										Description: `["Custom flex group name"]
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"device_name": &schema.Schema{
										Description: `Device name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"rf_profile": &schema.Schema{
										Description: `Radio frequency profile name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"site_name_hierarchy": &schema.Schema{
										Description: `Site name hierarchy(ex: Global/...)
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Description: `ApWirelessConfiguration
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

func resourceWirelessProvisionAccessPointCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vPersistbapioutput := resourceItem["persistbapioutput"]

	request1 := expandRequestWirelessProvisionAccessPointApProvisionConnectivity(ctx, "parameters.0", d)

	headerParams1 := catalystcentersdkgo.ApProvisionConnectivityHeaderParams{}

	headerParams1.Persistbapioutput = vPersistbapioutput.(string)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.ApProvisionConnectivity(request1, &headerParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ApProvisionConnectivity", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp1, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing APProvisionConnectivity", err,
				"Failure at APProvisionConnectivity execution", bapiError))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenWirelessApProvisionConnectivityItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ApProvisionConnectivity response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessProvisionAccessPointRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessProvisionAccessPointDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessProvisionAccessPointApProvisionConnectivity(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessApProvisionConnectivity {
	request := catalystcentersdkgo.RequestWirelessApProvisionConnectivity{}
	if v := expandRequestWirelessProvisionAccessPointApProvisionConnectivityItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestWirelessProvisionAccessPointApProvisionConnectivityItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemWirelessApProvisionConnectivity {
	request := []catalystcentersdkgo.RequestItemWirelessApProvisionConnectivity{}
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
		i := expandRequestWirelessProvisionAccessPointApProvisionConnectivityItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessProvisionAccessPointApProvisionConnectivityItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemWirelessApProvisionConnectivity {
	request := catalystcentersdkgo.RequestItemWirelessApProvisionConnectivity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_profile")))) {
		request.RfProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_name")))) {
		request.DeviceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_ap_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_ap_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_ap_group_name")))) {
		request.CustomApGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_flex_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_flex_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_flex_group_name")))) {
		request.CustomFlexGroupName = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	return &request
}

func flattenWirelessApProvisionConnectivityItem(item *catalystcentersdkgo.ResponseWirelessApProvisionConnectivity) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["execution_id"] = item.ExecutionID
	respItem["execution_status_url"] = item.ExecutionStatusURL
	respItem["message"] = item.Message
	return []map[string]interface{}{
		respItem,
	}
}
