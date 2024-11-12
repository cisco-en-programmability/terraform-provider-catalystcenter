package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesImageDistributionSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieve image distribution settings for a site; *null* values indicate that the setting will be inherited from the
parent site; empty objects (*{}*) indicate that the setting is unset at a site.
`,

		ReadContext: dataSourceSitesImageDistributionSettingsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Site Id
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"inherited": &schema.Schema{
				Description: `_inherited query parameter. Include settings explicitly set for this site and settings inherited from sites higher in the site hierarchy; when *false*, *null* values indicate that the site inherits that setting from the parent site or a site higher in the site hierarchy.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"image_distribution": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"servers": &schema.Schema{
										Description: `This field holds an array of unique identifiers representing image distribution servers. SFTP servers to act as image distribution servers. A distributed SWIM architecture, using suitably located SFTP servers, can help support large-scale device software image upgrades and conserve WAN bandwidth.
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceSitesImageDistributionSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vInherited, okInherited := d.GetOk("inherited")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveImageDistributionSettingsForASiteV1")
		vvID := vID.(string)
		queryParams1 := catalystcentersdkgo.RetrieveImageDistributionSettingsForASiteV1QueryParams{}

		if okInherited {
			queryParams1.Inherited = vInherited.(bool)
		}

		response1, restyResp1, err := client.NetworkSettings.RetrieveImageDistributionSettingsForASiteV1(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveImageDistributionSettingsForASiteV1", err,
				"Failure at RetrieveImageDistributionSettingsForASiteV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveImageDistributionSettingsForASiteV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveImageDistributionSettingsForASiteV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrieveImageDistributionSettingsForASiteV1Item(item *catalystcentersdkgo.ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["image_distribution"] = flattenNetworkSettingsRetrieveImageDistributionSettingsForASiteV1ItemImageDistribution(item.ImageDistribution)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrieveImageDistributionSettingsForASiteV1ItemImageDistribution(item *catalystcentersdkgo.ResponseNetworkSettingsRetrieveImageDistributionSettingsForASiteV1ResponseImageDistribution) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["servers"] = item.Servers
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}
