package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSitesDeviceCredentials() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Network Settings.

- Gets device credential settings for a site; *null* values indicate that the setting will be inherited from the parent
site; empty objects (*{}*) indicate that the credential is unset, and that no credential of that type will be used for
the site.
`,

		ReadContext: dataSourceSitesDeviceCredentialsRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Site Id, retrievable from the *id* attribute in */dna/intent/api/v1/sites*
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

						"cli_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"http_read_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"http_write_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmpv2c_read_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmpv2c_write_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"snmpv3_credentials_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"credentials_id": &schema.Schema{
										Description: `The *id* of the credentials.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"inherited_site_id": &schema.Schema{
										Description: `Inherited Site Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"inherited_site_name": &schema.Schema{
										Description: `Inherited Site Name`,
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

func dataSourceSitesDeviceCredentialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vInherited, okInherited := d.GetOk("inherited")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceCredentialSettingsForASiteV1")
		vvID := vID.(string)
		queryParams1 := catalystcentersdkgo.GetDeviceCredentialSettingsForASiteV1QueryParams{}

		if okInherited {
			queryParams1.Inherited = vInherited.(bool)
		}

		response1, restyResp1, err := client.NetworkSettings.GetDeviceCredentialSettingsForASiteV1(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceCredentialSettingsForASiteV1", err,
				"Failure at GetDeviceCredentialSettingsForASiteV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceCredentialSettingsForASiteV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1Item(item *catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["cli_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemCliCredentialsID(item.CliCredentialsID)
	respItem["snmpv2c_read_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemSNMPv2CReadCredentialsID(item.SNMPv2CReadCredentialsID)
	respItem["snmpv2c_write_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemSNMPv2CWriteCredentialsID(item.SNMPv2CWriteCredentialsID)
	respItem["snmpv3_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemSNMPv3CredentialsID(item.SNMPv3CredentialsID)
	respItem["http_read_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemHTTPReadCredentialsID(item.HTTPReadCredentialsID)
	respItem["http_write_credentials_id"] = flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemHTTPWriteCredentialsID(item.HTTPWriteCredentialsID)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemCliCredentialsID(item *catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseCliCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemSNMPv2CReadCredentialsID(item *catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv2CReadCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemSNMPv2CWriteCredentialsID(item *catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv2CWriteCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemSNMPv3CredentialsID(item *catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseSNMPv3CredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemHTTPReadCredentialsID(item *catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseHTTPReadCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}

func flattenNetworkSettingsGetDeviceCredentialSettingsForASiteV1ItemHTTPWriteCredentialsID(item *catalystcentersdkgo.ResponseNetworkSettingsGetDeviceCredentialSettingsForASiteV1ResponseHTTPWriteCredentialsID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["credentials_id"] = item.CredentialsID
	respItem["inherited_site_id"] = item.InheritedSiteID
	respItem["inherited_site_name"] = item.InheritedSiteName

	return []map[string]interface{}{
		respItem,
	}

}
