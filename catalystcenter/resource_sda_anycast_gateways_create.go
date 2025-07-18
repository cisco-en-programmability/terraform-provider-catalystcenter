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
func resourceSdaAnycastGatewaysCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on SDA.

- Adds anycast gateways based on user input.
`,

		CreateContext: resourceSdaAnycastGatewaysCreateCreate,
		ReadContext:   resourceSdaAnycastGatewaysCreateRead,
		DeleteContext: resourceSdaAnycastGatewaysCreateDelete,
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
							Description: `ID of the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Task status lookup url.
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
							Description: `Array of RequestSdaAddAnycastGateways`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auto_generate_vlan_name": &schema.Schema{
										Description: `This field cannot be true when vlanName is provided. the vlanName will be generated as "{ipPoolGroupV4Cidr}-{virtualNetworkName}" for non-critical VLANs. for critical VLANs with DATA trafficType, vlanName will be "CRITICAL_VLAN". for critical VLANs with VOICE trafficType, vlanName will be "VOICE_VLAN".
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric this anycast gateway is to be assigned to.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ip_pool_name": &schema.Schema{
										Description: `Name of the IP pool associated with the anycast gateway.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"is_critical_pool": &schema.Schema{
										Description: `Enable/disable critical VLAN. if true, autoGenerateVlanName must also be true. (isCriticalPool is not applicable to INFRA_VN).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"is_group_based_policy_enforcement_enabled": &schema.Schema{
										Description: `Enable/disable Group-Based Policy Enforcement (applicable only to INFRA_VN; defaults to false).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"is_intra_subnet_routing_enabled": &schema.Schema{
										Description: `Enable/disable Intra-Subnet Routing (not applicable to INFRA_VN).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"is_ip_directed_broadcast": &schema.Schema{
										Description: `Enable/disable IP-directed broadcast (not applicable to INFRA_VN).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"is_layer2_flooding_enabled": &schema.Schema{
										Description: `Enable/disable layer 2 flooding (not applicable to INFRA_VN).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"is_multiple_ip_to_mac_addresses": &schema.Schema{
										Description: `Enable/disable multiple IP-to-MAC Addresses (Wireless Bridged-Network Virtual Machine; not applicable to INFRA_VN).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"is_supplicant_based_extended_node_onboarding": &schema.Schema{
										Description: `Enable/disable Supplicant-Based Extended Node Onboarding (applicable only to INFRA_VN).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"is_wireless_pool": &schema.Schema{
										Description: `Enable/disable fabric-enabled wireless (not applicable to INFRA_VN).
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"pool_type": &schema.Schema{
										Description: `The pool type of the anycast gateway (required for & applicable only to INFRA_VN).
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"security_group_name": &schema.Schema{
										Description: `Name of the associated Security Group (not applicable to INFRA_VN).
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"tcp_mss_adjustment": &schema.Schema{
										Description: `TCP maximum segment size adjustment.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"traffic_type": &schema.Schema{
										Description: `The type of traffic the anycast gateway serves.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"virtual_network_name": &schema.Schema{
										Description: `Name of the layer 3 virtual network associated with the anycast gateway. the virtual network must have already been added to the site before creating an anycast gateway with it.
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"vlan_id": &schema.Schema{
										Description: `ID of the VLAN of the anycast gateway. allowed VLAN range is 2-4093 except for reserved VLANs 1002-1005, 2046, and 4094. if deploying an anycast gateway on a fabric zone, this vlanId must match the vlanId of the corresponding anycast gateway on the fabric site.
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"vlan_name": &schema.Schema{
										Description: `Name of the VLAN of the anycast gateway.
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

func resourceSdaAnycastGatewaysCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSdaAnycastGatewaysCreateAddAnycastGateways(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sda.AddAnycastGateways(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing AddAnycastGateways", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddAnycastGateways", err))
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
				"Failure when executing AddAnycastGateways", err1))
			return diags
		}
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vItem1 := flattenSdaAddAnycastGatewaysItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AddAnycastGateways response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceSdaAnycastGatewaysCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSdaAnycastGatewaysCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestSdaAnycastGatewaysCreateAddAnycastGateways(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaAddAnycastGateways {
	request := catalystcentersdkgo.RequestSdaAddAnycastGateways{}
	if v := expandRequestSdaAnycastGatewaysCreateAddAnycastGatewaysItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestSdaAnycastGatewaysCreateAddAnycastGatewaysItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSdaAddAnycastGateways {
	request := []catalystcentersdkgo.RequestItemSdaAddAnycastGateways{}
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
		i := expandRequestSdaAnycastGatewaysCreateAddAnycastGatewaysItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSdaAnycastGatewaysCreateAddAnycastGatewaysItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSdaAddAnycastGateways {
	request := catalystcentersdkgo.RequestItemSdaAddAnycastGateways{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tcp_mss_adjustment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tcp_mss_adjustment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tcp_mss_adjustment")))) {
		request.TCPMssAdjustment = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_name")))) {
		request.VLANName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_type")))) {
		request.TrafficType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pool_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pool_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pool_type")))) {
		request.PoolType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".security_group_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".security_group_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".security_group_name")))) {
		request.SecurityGroupName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_critical_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_critical_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_critical_pool")))) {
		request.IsCriticalPool = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_layer2_flooding_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_layer2_flooding_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_layer2_flooding_enabled")))) {
		request.IsLayer2FloodingEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer2_flooding_address_assignment")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer2_flooding_address_assignment")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer2_flooding_address_assignment")))) {
		request.Layer2FloodingAddressAssignment = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer2_flooding_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer2_flooding_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer2_flooding_address")))) {
		request.Layer2FloodingAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_wireless_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_wireless_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_wireless_pool")))) {
		request.IsWirelessPool = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_wireless_flooding_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_wireless_flooding_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_wireless_flooding_enabled")))) {
		request.IsWirelessFloodingEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_resource_guard_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_resource_guard_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_resource_guard_enabled")))) {
		request.IsResourceGuardEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_ip_directed_broadcast")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_ip_directed_broadcast")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_ip_directed_broadcast")))) {
		request.IsIPDirectedBroadcast = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_intra_subnet_routing_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_intra_subnet_routing_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_intra_subnet_routing_enabled")))) {
		request.IsIntraSubnetRoutingEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_multiple_ip_to_mac_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_multiple_ip_to_mac_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_multiple_ip_to_mac_addresses")))) {
		request.IsMultipleIPToMacAddresses = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_supplicant_based_extended_node_onboarding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_supplicant_based_extended_node_onboarding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_supplicant_based_extended_node_onboarding")))) {
		request.IsSupplicantBasedExtendedNodeOnboarding = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_group_based_policy_enforcement_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_group_based_policy_enforcement_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_group_based_policy_enforcement_enabled")))) {
		request.IsGroupBasedPolicyEnforcementEnabled = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auto_generate_vlan_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auto_generate_vlan_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auto_generate_vlan_name")))) {
		request.AutoGenerateVLANName = interfaceToBoolPtr(v)
	}
	return &request
}

func flattenSdaAddAnycastGatewaysItem(item *catalystcentersdkgo.ResponseSdaAddAnycastGatewaysResponse) []map[string]interface{} {
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
