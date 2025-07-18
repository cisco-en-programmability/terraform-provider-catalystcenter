package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceAppPolicyIntentCreate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Application Policy.

- Create/Update/Delete application policy
`,

		CreateContext: resourceAppPolicyIntentCreateCreate,
		ReadContext:   resourceAppPolicyIntentCreateRead,
		DeleteContext: resourceAppPolicyIntentCreateDelete,
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

						"task_id": &schema.Schema{
							Description: `Task id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Task url
`,
							Type:     schema.TypeString,
							Computed: true,
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
						"create_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_policy_scope": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_policy_scope_element": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"group_id": &schema.Schema{
																Description: `The site(s) ID where the Application QoS Policy will be deployed.
`,
																Type:     schema.TypeList,
																Optional: true,
																ForceNew: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"ssid": &schema.Schema{
																Description: `Ssid
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
												"name": &schema.Schema{
													Description: `Policy name
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"consumer": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"scalable_group": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id_ref": &schema.Schema{
																Description: `Id ref to application Scalable group
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"contract": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id ref to Queueing profile
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"delete_policy_status": &schema.Schema{
										Description: `NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"exclusive_contract": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"clause": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"device_removal_behavior": &schema.Schema{
																Description: `Device eemoval behavior
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"host_tracking_enabled": &schema.Schema{
																Description: `Is host tracking enabled
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"relevance_level": &schema.Schema{
																Description: `Relevance level
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"type": &schema.Schema{
																Description: `Type
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"policy_scope": &schema.Schema{
										Description: `Policy name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"priority": &schema.Schema{
										Description: `Set to 4095 while producer refer to application Scalable group otherwise 100
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"producer": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"scalable_group": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id_ref": &schema.Schema{
																Description: `Id ref to application-set or application Scalable group
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
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
						"delete_list": &schema.Schema{
							Description: `Delete list of Group Based Policy ids
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"update_list": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_policy_scope": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"advanced_policy_scope_element": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"group_id": &schema.Schema{
																Description: `The site(s) ID where the Application QoS Policy will be deployed.
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
																Description: `Id of Advance policy scope element
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"ssid": &schema.Schema{
																Description: `Ssid
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
												"id": &schema.Schema{
													Description: `Id of Advance policy scope
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Policy name
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"consumer": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id of Consumer
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"scalable_group": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id_ref": &schema.Schema{
																Description: `Id ref to application Scalable group
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"contract": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id ref to Queueing profile
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"delete_policy_status": &schema.Schema{
										Description: `NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"exclusive_contract": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"clause": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"device_removal_behavior": &schema.Schema{
																Description: `Device removal behavior
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"host_tracking_enabled": &schema.Schema{
																Description: `Host tracking enabled
`,
																// Type:        schema.TypeBool,
																Type:         schema.TypeString,
																ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
																Optional:     true,
																ForceNew:     true,
																Computed:     true,
															},
															"id": &schema.Schema{
																Description: `Id of Business relevance or Application policy knobs clause
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"relevance_level": &schema.Schema{
																Description: `Relevance level
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"type": &schema.Schema{
																Description: `Type
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Description: `Id of Exclusive contract
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `Id of Group based policy
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"policy_scope": &schema.Schema{
										Description: `Policy name
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"priority": &schema.Schema{
										Description: `Set to 4095 while producer refer to application Scalable group otherwise 100
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"producer": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Id of Producer
`,
													Type:     schema.TypeString,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"scalable_group": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id_ref": &schema.Schema{
																Description: `Id ref to application-set or application Scalable group
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
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
			},
		},
	}
}

func resourceAppPolicyIntentCreateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestAppPolicyIntentCreateApplicationPolicyIntent(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.ApplicationPolicy.ApplicationPolicyIntent(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ApplicationPolicyIntent", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ApplicationPolicyIntent", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil || strings.Contains(restyResp3.String(), "<!doctype html>") {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ApplicationPolicyIntent", err1))
			return diags
		}
	}

	vItem1 := flattenApplicationPolicyApplicationPolicyIntentItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ApplicationPolicyIntent response",
			err))
		return diags
	}
	d.SetId(getUnixTimeString())
	return diags
}
func resourceAppPolicyIntentCreateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceAppPolicyIntentCreateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntent(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntent {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntent{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_list")))) {
		request.CreateList = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListArray(ctx, key+".create_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".update_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".update_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".update_list")))) {
		request.UpdateList = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListArray(ctx, key+".update_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".delete_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".delete_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".delete_list")))) {
		request.DeleteList = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateList {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateList{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateList {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".delete_policy_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".delete_policy_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".delete_policy_status")))) {
		request.DeletePolicyStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_scope")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_scope")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_scope")))) {
		request.PolicyScope = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope")))) {
		request.AdvancedPolicyScope = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListAdvancedPolicyScope(ctx, key+".advanced_policy_scope.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exclusive_contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exclusive_contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exclusive_contract")))) {
		request.ExclusiveContract = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListExclusiveContract(ctx, key+".exclusive_contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".contract")))) {
		request.Contract = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListContract(ctx, key+".contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".producer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".producer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".producer")))) {
		request.Producer = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListProducer(ctx, key+".producer.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".consumer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".consumer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".consumer")))) {
		request.Consumer = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListConsumer(ctx, key+".consumer.0", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListAdvancedPolicyScope(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScope {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScope{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope_element")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) {
		request.AdvancedPolicyScopeElement = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx, key+".advanced_policy_scope_element", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListExclusiveContract(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContract {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListExclusiveContractClauseArray(ctx, key+".clause", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListExclusiveContractClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListExclusiveContractClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListExclusiveContractClause(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListExclusiveContractClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".relevance_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".relevance_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".relevance_level")))) {
		request.RelevanceLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_removal_behavior")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_removal_behavior")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_removal_behavior")))) {
		request.DeviceRemovalBehavior = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_tracking_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_tracking_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_tracking_enabled")))) {
		request.HostTrackingEnabled = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListContract(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListContract {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListProducer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducer {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListProducerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListProducerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListProducerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListProducerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListProducerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListConsumer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumer {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListConsumerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListConsumerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListConsumerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentCreateListConsumerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentCreateListConsumerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateList {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateList{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateList {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateList{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".delete_policy_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".delete_policy_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".delete_policy_status")))) {
		request.DeletePolicyStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".policy_scope")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".policy_scope")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".policy_scope")))) {
		request.PolicyScope = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".priority")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".priority")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".priority")))) {
		request.Priority = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope")))) {
		request.AdvancedPolicyScope = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListAdvancedPolicyScope(ctx, key+".advanced_policy_scope.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exclusive_contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exclusive_contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exclusive_contract")))) {
		request.ExclusiveContract = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListExclusiveContract(ctx, key+".exclusive_contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".contract")))) {
		request.Contract = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListContract(ctx, key+".contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".producer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".producer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".producer")))) {
		request.Producer = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListProducer(ctx, key+".producer.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".consumer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".consumer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".consumer")))) {
		request.Consumer = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListConsumer(ctx, key+".consumer.0", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListAdvancedPolicyScope(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScope {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScope{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope_element")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) {
		request.AdvancedPolicyScopeElement = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx, key+".advanced_policy_scope_element", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListExclusiveContract(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContract {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListExclusiveContractClauseArray(ctx, key+".clause", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListExclusiveContractClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListExclusiveContractClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListExclusiveContractClause(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListExclusiveContractClause{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".relevance_level")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".relevance_level")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".relevance_level")))) {
		request.RelevanceLevel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_removal_behavior")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_removal_behavior")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_removal_behavior")))) {
		request.DeviceRemovalBehavior = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".host_tracking_enabled")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".host_tracking_enabled")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".host_tracking_enabled")))) {
		request.HostTrackingEnabled = interfaceToBoolPtr(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListContract(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListContract {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListProducer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducer {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListProducerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListProducerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListProducerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListProducerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListProducerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListConsumer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumer {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListConsumerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListConsumerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListConsumerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentUpdateListConsumerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentUpdateListConsumerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func flattenApplicationPolicyApplicationPolicyIntentItem(item *catalystcentersdkgo.ResponseApplicationPolicyApplicationPolicyIntentResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
