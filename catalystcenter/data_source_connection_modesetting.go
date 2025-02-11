package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceConnectionModesetting() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Retrieves Cisco Smart Software Manager (CSSM) connection mode setting.
`,

		ReadContext: dataSourceConnectionModesettingRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"connection_mode": &schema.Schema{
							Description: `The CSSM connection modes of Catalyst Center are DIRECT, ON_PREMISE and SMART_PROXY
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parameters": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"client_id": &schema.Schema{
										Description: `On-premise CSSM client id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"on_premise_host": &schema.Schema{
										Description: `On-premise host
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"smart_account_name": &schema.Schema{
										Description: `On-premise CSSM local smart account name
`,
										Type:     schema.TypeString,
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

func dataSourceConnectionModesettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesCSSMConnectionModeV1")

		// has_unknown_response: None

		response1, restyResp1, err := client.Licenses.RetrievesCSSMConnectionModeV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesCSSMConnectionModeV1", err,
				"Failure at RetrievesCSSMConnectionModeV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesRetrievesCSSMConnectionModeV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesCSSMConnectionModeV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesRetrievesCSSMConnectionModeV1Item(item *catalystcentersdkgo.ResponseLicensesRetrievesCSSMConnectionModeV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connection_mode"] = item.ConnectionMode
	respItem["parameters"] = flattenLicensesRetrievesCSSMConnectionModeV1ItemParameters(item.Parameters)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenLicensesRetrievesCSSMConnectionModeV1ItemParameters(item *catalystcentersdkgo.ResponseLicensesRetrievesCSSMConnectionModeV1ResponseParameters) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["on_premise_host"] = item.OnPremiseHost
	respItem["smart_account_name"] = item.SmartAccountName
	respItem["client_id"] = item.ClientID

	return []map[string]interface{}{
		respItem,
	}

}
