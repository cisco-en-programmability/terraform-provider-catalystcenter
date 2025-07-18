package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceSensorTestCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Sensors.

- Intent API to create a SENSOR test template with a new SSID, existing SSID, or both new and existing SSID
`,

		CreateContext: resourceSensorTestCreateCreate,
		ReadContext:   resourceSensorTestCreateRead,
		DeleteContext: resourceSensorTestCreateDelete,
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

						"id": &schema.Schema{
							Description: `(Used in edit only) The sensor test template unique identifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"action_in_progress": &schema.Schema{
							Description: `Indication of inprogress action
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ap_coverage": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bands": &schema.Schema{
										Description: `The WIFI bands
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"number_of_aps_to_test": &schema.Schema{
										Description: `Number of APs to test
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"rssi_threshold": &schema.Schema{
										Description: `RSSI threshold
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"connection": &schema.Schema{
							Description: `connection type of test: WIRED, WIRELESS, BOTH
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"encryption_mode": &schema.Schema{
							Description: `Encryption mode
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"frequency": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"unit": &schema.Schema{
										Description: `Unit of the time value (NANOSECONDS, MICROSECONDS, MILLISECONDS, SECONDS, MINUTES, HOURS, DAYS)
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": &schema.Schema{
										Description: `Value of the unit
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"last_modified_time": &schema.Schema{
							Description: `Last modify time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"location": &schema.Schema{
							Description: `Location string
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"location_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensors": &schema.Schema{
										Description: `Use all sensors in the site for test
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"custom_management_vlan": &schema.Schema{
										Description: `Custom Management VLAN
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_type": &schema.Schema{
										Description: `Site type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mac_address_list": &schema.Schema{
										Description: `MAC addresses
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"management_vlan": &schema.Schema{
										Description: `Management VLAN
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site name hierarhy
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"model_version": &schema.Schema{
							Description: `Test template object model version (must be 2)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `The sensor test template name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"num_associated_sensor": &schema.Schema{
							Description: `Number of associated sensor
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"num_neighbor_apthreshold": &schema.Schema{
							Description: `Number of neighboring AP threshold
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"profiles": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"device_type": &schema.Schema{
										Description: `Device Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `External authentication portal
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_vlan_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"location_id": &schema.Schema{
													Description: `Site UUID
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"vlans": &schema.Schema{
													Description: `Array of VLANs
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
									"password": &schema.Schema{
										Description: `Password string for onboarding SSID
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `Profile name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"downlink_test": &schema.Schema{
																Description: `Downlink test
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"shared_secret": &schema.Schema{
																Description: `Shared secret
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name of the test
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `User name string for onboarding SSID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"vlan": &schema.Schema{
										Description: `VLAN
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"radio_as_sensor_removed": &schema.Schema{
							Description: `Radio as sensor removed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rssi_threshold": &schema.Schema{
							Description: `RSSI threshold
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"run_now": &schema.Schema{
							Description: `Run now (YES, NO)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"schedule_in_days": &schema.Schema{
							Description: `Bit-wise value of scheduled test days
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"sensors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensor_addition": &schema.Schema{
										Description: `Is all sensor addition
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"assigned": &schema.Schema{
										Description: `Is assigned
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"config_updated": &schema.Schema{
										Description: `Configuration updated: YES, NO
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"host_name": &schema.Schema{
										Description: `Host name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"i_perf_info": &schema.Schema{
										Description: `A string-stringList iPerf information
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Description: `Sensor ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Description: `IP address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `MAC address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"marked_for_uninstall": &schema.Schema{
										Description: `Is marked for uninstall
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Sensor name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"run_now": &schema.Schema{
										Description: `Run now: YES, NO
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"sensor_type": &schema.Schema{
										Description: `Sensor type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_policy": &schema.Schema{
										Description: `Service policy
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `Sensor device status: UP, DOWN, REBOOT
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"switch_mac": &schema.Schema{
										Description: `Switch MAC address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"switch_serial_number": &schema.Schema{
										Description: `Switch serial number
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"switch_uuid": &schema.Schema{
										Description: `Switch device UUID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"target_a_ps": &schema.Schema{
										Description: `Array of target APs
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"test_mac_addresses": &schema.Schema{
										Description: `A string-string test MAC address
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},
									"wired_application_message": &schema.Schema{
										Description: `Wired application message
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"wired_application_status": &schema.Schema{
										Description: `Wired application status
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"xor_sensor": &schema.Schema{
										Description: `Is XOR sensor
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"show_wlc_upgrade_banner": &schema.Schema{
							Description: `Show WLC upgrade banner
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_hierarchy": &schema.Schema{
							Description: `Site hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ssids": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"bands": &schema.Schema{
										Description: `WIFI bands: 2.4GHz or 5GHz
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `External authentication portal
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Identification number
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"layer3web_auth_email_address": &schema.Schema{
										Description: `Layer 3 WEB Auth email address
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"layer3web_authpassword": &schema.Schema{
										Description: `Layer 3 WEB Auth password
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"layer3web_authsecurity": &schema.Schema{
										Description: `Layer 3 WEB Auth security
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"layer3web_authuser_name": &schema.Schema{
										Description: `Layer 3 WEB Auth user name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"num_aps": &schema.Schema{
										Description: `Number of APs in the test
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"num_sensors": &schema.Schema{
										Description: `Number of Sensors in the test
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"password": &schema.Schema{
										Description: `Password string for onboarding SSID
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `The SSID profile name string
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_password": &schema.Schema{
										Description: `Proxy server password
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_port": &schema.Schema{
										Description: `Proxy server port
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_server": &schema.Schema{
										Description: `Proxy server for onboarding SSID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"proxy_user_name": &schema.Schema{
										Description: `Proxy server user name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"ssid": &schema.Schema{
										Description: `The SSID string
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `WLAN status: ENABLED or DISABLED
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"downlink_test": &schema.Schema{
																Description: `Downlink test
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"shared_secret": &schema.Schema{
																Description: `Shared secret
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name of the test
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"third_party": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"selected": &schema.Schema{
													Description: `true: the SSID is third party
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `User name string for onboarding SSID
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"valid_from": &schema.Schema{
										Description: `Valid From UTC timestamp
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"valid_to": &schema.Schema{
										Description: `Valid To UTC timestamp
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"wlan_id": &schema.Schema{
										Description: `WLAN ID
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"wlc": &schema.Schema{
										Description: `WLC IP addres
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"start_time": &schema.Schema{
							Description: `Start time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `Status of the test (RUNNING, NOTRUNNING)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"test_schedule_mode": &schema.Schema{
							Description: `Test schedule mode (ONDEMAND, DEDICATED, SCHEDULED, CONTINUOUS, RUNNOW)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
							Description: `The sensor test template version (must be 2)
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"wlans": &schema.Schema{
							Description: `WLANs list
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ap_coverage": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"bands": &schema.Schema{
										Description: `The WIFI bands
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"number_of_aps_to_test": &schema.Schema{
										Description: `Number of APs to test
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"rssi_threshold": &schema.Schema{
										Description: `RSSI threshold
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"connection": &schema.Schema{
							Description: `connection type of test: WIRED, WIRELESS, BOTH
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"encryption_mode": &schema.Schema{
							Description: `Encryption mode
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"location_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensors": &schema.Schema{
										Description: `Use all sensors in the site for test
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"custom_management_vlan": &schema.Schema{
										Description: `Custom Management VLAN
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"location_type": &schema.Schema{
										Description: `Site type
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mac_address_list": &schema.Schema{
										Description: `MAC addresses
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"management_vlan": &schema.Schema{
										Description: `Management VLAN
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"site_hierarchy": &schema.Schema{
										Description: `Site name hierarhy
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"model_version": &schema.Schema{
							Description: `Test template object model version (must be 2)
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `The sensor test template name
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"profiles": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"device_type": &schema.Schema{
										Description: `Device Type
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `External authentication portal
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"location_vlan_list": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"location_id": &schema.Schema{
													Description: `Site UUID
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"vlans": &schema.Schema{
													Description: `Array of VLANs
`,
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"password": &schema.Schema{
										Description: `Password string for onboarding SSID
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `Profile name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"downlink_test": &schema.Schema{
																Description: `Downlink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Optional:  true,
																ForceNew:  true,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Optional:  true,
																ForceNew:  true,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"shared_secret": &schema.Schema{
																Description: `Shared secret
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name of the test
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `User name string for onboarding SSID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"vlan": &schema.Schema{
										Description: `VLAN
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
								},
							},
						},
						"run_now": &schema.Schema{
							Description: `Run now (YES, NO)
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"sensors": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"all_sensor_addition": &schema.Schema{
										Description: `Is all sensor addition
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"assigned": &schema.Schema{
										Description: `Is assigned
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"config_updated": &schema.Schema{
										Description: `Configuration updated: YES, NO
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"host_name": &schema.Schema{
										Description: `Host name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"i_perf_info": &schema.Schema{
										Description: `A string-stringList iPerf information
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Description: `Sensor ID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ip_address": &schema.Schema{
										Description: `IP address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"location_id": &schema.Schema{
										Description: `Site UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"mac_address": &schema.Schema{
										Description: `MAC address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"marked_for_uninstall": &schema.Schema{
										Description: `Is marked for uninstall
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"name": &schema.Schema{
										Description: `Sensor name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"run_now": &schema.Schema{
										Description: `Run now: YES, NO
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"sensor_type": &schema.Schema{
										Description: `Sensor type
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"service_policy": &schema.Schema{
										Description: `Service policy
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"status": &schema.Schema{
										Description: `Sensor device status: UP, DOWN, REBOOT
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"switch_mac": &schema.Schema{
										Description: `Switch MAC address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"switch_serial_number": &schema.Schema{
										Description: `Switch serial number
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"switch_uuid": &schema.Schema{
										Description: `Switch device UUID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"target_a_ps": &schema.Schema{
										Description: `Array of target APs
`,
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"test_mac_addresses": &schema.Schema{
										Description: `A string-string test MAC address
`,
										Type:     schema.TypeString, //TEST,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"wired_application_message": &schema.Schema{
										Description: `Wired application message
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"wired_application_status": &schema.Schema{
										Description: `Wired application status
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"xor_sensor": &schema.Schema{
										Description: `Is XOR sensor
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
								},
							},
						},
						"ssids": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"auth_protocol": &schema.Schema{
										Description: `Auth protocol
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"auth_type": &schema.Schema{
										Description: `Authentication type: OPEN, WPA2_PSK, WPA2_EaP, WEB_AUTH, MAB, DOT1X, OTHER
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"bands": &schema.Schema{
										Description: `WIFI bands: 2.4GHz or 5GHz
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certdownloadurl": &schema.Schema{
										Description: `Certificate download URL
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certfilename": &schema.Schema{
										Description: `Auth certificate file name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certpassphrase": &schema.Schema{
										Description: `Certificate password phrase
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certstatus": &schema.Schema{
										Description: `Certificate status: INACTIVE or ACTIVE
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"certxferprotocol": &schema.Schema{
										Description: `Certificate transfering protocol: HTTP or HTTPS
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"eap_method": &schema.Schema{
										Description: `WPA2_EAP methods: EAP-FAST, PEAP-MSCHAPv2, EAP-TLS, PEAP-TLS, EAP-TTLS-MSCHAPv2, EAP-TTLS-PAP, EAP-TTLS-CHAP, EAP-FAST-GTC, EAP-PEAP-GTC
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth": &schema.Schema{
										Description: `Indication of using external WEB Auth
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"ext_web_auth_access_url": &schema.Schema{
										Description: `External WEB Auth access URL
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth_html_tag": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"label": &schema.Schema{
													Description: `Label`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"tag": &schema.Schema{
													Description: `Tag`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
									"ext_web_auth_portal": &schema.Schema{
										Description: `External authentication portal
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"ext_web_auth_virtual_ip": &schema.Schema{
										Description: `External WEB Auth virtual IP
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"layer3web_auth_email_address": &schema.Schema{
										Description: `Layer 3 WEB Auth email address
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"layer3web_authpassword": &schema.Schema{
										Description: `Layer 3 WEB Auth password
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"layer3web_authsecurity": &schema.Schema{
										Description: `Layer 3 WEB Auth security
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"layer3web_authuser_name": &schema.Schema{
										Description: `Layer 3 WEB Auth user name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"password": &schema.Schema{
										Description: `Password string for onboarding SSID
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"password_type": &schema.Schema{
										Description: `SSID password type: ASCII or HEX
`,
										Type:      schema.TypeString,
										Optional:  true,
										ForceNew:  true,
										Sensitive: true,
										Computed:  true,
									},
									"profile_name": &schema.Schema{
										Description: `The SSID profile name string
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy_password": &schema.Schema{
										Description: `Proxy server password
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy_port": &schema.Schema{
										Description: `Proxy server port
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy_server": &schema.Schema{
										Description: `Proxy server for onboarding SSID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"proxy_user_name": &schema.Schema{
										Description: `Proxy server user name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"psk": &schema.Schema{
										Description: `Password of SSID when passwordType is ASCII
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"qos_policy": &schema.Schema{
										Description: `QoS policy: PlATINUM, GOLD, SILVER, BRONZE
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"scep": &schema.Schema{
										Description: `Secure certificate enrollment protocol: true or false or null for not applicable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"ssid": &schema.Schema{
										Description: `The SSID string
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"tests": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"direction": &schema.Schema{
																Description: `IPerf direction (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"domains": &schema.Schema{
																Description: `DNS domain name
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"downlink_test": &schema.Schema{
																Description: `Downlink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"end_port": &schema.Schema{
																Description: `IPerf end port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"exit_command": &schema.Schema{
																Description: `Exit command
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"final_prompt": &schema.Schema{
																Description: `Final prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server": &schema.Schema{
																Description: `NDT server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server_path": &schema.Schema{
																Description: `NDT server path
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ndt_server_port": &schema.Schema{
																Description: `NDT server port
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"num_packets": &schema.Schema{
																Description: `Number of packets
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"password": &schema.Schema{
																Description: `Password
`,
																Type:      schema.TypeString,
																Optional:  true,
																ForceNew:  true,
																Sensitive: true,
																Computed:  true,
															},
															"password_prompt": &schema.Schema{
																Description: `Password prompt
`,
																Type:      schema.TypeString,
																Optional:  true,
																ForceNew:  true,
																Sensitive: true,
																Computed:  true,
															},
															"path_to_download": &schema.Schema{
																Description: `File path for file transfer
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"port": &schema.Schema{
																Description: `Radius or WEB server port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"probe_type": &schema.Schema{
																Description: `Probe type
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"protocol": &schema.Schema{
																Description: `Protocol used by file transfer, IPerf, mail server, and radius (TCP, UDP, FTP, POP3, IMAP, CHAP, PAP)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_password": &schema.Schema{
																Description: `Proxy password
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_port": &schema.Schema{
																Description: `Proxy port
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_server": &schema.Schema{
																Description: `Proxy server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"proxy_user_name": &schema.Schema{
																Description: `Proxy user name
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"server": &schema.Schema{
																Description: `Ping, file transfer, mail, radius, ssh, or telnet server
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"servers": &schema.Schema{
																Description: `IPerf server list
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"shared_secret": &schema.Schema{
																Description: `Shared secret
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"start_port": &schema.Schema{
																Description: `IPerf start port
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"transfer_type": &schema.Schema{
																Description: `File transfer type (UPLOAD, DOWNLOAD, BOTH)
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"udp_bandwidth": &schema.Schema{
																Description: `IPerf UDP bandwidth
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"uplink_test": &schema.Schema{
																Description: `Uplink test
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"url": &schema.Schema{
																Description: `URL
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"user_name": &schema.Schema{
																Description: `User name
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"user_name_prompt": &schema.Schema{
																Description: `User name prompt
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
												"name": &schema.Schema{
													Description: `Name of the test
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"third_party": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"selected": &schema.Schema{
													Description: `true: the SSID is third party
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													ForceNew:     true,
													Computed:     true,
												},
											},
										},
									},
									"username": &schema.Schema{
										Description: `User name string for onboarding SSID
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"white_list": &schema.Schema{
										Description: `Indication of being on allowed list
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
									"wlan_id": &schema.Schema{
										Description: `WLAN ID
`,
										Type:     schema.TypeInt,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"wlc": &schema.Schema{
										Description: `WLC IP addres
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Description: `The sensor test template version (must be 2)
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSensorTestCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestSensorTestCreateCreateSensorTestTemplate(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.Sensors.CreateSensorTestTemplate(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSensorTestTemplate", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenSensorsCreateSensorTestTemplateItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting CreateSensorTestTemplate response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceSensorTestCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceSensorTestCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func flattenSensorsCreateSensorTestTemplateItem(item *catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["type_id"] = item.TypeID
	respItem["version"] = item.Version
	respItem["model_version"] = item.ModelVersion
	respItem["start_time"] = item.StartTime
	respItem["last_modified_time"] = item.LastModifiedTime
	respItem["num_associated_sensor"] = item.NumAssociatedSensor
	respItem["location"] = item.Location
	respItem["site_hierarchy"] = item.SiteHierarchy
	respItem["status"] = item.Status
	respItem["connection"] = item.Connection
	respItem["action_in_progress"] = item.ActionInProgress
	respItem["frequency"] = flattenSensorsCreateSensorTestTemplateItemFrequency(item.Frequency)
	respItem["rssi_threshold"] = item.RssiThreshold
	respItem["num_neighbor_apthreshold"] = item.NumNeighborApThreshold
	respItem["schedule_in_days"] = item.ScheduleInDays
	respItem["wlans"] = item.WLANs
	respItem["ssids"] = flattenSensorsCreateSensorTestTemplateItemSSIDs(item.SSIDs)
	respItem["profiles"] = flattenSensorsCreateSensorTestTemplateItemProfiles(item.Profiles)
	respItem["test_schedule_mode"] = item.TestScheduleMode
	respItem["show_wlc_upgrade_banner"] = boolPtrToString(item.ShowWlcUpgradeBanner)
	respItem["radio_as_sensor_removed"] = boolPtrToString(item.RadioAsSensorRemoved)
	respItem["encryption_mode"] = item.EncryptionMode
	respItem["run_now"] = item.RunNow
	respItem["location_info_list"] = flattenSensorsCreateSensorTestTemplateItemLocationInfoList(item.LocationInfoList)
	respItem["sensors"] = flattenSensorsCreateSensorTestTemplateItemSensors(item.Sensors)
	respItem["ap_coverage"] = flattenSensorsCreateSensorTestTemplateItemApCoverage(item.ApCoverage)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSensorsCreateSensorTestTemplateItemFrequency(item *catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseFrequency) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["value"] = item.Value
	respItem["unit"] = item.Unit

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsCreateSensorTestTemplateItemSSIDs(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = item.Bands
		respItem["ssid"] = item.SSID
		respItem["profile_name"] = item.ProfileName
		respItem["num_aps"] = item.NumAps
		respItem["num_sensors"] = item.NumSensors
		respItem["layer3web_authsecurity"] = item.Layer3WebAuthsecurity
		respItem["layer3web_authuser_name"] = item.Layer3WebAuthuserName
		respItem["layer3web_authpassword"] = item.Layer3WebAuthpassword
		respItem["layer3web_auth_email_address"] = item.Layer3WebAuthEmailAddress
		respItem["third_party"] = flattenSensorsCreateSensorTestTemplateItemSSIDsThirdParty(item.ThirdParty)
		respItem["id"] = item.ID
		respItem["wlan_id"] = item.WLANID
		respItem["wlc"] = item.Wlc
		respItem["valid_from"] = item.ValidFrom
		respItem["valid_to"] = item.ValidTo
		respItem["status"] = item.Status
		respItem["proxy_server"] = item.ProxyServer
		respItem["proxy_port"] = item.ProxyPort
		respItem["proxy_user_name"] = item.ProxyUserName
		respItem["proxy_password"] = item.ProxyPassword
		respItem["auth_type"] = item.AuthType
		respItem["psk"] = item.Psk
		respItem["username"] = item.Username
		respItem["password"] = item.Password
		respItem["password_type"] = item.PasswordType
		respItem["eap_method"] = item.EapMethod
		respItem["scep"] = boolPtrToString(item.Scep)
		respItem["auth_protocol"] = item.AuthProtocol
		respItem["certfilename"] = item.Certfilename
		respItem["certxferprotocol"] = item.Certxferprotocol
		respItem["certstatus"] = item.Certstatus
		respItem["certpassphrase"] = item.Certpassphrase
		respItem["certdownloadurl"] = item.Certdownloadurl
		respItem["ext_web_auth_virtual_ip"] = item.ExtWebAuthVirtualIP
		respItem["ext_web_auth"] = boolPtrToString(item.ExtWebAuth)
		respItem["white_list"] = boolPtrToString(item.WhiteList)
		respItem["ext_web_auth_portal"] = item.ExtWebAuthPortal
		respItem["ext_web_auth_access_url"] = item.ExtWebAuthAccessURL
		respItem["ext_web_auth_html_tag"] = flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["qos_policy"] = item.QosPolicy
		respItem["tests"] = flattenSensorsCreateSensorTestTemplateItemSSIDsTests(item.Tests)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemSSIDsThirdParty(item *catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsThirdParty) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["selected"] = boolPtrToString(item.Selected)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSensorsCreateSensorTestTemplateItemSSIDsExtWebAuthHTMLTag(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsExtWebAuthHTMLTag) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["label"] = item.Label
		respItem["tag"] = item.Tag
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemSSIDsTests(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsCreateSensorTestTemplateItemSSIDsTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemSSIDsTestsConfig(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSSIDsTestsConfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["domains"] = item.Domains
		respItem["server"] = item.Server
		respItem["user_name"] = item.UserName
		respItem["password"] = item.Password
		respItem["url"] = item.URL
		respItem["port"] = item.Port
		respItem["protocol"] = item.Protocol
		respItem["servers"] = item.Servers
		respItem["direction"] = item.Direction
		respItem["start_port"] = item.StartPort
		respItem["end_port"] = item.EndPort
		respItem["udp_bandwidth"] = item.UDPBandwidth
		respItem["probe_type"] = item.ProbeType
		respItem["num_packets"] = item.NumPackets
		respItem["path_to_download"] = item.PathToDownload
		respItem["transfer_type"] = item.TransferType
		respItem["shared_secret"] = item.SharedSecret
		respItem["ndt_server"] = item.NdtServer
		respItem["ndt_server_port"] = item.NdtServerPort
		respItem["ndt_server_path"] = item.NdtServerPath
		respItem["uplink_test"] = boolPtrToString(item.UplinkTest)
		respItem["downlink_test"] = boolPtrToString(item.DownlinkTest)
		respItem["proxy_server"] = item.ProxyServer
		respItem["proxy_port"] = item.ProxyPort
		respItem["proxy_user_name"] = item.ProxyUserName
		respItem["proxy_password"] = item.ProxyPassword
		respItem["user_name_prompt"] = item.UserNamePrompt
		respItem["password_prompt"] = item.PasswordPrompt
		respItem["exit_command"] = item.ExitCommand
		respItem["final_prompt"] = item.FinalPrompt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemProfiles(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseProfiles) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["auth_type"] = item.AuthType
		respItem["psk"] = item.Psk
		respItem["username"] = item.Username
		respItem["password"] = item.Password
		respItem["password_type"] = item.PasswordType
		respItem["eap_method"] = item.EapMethod
		respItem["scep"] = boolPtrToString(item.Scep)
		respItem["auth_protocol"] = item.AuthProtocol
		respItem["certfilename"] = item.Certfilename
		respItem["certxferprotocol"] = item.Certxferprotocol
		respItem["certstatus"] = item.Certstatus
		respItem["certpassphrase"] = item.Certpassphrase
		respItem["certdownloadurl"] = item.Certdownloadurl
		respItem["ext_web_auth_virtual_ip"] = item.ExtWebAuthVirtualIP
		respItem["ext_web_auth"] = boolPtrToString(item.ExtWebAuth)
		respItem["white_list"] = boolPtrToString(item.WhiteList)
		respItem["ext_web_auth_portal"] = item.ExtWebAuthPortal
		respItem["ext_web_auth_access_url"] = item.ExtWebAuthAccessURL
		respItem["ext_web_auth_html_tag"] = flattenSensorsCreateSensorTestTemplateItemProfilesExtWebAuthHTMLTag(item.ExtWebAuthHTMLTag)
		respItem["qos_policy"] = item.QosPolicy
		respItem["tests"] = flattenSensorsCreateSensorTestTemplateItemProfilesTests(item.Tests)
		respItem["profile_name"] = item.ProfileName
		respItem["device_type"] = item.DeviceType
		respItem["vlan"] = item.VLAN
		respItem["location_vlan_list"] = flattenSensorsCreateSensorTestTemplateItemProfilesLocationVLANList(item.LocationVLANList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemProfilesExtWebAuthHTMLTag(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseProfilesExtWebAuthHTMLTag) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["label"] = item.Label
		respItem["tag"] = item.Tag
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemProfilesTests(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseProfilesTests) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["config"] = flattenSensorsCreateSensorTestTemplateItemProfilesTestsConfig(item.Config)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemProfilesTestsConfig(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseProfilesTestsConfig) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["domains"] = item.Domains
		respItem["server"] = item.Server
		respItem["user_name"] = item.UserName
		respItem["password"] = item.Password
		respItem["url"] = item.URL
		respItem["port"] = item.Port
		respItem["protocol"] = item.Protocol
		respItem["servers"] = item.Servers
		respItem["direction"] = item.Direction
		respItem["start_port"] = item.StartPort
		respItem["end_port"] = item.EndPort
		respItem["udp_bandwidth"] = item.UDPBandwidth
		respItem["probe_type"] = item.ProbeType
		respItem["num_packets"] = item.NumPackets
		respItem["path_to_download"] = item.PathToDownload
		respItem["transfer_type"] = item.TransferType
		respItem["shared_secret"] = item.SharedSecret
		respItem["ndt_server"] = item.NdtServer
		respItem["ndt_server_port"] = item.NdtServerPort
		respItem["ndt_server_path"] = item.NdtServerPath
		respItem["uplink_test"] = boolPtrToString(item.UplinkTest)
		respItem["downlink_test"] = boolPtrToString(item.DownlinkTest)
		respItem["proxy_server"] = item.ProxyServer
		respItem["proxy_port"] = item.ProxyPort
		respItem["proxy_user_name"] = item.ProxyUserName
		respItem["proxy_password"] = item.ProxyPassword
		respItem["user_name_prompt"] = item.UserNamePrompt
		respItem["password_prompt"] = item.PasswordPrompt
		respItem["exit_command"] = item.ExitCommand
		respItem["final_prompt"] = item.FinalPrompt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemProfilesLocationVLANList(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseProfilesLocationVLANList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["location_id"] = item.LocationID
		respItem["vlans"] = item.VLANs
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemLocationInfoList(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseLocationInfoList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["location_id"] = item.LocationID
		respItem["location_type"] = item.LocationType
		respItem["all_sensors"] = boolPtrToString(item.AllSensors)
		respItem["site_hierarchy"] = item.SiteHierarchy
		respItem["mac_address_list"] = item.MacAddressList
		respItem["management_vlan"] = item.ManagementVLAN
		respItem["custom_management_vlan"] = boolPtrToString(item.CustomManagementVLAN)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemSensors(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSensors) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["mac_address"] = item.MacAddress
		respItem["switch_mac"] = item.SwitchMac
		respItem["switch_uuid"] = item.SwitchUUID
		respItem["switch_serial_number"] = item.SwitchSerialNumber
		respItem["marked_for_uninstall"] = boolPtrToString(item.MarkedForUninstall)
		respItem["ip_address"] = item.IPAddress
		respItem["host_name"] = item.HostName
		respItem["wired_application_status"] = item.WiredApplicationStatus
		respItem["wired_application_message"] = item.WiredApplicationMessage
		respItem["assigned"] = boolPtrToString(item.Assigned)
		respItem["status"] = item.Status
		respItem["xor_sensor"] = boolPtrToString(item.XorSensor)
		respItem["target_a_ps"] = item.TargetAPs
		respItem["run_now"] = item.RunNow
		respItem["location_id"] = item.LocationID
		respItem["all_sensor_addition"] = boolPtrToString(item.AllSensorAddition)
		respItem["config_updated"] = item.ConfigUpdated
		respItem["sensor_type"] = item.SensorType
		respItem["test_mac_addresses"] = flattenSensorsCreateSensorTestTemplateItemSensorsTestMacAddresses(item.TestMacAddresses)
		respItem["id"] = item.ID
		respItem["service_policy"] = item.ServicePolicy
		respItem["i_perf_info"] = flattenSensorsCreateSensorTestTemplateItemSensorsIPerfInfo(item.IPerfInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSensorsCreateSensorTestTemplateItemSensorsTestMacAddresses(item *catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSensorsTestMacAddresses) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemSensorsIPerfInfo(item *catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseSensorsIPerfInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSensorsCreateSensorTestTemplateItemApCoverage(items *[]catalystcentersdkgo.ResponseSensorsCreateSensorTestTemplateResponseApCoverage) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["bands"] = item.Bands
		respItem["number_of_aps_to_test"] = item.NumberOfApsToTest
		respItem["rssi_threshold"] = item.RssiThreshold
		respItems = append(respItems, respItem)
	}
	return respItems
}

func expandRequestSensorTestCreateCreateSensorTestTemplate(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplate {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".model_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".model_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".model_version")))) {
		request.ModelVersion = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection")))) {
		request.Connection = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssids")))) {
		request.SSIDs = expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsArray(ctx, key+".ssids", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profiles")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profiles")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profiles")))) {
		request.Profiles = expandRequestSensorTestCreateCreateSensorTestTemplateProfilesArray(ctx, key+".profiles", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".encryption_mode")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".encryption_mode")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".encryption_mode")))) {
		request.EncryptionMode = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_now")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_now")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_now")))) {
		request.RunNow = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_info_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_info_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_info_list")))) {
		request.LocationInfoList = expandRequestSensorTestCreateCreateSensorTestTemplateLocationInfoListArray(ctx, key+".location_info_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sensors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sensors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sensors")))) {
		request.Sensors = expandRequestSensorTestCreateCreateSensorTestTemplateSensorsArray(ctx, key+".sensors", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ap_coverage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ap_coverage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ap_coverage")))) {
		request.ApCoverage = expandRequestSensorTestCreateCreateSensorTestTemplateApCoverageArray(ctx, key+".ap_coverage", d)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateSSIDs(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDs(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDs{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bands")))) {
		request.Bands = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3web_authsecurity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3web_authsecurity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3web_authsecurity")))) {
		request.Layer3WebAuthsecurity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3web_authuser_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3web_authuser_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3web_authuser_name")))) {
		request.Layer3WebAuthuserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3web_authpassword")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3web_authpassword")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3web_authpassword")))) {
		request.Layer3WebAuthpassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".layer3web_auth_email_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".layer3web_auth_email_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".layer3web_auth_email_address")))) {
		request.Layer3WebAuthEmailAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".third_party")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".third_party")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".third_party")))) {
		request.ThirdParty = expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsThirdParty(ctx, key+".third_party.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlan_id")))) {
		request.WLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wlc")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wlc")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wlc")))) {
		request.Wlc = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_server")))) {
		request.ProxyServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_port")))) {
		request.ProxyPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_user_name")))) {
		request.ProxyUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_password")))) {
		request.ProxyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psk")))) {
		request.Psk = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_type")))) {
		request.PasswordType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_method")))) {
		request.EapMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scep")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scep")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scep")))) {
		request.Scep = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_protocol")))) {
		request.AuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certfilename")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certfilename")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certfilename")))) {
		request.Certfilename = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certxferprotocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certxferprotocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certxferprotocol")))) {
		request.Certxferprotocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certstatus")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certstatus")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certstatus")))) {
		request.Certstatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certpassphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certpassphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certpassphrase")))) {
		request.Certpassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certdownloadurl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certdownloadurl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certdownloadurl")))) {
		request.Certdownloadurl = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_virtual_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_virtual_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_virtual_ip")))) {
		request.ExtWebAuthVirtualIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth")))) {
		request.ExtWebAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".white_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".white_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".white_list")))) {
		request.WhiteList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_portal")))) {
		request.ExtWebAuthPortal = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_access_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_access_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_access_url")))) {
		request.ExtWebAuthAccessURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_html_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_html_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_html_tag")))) {
		request.ExtWebAuthHTMLTag = expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsExtWebAuthHTMLTagArray(ctx, key+".ext_web_auth_html_tag", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_policy")))) {
		request.QosPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tests")))) {
		request.Tests = expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsTestsArray(ctx, key+".tests", d)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsThirdParty(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsThirdParty {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsThirdParty{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selected")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selected")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selected")))) {
		request.Selected = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsExtWebAuthHTMLTagArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsExtWebAuthHTMLTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag")))) {
		request.Tag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsTestsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsTests(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsTests(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTests{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config")))) {
		request.Config = expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsTestsConfigArray(ctx, key+".config", d)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsTestsConfigArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsTestsConfig(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSSIDsTestsConfig(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSSIDsTestsConfig{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains")))) {
		request.Domains = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server")))) {
		request.Server = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".servers")))) {
		request.Servers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".direction")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".direction")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".direction")))) {
		request.Direction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_port")))) {
		request.StartPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_port")))) {
		request.EndPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".udp_bandwidth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".udp_bandwidth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".udp_bandwidth")))) {
		request.UDPBandwidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".probe_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".probe_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".probe_type")))) {
		request.ProbeType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_packets")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_packets")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_packets")))) {
		request.NumPackets = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".path_to_download")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".path_to_download")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".path_to_download")))) {
		request.PathToDownload = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transfer_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transfer_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transfer_type")))) {
		request.TransferType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server")))) {
		request.NdtServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server_port")))) {
		request.NdtServerPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server_path")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server_path")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server_path")))) {
		request.NdtServerPath = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".uplink_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".uplink_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".uplink_test")))) {
		request.UplinkTest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".downlink_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".downlink_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".downlink_test")))) {
		request.DownlinkTest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_server")))) {
		request.ProxyServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_port")))) {
		request.ProxyPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_user_name")))) {
		request.ProxyUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_password")))) {
		request.ProxyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_prompt")))) {
		request.UserNamePrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_prompt")))) {
		request.PasswordPrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exit_command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exit_command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exit_command")))) {
		request.ExitCommand = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".final_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".final_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".final_prompt")))) {
		request.FinalPrompt = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfiles {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfiles{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateProfiles(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfiles(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfiles {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfiles{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_type")))) {
		request.AuthType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".psk")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".psk")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".psk")))) {
		request.Psk = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_type")))) {
		request.PasswordType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".eap_method")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".eap_method")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".eap_method")))) {
		request.EapMethod = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scep")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scep")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scep")))) {
		request.Scep = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_protocol")))) {
		request.AuthProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certfilename")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certfilename")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certfilename")))) {
		request.Certfilename = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certxferprotocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certxferprotocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certxferprotocol")))) {
		request.Certxferprotocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certstatus")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certstatus")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certstatus")))) {
		request.Certstatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certpassphrase")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certpassphrase")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certpassphrase")))) {
		request.Certpassphrase = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".certdownloadurl")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".certdownloadurl")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".certdownloadurl")))) {
		request.Certdownloadurl = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_virtual_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_virtual_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_virtual_ip")))) {
		request.ExtWebAuthVirtualIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth")))) {
		request.ExtWebAuth = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".white_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".white_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".white_list")))) {
		request.WhiteList = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_portal")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_portal")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_portal")))) {
		request.ExtWebAuthPortal = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_access_url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_access_url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_access_url")))) {
		request.ExtWebAuthAccessURL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ext_web_auth_html_tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ext_web_auth_html_tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ext_web_auth_html_tag")))) {
		request.ExtWebAuthHTMLTag = expandRequestSensorTestCreateCreateSensorTestTemplateProfilesExtWebAuthHTMLTagArray(ctx, key+".ext_web_auth_html_tag", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_policy")))) {
		request.QosPolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tests")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tests")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tests")))) {
		request.Tests = expandRequestSensorTestCreateCreateSensorTestTemplateProfilesTestsArray(ctx, key+".tests", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".profile_name")))) {
		request.ProfileName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_type")))) {
		request.DeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan")))) {
		request.VLAN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_vlan_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_vlan_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_vlan_list")))) {
		request.LocationVLANList = expandRequestSensorTestCreateCreateSensorTestTemplateProfilesLocationVLANListArray(ctx, key+".location_vlan_list", d)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesExtWebAuthHTMLTagArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateProfilesExtWebAuthHTMLTag(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesExtWebAuthHTMLTag(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesExtWebAuthHTMLTag{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tag")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tag")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tag")))) {
		request.Tag = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesTestsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTests {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTests{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateProfilesTests(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesTests(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTests {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTests{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config")))) {
		request.Config = expandRequestSensorTestCreateCreateSensorTestTemplateProfilesTestsConfigArray(ctx, key+".config", d)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesTestsConfigArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTestsConfig {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTestsConfig{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateProfilesTestsConfig(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesTestsConfig(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTestsConfig {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesTestsConfig{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains")))) {
		request.Domains = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server")))) {
		request.Server = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name")))) {
		request.UserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".servers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".servers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".servers")))) {
		request.Servers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".direction")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".direction")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".direction")))) {
		request.Direction = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".start_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".start_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".start_port")))) {
		request.StartPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".end_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".end_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".end_port")))) {
		request.EndPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".udp_bandwidth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".udp_bandwidth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".udp_bandwidth")))) {
		request.UDPBandwidth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".probe_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".probe_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".probe_type")))) {
		request.ProbeType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".num_packets")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".num_packets")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".num_packets")))) {
		request.NumPackets = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".path_to_download")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".path_to_download")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".path_to_download")))) {
		request.PathToDownload = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".transfer_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".transfer_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".transfer_type")))) {
		request.TransferType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".shared_secret")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".shared_secret")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".shared_secret")))) {
		request.SharedSecret = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server")))) {
		request.NdtServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server_port")))) {
		request.NdtServerPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ndt_server_path")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ndt_server_path")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ndt_server_path")))) {
		request.NdtServerPath = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".uplink_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".uplink_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".uplink_test")))) {
		request.UplinkTest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".downlink_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".downlink_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".downlink_test")))) {
		request.DownlinkTest = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_server")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_server")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_server")))) {
		request.ProxyServer = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_port")))) {
		request.ProxyPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_user_name")))) {
		request.ProxyUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".proxy_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".proxy_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".proxy_password")))) {
		request.ProxyPassword = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_name_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_name_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_name_prompt")))) {
		request.UserNamePrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password_prompt")))) {
		request.PasswordPrompt = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exit_command")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exit_command")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exit_command")))) {
		request.ExitCommand = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".final_prompt")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".final_prompt")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".final_prompt")))) {
		request.FinalPrompt = interfaceToString(v)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesLocationVLANListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateProfilesLocationVLANList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateProfilesLocationVLANList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateProfilesLocationVLANList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_id")))) {
		request.LocationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlans")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlans")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlans")))) {
		request.VLANs = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateLocationInfoListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateLocationInfoList {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateLocationInfoList{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateLocationInfoList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateLocationInfoList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateLocationInfoList {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateLocationInfoList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_id")))) {
		request.LocationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_type")))) {
		request.LocationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".all_sensors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".all_sensors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".all_sensors")))) {
		request.AllSensors = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_hierarchy")))) {
		request.SiteHierarchy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address_list")))) {
		request.MacAddressList = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".management_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".management_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".management_vlan")))) {
		request.ManagementVLAN = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_management_vlan")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_management_vlan")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_management_vlan")))) {
		request.CustomManagementVLAN = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSensorsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSensors {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSensors{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateSensors(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSensors(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSensors {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSensors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".switch_mac")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".switch_mac")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".switch_mac")))) {
		request.SwitchMac = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".switch_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".switch_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".switch_uuid")))) {
		request.SwitchUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".switch_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".switch_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".switch_serial_number")))) {
		request.SwitchSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".marked_for_uninstall")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".marked_for_uninstall")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".marked_for_uninstall")))) {
		request.MarkedForUninstall = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_address")))) {
		request.IPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_name")))) {
		request.HostName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wired_application_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wired_application_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wired_application_status")))) {
		request.WiredApplicationStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".wired_application_message")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".wired_application_message")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".wired_application_message")))) {
		request.WiredApplicationMessage = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".assigned")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".assigned")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".assigned")))) {
		request.Assigned = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".status")))) {
		request.Status = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".xor_sensor")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".xor_sensor")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".xor_sensor")))) {
		request.XorSensor = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".target_a_ps")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".target_a_ps")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".target_a_ps")))) {
		request.TargetAPs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".run_now")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".run_now")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".run_now")))) {
		request.RunNow = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".location_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".location_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".location_id")))) {
		request.LocationID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".all_sensor_addition")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".all_sensor_addition")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".all_sensor_addition")))) {
		request.AllSensorAddition = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_updated")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_updated")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_updated")))) {
		request.ConfigUpdated = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sensor_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sensor_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sensor_type")))) {
		request.SensorType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".test_mac_addresses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".test_mac_addresses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".test_mac_addresses")))) {
		request.TestMacAddresses = expandRequestSensorTestCreateCreateSensorTestTemplateSensorsTestMacAddresses(ctx, key+".test_mac_addresses.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".service_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".service_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".service_policy")))) {
		request.ServicePolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".i_perf_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".i_perf_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".i_perf_info")))) {
		request.IPerfInfo = expandRequestSensorTestCreateCreateSensorTestTemplateSensorsIPerfInfo(ctx, key+".i_perf_info.0", d)
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSensorsTestMacAddresses(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSensorsTestMacAddresses {
	var request catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSensorsTestMacAddresses
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateSensorsIPerfInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSensorsIPerfInfo {
	var request catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateSensorsIPerfInfo
	request = d.Get(fixKeyAccess(key))
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateApCoverageArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage {
	request := []catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage{}
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
		i := expandRequestSensorTestCreateCreateSensorTestTemplateApCoverage(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestSensorTestCreateCreateSensorTestTemplateApCoverage(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage {
	request := catalystcentersdkgo.RequestSensorsCreateSensorTestTemplateApCoverage{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bands")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bands")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bands")))) {
		request.Bands = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".number_of_aps_to_test")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".number_of_aps_to_test")))) {
		request.NumberOfApsToTest = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rssi_threshold")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rssi_threshold")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rssi_threshold")))) {
		request.RssiThreshold = interfaceToIntPtr(v)
	}
	return &request
}
