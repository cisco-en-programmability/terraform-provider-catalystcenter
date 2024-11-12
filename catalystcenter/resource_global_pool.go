package catalystcenter

import (
	"context"
	"fmt"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceGlobalPool() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Settings.

- API to update global pool

- API to create global pool.

- API to delete global IP pool.
`,

		CreateContext: resourceGlobalPoolCreate,
		ReadContext:   resourceGlobalPoolRead,
		UpdateContext: resourceGlobalPoolUpdate,
		DeleteContext: resourceGlobalPoolDelete,
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

						"available_ip_address_count": &schema.Schema{
							Description: `Available Ip Address Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"client_options": &schema.Schema{
							Description: `Client Options`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},
						"configure_external_dhcp": &schema.Schema{
							Description: `Configure External Dhcp`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"context": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"context_key": &schema.Schema{
										Description: `Context Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"context_value": &schema.Schema{
										Description: `Context Value`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"owner": &schema.Schema{
										Description: `Owner`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"create_time": &schema.Schema{
							Description: `Create Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"default_assigned_ip_address_count": &schema.Schema{
							Description: `Default Assigned Ip Address Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"dhcp_server_ips": &schema.Schema{
							Description: `Dhcp Server Ips`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"dns_server_ips": &schema.Schema{
							Description: `Dns Server Ips`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"gateways": &schema.Schema{
							Description: `Gateways`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"has_subpools": &schema.Schema{
							Description: `Has Subpools`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ip_pool_cidr": &schema.Schema{
							Description: `Ip Pool Cidr`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ip_pool_name": &schema.Schema{
							Description: `Ip Pool Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ip_pool_type": &schema.Schema{
							Description: `Ip Pool Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ipv6": &schema.Schema{
							Description: `Ipv6`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Last Update Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"overlapping": &schema.Schema{
							Description: `Overlapping`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"owner": &schema.Schema{
							Description: `Owner`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"parent_uuid": &schema.Schema{
							Description: `Parent Uuid`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"shared": &schema.Schema{
							Description: `Shared`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"total_assignable_ip_address_count": &schema.Schema{
							Description: `Total Assignable Ip Address Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"total_ip_address_count": &schema.Schema{
							Description: `Total Ip Address Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"unavailable_ip_address_count": &schema.Schema{
							Description: `Unavailable Ip Address Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
						"used_ip_address_count": &schema.Schema{
							Description: `Used Ip Address Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"used_percentage": &schema.Schema{
							Description: `Used Percentage`,
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
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `id path parameter. global pool id
`,
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"settings": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ippool": &schema.Schema{
										Type:     schema.TypeList,
										Required: true,
										MaxItems: 1,
										MinItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ip_address_space": &schema.Schema{
													Description: `Ip Address Space`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"dhcp_server_ips": &schema.Schema{
													Description: `Dhcp Server Ips`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"dns_server_ips": &schema.Schema{
													Description: `Dns Server Ips`,
													Type:        schema.TypeList,
													Optional:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"gateway": &schema.Schema{
													Description: `Gateway`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"ip_pool_cidr": &schema.Schema{
													Description: `Ip Pool Cidr`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"ip_pool_name": &schema.Schema{
													Description: `Ip Pool Name`,
													Type:        schema.TypeString,
													Required:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
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

func resourceGlobalPoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG] call resourceGlobalPoolCreate")
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	request1 := expandRequestGlobalPoolCreateGlobalPoolV1(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	resourceItem := *getResourceItem(d.Get("parameters"))
	vvID := interfaceToString(resourceItem["id"])
	vvIpPoolName := ""
	if _, ok := d.GetOk("parameters.0"); ok {
		if _, ok := d.GetOk("parameters.0.settings"); ok {
			if _, ok := d.GetOk("parameters.0.settings.0"); ok {
				if _, ok := d.GetOk("parameters.0.settings.0.ippool"); ok {
					if _, ok := d.GetOk("parameters.0.settings.0.ippool.0"); ok {
						if v, ok := d.GetOk("parameters.0.settings.0.ippool.0.ip_pool_name"); ok {
							vvIpPoolName = interfaceToString(v)
						}
					}
				}
			}
		}
	}
	queryParams1 := catalystcentersdkgo.GetGlobalPoolV1QueryParams{}

	response1, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vvID, vvIpPoolName)
	if response1 != nil {
		log.Printf("[DEBUG] searchNetworkSettingsGetGlobalPool result %v", responseInterfaceToString(*response1))
	}
	if err == nil && (response1 != nil && len(*response1) > 0) {
		resourceMap := make(map[string]string)
		resourceMap["ip_pool_name"] = vvIpPoolName
		resourceMap["id"] = vvID
		d.SetId(joinResourceID(resourceMap))
		return resourceGlobalPoolRead(ctx, d, m)
	}

	resp1, restyResp1, err := client.NetworkSettings.CreateGlobalPool(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateGlobalPool", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateGlobalPool", err))
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
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
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
		if statusIsFailure(response2.Status) {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing CreateGlobalPool", err,
				"Failure at CreateGlobalPool execution", bapiError))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["ip_pool_name"] = vvIpPoolName
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalPoolRead(ctx, d, m)
}

func resourceGlobalPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG] call resourceGlobalPoolRead")

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIpPoolName := resourceMap["ip_pool_name"]
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalPool")
		queryParams1 := catalystcentersdkgo.GetGlobalPoolV1QueryParams{}

		response1, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vID, vIpPoolName)
		if err != nil || response1 == nil || len(*response1) <= 0 {
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsGetGlobalPoolV1Items(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalPool search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceGlobalPoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG] call resourceGlobalPoolUpdate")
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIpPoolName := resourceMap["ip_pool_name"]
	vID := resourceMap["id"]

	queryParams1 := catalystcentersdkgo.GetGlobalPoolV1QueryParams{}
	item, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vID, vIpPoolName)
	if err != nil || item == nil || len(*item) <= 0 {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalPool", err,
			"Failure at GetGlobalPool, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	if d.HasChange("parameters") {
		request1 := expandRequestGlobalPoolUpdateGlobalPoolV1(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		if item != nil && len(*item) > 0 {
			if request1 != nil && request1.Settings != nil && request1.Settings.IPpool != nil && len(*request1.Settings.IPpool) > 0 {
				found := *item
				req := *request1.Settings.IPpool
				req[0].ID = found[0].ID
			}
		}
		response1, restyResp1, err := client.NetworkSettings.UpdateGlobalPool(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateGlobalPool", err, restyResp1.String(),
					"Failure at UpdateGlobalPool, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGlobalPool", err,
				"Failure at UpdateGlobalPool, unexpected response", ""))
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
					"Failure when executing GetBusinessAPIExecutionDetails", err,
					"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
				return diags
			}
			for statusIsPending(response2.Status) {
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
			if statusIsFailure(response2.Status) {
				bapiError := response2.BapiError
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing UpdateGlobalPool", err,
					"Failure at UpdateGlobalPool execution", bapiError))
				return diags
			}
		}

	}

	return resourceGlobalPoolRead(ctx, d, m)
}

func resourceGlobalPoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Println("[DEBUG] call resourceGlobalPoolDelete")

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vIpPoolName := resourceMap["ip_pool_name"]
	vID := resourceMap["id"]

	queryParams1 := catalystcentersdkgo.GetGlobalPoolV1QueryParams{}
	item, err := searchNetworkSettingsGetGlobalPool(m, queryParams1, vID, vIpPoolName)
	if err != nil {
		return diags
	}
	if item == nil || len(*item) == 0 {
		return diags
	}
	if vID == "" && item != nil && len(*item) > 0 {
		vID = (*item)[0].ID
	}

	response1, restyResp1, err := client.NetworkSettings.DeleteGlobalIPPool(vID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteGlobalIPPool", err, restyResp1.String(),
				"Failure at DeleteGlobalIPPool, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteGlobalIPPool", err,
			"Failure at DeleteGlobalIPPool, unexpected response", ""))
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
				"Failure when executing GetBusinessAPIExecutionDetails", err,
				"Failure at GetBusinessAPIExecutionDetails, unexpected response", ""))
			return diags
		}
		for statusIsPending(response2.Status) {
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
		if statusIsFailure(response2.Status) {
			bapiError := response2.BapiError
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateGlobalPool", err,
				"Failure at DeleteGlobalIPPool execution", bapiError))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestGlobalPoolCreateGlobalPoolV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateGlobalPoolV1 {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateGlobalPoolV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".settings")))) {
		request.Settings = expandRequestGlobalPoolCreateGlobalPoolV1Settings(ctx, key+".settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolV1Settings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateGlobalPoolV1Settings {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateGlobalPoolV1Settings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ippool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ippool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ippool")))) {
		request.IPpool = expandRequestGlobalPoolCreateGlobalPoolV1SettingsIPpoolArray(ctx, key+".ippool", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolV1SettingsIPpoolArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestNetworkSettingsCreateGlobalPoolV1SettingsIPpool {
	request := []catalystcentersdkgo.RequestNetworkSettingsCreateGlobalPoolV1SettingsIPpool{}
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
		i := expandRequestGlobalPoolCreateGlobalPoolV1SettingsIPpool(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalPoolCreateGlobalPoolV1SettingsIPpool(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsCreateGlobalPoolV1SettingsIPpool {
	request := catalystcentersdkgo.RequestNetworkSettingsCreateGlobalPoolV1SettingsIPpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_cidr")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_cidr")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_cidr")))) {
		request.IPPoolCidr = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server_ips")))) {
		request.DhcpServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server_ips")))) {
		request.DNSServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address_space")))) {
		request.IPAddressSpace = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsUpdateGlobalPoolV1 {
	request := catalystcentersdkgo.RequestNetworkSettingsUpdateGlobalPoolV1{}
	request.Settings = expandRequestGlobalPoolUpdateGlobalPoolV1Settings(ctx, key, d)
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolV1Settings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsUpdateGlobalPoolV1Settings {
	request := catalystcentersdkgo.RequestNetworkSettingsUpdateGlobalPoolV1Settings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ippool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ippool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ippool")))) {
		request.IPpool = expandRequestGlobalPoolUpdateGlobalPoolV1SettingsIPpoolArray(ctx, key+".ippool", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolV1SettingsIPpoolArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestNetworkSettingsUpdateGlobalPoolV1SettingsIPpool {
	request := []catalystcentersdkgo.RequestNetworkSettingsUpdateGlobalPoolV1SettingsIPpool{}
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
		i := expandRequestGlobalPoolUpdateGlobalPoolV1SettingsIPpool(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestGlobalPoolUpdateGlobalPoolV1SettingsIPpool(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsUpdateGlobalPoolV1SettingsIPpool {
	request := catalystcentersdkgo.RequestNetworkSettingsUpdateGlobalPoolV1SettingsIPpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_pool_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_pool_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_pool_name")))) {
		request.IPPoolName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dhcp_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dhcp_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dhcp_server_ips")))) {
		request.DhcpServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dns_server_ips")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dns_server_ips")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dns_server_ips")))) {
		request.DNSServerIPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchNetworkSettingsGetGlobalPool(m interface{}, queryParams catalystcentersdkgo.GetGlobalPoolV1QueryParams, vID string, vIPPoolName string) (*[]catalystcentersdkgo.ResponseNetworkSettingsGetGlobalPoolV1Response, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItems []catalystcentersdkgo.ResponseNetworkSettingsGetGlobalPoolV1Response
	offset := 1
	queryParams.Offset = float64(offset)

	nResponse, _, err := client.NetworkSettings.GetGlobalPool(&queryParams)
	if err != nil {
		log.Printf("[DEBUG] GetGlobalPool error %s", err.Error())
		return nil, err
	}
	if nResponse == nil {
		return nil, err
	}
	if nResponse.Response == nil {
		return nil, err
	}
	maxPageSize := len(*nResponse.Response)
	//maxPageSize := 10
	for nResponse != nil && nResponse.Response != nil && len(*nResponse.Response) > 0 {
		for _, item := range *nResponse.Response {
			if vIPPoolName == item.IPPoolName {
				foundItems = append(foundItems, item)
				return &foundItems, err
			}
			if vID == item.ID {
				foundItems = append(foundItems, item)
				return &foundItems, err
			}
		}

		queryParams.Limit = float64(maxPageSize)
		offset += maxPageSize
		queryParams.Offset = float64(offset)

		nResponse, _, err = client.NetworkSettings.GetGlobalPool(&queryParams)
		if nResponse == nil || nResponse.Response == nil {
			break
		}
	}
	return &foundItems, err
}
