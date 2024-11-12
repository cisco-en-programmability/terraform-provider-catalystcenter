package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDnaCommandRunnerKeywords() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Command Runner.

- Get valid keywords
`,

		ReadContext: dataSourceDnaCommandRunnerKeywordsRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDnaCommandRunnerKeywordsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllKeywordsOfClisAcceptedByCommandRunnerV1")

		response1, restyResp1, err := client.CommandRunner.GetAllKeywordsOfClisAcceptedByCommandRunnerV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllKeywordsOfClisAcceptedByCommandRunnerV1", err,
				"Failure at GetAllKeywordsOfClisAcceptedByCommandRunnerV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenCommandRunnerGetAllKeywordsOfClisAcceptedByCommandRunnerV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllKeywordsOfClisAcceptedByCommandRunnerV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenCommandRunnerGetAllKeywordsOfClisAcceptedByCommandRunnerV1Items(items *catalystcentersdkgo.ResponseCommandRunnerGetAllKeywordsOfClisAcceptedByCommandRunnerV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = items.Response
	respItem["version"] = items.Version
	return []map[string]interface{}{
		respItem,
	}
}
