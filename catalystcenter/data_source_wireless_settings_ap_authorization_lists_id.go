package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsApAuthorizationListsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get an AP Authorization List by AP Authorization List ID that captured in wireless
settings design.
`,

		ReadContext: dataSourceWirelessSettingsApAuthorizationListsIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. AP Authorization List ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_authorization_list_name": &schema.Schema{
							Description: `Ap Authorization List Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"local_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_mac_entries": &schema.Schema{
										Description: `AP Mac Addresses`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ap_serial_number_entries": &schema.Schema{
										Description: `AP Serial Number Entries`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"remote_authorization": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"aaa_servers": &schema.Schema{
										Description: `AAA Servers`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"authorize_ap_with_mac": &schema.Schema{
										Description: `Authorize AP With Mac`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
										Computed: true,
									},

									"authorize_ap_with_serial_number": &schema.Schema{
										Description: `Authorize AP With Serial Number`,
										// Type:        schema.TypeBool,
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

func dataSourceWirelessSettingsApAuthorizationListsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApAuthorizationListByIDV1")
		vvID := vID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.GetApAuthorizationListByIDV1(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetApAuthorizationListByIDV1", err,
				"Failure at GetApAuthorizationListByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetApAuthorizationListByIDV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApAuthorizationListByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetApAuthorizationListByIDV1Item(item *catalystcentersdkgo.ResponseWirelessGetApAuthorizationListByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["ap_authorization_list_name"] = item.ApAuthorizationListName
	respItem["local_authorization"] = flattenWirelessGetApAuthorizationListByIDV1ItemLocalAuthorization(item.LocalAuthorization)
	respItem["remote_authorization"] = flattenWirelessGetApAuthorizationListByIDV1ItemRemoteAuthorization(item.RemoteAuthorization)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetApAuthorizationListByIDV1ItemLocalAuthorization(item *catalystcentersdkgo.ResponseWirelessGetApAuthorizationListByIDV1ResponseLocalAuthorization) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ap_mac_entries"] = item.ApMacEntries
	respItem["ap_serial_number_entries"] = item.ApSerialNumberEntries

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetApAuthorizationListByIDV1ItemRemoteAuthorization(item *catalystcentersdkgo.ResponseWirelessGetApAuthorizationListByIDV1ResponseRemoteAuthorization) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aaa_servers"] = item.AAAServers
	respItem["authorize_ap_with_mac"] = boolPtrToString(item.AuthorizeApWithMac)
	respItem["authorize_ap_with_serial_number"] = boolPtrToString(item.AuthorizeApWithSerialNumber)

	return []map[string]interface{}{
		respItem,
	}

}
