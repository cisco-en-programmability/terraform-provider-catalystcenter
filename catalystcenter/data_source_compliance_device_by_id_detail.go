package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDeviceByIDDetail() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return compliance detailed report for a device.
`,

		ReadContext: dataSourceComplianceDeviceByIDDetailRead,
		Schema: map[string]*schema.Schema{
			"category": &schema.Schema{
				Description: `category query parameter. category can have any value among 'INTENT', 'RUNNING_CONFIG' , 'IMAGE' , 'PSIRT' , 'DESIGN_OOD' , 'EoX' , 'NETWORK_SETTINGS'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_type": &schema.Schema{
				Description: `complianceType query parameter. Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EoX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"device_uuid": &schema.Schema{
				Description: `deviceUuid path parameter. Device Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"diff_list": &schema.Schema{
				Description: `diffList query parameter. diff list [ pass true to fetch the diff list ]
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"remediation_supported": &schema.Schema{
				Description: `remediationSupported query parameter. The 'remediationSupported' parameter can be set to 'true' or 'false'. The result will be a combination of both values if it is not provided.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. 'COMPLIANT', 'NON_COMPLIANT', 'ERROR', 'IN_PROGRESS', 'NOT_APPLICABLE', 'NOT_AVAILABLE', 'WARNING', 'REMEDIATION_IN_PROGRESS' can be the value of the compliance 'status' parameter. [COMPLIANT: Device currently meets the compliance requirements.  NON_COMPLIANT: One of the compliance requirements like Software Image, PSIRT, Network Profile, Startup vs Running, etc. are not met. ERROR: Compliance is unable to compute status due to underlying errors. IN_PROGRESS: Compliance check is in progress for the device. NOT_APPLICABLE: Device is not supported for compliance, or minimum license requirement is not met. NOT_AVAILABLE: Compliance is not available for the device. COMPLIANT_WARNING: The device is compliant with warning if the last date of support is nearing. REMEDIATION_IN_PROGRESS: Compliance remediation is in progress for the device.]
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ack_status": &schema.Schema{
							Description: `Acknowledgment status of the compliance type. UNACKNOWLEDGED if none of the violations under the compliance type are acknowledged. Otherwise it will be ACKNOWLEDGED.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"compliance_type": &schema.Schema{
							Description: `Compliance type corresponds to a tile on the UI that will be one of NETWORK_PROFILE, IMAGE, APPLICATION_VISIBILITY, FABRIC, PSIRT, RUNNING_CONFIG, NETWORK_SETTINGS, WORKFLOW, or EoX.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_uuid": &schema.Schema{
							Description: `UUID of the device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_sync_time": &schema.Schema{
							Description: `Timestamp when the status changed from a different value to the current value.
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Timestamp of the latest compliance check that was run.
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"remediation_supported": &schema.Schema{
							Description: `Indicates whether remediation is supported for this compliance type or not.
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"source_info_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ack_status": &schema.Schema{
										Description: `Acknowledgment status of violations. UNACKNOWLEDGED if none of the violations are acknowledged. Otherwise it will be ACKNOWLEDGED.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"app_name": &schema.Schema{
										Description: `Application name that is used to club the violations.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"business_key": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"business_key_attributes": &schema.Schema{
													Description: `Attributes that together uniquely identify the configuration instance.
`,
													Type:     schema.TypeString, //TEST,
													Computed: true,
												},

												"other_attributes": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"cfs_attributes": &schema.Schema{
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{

																		"app_name": &schema.Schema{
																			Description: `Same as appName above.
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"description": &schema.Schema{
																			Description: `Description for the configuration, if available.
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"display_name": &schema.Schema{
																			Description: `User friendly name for the configuration.
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"source": &schema.Schema{
																			Description: `Will be same as compliance type.
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},

																		"type": &schema.Schema{
																			Description: `The type of this attribute (for example, type can be Intent).
`,
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																	},
																},
															},

															"name": &schema.Schema{
																Description: `Name of the attributes.
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"resource_name": &schema.Schema{
													Description: `Name of the top level resource. Same as name above.
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"count": &schema.Schema{
										Description: `Number of violations present.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"diff_list": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ack_status": &schema.Schema{
													Description: `Acknowledgment status of the violation. ACKNOWLEDGED if the violation is acknowledged or at the top-level configuration. Otherwise it will be UNACKNOWLEDGED.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"business_key": &schema.Schema{
													Description: `The Unique key of the individual violation does not change after every compliance check, as long as the deployment data doesn't change.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"configured_value": &schema.Schema{
													Description: `Configured value i.e. running / current value. It will be empty for the template violations due to potentially large size of the template. Use a dedicated API to get the template data.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"display_name": &schema.Schema{
													Description: `Display name for attribute in ui .If business key is null or of type owning entity type.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"extended_attributes": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"attribute_display_name": &schema.Schema{
																Description: `Display name for attribute in ui .if business key is null or only owning entity type.
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"data_converter": &schema.Schema{
																Description: `Name of the converter used to display configurations in user-friendly format, if available.
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"path": &schema.Schema{
																Description: `Path to be displayed on the UI, instead of the above path, if available.
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"type": &schema.Schema{
																Description: `Type of this attribute.(example type can be Intent)
`,
																Type:     schema.TypeString,
																Computed: true,
															},
														},
													},
												},

												"instance_uui_d": &schema.Schema{
													Description: `UUID of the individual violation. Changes after every compliance check.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"intended_value": &schema.Schema{
													Description: `Enable", Intended value. It will be empty for the template violations due to potentially large size of the template. Use a dedicated API to get the template data.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"move_from_path": &schema.Schema{
													Description: `Additional URI to fetch more details, if available.
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"op": &schema.Schema{
													Description: `Type of change (add, remove, or update).
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"path": &schema.Schema{
													Description: `Path of the configuration relative to the top-level configuration. Use it along with a name to identify certain set of violations.
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"display_name": &schema.Schema{
										Description: `Model display name.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the type of top level configuration.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name_with_business_key": &schema.Schema{
										Description: `Name With Business Key`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"source_enum": &schema.Schema{
										Description: `Will be same as compliance type.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `Type of the top level configuration.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"state": &schema.Schema{
							Description: `State of the compliance check for the compliance type, will be one of SUCCESS, FAILED, or IN_PROGRESS.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status of compliance for the compliance type, will be one of COMPLIANT, NON_COMPLIANT, ERROR, IN_PROGRESS, NOT_APPLICABLE, NOT_AVAILABLE, COMPLIANT_WARNING, REMEDIATION_IN_PROGRESS, or ABORTED.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Version of the API.
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

func dataSourceComplianceDeviceByIDDetailRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceUUID := d.Get("device_uuid")
	vCategory, okCategory := d.GetOk("category")
	vComplianceType, okComplianceType := d.GetOk("compliance_type")
	vDiffList, okDiffList := d.GetOk("diff_list")
	vStatus, okStatus := d.GetOk("status")
	vRemediationSupported, okRemediationSupported := d.GetOk("remediation_supported")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ComplianceDetailsOfDeviceV1")
		vvDeviceUUID := vDeviceUUID.(string)
		queryParams1 := catalystcentersdkgo.ComplianceDetailsOfDeviceV1QueryParams{}

		if okCategory {
			queryParams1.Category = vCategory.(string)
		}
		if okComplianceType {
			queryParams1.ComplianceType = vComplianceType.(string)
		}
		if okDiffList {
			queryParams1.DiffList = vDiffList.(bool)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okRemediationSupported {
			queryParams1.RemediationSupported = vRemediationSupported.(bool)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Compliance.ComplianceDetailsOfDeviceV1(vvDeviceUUID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ComplianceDetailsOfDeviceV1", err,
				"Failure at ComplianceDetailsOfDeviceV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenComplianceComplianceDetailsOfDeviceV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ComplianceDetailsOfDeviceV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceComplianceDetailsOfDeviceV1Items(items *[]catalystcentersdkgo.ResponseComplianceComplianceDetailsOfDeviceV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_uuid"] = item.DeviceUUID
		respItem["compliance_type"] = item.ComplianceType
		respItem["status"] = item.Status
		respItem["state"] = item.State
		respItem["last_sync_time"] = item.LastSyncTime
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["source_info_list"] = flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoList(item.SourceInfoList)
		respItem["ack_status"] = item.AckStatus
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoList(items *[]catalystcentersdkgo.ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["name_with_business_key"] = item.NameWithBusinessKey
		respItem["source_enum"] = item.SourceEnum
		respItem["type"] = item.Type
		respItem["app_name"] = item.AppName
		respItem["count"] = item.Count
		respItem["ack_status"] = item.AckStatus
		respItem["business_key"] = flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListBusinessKey(item.BusinessKey)
		respItem["diff_list"] = flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListDiffList(item.DiffList)
		respItem["display_name"] = item.DisplayName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListBusinessKey(item *catalystcentersdkgo.ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKey) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["resource_name"] = item.ResourceName
	respItem["business_key_attributes"] = flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListBusinessKeyBusinessKeyAttributes(item.BusinessKeyAttributes)
	respItem["other_attributes"] = flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListBusinessKeyOtherAttributes(item.OtherAttributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListBusinessKeyBusinessKeyAttributes(item *catalystcentersdkgo.ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyBusinessKeyAttributes) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListBusinessKeyOtherAttributes(item *catalystcentersdkgo.ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyOtherAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["cfs_attributes"] = flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListBusinessKeyOtherAttributesCfsAttributes(item.CfsAttributes)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListBusinessKeyOtherAttributesCfsAttributes(item *catalystcentersdkgo.ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListBusinessKeyOtherAttributesCfsAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["display_name"] = item.DisplayName
	respItem["app_name"] = item.AppName
	respItem["description"] = item.Description
	respItem["source"] = item.Source
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListDiffList(items *[]catalystcentersdkgo.ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListDiffList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["op"] = item.Op
		respItem["configured_value"] = item.ConfiguredValue
		respItem["intended_value"] = item.IntendedValue
		respItem["move_from_path"] = item.MoveFromPath
		respItem["business_key"] = item.BusinessKey
		respItem["path"] = item.Path
		respItem["extended_attributes"] = flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListDiffListExtendedAttributes(item.ExtendedAttributes)
		respItem["ack_status"] = item.AckStatus
		respItem["instance_uui_d"] = item.InstanceUUID
		respItem["display_name"] = item.DisplayName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenComplianceComplianceDetailsOfDeviceV1ItemsSourceInfoListDiffListExtendedAttributes(item *catalystcentersdkgo.ResponseComplianceComplianceDetailsOfDeviceV1ResponseSourceInfoListDiffListExtendedAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attribute_display_name"] = item.AttributeDisplayName
	respItem["path"] = item.Path
	respItem["data_converter"] = item.DataConverter
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
