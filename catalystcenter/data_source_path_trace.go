package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

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
		log.Printf("[DEBUG] Selected method: RetrievesAllPreviousPathtracesSummaryV1")
		queryParams1 := catalystcentersdkgo.RetrievesAllPreviousPathtracesSummaryV1QueryParams{}

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

		response1, restyResp1, err := client.PathTrace.RetrievesAllPreviousPathtracesSummaryV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesAllPreviousPathtracesSummaryV1", err,
				"Failure at RetrievesAllPreviousPathtracesSummaryV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenPathTraceRetrievesAllPreviousPathtracesSummaryV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesAllPreviousPathtracesSummaryV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: RetrievesPreviousPathtraceV1")
		vvFlowAnalysisID := vFlowAnalysisID.(string)

		response2, restyResp2, err := client.PathTrace.RetrievesPreviousPathtraceV1(vvFlowAnalysisID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesPreviousPathtraceV1", err,
				"Failure at RetrievesPreviousPathtraceV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenPathTraceRetrievesPreviousPathtraceV1Item(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesPreviousPathtraceV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPathTraceRetrievesAllPreviousPathtracesSummaryV1Items(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesAllPreviousPathtracesSummaryV1Response) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1Item(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["detailed_status"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemDetailedStatus(item.DetailedStatus)
	respItem["last_update"] = item.LastUpdate
	respItem["network_elements"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElements(item.NetworkElements)
	respItem["network_elements_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfo(item.NetworkElementsInfo)
	respItem["properties"] = item.Properties
	respItem["request"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemRequest(item.Request)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemDetailedStatus(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseDetailedStatus) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElements(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElements) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["accuracy_list"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsAccuracyList(item.AccuracyList)
		respItem["detailed_status"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsDetailedStatus(item.DetailedStatus)
		respItem["device_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsDeviceStatistics(item.DeviceStatistics)
		respItem["device_stats_collection"] = item.DeviceStatsCollection
		respItem["device_stats_collection_failure_reason"] = item.DeviceStatsCollectionFailureReason
		respItem["egress_physical_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterface(item.EgressPhysicalInterface)
		respItem["egress_virtual_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterface(item.EgressVirtualInterface)
		respItem["flex_connect"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnect(item.FlexConnect)
		respItem["id"] = item.ID
		respItem["ingress_physical_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterface(item.IngressPhysicalInterface)
		respItem["ingress_virtual_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterface(item.IngressVirtualInterface)
		respItem["ip"] = item.IP
		respItem["link_information_source"] = item.LinkInformationSource
		respItem["name"] = item.Name
		respItem["perf_mon_collection"] = item.PerfMonCollection
		respItem["perf_mon_collection_failure_reason"] = item.PerfMonCollectionFailureReason
		respItem["perf_mon_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsPerfMonStatistics(item.PerfMonStatistics)
		respItem["role"] = item.Role
		respItem["ssid"] = item.SSID
		respItem["tunnels"] = item.Tunnels
		respItem["type"] = item.Type
		respItem["wlan_id"] = item.WLANID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsAccuracyList(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsAccuracyList) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsDetailedStatus(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDetailedStatus) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsDeviceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cpu_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsDeviceStatisticsCPUStatistics(item.CPUStatistics)
	respItem["memory_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsDeviceStatisticsMemoryStatistics(item.MemoryStatistics)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsDeviceStatisticsCPUStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatisticsCPUStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsDeviceStatisticsMemoryStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsDeviceStatisticsMemoryStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceInterfaceStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfo) []map[string]interface{} {
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
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressPhysicalInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressPhysicalInterfaceQosStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceInterfaceStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfacePathOverlayInfo) []map[string]interface{} {
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
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsEgressVirtualInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsEgressVirtualInterfaceQosStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnect(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnect) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["authentication"] = item.Authentication
	respItem["data_switching"] = item.DataSwitching
	respItem["egress_acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectEgressACLAnalysis(item.EgressACLAnalysis)
	respItem["ingress_acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectIngressACLAnalysis(item.IngressACLAnalysis)
	respItem["wireless_lan_controller_id"] = item.WirelessLanControllerID
	respItem["wireless_lan_controller_name"] = item.WirelessLanControllerName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectEgressACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectIngressACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceInterfaceStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfo) []map[string]interface{} {
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
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressPhysicalInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressPhysicalInterfaceQosStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceInterfaceStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfacePathOverlayInfo) []map[string]interface{} {
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
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsIngressVirtualInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsIngressVirtualInterfaceQosStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsPerfMonStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsPerfMonStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["accuracy_list"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoAccuracyList(item.AccuracyList)
		respItem["detailed_status"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoDetailedStatus(item.DetailedStatus)
		respItem["device_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoDeviceStatistics(item.DeviceStatistics)
		respItem["device_stats_collection"] = item.DeviceStatsCollection
		respItem["device_stats_collection_failure_reason"] = item.DeviceStatsCollectionFailureReason
		respItem["egress_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterface(item.EgressInterface)
		respItem["flex_connect"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnect(item.FlexConnect)
		respItem["id"] = item.ID
		respItem["ingress_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterface(item.IngressInterface)
		respItem["ip"] = item.IP
		respItem["link_information_source"] = item.LinkInformationSource
		respItem["name"] = item.Name
		respItem["perf_mon_collection"] = item.PerfMonCollection
		respItem["perf_mon_collection_failure_reason"] = item.PerfMonCollectionFailureReason
		respItem["perf_monitor_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoPerfMonitorStatistics(item.PerfMonitorStatistics)
		respItem["role"] = item.Role
		respItem["ssid"] = item.SSID
		respItem["tunnels"] = item.Tunnels
		respItem["type"] = item.Type
		respItem["wlan_id"] = item.WLANID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoAccuracyList(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoAccuracyList) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoDetailedStatus(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDetailedStatus) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoDeviceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatistics) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cpu_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoDeviceStatisticsCPUStatistics(item.CPUStatistics)
	respItem["memory_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoDeviceStatisticsMemoryStatistics(item.MemoryStatistics)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoDeviceStatisticsCPUStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatisticsCPUStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoDeviceStatisticsMemoryStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoDeviceStatisticsMemoryStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["physical_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterface(item.PhysicalInterface)
	respItem["virtual_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterface(item.VirtualInterface)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceInterfaceStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfo) []map[string]interface{} {
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
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfacePhysicalInterfaceQosStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterface(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterface) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis(item.ACLAnalysis)
		respItem["id"] = item.ID
		respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics(item.InterfaceStatistics)
		respItem["interface_stats_collection"] = item.InterfaceStatsCollection
		respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
		respItem["name"] = item.Name
		respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo(item.PathOverlayInfo)
		respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics(item.QosStatistics)
		respItem["qos_stats_collection"] = item.QosStatsCollection
		respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
		respItem["used_vlan"] = item.UsedVLAN
		respItem["vrf_name"] = item.VrfName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceInterfaceStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfo) []map[string]interface{} {
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
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoEgressInterfaceVirtualInterfaceQosStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnect(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnect) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["authentication"] = item.Authentication
	respItem["data_switching"] = item.DataSwitching
	respItem["egress_acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectEgressACLAnalysis(item.EgressACLAnalysis)
	respItem["ingress_acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectIngressACLAnalysis(item.IngressACLAnalysis)
	respItem["wireless_lan_controller_id"] = item.WirelessLanControllerID
	respItem["wireless_lan_controller_name"] = item.WirelessLanControllerName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectEgressACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectEgressACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectIngressACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoFlexConnectIngressACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["physical_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterface(item.PhysicalInterface)
	respItem["virtual_interface"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterface(item.VirtualInterface)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterface(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterface) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis(item.ACLAnalysis)
	respItem["id"] = item.ID
	respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics(item.InterfaceStatistics)
	respItem["interface_stats_collection"] = item.InterfaceStatsCollection
	respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
	respItem["name"] = item.Name
	respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo(item.PathOverlayInfo)
	respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics(item.QosStatistics)
	respItem["qos_stats_collection"] = item.QosStatsCollection
	respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
	respItem["used_vlan"] = item.UsedVLAN
	respItem["vrf_name"] = item.VrfName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceInterfaceStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfo) []map[string]interface{} {
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
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfacePhysicalInterfaceQosStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterface(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterface) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["acl_analysis"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis(item.ACLAnalysis)
		respItem["id"] = item.ID
		respItem["interface_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics(item.InterfaceStatistics)
		respItem["interface_stats_collection"] = item.InterfaceStatsCollection
		respItem["interface_stats_collection_failure_reason"] = item.InterfaceStatsCollectionFailureReason
		respItem["name"] = item.Name
		respItem["path_overlay_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo(item.PathOverlayInfo)
		respItem["qos_statistics"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics(item.QosStatistics)
		respItem["qos_stats_collection"] = item.QosStatsCollection
		respItem["qos_stats_collection_failure_reason"] = item.QosStatsCollectionFailureReason
		respItem["used_vlan"] = item.UsedVLAN
		respItem["vrf_name"] = item.VrfName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysis) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["acl_name"] = item.ACLName
	respItem["matching_aces"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces(item.MatchingAces)
	respItem["result"] = item.Result

	return []map[string]interface{}{
		respItem,
	}

}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ace"] = item.Ace
		respItem["matching_ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(item.MatchingPorts)
		respItem["result"] = item.Result
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPorts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ports"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(item.Ports)
		respItem["protocol"] = item.Protocol
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceACLAnalysisMatchingAcesMatchingPortsPorts) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceInterfaceStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfo) []map[string]interface{} {
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
		respItem["vxlan_info"] = flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo(item.VxlanInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfacePathOverlayInfoVxlanInfo) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoIngressInterfaceVirtualInterfaceQosStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemNetworkElementsInfoPerfMonitorStatistics(items *[]catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseNetworkElementsInfoPerfMonitorStatistics) []map[string]interface{} {
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

func flattenPathTraceRetrievesPreviousPathtraceV1ItemRequest(item *catalystcentersdkgo.ResponsePathTraceRetrievesPreviousPathtraceV1ResponseRequest) []map[string]interface{} {
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
