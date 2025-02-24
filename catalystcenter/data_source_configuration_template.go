package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfigurationTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- List the templates available

- Details of the template by its id
`,

		ReadContext: dataSourceConfigurationTemplateRead,
		Schema: map[string]*schema.Schema{
			"filter_conflicting_templates": &schema.Schema{
				Description: `filterConflictingTemplates query parameter. Filter template(s) based on confliting templates
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"latest_version": &schema.Schema{
				Description: `latestVersion query parameter. latestVersion flag to get the latest versioned template
`,
				Type:     schema.TypeBool,
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
				Description: `projectId query parameter. Filter template(s) based on project UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_names": &schema.Schema{
				Description: `projectNames query parameter. Filter template(s) based on project names
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
				Description: `sortOrder query parameter. Sort Order Ascending (asc) or Descending (des)
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
			"template_id": &schema.Schema{
				Description: `templateId path parameter. TemplateId(UUID) to get details of the template
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"un_committed": &schema.Schema{
				Description: `unCommitted query parameter. Filter template(s) based on template commited or not
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
					},
				},
			},

			"items": &schema.Schema{
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

						"name": &schema.Schema{
							Description: `Name of template
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"project_name": &schema.Schema{
							Description: `Name of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"template_id": &schema.Schema{
							Description: `UUID of template
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
										Description: `Author of version template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of template
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

									"version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_comment": &schema.Schema{
										Description: `Version comment
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"version_time": &schema.Schema{
										Description: `Template version time
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
		},
	}
}

func dataSourceConfigurationTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vProjectID, okProjectID := d.GetOk("project_id")
	vSoftwareType, okSoftwareType := d.GetOk("software_type")
	vSoftwareVersion, okSoftwareVersion := d.GetOk("software_version")
	vProductFamily, okProductFamily := d.GetOk("product_family")
	vProductSeries, okProductSeries := d.GetOk("product_series")
	vProductType, okProductType := d.GetOk("product_type")
	vFilterConflictingTemplates, okFilterConflictingTemplates := d.GetOk("filter_conflicting_templates")
	vTags, okTags := d.GetOk("tags")
	vProjectNames, okProjectNames := d.GetOk("project_names")
	vUnCommitted, okUnCommitted := d.GetOk("un_committed")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vTemplateID, okTemplateID := d.GetOk("template_id")
	vLatestVersion, okLatestVersion := d.GetOk("latest_version")

	method1 := []bool{okProjectID, okSoftwareType, okSoftwareVersion, okProductFamily, okProductSeries, okProductType, okFilterConflictingTemplates, okTags, okProjectNames, okUnCommitted, okSortOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okTemplateID, okLatestVersion}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsTheTemplatesAvailableV1")
		queryParams1 := catalystcentersdkgo.GetsTheTemplatesAvailableV1QueryParams{}

		if okProjectID {
			queryParams1.ProjectID = vProjectID.(string)
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
		if okProjectNames {
			queryParams1.ProjectNames = interfaceToSliceString(vProjectNames)
		}
		if okUnCommitted {
			queryParams1.UnCommitted = vUnCommitted.(bool)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}

		response1, restyResp1, err := client.ConfigurationTemplates.GetsTheTemplatesAvailableV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsTheTemplatesAvailableV1", err,
				"Failure at GetsTheTemplatesAvailableV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationTemplatesGetsTheTemplatesAvailableV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheTemplatesAvailableV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetsDetailsOfAGivenTemplateV1")
		vvTemplateID := vTemplateID.(string)
		queryParams2 := catalystcentersdkgo.GetsDetailsOfAGivenTemplateV1QueryParams{}

		if okLatestVersion {
			queryParams2.LatestVersion = vLatestVersion.(bool)
		}

		response2, restyResp2, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplateV1(vvTemplateID, &queryParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsDetailsOfAGivenTemplateV1", err,
				"Failure at GetsDetailsOfAGivenTemplateV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1Item(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsDetailsOfAGivenTemplateV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetsTheTemplatesAvailableV1Items(items *catalystcentersdkgo.ResponseConfigurationTemplatesGetsTheTemplatesAvailableV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["name"] = item.Name
		respItem["project_id"] = item.ProjectID
		respItem["project_name"] = item.ProjectName
		respItem["template_id"] = item.TemplateID
		respItem["versions_info"] = flattenConfigurationTemplatesGetsTheTemplatesAvailableV1ItemsVersionsInfo(item.VersionsInfo)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsTheTemplatesAvailableV1ItemsVersionsInfo(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsTheTemplatesAvailableV1VersionsInfo) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1Item(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tags"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTags(item.Tags)
	respItem["author"] = item.Author
	respItem["composite"] = boolPtrToString(item.Composite)
	respItem["containing_templates"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplates(item.ContainingTemplates)
	respItem["create_time"] = item.CreateTime
	respItem["custom_params_order"] = boolPtrToString(item.CustomParamsOrder)
	respItem["description"] = item.Description
	respItem["device_types"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemDeviceTypes(item.DeviceTypes)
	respItem["failure_policy"] = item.FailurePolicy
	respItem["id"] = item.ID
	respItem["language"] = item.Language
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["latest_version_time"] = item.LatestVersionTime
	respItem["name"] = item.Name
	respItem["parent_template_id"] = item.ParentTemplateID
	respItem["project_id"] = item.ProjectID
	respItem["project_name"] = item.ProjectName
	respItem["rollback_template_content"] = item.RollbackTemplateContent
	respItem["rollback_template_params"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemRollbackTemplateParams(item.RollbackTemplateParams)
	respItem["software_type"] = item.SoftwareType
	respItem["software_variant"] = item.SoftwareVariant
	respItem["software_version"] = item.SoftwareVersion
	respItem["template_content"] = item.TemplateContent
	respItem["template_params"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTemplateParams(item.TemplateParams)
	respItem["validation_errors"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemValidationErrors(item.ValidationErrors)
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTags(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1Tags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplates(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTags(item.Tags)
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesDeviceTypes(item.DeviceTypes)
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["name"] = item.Name
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTemplateParams(item.TemplateParams)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTags(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesDeviceTypes(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesDeviceTypes) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesRollbackTemplateParams(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesRollbackTemplateParamsRange(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesRollbackTemplateParamsSelection(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTemplateParams(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTemplateParamsRange(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTemplateParamsSelection(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemContainingTemplatesTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ContainingTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemDeviceTypes(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1DeviceTypes) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemRollbackTemplateParams(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemRollbackTemplateParamsRange(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemRollbackTemplateParamsSelection(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemRollbackTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1RollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTemplateParams(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTemplateParamsRange(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTemplateParamsSelection(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1TemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemValidationErrors(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrors) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rollback_template_errors"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemValidationErrorsRollbackTemplateErrors(item.RollbackTemplateErrors)
	respItem["template_errors"] = flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemValidationErrorsTemplateErrors(item.TemplateErrors)
	respItem["template_id"] = item.TemplateID
	respItem["template_version"] = item.TemplateVersion

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemValidationErrorsRollbackTemplateErrors(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrorsRollbackTemplateErrors) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ItemValidationErrorsTemplateErrors(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplateV1ValidationErrorsTemplateErrors) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
