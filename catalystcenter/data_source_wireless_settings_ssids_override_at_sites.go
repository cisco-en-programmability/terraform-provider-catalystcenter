package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsSSIDsOverrideAtSites() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieve list of siteId(s) with information of SSID(s) which are overridden
`,

		ReadContext: dataSourceWirelessSettingsSSIDsOverrideAtSitesRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Site UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"site_id": &schema.Schema{
							Description: `Site ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssids": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `SSID ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssid": &schema.Schema{
										Description: `SSID
`,
										Type:     schema.TypeString,
										Computed: true,
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

func dataSourceWirelessSettingsSSIDsOverrideAtSitesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveSitesWithOverriddenSSIDsV1")
		queryParams1 := catalystcentersdkgo.RetrieveSitesWithOverriddenSSIDsV1QueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.RetrieveSitesWithOverriddenSSIDsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveSitesWithOverriddenSSIDsV1", err,
				"Failure at RetrieveSitesWithOverriddenSSIDsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessRetrieveSitesWithOverriddenSSIDsV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveSitesWithOverriddenSSIDsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessRetrieveSitesWithOverriddenSSIDsV1Items(items *[]catalystcentersdkgo.ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["site_id"] = item.SiteID
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItem["ssids"] = flattenWirelessRetrieveSitesWithOverriddenSSIDsV1ItemsSSIDs(item.SSIDs)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessRetrieveSitesWithOverriddenSSIDsV1ItemsSSIDs(items *[]catalystcentersdkgo.ResponseWirelessRetrieveSitesWithOverriddenSSIDsV1ResponseSSIDs) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["ssid"] = item.SSID
		respItems = append(respItems, respItem)
	}
	return respItems
}
