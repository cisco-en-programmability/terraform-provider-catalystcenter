package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDevicesNetworkProfilesForSites() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Retrieves the list of network profiles for sites.

- Retrieves a network profile for sites by id.
`,

		ReadContext: dataSourceNetworkDevicesNetworkProfilesForSitesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The *id* of the network profile, retrievable from *GET /intent/api/v1/networkProfilesForSites*
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page.
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
			"order": &schema.Schema{
				Description: `order query parameter. Whether ascending or descending order should be used to sort the response.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A property within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
				Description: `type query parameter. Filter responses to only include profiles of a given type
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `The ID of this network profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `The name of the network profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `The ID of this network profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `The name of the network profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkDevicesNetworkProfilesForSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vType, okType := d.GetOk("type")
	vID, okID := d.GetOk("id")

	method1 := []bool{okOffset, okLimit, okSortBy, okOrder, okType}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheListOfNetworkProfilesForSitesV1")
		queryParams1 := catalystcentersdkgo.RetrievesTheListOfNetworkProfilesForSitesV1QueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okType {
			queryParams1.Type = vType.(string)
		}

		response1, restyResp1, err := client.SiteDesign.RetrievesTheListOfNetworkProfilesForSitesV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheListOfNetworkProfilesForSitesV1", err,
				"Failure at RetrievesTheListOfNetworkProfilesForSitesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheListOfNetworkProfilesForSitesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: RetrieveANetworkProfileForSitesByIDV1")
		vvID := vID.(string)

		response2, restyResp2, err := client.SiteDesign.RetrieveANetworkProfileForSitesByIDV1(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveANetworkProfileForSitesByIDV1", err,
				"Failure at RetrieveANetworkProfileForSitesByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenSiteDesignRetrieveANetworkProfileForSitesByIDV1Item(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveANetworkProfileForSitesByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1Items(items *[]catalystcentersdkgo.ResponseSiteDesignRetrievesTheListOfNetworkProfilesForSitesV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSiteDesignRetrieveANetworkProfileForSitesByIDV1Item(item *catalystcentersdkgo.ResponseSiteDesignRetrieveANetworkProfileForSitesByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["type"] = item.Type
	return []map[string]interface{}{
		respItem,
	}
}
