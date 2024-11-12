package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesNtpSettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieve NTP settings for a site; *null* values indicate that the setting will be inherited from the parent site;
empty objects (*{}*) indicate that the setting is unset at a site.
`,

		ReadContext: dataSourceSitesNtpSettingsRead,
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

						"ntp": &schema.Schema{
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
										Description: `NTP servers to facilitate system clock synchronization for your network.
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

func dataSourceSitesNtpSettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vInherited, okInherited := d.GetOk("inherited")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveNTPSettingsForASiteV1")
		vvID := vID.(string)
		queryParams1 := catalystcentersdkgo.RetrieveNTPSettingsForASiteV1QueryParams{}

		if okInherited {
			queryParams1.Inherited = vInherited.(bool)
		}

		response1, restyResp1, err := client.NetworkSettings.RetrieveNTPSettingsForASiteV1(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveNTPSettingsForASiteV1", err,
				"Failure at RetrieveNTPSettingsForASiteV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveNTPSettingsForASiteV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveNTPSettingsForASiteV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrieveNTPSettingsForASiteV1Item(item *catalystcentersdkgo.ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ntp"] = flattenNetworkSettingsRetrieveNTPSettingsForASiteV1ItemNtp(item.Ntp)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrieveNTPSettingsForASiteV1ItemNtp(item *catalystcentersdkgo.ResponseNetworkSettingsRetrieveNTPSettingsForASiteV1ResponseNtp) []map[string]interface{} {
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
