package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplatesDetailsV2() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get template(s) details
`,

		ReadContext: dataSourceTemplatesDetailsV2Read,
		Schema: map[string]*schema.Schema{
			"all_template_attributes": &schema.Schema{
				Description: `allTemplateAttributes query parameter. Return all template attributes
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"filter_conflicting_templates": &schema.Schema{
				Description: `filterConflictingTemplates query parameter. Filter template(s) based on confliting templates
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Id of template to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"include_version_details": &schema.Schema{
				Description: `includeVersionDetails query parameter. Include template version details
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page;The minimum is 1, and the maximum is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Name of template to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Index of first result
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"product_family": &schema.Schema{
				Description: `productFamily query parameter. Filter template(s) based on device family
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_series": &schema.Schema{
				Description: `productSeries query parameter. Filter template(s) based on device series
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_type": &schema.Schema{
				Description: `productType query parameter. Filter template(s) based on device type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": &schema.Schema{
				Description: `projectId query parameter. Filter template(s) based on project id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_name": &schema.Schema{
				Description: `projectName query parameter. Filter template(s) based on project name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_type": &schema.Schema{
				Description: `softwareType query parameter. Filter template(s) based software type
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"software_version": &schema.Schema{
				Description: `softwareVersion query parameter. Filter template(s) based softwareVersion
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_order": &schema.Schema{
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (dsc)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Description: `tags query parameter. Filter template(s) based on tags
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"un_committed": &schema.Schema{
				Description: `unCommitted query parameter. Return uncommitted template
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"author": &schema.Schema{
							Description: `Author of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"composite": &schema.Schema{
							Description: `Is it composite template
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"containing_templates": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"composite": &schema.Schema{
										Description: `Is it composite template
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"device_types": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Device family
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_series": &schema.Schema{
													Description: `Device series
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_type": &schema.Schema{
													Description: `Device type
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"language": &schema.Schema{
										Description: `Template language (JINJA or VELOCITY)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"project_name": &schema.Schema{
										Description: `Project name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rollback_template_params": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"range": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"tags": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of tag
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Name of tag
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"template_content": &schema.Schema{
										Description: `Template content
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"template_params": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"range": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},

															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Computed: true,
															},

															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeString, //TEST,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"create_time": &schema.Schema{
							Description: `Create time of template
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"custom_params_order": &schema.Schema{
							Description: `Custom Params Order
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_types": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"product_family": &schema.Schema{
										Description: `Device family
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"product_series": &schema.Schema{
										Description: `Device series
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"product_type": &schema.Schema{
										Description: `Device type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"failure_policy": &schema.Schema{
							Description: `Define failure policy if template provisioning fails
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"language": &schema.Schema{
							Description: `Template language (JINJA or VELOCITY)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Update time of template
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"latest_version_time": &schema.Schema{
							Description: `Latest versioned template time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_template_id": &schema.Schema{
							Description: `Parent templateID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_associated": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_id": &schema.Schema{
							Description: `Project UUID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_name": &schema.Schema{
							Description: `Project name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rollback_template_content": &schema.Schema{
							Description: `Rollback template content
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rollback_template_params": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"range": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"required": &schema.Schema{
										Description: `Is param required
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"selection_values": &schema.Schema{
													Description: `Selection values
`,
													Type:     schema.TypeString, //TEST,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"software_type": &schema.Schema{
							Description: `Applicable device software type
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_variant": &schema.Schema{
							Description: `Applicable device software variant
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"software_version": &schema.Schema{
							Description: `Applicable device software version
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"template_content": &schema.Schema{
							Description: `Template content
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"template_params": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"range": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"required": &schema.Schema{
										Description: `Is param required
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"selection_values": &schema.Schema{
													Description: `Selection values
`,
													Type:     schema.TypeString, //TEST,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"validation_errors": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rollback_template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors of rollback template
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"template_id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"template_version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `Current version of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"versions_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"version": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_comment": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_time": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
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

func dataSourceTemplatesDetailsV2Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vName, okName := d.GetOk("name")
	vProjectID, okProjectID := d.GetOk("project_id")
	vProjectName, okProjectName := d.GetOk("project_name")
	vSoftwareType, okSoftwareType := d.GetOk("software_type")
	vSoftwareVersion, okSoftwareVersion := d.GetOk("software_version")
	vProductFamily, okProductFamily := d.GetOk("product_family")
	vProductSeries, okProductSeries := d.GetOk("product_series")
	vProductType, okProductType := d.GetOk("product_type")
	vFilterConflictingTemplates, okFilterConflictingTemplates := d.GetOk("filter_conflicting_templates")
	vTags, okTags := d.GetOk("tags")
	vUnCommitted, okUnCommitted := d.GetOk("un_committed")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vAllTemplateAttributes, okAllTemplateAttributes := d.GetOk("all_template_attributes")
	vIncludeVersionDetails, okIncludeVersionDetails := d.GetOk("include_version_details")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTemplatesDetailsV2")
		queryParams1 := catalystcentersdkgo.GetTemplatesDetailsV2QueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okProjectID {
			queryParams1.ProjectID = vProjectID.(string)
		}
		if okProjectName {
			queryParams1.ProjectName = vProjectName.(string)
		}
		if okSoftwareType {
			queryParams1.SoftwareType = vSoftwareType.(string)
		}
		if okSoftwareVersion {
			queryParams1.SoftwareVersion = vSoftwareVersion.(string)
		}
		if okProductFamily {
			queryParams1.ProductFamily = vProductFamily.(string)
		}
		if okProductSeries {
			queryParams1.ProductSeries = vProductSeries.(string)
		}
		if okProductType {
			queryParams1.ProductType = vProductType.(string)
		}
		if okFilterConflictingTemplates {
			queryParams1.FilterConflictingTemplates = vFilterConflictingTemplates.(bool)
		}
		if okTags {
			queryParams1.Tags = interfaceToSliceString(vTags)
		}
		if okUnCommitted {
			queryParams1.UnCommitted = vUnCommitted.(bool)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}
		if okAllTemplateAttributes {
			queryParams1.AllTemplateAttributes = vAllTemplateAttributes.(bool)
		}
		if okIncludeVersionDetails {
			queryParams1.IncludeVersionDetails = vIncludeVersionDetails.(bool)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.ConfigurationTemplates.GetTemplatesDetailsV2(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTemplatesDetailsV2", err,
				"Failure at GetTemplatesDetailsV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTemplatesDetailsV2", err,
				"Failure at GetTemplatesDetailsV2, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesGetTemplatesDetailsV2Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTemplatesDetailsV2 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2Item(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["author"] = item.Author
	respItem["composite"] = boolPtrToString(item.Composite)
	respItem["containing_templates"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplates(item.ContainingTemplates)
	respItem["create_time"] = item.CreateTime
	respItem["custom_params_order"] = boolPtrToString(item.CustomParamsOrder)
	respItem["description"] = item.Description
	respItem["device_types"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemDeviceTypes(item.DeviceTypes)
	respItem["failure_policy"] = item.FailurePolicy
	respItem["id"] = item.ID
	respItem["language"] = item.Language
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["latest_version_time"] = item.LatestVersionTime
	respItem["name"] = item.Name
	respItem["parent_template_id"] = item.ParentTemplateID
	respItem["project_associated"] = boolPtrToString(item.ProjectAssociated)
	respItem["project_id"] = item.ProjectID
	respItem["project_name"] = item.ProjectName
	respItem["rollback_template_content"] = item.RollbackTemplateContent
	respItem["rollback_template_params"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemRollbackTemplateParams(item.RollbackTemplateParams)
	respItem["software_type"] = item.SoftwareType
	respItem["software_variant"] = item.SoftwareVariant
	respItem["software_version"] = item.SoftwareVersion
	respItem["tags"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTags(item.Tags)
	respItem["template_content"] = item.TemplateContent
	respItem["template_params"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTemplateParams(item.TemplateParams)
	respItem["validation_errors"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemValidationErrors(item.ValidationErrors)
	respItem["version"] = item.Version
	respItem["versions_info"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemVersionsInfo(item.VersionsInfo)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplates(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesDeviceTypes(item.DeviceTypes)
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["name"] = item.Name
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["tags"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTags(item.Tags)
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTemplateParams(item.TemplateParams)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesDeviceTypes(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesDeviceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_type"] = item.ProductType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesRollbackTemplateParams(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesRollbackTemplateParamsRange(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesRollbackTemplateParamsSelection(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTags(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTemplateParams(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTemplateParamsRange(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTemplateParamsSelection(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemContainingTemplatesTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ContainingTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemDeviceTypes(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2DeviceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_type"] = item.ProductType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemRollbackTemplateParams(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemRollbackTemplateParamsRange(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemRollbackTemplateParamsSelection(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemRollbackTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2RollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTags(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2Tags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTemplateParams(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParams) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["binding"] = item.Binding
		respItem["custom_order"] = item.CustomOrder
		respItem["data_type"] = item.DataType
		respItem["default_value"] = item.DefaultValue
		respItem["description"] = item.Description
		respItem["display_name"] = item.DisplayName
		respItem["group"] = item.Group
		respItem["id"] = item.ID
		respItem["instruction_text"] = item.InstructionText
		respItem["key"] = item.Key
		respItem["not_param"] = boolPtrToString(item.NotParam)
		respItem["order"] = item.Order
		respItem["param_array"] = boolPtrToString(item.ParamArray)
		respItem["parameter_name"] = item.ParameterName
		respItem["provider"] = item.Provider
		respItem["range"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTemplateParamsRange(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsRange) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["max_value"] = item.MaxValue
		respItem["min_value"] = item.MinValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTemplateParamsSelection(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2TemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemValidationErrors(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrors) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rollback_template_errors"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemValidationErrorsRollbackTemplateErrors(item.RollbackTemplateErrors)
	respItem["template_errors"] = flattenConfigurationTemplatesGetTemplatesDetailsV2ItemValidationErrorsTemplateErrors(item.TemplateErrors)
	respItem["template_id"] = item.TemplateID
	respItem["template_version"] = item.TemplateVersion

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemValidationErrorsRollbackTemplateErrors(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsRollbackTemplateErrors) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemValidationErrorsTemplateErrors(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2ValidationErrorsTemplateErrors) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetTemplatesDetailsV2ItemVersionsInfo(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplatesDetailsV2VersionsInfo) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["author"] = item.Author
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["version"] = item.Version
		respItem["version_comment"] = item.VersionComment
		respItem["version_time"] = item.VersionTime
		respItems = append(respItems, respItem)
	}
	return respItems
}
