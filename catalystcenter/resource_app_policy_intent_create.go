package catalystcenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

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

	request1 := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.ApplicationPolicy.ApplicationPolicyIntentV1(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ApplicationPolicyIntentV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ApplicationPolicyIntentV1", err))
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
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ApplicationPolicyIntentV1", err1))
			return diags
		}
	}

	vItem1 := flattenApplicationPolicyApplicationPolicyIntentV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ApplicationPolicyIntentV1 response",
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

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1 {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_list")))) {
		request.CreateList = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListArray(ctx, key+".create_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".update_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".update_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".update_list")))) {
		request.UpdateList = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListArray(ctx, key+".update_list", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".delete_list")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".delete_list")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".delete_list")))) {
		request.DeleteList = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateList {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateList{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateList {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateList{}
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
		request.AdvancedPolicyScope = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListAdvancedPolicyScope(ctx, key+".advanced_policy_scope.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exclusive_contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exclusive_contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exclusive_contract")))) {
		request.ExclusiveContract = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListExclusiveContract(ctx, key+".exclusive_contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".contract")))) {
		request.Contract = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListContract(ctx, key+".contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".producer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".producer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".producer")))) {
		request.Producer = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListProducer(ctx, key+".producer.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".consumer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".consumer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".consumer")))) {
		request.Consumer = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListConsumer(ctx, key+".consumer.0", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListAdvancedPolicyScope(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScope {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScope{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope_element")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) {
		request.AdvancedPolicyScopeElement = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx, key+".advanced_policy_scope_element", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group_id")))) {
		request.GroupID = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ssid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ssid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ssid")))) {
		request.SSID = interfaceToSliceString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListExclusiveContract(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContract {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListExclusiveContractClauseArray(ctx, key+".clause", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListExclusiveContractClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContractClause {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContractClause{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListExclusiveContractClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListExclusiveContractClause(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContractClause {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListExclusiveContractClause{}
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

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListContract(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListContract {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListProducer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducer {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListProducerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListProducerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducerScalableGroup {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducerScalableGroup{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListProducerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListProducerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducerScalableGroup {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListProducerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListConsumer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumer {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListConsumerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListConsumerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumerScalableGroup {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumerScalableGroup{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListConsumerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1CreateListConsumerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumerScalableGroup {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1CreateListConsumerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateList {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateList{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateList(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateList(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateList {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateList{}
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
		request.AdvancedPolicyScope = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListAdvancedPolicyScope(ctx, key+".advanced_policy_scope.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".exclusive_contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".exclusive_contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".exclusive_contract")))) {
		request.ExclusiveContract = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListExclusiveContract(ctx, key+".exclusive_contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".contract")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".contract")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".contract")))) {
		request.Contract = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListContract(ctx, key+".contract.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".producer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".producer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".producer")))) {
		request.Producer = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListProducer(ctx, key+".producer.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".consumer")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".consumer")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".consumer")))) {
		request.Consumer = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListConsumer(ctx, key+".consumer.0", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListAdvancedPolicyScope(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScope {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScope{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".advanced_policy_scope_element")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".advanced_policy_scope_element")))) {
		request.AdvancedPolicyScopeElement = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx, key+".advanced_policy_scope_element", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElementArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListAdvancedPolicyScopeAdvancedPolicyScopeElement{}
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

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListExclusiveContract(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContract {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".clause")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".clause")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".clause")))) {
		request.Clause = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListExclusiveContractClauseArray(ctx, key+".clause", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListExclusiveContractClauseArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContractClause {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContractClause{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListExclusiveContractClause(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListExclusiveContractClause(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContractClause {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListExclusiveContractClause{}
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

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListContract(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListContract {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListContract{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListProducer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducer {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListProducerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListProducerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducerScalableGroup {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducerScalableGroup{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListProducerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListProducerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducerScalableGroup {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListProducerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListConsumer(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumer {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumer{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group")))) {
		request.ScalableGroup = expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListConsumerScalableGroupArray(ctx, key+".scalable_group", d)
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListConsumerScalableGroupArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumerScalableGroup {
	request := []catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumerScalableGroup{}
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
		i := expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListConsumerScalableGroup(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestAppPolicyIntentCreateApplicationPolicyIntentV1UpdateListConsumerScalableGroup(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumerScalableGroup {
	request := catalystcentersdkgo.RequestApplicationPolicyApplicationPolicyIntentV1UpdateListConsumerScalableGroup{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	return &request
}

func flattenApplicationPolicyApplicationPolicyIntentV1Item(item *catalystcentersdkgo.ResponseApplicationPolicyApplicationPolicyIntentV1Response) []map[string]interface{} {
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
