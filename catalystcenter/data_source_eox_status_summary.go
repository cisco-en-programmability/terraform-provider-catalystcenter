package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEoxStatusSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on EoX.

- Retrieves EoX summary for all devices in the network
`,

		ReadContext: dataSourceEoxStatusSummaryRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"hardware_count": &schema.Schema{
							Description: `Number of hardware EoX alerts detected on the network
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"module_count": &schema.Schema{
							Description: `Number of module EoX alerts detected on the network
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"software_count": &schema.Schema{
							Description: `Number of software EoX alerts detected on the network
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"total_count": &schema.Schema{
							Description: `Total number of EoX alerts detected on the network. This is the sum of hardwareCount, softwareCount and moduleCount.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEoxStatusSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEoxSummary")

		// has_unknown_response: None

		response1, restyResp1, err := client.Eox.GetEoxSummary()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetEoxSummary", err,
				"Failure at GetEoxSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetEoxSummary", err,
				"Failure at GetEoxSummary, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenEoxGetEoxSummaryItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEoxSummary response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEoxGetEoxSummaryItem(item *catalystcentersdkgo.ResponseEoxGetEoxSummaryResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["hardware_count"] = item.HardwareCount
	respItem["software_count"] = item.SoftwareCount
	respItem["module_count"] = item.ModuleCount
	respItem["total_count"] = item.TotalCount
	return []map[string]interface{}{
		respItem,
	}
}
