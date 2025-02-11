package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTemplatesTemplateIDVersionsVersionID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- Get a template's version by the version ID.
`,

		ReadContext: dataSourceTemplatesTemplateIDVersionsVersionIDRead,
		Schema: map[string]*schema.Schema{
			"template_id": &schema.Schema{
				Description: `templateId path parameter. The id of the template to get versions of, retrieveable from GET /dna/intent/api/v1/templates
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"version_id": &schema.Schema{
				Description: `versionId path parameter. The id of the versioned template to get versions of, retrieveable from GET /dna/intent/api/v1/templates/{id}/versions
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"composite_template": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Description: `Author of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_policy": &schema.Schema{
										Description: `Policy to handle failure only applicable for composite templates  CONTINUE_ON_ERROR: If a composed template fails while deploying a device, continue deploying the next composed template  ABORT_TARGET_ON_ERROR: If a composed template fails while deploying to a device, abort the subsequent composed templates to that device if there any remaining
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_update_time": &schema.Schema{
										Description: `Timestamp of when the template was updated or modified
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"products": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Family name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_name": &schema.Schema{
													Description: `Name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_series": &schema.Schema{
													Description: `Series name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"project_id": &schema.Schema{
										Description: `Id of the project
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"software_family": &schema.Schema{
										Description: `Software Family`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"template_id": &schema.Schema{
										Description: `The id of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `The type of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"regular_template": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"author": &schema.Schema{
										Description: `Author of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"language": &schema.Schema{
										Description: `Language of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_update_time": &schema.Schema{
										Description: `Timestamp of when the template was updated or modified
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"products": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"product_family": &schema.Schema{
													Description: `Family name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_name": &schema.Schema{
													Description: `Name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"product_series": &schema.Schema{
													Description: `Series name of the product
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"project_id": &schema.Schema{
										Description: `Id of the project
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"software_family": &schema.Schema{
										Description: `Software Family`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"template_content": &schema.Schema{
										Description: `Template content (uses LF styling for line-breaks)
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"template_id": &schema.Schema{
										Description: `The id of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type": &schema.Schema{
										Description: `The type of the template
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"version": &schema.Schema{
							Description: `The version number of this version
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version_id": &schema.Schema{
							Description: `The id of this version
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version_time": &schema.Schema{
							Description: `Time at which this version was committed
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTemplatesTemplateIDVersionsVersionIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vTemplateID := d.Get("template_id")
	vVersionID := d.Get("version_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTemplateVersionV1")
		vvTemplateID := vTemplateID.(string)
		vvVersionID := vVersionID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.ConfigurationTemplates.GetTemplateVersionV1(vvTemplateID, vvVersionID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTemplateVersionV1", err,
				"Failure at GetTemplateVersionV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationTemplatesGetTemplateVersionV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTemplateVersionV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetTemplateVersionV1Item(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["version_id"] = item.VersionID
	respItem["version"] = item.Version
	respItem["version_time"] = item.VersionTime
	respItem["regular_template"] = flattenConfigurationTemplatesGetTemplateVersionV1ItemRegularTemplate(item.RegularTemplate)
	respItem["composite_template"] = flattenConfigurationTemplatesGetTemplateVersionV1ItemCompositeTemplate(item.CompositeTemplate)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesGetTemplateVersionV1ItemRegularTemplate(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionV1ResponseRegularTemplate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["template_id"] = item.TemplateID
	respItem["name"] = item.Name
	respItem["project_id"] = item.ProjectID
	respItem["description"] = item.Description
	respItem["software_family"] = item.SoftwareFamily
	respItem["author"] = item.Author
	respItem["products"] = flattenConfigurationTemplatesGetTemplateVersionV1ItemRegularTemplateProducts(item.Products)
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["type"] = item.Type
	respItem["language"] = item.Language
	respItem["template_content"] = item.TemplateContent

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplateVersionV1ItemRegularTemplateProducts(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionV1ResponseRegularTemplateProducts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_name"] = item.ProductName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetTemplateVersionV1ItemCompositeTemplate(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionV1ResponseCompositeTemplate) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["template_id"] = item.TemplateID
	respItem["name"] = item.Name
	respItem["project_id"] = item.ProjectID
	respItem["description"] = item.Description
	respItem["software_family"] = item.SoftwareFamily
	respItem["author"] = item.Author
	respItem["products"] = flattenConfigurationTemplatesGetTemplateVersionV1ItemCompositeTemplateProducts(item.Products)
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["type"] = item.Type
	respItem["failure_policy"] = item.FailurePolicy

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetTemplateVersionV1ItemCompositeTemplateProducts(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetTemplateVersionV1ResponseCompositeTemplateProducts) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["product_family"] = item.ProductFamily
		respItem["product_series"] = item.ProductSeries
		respItem["product_name"] = item.ProductName
		respItems = append(respItems, respItem)
	}
	return respItems
}
