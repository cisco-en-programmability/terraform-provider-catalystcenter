package catalystcenter

import (
	"context"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaVirtualNetworkV1() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Add virtual network with scalable groups at global level

- Delete virtual network with scalable groups

- Update virtual network with scalable groups
`,

		CreateContext: resourceSdaVirtualNetworkV1Create,
		ReadContext:   resourceSdaVirtualNetworkV1Read,
		UpdateContext: resourceSdaVirtualNetworkV1Update,
		DeleteContext: resourceSdaVirtualNetworkV1Delete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"description": &schema.Schema{
							Description: `Virtual network info retrieved successfully
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"execution_id": &schema.Schema{
							Description: `Execution Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_guest_virtual_network": &schema.Schema{
							Description: `Guest Virtual Network
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"scalable_group_names": &schema.Schema{
							Description: `Scalable Group Names
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"status": &schema.Schema{
							Description: `Status
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"v_manage_vpn_id": &schema.Schema{
							Description: `vManage vpn id for SD-WAN
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_network_context_id": &schema.Schema{
							Description: `Virtual Network Context Id for Global Virtual Network
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name to be assigned at global level
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"is_guest_virtual_network": &schema.Schema{
							Description: `Guest Virtual Network enablement flag, default value is False.
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"scalable_group_names": &schema.Schema{
							Description: `Scalable Group to be associated to virtual network
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"v_manage_vpn_id": &schema.Schema{
							Description: `vManage vpn id for SD-WAN
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"virtual_network_name": &schema.Schema{
							Description: `Virtual Network Name to be assigned at global level
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaVirtualNetworkV1Create(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaVirtualNetworkV2AddVirtualNetworkWithScalableGroupsV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vVirtualNetworkName := resourceItem["virtual_network_name"]
	vvVirtualNetworkName := interfaceToString(vVirtualNetworkName)
	queryParamImport := catalystcentersdkgo.GetVirtualNetworkWithScalableGroupsV1QueryParams{}
	queryParamImport.VirtualNetworkName = vvVirtualNetworkName
	item2, _, err := client.Sda.GetVirtualNetworkWithScalableGroupsV1(&queryParamImport)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["virtual_network_name"] = item2.VirtualNetworkName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaVirtualNetworkV1Read(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddVirtualNetworkWithScalableGroupsV1(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddVirtualNetworkWithScalableGroupsV1", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddVirtualNetworkWithScalableGroupsV1", err))
		return diags
	}
	executionId := resp1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing AddVirtualNetworkWithScalableGroupsV1", err))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetVirtualNetworkWithScalableGroupsV1QueryParams{}
	queryParamValidate.VirtualNetworkName = vvVirtualNetworkName
	item3, _, err := client.Sda.GetVirtualNetworkWithScalableGroupsV1(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddVirtualNetworkWithScalableGroupsV1", err,
			"Failure at AddVirtualNetworkWithScalableGroupsV1, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["virtual_network_name"] = item3.VirtualNetworkName

	d.SetId(joinResourceID(resourceMap))
	return resourceSdaVirtualNetworkV1Read(ctx, d, m)
}

func resourceSdaVirtualNetworkV1Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vVirtualNetworkName := resourceMap["virtual_network_name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetVirtualNetworkWithScalableGroupsV1")
		queryParams1 := catalystcentersdkgo.GetVirtualNetworkWithScalableGroupsV1QueryParams{}

		queryParams1.VirtualNetworkName = vVirtualNetworkName

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.GetVirtualNetworkWithScalableGroupsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetVirtualNetworkWithScalableGroupsV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetVirtualNetworkWithScalableGroupsV1 response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSdaVirtualNetworkV1Update(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestSdaVirtualNetworkV2UpdateVirtualNetworkWithScalableGroupsV1(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Sda.UpdateVirtualNetworkWithScalableGroupsV1(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateVirtualNetworkWithScalableGroupsV1", err, restyResp1.String(),
					"Failure at UpdateVirtualNetworkWithScalableGroupsV1, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateVirtualNetworkWithScalableGroupsV1", err,
				"Failure at UpdateVirtualNetworkWithScalableGroupsV1, unexpected response", ""))
			return diags
		}

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
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
			for response2.Status == "IN_PROGRESS" {
				time.Sleep(10 * time.Second)
				response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
				if err != nil || response2 == nil {
					if restyResp2 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
					}
					diags = append(diags, diagErrorWithAlt(
						"Failure when executing GetExecutionByID", err,
						"Failure at GetExecutionByID, unexpected response", ""))
					return diags
				}
			}
			if response2.Status == "FAILURE" {
				log.Printf("[DEBUG] Error %s", response2.BapiError)
				diags = append(diags, diagError(
					"Failure when executing UpdateVirtualNetworkWithScalableGroupsV1", err))
				return diags
			}
		}

	}

	return resourceSdaVirtualNetworkV1Read(ctx, d, m)
}

func resourceSdaVirtualNetworkV1Delete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	queryParamDelete := catalystcentersdkgo.DeleteVirtualNetworkWithScalableGroupsV1QueryParams{}

	vvVirtualNetworkName := resourceMap["virtual_network_name"]
	queryParamDelete.VirtualNetworkName = vvVirtualNetworkName

	response1, restyResp1, err := client.Sda.DeleteVirtualNetworkWithScalableGroupsV1(&queryParamDelete)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteVirtualNetworkWithScalableGroupsV1", err, restyResp1.String(),
				"Failure at DeleteVirtualNetworkWithScalableGroupsV1, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteVirtualNetworkWithScalableGroupsV1", err,
			"Failure at DeleteVirtualNetworkWithScalableGroupsV1, unexpected response", ""))
		return diags
	}

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
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteVirtualNetworkWithScalableGroupsV1", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaVirtualNetworkV2AddVirtualNetworkWithScalableGroupsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaAddVirtualNetworkWithScalableGroupsV1 {
	request := catalystcentersdkgo.RequestSdaAddVirtualNetworkWithScalableGroupsV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_guest_virtual_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_guest_virtual_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_guest_virtual_network")))) {
		request.IsGuestVirtualNetwork = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_names")))) {
		request.ScalableGroupNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".v_manage_vpn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".v_manage_vpn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".v_manage_vpn_id")))) {
		request.VManageVpnID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaVirtualNetworkV2UpdateVirtualNetworkWithScalableGroupsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaUpdateVirtualNetworkWithScalableGroupsV1 {
	request := catalystcentersdkgo.RequestSdaUpdateVirtualNetworkWithScalableGroupsV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".virtual_network_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".virtual_network_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".virtual_network_name")))) {
		request.VirtualNetworkName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_guest_virtual_network")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_guest_virtual_network")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_guest_virtual_network")))) {
		request.IsGuestVirtualNetwork = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_names")))) {
		request.ScalableGroupNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".v_manage_vpn_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".v_manage_vpn_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".v_manage_vpn_id")))) {
		request.VManageVpnID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
