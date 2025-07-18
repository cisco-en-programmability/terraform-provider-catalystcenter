package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePathTrace() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Path Trace.

- Returns a summary of all flow analyses stored. Results can be filtered by specified parameters.

- Returns result of a previously requested flow analysis by its Flow Analysis id
`,

		ReadContext: dataSourcePathTraceRead,
		Schema: map[string]*schema.Schema{
			"dest_ip": &schema.Schema{
				Description: `destIP query parameter. Destination IP address
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"dest_port": &schema.Schema{
				Description: `destPort query parameter. Destination port
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"flow_analysis_id": &schema.Schema{
				Description: `flowAnalysisId path parameter. Flow analysis request id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"gt_create_time": &schema.Schema{
				Description: `gtCreateTime query parameter. Analyses requested after this time
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"last_update_time": &schema.Schema{
				Description: `lastUpdateTime query parameter. Last update time
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of resources returned
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"lt_create_time": &schema.Schema{
				Description: `ltCreateTime query parameter. Analyses requested before this time
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Start index of resources returned (1-based)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Order by this field
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"periodic_refresh": &schema.Schema{
				Description: `periodicRefresh query parameter. Is analysis periodically refreshed?
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"protocol": &schema.Schema{
				Description: `protocol query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Sort by this field
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_ip": &schema.Schema{
				Description: `sourceIP query parameter. Source IP address
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"source_port": &schema.Schema{
				Description: `sourcePort query parameter. Source port
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"task_id": &schema.Schema{
				Description: `taskId query parameter. Task ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"detailed_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"acl_trace_calculation": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"acl_trace_calculation_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"last_update": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_elements": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accuracy_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"percent": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"detailed_status": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_trace_calculation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"acl_trace_calculation_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"device_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cpu_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"five_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"five_secs_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"one_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"memory_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"memory_usage": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"total_memory": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"device_stats_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_stats_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"egress_physical_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},

																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},

												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"egress_virtual_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},

																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},

												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"authentication": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"data_switching": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"egress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},

																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"ingress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},

																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"wireless_lan_controller_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"wireless_lan_controller_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"ingress_physical_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},

																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},

												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"ingress_virtual_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},

																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"admin_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"input_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_flushes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_queue_max_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"input_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"operational_status": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"output_drop": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_queue_count": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"output_ratebps": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"interface_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"interface_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"path_overlay_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"control_plane": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"data_packet_encapsulation": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"dest_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"dest_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"protocol": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"source_ip": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"source_port": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"vxlan_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"dscp": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"vnid": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},

												"qos_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"class_map_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"drop_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"num_bytes": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"num_packets": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"offered_rate": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_bandwidthbps": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"queue_depth": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_no_buffer_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"queue_total_drops": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"qos_stats_collection": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"qos_stats_collection_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"used_vlan": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"vrf_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"link_information_source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"perf_mon_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"perf_mon_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"perf_mon_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byte_rate": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dest_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"dest_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"input_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipv4_dsc_p": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipv4_ttl": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"output_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"packet_bytes": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"packet_count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"packet_loss": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"packet_loss_percentage": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"protocol": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"refreshed_at": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"rtp_jitter_max": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"rtp_jitter_mean": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"rtp_jitter_min": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"source_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"source_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"role": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"tunnels": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"wlan_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"network_elements_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"accuracy_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"percent": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"detailed_status": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"acl_trace_calculation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"acl_trace_calculation_failure_reason": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"device_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"cpu_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"five_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"five_secs_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"one_min_usage_in_percentage": &schema.Schema{
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"memory_statistics": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"memory_usage": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"refreshed_at": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},

															"total_memory": &schema.Schema{
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"device_stats_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_stats_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"egress_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"physical_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},

																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},

																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},

																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},

																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},

															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},

																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},

															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"virtual_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},

																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},

																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},

																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},

																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},

															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},

																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},

															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"flex_connect": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"authentication": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"data_switching": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"egress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},

																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"ingress_acl_analysis": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"matching_aces": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"ace": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_ports": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"dest_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},

																								"source_ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Schema{
																										Type: schema.TypeString,
																									},
																								},
																							},
																						},
																					},

																					"protocol": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"result": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"wireless_lan_controller_id": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"wireless_lan_controller_name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"ingress_interface": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"physical_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},

																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},

																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},

																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},

																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},

															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},

																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},

															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"virtual_interface": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"acl_analysis": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"acl_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"matching_aces": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"ace": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},

																					"matching_ports": &schema.Schema{
																						Type:     schema.TypeList,
																						Computed: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{

																								"ports": &schema.Schema{
																									Type:     schema.TypeList,
																									Computed: true,
																									Elem: &schema.Resource{
																										Schema: map[string]*schema.Schema{

																											"dest_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},

																											"source_ports": &schema.Schema{
																												Type:     schema.TypeList,
																												Computed: true,
																												Elem: &schema.Schema{
																													Type: schema.TypeString,
																												},
																											},
																										},
																									},
																								},

																								"protocol": &schema.Schema{
																									Type:     schema.TypeString,
																									Computed: true,
																								},
																							},
																						},
																					},

																					"result": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},

																		"result": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"id": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"interface_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"admin_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"input_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_flushes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_queue_max_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"input_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"operational_status": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"output_drop": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_queue_count": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"output_ratebps": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},

															"interface_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"interface_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"path_overlay_info": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"control_plane": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"data_packet_encapsulation": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"dest_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"dest_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"protocol": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source_ip": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source_port": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"vxlan_info": &schema.Schema{
																			Type:     schema.TypeList,
																			Computed: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{

																					"dscp": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},

																					"vnid": &schema.Schema{
																						Type:     schema.TypeString,
																						Computed: true,
																					},
																				},
																			},
																		},
																	},
																},
															},

															"qos_statistics": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"class_map_name": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"drop_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"num_bytes": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"num_packets": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"offered_rate": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_bandwidthbps": &schema.Schema{
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"queue_depth": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_no_buffer_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"queue_total_drops": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},

																		"refreshed_at": &schema.Schema{
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																	},
																},
															},

															"qos_stats_collection": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"qos_stats_collection_failure_reason": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"used_vlan": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},

															"vrf_name": &schema.Schema{
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"link_information_source": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"perf_mon_collection": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"perf_mon_collection_failure_reason": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"perf_monitor_statistics": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"byte_rate": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"dest_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"dest_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"input_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipv4_dsc_p": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"ipv4_ttl": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"output_interface": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"packet_bytes": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"packet_count": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"packet_loss": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"packet_loss_percentage": &schema.Schema{
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"protocol": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"refreshed_at": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"rtp_jitter_max": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"rtp_jitter_mean": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"rtp_jitter_min": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},

												"source_ip_address": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"source_port": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"role": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssid": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"tunnels": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"wlan_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"request": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"control_path": &schema.Schema{
										Description: `Control path tracing
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"create_time": &schema.Schema{
										Description: `Timestamp when the Path Trace request was first received
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"dest_ip": &schema.Schema{
										Description: `IP Address of the destination device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dest_port": &schema.Schema{
										Description: `Port on the destination device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_reason": &schema.Schema{
										Description: `Reason for failure
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Unique ID for the Path Trace request
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inclusions": &schema.Schema{
										Description: `Subset of {INTERFACE-STATS, QOS-STATS, DEVICE-STATS, PERFORMANCE-STATS, ACL-TRACE}
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"last_update_time": &schema.Schema{
										Description: `Last timestamp when the path trace response was updated
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"periodic_refresh": &schema.Schema{
										Description: `Re-run the Path Trace every 30 seconds
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"previous_flow_analysis_id": &schema.Schema{
										Description: `When periodicRefresh is true, this field holds the original Path Trace request ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"protocol": &schema.Schema{
										Description: `One of TCP/UDP or either (null)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"source_ip": &schema.Schema{
										Description: `IP Address of the source device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"source_port": &schema.Schema{
										Description: `Port on the source device
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `One of {SUCCESS, INPROGRESS, FAILED, SCHEDULED, PENDING, COMPLETED}
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"control_path": &schema.Schema{
							Description: `Control path tracing
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"create_time": &schema.Schema{
							Description: `Timestamp when the Path Trace request was first received
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"dest_ip": &schema.Schema{
							Description: `IP Address of the destination device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"dest_port": &schema.Schema{
							Description: `Port on the destination device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"failure_reason": &schema.Schema{
							Description: `Reason for failure
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Unique ID for the Path Trace request
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"inclusions": &schema.Schema{
							Description: `Subset of {INTERFACE-STATS, QOS-STATS, DEVICE-STATS, PERFORMANCE-STATS, ACL-TRACE}
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"last_update_time": &schema.Schema{
							Description: `Last timestamp when the path trace response was updated
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"periodic_refresh": &schema.Schema{
							Description: `Re-run the Path Trace every 30 seconds
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"previous_flow_analysis_id": &schema.Schema{
							Description: `When periodicRefresh is true, this field holds the original Path Trace request ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"protocol": &schema.Schema{
							Description: `One of TCP/UDP or either (null)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"source_ip": &schema.Schema{
							Description: `IP Address of the source device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"source_port": &schema.Schema{
							Description: `Port on the source device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `One of {SUCCESS, INPROGRESS, FAILED, SCHEDULED, PENDING, COMPLETED}
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

func dataSourcePathTraceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vPeriodicRefresh, okPeriodicRefresh := d.GetOk("periodic_refresh")
	vSourceIP, okSourceIP := d.GetOk("source_ip")
	vDestIP, okDestIP := d.GetOk("dest_ip")
	vSourcePort, okSourcePort := d.GetOk("source_port")
	vDestPort, okDestPort := d.GetOk("dest_port")
	vGtCreateTime, okGtCreateTime := d.GetOk("gt_create_time")
	vLtCreateTime, okLtCreateTime := d.GetOk("lt_create_time")
	vProtocol, okProtocol := d.GetOk("protocol")
	vStatus, okStatus := d.GetOk("status")
	vTaskID, okTaskID := d.GetOk("task_id")
	vLastUpdateTime, okLastUpdateTime := d.GetOk("last_update_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vOrder, okOrder := d.GetOk("order")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vFlowAnalysisID, okFlowAnalysisID := d.GetOk("flow_analysis_id")

	method1 := []bool{okPeriodicRefresh, okSourceIP, okDestIP, okSourcePort, okDestPort, okGtCreateTime, okLtCreateTime, okProtocol, okStatus, okTaskID, okLastUpdateTime, okLimit, okOffset, okOrder, okSortBy}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okFlowAnalysisID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesAllPreviousPathtracesSummary")
		queryParams1 := catalystcentersdkgo.RetrievesAllPreviousPathtracesSummaryQueryParams{}

		if okPeriodicRefresh {
			queryParams1.PeriodicRefresh = vPeriodicRefresh.(bool)
		}
		if okSourceIP {
			queryParams1.SourceIP = vSourceIP.(string)
		}
		if okDestIP {
			queryParams1.DestIP = vDestIP.(string)
		}
		if okSourcePort {
			queryParams1.SourcePort = vSourcePort.(float64)
		}
		if okDestPort {
			queryParams1.DestPort = vDestPort.(float64)
		}
		if okGtCreateTime {
			queryParams1.GtCreateTime = vGtCreateTime.(float64)
		}
		if okLtCreateTime {
			queryParams1.LtCreateTime = vLtCreateTime.(float64)
		}
		if okProtocol {
			queryParams1.Protocol = vProtocol.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}
		if okLastUpdateTime {
			queryParams1.LastUpdateTime = vLastUpdateTime.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.PathTrace.RetrievesAllPreviousPathtracesSummary(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesAllPreviousPathtracesSummary", err,
				"Failure at RetrievesAllPreviousPathtracesSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesAllPreviousPathtracesSummary", err,
				"Failure at RetrievesAllPreviousPathtracesSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenPathTraceRetrievesAllPreviousPathtracesSummaryItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesAllPreviousPathtracesSummary response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: RetrievesPreviousPathtrace")
		vvFlowAnalysisID := vFlowAnalysisID.(string)

		// has_unknown_response: None

		response2, restyResp2, err := client.PathTrace.RetrievesPreviousPathtrace(vvFlowAnalysisID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesPreviousPathtrace", err,
				"Failure at RetrievesPreviousPathtrace, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesPreviousPathtrace", err,
				"Failure at RetrievesPreviousPathtrace, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenPathTraceRetrievesPreviousPathtraceItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesPreviousPathtrace response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPathTraceRetrievesAllPreviousPathtracesSummaryItems(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesAllPreviousPathtracesSummaryResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_path"] = boolPtrToString(item.ControlPath)
		respItem["create_time"] = item.CreateTime
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["failure_reason"] = item.FailureReason
		respItem["id"] = item.ID
		respItem["inclusions"] = item.Inclusions
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["periodic_refresh"] = boolPtrToString(item.PeriodicRefresh)
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["status"] = item.Status
		respItem["previous_flow_analysis_id"] = item.PreviousFlowAnalysisID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItem(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["detailed_status"] = flattenPathTraceRetrievesPreviousPathtraceItemDetailedStatus(item.DetailedStatus)
	respItem["last_update"] = item.LastUpdate
	respItem["network_elements"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElements(item.NetworkElements)
	respItem["network_elements_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfo(item.NetworkElementsInfo)
	respItem["properties"] = item.Properties
	respItem["request"] = flattenPathTraceRetrievesPreviousPathtraceItemRequest(item.Request)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPathTraceRetrievesPreviousPathtraceItemDetailedStatus(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseDetailedStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_trace_calculation"] = item.ACLTraceCalculation
	respItem["acl_trace_calculation_failure_reason"] = item.ACLTraceCalculationFailureReason

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElements(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElements) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["accuracy_list"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsAccuracyList(item.AccuracyList)
		respItem["detailed_status"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsDetailedStatus(item.DetailedStatus)
		respItem["device_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsDeviceStatistics(item.DeviceStatistics)
		respItem["device_stats_collection"] = item.DeviceStatsCollection
		respItem["device_stats_collection_failure_reason"] = item.DeviceStatsCollectionFailureReason
		respItem["egress_physical_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterface(item.EgressPhysicalInterface)
		respItem["egress_virtual_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterface(item.EgressVirtualInterface)
		respItem["flex_connect"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnect(item.FlexConnect)
		respItem["id"] = item.ID
		respItem["ingress_physical_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterface(item.IngressPhysicalInterface)
		respItem["ingress_virtual_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterface(item.IngressVirtualInterface)
		respItem["ip"] = item.IP
		respItem["link_information_source"] = item.LinkInformationSource
		respItem["name"] = item.Name
		respItem["perf_mon_collection"] = item.PerfMonCollection
		respItem["perf_mon_collection_failure_reason"] = item.PerfMonCollectionFailureReason
		respItem["perf_mon_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsPerfMonStatistics(item.PerfMonStatistics)
		respItem["role"] = item.Role
		respItem["ssid"] = item.SSID
		respItem["tunnels"] = item.Tunnels
		respItem["type"] = item.Type
		respItem["wlan_id"] = item.WLANID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsAccuracyList(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsAccuracyList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["percent"] = item.Percent
		respItem["reason"] = item.Reason
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsDetailedStatus(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDetailedStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_trace_calculation"] = item.ACLTraceCalculation
	respItem["acl_trace_calculation_failure_reason"] = item.ACLTraceCalculationFailureReason

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsDeviceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cpu_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsDeviceStatisticsCPUStatistics(item.CPUStatistics)
	respItem["memory_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsDeviceStatisticsMemoryStatistics(item.MemoryStatistics)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsDeviceStatisticsCPUStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsCPUStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["five_min_usage_in_percentage"] = item.FiveMinUsageInPercentage
	respItem["five_secs_usage_in_percentage"] = item.FiveSecsUsageInPercentage
	respItem["one_min_usage_in_percentage"] = item.OneMinUsageInPercentage
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsDeviceStatisticsMemoryStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsDeviceStatisticsMemoryStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["memory_usage"] = item.MemoryUsage
	respItem["refreshed_at"] = item.RefreshedAt
	respItem["total_memory"] = item.TotalMemory

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["input_packets"] = item.InputPackets
	respItem["input_queue_count"] = item.InputQueueCount
	respItem["input_queue_drops"] = item.InputQueueDrops
	respItem["input_queue_flushes"] = item.InputQueueFlushes
	respItem["input_queue_max_depth"] = item.InputQueueMaxDepth
	respItem["input_ratebps"] = item.InputRatebps
	respItem["operational_status"] = item.OperationalStatus
	respItem["output_drop"] = item.OutputDrop
	respItem["output_packets"] = item.OutputPackets
	respItem["output_queue_count"] = item.OutputQueueCount
	respItem["output_queue_depth"] = item.OutputQueueDepth
	respItem["output_ratebps"] = item.OutputRatebps
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_plane"] = item.ControlPlane
		respItem["data_packet_encapsulation"] = item.DataPacketEncapsulation
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dscp"] = item.Dscp
	respItem["vnid"] = item.Vnid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressPhysicalInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressPhysicalInterfaceQosStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["class_map_name"] = item.ClassMapName
		respItem["drop_rate"] = item.DropRate
		respItem["num_bytes"] = item.NumBytes
		respItem["num_packets"] = item.NumPackets
		respItem["offered_rate"] = item.OfferedRate
		respItem["queue_bandwidthbps"] = item.QueueBandwidthbps
		respItem["queue_depth"] = item.QueueDepth
		respItem["queue_no_buffer_drops"] = item.QueueNoBufferDrops
		respItem["queue_total_drops"] = item.QueueTotalDrops
		respItem["refreshed_at"] = item.RefreshedAt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["input_packets"] = item.InputPackets
	respItem["input_queue_count"] = item.InputQueueCount
	respItem["input_queue_drops"] = item.InputQueueDrops
	respItem["input_queue_flushes"] = item.InputQueueFlushes
	respItem["input_queue_max_depth"] = item.InputQueueMaxDepth
	respItem["input_ratebps"] = item.InputRatebps
	respItem["operational_status"] = item.OperationalStatus
	respItem["output_drop"] = item.OutputDrop
	respItem["output_packets"] = item.OutputPackets
	respItem["output_queue_count"] = item.OutputQueueCount
	respItem["output_queue_depth"] = item.OutputQueueDepth
	respItem["output_ratebps"] = item.OutputRatebps
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_plane"] = item.ControlPlane
		respItem["data_packet_encapsulation"] = item.DataPacketEncapsulation
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dscp"] = item.Dscp
	respItem["vnid"] = item.Vnid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsEgressVirtualInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsEgressVirtualInterfaceQosStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["class_map_name"] = item.ClassMapName
		respItem["drop_rate"] = item.DropRate
		respItem["num_bytes"] = item.NumBytes
		respItem["num_packets"] = item.NumPackets
		respItem["offered_rate"] = item.OfferedRate
		respItem["queue_bandwidthbps"] = item.QueueBandwidthbps
		respItem["queue_depth"] = item.QueueDepth
		respItem["queue_no_buffer_drops"] = item.QueueNoBufferDrops
		respItem["queue_total_drops"] = item.QueueTotalDrops
		respItem["refreshed_at"] = item.RefreshedAt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnect(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnect) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["authentication"] = item.Authentication
	respItem["data_switching"] = item.DataSwitching
	respItem["egress_acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectEgressACLAnalysis(item.EgressACLAnalysis)
	respItem["ingress_acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectIngressACLAnalysis(item.IngressACLAnalysis)
	respItem["wireless_lan_controller_id"] = item.WirelessLanControllerID
	respItem["wireless_lan_controller_name"] = item.WirelessLanControllerName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectEgressACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectIngressACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["input_packets"] = item.InputPackets
	respItem["input_queue_count"] = item.InputQueueCount
	respItem["input_queue_drops"] = item.InputQueueDrops
	respItem["input_queue_flushes"] = item.InputQueueFlushes
	respItem["input_queue_max_depth"] = item.InputQueueMaxDepth
	respItem["input_ratebps"] = item.InputRatebps
	respItem["operational_status"] = item.OperationalStatus
	respItem["output_drop"] = item.OutputDrop
	respItem["output_packets"] = item.OutputPackets
	respItem["output_queue_count"] = item.OutputQueueCount
	respItem["output_queue_depth"] = item.OutputQueueDepth
	respItem["output_ratebps"] = item.OutputRatebps
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_plane"] = item.ControlPlane
		respItem["data_packet_encapsulation"] = item.DataPacketEncapsulation
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dscp"] = item.Dscp
	respItem["vnid"] = item.Vnid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressPhysicalInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressPhysicalInterfaceQosStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["class_map_name"] = item.ClassMapName
		respItem["drop_rate"] = item.DropRate
		respItem["num_bytes"] = item.NumBytes
		respItem["num_packets"] = item.NumPackets
		respItem["offered_rate"] = item.OfferedRate
		respItem["queue_bandwidthbps"] = item.QueueBandwidthbps
		respItem["queue_depth"] = item.QueueDepth
		respItem["queue_no_buffer_drops"] = item.QueueNoBufferDrops
		respItem["queue_total_drops"] = item.QueueTotalDrops
		respItem["refreshed_at"] = item.RefreshedAt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["input_packets"] = item.InputPackets
	respItem["input_queue_count"] = item.InputQueueCount
	respItem["input_queue_drops"] = item.InputQueueDrops
	respItem["input_queue_flushes"] = item.InputQueueFlushes
	respItem["input_queue_max_depth"] = item.InputQueueMaxDepth
	respItem["input_ratebps"] = item.InputRatebps
	respItem["operational_status"] = item.OperationalStatus
	respItem["output_drop"] = item.OutputDrop
	respItem["output_packets"] = item.OutputPackets
	respItem["output_queue_count"] = item.OutputQueueCount
	respItem["output_queue_depth"] = item.OutputQueueDepth
	respItem["output_ratebps"] = item.OutputRatebps
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_plane"] = item.ControlPlane
		respItem["data_packet_encapsulation"] = item.DataPacketEncapsulation
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dscp"] = item.Dscp
	respItem["vnid"] = item.Vnid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsIngressVirtualInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsIngressVirtualInterfaceQosStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["class_map_name"] = item.ClassMapName
		respItem["drop_rate"] = item.DropRate
		respItem["num_bytes"] = item.NumBytes
		respItem["num_packets"] = item.NumPackets
		respItem["offered_rate"] = item.OfferedRate
		respItem["queue_bandwidthbps"] = item.QueueBandwidthbps
		respItem["queue_depth"] = item.QueueDepth
		respItem["queue_no_buffer_drops"] = item.QueueNoBufferDrops
		respItem["queue_total_drops"] = item.QueueTotalDrops
		respItem["refreshed_at"] = item.RefreshedAt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsPerfMonStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsPerfMonStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["byte_rate"] = item.ByteRate
		respItem["dest_ip_address"] = item.DestIPAddress
		respItem["dest_port"] = item.DestPort
		respItem["input_interface"] = item.InputInterface
		respItem["ipv4_dsc_p"] = item.IPv4DSCP
		respItem["ipv4_ttl"] = item.IPv4TTL
		respItem["output_interface"] = item.OutputInterface
		respItem["packet_bytes"] = item.PacketBytes
		respItem["packet_count"] = item.PacketCount
		respItem["packet_loss"] = item.PacketLoss
		respItem["packet_loss_percentage"] = item.PacketLossPercentage
		respItem["protocol"] = item.Protocol
		respItem["refreshed_at"] = item.RefreshedAt
		respItem["rtp_jitter_max"] = item.RtpJitterMax
		respItem["rtp_jitter_mean"] = item.RtpJitterMean
		respItem["rtp_jitter_min"] = item.RtpJitterMin
		respItem["source_ip_address"] = item.SourceIPAddress
		respItem["source_port"] = item.SourcePort
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["accuracy_list"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoAccuracyList(item.AccuracyList)
		respItem["detailed_status"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoDetailedStatus(item.DetailedStatus)
		respItem["device_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoDeviceStatistics(item.DeviceStatistics)
		respItem["device_stats_collection"] = item.DeviceStatsCollection
		respItem["device_stats_collection_failure_reason"] = item.DeviceStatsCollectionFailureReason
		respItem["egress_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterface(item.EgressInterface)
		respItem["flex_connect"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnect(item.FlexConnect)
		respItem["id"] = item.ID
		respItem["ingress_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterface(item.IngressInterface)
		respItem["ip"] = item.IP
		respItem["link_information_source"] = item.LinkInformationSource
		respItem["name"] = item.Name
		respItem["perf_mon_collection"] = item.PerfMonCollection
		respItem["perf_mon_collection_failure_reason"] = item.PerfMonCollectionFailureReason
		respItem["perf_monitor_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoPerfMonitorStatistics(item.PerfMonitorStatistics)
		respItem["role"] = item.Role
		respItem["ssid"] = item.SSID
		respItem["tunnels"] = item.Tunnels
		respItem["type"] = item.Type
		respItem["wlan_id"] = item.WLANID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoAccuracyList(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoAccuracyList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["percent"] = item.Percent
		respItem["reason"] = item.Reason
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoDetailedStatus(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDetailedStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_trace_calculation"] = item.ACLTraceCalculation
	respItem["acl_trace_calculation_failure_reason"] = item.ACLTraceCalculationFailureReason

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoDeviceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cpu_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoDeviceStatisticsCPUStatistics(item.CPUStatistics)
	respItem["memory_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoDeviceStatisticsMemoryStatistics(item.MemoryStatistics)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoDeviceStatisticsCPUStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsCPUStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["five_min_usage_in_percentage"] = item.FiveMinUsageInPercentage
	respItem["five_secs_usage_in_percentage"] = item.FiveSecsUsageInPercentage
	respItem["one_min_usage_in_percentage"] = item.OneMinUsageInPercentage
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoDeviceStatisticsMemoryStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["memory_usage"] = item.MemoryUsage
	respItem["refreshed_at"] = item.RefreshedAt
	respItem["total_memory"] = item.TotalMemory

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["physical_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterface(item.PhysicalInterface)
	respItem["virtual_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterface(item.VirtualInterface)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["input_packets"] = item.InputPackets
	respItem["input_queue_count"] = item.InputQueueCount
	respItem["input_queue_drops"] = item.InputQueueDrops
	respItem["input_queue_flushes"] = item.InputQueueFlushes
	respItem["input_queue_max_depth"] = item.InputQueueMaxDepth
	respItem["input_ratebps"] = item.InputRatebps
	respItem["operational_status"] = item.OperationalStatus
	respItem["output_drop"] = item.OutputDrop
	respItem["output_packets"] = item.OutputPackets
	respItem["output_queue_count"] = item.OutputQueueCount
	respItem["output_queue_depth"] = item.OutputQueueDepth
	respItem["output_ratebps"] = item.OutputRatebps
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_plane"] = item.ControlPlane
		respItem["data_packet_encapsulation"] = item.DataPacketEncapsulation
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dscp"] = item.Dscp
	respItem["vnid"] = item.Vnid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["class_map_name"] = item.ClassMapName
		respItem["drop_rate"] = item.DropRate
		respItem["num_bytes"] = item.NumBytes
		respItem["num_packets"] = item.NumPackets
		respItem["offered_rate"] = item.OfferedRate
		respItem["queue_bandwidthbps"] = item.QueueBandwidthbps
		respItem["queue_depth"] = item.QueueDepth
		respItem["queue_no_buffer_drops"] = item.QueueNoBufferDrops
		respItem["queue_total_drops"] = item.QueueTotalDrops
		respItem["refreshed_at"] = item.RefreshedAt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterface(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterface) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis(item.ACLAnalysis)
		respItem["id"] = item.ID
		respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics(item.InterfaceStatistics)
		respItem["interface_stats_collection"] = item.InterfaceStatsCollection
		respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
		respItem["name"] = item.Name
		respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo(item.PathOverlayInfo)
		respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics(item.QosStatistics)
		respItem["qos_stats_collection"] = item.QosStatsCollection
		respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
		respItem["used_vlan"] = item.UsedVLAN
		respItem["vrf_name"] = item.VrfName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["input_packets"] = item.InputPackets
	respItem["input_queue_count"] = item.InputQueueCount
	respItem["input_queue_drops"] = item.InputQueueDrops
	respItem["input_queue_flushes"] = item.InputQueueFlushes
	respItem["input_queue_max_depth"] = item.InputQueueMaxDepth
	respItem["input_ratebps"] = item.InputRatebps
	respItem["operational_status"] = item.OperationalStatus
	respItem["output_drop"] = item.OutputDrop
	respItem["output_packets"] = item.OutputPackets
	respItem["output_queue_count"] = item.OutputQueueCount
	respItem["output_queue_depth"] = item.OutputQueueDepth
	respItem["output_ratebps"] = item.OutputRatebps
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_plane"] = item.ControlPlane
		respItem["data_packet_encapsulation"] = item.DataPacketEncapsulation
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dscp"] = item.Dscp
	respItem["vnid"] = item.Vnid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["class_map_name"] = item.ClassMapName
		respItem["drop_rate"] = item.DropRate
		respItem["num_bytes"] = item.NumBytes
		respItem["num_packets"] = item.NumPackets
		respItem["offered_rate"] = item.OfferedRate
		respItem["queue_bandwidthbps"] = item.QueueBandwidthbps
		respItem["queue_depth"] = item.QueueDepth
		respItem["queue_no_buffer_drops"] = item.QueueNoBufferDrops
		respItem["queue_total_drops"] = item.QueueTotalDrops
		respItem["refreshed_at"] = item.RefreshedAt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnect(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnect) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["authentication"] = item.Authentication
	respItem["data_switching"] = item.DataSwitching
	respItem["egress_acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectEgressACLAnalysis(item.EgressACLAnalysis)
	respItem["ingress_acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectIngressACLAnalysis(item.IngressACLAnalysis)
	respItem["wireless_lan_controller_id"] = item.WirelessLanControllerID
	respItem["wireless_lan_controller_name"] = item.WirelessLanControllerName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectEgressACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectIngressACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["physical_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterface(item.PhysicalInterface)
	respItem["virtual_interface"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterface(item.VirtualInterface)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["input_packets"] = item.InputPackets
	respItem["input_queue_count"] = item.InputQueueCount
	respItem["input_queue_drops"] = item.InputQueueDrops
	respItem["input_queue_flushes"] = item.InputQueueFlushes
	respItem["input_queue_max_depth"] = item.InputQueueMaxDepth
	respItem["input_ratebps"] = item.InputRatebps
	respItem["operational_status"] = item.OperationalStatus
	respItem["output_drop"] = item.OutputDrop
	respItem["output_packets"] = item.OutputPackets
	respItem["output_queue_count"] = item.OutputQueueCount
	respItem["output_queue_depth"] = item.OutputQueueDepth
	respItem["output_ratebps"] = item.OutputRatebps
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_plane"] = item.ControlPlane
		respItem["data_packet_encapsulation"] = item.DataPacketEncapsulation
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dscp"] = item.Dscp
	respItem["vnid"] = item.Vnid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["class_map_name"] = item.ClassMapName
		respItem["drop_rate"] = item.DropRate
		respItem["num_bytes"] = item.NumBytes
		respItem["num_packets"] = item.NumPackets
		respItem["offered_rate"] = item.OfferedRate
		respItem["queue_bandwidthbps"] = item.QueueBandwidthbps
		respItem["queue_depth"] = item.QueueDepth
		respItem["queue_no_buffer_drops"] = item.QueueNoBufferDrops
		respItem["queue_total_drops"] = item.QueueTotalDrops
		respItem["refreshed_at"] = item.RefreshedAt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterface(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterface) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis(item.ACLAnalysis)
		respItem["id"] = item.ID
		respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics(item.InterfaceStatistics)
		respItem["interface_stats_collection"] = item.InterfaceStatsCollection
		respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
		respItem["name"] = item.Name
		respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo(item.PathOverlayInfo)
		respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics(item.QosStatistics)
		respItem["qos_stats_collection"] = item.QosStatsCollection
		respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
		respItem["used_vlan"] = item.UsedVLAN
		respItem["vrf_name"] = item.VrfName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dest_ports"] = item.DestPorts
		respItem["source_ports"] = item.SourcePorts
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["admin_status"] = item.AdminStatus
	respItem["input_packets"] = item.InputPackets
	respItem["input_queue_count"] = item.InputQueueCount
	respItem["input_queue_drops"] = item.InputQueueDrops
	respItem["input_queue_flushes"] = item.InputQueueFlushes
	respItem["input_queue_max_depth"] = item.InputQueueMaxDepth
	respItem["input_ratebps"] = item.InputRatebps
	respItem["operational_status"] = item.OperationalStatus
	respItem["output_drop"] = item.OutputDrop
	respItem["output_packets"] = item.OutputPackets
	respItem["output_queue_count"] = item.OutputQueueCount
	respItem["output_queue_depth"] = item.OutputQueueDepth
	respItem["output_ratebps"] = item.OutputRatebps
	respItem["refreshed_at"] = item.RefreshedAt

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["control_plane"] = item.ControlPlane
		respItem["data_packet_encapsulation"] = item.DataPacketEncapsulation
		respItem["dest_ip"] = item.DestIP
		respItem["dest_port"] = item.DestPort
		respItem["protocol"] = item.Protocol
		respItem["source_ip"] = item.SourceIP
		respItem["source_port"] = item.SourcePort
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dscp"] = item.Dscp
	respItem["vnid"] = item.Vnid

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["class_map_name"] = item.ClassMapName
		respItem["drop_rate"] = item.DropRate
		respItem["num_bytes"] = item.NumBytes
		respItem["num_packets"] = item.NumPackets
		respItem["offered_rate"] = item.OfferedRate
		respItem["queue_bandwidthbps"] = item.QueueBandwidthbps
		respItem["queue_depth"] = item.QueueDepth
		respItem["queue_no_buffer_drops"] = item.QueueNoBufferDrops
		respItem["queue_total_drops"] = item.QueueTotalDrops
		respItem["refreshed_at"] = item.RefreshedAt
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemNetworkElementsInfoPerfMonitorStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseNetworkElementsInfoPerfMonitorStatistics) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["byte_rate"] = item.ByteRate
		respItem["dest_ip_address"] = item.DestIPAddress
		respItem["dest_port"] = item.DestPort
		respItem["input_interface"] = item.InputInterface
		respItem["ipv4_dsc_p"] = item.IPv4DSCP
		respItem["ipv4_ttl"] = item.IPv4TTL
		respItem["output_interface"] = item.OutputInterface
		respItem["packet_bytes"] = item.PacketBytes
		respItem["packet_count"] = item.PacketCount
		respItem["packet_loss"] = item.PacketLoss
		respItem["packet_loss_percentage"] = item.PacketLossPercentage
		respItem["protocol"] = item.Protocol
		respItem["refreshed_at"] = item.RefreshedAt
		respItem["rtp_jitter_max"] = item.RtpJitterMax
		respItem["rtp_jitter_mean"] = item.RtpJitterMean
		respItem["rtp_jitter_min"] = item.RtpJitterMin
		respItem["source_ip_address"] = item.SourceIPAddress
		respItem["source_port"] = item.SourcePort
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceItemRequest(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceResponseRequest) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["control_path"] = boolPtrToString(item.ControlPath)
	respItem["create_time"] = item.CreateTime
	respItem["dest_ip"] = item.DestIP
	respItem["dest_port"] = item.DestPort
	respItem["failure_reason"] = item.FailureReason
	respItem["id"] = item.ID
	respItem["inclusions"] = item.Inclusions
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["periodic_refresh"] = boolPtrToString(item.PeriodicRefresh)
	respItem["protocol"] = item.Protocol
	respItem["source_ip"] = item.SourceIP
	respItem["source_port"] = item.SourcePort
	respItem["status"] = item.Status
	respItem["previous_flow_analysis_id"] = item.PreviousFlowAnalysisID

	return []map[string]interface{}{
		respItem,
	}

}
