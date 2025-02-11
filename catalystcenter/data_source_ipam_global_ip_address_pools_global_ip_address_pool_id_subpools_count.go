package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPamGlobalIPAddressPoolsGlobalIPAddressPoolIDSubpoolsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Counts subpools of a global IP address pool.
`,

		ReadContext: dataSourceIPamGlobalIPAddressPoolsGlobalIPAddressPoolIDSubpoolsCountRead,
		Schema: map[string]*schema.Schema{
			"global_ip_address_pool_id": &schema.Schema{
				Description: `globalIpAddressPoolId path parameter. The id of the global IP address pool for which to count subpools.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count.
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

func dataSourceIPamGlobalIPAddressPoolsGlobalIPAddressPoolIDSubpoolsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vGlobalIPAddressPoolID := d.Get("global_ip_address_pool_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CountsSubpoolsOfAGlobalIPAddressPoolV1")
		vvGlobalIPAddressPoolID := vGlobalIPAddressPoolID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.NetworkSettings.CountsSubpoolsOfAGlobalIPAddressPoolV1(vvGlobalIPAddressPoolID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CountsSubpoolsOfAGlobalIPAddressPoolV1", err,
				"Failure at CountsSubpoolsOfAGlobalIPAddressPoolV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountsSubpoolsOfAGlobalIPAddressPoolV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1Item(item *catalystcentersdkgo.ResponseNetworkSettingsCountsSubpoolsOfAGlobalIPAddressPoolV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
