package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsRfProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get all RF Profiles

- This data source allows the user to get a RF Profile by RF Profile ID
`,

		ReadContext: dataSourceWirelessSettingsRfProfilesRead,
		Schema: map[string]*schema.Schema{
			"enable_radio_type6_g_hz": &schema.Schema{
				Description: `enableRadioType6GHz query parameter. Enable Radio Type6GHz
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_radio_type_a": &schema.Schema{
				Description: `enableRadioTypeA query parameter. Enable Radio TypeA
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_radio_type_b": &schema.Schema{
				Description: `enableRadioTypeB query parameter. Enable Radio TypeB
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. RF Profile ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"rf_profile_name": &schema.Schema{
				Description: `rfProfileName query parameter. RF Profile Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"default_rf_profile": &schema.Schema{
							Description: `True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_custom": &schema.Schema{
							Description: `True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_radio_type6_g_hz": &schema.Schema{
							Description: `True if 6 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_radio_type_a": &schema.Schema{
							Description: `True if 5 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_radio_type_b": &schema.Schema{
							Description: `True if 2.4 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `RF Profile ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"radio_type6_g_hz_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"broadcast_probe_response_interval": &schema.Schema{
										Description: `Broadcast Probe Response Interval of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"data_rates": &schema.Schema{
										Description: `Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"discovery_frames6_g_hz": &schema.Schema{
										Description: `Discovery Frames of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_standard_power_service": &schema.Schema{
										Description: `True if Standard Power Service is enabled, else False
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fra_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_reset_count": &schema.Schema{
													Description: `Client Reset Count of 6 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"client_utilization_threshold": &schema.Schema{
													Description: `Client Utilization Threshold of 6 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_dbs_width": &schema.Schema{
										Description: `Maximum DBS Width (Permissible Values: 20,40,80,160,320)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_dbs_width": &schema.Schema{
										Description: `Minimum DBS Width ( Permissible values : 20,40,80,160,320)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"multi_bssid_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mu_mimo_down_link": &schema.Schema{
																Description: `MU-MIMO Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"mu_mimo_up_link": &schema.Schema{
																Description: `MU-MIMO Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_down_link": &schema.Schema{
																Description: `OFDMA Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_up_link": &schema.Schema{
																Description: `OFDMA Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"dot11be_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mu_mimo_down_link": &schema.Schema{
																Description: `MU-MIMO Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"mu_mimo_up_link": &schema.Schema{
																Description: `MU-MIMO Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_down_link": &schema.Schema{
																Description: `OFDMA Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_multi_ru": &schema.Schema{
																Description: `OFDMA Multi-RU
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_up_link": &schema.Schema{
																Description: `OFDMA Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"target_wake_time": &schema.Schema{
													Description: `Target Wake Time
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"twt_broadcast_support": &schema.Schema{
													Description: `TWT Broadcast Support
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"parent_profile": &schema.Schema{
										Description: `Parent profile of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"preamble_puncture": &schema.Schema{
										Description: `Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"psc_enforcing_enabled": &schema.Schema{
										Description: `PSC Enforcing Enable for 6 GHz radio band
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"radio_channels": &schema.Schema{
										Description: `DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"radio_type_a_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"channel_width": &schema.Schema{
										Description: `Channel Width
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"data_rates": &schema.Schema{
										Description: `Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fra_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_aware": &schema.Schema{
													Description: `Client Aware of 5 GHz radio band
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"client_reset": &schema.Schema{
													Description: `Client Reset(%) of 5 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"client_select": &schema.Schema{
													Description: `Client Select(%) of 5 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent profile of 5 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"preamble_puncture": &schema.Schema{
										Description: `Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"radio_channels": &schema.Schema{
										Description: `DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 5 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"zero_wait_dfs_enable": &schema.Schema{
										Description: `Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"radio_type_b_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"data_rates": &schema.Schema{
										Description: `Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent profile of 2.4 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"radio_channels": &schema.Schema{
										Description: `DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 2.4 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"rf_profile_name": &schema.Schema{
							Description: `RF Profile Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"default_rf_profile": &schema.Schema{
							Description: `True if RF Profile is default, else False. Maximum of only 1 RF Profile can be marked as default at any given time
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_custom": &schema.Schema{
							Description: `True if RF Profile is custom, else False for system RF profiles like Low, High and Medium (Typical)
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_radio_type6_g_hz": &schema.Schema{
							Description: `True if 6 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_radio_type_a": &schema.Schema{
							Description: `True if 5 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"enable_radio_type_b": &schema.Schema{
							Description: `True if 2.4 GHz radio band is enabled in the RF Profile, else False
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `RF Profile ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"radio_type6_g_hz_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"broadcast_probe_response_interval": &schema.Schema{
										Description: `Broadcast Probe Response Interval of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"data_rates": &schema.Schema{
										Description: `Data rates of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"discovery_frames6_g_hz": &schema.Schema{
										Description: `Discovery Frames of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_standard_power_service": &schema.Schema{
										Description: `True if Standard Power Service is enabled, else False
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fra_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_reset_count": &schema.Schema{
													Description: `Client Reset Count of 6 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"client_utilization_threshold": &schema.Schema{
													Description: `Client Utilization Threshold of 6 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 6 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_dbs_width": &schema.Schema{
										Description: `Maximum DBS Width (Permissible Values: 20,40,80,160,320)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_dbs_width": &schema.Schema{
										Description: `Minimum DBS Width ( Permissible values : 20,40,80,160,320)
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"multi_bssid_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mu_mimo_down_link": &schema.Schema{
																Description: `MU-MIMO Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"mu_mimo_up_link": &schema.Schema{
																Description: `MU-MIMO Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_down_link": &schema.Schema{
																Description: `OFDMA Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_up_link": &schema.Schema{
																Description: `OFDMA Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"dot11be_parameters": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"mu_mimo_down_link": &schema.Schema{
																Description: `MU-MIMO Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"mu_mimo_up_link": &schema.Schema{
																Description: `MU-MIMO Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_down_link": &schema.Schema{
																Description: `OFDMA Downlink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_multi_ru": &schema.Schema{
																Description: `OFDMA Multi-RU
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"ofdma_up_link": &schema.Schema{
																Description: `OFDMA Uplink
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"target_wake_time": &schema.Schema{
													Description: `Target Wake Time
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"twt_broadcast_support": &schema.Schema{
													Description: `TWT Broadcast Support
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"parent_profile": &schema.Schema{
										Description: `Parent profile of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 6 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"preamble_puncture": &schema.Schema{
										Description: `Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"psc_enforcing_enabled": &schema.Schema{
										Description: `PSC Enforcing Enable for 6 GHz radio band
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"radio_channels": &schema.Schema{
										Description: `DCA channels of 6 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 5, 9, 13, 17, 21, 25, 29, 33, 37, 41, 45, 49, 53, 57, 61, 65, 69, 73, 77, 81, 85, 89, 93, 97, 101, 105, 109, 113, 117, 121, 125, 129, 133, 137, 141, 145, 149, 153, 157, 161, 165, 169, 173, 177, 181, 185, 189, 193, 197, 201, 205, 209, 213, 217, 221, 225, 229, 233
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 6 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"radio_type_a_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"channel_width": &schema.Schema{
										Description: `Channel Width
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"data_rates": &schema.Schema{
										Description: `Data rates of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"fra_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"client_aware": &schema.Schema{
													Description: `Client Aware of 5 GHz radio band
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"client_reset": &schema.Schema{
													Description: `Client Reset(%) of 5 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"client_select": &schema.Schema{
													Description: `Client Select(%) of 5 GHz radio band
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 5 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 6, 9, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent profile of 5 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 5 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"preamble_puncture": &schema.Schema{
										Description: `Enable or Disable Preamble Puncturing. This Wifi 7 configuration is applicable to wireless IOS devices supporting 17.15 and higher
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"radio_channels": &schema.Schema{
										Description: `DCA channels of 5 GHz radio band passed in comma separated format without any spaces. Permissible values: 36, 40, 44, 48, 52, 56, 60, 64, 100, 104, 108, 112, 116, 120, 124, 128, 132, 136, 140, 144, 149, 153, 157, 161, 165, 169, 173
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 5 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"zero_wait_dfs_enable": &schema.Schema{
										Description: `Zero Wait DFS is applicable only for IOS-XE based Wireless Controllers running 17.9.1 and above versions
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"radio_type_b_properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"coverage_hole_detection_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"chd_client_level": &schema.Schema{
													Description: `Coverage Hole Detection Client Level
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_data_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Data Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_exception_level": &schema.Schema{
													Description: `Coverage Hole Detection Exception Level(%)
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"chd_voice_rssi_threshold": &schema.Schema{
													Description: `Coverage Hole Detection Voice Rssi Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"custom_rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold custom configuration of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"data_rates": &schema.Schema{
										Description: `Data rates of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mandatory_data_rates": &schema.Schema{
										Description: `Mandatory data rates of 2.4 GHz radio band passed in comma separated format without any spaces and must be a subset of selected dataRates with maximum of 2 values. Permissible values: 1, 2, 5.5, 6, 9, 11, 12, 18, 24, 36, 48, 54
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"max_power_level": &schema.Schema{
										Description: `Maximum power level of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"max_radio_clients": &schema.Schema{
										Description: `Client Limit of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"min_power_level": &schema.Schema{
										Description: `Minimum power level of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"parent_profile": &schema.Schema{
										Description: `Parent profile of 2.4 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"power_threshold_v1": &schema.Schema{
										Description: `Power threshold of 2.4 GHz radio band
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"radio_channels": &schema.Schema{
										Description: `DCA channels of 2.4 GHz radio band passed in comma separated format without any spaces. Permissible values: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rx_sop_threshold": &schema.Schema{
										Description: `RX-SOP threshold of 2.4 GHz radio band
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"spatial_reuse_properties": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dot11ax_non_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_non_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax Non SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_max_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Max Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dot11ax_srg_obss_packet_detect_min_threshold": &schema.Schema{
													Description: `Dot11ax SRG OBSS PD Min Threshold
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"rf_profile_name": &schema.Schema{
							Description: `RF Profile Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessSettingsRfProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vRfProfileName, okRfProfileName := d.GetOk("rf_profile_name")
	vEnableRadioTypeA, okEnableRadioTypeA := d.GetOk("enable_radio_type_a")
	vEnableRadioTypeB, okEnableRadioTypeB := d.GetOk("enable_radio_type_b")
	vEnableRadioType6GHz, okEnableRadioType6GHz := d.GetOk("enable_radio_type6_g_hz")
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset, okRfProfileName, okEnableRadioTypeA, okEnableRadioTypeB, okEnableRadioType6GHz}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetRfProfilesV1")
		queryParams1 := catalystcentersdkgo.GetRfProfilesV1QueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okRfProfileName {
			queryParams1.RfProfileName = vRfProfileName.(string)
		}
		if okEnableRadioTypeA {
			queryParams1.EnableRadioTypeA = vEnableRadioTypeA.(bool)
		}
		if okEnableRadioTypeB {
			queryParams1.EnableRadioTypeB = vEnableRadioTypeB.(bool)
		}
		if okEnableRadioType6GHz {
			queryParams1.EnableRadioType6GHz = vEnableRadioType6GHz.(bool)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.GetRfProfilesV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetRfProfilesV1", err,
				"Failure at GetRfProfilesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetRfProfilesV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRfProfilesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetRfProfileByIDV1")
		vvID := vID.(string)

		// has_unknown_response: None

		response2, restyResp2, err := client.Wireless.GetRfProfileByIDV1(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetRfProfileByIDV1", err,
				"Failure at GetRfProfileByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenWirelessGetRfProfileByIDV1Item(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetRfProfileByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetRfProfilesV1Items(items *[]catalystcentersdkgo.ResponseWirelessGetRfProfilesV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["rf_profile_name"] = item.RfProfileName
		respItem["default_rf_profile"] = boolPtrToString(item.DefaultRfProfile)
		respItem["enable_radio_type_a"] = boolPtrToString(item.EnableRadioTypeA)
		respItem["enable_radio_type_b"] = boolPtrToString(item.EnableRadioTypeB)
		respItem["enable_radio_type6_g_hz"] = boolPtrToString(item.EnableRadioType6GHz)
		respItem["enable_custom"] = boolPtrToString(item.EnableCustom)
		respItem["radio_type_a_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioTypeAProperties(item.RadioTypeAProperties)
		respItem["radio_type_b_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioTypeBProperties(item.RadioTypeBProperties)
		respItem["radio_type6_g_hz_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioType6GHzProperties(item.RadioType6GHzProperties)
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetRfProfilesV1ItemsRadioTypeAProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioTypeAProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThresholdV1
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["channel_width"] = item.ChannelWidth
	respItem["preamble_puncture"] = boolPtrToString(item.PreamblePuncture)
	respItem["zero_wait_dfs_enable"] = boolPtrToString(item.ZeroWaitDfsEnable)
	respItem["custom_rx_sop_threshold"] = item.CustomRxSopThreshold
	respItem["max_radio_clients"] = item.MaxRadioClients
	respItem["fra_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioTypeAPropertiesFraProperties(item.FraProperties)
	respItem["coverage_hole_detection_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioTypeAPropertiesCoverageHoleDetectionProperties(item.CoverageHoleDetectionProperties)
	respItem["spatial_reuse_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioTypeAPropertiesSpatialReuseProperties(item.SpatialReuseProperties)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioTypeAPropertiesFraProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesFraProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["client_aware"] = boolPtrToString(item.ClientAware)
	respItem["client_select"] = item.ClientSelect
	respItem["client_reset"] = item.ClientReset

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioTypeAPropertiesCoverageHoleDetectionProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesCoverageHoleDetectionProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["chd_client_level"] = item.ChdClientLevel
	respItem["chd_data_rssi_threshold"] = item.ChdDataRssiThreshold
	respItem["chd_voice_rssi_threshold"] = item.ChdVoiceRssiThreshold
	respItem["chd_exception_level"] = item.ChdExceptionLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioTypeAPropertiesSpatialReuseProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioTypeAPropertiesSpatialReuseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_non_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxNonSrgObssPacketDetect)
	respItem["dot11ax_non_srg_obss_packet_detect_max_threshold"] = item.Dot11AxNonSrgObssPacketDetectMaxThreshold
	respItem["dot11ax_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxSrgObssPacketDetect)
	respItem["dot11ax_srg_obss_packet_detect_min_threshold"] = item.Dot11AxSrgObssPacketDetectMinThreshold
	respItem["dot11ax_srg_obss_packet_detect_max_threshold"] = item.Dot11AxSrgObssPacketDetectMaxThreshold

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioTypeBProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioTypeBProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThresholdV1
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["custom_rx_sop_threshold"] = item.CustomRxSopThreshold
	respItem["max_radio_clients"] = item.MaxRadioClients
	respItem["coverage_hole_detection_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioTypeBPropertiesCoverageHoleDetectionProperties(item.CoverageHoleDetectionProperties)
	respItem["spatial_reuse_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioTypeBPropertiesSpatialReuseProperties(item.SpatialReuseProperties)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioTypeBPropertiesCoverageHoleDetectionProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioTypeBPropertiesCoverageHoleDetectionProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["chd_client_level"] = item.ChdClientLevel
	respItem["chd_data_rssi_threshold"] = item.ChdDataRssiThreshold
	respItem["chd_voice_rssi_threshold"] = item.ChdVoiceRssiThreshold
	respItem["chd_exception_level"] = item.ChdExceptionLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioTypeBPropertiesSpatialReuseProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioTypeBPropertiesSpatialReuseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_non_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxNonSrgObssPacketDetect)
	respItem["dot11ax_non_srg_obss_packet_detect_max_threshold"] = item.Dot11AxNonSrgObssPacketDetectMaxThreshold
	respItem["dot11ax_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxSrgObssPacketDetect)
	respItem["dot11ax_srg_obss_packet_detect_min_threshold"] = item.Dot11AxSrgObssPacketDetectMinThreshold
	respItem["dot11ax_srg_obss_packet_detect_max_threshold"] = item.Dot11AxSrgObssPacketDetectMaxThreshold

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioType6GHzProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThresholdV1
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["enable_standard_power_service"] = boolPtrToString(item.EnableStandardPowerService)
	respItem["multi_bssid_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesMultiBssidProperties(item.MultiBssidProperties)
	respItem["preamble_puncture"] = boolPtrToString(item.PreamblePuncture)
	respItem["min_dbs_width"] = item.MinDbsWidth
	respItem["max_dbs_width"] = item.MaxDbsWidth
	respItem["custom_rx_sop_threshold"] = item.CustomRxSopThreshold
	respItem["max_radio_clients"] = item.MaxRadioClients
	respItem["psc_enforcing_enabled"] = boolPtrToString(item.PscEnforcingEnabled)
	respItem["discovery_frames6_g_hz"] = item.DiscoveryFrames6GHz
	respItem["broadcast_probe_response_interval"] = item.BroadcastProbeResponseInterval
	respItem["fra_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesFraProperties(item.FraProperties)
	respItem["coverage_hole_detection_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesCoverageHoleDetectionProperties(item.CoverageHoleDetectionProperties)
	respItem["spatial_reuse_properties"] = flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesSpatialReuseProperties(item.SpatialReuseProperties)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesMultiBssidProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_parameters"] = flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(item.Dot11AxParameters)
	respItem["dot11be_parameters"] = flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(item.Dot11BeParameters)
	respItem["target_wake_time"] = boolPtrToString(item.TargetWakeTime)
	respItem["twt_broadcast_support"] = boolPtrToString(item.TwtBroadcastSupport)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)
	respItem["ofdma_multi_ru"] = boolPtrToString(item.OfdmaMultiRu)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesFraProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesFraProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["client_reset_count"] = item.ClientResetCount
	respItem["client_utilization_threshold"] = item.ClientUtilizationThreshold

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesCoverageHoleDetectionProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["chd_client_level"] = item.ChdClientLevel
	respItem["chd_data_rssi_threshold"] = item.ChdDataRssiThreshold
	respItem["chd_voice_rssi_threshold"] = item.ChdVoiceRssiThreshold
	respItem["chd_exception_level"] = item.ChdExceptionLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfilesV1ItemsRadioType6GHzPropertiesSpatialReuseProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfilesV1ResponseRadioType6GHzPropertiesSpatialReuseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_non_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxNonSrgObssPacketDetect)
	respItem["dot11ax_non_srg_obss_packet_detect_max_threshold"] = item.Dot11AxNonSrgObssPacketDetectMaxThreshold
	respItem["dot11ax_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxSrgObssPacketDetect)
	respItem["dot11ax_srg_obss_packet_detect_min_threshold"] = item.Dot11AxSrgObssPacketDetectMinThreshold
	respItem["dot11ax_srg_obss_packet_detect_max_threshold"] = item.Dot11AxSrgObssPacketDetectMaxThreshold

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1Item(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rf_profile_name"] = item.RfProfileName
	respItem["default_rf_profile"] = boolPtrToString(item.DefaultRfProfile)
	respItem["enable_radio_type_a"] = boolPtrToString(item.EnableRadioTypeA)
	respItem["enable_radio_type_b"] = boolPtrToString(item.EnableRadioTypeB)
	respItem["enable_radio_type6_g_hz"] = boolPtrToString(item.EnableRadioType6GHz)
	respItem["enable_custom"] = boolPtrToString(item.EnableCustom)
	respItem["radio_type_a_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioTypeAProperties(item.RadioTypeAProperties)
	respItem["radio_type_b_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioTypeBProperties(item.RadioTypeBProperties)
	respItem["radio_type6_g_hz_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzProperties(item.RadioType6GHzProperties)
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetRfProfileByIDV1ItemRadioTypeAProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThresholdV1
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["channel_width"] = item.ChannelWidth
	respItem["preamble_puncture"] = boolPtrToString(item.PreamblePuncture)
	respItem["zero_wait_dfs_enable"] = boolPtrToString(item.ZeroWaitDfsEnable)
	respItem["custom_rx_sop_threshold"] = item.CustomRxSopThreshold
	respItem["max_radio_clients"] = item.MaxRadioClients
	respItem["fra_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioTypeAPropertiesFraProperties(item.FraProperties)
	respItem["coverage_hole_detection_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioTypeAPropertiesCoverageHoleDetectionProperties(item.CoverageHoleDetectionProperties)
	respItem["spatial_reuse_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioTypeAPropertiesSpatialReuseProperties(item.SpatialReuseProperties)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioTypeAPropertiesFraProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesFraProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["client_aware"] = boolPtrToString(item.ClientAware)
	respItem["client_select"] = item.ClientSelect
	respItem["client_reset"] = item.ClientReset

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioTypeAPropertiesCoverageHoleDetectionProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesCoverageHoleDetectionProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["chd_client_level"] = item.ChdClientLevel
	respItem["chd_data_rssi_threshold"] = item.ChdDataRssiThreshold
	respItem["chd_voice_rssi_threshold"] = item.ChdVoiceRssiThreshold
	respItem["chd_exception_level"] = item.ChdExceptionLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioTypeAPropertiesSpatialReuseProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeAPropertiesSpatialReuseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_non_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxNonSrgObssPacketDetect)
	respItem["dot11ax_non_srg_obss_packet_detect_max_threshold"] = item.Dot11AxNonSrgObssPacketDetectMaxThreshold
	respItem["dot11ax_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxSrgObssPacketDetect)
	respItem["dot11ax_srg_obss_packet_detect_min_threshold"] = item.Dot11AxSrgObssPacketDetectMinThreshold
	respItem["dot11ax_srg_obss_packet_detect_max_threshold"] = item.Dot11AxSrgObssPacketDetectMaxThreshold

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioTypeBProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThresholdV1
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["custom_rx_sop_threshold"] = item.CustomRxSopThreshold
	respItem["max_radio_clients"] = item.MaxRadioClients
	respItem["coverage_hole_detection_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioTypeBPropertiesCoverageHoleDetectionProperties(item.CoverageHoleDetectionProperties)
	respItem["spatial_reuse_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioTypeBPropertiesSpatialReuseProperties(item.SpatialReuseProperties)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioTypeBPropertiesCoverageHoleDetectionProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBPropertiesCoverageHoleDetectionProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["chd_client_level"] = item.ChdClientLevel
	respItem["chd_data_rssi_threshold"] = item.ChdDataRssiThreshold
	respItem["chd_voice_rssi_threshold"] = item.ChdVoiceRssiThreshold
	respItem["chd_exception_level"] = item.ChdExceptionLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioTypeBPropertiesSpatialReuseProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioTypeBPropertiesSpatialReuseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_non_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxNonSrgObssPacketDetect)
	respItem["dot11ax_non_srg_obss_packet_detect_max_threshold"] = item.Dot11AxNonSrgObssPacketDetectMaxThreshold
	respItem["dot11ax_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxSrgObssPacketDetect)
	respItem["dot11ax_srg_obss_packet_detect_min_threshold"] = item.Dot11AxSrgObssPacketDetectMinThreshold
	respItem["dot11ax_srg_obss_packet_detect_max_threshold"] = item.Dot11AxSrgObssPacketDetectMaxThreshold

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["parent_profile"] = item.ParentProfile
	respItem["radio_channels"] = item.RadioChannels
	respItem["data_rates"] = item.DataRates
	respItem["mandatory_data_rates"] = item.MandatoryDataRates
	respItem["power_threshold_v1"] = item.PowerThresholdV1
	respItem["rx_sop_threshold"] = item.RxSopThreshold
	respItem["min_power_level"] = item.MinPowerLevel
	respItem["max_power_level"] = item.MaxPowerLevel
	respItem["enable_standard_power_service"] = boolPtrToString(item.EnableStandardPowerService)
	respItem["multi_bssid_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesMultiBssidProperties(item.MultiBssidProperties)
	respItem["preamble_puncture"] = boolPtrToString(item.PreamblePuncture)
	respItem["min_dbs_width"] = item.MinDbsWidth
	respItem["max_dbs_width"] = item.MaxDbsWidth
	respItem["custom_rx_sop_threshold"] = item.CustomRxSopThreshold
	respItem["max_radio_clients"] = item.MaxRadioClients
	respItem["psc_enforcing_enabled"] = boolPtrToString(item.PscEnforcingEnabled)
	respItem["discovery_frames6_g_hz"] = item.DiscoveryFrames6GHz
	respItem["broadcast_probe_response_interval"] = item.BroadcastProbeResponseInterval
	respItem["fra_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesFraProperties(item.FraProperties)
	respItem["coverage_hole_detection_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesCoverageHoleDetectionProperties(item.CoverageHoleDetectionProperties)
	respItem["spatial_reuse_properties"] = flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesSpatialReuseProperties(item.SpatialReuseProperties)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesMultiBssidProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_parameters"] = flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(item.Dot11AxParameters)
	respItem["dot11be_parameters"] = flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(item.Dot11BeParameters)
	respItem["target_wake_time"] = boolPtrToString(item.TargetWakeTime)
	respItem["twt_broadcast_support"] = boolPtrToString(item.TwtBroadcastSupport)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11AxParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesMultiBssidPropertiesDot11BeParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ofdma_down_link"] = boolPtrToString(item.OfdmaDownLink)
	respItem["ofdma_up_link"] = boolPtrToString(item.OfdmaUpLink)
	respItem["mu_mimo_up_link"] = boolPtrToString(item.MuMimoUpLink)
	respItem["mu_mimo_down_link"] = boolPtrToString(item.MuMimoDownLink)
	respItem["ofdma_multi_ru"] = boolPtrToString(item.OfdmaMultiRu)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesFraProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesFraProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["client_reset_count"] = item.ClientResetCount
	respItem["client_utilization_threshold"] = item.ClientUtilizationThreshold

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesCoverageHoleDetectionProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesCoverageHoleDetectionProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["chd_client_level"] = item.ChdClientLevel
	respItem["chd_data_rssi_threshold"] = item.ChdDataRssiThreshold
	respItem["chd_voice_rssi_threshold"] = item.ChdVoiceRssiThreshold
	respItem["chd_exception_level"] = item.ChdExceptionLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetRfProfileByIDV1ItemRadioType6GHzPropertiesSpatialReuseProperties(item *catalystcentersdkgo.ResponseWirelessGetRfProfileByIDV1ResponseRadioType6GHzPropertiesSpatialReuseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dot11ax_non_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxNonSrgObssPacketDetect)
	respItem["dot11ax_non_srg_obss_packet_detect_max_threshold"] = item.Dot11AxNonSrgObssPacketDetectMaxThreshold
	respItem["dot11ax_srg_obss_packet_detect"] = boolPtrToString(item.Dot11AxSrgObssPacketDetect)
	respItem["dot11ax_srg_obss_packet_detect_min_threshold"] = item.Dot11AxSrgObssPacketDetectMinThreshold
	respItem["dot11ax_srg_obss_packet_detect_max_threshold"] = item.Dot11AxSrgObssPacketDetectMaxThreshold

	return []map[string]interface{}{
		respItem,
	}

}
