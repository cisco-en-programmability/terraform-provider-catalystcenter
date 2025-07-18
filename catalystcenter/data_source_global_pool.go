package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGlobalPool() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- API to get the global pool.
`,

		ReadContext: dataSourceGlobalPoolRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of Global Pools to be retrieved. Default is 25 if not specified.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Offset/starting row. Indexed from 1. Default value of 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
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
		},
	}
}

func dataSourceGlobalPoolRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetGlobalPool")
		queryParams1 := catalystcentersdkgo.GetGlobalPoolQueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.NetworkSettings.GetGlobalPool(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetGlobalPool", err,
				"Failure at GetGlobalPool, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetGlobalPool", err,
				"Failure at GetGlobalPool, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsGetGlobalPoolItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalPool response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetGlobalPoolItems(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ip_pool_name"] = item.IPPoolName
		respItem["dhcp_server_ips"] = item.DhcpServerIPs
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
		respItem["client_options"] = flattenNetworkSettingsGetGlobalPoolItemsClientOptions(item.ClientOptions)
		respItem["ip_pool_type"] = item.IPPoolType
		respItem["unavailable_ip_address_count"] = item.UnavailableIPAddressCount
		respItem["available_ip_address_count"] = item.AvailableIPAddressCount
		respItem["total_assignable_ip_address_count"] = item.TotalAssignableIPAddressCount
		respItem["dns_server_ips"] = item.DNSServerIPs
		respItem["has_subpools"] = boolPtrToString(item.HasSubpools)
		respItem["default_assigned_ip_address_count"] = item.DefaultAssignedIPAddressCount
		respItem["context"] = flattenNetworkSettingsGetGlobalPoolItemsContext(item.Context)
		respItem["ipv6"] = boolPtrToString(item.IPv6)
		respItem["id"] = item.ID
		respItem["ip_pool_cidr"] = item.IPPoolCidr
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenNetworkSettingsGetGlobalPoolItemsClientOptions(item *catalystcentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponseClientOptions) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenNetworkSettingsGetGlobalPoolItemsContext(items *[]catalystcentersdkgo.ResponseNetworkSettingsGetGlobalPoolResponseContext) []map[string]interface{} {
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
