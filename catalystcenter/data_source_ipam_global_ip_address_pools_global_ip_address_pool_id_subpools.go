package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIPamGlobalIPAddressPoolsGlobalIPAddressPoolIDSubpools() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieves subpools IDs of a global IP address pool.  The IDs can be fetched with
/dna/intent/api/v1/ipam/siteIpAddressPools/{id}
`,

		ReadContext: dataSourceIPamGlobalIPAddressPoolsGlobalIPAddressPoolIDSubpoolsRead,
		Schema: map[string]*schema.Schema{
			"global_ip_address_pool_id": &schema.Schema{
				Description: `globalIpAddressPoolId path parameter. The id of the global IP address pool for which to retrieve subpool IDs.
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page;The minimum is 1, and the maximum is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `ID of the subpool
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

func dataSourceIPamGlobalIPAddressPoolsGlobalIPAddressPoolIDSubpoolsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vGlobalIPAddressPoolID := d.Get("global_ip_address_pool_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1")
		vvGlobalIPAddressPoolID := vGlobalIPAddressPoolID.(string)
		queryParams1 := catalystcentersdkgo.RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1QueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.NetworkSettings.RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1(vvGlobalIPAddressPoolID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1", err,
				"Failure at RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1Items(items *[]catalystcentersdkgo.ResponseNetworkSettingsRetrievesSubpoolsIDsOfAGlobalIPAddressPoolV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}
