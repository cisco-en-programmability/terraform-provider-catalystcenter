package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesNetworkProfilesForSitesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieves the count of network profiles for sites
`,

		ReadContext: dataSourceNetworkDevicesNetworkProfilesForSitesCountRead,
		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Description: `type query parameter. Filter the response to only count profiles of a given type
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

func dataSourceNetworkDevicesNetworkProfilesForSitesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vType, okType := d.GetOk("type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCountOfNetworkProfilesForSites")
		queryParams1 := catalystcentersdkgo.RetrievesTheCountOfNetworkProfilesForSitesQueryParams{}

		if okType {
			queryParams1.Type = vType.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.SiteDesign.RetrievesTheCountOfNetworkProfilesForSites(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCountOfNetworkProfilesForSites", err,
				"Failure at RetrievesTheCountOfNetworkProfilesForSites, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCountOfNetworkProfilesForSites", err,
				"Failure at RetrievesTheCountOfNetworkProfilesForSites, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignRetrievesTheCountOfNetworkProfilesForSitesItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCountOfNetworkProfilesForSites response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignRetrievesTheCountOfNetworkProfilesForSitesItem(item *catalystcentersdkgo.ResponseSiteDesignRetrievesTheCountOfNetworkProfilesForSitesResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
