package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesProfileAssignmentsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieves the count of profiles that the given site has been assigned.  These profiles may either be directly assigned
to this site, or were assigned to a parent site and have been inherited.
`,

		ReadContext: dataSourceSitesProfileAssignmentsCountRead,
		Schema: map[string]*schema.Schema{
			"site_id": &schema.Schema{
				Description: `siteId path parameter. The *id* of the site, retrievable from */dna/intent/api/v1/sites*
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

func dataSourceSitesProfileAssignmentsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1")
		vvSiteID := vSiteID.(string)

		response1, restyResp1, err := client.SiteDesign.RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1(vvSiteID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1", err,
				"Failure at RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1Item(item *catalystcentersdkgo.ResponseSiteDesignRetrievesTheCountOfProfilesThatTheGivenSiteHasBeenAssignedV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
