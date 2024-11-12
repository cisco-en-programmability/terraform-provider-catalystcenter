package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLanAutomationSessions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on LAN Automation.

- Invoke this API to get the LAN Automation active session information
`,

		ReadContext: dataSourceLanAutomationSessionsRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"active_session_ids": &schema.Schema{
							Description: `List of Active LAN Automation IDs
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"active_sessions": &schema.Schema{
							Description: `Current active sessions count
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"max_supported_count": &schema.Schema{
							Description: `Maximum supported parallel sessions count
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

func dataSourceLanAutomationSessionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LanAutomationActiveSessionsV1")

		response1, restyResp1, err := client.LanAutomation.LanAutomationActiveSessionsV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LanAutomationActiveSessionsV1", err,
				"Failure at LanAutomationActiveSessionsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLanAutomationLanAutomationActiveSessionsV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomationActiveSessionsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLanAutomationLanAutomationActiveSessionsV1Item(item *catalystcentersdkgo.ResponseLanAutomationLanAutomationActiveSessionsV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["max_supported_count"] = item.MaxSupportedCount
	respItem["active_sessions"] = item.ActiveSessions
	respItem["active_session_ids"] = item.ActiveSessionIDs
	return []map[string]interface{}{
		respItem,
	}
}
