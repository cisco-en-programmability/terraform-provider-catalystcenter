package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceClientEnrichmentDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Clients.

- Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the
user, the devices that the user is connected to and the assurance issues that the user is impacted by
`,

		ReadContext: dataSourceClientEnrichmentDetailsRead,
		Schema: map[string]*schema.Schema{
			"entity_type": &schema.Schema{
				Description: `entity_type header parameter. Client enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"entity_value": &schema.Schema{
				Description: `entity_value header parameter. Contains the actual value for the entity type that has been defined
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"issue_category": &schema.Schema{
				Description: `issueCategory header parameter. The category of the DNA event based on which the underlying issues need to be fetched
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connected_device": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ap_manager_interface_ip": &schema.Schema{
													Description: `Ap Manager Interface Ip`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"associated_wlc_ip": &schema.Schema{
													Description: `Associated Wlc Ip`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"boot_date_time": &schema.Schema{
													Description: `Boot Date Time`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"cisco360view": &schema.Schema{
													Description: `Cisco360view`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"collection_interval": &schema.Schema{
													Description: `Collection Interval`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"collection_status": &schema.Schema{
													Description: `Collection Status`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"error_code": &schema.Schema{
													Description: `Error Code`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"error_description": &schema.Schema{
													Description: `Error Description`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"family": &schema.Schema{
													Description: `Family`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"hostname": &schema.Schema{
													Description: `Hostname`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"instance_uuid": &schema.Schema{
													Description: `Instance Uuid`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"interface_count": &schema.Schema{
													Description: `Interface Count`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"inventory_status_detail": &schema.Schema{
													Description: `Inventory Status Detail`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"last_update_time": &schema.Schema{
													Description: `Last Update Time`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"last_updated": &schema.Schema{
													Description: `Last Updated`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"line_card_count": &schema.Schema{
													Description: `Line Card Count`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"line_card_id": &schema.Schema{
													Description: `Line Card Id`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"location": &schema.Schema{
													Description: `Location`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"location_name": &schema.Schema{
													Description: `Location Name`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"mac_address": &schema.Schema{
													Description: `Mac Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"management_ip_address": &schema.Schema{
													Description: `Management Ip Address`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"memory_size": &schema.Schema{
													Description: `Memory Size`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"neighbor_topology": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"links": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"id": &schema.Schema{
																			Description: `Id`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"label": &schema.Schema{
																			Description: `Label`,
																			Type:        schema.TypeList,
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"link_status": &schema.Schema{
																			Description: `Link Status`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"port_utilization": &schema.Schema{
																			Description: `Port Utilization`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"source": &schema.Schema{
																			Description: `Source`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"target": &schema.Schema{
																			Description: `Target`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},

															"nodes": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"clients": &schema.Schema{
																			Description: `Clients`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},

																		"count": &schema.Schema{
																			Description: `Count`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"description": &schema.Schema{
																			Description: `Description`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"device_type": &schema.Schema{
																			Description: `Device Type`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"fabric_group": &schema.Schema{
																			Description: `Fabric Group`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"family": &schema.Schema{
																			Description: `Family`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"health_score": &schema.Schema{
																			Description: `Health Score`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"id": &schema.Schema{
																			Description: `Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"ip": &schema.Schema{
																			Description: `Ip`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"level": &schema.Schema{
																			Description: `Level`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},

																		"name": &schema.Schema{
																			Description: `Name`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"node_type": &schema.Schema{
																			Description: `Node Type`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"platform_id": &schema.Schema{
																			Description: `Platform Id`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"radio_frequency": &schema.Schema{
																			Description: `Radio Frequency`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"role": &schema.Schema{
																			Description: `Role`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"software_version": &schema.Schema{
																			Description: `Software Version`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"user_id": &schema.Schema{
																			Description: `User Id`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},

												"platform_id": &schema.Schema{
													Description: `Platform Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"reachability_failure_reason": &schema.Schema{
													Description: `Reachability Failure Reason`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"reachability_status": &schema.Schema{
													Description: `Reachability Status`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"role": &schema.Schema{
													Description: `Role`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"role_source": &schema.Schema{
													Description: `Role Source`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"serial_number": &schema.Schema{
													Description: `Serial Number`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"series": &schema.Schema{
													Description: `Series`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"snmp_contact": &schema.Schema{
													Description: `Snmp Contact`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"snmp_location": &schema.Schema{
													Description: `Snmp Location`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"software_version": &schema.Schema{
													Description: `Software Version`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"tag_count": &schema.Schema{
													Description: `Tag Count`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"tunnel_udp_port": &schema.Schema{
													Description: `Tunnel Udp Port`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"up_time": &schema.Schema{
													Description: `Up Time`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"waas_device_mode": &schema.Schema{
													Description: `Waas Device Mode`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},

						"issue_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"issue": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"impacted_hosts": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connected_interface": &schema.Schema{
																Description: `Connected Interface`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"failed_attempts": &schema.Schema{
																Description: `Failed Attempts`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"host_name": &schema.Schema{
																Description: `Host Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"host_os": &schema.Schema{
																Description: `Host Os`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"host_type": &schema.Schema{
																Description: `Host Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"location": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"aps_impacted": &schema.Schema{
																			Description: `Aps Impacted`,
																			Type:        schema.TypeList,
																			Computed:    true,
																			Elem: &schema.Schema{
																				Type: schema.TypeString,
																			},
																		},

																		"area": &schema.Schema{
																			Description: `Area`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"building": &schema.Schema{
																			Description: `Building`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"floor": &schema.Schema{
																			Description: `Floor`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},

																		"site_id": &schema.Schema{
																			Description: `Site Id`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},

																		"site_type": &schema.Schema{
																			Description: `Site Type`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},

															"mac_address": &schema.Schema{
																Description: `Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"ssid": &schema.Schema{
																Description: `Ssid`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"timestamp": &schema.Schema{
																Description: `Timestamp`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
														},
													},
												},

												"issue_category": &schema.Schema{
													Description: `Issue Category`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_description": &schema.Schema{
													Description: `Issue Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_entity": &schema.Schema{
													Description: `Issue Entity`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_entity_value": &schema.Schema{
													Description: `Issue Entity Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_id": &schema.Schema{
													Description: `Issue Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_name": &schema.Schema{
													Description: `Issue Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_priority": &schema.Schema{
													Description: `Issue Priority`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_severity": &schema.Schema{
													Description: `Issue Severity`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_source": &schema.Schema{
													Description: `Issue Source`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_summary": &schema.Schema{
													Description: `Issue Summary`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"issue_timestamp": &schema.Schema{
													Description: `Issue Timestamp`,
													Type:        schema.TypeInt,
													Computed:    true,
												},

												"suggested_actions": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"message": &schema.Schema{
																Description: `Message`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"steps": &schema.Schema{
																Description: `Steps`,
																Type:        schema.TypeList,
																Computed:    true,
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

						"user_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"client_connection": &schema.Schema{
										Description: `Client Connection`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"connected_device": &schema.Schema{
										Description: `Connected Device`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"connection_status": &schema.Schema{
										Description: `Connection Status`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"data_rate": &schema.Schema{
										Description: `Data Rate`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"health_score": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"health_type": &schema.Schema{
													Description: `Health Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"reason": &schema.Schema{
													Description: `Reason`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"score": &schema.Schema{
													Description: `Score`,
													Type:        schema.TypeInt,
													Computed:    true,
												},
											},
										},
									},

									"host_ip_v4": &schema.Schema{
										Description: `Host Ip V4`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_ip_v6": &schema.Schema{
										Description: `Host Ip V6`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"host_mac": &schema.Schema{
										Description: `Host Mac`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_name": &schema.Schema{
										Description: `Host Name`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"host_os": &schema.Schema{
										Description: `Host Os`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"host_type": &schema.Schema{
										Description: `Host Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"host_version": &schema.Schema{
										Description: `Host Version`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"issue_count": &schema.Schema{
										Description: `Issue Count`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},

									"last_updated": &schema.Schema{
										Description: `Last Updated`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"location": &schema.Schema{
										Description: `Location`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"port": &schema.Schema{
										Description: `Port`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"rssi": &schema.Schema{
										Description: `Rssi`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"snr": &schema.Schema{
										Description: `Snr`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"ssid": &schema.Schema{
										Description: `Ssid`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"sub_type": &schema.Schema{
										Description: `Sub Type`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"user_id": &schema.Schema{
										Description: `User Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"vlan_id": &schema.Schema{
										Description: `Vlan Id`,
										Type:        schema.TypeString,
										Computed:    true,
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

func dataSourceClientEnrichmentDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vEntityType := d.Get("entity_type")
	vEntityValue := d.Get("entity_value")
	vIssueCategory := d.Get("issue_category")
	vPersistbapioutput := d.Get("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetClientEnrichmentDetailsV1")

		headerParams1 := catalystcentersdkgo.GetClientEnrichmentDetailsV1HeaderParams{}

		headerParams1.EntityType = vEntityType.(string)

		headerParams1.EntityValue = vEntityValue.(string)

		headerParams1.IssueCategory = vIssueCategory.(string)

		headerParams1.Persistbapioutput = vPersistbapioutput.(string)

		response1, restyResp1, err := client.Clients.GetClientEnrichmentDetailsV1(&headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetClientEnrichmentDetailsV1", err,
				"Failure at GetClientEnrichmentDetailsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenClientsGetClientEnrichmentDetailsV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetClientEnrichmentDetailsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenClientsGetClientEnrichmentDetailsV1Items(items *catalystcentersdkgo.ResponseClientsGetClientEnrichmentDetailsV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["user_details"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetails(item.UserDetails)
		respItem["connected_device"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDevice(item.ConnectedDevice)
		respItem["issue_details"] = flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetails(item.IssueDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetails(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["connection_status"] = item.ConnectionStatus
	respItem["host_type"] = item.HostType
	respItem["user_id"] = item.UserID
	respItem["host_name"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHostName(item.HostName)
	respItem["host_os"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHostOs(item.HostOs)
	respItem["host_version"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHostVersion(item.HostVersion)
	respItem["sub_type"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsSubType(item.SubType)
	respItem["last_updated"] = item.LastUpdated
	respItem["health_score"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHealthScore(item.HealthScore)
	respItem["host_mac"] = item.HostMac
	respItem["host_ip_v4"] = item.HostIPV4
	respItem["host_ip_v6"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHostIPV6(item.HostIPV6)
	respItem["auth_type"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsAuthType(item.AuthType)
	respItem["vlan_id"] = item.VLANID
	respItem["ssid"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsSSID(item.SSID)
	respItem["location"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsLocation(item.Location)
	respItem["client_connection"] = item.ClientConnection
	respItem["connected_device"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsConnectedDevice(item.ConnectedDevice)
	respItem["issue_count"] = item.IssueCount
	respItem["rssi"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsRssi(item.Rssi)
	respItem["snr"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsSnr(item.Snr)
	respItem["data_rate"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsDataRate(item.DataRate)
	respItem["port"] = flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsPort(item.Port)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHostName(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHostOs(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostOs) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHostVersion(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostVersion) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsSubType(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSubType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHealthScore(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHealthScore) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["health_type"] = item.HealthType
		respItem["reason"] = item.Reason
		respItem["score"] = item.Score
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsHostIPV6(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsHostIPV6) []interface{} {
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

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsAuthType(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsAuthType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsSSID(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSSID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsLocation(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsConnectedDevice(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsConnectedDevice) []interface{} {
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

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsRssi(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsRssi) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsSnr(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsSnr) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsDataRate(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsDataRate) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsUserDetailsPort(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1UserDetailsPort) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDevice(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDevice) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_details"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetails(item.DeviceDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetails(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["family"] = item.Family
	respItem["type"] = item.Type
	respItem["location"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLocation(item.Location)
	respItem["error_code"] = item.ErrorCode
	respItem["mac_address"] = item.MacAddress
	respItem["role"] = item.Role
	respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
	respItem["associated_wlc_ip"] = item.AssociatedWlcIP
	respItem["boot_date_time"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsBootDateTime(item.BootDateTime)
	respItem["collection_status"] = item.CollectionStatus
	respItem["interface_count"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsInterfaceCount(item.InterfaceCount)
	respItem["line_card_count"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLineCardCount(item.LineCardCount)
	respItem["line_card_id"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLineCardID(item.LineCardID)
	respItem["management_ip_address"] = item.ManagementIPAddress
	respItem["memory_size"] = item.MemorySize
	respItem["platform_id"] = item.PlatformID
	respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
	respItem["reachability_status"] = item.ReachabilityStatus
	respItem["snmp_contact"] = item.SNMPContact
	respItem["snmp_location"] = item.SNMPLocation
	respItem["tunnel_udp_port"] = item.TunnelUDPPort
	respItem["waas_device_mode"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsWaasDeviceMode(item.WaasDeviceMode)
	respItem["series"] = item.Series
	respItem["inventory_status_detail"] = item.InventoryStatusDetail
	respItem["collection_interval"] = item.CollectionInterval
	respItem["serial_number"] = item.SerialNumber
	respItem["software_version"] = item.SoftwareVersion
	respItem["role_source"] = item.RoleSource
	respItem["hostname"] = item.Hostname
	respItem["up_time"] = item.UpTime
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["error_description"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsErrorDescription(item.ErrorDescription)
	respItem["location_name"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLocationName(item.LocationName)
	respItem["tag_count"] = item.TagCount
	respItem["last_updated"] = item.LastUpdated
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["id"] = item.ID
	respItem["neighbor_topology"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopology(item.NeighborTopology)
	respItem["cisco360view"] = item.Cisco360View

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLocation(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsBootDateTime(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsBootDateTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsInterfaceCount(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsInterfaceCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLineCardCount(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLineCardCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLineCardID(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLineCardID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsWaasDeviceMode(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsWaasDeviceMode) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsErrorDescription(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorDescription) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLocationName(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocationName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopology(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopology) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["nodes"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodes(item.Nodes)
		respItem["links"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyLinks(item.Links)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodes(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["role"] = item.Role
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["description"] = item.Description
		respItem["device_type"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType(item.DeviceType)
		respItem["platform_id"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID(item.PlatformID)
		respItem["family"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily(item.Family)
		respItem["ip"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP(item.IP)
		respItem["software_version"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion(item.SoftwareVersion)
		respItem["user_id"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID(item.UserID)
		respItem["node_type"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType(item.NodeType)
		respItem["radio_frequency"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency(item.RadioFrequency)
		respItem["clients"] = item.Clients
		respItem["count"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount(item.Count)
		respItem["health_score"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore(item.HealthScore)
		respItem["level"] = item.Level
		respItem["fabric_group"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup(item.FabricGroup)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesDeviceType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesPlatformID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesFamily) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesIP(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesIP) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesSoftwareVersion) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesUserID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesNodeType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesRadioFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesCount(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesCount) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesHealthScore) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyNodesFabricGroup) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyLinks(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["source"] = item.Source
		respItem["link_status"] = item.LinkStatus
		respItem["label"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel(item.Label)
		respItem["target"] = item.Target
		respItem["id"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksID(item.ID)
		respItem["port_utilization"] = flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization(item.PortUtilization)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksLabel) []interface{} {
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

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksID(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopologyLinksPortUtilization) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetails(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["issue"] = flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssue(item.Issue)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssue(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssue) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["issue_id"] = item.IssueID
		respItem["issue_source"] = item.IssueSource
		respItem["issue_category"] = item.IssueCategory
		respItem["issue_name"] = item.IssueName
		respItem["issue_description"] = item.IssueDescription
		respItem["issue_entity"] = item.IssueEntity
		respItem["issue_entity_value"] = item.IssueEntityValue
		respItem["issue_severity"] = item.IssueSeverity
		respItem["issue_priority"] = item.IssuePriority
		respItem["issue_summary"] = item.IssueSummary
		respItem["issue_timestamp"] = item.IssueTimestamp
		respItem["suggested_actions"] = flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueSuggestedActions(item.SuggestedActions)
		respItem["impacted_hosts"] = flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueImpactedHosts(item.ImpactedHosts)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueSuggestedActions(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueSuggestedActions) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItem["steps"] = flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueSuggestedActionsSteps(item.Steps)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueSuggestedActionsSteps(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueSuggestedActionsSteps) []interface{} {
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

func flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueImpactedHosts(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHosts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["host_type"] = item.HostType
		respItem["host_name"] = item.HostName
		respItem["host_os"] = item.HostOs
		respItem["ssid"] = item.SSID
		respItem["connected_interface"] = item.ConnectedInterface
		respItem["mac_address"] = item.MacAddress
		respItem["failed_attempts"] = item.FailedAttempts
		respItem["location"] = flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueImpactedHostsLocation(item.Location)
		respItem["timestamp"] = item.Timestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueImpactedHostsLocation(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site_id"] = item.SiteID
	respItem["site_type"] = item.SiteType
	respItem["area"] = item.Area
	respItem["building"] = item.Building
	respItem["floor"] = flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueImpactedHostsLocationFloor(item.Floor)
	respItem["aps_impacted"] = flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueImpactedHostsLocationApsImpacted(item.ApsImpacted)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueImpactedHostsLocationFloor(item *catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocationFloor) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenClientsGetClientEnrichmentDetailsV1ItemsIssueDetailsIssueImpactedHostsLocationApsImpacted(items *[]catalystcentersdkgo.ResponseItemClientsGetClientEnrichmentDetailsV1IssueDetailsIssueImpactedHostsLocationApsImpacted) []interface{} {
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
