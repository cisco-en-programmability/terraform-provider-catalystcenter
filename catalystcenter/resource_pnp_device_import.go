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
func resourcePnpDeviceImport() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Add devices to PnP in bulk
`,

		CreateContext: resourcePnpDeviceImportCreate,
		ReadContext:   resourcePnpDeviceImportRead,
		DeleteContext: resourcePnpDeviceImportDelete,
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

						"failure_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"index": &schema.Schema{
										Description: `Index`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"msg": &schema.Schema{
										Description: `Msg`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"serial_num": &schema.Schema{
										Description: `Serial Num`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"success_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"day_zero_config": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config": &schema.Schema{
													Description: `Config`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"day_zero_config_preview": &schema.Schema{
										Description: `Day Zero Config Preview`,
										Type:        schema.TypeString, //TEST,
										Computed:    true,
									},
									"device_info": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"aaa_credentials": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"password": &schema.Schema{
																Description: `Password`,
																Type:        schema.TypeString,
																Sensitive:   true,
																Computed:    true,
															},
															"username": &schema.Schema{
																Description: `Username`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"added_on": &schema.Schema{
													Description: `Added On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"addn_mac_addrs": &schema.Schema{
													Description: `Addn Mac Addrs`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"agent_type": &schema.Schema{
													Description: `Agent Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"auth_status": &schema.Schema{
													Description: `Auth Status`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"authenticated_mic_number": &schema.Schema{
													Description: `Authenticated Mic Number`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"authenticated_sudi_serial_no": &schema.Schema{
													Description: `Authenticated Sudi Serial No`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"capabilities_supported": &schema.Schema{
													Description: `Capabilities Supported`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"cm_state": &schema.Schema{
													Description: `Cm State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"device_sudi_serial_nos": &schema.Schema{
													Description: `Device Sudi Serial Nos`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"device_type": &schema.Schema{
													Description: `Device Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"features_supported": &schema.Schema{
													Description: `Features Supported`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"file_system_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"freespace": &schema.Schema{
																Description: `Freespace`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"readable": &schema.Schema{
																Description: `Readable`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"size": &schema.Schema{
																Description: `Size`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"writeable": &schema.Schema{
																Description: `Writeable`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},
												"first_contact": &schema.Schema{
													Description: `First Contact`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"hostname": &schema.Schema{
													Description: `Hostname`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"http_headers": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"key": &schema.Schema{
																Description: `Key`,
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
												"image_file": &schema.Schema{
													Description: `Image File`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"image_version": &schema.Schema{
													Description: `Image Version`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"ip_interfaces": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"ipv4_address": &schema.Schema{
																Description: `Ipv4 Address`,
																Type:        schema.TypeString, //TEST,
																Computed:    true,
															},
															"ipv6_address_list": &schema.Schema{
																Description: `Ipv6 Address List`,
																Type:        schema.TypeList,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"mac_address": &schema.Schema{
																Description: `Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"status": &schema.Schema{
																Description: `Status`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"last_contact": &schema.Schema{
													Description: `Last Contact`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"last_sync_time": &schema.Schema{
													Description: `Last Sync Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"last_update_on": &schema.Schema{
													Description: `Last Update On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"location": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"address": &schema.Schema{
																Description: `Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"altitude": &schema.Schema{
																Description: `Altitude`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"latitude": &schema.Schema{
																Description: `Latitude`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"longitude": &schema.Schema{
																Description: `Longitude`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"site_id": &schema.Schema{
																Description: `Site Id`,
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
												"mode": &schema.Schema{
													Description: `Mode`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"neighbor_links": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"local_interface_name": &schema.Schema{
																Description: `Local Interface Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"local_mac_address": &schema.Schema{
																Description: `Local Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"local_short_interface_name": &schema.Schema{
																Description: `Local Short Interface Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_device_name": &schema.Schema{
																Description: `Remote Device Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_interface_name": &schema.Schema{
																Description: `Remote Interface Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_mac_address": &schema.Schema{
																Description: `Remote Mac Address`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_platform": &schema.Schema{
																Description: `Remote Platform`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_short_interface_name": &schema.Schema{
																Description: `Remote Short Interface Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"remote_version": &schema.Schema{
																Description: `Remote Version`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"onb_state": &schema.Schema{
													Description: `Onb State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"pid": &schema.Schema{
													Description: `Pid`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"pnp_profile_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"created_by": &schema.Schema{
																Description: `Created By`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"discovery_created": &schema.Schema{
																Description: `Discovery Created`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"primary_endpoint": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"certificate": &schema.Schema{
																			Description: `Certificate`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"fqdn": &schema.Schema{
																			Description: `Fqdn`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"ipv4_address": &schema.Schema{
																			Description: `Ipv4 Address`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},
																		"ipv6_address": &schema.Schema{
																			Description: `Ipv6 Address`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},
																		"port": &schema.Schema{
																			Description: `Port`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"protocol": &schema.Schema{
																			Description: `Protocol`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
															"profile_name": &schema.Schema{
																Description: `Profile Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"secondary_endpoint": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"certificate": &schema.Schema{
																			Description: `Certificate`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"fqdn": &schema.Schema{
																			Description: `Fqdn`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"ipv4_address": &schema.Schema{
																			Description: `Ipv4 Address`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},
																		"ipv6_address": &schema.Schema{
																			Description: `Ipv6 Address`,
																			Type:        schema.TypeString, //TEST,
																			Computed:    true,
																		},
																		"port": &schema.Schema{
																			Description: `Port`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"protocol": &schema.Schema{
																			Description: `Protocol`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"populate_inventory": &schema.Schema{
													Description: `Populate Inventory`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"pre_workflow_cli_ouputs": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"cli": &schema.Schema{
																Description: `Cli`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"cli_output": &schema.Schema{
																Description: `Cli Output`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"project_id": &schema.Schema{
													Description: `Project Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"project_name": &schema.Schema{
													Description: `Project Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"reload_requested": &schema.Schema{
													Description: `Reload Requested`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"serial_number": &schema.Schema{
													Description: `Serial Number`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"site_id": &schema.Schema{
													Description: `Site Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"site_name": &schema.Schema{
													Description: `Site Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"smart_account_id": &schema.Schema{
													Description: `Smart Account Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"source": &schema.Schema{
													Description: `Source`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"stack": &schema.Schema{
													Description: `Stack`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"stack_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"is_full_ring": &schema.Schema{
																Description: `Is Full Ring`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"stack_member_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"hardware_version": &schema.Schema{
																			Description: `Hardware Version`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"license_level": &schema.Schema{
																			Description: `License Level`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"license_type": &schema.Schema{
																			Description: `License Type`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"mac_address": &schema.Schema{
																			Description: `Mac Address`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"pid": &schema.Schema{
																			Description: `Pid`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"priority": &schema.Schema{
																			Description: `Priority`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"role": &schema.Schema{
																			Description: `Role`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"serial_number": &schema.Schema{
																			Description: `Serial Number`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"software_version": &schema.Schema{
																			Description: `Software Version`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"stack_number": &schema.Schema{
																			Description: `Stack Number`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"sudi_serial_number": &schema.Schema{
																			Description: `Sudi Serial Number`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																	},
																},
															},
															"stack_ring_protocol": &schema.Schema{
																Description: `Stack Ring Protocol`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"supports_stack_workflows": &schema.Schema{
																Description: `Supports Stack Workflows`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},
															"total_member_count": &schema.Schema{
																Description: `Total Member Count`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"valid_license_levels": &schema.Schema{
																Description: `Valid License Levels`,
																Type:        schema.TypeList,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"sudi_required": &schema.Schema{
													Description: `Sudi Required`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"tags": &schema.Schema{
													Description: `Tags`,
													Type:        schema.TypeString, //TEST,
													Computed:    true,
												},
												"user_mic_numbers": &schema.Schema{
													Description: `User Mic Numbers`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"user_sudi_serial_nos": &schema.Schema{
													Description: `User Sudi Serial Nos`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"virtual_account_id": &schema.Schema{
													Description: `Virtual Account Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"workflow_id": &schema.Schema{
													Description: `Workflow Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"workflow_name": &schema.Schema{
													Description: `Workflow Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"run_summary_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"details": &schema.Schema{
													Description: `Details`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"error_flag": &schema.Schema{
													Description: `Error Flag`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"history_task_info": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"addn_details": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"key": &schema.Schema{
																			Description: `Key`,
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
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"time_taken": &schema.Schema{
																Description: `Time Taken`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"work_item_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"command": &schema.Schema{
																			Description: `Command`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"end_time": &schema.Schema{
																			Description: `End Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"output_str": &schema.Schema{
																			Description: `Output Str`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"start_time": &schema.Schema{
																			Description: `Start Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"time_taken": &schema.Schema{
																			Description: `Time Taken`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"timestamp": &schema.Schema{
													Description: `Timestamp`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"system_reset_workflow": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"add_to_inventory": &schema.Schema{
													Description: `Add To Inventory`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"added_on": &schema.Schema{
													Description: `Added On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"config_id": &schema.Schema{
													Description: `Config Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"curr_task_idx": &schema.Schema{
													Description: `Curr Task Idx`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"exec_time": &schema.Schema{
													Description: `Exec Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"image_id": &schema.Schema{
													Description: `Image Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"instance_type": &schema.Schema{
													Description: `Instance Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"lastupdate_on": &schema.Schema{
													Description: `Lastupdate On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tasks": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"curr_work_item_idx": &schema.Schema{
																Description: `Curr Work Item Idx`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"end_time": &schema.Schema{
																Description: `End Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"start_time": &schema.Schema{
																Description: `Start Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"state": &schema.Schema{
																Description: `State`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"task_seq_no": &schema.Schema{
																Description: `Task Seq No`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"time_taken": &schema.Schema{
																Description: `Time Taken`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"work_item_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"command": &schema.Schema{
																			Description: `Command`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"end_time": &schema.Schema{
																			Description: `End Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"output_str": &schema.Schema{
																			Description: `Output Str`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"start_time": &schema.Schema{
																			Description: `Start Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"time_taken": &schema.Schema{
																			Description: `Time Taken`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"tenant_id": &schema.Schema{
													Description: `Tenant Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_state": &schema.Schema{
													Description: `Use State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"version": &schema.Schema{
													Description: `Version`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"system_workflow": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"add_to_inventory": &schema.Schema{
													Description: `Add To Inventory`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"added_on": &schema.Schema{
													Description: `Added On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"config_id": &schema.Schema{
													Description: `Config Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"curr_task_idx": &schema.Schema{
													Description: `Curr Task Idx`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"exec_time": &schema.Schema{
													Description: `Exec Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"image_id": &schema.Schema{
													Description: `Image Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"instance_type": &schema.Schema{
													Description: `Instance Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"lastupdate_on": &schema.Schema{
													Description: `Lastupdate On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tasks": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"curr_work_item_idx": &schema.Schema{
																Description: `Curr Work Item Idx`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"end_time": &schema.Schema{
																Description: `End Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"start_time": &schema.Schema{
																Description: `Start Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"state": &schema.Schema{
																Description: `State`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"task_seq_no": &schema.Schema{
																Description: `Task Seq No`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"time_taken": &schema.Schema{
																Description: `Time Taken`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"work_item_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"command": &schema.Schema{
																			Description: `Command`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"end_time": &schema.Schema{
																			Description: `End Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"output_str": &schema.Schema{
																			Description: `Output Str`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"start_time": &schema.Schema{
																			Description: `Start Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"time_taken": &schema.Schema{
																			Description: `Time Taken`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"tenant_id": &schema.Schema{
													Description: `Tenant Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_state": &schema.Schema{
													Description: `Use State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"version": &schema.Schema{
													Description: `Version`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"tenant_id": &schema.Schema{
										Description: `Tenant Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"version": &schema.Schema{
										Description: `Version`,
										Type:        schema.TypeFloat,
										Computed:    true,
									},
									"workflow": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"add_to_inventory": &schema.Schema{
													Description: `Add To Inventory`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},
												"added_on": &schema.Schema{
													Description: `Added On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"config_id": &schema.Schema{
													Description: `Config Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"curr_task_idx": &schema.Schema{
													Description: `Curr Task Idx`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"end_time": &schema.Schema{
													Description: `End Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"exec_time": &schema.Schema{
													Description: `Exec Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"image_id": &schema.Schema{
													Description: `Image Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"instance_type": &schema.Schema{
													Description: `Instance Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"lastupdate_on": &schema.Schema{
													Description: `Lastupdate On`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"start_time": &schema.Schema{
													Description: `Start Time`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
												"state": &schema.Schema{
													Description: `State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"tasks": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"curr_work_item_idx": &schema.Schema{
																Description: `Curr Work Item Idx`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"end_time": &schema.Schema{
																Description: `End Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"start_time": &schema.Schema{
																Description: `Start Time`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"state": &schema.Schema{
																Description: `State`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"task_seq_no": &schema.Schema{
																Description: `Task Seq No`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"time_taken": &schema.Schema{
																Description: `Time Taken`,
																Type:        schema.TypeFloat,
																Computed:    true,
															},
															"type": &schema.Schema{
																Description: `Type`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"work_item_list": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"command": &schema.Schema{
																			Description: `Command`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"end_time": &schema.Schema{
																			Description: `End Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"output_str": &schema.Schema{
																			Description: `Output Str`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"start_time": &schema.Schema{
																			Description: `Start Time`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Computed:    true,
																		},
																		"time_taken": &schema.Schema{
																			Description: `Time Taken`,
																			Type:        schema.TypeFloat,
																			Computed:    true,
																		},
																	},
																},
															},
														},
													},
												},
												"tenant_id": &schema.Schema{
													Description: `Tenant Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"use_state": &schema.Schema{
													Description: `Use State`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"version": &schema.Schema{
													Description: `Version`,
													Type:        schema.TypeFloat,
													Computed:    true,
												},
											},
										},
									},
									"workflow_parameters": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"config_list": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"config_id": &schema.Schema{
																Description: `Config Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"config_parameters": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"key": &schema.Schema{
																			Description: `Key`,
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
														},
													},
												},
												"license_level": &schema.Schema{
													Description: `License Level`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"license_type": &schema.Schema{
													Description: `License Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"top_of_stack_serial_number": &schema.Schema{
													Description: `Top Of Stack Serial Number`,
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
						"payload": &schema.Schema{
							Description: `Array of RequestDeviceOnboardingPnpImportDevicesInBulk`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"device_info": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"device_sudi_serial_nos": &schema.Schema{
													Description: `Device Sudi Serial Nos`,
													Type:        schema.TypeList,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"hostname": &schema.Schema{
													Description: `Hostname`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"mac_address": &schema.Schema{
													Description: `Mac Address`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"pid": &schema.Schema{
													Description: `Pid`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"serial_number": &schema.Schema{
													Description: `Serial Number`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"site_id": &schema.Schema{
													Description: `Site Id`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"stack": &schema.Schema{
													Description: `Stack`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													ForceNew:     true,
													Computed:     true,
												},
												"stack_info": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"is_full_ring": &schema.Schema{
																Description: `Is Full Ring`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"stack_member_list": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"hardware_version": &schema.Schema{
																			Description: `Hardware Version`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"license_level": &schema.Schema{
																			Description: `License Level`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"license_type": &schema.Schema{
																			Description: `License Type`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"mac_address": &schema.Schema{
																			Description: `Mac Address`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"pid": &schema.Schema{
																			Description: `Pid`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"priority": &schema.Schema{
																			Description: `Priority`,
																			Type:        schema.TypeFloat,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"role": &schema.Schema{
																			Description: `Role`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"serial_number": &schema.Schema{
																			Description: `Serial Number`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"software_version": &schema.Schema{
																			Description: `Software Version`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"stack_number": &schema.Schema{
																			Description: `Stack Number`,
																			Type:        schema.TypeFloat,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"state": &schema.Schema{
																			Description: `State`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																		"sudi_serial_number": &schema.Schema{
																			Description: `Sudi Serial Number`,
																			Type:        schema.TypeString,
																			Optional:    true,
																			ForceNew:    true,
																			Computed:    true,
																		},
																	},
																},
															},
															"stack_ring_protocol": &schema.Schema{
																Description: `Stack Ring Protocol`,
																Type:        schema.TypeString,
																Optional:    true,
																ForceNew:    true,
																Computed:    true,
															},
															"supports_stack_workflows": &schema.Schema{
																Description: `Supports Stack Workflows`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"total_member_count": &schema.Schema{
																Description: `Total Member Count`,
																Type:        schema.TypeFloat,
																Optional:    true,
																ForceNew:    true,
																Computed:    true,
															},
															"valid_license_levels": &schema.Schema{
																Description: `Valid License Levels`,
																Type:        schema.TypeList,
																Optional:    true,
																ForceNew:    true,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"sudi_required": &schema.Schema{
													Description: `Is Sudi Required`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													ForceNew:     true,
													Computed:     true,
												},
												"user_mic_numbers": &schema.Schema{
													Description: `User Mic Numbers`,
													Type:        schema.TypeList,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"user_sudi_serial_nos": &schema.Schema{
													Description: `User Sudi Serial Nos`,
													Type:        schema.TypeList,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"workflow_id": &schema.Schema{
													Description: `Workflow Id`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"workflow_name": &schema.Schema{
													Description: `Workflow Name`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
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

func resourcePnpDeviceImportCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestPnpDeviceImportImportDevicesInBulk(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.DeviceOnboardingPnp.ImportDevicesInBulk(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ImportDevicesInBulk", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDeviceOnboardingPnpImportDevicesInBulkItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ImportDevicesInBulk response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourcePnpDeviceImportRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpDeviceImportDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestPnpDeviceImportImportDevicesInBulk(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpImportDevicesInBulk {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpImportDevicesInBulk{}
	if v := expandRequestPnpDeviceImportImportDevicesInBulkItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulk {
	request := []catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulk{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulk {
	request := catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulk{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.TypeID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_info")))) {
		request.DeviceInfo = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfo(ctx, key+".device_info.0", d)
	}
	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfo {
	request := catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack")))) {
		request.Stack = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_required")))) {
		request.SudiRequired = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_sudi_serial_nos")))) {
		request.DeviceSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_mic_numbers")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_mic_numbers")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_mic_numbers")))) {
		request.UserMicNumbers = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".user_sudi_serial_nos")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".user_sudi_serial_nos")))) {
		request.UserSudiSerialNos = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_id")))) {
		request.WorkflowID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".workflow_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".workflow_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".workflow_name")))) {
		request.WorkflowName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_info")))) {
		request.StackInfo = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfo(ctx, key+".stack_info.0", d)
	}
	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfo {
	request := catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".supports_stack_workflows")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".supports_stack_workflows")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".supports_stack_workflows")))) {
		request.SupportsStackWorkflows = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_full_ring")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_full_ring")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_full_ring")))) {
		request.IsFullRing = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_member_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_member_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_member_list")))) {
		request.StackMemberList = expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfoStackMemberListArray(ctx, key+".stack_member_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_ring_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_ring_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_ring_protocol")))) {
		request.StackRingProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".valid_license_levels")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".valid_license_levels")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".valid_license_levels")))) {
		request.ValidLicenseLevels = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".total_member_count")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".total_member_count")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".total_member_count")))) {
		request.TotalMemberCount = interfaceToFloat64Ptr(v)
	}
	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfoStackMemberListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList {
	request := []catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList{}
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
		i := expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfoStackMemberList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceImportImportDevicesInBulkItemDeviceInfoStackInfoStackMemberList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList {
	request := catalystcentersdkgo.RequestItemDeviceOnboardingPnpImportDevicesInBulkDeviceInfoStackInfoStackMemberList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".serial_number")))) {
		request.SerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".state")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".state")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".state")))) {
		request.State = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".mac_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".mac_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".mac_address")))) {
		request.MacAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".pid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".pid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".pid")))) {
		request.Pid = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_level")))) {
		request.LicenseLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".license_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".license_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".license_type")))) {
		request.LicenseType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sudi_serial_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sudi_serial_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sudi_serial_number")))) {
		request.SudiSerialNumber = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hardware_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hardware_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hardware_version")))) {
		request.HardwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".stack_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".stack_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".stack_number")))) {
		request.StackNumber = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToFloat64Ptr(v)
	}
	return &request
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItem(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulk) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["success_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessList(item.SuccessList)
	respItem["failure_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemFailureList(item.FailureList)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["device_info"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfo(item.DeviceInfo)
		respItem["system_reset_workflow"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflow(item.SystemResetWorkflow)
		respItem["system_workflow"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflow(item.SystemWorkflow)
		respItem["workflow"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflow(item.Workflow)
		respItem["run_summary_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryList(item.RunSummaryList)
		respItem["workflow_parameters"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParameters(item.WorkflowParameters)
		respItem["day_zero_config"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDayZeroConfig(item.DayZeroConfig)
		respItem["day_zero_config_preview"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDayZeroConfigPreview(item.DayZeroConfigPreview)
		respItem["version"] = item.Version
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfo(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["source"] = item.Source
	respItem["serial_number"] = item.SerialNumber
	respItem["stack"] = boolPtrToString(item.Stack)
	respItem["mode"] = item.Mode
	respItem["state"] = item.State
	respItem["location"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoLocation(item.Location)
	respItem["description"] = item.Description
	respItem["onb_state"] = item.OnbState
	respItem["authenticated_mic_number"] = item.AuthenticatedMicNumber
	respItem["authenticated_sudi_serial_no"] = item.AuthenticatedSudiSerialNo
	respItem["capabilities_supported"] = item.CapabilitiesSupported
	respItem["features_supported"] = item.FeaturesSupported
	respItem["cm_state"] = item.CmState
	respItem["first_contact"] = item.FirstContact
	respItem["last_contact"] = item.LastContact
	respItem["mac_address"] = item.MacAddress
	respItem["pid"] = item.Pid
	respItem["device_sudi_serial_nos"] = item.DeviceSudiSerialNos
	respItem["last_update_on"] = item.LastUpdateOn
	respItem["workflow_id"] = item.WorkflowID
	respItem["workflow_name"] = item.WorkflowName
	respItem["project_id"] = item.ProjectID
	respItem["project_name"] = item.ProjectName
	respItem["device_type"] = item.DeviceType
	respItem["agent_type"] = item.AgentType
	respItem["image_version"] = item.ImageVersion
	respItem["file_system_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoFileSystemList(item.FileSystemList)
	respItem["pnp_profile_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileList(item.PnpProfileList)
	respItem["image_file"] = item.ImageFile
	respItem["http_headers"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoHTTPHeaders(item.HTTPHeaders)
	respItem["neighbor_links"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoNeighborLinks(item.NeighborLinks)
	respItem["last_sync_time"] = item.LastSyncTime
	respItem["ip_interfaces"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfaces(item.IPInterfaces)
	respItem["hostname"] = item.Hostname
	respItem["auth_status"] = item.AuthStatus
	respItem["stack_info"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoStackInfo(item.StackInfo)
	respItem["reload_requested"] = boolPtrToString(item.ReloadRequested)
	respItem["added_on"] = item.AddedOn
	respItem["site_id"] = item.SiteID
	respItem["aaa_credentials"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoAAACredentials(item.AAACredentials)
	respItem["user_mic_numbers"] = item.UserMicNumbers
	respItem["user_sudi_serial_nos"] = item.UserSudiSerialNos
	respItem["addn_mac_addrs"] = item.AddnMacAddrs
	respItem["pre_workflow_cli_ouputs"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPreWorkflowCliOuputs(item.PreWorkflowCliOuputs)
	respItem["tags"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoTags(item.Tags)
	respItem["sudi_required"] = boolPtrToString(item.SudiRequired)
	respItem["smart_account_id"] = item.SmartAccountID
	respItem["virtual_account_id"] = item.VirtualAccountID
	respItem["populate_inventory"] = boolPtrToString(item.PopulateInventory)
	respItem["site_name"] = item.SiteName
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoLocation(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoLocation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site_id"] = item.SiteID
	respItem["address"] = item.Address
	respItem["latitude"] = item.Latitude
	respItem["longitude"] = item.Longitude
	respItem["altitude"] = item.Altitude

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoFileSystemList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoFileSystemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["writeable"] = boolPtrToString(item.Writeable)
		respItem["freespace"] = item.Freespace
		respItem["name"] = item.Name
		respItem["readable"] = boolPtrToString(item.Readable)
		respItem["size"] = item.Size
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["profile_name"] = item.ProfileName
		respItem["discovery_created"] = boolPtrToString(item.DiscoveryCreated)
		respItem["created_by"] = item.CreatedBy
		respItem["primary_endpoint"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpoint(item.PrimaryEndpoint)
		respItem["secondary_endpoint"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpoint(item.SecondaryEndpoint)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpoint(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol
	respItem["ipv4_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(item.IPv4Address)
	respItem["ipv6_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(item.IPv6Address)
	respItem["fqdn"] = item.Fqdn
	respItem["certificate"] = item.Certificate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListPrimaryEndpointIPv6Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpoint(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpoint) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["port"] = item.Port
	respItem["protocol"] = item.Protocol
	respItem["ipv4_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(item.IPv4Address)
	respItem["ipv6_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(item.IPv6Address)
	respItem["fqdn"] = item.Fqdn
	respItem["certificate"] = item.Certificate

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPnpProfileListSecondaryEndpointIPv6Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoHTTPHeaders(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoHTTPHeaders) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoNeighborLinks(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoNeighborLinks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["local_interface_name"] = item.LocalInterfaceName
		respItem["local_short_interface_name"] = item.LocalShortInterfaceName
		respItem["local_mac_address"] = item.LocalMacAddress
		respItem["remote_interface_name"] = item.RemoteInterfaceName
		respItem["remote_short_interface_name"] = item.RemoteShortInterfaceName
		respItem["remote_mac_address"] = item.RemoteMacAddress
		respItem["remote_device_name"] = item.RemoteDeviceName
		respItem["remote_platform"] = item.RemotePlatform
		respItem["remote_version"] = item.RemoteVersion
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfaces(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfaces) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["status"] = item.Status
		respItem["mac_address"] = item.MacAddress
		respItem["ipv4_address"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfacesIPv4Address(item.IPv4Address)
		respItem["ipv6_address_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfacesIPv6AddressList(item.IPv6AddressList)
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfacesIPv4Address(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfacesIPv4Address) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoIPInterfacesIPv6AddressList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoIPInterfacesIPv6AddressList) []interface{} {
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

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoStackInfo(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoStackInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["supports_stack_workflows"] = boolPtrToString(item.SupportsStackWorkflows)
	respItem["is_full_ring"] = boolPtrToString(item.IsFullRing)
	respItem["stack_member_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoStackInfoStackMemberList(item.StackMemberList)
	respItem["stack_ring_protocol"] = item.StackRingProtocol
	respItem["valid_license_levels"] = item.ValidLicenseLevels
	respItem["total_member_count"] = item.TotalMemberCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoStackInfoStackMemberList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoStackInfoStackMemberList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["serial_number"] = item.SerialNumber
		respItem["state"] = item.State
		respItem["role"] = item.Role
		respItem["mac_address"] = item.MacAddress
		respItem["pid"] = item.Pid
		respItem["license_level"] = item.LicenseLevel
		respItem["license_type"] = item.LicenseType
		respItem["sudi_serial_number"] = item.SudiSerialNumber
		respItem["hardware_version"] = item.HardwareVersion
		respItem["stack_number"] = item.StackNumber
		respItem["software_version"] = item.SoftwareVersion
		respItem["priority"] = item.Priority
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoAAACredentials(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoAAACredentials) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["password"] = item.Password
	respItem["username"] = item.Username

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoPreWorkflowCliOuputs(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoPreWorkflowCliOuputs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["cli"] = item.Cli
		respItem["cli_output"] = item.CliOutput
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDeviceInfoTags(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDeviceInfoTags) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflow(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["state"] = item.State
	respItem["type"] = item.Type
	respItem["description"] = item.Description
	respItem["lastupdate_on"] = item.LastupdateOn
	respItem["image_id"] = item.ImageID
	respItem["curr_task_idx"] = item.CurrTaskIDx
	respItem["added_on"] = item.AddedOn
	respItem["tasks"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflowTasks(item.Tasks)
	respItem["add_to_inventory"] = boolPtrToString(item.AddToInventory)
	respItem["instance_type"] = item.InstanceType
	respItem["end_time"] = item.EndTime
	respItem["exec_time"] = item.ExecTime
	respItem["start_time"] = item.StartTime
	respItem["use_state"] = item.UseState
	respItem["config_id"] = item.ConfigID
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflowTasks(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflowTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["curr_work_item_idx"] = item.CurrWorkItemIDx
		respItem["task_seq_no"] = item.TaskSeqNo
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["work_item_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemResetWorkflowTasksWorkItemList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemResetWorkflowTasksWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflow(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["state"] = item.State
	respItem["type"] = item.Type
	respItem["description"] = item.Description
	respItem["lastupdate_on"] = item.LastupdateOn
	respItem["image_id"] = item.ImageID
	respItem["curr_task_idx"] = item.CurrTaskIDx
	respItem["added_on"] = item.AddedOn
	respItem["tasks"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflowTasks(item.Tasks)
	respItem["add_to_inventory"] = boolPtrToString(item.AddToInventory)
	respItem["instance_type"] = item.InstanceType
	respItem["end_time"] = item.EndTime
	respItem["exec_time"] = item.ExecTime
	respItem["start_time"] = item.StartTime
	respItem["use_state"] = item.UseState
	respItem["config_id"] = item.ConfigID
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflowTasks(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflowTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["curr_work_item_idx"] = item.CurrWorkItemIDx
		respItem["task_seq_no"] = item.TaskSeqNo
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["work_item_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListSystemWorkflowTasksWorkItemList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListSystemWorkflowTasksWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflow(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflow) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["state"] = item.State
	respItem["type"] = item.Type
	respItem["description"] = item.Description
	respItem["lastupdate_on"] = item.LastupdateOn
	respItem["image_id"] = item.ImageID
	respItem["curr_task_idx"] = item.CurrTaskIDx
	respItem["added_on"] = item.AddedOn
	respItem["tasks"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowTasks(item.Tasks)
	respItem["add_to_inventory"] = boolPtrToString(item.AddToInventory)
	respItem["instance_type"] = item.InstanceType
	respItem["end_time"] = item.EndTime
	respItem["exec_time"] = item.ExecTime
	respItem["start_time"] = item.StartTime
	respItem["use_state"] = item.UseState
	respItem["config_id"] = item.ConfigID
	respItem["name"] = item.Name
	respItem["version"] = item.Version
	respItem["tenant_id"] = item.TenantID

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowTasks(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowTasks) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["type"] = item.Type
		respItem["curr_work_item_idx"] = item.CurrWorkItemIDx
		respItem["task_seq_no"] = item.TaskSeqNo
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["work_item_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowTasksWorkItemList(item.WorkItemList)
		respItem["time_taken"] = item.TimeTaken
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowTasksWorkItemList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowTasksWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["details"] = item.Details
		respItem["history_task_info"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfo(item.HistoryTaskInfo)
		respItem["error_flag"] = boolPtrToString(item.ErrorFlag)
		respItem["timestamp"] = item.Timestamp
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfo(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfo) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["work_item_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfoWorkItemList(item.WorkItemList)
	respItem["time_taken"] = item.TimeTaken
	respItem["addn_details"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfoAddnDetails(item.AddnDetails)
	respItem["name"] = item.Name

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfoWorkItemList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfoWorkItemList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["state"] = item.State
		respItem["command"] = item.Command
		respItem["output_str"] = item.OutputStr
		respItem["end_time"] = item.EndTime
		respItem["start_time"] = item.StartTime
		respItem["time_taken"] = item.TimeTaken
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListRunSummaryListHistoryTaskInfoAddnDetails(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListRunSummaryListHistoryTaskInfoAddnDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParameters(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["top_of_stack_serial_number"] = item.TopOfStackSerialNumber
	respItem["license_level"] = item.LicenseLevel
	respItem["license_type"] = item.LicenseType
	respItem["config_list"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParametersConfigList(item.ConfigList)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParametersConfigList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParametersConfigList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["config_parameters"] = flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParametersConfigListConfigParameters(item.ConfigParameters)
		respItem["config_id"] = item.ConfigID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListWorkflowParametersConfigListConfigParameters(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListWorkflowParametersConfigListConfigParameters) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDayZeroConfig(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDayZeroConfig) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["config"] = item.Config

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemSuccessListDayZeroConfigPreview(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkSuccessListDayZeroConfigPreview) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDeviceOnboardingPnpImportDevicesInBulkItemFailureList(items *[]catalystcentersdkgo.ResponseDeviceOnboardingPnpImportDevicesInBulkFailureList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["index"] = item.Index
		respItem["serial_num"] = item.SerialNum
		respItem["id"] = item.ID
		respItem["msg"] = item.Msg
		respItems = append(respItems, respItem)
	}
	return respItems
}
