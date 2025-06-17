package catalystcenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfigurationTemplate() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Configuration Templates.

- API to create a template by project id.

- API to update a template.

- Deletes the template by its id
`,

		CreateContext: resourceConfigurationTemplateCreate,
		ReadContext:   resourceConfigurationTemplateRead,
		UpdateContext: resourceConfigurationTemplateUpdate,
		DeleteContext: resourceConfigurationTemplateDelete,
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
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"author": &schema.Schema{
							Description: `Author of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"composite": &schema.Schema{
							Description: `Is it composite template
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"containing_templates": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"composite": &schema.Schema{
										Description: `Is it composite template
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"description": &schema.Schema{
										Description: `Description of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"device_types": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Device family
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"product_series": &schema.Schema{
													Description: `Device series
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"product_type": &schema.Schema{
													Description: `Device type
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"language": &schema.Schema{
										Description: `Template language (JINJA or VELOCITY)
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"project_name": &schema.Schema{
										Description: `Project name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"rollback_template_params": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeString, //TEST,
																Optional: true,
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
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of tag
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Description: `Name of tag
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"template_content": &schema.Schema{
										Description: `Template content
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"template_params": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"binding": &schema.Schema{
													Description: `Bind to source
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"custom_order": &schema.Schema{
													Description: `CustomOrder of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"data_type": &schema.Schema{
													Description: `Datatype of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"default_value": &schema.Schema{
													Description: `Default value of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"description": &schema.Schema{
													Description: `Description of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"display_name": &schema.Schema{
													Description: `Display name of param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"group": &schema.Schema{
													Description: `group
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"id": &schema.Schema{
													Description: `UUID of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"instruction_text": &schema.Schema{
													Description: `Instruction text for param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"key": &schema.Schema{
													Description: `key
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"not_param": &schema.Schema{
													Description: `Is it not a variable
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"order": &schema.Schema{
													Description: `Order of template param
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"param_array": &schema.Schema{
													Description: `Is it an array
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"parameter_name": &schema.Schema{
													Description: `Name of template param
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"provider": &schema.Schema{
													Description: `provider
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"range": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"id": &schema.Schema{
																Description: `UUID of range
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"max_value": &schema.Schema{
																Description: `Max value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
															"min_value": &schema.Schema{
																Description: `Min value of range
`,
																Type:     schema.TypeInt,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
												"required": &schema.Schema{
													Description: `Is param required
`,
													// Type:        schema.TypeBool,
													Type:         schema.TypeString,
													ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
													Optional:     true,
													Computed:     true,
												},
												"selection": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"default_selected_values": &schema.Schema{
																Description: `Default selection values
`,
																Type:     schema.TypeList,
																Optional: true,
																Computed: true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
															"id": &schema.Schema{
																Description: `UUID of selection
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"selection_type": &schema.Schema{
																Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
															"selection_values": &schema.Schema{
																Description: `Selection values
`,
																Type:     schema.TypeString, //TEST,
																Optional: true,
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
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"create_time": &schema.Schema{
							Description: `Create time of template
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"custom_params_order": &schema.Schema{
							Description: `Custom Params Order
`,
							// Type:        schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
							Computed:     true,
						},
						"description": &schema.Schema{
							Description: `Description of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"device_types": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"product_family": &schema.Schema{
										Description: `Device family
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"product_series": &schema.Schema{
										Description: `Device series
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"product_type": &schema.Schema{
										Description: `Device type
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"failure_policy": &schema.Schema{
							Description: `Define failure policy if template provisioning fails
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `UUID of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"language": &schema.Schema{
							Description: `Template language (JINJA or VELOCITY)
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Update time of template
`,
							Type:             schema.TypeInt,
							Optional:         true,
							Computed:         true,
							DiffSuppressFunc: diffSuppressAlways(),
						},
						"latest_version_time": &schema.Schema{
							Description: `Latest versioned template time
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parent_template_id": &schema.Schema{
							Description: `Parent templateID
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"project_id": &schema.Schema{
							Description: `Project UUID
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"project_name": &schema.Schema{
							Description: `Project name
`,
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "",
							DiffSuppressFunc: diffSuppressAlways(),
						},
						"rollback_template_content": &schema.Schema{
							Description: `Rollback template content
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rollback_template_params": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"required": &schema.Schema{
										Description: `Is param required
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"selection_values": &schema.Schema{
													Description: `Selection values
`,
													Type:     schema.TypeString, //TEST,
													Optional: true,
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
							Optional: true,
							Computed: true,
						},
						"software_variant": &schema.Schema{
							Description: `Applicable device software variant
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"software_version": &schema.Schema{
							Description: `Applicable device software version
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `UUID of tag
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Name of tag
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"template_content": &schema.Schema{
							Description: `Template content
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"template_id": &schema.Schema{
							Description: `templateId path parameter. templateId(UUID) of template to be deleted
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"template_params": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"binding": &schema.Schema{
										Description: `Bind to source
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"custom_order": &schema.Schema{
										Description: `CustomOrder of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"data_type": &schema.Schema{
										Description: `Datatype of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"default_value": &schema.Schema{
										Description: `Default value of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"display_name": &schema.Schema{
										Description: `Display name of param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"group": &schema.Schema{
										Description: `group
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `UUID of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"instruction_text": &schema.Schema{
										Description: `Instruction text for param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"key": &schema.Schema{
										Description: `key
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"not_param": &schema.Schema{
										Description: `Is it not a variable
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"order": &schema.Schema{
										Description: `Order of template param
`,
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"param_array": &schema.Schema{
										Description: `Is it an array
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"parameter_name": &schema.Schema{
										Description: `Name of template param
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"provider": &schema.Schema{
										Description: `provider
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"range": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `UUID of range
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"max_value": &schema.Schema{
													Description: `Max value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"min_value": &schema.Schema{
													Description: `Min value of range
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
											},
										},
									},
									"required": &schema.Schema{
										Description: `Is param required
`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										Computed:     true,
									},
									"selection": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"default_selected_values": &schema.Schema{
													Description: `Default selection values
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"id": &schema.Schema{
													Description: `UUID of selection
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"selection_type": &schema.Schema{
													Description: `Type of selection(SINGLE_SELECT or MULTI_SELECT)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"selection_values": &schema.Schema{
													Description: `Selection values
`,
													Type:     schema.TypeString, //TEST,
													Optional: true,
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
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"rollback_template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors of rollback template
`,
										Type:     schema.TypeString, //TEST,
										Optional: true,
										Computed: true,
									},
									"template_errors": &schema.Schema{
										Description: `Validation or design conflicts errors
`,
										Type:     schema.TypeString, //TEST,
										Optional: true,
										Computed: true,
									},
									"template_id": &schema.Schema{
										Description: `UUID of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"template_version": &schema.Schema{
										Description: `Current version of template
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"version": &schema.Schema{
							Description: `Current version of template
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceConfigurationTemplateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestConfigurationTemplateCreateTemplate(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vTemplateID := resourceItem["template_id"]
	vvTemplateID := interfaceToString(vTemplateID)
	vProjectName := resourceItem["project_name"]
	vvProjectName := interfaceToString(vProjectName)
	vProjectID := resourceItem["project_id"]
	vvProjectID := interfaceToString(vProjectID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if vvProjectName == "" {
		project, restyResp1, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)
		if err != nil {
			if restyResp1 != nil {
				diags = append(diags, diagErrorWithResponse(
					"Failure when executing CreateTemplate", err, restyResp1.String()))
				return diags
			}
			diags = append(diags, diagError(
				"Failure when executing CreateTemplate", err))
			return diags
		}
		vvProjectName = project.Name
	}
	if vvTemplateID != "" {
		getResponse2, _, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplate(vvTemplateID, nil)
		// log.Printf("[DEBUG] getResponse2 sent => %v", responseInterfaceToString(*getResponse2))
		if err == nil && getResponse2 != nil {
			resourceMap := make(map[string]string)
			resourceMap["template_id"] = vvTemplateID
			resourceMap["project_name"] = getResponse2.ProjectName
			resourceMap["project_id"] = getResponse2.ProjectID
			resourceMap["template_name"] = getResponse2.Name
			d.SetId(joinResourceID(resourceMap))
			return resourceConfigurationTemplateRead(ctx, d, m)
		}
	} else {
		queryParamImport := catalystcentersdkgo.GetsAListOfProjectsQueryParams{}
		queryParamImport.Name = vvProjectName
		log.Print("[DEBUG] 1")
		response2, err := searchConfigurationTemplatesGetsTheTemplatesAvailable(m, queryParamImport, vvName)
		if response2 != nil {
			log.Printf("[DEBUG] response2 sent => %v", responseInterfaceToString(*response2))
		}
		if response2 != nil && err == nil {
			log.Print("[DEBUG] 2")
			resourceMap := make(map[string]string)
			resourceMap["template_name"] = response2.Name
			resourceMap["project_name"] = response2.ProjectName
			resourceMap["template_id"] = response2.ID
			resourceMap["project_id"] = response2.ProjectID
			d.SetId(joinResourceID(resourceMap))
			return resourceConfigurationTemplateRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.ConfigurationTemplates.CreateTemplate(vvProjectID, request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateTemplate", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateTemplate", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateTemplate", err))
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
				"Failure when executing CreateTemplate", err1))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetsAListOfProjectsQueryParams{}
	queryParamValidate.Name = vvProjectName
	item3, err := searchConfigurationTemplatesGetsTheTemplatesAvailable(m, queryParamValidate, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateTemplate", err,
			"Failure at CreateTemplate, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["template_name"] = item3.Name
	resourceMap["project_name"] = item3.ProjectName
	resourceMap["template_id"] = item3.ID
	resourceMap["project_id"] = item3.ProjectID
	d.SetId(joinResourceID(resourceMap))
	return resourceConfigurationTemplateRead(ctx, d, m)
}

func resourceConfigurationTemplateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vProjectName := resourceMap["project_name"]
	// vTemplateName := resourceMap["template_name"]
	vTemplateID := resourceMap["template_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsDetailsOfAGivenTemplate")
		queryParams1 := catalystcentersdkgo.GetsAListOfProjectsQueryParams{}
		queryParams1.Name = vProjectName

		// response1, err := searchConfigurationTemplatesGetsTheTemplatesAvailable(m, queryParams1, vTemplateName)
		response1, restyResp1, err := client.ConfigurationTemplates.GetsDetailsOfAGivenTemplate(vTemplateID, nil)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("resty response= %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}
		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		// Review flatten function used
		// items := catalystcentersdkgo.ResponseConfigurationTemplatesGetsDetailsOfAGivenTemplate{
		// 	*response1,

		vItem1 := flattenConfigurationTemplatesGetsDetailsOfAGivenTemplateItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheTemplatesAvailable search response",
				err))
			return diags
		}

		if err := d.Set("parameters", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheTemplatesAvailable search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceConfigurationTemplateUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["template_id"]
	if d.HasChange("parameters") {
		request1 := expandRequestConfigurationTemplateUpdateTemplate(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && request1.ID == "" {
			request1.ID = vID
		}
		response1, restyResp1, err := client.ConfigurationTemplates.UpdateTemplate(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateTemplate", err, restyResp1.String(),
					"Failure at UpdateTemplate, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateTemplate", err,
				"Failure at UpdateTemplate, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateTemplate", err))
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
					"Failure when executing UpdateTemplate", err1))
				return diags
			}
		}

	}

	return resourceConfigurationTemplateRead(ctx, d, m)
}

func resourceConfigurationTemplateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvTemplateID := resourceMap["template_id"]

	response1, restyResp1, err := client.ConfigurationTemplates.DeletesTheTemplate(vvTemplateID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesTheTemplate", err, restyResp1.String(),
				"Failure at DeletesTheTemplate, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesTheTemplate", err,
			"Failure at DeletesTheTemplate, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletesTheTemplate", err))
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
				"Failure when executing DeletesTheTemplate", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestConfigurationTemplateCreateTemplate(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplate {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateCreateTemplateTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".author")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".author")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".author")))) {
		request.Author = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".containing_templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".containing_templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".containing_templates")))) {
		request.ContainingTemplates = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesArray(ctx, key+".containing_templates", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_params_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_params_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_params_order")))) {
		request.CustomParamsOrder = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateCreateTemplateDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failure_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failure_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failure_policy")))) {
		request.FailurePolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latest_version_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latest_version_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latest_version_time")))) {
		request.LatestVersionTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_template_id")))) {
		request.ParentTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_content")))) {
		request.RollbackTemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_type")))) {
		request.SoftwareType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_variant")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_variant")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_variant")))) {
		request.SoftwareVariant = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateCreateTemplateTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validation_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validation_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validation_errors")))) {
		request.ValidationErrors = expandRequestConfigurationTemplateCreateTemplateValidationErrors(ctx, key+".validation_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTags {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTags{}
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
		i := expandRequestConfigurationTemplateCreateTemplateTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTags(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTags {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplates {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplates{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplates(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplates {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelection {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelection {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues {
	var request catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateDeviceTypes {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateDeviceTypes{}
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
		i := expandRequestConfigurationTemplateCreateTemplateDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateDeviceTypes {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelection {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelectionSelectionValues {
	var request catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParams {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParams{}
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
		i := expandRequestConfigurationTemplateCreateTemplateTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParams {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateCreateTemplateTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateCreateTemplateTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsRange {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateCreateTemplateTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsRange {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsSelection {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateCreateTemplateTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsSelectionSelectionValues {
	var request catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateValidationErrors(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrors {
	request := catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_errors")))) {
		request.RollbackTemplateErrors = expandRequestConfigurationTemplateCreateTemplateValidationErrorsRollbackTemplateErrors(ctx, key+".rollback_template_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_errors")))) {
		request.TemplateErrors = expandRequestConfigurationTemplateCreateTemplateValidationErrorsTemplateErrors(ctx, key+".template_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_version")))) {
		request.TemplateVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateValidationErrorsRollbackTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrorsRollbackTemplateErrors {
	var request catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrorsRollbackTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateCreateTemplateValidationErrorsTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrorsTemplateErrors {
	var request catalystcentersdkgo.RequestConfigurationTemplatesCreateTemplateValidationErrorsTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplate(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplate {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplate{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateUpdateTemplateTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".author")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".author")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".author")))) {
		request.Author = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".containing_templates")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".containing_templates")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".containing_templates")))) {
		request.ContainingTemplates = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesArray(ctx, key+".containing_templates", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".create_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".create_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".create_time")))) {
		request.CreateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_params_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_params_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_params_order")))) {
		request.CustomParamsOrder = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateUpdateTemplateDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".failure_policy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".failure_policy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".failure_policy")))) {
		request.FailurePolicy = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".last_update_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".last_update_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".last_update_time")))) {
		request.LastUpdateTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latest_version_time")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latest_version_time")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latest_version_time")))) {
		request.LatestVersionTime = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_template_id")))) {
		request.ParentTemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_id")))) {
		request.ProjectID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_content")))) {
		request.RollbackTemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_type")))) {
		request.SoftwareType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_variant")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_variant")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_variant")))) {
		request.SoftwareVariant = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".software_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".software_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".software_version")))) {
		request.SoftwareVersion = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateUpdateTemplateTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".validation_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".validation_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".validation_errors")))) {
		request.ValidationErrors = expandRequestConfigurationTemplateUpdateTemplateValidationErrors(ctx, key+".validation_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTags {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTags{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTags(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTags {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplates {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplates{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplates(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplates(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplates {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplates{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTagsArray(ctx, key+".tags", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".composite")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".composite")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".composite")))) {
		request.Composite = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_types")))) {
		request.DeviceTypes = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesDeviceTypesArray(ctx, key+".device_types", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".language")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".language")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".language")))) {
		request.Language = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".project_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".project_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".project_name")))) {
		request.ProjectName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_params")))) {
		request.RollbackTemplateParams = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsArray(ctx, key+".rollback_template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_content")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_content")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_content")))) {
		request.TemplateContent = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_params")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_params")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_params")))) {
		request.TemplateParams = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsArray(ctx, key+".template_params", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTags(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues {
	var request catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelection {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues {
	var request catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateContainingTemplatesTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateDeviceTypesArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateDeviceTypes {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateDeviceTypes{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateDeviceTypes(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateDeviceTypes(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateDeviceTypes {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateDeviceTypes{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_family")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_family")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_family")))) {
		request.ProductFamily = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_series")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_series")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_series")))) {
		request.ProductSeries = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".product_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".product_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".product_type")))) {
		request.ProductType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelection {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateRollbackTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelectionSelectionValues {
	var request catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateRollbackTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParams {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParams{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateTemplateParams(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParams(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParams {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParams{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".binding")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".binding")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".binding")))) {
		request.Binding = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".custom_order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".custom_order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".custom_order")))) {
		request.CustomOrder = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data_type")))) {
		request.DataType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_value")))) {
		request.DefaultValue = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".group")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".group")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".group")))) {
		request.Group = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instruction_text")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instruction_text")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instruction_text")))) {
		request.InstructionText = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".not_param")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".not_param")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".not_param")))) {
		request.NotParam = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".order")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".order")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".order")))) {
		request.Order = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".param_array")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".param_array")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".param_array")))) {
		request.ParamArray = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parameter_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parameter_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parameter_name")))) {
		request.ParameterName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".provider")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".provider")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".provider")))) {
		request.Provider = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".range")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".range")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".range")))) {
		request.Range = expandRequestConfigurationTemplateUpdateTemplateTemplateParamsRangeArray(ctx, key+".range", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".required")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".required")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".required")))) {
		request.Required = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection")))) {
		request.Selection = expandRequestConfigurationTemplateUpdateTemplateTemplateParamsSelection(ctx, key+".selection.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsRangeArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange {
	request := []catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange{}
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
		i := expandRequestConfigurationTemplateUpdateTemplateTemplateParamsRange(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsRange(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsRange{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".max_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".max_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".max_value")))) {
		request.MaxValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".min_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".min_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".min_value")))) {
		request.MinValue = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsSelection(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelection {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelection{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_selected_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_selected_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_selected_values")))) {
		request.DefaultSelectedValues = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_type")))) {
		request.SelectionType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".selection_values")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".selection_values")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".selection_values")))) {
		request.SelectionValues = expandRequestConfigurationTemplateUpdateTemplateTemplateParamsSelectionSelectionValues(ctx, key+".selection_values.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateTemplateParamsSelectionSelectionValues(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelectionSelectionValues {
	var request catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateTemplateParamsSelectionSelectionValues
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateValidationErrors(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrors {
	request := catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrors{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rollback_template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rollback_template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rollback_template_errors")))) {
		request.RollbackTemplateErrors = expandRequestConfigurationTemplateUpdateTemplateValidationErrorsRollbackTemplateErrors(ctx, key+".rollback_template_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_errors")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_errors")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_errors")))) {
		request.TemplateErrors = expandRequestConfigurationTemplateUpdateTemplateValidationErrorsTemplateErrors(ctx, key+".template_errors.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_id")))) {
		request.TemplateID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".template_version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".template_version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".template_version")))) {
		request.TemplateVersion = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateValidationErrorsRollbackTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors {
	var request catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsRollbackTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestConfigurationTemplateUpdateTemplateValidationErrorsTemplateErrors(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors {
	var request catalystcentersdkgo.RequestConfigurationTemplatesUpdateTemplateValidationErrorsTemplateErrors
	request = d.Get(fixKeyAccess(key))
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchConfigurationTemplatesGetsTheTemplatesAvailable(m interface{}, queryParams catalystcentersdkgo.GetsAListOfProjectsQueryParams, vName string) (*catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplates, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplates
	log.Print("[DEBUG] 4")
	nResponse, _, err := client.ConfigurationTemplates.GetsAListOfProjects(&queryParams)
	log.Print("[DEBUG] 5")
	if err != nil {
		log.Printf("[DEBUG] Error when search => %s", err.Error())
		return foundItem, err
	}
	if nResponse == nil {
		log.Printf("[DEBUG] Error when search => %s", err.Error())
		return foundItem, err
	}
	//maxPageSize := 10
	log.Printf("[DEBUG] Start to find a concidence")
	for _, item := range *nResponse {
		for _, item2 := range *item.Templates {
			if vName == item2.Name {
				log.Printf("Vname %s , item2.Name %s", vName, item2.Name)
				foundItem = &item2
				return foundItem, err
			}
		}
	}
	return foundItem, err
}
