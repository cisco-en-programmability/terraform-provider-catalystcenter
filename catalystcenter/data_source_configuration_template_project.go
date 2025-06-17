package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConfigurationTemplateProject() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Templates.

- List the projects

- Get the details of the given project by its id.
`,

		ReadContext: dataSourceConfigurationTemplateProjectRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `name query parameter. Name of project to be searched
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"project_id": &schema.Schema{
				Description: `projectId path parameter. projectId(UUID) of project to get project details
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

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of project
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

						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
							Type:     schema.TypeString, //TEST,
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

						"create_time": &schema.Schema{
							Description: `Create time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `UUID of project
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_update_time": &schema.Schema{
							Description: `Update time of project
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of project
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

						"templates": &schema.Schema{
							Description: `List of templates within the project
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceConfigurationTemplateProjectRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vSortOrder, okSortOrder := d.GetOk("sort_order")
	vProjectID, okProjectID := d.GetOk("project_id")

	method1 := []bool{okName, okSortOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okProjectID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsAListOfProjects")
		queryParams1 := catalystcentersdkgo.GetsAListOfProjectsQueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okSortOrder {
			queryParams1.SortOrder = vSortOrder.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.ConfigurationTemplates.GetsAListOfProjects(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsAListOfProjects", err,
				"Failure at GetsAListOfProjects, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsAListOfProjects", err,
				"Failure at GetsAListOfProjects, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenConfigurationTemplatesGetsAListOfProjectsItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAListOfProjects response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetsTheDetailsOfAGivenProject")
		vvProjectID := vProjectID.(string)

		// has_unknown_response: None

		response2, restyResp2, err := client.ConfigurationTemplates.GetsTheDetailsOfAGivenProject(vvProjectID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsTheDetailsOfAGivenProject", err,
				"Failure at GetsTheDetailsOfAGivenProject, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItem(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsTheDetailsOfAGivenProject response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationTemplatesGetsAListOfProjectsItems(items *catalystcentersdkgo.ResponseConfigurationTemplatesGetsAListOfProjects) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTags(item.Tags)
		respItem["create_time"] = item.CreateTime
		respItem["description"] = item.Description
		respItem["id"] = item.ID
		respItem["last_update_time"] = item.LastUpdateTime
		respItem["name"] = item.Name
		respItem["templates"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplates(item.Templates)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTags(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplates(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTags(item.Tags)
		respItem["author"] = item.Author
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["containing_templates"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplates(item.ContainingTemplates)
		respItem["create_time"] = item.CreateTime
		respItem["custom_params_order"] = boolPtrToString(item.CustomParamsOrder)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesDeviceTypes(item.DeviceTypes)
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
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["software_type"] = item.SoftwareType
		respItem["software_variant"] = item.SoftwareVariant
		respItem["software_version"] = item.SoftwareVersion
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParams(item.TemplateParams)
		respItem["validation_errors"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrors(item.ValidationErrors)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTags(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplates(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplates) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["tags"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTags(item.Tags)
		respItem["composite"] = boolPtrToString(item.Composite)
		respItem["description"] = item.Description
		respItem["device_types"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesDeviceTypes(item.DeviceTypes)
		respItem["id"] = item.ID
		respItem["language"] = item.Language
		respItem["name"] = item.Name
		respItem["project_name"] = item.ProjectName
		respItem["rollback_template_params"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParams(item.RollbackTemplateParams)
		respItem["template_content"] = item.TemplateContent
		respItem["template_params"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParams(item.TemplateParams)
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTags(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesDeviceTypes(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesDeviceTypes) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParams(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsRange(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelection(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParams(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsRange(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsSelection(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesContainingTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesDeviceTypes(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesDeviceTypes) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParams(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsRange(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsSelection(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesRollbackTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesRollbackTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParams(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParams) []map[string]interface{} {
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
		respItem["range"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsRange(item.Range)
		respItem["required"] = boolPtrToString(item.Required)
		respItem["selection"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsSelection(item.Selection)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsRange(items *[]catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsRange) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsSelection(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelection) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_selected_values"] = item.DefaultSelectedValues
	respItem["id"] = item.ID
	respItem["selection_type"] = item.SelectionType
	respItem["selection_values"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsSelectionSelectionValues(item.SelectionValues)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesTemplateParamsSelectionSelectionValues(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesTemplateParamsSelectionSelectionValues) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrors(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrors) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["rollback_template_errors"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrorsRollbackTemplateErrors(item.RollbackTemplateErrors)
	respItem["template_errors"] = flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrorsTemplateErrors(item.TemplateErrors)
	respItem["template_id"] = item.TemplateID
	respItem["template_version"] = item.TemplateVersion

	return []map[string]interface{}{
		respItem,
	}

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrorsRollbackTemplateErrors(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsRollbackTemplateErrors) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsAListOfProjectsItemsTemplatesValidationErrorsTemplateErrors(item *catalystcentersdkgo.ResponseItemConfigurationTemplatesGetsAListOfProjectsTemplatesValidationErrorsTemplateErrors) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItem(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProject) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tags"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTags(item.Tags)
	respItem["create_time"] = item.CreateTime
	respItem["description"] = item.Description
	respItem["id"] = item.ID
	respItem["last_update_time"] = item.LastUpdateTime
	respItem["name"] = item.Name
	respItem["templates"] = flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplates(item.Templates)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTags(items *[]catalystcentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTags) []map[string]interface{} {
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

func flattenConfigurationTemplatesGetsTheDetailsOfAGivenProjectItemTemplates(item *catalystcentersdkgo.ResponseConfigurationTemplatesGetsTheDetailsOfAGivenProjectTemplates) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
