package catalystcenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAppPolicyQueuingProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Update existing custom application queuing profile

- Create new custom application queuing profile

- Delete existing custom application policy queuing profile by id
`,

		CreateContext: resourceAppPolicyQueuingProfileCreate,
		ReadContext:   resourceAppPolicyQueuingProfileRead,
		UpdateContext: resourceAppPolicyQueuingProfileUpdate,
		DeleteContext: resourceAppPolicyQueuingProfileDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

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

						"cfs_change_info": &schema.Schema{
							Description: `Cfs change info
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"clause": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"instance_version": &schema.Schema{
										Description: `Instance version
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"interface_speed_bandwidth_clauses": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `Id
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"instance_created_on": &schema.Schema{
													Description: `Instance created on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"instance_updated_on": &schema.Schema{
													Description: `Instance updated on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"instance_version": &schema.Schema{
													Description: `Instance version
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"interface_speed": &schema.Schema{
													Description: `Interface speed
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"tc_bandwidth_settings": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"bandwidth_percentage": &schema.Schema{
																Description: `Bandwidth percentage
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"display_name": &schema.Schema{
																Description: `Display name
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"id": &schema.Schema{
																Description: `Id
`,
																Type:     schema.TypeString,
																Computed: true,
															},
															"instance_created_on": &schema.Schema{
																Description: `Instance created on
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"instance_id": &schema.Schema{
																Description: `Instance id
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"instance_updated_on": &schema.Schema{
																Description: `Instance updated on
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
															"instance_version": &schema.Schema{
																Description: `Instance version
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},
															"traffic_class": &schema.Schema{
																Description: `Traffic Class
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
									"is_common_between_all_interface_speeds": &schema.Schema{
										Description: `Is common between all interface speeds
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},
									"priority": &schema.Schema{
										Description: `Priority
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"tc_dscp_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `Display name
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"dscp": &schema.Schema{
													Description: `Dscp value
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `Id
`,
													Type:     schema.TypeString,
													Computed: true,
												},
												"instance_created_on": &schema.Schema{
													Description: `Instance created on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"instance_updated_on": &schema.Schema{
													Description: `Instance updated on
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
												"instance_version": &schema.Schema{
													Description: `Instance version
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},
												"traffic_class": &schema.Schema{
													Description: `Traffic Class
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"type": &schema.Schema{
										Description: `Type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"contract_classifier": &schema.Schema{
							Description: `Contract classifier
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"create_time": &schema.Schema{
							Description: `Create time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"custom_provisions": &schema.Schema{
							Description: `Custom provisions
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"deployed": &schema.Schema{
							Description: `Deployed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Free test description
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": &schema.Schema{
							Description: `Display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"gen_id": &schema.Schema{
							Description: `Gen id
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Id of Queueing profile
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_created_on": &schema.Schema{
							Description: `Instance created on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instance_id": &schema.Schema{
							Description: `Instance id
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instance_updated_on": &schema.Schema{
							Description: `Instance updated on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instance_version": &schema.Schema{
							Description: `Instance version
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"internal": &schema.Schema{
							Description: `Internal
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_deleted": &schema.Schema{
							Description: `Is deleted
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_seeded": &schema.Schema{
							Description: `Is seeded
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_stale": &schema.Schema{
							Description: `Is stale
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"ise_reserved": &schema.Schema{
							Description: `Is reserved
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Last update time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Queueing profile name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": &schema.Schema{
							Description: `Namespace
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"provisioning_state": &schema.Schema{
							Description: `Provisioning State
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"pushed": &schema.Schema{
							Description: `Pushed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"qualifier": &schema.Schema{
							Description: `Qualifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_version": &schema.Schema{
							Description: `Resource version
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"target_id_list": &schema.Schema{
							Description: `Target id list
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Description: `Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationPolicyQueuingProfileV1`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"clause": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"interface_speed_bandwidth_clauses": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"instance_id": &schema.Schema{
																Description: `Instance id
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"interface_speed": &schema.Schema{
																Description: `Interface speed
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"tc_bandwidth_settings": &schema.Schema{
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"bandwidth_percentage": &schema.Schema{
																			Description: `Bandwidth percentage
`,
																			Type:     schema.TypeInt,
																			Optional: true,
																			Computed: true,
																		},
																		"instance_id": &schema.Schema{
																			Description: `Instance id
`,
																			Type:     schema.TypeInt,
																			Optional: true,
																			Computed: true,
																		},
																		"traffic_class": &schema.Schema{
																			Description: `Traffic Class
`,
																			Type:     schema.TypeString,
																			Optional: true,
																			Computed: true,
																		},
																	},
																},
															},
														},
													},
												},
												"is_common_between_all_interface_speeds": &schema.Schema{
													Description: `Is common between all interface speeds
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"tc_dscp_settings": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"dscp": &schema.Schema{
																Description: `Dscp value
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"instance_id": &schema.Schema{
																Description: `Instance id
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"traffic_class": &schema.Schema{
																Description: `Traffic Class
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"type": &schema.Schema{
													Description: `The allowed clause types are: BANDWIDTH, DSCP_CUSTOMIZATION
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"description": &schema.Schema{
										Description: `Free test description
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Id of Queueing profile
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"name": &schema.Schema{
										Description: `Queueing profile name
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
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

func resourceAppPolicyQueuingProfileCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	queryParamImport := catalystcentersdkgo.GetApplicationPolicyQueuingProfileV1QueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchApplicationPolicyGetApplicationPolicyQueuingProfileV1(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceAppPolicyQueuingProfileRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationPolicyQueuingProfileV1(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationPolicyQueuingProfileV1", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationPolicyQueuingProfileV1", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationPolicyQueuingProfileV1", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateApplicationPolicyQueuingProfileV1", err1))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetApplicationPolicyQueuingProfileV1QueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchApplicationPolicyGetApplicationPolicyQueuingProfileV1(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateApplicationPolicyQueuingProfileV1", err,
			"Failure at CreateApplicationPolicyQueuingProfileV1, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceAppPolicyQueuingProfileRead(ctx, d, m)
}

func resourceAppPolicyQueuingProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	vvName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplicationPolicyQueuingProfileV1")
		queryParams1 := catalystcentersdkgo.GetApplicationPolicyQueuingProfileV1QueryParams{}
		queryParams1.Name = vvName
		item1, err := searchApplicationPolicyGetApplicationPolicyQueuingProfileV1(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1Response{
			*item1,
		}
		vItem1 := flattenApplicationPolicyGetApplicationPolicyQueuingProfileV1Items(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationPolicyQueuingProfileV1 search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceAppPolicyQueuingProfileUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.ApplicationPolicy.UpdateApplicationPolicyQueuingProfile(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateApplicationPolicyQueuingProfile", err, restyResp1.String(),
					"Failure at UpdateApplicationPolicyQueuingProfile, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateApplicationPolicyQueuingProfile", err,
				"Failure at UpdateApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateApplicationPolicyQueuingProfile", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateApplicationPolicyQueuingProfile", err1))
				return diags
			}
		}

	}

	return resourceAppPolicyQueuingProfileRead(ctx, d, m)
}

func resourceAppPolicyQueuingProfileDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplicationPolicyQueuingProfile(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteApplicationPolicyQueuingProfile", err, restyResp1.String(),
				"Failure at DeleteApplicationPolicyQueuingProfile, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteApplicationPolicyQueuingProfile", err,
			"Failure at DeleteApplicationPolicyQueuingProfile, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteApplicationPolicyQueuingProfile", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing DeleteApplicationPolicyQueuingProfile", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileV1 {
	request := catalystcentersdkgo.RequestApplicationPolicyCreateApplicationPolicyQueuingProfileV1{}
	if v := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1 {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1 {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1Clause {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1Clause{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClause(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1Clause {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1Clause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileCreateApplicationPolicyQueuingProfileV1ItemClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 {
	request := catalystcentersdkgo.RequestApplicationPolicyUpdateApplicationPolicyQueuingProfileV1{}
	if v := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1 {
	request := catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseArray(ctx, key+".clause", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1Clause {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1Clause{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClause(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1Clause {
	request := catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1Clause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".is_common_between_all_interface_speeds")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".is_common_between_all_interface_speeds")))) {
		request.IsCommonBetweenAllInterfaceSpeeds = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed_bandwidth_clauses")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed_bandwidth_clauses")))) {
		request.InterfaceSpeedBandwidthClauses = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesArray(ctx, key+".interface_speed_bandwidth_clauses", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_dscp_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_dscp_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_dscp_settings")))) {
		request.TcDscpSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseTcDscpSettingsArray(ctx, key+".tc_dscp_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClauses(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClauses(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses {
	request := catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClauses{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_speed")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_speed")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_speed")))) {
		request.InterfaceSpeed = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tc_bandwidth_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tc_bandwidth_settings")))) {
		request.TcBandwidthSettings = expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx, key+".tc_bandwidth_settings", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings {
	request := catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseInterfaceSpeedBandwidthClausesTcBandwidthSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".bandwidth_percentage")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".bandwidth_percentage")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".bandwidth_percentage")))) {
		request.BandwidthPercentage = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseTcDscpSettingsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings{}
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
		i := expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseTcDscpSettings(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAppPolicyQueuingProfileUpdateApplicationPolicyQueuingProfileV1ItemClauseTcDscpSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings {
	request := catalystcentersdkgo.RequestItemApplicationPolicyUpdateApplicationPolicyQueuingProfileV1ClauseTcDscpSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchApplicationPolicyGetApplicationPolicyQueuingProfileV1(m interface{}, queryParams catalystcentersdkgo.GetApplicationPolicyQueuingProfileV1QueryParams, vID string) (*catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1Response, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1Response
	var ite *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1
	ite, _, err = client.ApplicationPolicy.GetApplicationPolicyQueuingProfileV1(&queryParams)
	if err != nil || ite == nil {
		return foundItem, err

	}
	items := ite
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyQueuingProfileV1Response
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
