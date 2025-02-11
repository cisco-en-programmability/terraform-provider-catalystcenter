package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkBugsResultsTrendCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of network bugs results trend over time
`,

		ReadContext: dataSourceNetworkBugsResultsTrendCountRead,
		Schema: map[string]*schema.Schema{
			"scan_time": &schema.Schema{
				Description: `scanTime query parameter. Return bugs trend with scanTime greater than this scanTime
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkBugsResultsTrendCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vScanTime, okScanTime := d.GetOk("scan_time")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfNetworkBugsResultsTrendOverTimeV1")
		queryParams1 := catalystcentersdkgo.GetCountOfNetworkBugsResultsTrendOverTimeV1QueryParams{}

		if okScanTime {
			queryParams1.ScanTime = vScanTime.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Compliance.GetCountOfNetworkBugsResultsTrendOverTimeV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfNetworkBugsResultsTrendOverTimeV1", err,
				"Failure at GetCountOfNetworkBugsResultsTrendOverTimeV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfNetworkBugsResultsTrendOverTimeV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1Item(item *catalystcentersdkgo.ResponseComplianceGetCountOfNetworkBugsResultsTrendOverTimeV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
