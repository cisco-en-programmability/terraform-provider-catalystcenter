package catalystcenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceReserveIPSubpool() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Network Settings.

- API to delete the reserved ip subpool

- API to reserve an ip subpool from the global pool

- API to update ip subpool from the global pool
`,

		CreateContext: resourceReserveIPSubpoolCreate,
		ReadContext:   resourceReserveIPSubpoolRead,
		UpdateContext: resourceReserveIPSubpoolUpdate,
		DeleteContext: resourceReserveIPSubpoolDelete,
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

						"group_name": &schema.Schema{
							Description: `Group Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"group_owner": &schema.Schema{
							Description: `Group Owner`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"ip_pools": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

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
									"group_uuid": &schema.Schema{
										Description: `Group Uuid`,
										Type:        schema.TypeString,
										Computed:    true,
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
									"total_ip_address_count": &schema.Schema{
										Description: `Total Ip Address Count`,
										Type:        schema.TypeInt,
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
						"site_hierarchy": &schema.Schema{
							Description: `Site Hierarchy`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
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

						"id": &schema.Schema{
							Description: `id path parameter. Id of reserve ip subpool to be deleted.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"ipv4_dhcp_servers": &schema.Schema{
							Description: `IPv4 input for dhcp server ip example: ["1.1.1.1"]
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4_dns_servers": &schema.Schema{
							Description: `IPv4 input for dns server ip example: ["4.4.4.4"]
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv4_gate_way": &schema.Schema{
							Description: `Gateway ip address details, example: 175.175.0.1
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_global_pool": &schema.Schema{
							Description: `IP v4 Global pool address with cidr, example: 175.175.0.0/16
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_prefix": &schema.Schema{
							Description: `IPv4 prefix value is true, the ip4 prefix length input field is enabled , if it is false ipv4 total Host input is enable
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"ipv4_prefix_length": &schema.Schema{
							Description: `The ipv4 prefix length is required when ipv4prefix value is true.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ipv4_subnet": &schema.Schema{
							Description: `IPv4 Subnet address, example: 175.175.0.0. Either ipv4Subnet or ipv4TotalHost needs to be passed if creating IPv4 subpool.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv4_total_host": &schema.Schema{
							Description: `IPv4 total host is required when ipv4prefix value is false.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ipv6_address_space": &schema.Schema{
							Description: `If the value is omitted or false only ipv4 input are required, otherwise both ipv6 and ipv4 are required
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"ipv6_dhcp_servers": &schema.Schema{
							Description: `IPv6 format dhcp server as input example : ["2001:db8::1234"]
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv6_dns_servers": &schema.Schema{
							Description: `IPv6 format dns server input example: ["2001:db8::1234"]
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"ipv6_gate_way": &schema.Schema{
							Description: `Gateway ip address details, example: 2001:db8:85a3:0:100::1
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_global_pool": &schema.Schema{
							Description: `IPv6 Global pool address with cidr this is required when Ipv6AddressSpace value is true, example: 2001:db8:85a3::/64
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_prefix": &schema.Schema{
							Description: `Ipv6 prefix value is true, the ip6 prefix length input field is enabled , if it is false ipv6 total Host input is enable
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"ipv6_prefix_length": &schema.Schema{
							Description: `IPv6 prefix length is required when the ipv6prefix value is true
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ipv6_subnet": &schema.Schema{
							Description: `IPv6 Subnet address, example :2001:db8:85a3:0:100::. Either ipv6Subnet or ipv6TotalHost needs to be passed if creating IPv6 subpool.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ipv6_total_host": &schema.Schema{
							Description: `IPv6 total host is required when ipv6prefix value is false.
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of the reserve ip sub pool
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `siteId path parameter. Site id to reserve the ip sub pool.
`,
							Type:     schema.TypeString,
							Optional: true,
						},
						"slaac_support": &schema.Schema{
							Description: `Slaac Support`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"type": &schema.Schema{
							Description: `Type of the reserve ip sub pool
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

func resourceReserveIPSubpoolCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestReserveIPSubpoolReserveIPSubpool(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)

	queryParams1 := catalystcentersdkgo.GetReserveIPSubpoolQueryParams{}

	queryParams1.SiteID = vvSiteID

	response1, err := searchNetworkSettingsGetReserveIPSubpool(m, queryParams1, vvName)

	if err != nil || response1 != nil {
		resourceMap := make(map[string]string)
		resourceMap["site_id"] = response1.SiteID
		resourceMap["name"] = vvName
		d.SetId(joinResourceID(resourceMap))
		return resourceReserveIPSubpoolRead(ctx, d, m)
	}

	resp1, restyResp1, err := client.NetworkSettings.ReserveIPSubpool(vvSiteID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ReserveIPSubpool", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ReserveIPSubpool", err))
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
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing ReserveIPSubpool", err))
			return diags
		} else {
			// if strings.Contains(response2.BapiSyncResponse, "FailureReason:") {
			err1 := errors.New(response2.BapiError)
			// log.Printf("[DEBUG] Error %s", response2.BapiSyncResponse)
			diags = append(diags, diagError(
				"Failure when executing ReserveIPSubpool", err1))
			return diags
			// }
		}
	}

	resourceMap := make(map[string]string)
	resourceMap["site_id"] = vvSiteID
	resourceMap["name"] = vvName
	d.SetId(joinResourceID(resourceMap))
	return resourceReserveIPSubpoolRead(ctx, d, m)
}

func resourceReserveIPSubpoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID, okSiteID := resourceMap["site_id"]
	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetReserveIPSubpool")
		queryParams1 := catalystcentersdkgo.GetReserveIPSubpoolQueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID
		}

		response1, err := searchNetworkSettingsGetReserveIPSubpool(m, queryParams1, vName)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when setting searchNetworkSettingsGetReserveIPSubpool search response",
				err))
			return diags
		}
		if response1 == nil {
			//			log.Print("[DEBUG] Error response")
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		items := []catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponse{
			*response1,
		}
		vItem1 := flattenNetworkSettingsGetReserveIPSubpoolItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetReserveIPSubpool search response",
				err))
			return diags
		}
		request1 := expandRequestReserveIPSubpoolReserveIPSubpool(ctx, "parameters.0", d)
		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		updatedParameters := updateReserveIpPoolParameters(request1, response1)

		vParameters := flattenNetworkSettingsGetReserveIPSubpoolParameters(updatedParameters)
		vParameters[0]["site_id"] = vSiteID
		if err := d.Set("parameters", vParameters); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetReserveIPSubpool search response",
				err))
			return diags
		}

	}
	return diags
}

func flattenNetworkSettingsGetReserveIPSubpoolParameters(item *catalystcentersdkgo.RequestNetworkSettingsReserveIPSubpool) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["type"] = item.Type
	respItem["ipv6AddressSpace"] = item.IPv6AddressSpace
	respItem["ipv4GlobalPool"] = item.IPv4GlobalPool
	respItem["ipv4Prefix"] = item.IPv4Prefix
	respItem["ipv4PrefixLength"] = item.IPv4PrefixLength
	respItem["ipv4Subnet"] = item.IPv4Subnet
	respItem["ipv4GateWay"] = item.IPv4GateWay
	respItem["ipv4DhcpServers"] = item.IPv4DhcpServers
	respItem["ipv4DnsServers"] = item.IPv4DNSServers
	respItem["ipv6GlobalPool"] = item.IPv6GlobalPool
	respItem["ipv6Prefix"] = item.IPv6Prefix
	respItem["ipv6PrefixLength"] = item.IPv6PrefixLength
	respItem["ipv6Subnet"] = item.IPv6Subnet
	respItem["ipv6GateWay"] = item.IPv6GateWay
	respItem["ipv6DhcpServers"] = item.IPv6DhcpServers
	respItem["ipv6DnsServers"] = item.IPv6DNSServers
	respItem["ipv4TotalHost"] = item.IPv4TotalHost
	respItem["ipv6TotalHost"] = item.IPv6TotalHost
	respItem["slaacSupport"] = item.SLAacSupport
	return []map[string]interface{}{
		respItem,
	}
}

func resourceReserveIPSubpoolUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]
	vName := resourceMap["name"]

	queryParams1 := catalystcentersdkgo.GetReserveIPSubpoolQueryParams{}
	queryParams1.SiteID = vSiteID
	item, err := searchNetworkSettingsGetReserveIPSubpool(m, queryParams1, vName)
	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetReserveIPSubpool", err,
			"Failure at GetReserveIPSubpool, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vSiteID)
		request1 := expandRequestReserveIPSubpoolUpdateReserveIPSubpool(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		queryParams2 := catalystcentersdkgo.UpdateReserveIPSubpoolQueryParams{}
		queryParams2.ID = item.ID
		response1, restyResp1, err := client.NetworkSettings.UpdateReserveIPSubpool(vSiteID, request1, &queryParams2)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateReserveIPSubpool", err, restyResp1.String(),
					"Failure at UpdateReserveIPSubpool, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateReserveIPSubpool", err,
				"Failure at UpdateReserveIPSubpool, unexpected response", ""))
			return diags
		}
		executionId := response1.ExecutionID
		log.Printf("[DEBUG] ExecutionID => %s", executionId)
		if executionId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp1 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
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
				log.Printf("[DEBUG] Error %s", response2.BapiError)
				diags = append(diags, diagError(
					"Failure when executing UpdateReserveIPSubpool", err))
				return diags
			}
		}

	}

	return resourceReserveIPSubpoolRead(ctx, d, m)
}

func resourceReserveIPSubpoolDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSiteID := resourceMap["site_id"]
	vName := resourceMap["name"]

	queryParams1 := catalystcentersdkgo.GetReserveIPSubpoolQueryParams{}
	queryParams1.SiteID = vSiteID
	item, err := searchNetworkSettingsGetReserveIPSubpool(m, queryParams1, vName)
	if err != nil || item == nil {
		return diags
	}

	var vvID string
	// REVIEW: Add getAllItems and search function to get missing params

	vvID = item.ID

	response1, restyResp1, err := client.NetworkSettings.ReleaseReserveIPSubpool(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing ReleaseReserveIPSubpool", err, restyResp1.String(),
				"Failure at ReleaseReserveIPSubpool, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ReleaseReserveIPSubpool", err,
			"Failure at ReleaseReserveIPSubpool, unexpected response", ""))
		return diags
	}
	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp1, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
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
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing ReleaseReserveIPSubpool", err))
			return diags
		}
	}
	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestReserveIPSubpoolReserveIPSubpool(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsReserveIPSubpool {
	request := catalystcentersdkgo.RequestNetworkSettingsReserveIPSubpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_space")))) {
		request.IPv6AddressSpace = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_global_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_global_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_global_pool")))) {
		request.IPv4GlobalPool = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_prefix")))) {
		request.IPv4Prefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_prefix_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_prefix_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_prefix_length")))) {
		request.IPv4PrefixLength = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_subnet")))) {
		request.IPv4Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_gate_way")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_gate_way")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_gate_way")))) {
		request.IPv4GateWay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_dhcp_servers")))) {
		request.IPv4DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_dns_servers")))) {
		request.IPv4DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_global_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_global_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_global_pool")))) {
		request.IPv6GlobalPool = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_prefix")))) {
		request.IPv6Prefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_prefix_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_prefix_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_prefix_length")))) {
		request.IPv6PrefixLength = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_subnet")))) {
		request.IPv6Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_gate_way")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_gate_way")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_gate_way")))) {
		request.IPv6GateWay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_dhcp_servers")))) {
		request.IPv6DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_dns_servers")))) {
		request.IPv6DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_total_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_total_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_total_host")))) {
		request.IPv4TotalHost = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_total_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_total_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_total_host")))) {
		request.IPv6TotalHost = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".slaac_support")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".slaac_support")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".slaac_support")))) {
		request.SLAacSupport = interfaceToBoolPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestReserveIPSubpoolUpdateReserveIPSubpool(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestNetworkSettingsUpdateReserveIPSubpool {
	request := catalystcentersdkgo.RequestNetworkSettingsUpdateReserveIPSubpool{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_address_space")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_address_space")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_address_space")))) {
		request.IPv6AddressSpace = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_dhcp_servers")))) {
		request.IPv4DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_dns_servers")))) {
		request.IPv4DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_global_pool")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_global_pool")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_global_pool")))) {
		request.IPv6GlobalPool = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_prefix")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_prefix")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_prefix")))) {
		request.IPv6Prefix = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_prefix_length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_prefix_length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_prefix_length")))) {
		request.IPv6PrefixLength = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_subnet")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_subnet")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_subnet")))) {
		request.IPv6Subnet = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_total_host")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_total_host")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_total_host")))) {
		request.IPv6TotalHost = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_gate_way")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_gate_way")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_gate_way")))) {
		request.IPv6GateWay = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_dhcp_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_dhcp_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_dhcp_servers")))) {
		request.IPv6DhcpServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv6_dns_servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv6_dns_servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv6_dns_servers")))) {
		request.IPv6DNSServers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".slaac_support")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".slaac_support")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".slaac_support")))) {
		request.SLAacSupport = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ipv4_gate_way")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ipv4_gate_way")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ipv4_gate_way")))) {
		request.IPv4GateWay = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchNetworkSettingsGetReserveIPSubpool(m interface{}, queryParams catalystcentersdkgo.GetReserveIPSubpoolQueryParams, vName string) (*catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponse, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponse
	var nResponse *catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpool
	maxPageSize := 500
	offset := 1
	queryParams.Offset = float64(offset)
	queryParams.Limit = float64(maxPageSize)
	nResponse, _, err = client.NetworkSettings.GetReserveIPSubpool(&queryParams)
	if err != nil {
		return foundItem, err
	}

	if nResponse == nil {
		return foundItem, err
	}

	if nResponse.Response == nil {
		return nil, err
	}
	// maxPageSize := len(*nResponse.Response)
	for nResponse != nil && nResponse.Response != nil && len(*nResponse.Response) > 0 {
		for _, item := range *nResponse.Response {
			//			log.Printf("Vname: %s   GroupName: %s", vName, item.GroupName)
			if vName == item.GroupName {
				return &item, err
			}
		}

		queryParams.Limit = float64(maxPageSize)
		offset += maxPageSize
		queryParams.Offset = float64(offset)
		nResponse, _, err = client.NetworkSettings.GetReserveIPSubpool(&queryParams)
		if nResponse == nil || nResponse.Response == nil {
			break
		}
	}
	return foundItem, err
}

func updateReserveIpPoolParameters(request *catalystcentersdkgo.RequestNetworkSettingsReserveIPSubpool, response *catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolResponse) *catalystcentersdkgo.RequestNetworkSettingsReserveIPSubpool {
	log.Printf("IPPOOLREQUEST %s", responseInterfaceToString(request))
	if request == nil {
		request = &catalystcentersdkgo.RequestNetworkSettingsReserveIPSubpool{}
	}
	for _, v := range *response.IPPools {
		log.Printf("IPPOOL %s", responseInterfaceToString(v))
		log.Printf("IPPOOL Dhcp %s", responseInterfaceToString(v.DhcpServerIPs))
		if v.IPv6 != nil && *v.IPv6 {
			if v.IPPoolName == request.Name {
				request.Name = v.IPPoolName
			}
		} else {
			if v.IPPoolName == request.Name {
				request.Name = v.IPPoolName
			}
		}
	}

	log.Printf("IPPOOLREQUEST %s", responseInterfaceToString(request))
	return request
}
