package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAppPolicy() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get all existing application policies
`,

		ReadContext: dataSourceAppPolicyRead,
		Schema: map[string]*schema.Schema{
			"policy_scope": &schema.Schema{
				Description: `policyScope query parameter. policy scope name
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"advanced_policy_scope": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"advanced_policy_scope_element": &schema.Schema{
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

												"group_id": &schema.Schema{
													Description: `Group id
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"id": &schema.Schema{
													Description: `Id of Advanced policy scope element
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

												"ssid": &schema.Schema{
													Description: `Ssid
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

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of Advanced policy scope
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

									"name": &schema.Schema{
										Description: `Policy name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"cfs_change_info": &schema.Schema{
							Description: `Cfs change info
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"consumer": &schema.Schema{
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
										Description: `Id of Consumer
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

									"scalable_group": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id ref to application Scalable group
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

						"contract_list": &schema.Schema{
							Description: `Contract list
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

						"delete_policy_status": &schema.Schema{
							Description: `NONE: deployed policy to devices, DELETED: delete policy from devices, RESTORED: restored to original configuration
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"deployed": &schema.Schema{
							Description: `Deployed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"display_name": &schema.Schema{
							Description: `Display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"exclusive_contract": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"clause": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_removal_behavior": &schema.Schema{
													Description: `Device removal behavior
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

												"host_tracking_enabled": &schema.Schema{
													Description: `Host tracking enabled
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Id of Business relevance or Application policy knobs clause
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

												"priority": &schema.Schema{
													Description: `Priority
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"relevance_level": &schema.Schema{
													Description: `Relevance level
`,
													Type:     schema.TypeString,
													Computed: true,
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

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of Exclusive contract
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
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id of Group based policy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"identity_source": &schema.Schema{
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
										Description: `Id of Identity source
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

									"state": &schema.Schema{
										Description: `State
`,
										Type:     schema.TypeString,
										Computed: true,
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

						"is_enabled": &schema.Schema{
							Description: `Is enabled
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_scope_stale": &schema.Schema{
							Description: `Is scope stale
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
							Description: `Concatination of <polcy name>_<application-set-name> or <polcy name>_global_policy_configuration or <polcy name>_queuing_customization
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

						"policy_scope": &schema.Schema{
							Description: `Policy name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"policy_status": &schema.Schema{
							Description: `Policy status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"priority": &schema.Schema{
							Description: `Priority
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"producer": &schema.Schema{
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
										Description: `Id of Producer
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

									"scalable_group": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id ref to application-set or application Scalable group
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

						"provisioning_state": &schema.Schema{
							Description: `Provisioning state
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
		},
	}
}

func dataSourceAppPolicyRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vPolicyScope, okPolicyScope := d.GetOk("policy_scope")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplicationPolicyV1")
		queryParams1 := catalystcentersdkgo.GetApplicationPolicyV1QueryParams{}

		if okPolicyScope {
			queryParams1.PolicyScope = vPolicyScope.(string)
		}

		response1, restyResp1, err := client.ApplicationPolicy.GetApplicationPolicyV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApplicationPolicyV1", err,
				"Failure at GetApplicationPolicyV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenApplicationPolicyGetApplicationPolicyV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationPolicyV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetApplicationPolicyV1Items(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_created_on"] = item.InstanceCreatedOn
		respItem["instance_updated_on"] = item.InstanceUpdatedOn
		respItem["instance_version"] = item.InstanceVersion
		respItem["create_time"] = item.CreateTime
		respItem["deployed"] = boolPtrToString(item.Deployed)
		respItem["is_seeded"] = boolPtrToString(item.IsSeeded)
		respItem["is_stale"] = boolPtrToString(item.IsStale)
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["name"] = item.Name
		respItem["namespace"] = item.Namespace
		respItem["provisioning_state"] = item.ProvisioningState
		respItem["qualifier"] = item.Qualifier
		respItem["resource_version"] = item.ResourceVersion
		respItem["target_id_list"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsTargetIDList(item.TargetIDList)
		respItem["type"] = item.Type
		respItem["cfs_change_info"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsCfsChangeInfo(item.CfsChangeInfo)
		respItem["custom_provisions"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsCustomProvisions(item.CustomProvisions)
		respItem["delete_policy_status"] = item.DeletePolicyStatus
		respItem["internal"] = boolPtrToString(item.Internal)
		respItem["is_deleted"] = boolPtrToString(item.IsDeleted)
		respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
		respItem["is_scope_stale"] = boolPtrToString(item.IsScopeStale)
		respItem["ise_reserved"] = boolPtrToString(item.IseReserved)
		respItem["policy_scope"] = item.PolicyScope
		respItem["policy_status"] = item.PolicyStatus
		respItem["priority"] = item.Priority
		respItem["pushed"] = boolPtrToString(item.Pushed)
		respItem["advanced_policy_scope"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsAdvancedPolicyScope(item.AdvancedPolicyScope)
		respItem["contract_list"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsContractList(item.ContractList)
		respItem["exclusive_contract"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsExclusiveContract(item.ExclusiveContract)
		respItem["identity_source"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsIDentitySource(item.IDentitySource)
		respItem["producer"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsProducer(item.Producer)
		respItem["consumer"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsConsumer(item.Consumer)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsTargetIDList(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseTargetIDList) []interface{} {
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

func flattenApplicationPolicyGetApplicationPolicyV1ItemsCfsChangeInfo(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseCfsChangeInfo) []interface{} {
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

func flattenApplicationPolicyGetApplicationPolicyV1ItemsCustomProvisions(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseCustomProvisions) []interface{} {
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

func flattenApplicationPolicyGetApplicationPolicyV1ItemsAdvancedPolicyScope(item *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScope) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["name"] = item.Name
	respItem["advanced_policy_scope_element"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsAdvancedPolicyScopeAdvancedPolicyScopeElement(item.AdvancedPolicyScopeElement)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsAdvancedPolicyScopeAdvancedPolicyScopeElement(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScopeAdvancedPolicyScopeElement) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_created_on"] = item.InstanceCreatedOn
		respItem["instance_updated_on"] = item.InstanceUpdatedOn
		respItem["instance_version"] = item.InstanceVersion
		respItem["group_id"] = item.GroupID
		respItem["ssid"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsAdvancedPolicyScopeAdvancedPolicyScopeElementSSID(item.SSID)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsAdvancedPolicyScopeAdvancedPolicyScopeElementSSID(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseAdvancedPolicyScopeAdvancedPolicyScopeElementSSID) []interface{} {
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

func flattenApplicationPolicyGetApplicationPolicyV1ItemsContractList(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseContractList) []interface{} {
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

func flattenApplicationPolicyGetApplicationPolicyV1ItemsExclusiveContract(item *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseExclusiveContract) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["clause"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsExclusiveContractClause(item.Clause)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsExclusiveContractClause(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseExclusiveContractClause) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_id"] = item.InstanceID
		respItem["display_name"] = item.DisplayName
		respItem["instance_created_on"] = item.InstanceCreatedOn
		respItem["instance_updated_on"] = item.InstanceUpdatedOn
		respItem["instance_version"] = item.InstanceVersion
		respItem["priority"] = item.Priority
		respItem["type"] = item.Type
		respItem["relevance_level"] = item.RelevanceLevel
		respItem["device_removal_behavior"] = item.DeviceRemovalBehavior
		respItem["host_tracking_enabled"] = boolPtrToString(item.HostTrackingEnabled)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsIDentitySource(item *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseIDentitySource) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["state"] = item.State
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsProducer(item *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseProducer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["scalable_group"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsProducerScalableGroup(item.ScalableGroup)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsProducerScalableGroup(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseProducerScalableGroup) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id_ref"] = item.IDRef
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsConsumer(item *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseConsumer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_id"] = item.InstanceID
	respItem["display_name"] = item.DisplayName
	respItem["instance_created_on"] = item.InstanceCreatedOn
	respItem["instance_updated_on"] = item.InstanceUpdatedOn
	respItem["instance_version"] = item.InstanceVersion
	respItem["scalable_group"] = flattenApplicationPolicyGetApplicationPolicyV1ItemsConsumerScalableGroup(item.ScalableGroup)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenApplicationPolicyGetApplicationPolicyV1ItemsConsumerScalableGroup(items *[]catalystcentersdkgo.ResponseApplicationPolicyGetApplicationPolicyV1ResponseConsumerScalableGroup) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id_ref"] = item.IDRef
		respItems = append(respItems, respItem)
	}
	return respItems
}
