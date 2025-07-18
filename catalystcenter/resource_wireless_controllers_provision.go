package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceWirelessControllersProvision() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- This data source action is used to provision wireless controller
`,

		CreateContext: resourceWirelessControllersProvisionCreate,
		ReadContext:   resourceWirelessControllersProvisionRead,
		DeleteContext: resourceWirelessControllersProvisionDelete,
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

						"task_id": &schema.Schema{
							Description: `Task ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Task URL
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
						"device_id": &schema.Schema{
							Description: `deviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"ap_authorization_list_name": &schema.Schema{
							Description: `AP Authorization List name. 'Obtain the AP Authorization List names by using the API call GET: /intent/api/v1/wirelessSettings/apAuthorizationLists. During re-provision, obtain the AP Authorization List configured for the given provisioned network device Id using the API call GET: /intent/api/v1/wireless/apAuthorizationLists/{networkDeviceId}'
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"authorize_mesh_and_non_mesh_access_points": &schema.Schema{
							Description: `True if AP Authorization List should  authorize against All Mesh/Non-Mesh APs, else false if AP Authorization List should only authorize against Mesh APs (Applicable only when Mesh is enabled on sites)
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							ForceNew:     true,
							Computed:     true,
						},
						"feature_templates_overriden_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"edit_feature_templates": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"additional_identifiers": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"site_uuid": &schema.Schema{
																Description: `Site UUID. This must be provided if **featureTemplateId** belongs to **Flex Configuration** feature template.
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"wlan_profile_name": &schema.Schema{
																Description: `WLAN Profile Name. This must be passed if **featureTemplateId** belongs to **Advanced SSID Configuration** Feature Template.
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
												"attributes": &schema.Schema{
													Description: `This dynamic map should contain attribute name and overridden value of respective Feature Template whose **featureTemplateId**. List of attributes applicable to given **featureTemplateId** can be retrieved from its GET API call /dna/intent/api/v1/featureTemplates/wireless/<featureTemplateName>/featureTemplateId
`,
													Type:     schema.TypeString, //TEST,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"excluded_attributes": &schema.Schema{
													Description: `List of attributes which will NOT be provisioned.
`,
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"feature_template_id": &schema.Schema{
													Description: `Feature Template ID
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
						"interfaces": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_gateway": &schema.Schema{
										Description: `Interface Gateway
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"interface_ipaddress": &schema.Schema{
										Description: `Interface IP Address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface Name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"interface_netmask_in_cid_r": &schema.Schema{
										Description: `Interface Netmask In CIDR, range is 1-30
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"lag_or_port_number": &schema.Schema{
										Description: `Lag Or Port Number
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"vlan_id": &schema.Schema{
										Description: `VLAN ID range is 1 - 4094
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"rolling_ap_upgrade": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_reboot_percentage": &schema.Schema{
										Description: `AP Reboot Percentage. Permissible values - 5, 15, 25
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"enable_rolling_ap_upgrade": &schema.Schema{
										Description: `True if Rolling AP Upgrade is enabled, else False
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
						"skip_ap_provision": &schema.Schema{
							Description: `True if Skip AP Provision is enabled, else False
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

func resourceWirelessControllersProvisionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vDeviceID := resourceItem["device_id"]

	vvDeviceID := vDeviceID.(string)
	request1 := expandRequestWirelessControllersProvisionWirelessControllerProvision(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Wireless.WirelessControllerProvision(vvDeviceID, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing WirelessControllerProvision", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing WirelessControllerProvision", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil || strings.Contains(restyResp3.String(), "<!doctype html>") {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing WirelessControllerProvision", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenWirelessWirelessControllerProvisionItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting WirelessControllerProvision response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessControllersProvisionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessControllersProvisionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessControllersProvisionWirelessControllerProvision(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessWirelessControllerProvision {
	request := catalystcentersdkgo.RequestWirelessWirelessControllerProvision{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interfaces")))) {
		request.Interfaces = expandRequestWirelessControllersProvisionWirelessControllerProvisionInterfacesArray(ctx, key+".interfaces", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip_ap_provision")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip_ap_provision")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip_ap_provision")))) {
		request.SkipApProvision = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rolling_ap_upgrade")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rolling_ap_upgrade")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rolling_ap_upgrade")))) {
		request.RollingApUpgrade = expandRequestWirelessControllersProvisionWirelessControllerProvisionRollingApUpgrade(ctx, key+".rolling_ap_upgrade.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_authorization_list_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_authorization_list_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_authorization_list_name")))) {
		request.ApAuthorizationListName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authorize_mesh_and_non_mesh_access_points")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authorize_mesh_and_non_mesh_access_points")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authorize_mesh_and_non_mesh_access_points")))) {
		request.AuthorizeMeshAndNonMeshAccessPoints = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".feature_templates_overriden_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".feature_templates_overriden_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".feature_templates_overriden_attributes")))) {
		request.FeatureTemplatesOverridenAttributes = expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributes(ctx, key+".feature_templates_overriden_attributes.0", d)
	}
	return &request
}

func expandRequestWirelessControllersProvisionWirelessControllerProvisionInterfacesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestWirelessWirelessControllerProvisionInterfaces {
	request := []catalystcentersdkgo.RequestWirelessWirelessControllerProvisionInterfaces{}
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
		i := expandRequestWirelessControllersProvisionWirelessControllerProvisionInterfaces(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessControllersProvisionWirelessControllerProvisionInterfaces(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessWirelessControllerProvisionInterfaces {
	request := catalystcentersdkgo.RequestWirelessWirelessControllerProvisionInterfaces{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_ipaddress")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_ipaddress")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_ipaddress")))) {
		request.InterfaceIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_netmask_in_cid_r")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_netmask_in_cid_r")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_netmask_in_cid_r")))) {
		request.InterfaceNetmaskInCIDR = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_gateway")))) {
		request.InterfaceGateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lag_or_port_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lag_or_port_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lag_or_port_number")))) {
		request.LagOrPortNumber = interfaceToIntPtr(v)
	}
	return &request
}

func expandRequestWirelessControllersProvisionWirelessControllerProvisionRollingApUpgrade(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessWirelessControllerProvisionRollingApUpgrade {
	request := catalystcentersdkgo.RequestWirelessWirelessControllerProvisionRollingApUpgrade{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".enable_rolling_ap_upgrade")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".enable_rolling_ap_upgrade")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".enable_rolling_ap_upgrade")))) {
		request.EnableRollingApUpgrade = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_reboot_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_reboot_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_reboot_percentage")))) {
		request.ApRebootPercentage = interfaceToIntPtr(v)
	}
	return &request
}

func expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributes {
	request := catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".edit_feature_templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".edit_feature_templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".edit_feature_templates")))) {
		request.EditFeatureTemplates = expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesArray(ctx, key+".edit_feature_templates", d)
	}
	return &request
}

func expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplates {
	request := []catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplates{}
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
		i := expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplates(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplates {
	request := catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".feature_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".feature_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".feature_template_id")))) {
		request.FeatureTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".attributes")))) {
		request.Attributes = expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesAttributes(ctx, key+".attributes.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".additional_identifiers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".additional_identifiers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".additional_identifiers")))) {
		request.AdditionalIDentifiers = expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesAdditionalIDentifiers(ctx, key+".additional_identifiers.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_attributes")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_attributes")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_attributes")))) {
		request.ExcludedAttributes = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesAttributes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesAttributes {
	var request catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesAttributes
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestWirelessControllersProvisionWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesAdditionalIDentifiers(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesAdditionalIDentifiers {
	request := catalystcentersdkgo.RequestWirelessWirelessControllerProvisionFeatureTemplatesOverridenAttributesEditFeatureTemplatesAdditionalIDentifiers{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlan_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlan_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlan_profile_name")))) {
		request.WLANProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_uuid")))) {
		request.SiteUUID = interfaceToString(v)
	}
	return &request
}

func flattenWirelessWirelessControllerProvisionItem(item *catalystcentersdkgo.ResponseWirelessWirelessControllerProvisionResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
