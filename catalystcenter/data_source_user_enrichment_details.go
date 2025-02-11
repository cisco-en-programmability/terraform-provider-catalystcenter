package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUserEnrichmentDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Users.

- Enriches a given network End User context (a network user-id or end user’s device Mac Address) with details about the
user and devices that the user is connected to
`,

		ReadContext: dataSourceUserEnrichmentDetailsRead,
		Schema: map[string]*schema.Schema{
			"entity_type": &schema.Schema{
				Description: `entity_type header parameter. User enrichment details can be fetched based on either User ID or Client MAC address. This parameter value must either be network_user_id/mac_address
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
													Type:        schema.TypeString, //TEST,
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
													Type:        schema.TypeString,
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
													Type:        schema.TypeString,
													Computed:    true,
												},

												"line_card_id": &schema.Schema{
													Description: `Line Card Id`,
													Type:        schema.TypeString,
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

															"detail": &schema.Schema{
																Description: `Detail`,
																Type:        schema.TypeString,
																Computed:    true,
															},

															"error_code": &schema.Schema{
																Description: `Error Code`,
																Type:        schema.TypeInt,
																Computed:    true,
															},

															"message": &schema.Schema{
																Description: `Message`,
																Type:        schema.TypeString,
																Computed:    true,
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
													Type:        schema.TypeString, //TEST,
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

						"user_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_group": &schema.Schema{
										Description: `Ap Group`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"auth_type": &schema.Schema{
										Description: `Auth Type`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"avg_rssi": &schema.Schema{
										Description: `Avg Rssi`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"avg_snr": &schema.Schema{
										Description: `Avg Snr`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"channel": &schema.Schema{
										Description: `Channel`,
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

									"dns_failure": &schema.Schema{
										Description: `Dns Failure`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"dns_success": &schema.Schema{
										Description: `Dns Success`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"frequency": &schema.Schema{
										Description: `Frequency`,
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

									"onboarding": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aaa_server_ip": &schema.Schema{
													Description: `Aaa Server Ip`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"average_assoc_duration": &schema.Schema{
													Description: `Average Assoc Duration`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"average_auth_duration": &schema.Schema{
													Description: `Average Auth Duration`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"average_dhcp_duration": &schema.Schema{
													Description: `Average Dhcp Duration`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"average_run_duration": &schema.Schema{
													Description: `Average Run Duration`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"dhcp_server_ip": &schema.Schema{
													Description: `Dhcp Server Ip`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"max_assoc_duration": &schema.Schema{
													Description: `Max Assoc Duration`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"max_auth_duration": &schema.Schema{
													Description: `Max Auth Duration`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"max_dhcp_duration": &schema.Schema{
													Description: `Max Dhcp Duration`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},

												"max_run_duration": &schema.Schema{
													Description: `Max Run Duration`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},
											},
										},
									},

									"onboarding_time": &schema.Schema{
										Description: `Onboarding Time`,
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

									"rx_bytes": &schema.Schema{
										Description: `Rx Bytes`,
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
										Type:        schema.TypeString,
										Computed:    true,
									},

									"tx_bytes": &schema.Schema{
										Description: `Tx Bytes`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},

									"user_id": &schema.Schema{
										Description: `User Id`,
										Type:        schema.TypeString, //TEST,
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

func dataSourceUserEnrichmentDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vEntityType := d.Get("entity_type")
	vEntityValue := d.Get("entity_value")
	vPersistbapioutput := d.Get("persistbapioutput")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetUserEnrichmentDetailsV1")

		headerParams1 := catalystcentersdkgo.GetUserEnrichmentDetailsV1HeaderParams{}

		headerParams1.EntityType = vEntityType.(string)

		headerParams1.EntityValue = vEntityValue.(string)

		headerParams1.Persistbapioutput = vPersistbapioutput.(string)

		response1, restyResp1, err := client.Users.GetUserEnrichmentDetailsV1(&headerParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetUserEnrichmentDetailsV1", err,
				"Failure at GetUserEnrichmentDetailsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenUsersGetUserEnrichmentDetailsV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetUserEnrichmentDetailsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUsersGetUserEnrichmentDetailsV1Items(items *catalystcentersdkgo.ResponseUsersGetUserEnrichmentDetailsV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["user_details"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetails(item.UserDetails)
		respItem["connected_device"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDevice(item.ConnectedDevice)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetails(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["connection_status"] = item.ConnectionStatus
	respItem["host_type"] = item.HostType
	respItem["user_id"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsUserID(item.UserID)
	respItem["host_name"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHostName(item.HostName)
	respItem["host_os"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHostOs(item.HostOs)
	respItem["host_version"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHostVersion(item.HostVersion)
	respItem["sub_type"] = item.SubType
	respItem["last_updated"] = item.LastUpdated
	respItem["health_score"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHealthScore(item.HealthScore)
	respItem["host_mac"] = item.HostMac
	respItem["host_ip_v4"] = item.HostIPV4
	respItem["host_ip_v6"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHostIPV6(item.HostIPV6)
	respItem["auth_type"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsAuthType(item.AuthType)
	respItem["vlan_id"] = item.VLANID
	respItem["ssid"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsSSID(item.SSID)
	respItem["frequency"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsFrequency(item.Frequency)
	respItem["channel"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsChannel(item.Channel)
	respItem["ap_group"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsApGroup(item.ApGroup)
	respItem["location"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsLocation(item.Location)
	respItem["client_connection"] = item.ClientConnection
	respItem["connected_device"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsConnectedDevice(item.ConnectedDevice)
	respItem["issue_count"] = item.IssueCount
	respItem["rssi"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsRssi(item.Rssi)
	respItem["avg_rssi"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsAvgRssi(item.AvgRssi)
	respItem["snr"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsSnr(item.Snr)
	respItem["avg_snr"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsAvgSnr(item.AvgSnr)
	respItem["data_rate"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsDataRate(item.DataRate)
	respItem["tx_bytes"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsTxBytes(item.TxBytes)
	respItem["rx_bytes"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsRxBytes(item.RxBytes)
	respItem["dns_success"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsDNSSuccess(item.DNSSuccess)
	respItem["dns_failure"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsDNSFailure(item.DNSFailure)
	respItem["onboarding"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboarding(item.Onboarding)
	respItem["onboarding_time"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingTime(item.OnboardingTime)
	respItem["port"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsPort(item.Port)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsUserID(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsUserID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHostName(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHostOs(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostOs) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHostVersion(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostVersion) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHealthScore(items *[]catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHealthScore) []map[string]interface{} {
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

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsHostIPV6(items *[]catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsHostIPV6) []interface{} {
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

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsAuthType(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAuthType) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsSSID(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsSSID) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsFrequency(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsFrequency) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsChannel(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsChannel) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsApGroup(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsApGroup) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsLocation(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsConnectedDevice(items *[]catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsConnectedDevice) []interface{} {
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

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsRssi(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsRssi) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsAvgRssi(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAvgRssi) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsSnr(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsSnr) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsAvgSnr(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsAvgSnr) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsDataRate(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDataRate) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsTxBytes(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsTxBytes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsRxBytes(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsRxBytes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsDNSSuccess(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDNSSuccess) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsDNSFailure(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsDNSFailure) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboarding(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboarding) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["average_run_duration"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAverageRunDuration(item.AverageRunDuration)
	respItem["max_run_duration"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingMaxRunDuration(item.MaxRunDuration)
	respItem["average_assoc_duration"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAverageAssocDuration(item.AverageAssocDuration)
	respItem["max_assoc_duration"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingMaxAssocDuration(item.MaxAssocDuration)
	respItem["average_auth_duration"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAverageAuthDuration(item.AverageAuthDuration)
	respItem["max_auth_duration"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingMaxAuthDuration(item.MaxAuthDuration)
	respItem["average_dhcp_duration"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAverageDhcpDuration(item.AverageDhcpDuration)
	respItem["max_dhcp_duration"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingMaxDhcpDuration(item.MaxDhcpDuration)
	respItem["aaa_server_ip"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAAAServerIP(item.AAAServerIP)
	respItem["dhcp_server_ip"] = flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingDhcpServerIP(item.DhcpServerIP)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAverageRunDuration(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageRunDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingMaxRunDuration(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxRunDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAverageAssocDuration(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageAssocDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingMaxAssocDuration(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxAssocDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAverageAuthDuration(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageAuthDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingMaxAuthDuration(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxAuthDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAverageDhcpDuration(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAverageDhcpDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingMaxDhcpDuration(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingMaxDhcpDuration) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingAAAServerIP(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingAAAServerIP) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingDhcpServerIP(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingDhcpServerIP) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsOnboardingTime(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsOnboardingTime) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsUserDetailsPort(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1UserDetailsPort) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDevice(items *[]catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDevice) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_details"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetails(item.DeviceDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetails(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["family"] = item.Family
	respItem["type"] = item.Type
	respItem["location"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLocation(item.Location)
	respItem["error_code"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsErrorCode(item.ErrorCode)
	respItem["mac_address"] = item.MacAddress
	respItem["role"] = item.Role
	respItem["ap_manager_interface_ip"] = item.ApManagerInterfaceIP
	respItem["associated_wlc_ip"] = item.AssociatedWlcIP
	respItem["boot_date_time"] = item.BootDateTime
	respItem["collection_status"] = item.CollectionStatus
	respItem["interface_count"] = item.InterfaceCount
	respItem["line_card_count"] = item.LineCardCount
	respItem["line_card_id"] = item.LineCardID
	respItem["management_ip_address"] = item.ManagementIPAddress
	respItem["memory_size"] = item.MemorySize
	respItem["platform_id"] = item.PlatformID
	respItem["reachability_failure_reason"] = item.ReachabilityFailureReason
	respItem["reachability_status"] = item.ReachabilityStatus
	respItem["snmp_contact"] = item.SNMPContact
	respItem["snmp_location"] = item.SNMPLocation
	respItem["tunnel_udp_port"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsTunnelUDPPort(item.TunnelUDPPort)
	respItem["waas_device_mode"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsWaasDeviceMode(item.WaasDeviceMode)
	respItem["series"] = item.Series
	respItem["inventory_status_detail"] = item.InventoryStatusDetail
	respItem["collection_interval"] = item.CollectionInterval
	respItem["serial_number"] = item.SerialNumber
	respItem["software_version"] = item.SoftwareVersion
	respItem["role_source"] = item.RoleSource
	respItem["hostname"] = item.Hostname
	respItem["up_time"] = item.UpTime
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["error_description"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsErrorDescription(item.ErrorDescription)
	respItem["location_name"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLocationName(item.LocationName)
	respItem["tag_count"] = item.TagCount
	respItem["last_updated"] = item.LastUpdated
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["id"] = item.ID
	respItem["neighbor_topology"] = flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopology(item.NeighborTopology)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLocation(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocation) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsErrorCode(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorCode) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsTunnelUDPPort(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsTunnelUDPPort) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsWaasDeviceMode(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsWaasDeviceMode) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsErrorDescription(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsErrorDescription) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsLocationName(item *catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsLocationName) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenUsersGetUserEnrichmentDetailsV1ItemsConnectedDeviceDeviceDetailsNeighborTopology(items *[]catalystcentersdkgo.ResponseItemUsersGetUserEnrichmentDetailsV1ConnectedDeviceDeviceDetailsNeighborTopology) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["error_code"] = item.ErrorCode
		respItem["message"] = item.Message
		respItem["detail"] = item.Detail
		respItems = append(respItems, respItem)
	}
	return respItems
}
