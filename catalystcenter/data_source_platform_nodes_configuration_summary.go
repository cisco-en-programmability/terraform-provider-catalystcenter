package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePlatformNodesConfigurationSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Platform Configuration.

- Provides details about the current Cisco Catalyst Center node configuration, such as API version, node name, NTP
server, intracluster link, LACP mode, network static routes, DNS server, subnet mask, host IP, default gateway, and
interface information.
`,

		ReadContext: dataSourcePlatformNodesConfigurationSummaryRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"nodes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Cluster Identifier
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Node name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"network": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"inet": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dns_servers": &schema.Schema{
																Description: `DNS server
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"gateway": &schema.Schema{
																Description: `Default gateway
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"host_ip": &schema.Schema{
																Description: `IP assigned
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"netmask": &schema.Schema{
																Description: `Subnet mask
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"routes": &schema.Schema{
																Description: `Static route
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},

												"inet6": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"host_ip": &schema.Schema{
																Description: `IP assigned to the host
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"netmask": &schema.Schema{
																Description: `Subnet mask of the host
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"interface": &schema.Schema{
													Description: `Interface name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"intra_cluster_link": &schema.Schema{
													Description: `Flag to indicate which interface is configured as the inter-cluster link
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"lacp_mode": &schema.Schema{
													Description: `LACP Mode configuration on NIC
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"lacp_supported": &schema.Schema{
													Description: `LACP Support configuration on NIC
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"slave": &schema.Schema{
													Description: `Physical interface name
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"ntp": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"servers": &schema.Schema{
													Description: `NTP server
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"platform": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product": &schema.Schema{
													Description: `Product Identifier
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"serial": &schema.Schema{
													Description: `Serial number of chassis
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vendor": &schema.Schema{
													Description: `Product manufacturer
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"proxy": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"http_proxy": &schema.Schema{
													Description: `Not Supported
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"https_proxy": &schema.Schema{
													Description: `Https Proxy Server
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"https_proxy_password": &schema.Schema{
													Description: `Configured Https excrypted proxy password.
`,
													Type:      schema.TypeString,
													Sensitive: true,
													Computed:  true,
												},

												"https_proxy_username": &schema.Schema{
													Description: `Configured Https proxy username
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"no_proxy": &schema.Schema{
													Description: `Servers configured to explicitly use no proxy
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
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
		},
	}
}

func dataSourcePlatformNodesConfigurationSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CiscoCatalystCenterNodesConfigurationSummaryV1")

		// has_unknown_response: None

		response1, restyResp1, err := client.Platform.CiscoCatalystCenterNodesConfigurationSummaryV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CiscoCatalystCenterNodesConfigurationSummaryV1", err,
				"Failure at CiscoCatalystCenterNodesConfigurationSummaryV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CiscoCatalystCenterNodesConfigurationSummaryV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1Item(item *catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["nodes"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodes(item.Nodes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodes(items *[]catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ntp"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNtp(item.Ntp)
		respItem["network"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetwork(item.Network)
		respItem["proxy"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesProxy(item.Proxy)
		respItem["platform"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesPlatform(item.Platform)
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNtp(item *catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNtp) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["servers"] = item.Servers

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetwork(items *[]catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetwork) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["intra_cluster_link"] = boolPtrToString(item.IntraClusterLink)
		respItem["lacp_mode"] = boolPtrToString(item.LacpMode)
		respItem["inet"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetworkInet(item.Inet)
		respItem["interface"] = item.Interface
		respItem["inet6"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetworkInet6(item.Inet6)
		respItem["lacp_supported"] = boolPtrToString(item.LacpSupported)
		respItem["slave"] = item.SLAve
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetworkInet(item *catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["routes"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetworkInetRoutes(item.Routes)
	respItem["gateway"] = item.Gateway
	respItem["dns_servers"] = flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetworkInetDNSServers(item.DNSServers)
	respItem["netmask"] = item.Netmask
	respItem["host_ip"] = item.HostIP

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetworkInetRoutes(items *[]catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetRoutes) []interface{} {
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

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetworkInetDNSServers(items *[]catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInetDNSServers) []interface{} {
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

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesNetworkInet6(item *catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesNetworkInet6) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["host_ip"] = item.HostIP
	respItem["netmask"] = item.Netmask

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesProxy(item *catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesProxy) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["https_proxy"] = item.HTTPSProxy
	respItem["no_proxy"] = item.NoProxy
	respItem["https_proxy_username"] = item.HTTPSProxyUsername
	respItem["http_proxy"] = item.HTTPProxy
	respItem["https_proxy_password"] = item.HTTPSProxyPassword

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPlatformConfigurationCiscoCatalystCenterNodesConfigurationSummaryV1ItemNodesPlatform(item *catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterNodesConfigurationSummaryV1ResponseNodesPlatform) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["vendor"] = item.Vendor
	respItem["product"] = item.Product
	respItem["serial"] = item.Serial

	return []map[string]interface{}{
		respItem,
	}

}
