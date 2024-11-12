package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDisasterrecoverySystemStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Disaster Recovery.

- Detailed and Summarized status of DR components (Active, Standby and Witness system's health).
`,

		ReadContext: dataSourceDisasterrecoverySystemStatusRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ipconfig": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface": &schema.Schema{
										Description: `Enterprise or Management interface
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ip": &schema.Schema{
										Description: `This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.  
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"vip": &schema.Schema{
										Description: `Is this interface an Virtual IP address or not. This is true for Global DR VIP
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"ipsec_tunnel": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"side_a": &schema.Schema{
										Description: `A Side of the IPSec Tunnel
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"side_b": &schema.Schema{
										Description: `Other Side of the IPSec Tunnel 
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Status of this IPSec Tunnel
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"main": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipconfig": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Description: `Enterprise or Management interface
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip": &schema.Schema{
													Description: `This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vip": &schema.Schema{
													Description: `Is this interface an Virtual IP address or not. This is true for cluster level.
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"nodes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": &schema.Schema{
													Description: `Hostname of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipaddresses": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"interface": &schema.Schema{
																Description: `Enterprise or Management interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ip": &schema.Schema{
																Description: `Node IP address
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"vip": &schema.Schema{
																Description: `Is this interface a Virtual IP address or not. This is false for node level.
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"state": &schema.Schema{
													Description: `State of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"state": &schema.Schema{
										Description: `State of the Main Site. 
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"recovery": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipconfig": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Description: `Enterprise or Management interface
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip": &schema.Schema{
													Description: `This floating virtual IP address automatically moves to and operates on the site that is currently acting as your network's active site. If the interface is enterprise, then this address manages traffic between your disaster recovery system and your Enterprise network. If the interface is management, then this address manages traffic between your disaster recovery system and your Management network.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vip": &schema.Schema{
													Description: `Is this interface an Virtual IP address or not. This is true for cluster level.
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"nodes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": &schema.Schema{
													Description: `Hostname of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipconfig": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"interface": &schema.Schema{
																Description: `Enterprise or Management interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ip": &schema.Schema{
																Description: `Node IP Address
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"vip": &schema.Schema{
																Description: `Is this interface a Virtual IP Address or not. This is false for node level.
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"state": &schema.Schema{
													Description: `State of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"state": &schema.Schema{
										Description: `State of the Recovery site
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"site": &schema.Schema{
							Description: `Site of the disaster recovery system. 
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"state": &schema.Schema{
							Description: `State of the Disaster Recovery System.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"witness": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ipconfig": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"interface": &schema.Schema{
													Description: `Enterprise or Management interface
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ip": &schema.Schema{
													Description: `In case of witness, this is only an IP. 
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"vip": &schema.Schema{
													Description: `Is this interface an Virtual IP address or not. This is false for witness.
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"nodes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"hostname": &schema.Schema{
													Description: `Hostname of the witness node
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipconfig": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"interface": &schema.Schema{
																Description: `Enterprise or Management interface
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ip": &schema.Schema{
																Description: `In case of witness, this is only an IP
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"vip": &schema.Schema{
																Description: `Is this interface an Virtual IP address or not. This is false for Witness 
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"state": &schema.Schema{
													Description: `State of the node
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"state": &schema.Schema{
										Description: `State of the Witness Site
`,
										Type:     schema.TypeString,
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

func dataSourceDisasterrecoverySystemStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: DisasterRecoveryStatusV1")

		response1, restyResp1, err := client.DisasterRecovery.DisasterRecoveryStatusV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 DisasterRecoveryStatusV1", err,
				"Failure at DisasterRecoveryStatusV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDisasterRecoveryDisasterRecoveryStatusV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting DisasterRecoveryStatusV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1Item(item *catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemIPconfig(item.IPconfig)
	respItem["site"] = item.Site
	respItem["main"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemMain(item.Main)
	respItem["recovery"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemRecovery(item.Recovery)
	respItem["witness"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemWitness(item.Witness)
	respItem["state"] = item.State
	respItem["ipsec_tunnel"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemIPsecTunnel(item.IPsecTunnel)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemIPconfig(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1IPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemMain(item *catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1Main) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemMainIPconfig(item.IPconfig)
	respItem["state"] = item.State
	respItem["nodes"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemMainNodes(item.Nodes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemMainIPconfig(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1MainIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemMainNodes(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1MainNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hostname"] = item.Hostname
		respItem["state"] = item.State
		respItem["ipaddresses"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemMainNodesIPaddresses(item.IPaddresses)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemMainNodesIPaddresses(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1MainNodesIPaddresses) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemRecovery(item *catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1Recovery) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemRecoveryIPconfig(item.IPconfig)
	respItem["state"] = item.State
	respItem["nodes"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemRecoveryNodes(item.Nodes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemRecoveryIPconfig(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemRecoveryNodes(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hostname"] = item.Hostname
		respItem["state"] = item.State
		respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemRecoveryNodesIPconfig(item.IPconfig)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemRecoveryNodesIPconfig(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1RecoveryNodesIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemWitness(item *catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1Witness) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemWitnessIPconfig(item.IPconfig)
	respItem["state"] = item.State
	respItem["nodes"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemWitnessNodes(item.Nodes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemWitnessIPconfig(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemWitnessNodes(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["hostname"] = item.Hostname
		respItem["state"] = item.State
		respItem["ipconfig"] = flattenDisasterRecoveryDisasterRecoveryStatusV1ItemWitnessNodesIPconfig(item.IPconfig)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemWitnessNodesIPconfig(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1WitnessNodesIPconfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface"] = item.Interface
		respItem["vip"] = boolPtrToString(item.Vip)
		respItem["ip"] = item.IP
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDisasterRecoveryDisasterRecoveryStatusV1ItemIPsecTunnel(items *[]catalystcentersdkgo.ResponseDisasterRecoveryDisasterRecoveryStatusV1IPsecTunnel) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["side_a"] = item.SideA
		respItem["side_b"] = item.SideB
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
