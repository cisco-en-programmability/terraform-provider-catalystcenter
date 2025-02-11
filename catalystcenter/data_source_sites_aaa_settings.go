package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesAAASettings() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Retrieve AAA settings for a site; *null* values indicate that the setting will be inherited from the parent site;
empty objects (*{}*) indicate that the setting is unset at a site.
`,

		ReadContext: dataSourceSitesAAASettingsRead,
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

						"aaa_client": &schema.Schema{
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

									"pan": &schema.Schema{
										Description: `Administration Node. Required for ISE.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"primary_server_ip": &schema.Schema{
										Description: `The server to use as a primary.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"secondary_server_ip": &schema.Schema{
										Description: `The server to use as a secondary.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"server_type": &schema.Schema{
										Description: `Server Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"shared_secret": &schema.Schema{
										Description: `Shared Secret`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"aaa_network": &schema.Schema{
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

									"pan": &schema.Schema{
										Description: `Administration Node. Required for ISE.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"primary_server_ip": &schema.Schema{
										Description: `The server to use as a primary.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"secondary_server_ip": &schema.Schema{
										Description: `The server to use as a secondary.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"server_type": &schema.Schema{
										Description: `Server Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"shared_secret": &schema.Schema{
										Description: `Shared Secret`,
										Type:        schema.TypeString,
										Computed:    true,
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

func dataSourceSitesAAASettingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vInherited, okInherited := d.GetOk("inherited")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveAAASettingsForASiteV1")
		vvID := vID.(string)
		queryParams1 := catalystcentersdkgo.RetrieveAAASettingsForASiteV1QueryParams{}

		if okInherited {
			queryParams1.Inherited = vInherited.(bool)
		}

		response1, restyResp1, err := client.NetworkSettings.RetrieveAAASettingsForASiteV1(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveAAASettingsForASiteV1", err,
				"Failure at RetrieveAAASettingsForASiteV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsRetrieveAAASettingsForASiteV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveAAASettingsForASiteV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsRetrieveAAASettingsForASiteV1Item(item *catalystcentersdkgo.ResponseNetworkSettingsRetrieveAAASettingsForASiteV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aaa_network"] = flattenNetworkSettingsRetrieveAAASettingsForASiteV1ItemAAANetwork(item.AAANetwork)
	respItem["aaa_client"] = flattenNetworkSettingsRetrieveAAASettingsForASiteV1ItemAAAClient(item.AAAClient)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsRetrieveAAASettingsForASiteV1ItemAAANetwork(item *catalystcentersdkgo.ResponseNetworkSettingsRetrieveAAASettingsForASiteV1ResponseAAANetwork) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["server_type"] = item.ServerType
	respItem["protocol"] = item.Protocol
	respItem["pan"] = item.Pan
	respItem["primary_server_ip"] = item.PrimaryServerIP
	respItem["secondary_server_ip"] = item.SecondaryServerIP
	respItem["shared_secret"] = item.SharedSecret
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsRetrieveAAASettingsForASiteV1ItemAAAClient(item *catalystcentersdkgo.ResponseNetworkSettingsRetrieveAAASettingsForASiteV1ResponseAAAClient) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["server_type"] = item.ServerType
	respItem["protocol"] = item.Protocol
	respItem["pan"] = item.Pan
	respItem["primary_server_ip"] = item.PrimaryServerIP
	respItem["secondary_server_ip"] = item.SecondaryServerIP
	respItem["shared_secret"] = item.SharedSecret
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}
