package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessProfile() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Gets either one or all the wireless network profiles if no name is provided for network-profile.
`,

		ReadContext: dataSourceWirelessProfileRead,
		Schema: map[string]*schema.Schema{
			"profile_name": &schema.Schema{
				Description: `profileName query parameter. Wireless Network Profile Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"profile_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": &schema.Schema{
										Description: `Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sites": &schema.Schema{
										Description: `array of site name hierarchies(eg: ["Global/aaa/zzz", "Global/aaa/zzz"])
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"ssid_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"enable_fabric": &schema.Schema{
													Description: `true if fabric is enabled else false
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"flex_connect": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"enable_flex_connect": &schema.Schema{
																Description: `true if flex connect is enabled else false
`,
																// Type:        schema.TypeBool,
																Type:     schema.TypeString,
																Computed: true,
															},

															"local_to_vlan": &schema.Schema{
																Description: `Local To VLAN ID
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"interface_name": &schema.Schema{
													Description: `Interface Name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `SSID Name
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `SSID Type(enum: Enterprise/Guest)
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
				},
			},
		},
	}
}

func dataSourceWirelessProfileRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vProfileName, okProfileName := d.GetOk("profile_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetWirelessProfileV1")
		queryParams1 := catalystcentersdkgo.GetWirelessProfileV1QueryParams{}

		if okProfileName {
			queryParams1.ProfileName = vProfileName.(string)
		}

		response1, restyResp1, err := client.Wireless.GetWirelessProfileV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetWirelessProfileV1", err,
				"Failure at GetWirelessProfileV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetWirelessProfileV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessProfileV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetWirelessProfileV1Items(items *catalystcentersdkgo.ResponseWirelessGetWirelessProfileV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["profile_details"] = flattenWirelessGetWirelessProfileV1ItemsProfileDetails(item.ProfileDetails)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfileV1ItemsProfileDetails(item *catalystcentersdkgo.ResponseItemWirelessGetWirelessProfileV1ProfileDetails) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["sites"] = item.Sites
	respItem["ssid_details"] = flattenWirelessGetWirelessProfileV1ItemsProfileDetailsSSIDDetails(item.SSIDDetails)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessGetWirelessProfileV1ItemsProfileDetailsSSIDDetails(items *[]catalystcentersdkgo.ResponseItemWirelessGetWirelessProfileV1ProfileDetailsSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["type"] = item.Type
		respItem["enable_fabric"] = boolPtrToString(item.EnableFabric)
		respItem["flex_connect"] = flattenWirelessGetWirelessProfileV1ItemsProfileDetailsSSIDDetailsFlexConnect(item.FlexConnect)
		respItem["interface_name"] = item.InterfaceName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfileV1ItemsProfileDetailsSSIDDetailsFlexConnect(item *catalystcentersdkgo.ResponseItemWirelessGetWirelessProfileV1ProfileDetailsSSIDDetailsFlexConnect) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["enable_flex_connect"] = boolPtrToString(item.EnableFlexConnect)
	respItem["local_to_vlan"] = item.LocalToVLAN

	return []map[string]interface{}{
		respItem,
	}

}
