package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesWirelessSettingsSSIDsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get count of all SSIDs (Service Set Identifier) .
`,

		ReadContext: dataSourceSitesWirelessSettingsSSIDsCountRead,
		Schema: map[string]*schema.Schema{
			"inherited": &schema.Schema{
				Description: `_inherited query parameter. This query parameter indicates whether the current SSID count at the given 'siteId' is of the SSID(s) it is inheriting or count of non-inheriting SSID(s)
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId path parameter. Site UUID
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
							Description: `Count of the requested resource
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

func dataSourceSitesWirelessSettingsSSIDsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID := d.Get("site_id")
	vInherited, okInherited := d.GetOk("inherited")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSSIDCountBySite")
		vvSiteID := vSiteID.(string)
		queryParams1 := catalystcentersdkgo.GetSSIDCountBySiteQueryParams{}

		if okInherited {
			queryParams1.Inherited = vInherited.(bool)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.GetSSIDCountBySite(vvSiteID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSSIDCountBySite", err,
				"Failure at GetSSIDCountBySite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSSIDCountBySite", err,
				"Failure at GetSSIDCountBySite, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetSSIDCountBySiteItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSSIDCountBySite response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetSSIDCountBySiteItem(item *catalystcentersdkgo.ResponseWirelessGetSSIDCountBySiteResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
