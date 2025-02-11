package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReserveIPSubpool() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- API to get the ip subpool info.
`,

		ReadContext: dataSourceReserveIPSubpoolRead,
		Schema: map[string]*schema.Schema{
			"group_name": &schema.Schema{
				Description: `groupName query parameter. Name of the group
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ignore_inherited_groups": &schema.Schema{
				Description: `ignoreInheritedGroups query parameter. Ignores pools inherited from parent site. Either siteId or ignoreInheritedGroups must be passed. They can also be used together.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of reserve pools to be retrieved. Default is 25 if not specified. Maximum allowed limit is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. offset/starting row. Indexed from 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"pool_usage": &schema.Schema{
				Description: `poolUsage query parameter. Can take values empty, partially-full or empty-partially-full
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. site id of site from which to retrieve associated reserve pools. Either siteId (per site queries) or ignoreInheritedGroups must be used. They can also be used together.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
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
		},
	}
}

func dataSourceReserveIPSubpoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vIgnoreInheritedGroups, okIgnoreInheritedGroups := d.GetOk("ignore_inherited_groups")
	vPoolUsage, okPoolUsage := d.GetOk("pool_usage")
	vGroupName, okGroupName := d.GetOk("group_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetReserveIPSubpoolV1")
		queryParams1 := catalystcentersdkgo.GetReserveIPSubpoolV1QueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okIgnoreInheritedGroups {
			queryParams1.IgnoreInheritedGroups = vIgnoreInheritedGroups.(bool)
		}
		if okPoolUsage {
			queryParams1.PoolUsage = vPoolUsage.(string)
		}
		if okGroupName {
			queryParams1.GroupName = vGroupName.(string)
		}

		response1, restyResp1, err := client.NetworkSettings.GetReserveIPSubpoolV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetReserveIPSubpoolV1", err,
				"Failure at GetReserveIPSubpoolV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsGetReserveIPSubpoolV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetReserveIPSubpoolV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetReserveIPSubpoolV1Items(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["group_name"] = item.GroupName
		respItem["ip_pools"] = flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPools(item.IPPools)
		respItem["site_id"] = item.SiteID
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["type"] = item.Type
		respItem["group_owner"] = item.GroupOwner
		respItems = append(respItems, respItem)
	}
	return respItems
}
func flattenNetworkSettingsGetReserveIPSubpoolParameters(items *catalystcentersdkgo.RequestNetworkSettingsReserveIPSubpoolV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	respItem := make(map[string]interface{})
	respItem["name"] = items.Name
	respItem["type"] = items.Type
	respItem["ipv6_address_space"] = boolPtrToString(items.IPv6AddressSpace)
	respItem["ipv4_global_pool"] = items.IPv4GlobalPool
	respItem["ipv4_prefix"] = boolPtrToString(items.IPv4Prefix)
	respItem["ipv4_prefix_length"] = items.IPv4PrefixLength
	respItem["ipv4_subnet"] = items.IPv4Subnet
	respItem["ipv4_gate_way"] = items.IPv4GateWay
	respItem["ipv4_dhcp_servers"] = items.IPv4DhcpServers
	respItem["ipv4_dns_servers"] = items.IPv4DNSServers
	respItem["ipv6_global_pool"] = items.IPv6GlobalPool
	respItem["ipv6_prefix"] = boolPtrToString(items.IPv6Prefix)
	respItem["ipv6_prefix_length"] = items.IPv6PrefixLength
	respItem["ipv6_subnet"] = items.IPv6Subnet
	respItem["ipv6_gate_way"] = items.IPv6GateWay
	respItem["ipv6_dhcp_servers"] = items.IPv6DhcpServers
	respItem["ipv6_dns_servers"] = items.IPv6DNSServers
	respItem["ipv4_total_host"] = interfaceToIntPtr(items.IPv4TotalHost)
	respItem["ipv6_total_host"] = interfaceToIntPtr(items.IPv6TotalHost)
	respItem["slaac_support"] = boolPtrToString(items.SLAacSupport)

	respItems = append(respItems, respItem)
	return respItems
}

func flattenNetworkSettingsGetReserveIPSubpoolParametersDhcpServers(items []string) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenNetworkSettingsGetReserveIPSubpoolParametersDnsServers(items []string) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
func flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPools(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPools) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["dhcp_server_ips"] = flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPoolsDhcpServerIPs(item.DhcpServerIPs)
		respItem["gateways"] = item.Gateways
		respItem["create_time"] = item.CreateTime
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["total_ip_address_count"] = item.TotalIPAddressCount
		respItem["used_ip_address_count"] = item.UsedIPAddressCount
		respItem["parent_uuid"] = item.ParentUUID
		respItem["owner"] = item.Owner
		respItem["shared"] = boolPtrToString(item.Shared)
		respItem["overlapping"] = boolPtrToString(item.Overlapping)
		respItem["configure_external_dhcp"] = boolPtrToString(item.ConfigureExternalDhcp)
		respItem["used_percentage"] = item.UsedPercentage
		respItem["client_options"] = flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPoolsClientOptions(item.ClientOptions)
		respItem["group_uuid"] = item.GroupUUID
		respItem["dns_server_ips"] = flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPoolsDNSServerIPs(item.DNSServerIPs)
		respItem["context"] = flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPoolsContext(item.Context)
		respItem["ipv6"] = boolPtrToString(item.IPv6)
		respItem["id"] = item.ID
		respItem["ip_pool_cidr"] = item.IPPoolCidr
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPoolsDhcpServerIPs(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsDhcpServerIPs) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPoolsClientOptions(item *catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsClientOptions) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPoolsDNSServerIPs(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsDNSServerIPs) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}

func flattenNetworkSettingsGetReserveIPSubpoolV1ItemsIPPoolsContext(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetReserveIPSubpoolV1ResponseIPPoolsContext) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["owner"] = item.Owner
		respItem["context_key"] = item.ContextKey
		respItem["context_value"] = item.ContextValue
		respItems = append(respItems, respItem)
	}
	return respItems
}
