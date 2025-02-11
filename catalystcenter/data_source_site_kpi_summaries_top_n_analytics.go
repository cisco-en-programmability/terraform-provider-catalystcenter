package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSiteKpiSummariesTopNAnalytics() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Gets the topN analytics data for a given taskId. For detailed information about the usage of the API, please refer to
the Open API specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-SiteKpiSummaries-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceSiteKpiSummariesTopNAnalyticsRead,
		Schema: map[string]*schema.Schema{
			"task_id": &schema.Schema{
				Description: `taskId query parameter. used to retrieve asynchronously processed & stored data. When this parameter is used, the rest of the request params will be ignored.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
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

									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"value": &schema.Schema{
										Description: `Value`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSiteKpiSummariesTopNAnalyticsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vTaskID := d.Get("task_id")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1")

		headerParams1 := catalystcentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1HeaderParams{}
		queryParams1 := catalystcentersdkgo.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1QueryParams{}

		queryParams1.TaskID = vTaskID.(string)

		headerParams1.XCaLLERID = vXCaLLERID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Sites.GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1", err,
				"Failure at GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1Items(items *[]catalystcentersdkgo.ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["attributes"] = flattenSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1ItemsAttributes(item.Attributes)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1ItemsAttributes(items *[]catalystcentersdkgo.ResponseSitesGetTopNEntitiesRelatedToSiteAnalyticsForTheGivenTaskIDV1ResponseAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["value"] = item.Value
		respItems = append(respItems, respItem)
	}
	return respItems
}
