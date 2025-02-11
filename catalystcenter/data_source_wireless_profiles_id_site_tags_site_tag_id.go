package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessProfilesIDSiteTagsSiteTagID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This endpoint retrieves the details of a specific Site Tag associated with a given Wireless Profile. This data
source requires the id of the Wireless Profile and the siteTagId of the Site Tag.
`,

		ReadContext: dataSourceWirelessProfilesIDSiteTagsSiteTagIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Wireless Profile Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"site_tag_id": &schema.Schema{
				Description: `siteTagId path parameter. Site Tag Id
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_profile_name": &schema.Schema{
							Description: `Ap Profile Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"flex_profile_name": &schema.Schema{
							Description: `Flex Profile Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_ids": &schema.Schema{
							Description: `Site Ids`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"site_tag_id": &schema.Schema{
							Description: `Site Tag Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"site_tag_name": &schema.Schema{
							Description: `Use English letters, numbers, special characters except <, /, '.*', ? and leading/trailing space
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

func dataSourceWirelessProfilesIDSiteTagsSiteTagIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vSiteTagID := d.Get("site_tag_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveASpecificSiteTagForAWirelessProfileV1")
		vvID := vID.(string)
		vvSiteTagID := vSiteTagID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.RetrieveASpecificSiteTagForAWirelessProfileV1(vvID, vvSiteTagID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveASpecificSiteTagForAWirelessProfileV1", err,
				"Failure at RetrieveASpecificSiteTagForAWirelessProfileV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessRetrieveASpecificSiteTagForAWirelessProfileV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveASpecificSiteTagForAWirelessProfileV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessRetrieveASpecificSiteTagForAWirelessProfileV1Item(item *catalystcentersdkgo.ResponseWirelessRetrieveASpecificSiteTagForAWirelessProfileV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["site_ids"] = item.SiteIDs
	respItem["site_tag_name"] = item.SiteTagName
	respItem["flex_profile_name"] = item.FlexProfileName
	respItem["ap_profile_name"] = item.ApProfileName
	respItem["site_tag_id"] = item.SiteTagID
	return []map[string]interface{}{
		respItem,
	}
}
