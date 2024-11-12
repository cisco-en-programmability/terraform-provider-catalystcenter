package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProfilingRulesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on AIEndpointAnalytics.

- This data source fetches the count of profiling rules based on the filter values provided in the query parameters. The
filter parameters are same as that of 'GET /profiling-rules' API, excluding the pagination and sort parameters.
`,

		ReadContext: dataSourceProfilingRulesCountRead,
		Schema: map[string]*schema.Schema{
			"include_deleted": &schema.Schema{
				Description: `includeDeleted query parameter. Flag to indicate whether deleted rules should be part of the records fetched.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"rule_type": &schema.Schema{
				Description: `ruleType query parameter. Use comma-separated list of rule types to filter the data. Defaults to 'Custom Rule'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceProfilingRulesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vRuleType, okRuleType := d.GetOk("rule_type")
	vIncludeDeleted, okIncludeDeleted := d.GetOk("include_deleted")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfProfilingRulesV1")
		queryParams1 := catalystcentersdkgo.GetCountOfProfilingRulesV1QueryParams{}

		if okRuleType {
			queryParams1.RuleType = vRuleType.(string)
		}
		if okIncludeDeleted {
			queryParams1.IncludeDeleted = vIncludeDeleted.(bool)
		}

		response1, restyResp1, err := client.AIEndpointAnalytics.GetCountOfProfilingRulesV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfProfilingRulesV1", err,
				"Failure at GetCountOfProfilingRulesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenAIEndpointAnalyticsGetCountOfProfilingRulesV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfProfilingRulesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}
