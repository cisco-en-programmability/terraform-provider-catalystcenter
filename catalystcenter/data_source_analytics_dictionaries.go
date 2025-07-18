package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAnalyticsDictionaries() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on AI Endpoint Analytics.

- Fetches the list of attribute dictionaries.
`,

		ReadContext: dataSourceAnalyticsDictionariesRead,
		Schema: map[string]*schema.Schema{
			"include_attributes": &schema.Schema{
				Description: `includeAttributes query parameter. Flag to indicate whether attribute list for each dictionary should be included in response.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Description of the attribute.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the attribute.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"description": &schema.Schema{
							Description: `Description of the dictionary.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the dictionary.
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

func dataSourceAnalyticsDictionariesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vIncludeAttributes, okIncludeAttributes := d.GetOk("include_attributes")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAiEndpointAnalyticsAttributeDictionaries")
		queryParams1 := catalystcentersdkgo.GetAiEndpointAnalyticsAttributeDictionariesQueryParams{}

		if okIncludeAttributes {
			queryParams1.IncludeAttributes = vIncludeAttributes.(bool)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.AiEndpointAnalytics.GetAiEndpointAnalyticsAttributeDictionaries(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAiEndpointAnalyticsAttributeDictionaries", err,
				"Failure at GetAiEndpointAnalyticsAttributeDictionaries, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAiEndpointAnalyticsAttributeDictionaries", err,
				"Failure at GetAiEndpointAnalyticsAttributeDictionaries, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenAiEndpointAnalyticsGetAiEndpointAnalyticsAttributeDictionariesItems(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAiEndpointAnalyticsAttributeDictionaries response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenAiEndpointAnalyticsGetAiEndpointAnalyticsAttributeDictionariesItems(items *catalystcentersdkgo.ResponseAiEndpointAnalyticsGetAiEndpointAnalyticsAttributeDictionaries) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["attributes"] = flattenAiEndpointAnalyticsGetAiEndpointAnalyticsAttributeDictionariesItemsAttributes(item.Attributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenAiEndpointAnalyticsGetAiEndpointAnalyticsAttributeDictionariesItemsAttributes(items *[]catalystcentersdkgo.ResponseItemAiEndpointAnalyticsGetAiEndpointAnalyticsAttributeDictionariesAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItems = append(respItems, respItem)
	}
	return respItems
}
