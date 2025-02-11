package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessProfiles() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get all Wireless Network Profiles

- This data source allows the user to get a Wireless Network Profile by ID
`,

		ReadContext: dataSourceWirelessProfilesRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Wireless Profile Id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Default is 500 if not specified. Maximum allowed limit is 500
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
			"wireless_profile_name": &schema.Schema{
				Description: `wirelessProfileName query parameter. Wireless Profile Name
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_interfaces": &schema.Schema{
							Description: `Additional Interfaces
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"ap_zones": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_zone_name": &schema.Schema{
										Description: `AP Zone Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rf_profile_name": &schema.Schema{
										Description: `RF Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssids": &schema.Schema{
										Description: `ssids part of apZone
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

						"id": &schema.Schema{
							Description: `Wireless Profile Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"anchor_group_name": &schema.Schema{
										Description: `Anchor Group Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dot11be_profile_id": &schema.Schema{
										Description: `802.11be Profile ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_fabric": &schema.Schema{
										Description: `True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
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
													Description: `True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"local_to_vlan": &schema.Schema{
													Description: `Local to VLAN ID
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

									"policy_profile_name": &schema.Schema{
										Description: `Policy Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssid_name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"vlan_group_name": &schema.Schema{
										Description: `VLAN Group Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"wlan_profile_name": &schema.Schema{
										Description: `WLAN Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"wireless_profile_name": &schema.Schema{
							Description: `Wireless Profile Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"additional_interfaces": &schema.Schema{
							Description: `Additional Interfaces
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"ap_zones": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_zone_name": &schema.Schema{
										Description: `AP Zone Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"rf_profile_name": &schema.Schema{
										Description: `RF Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssids": &schema.Schema{
										Description: `ssids part of apZone
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

						"id": &schema.Schema{
							Description: `Wireless Profile Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ssid_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"anchor_group_name": &schema.Schema{
										Description: `Anchor Group Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dot11be_profile_id": &schema.Schema{
										Description: `802.11be Profile ID
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"enable_fabric": &schema.Schema{
										Description: `True if fabric is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
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
													Description: `True if flex connect is enabled, else False. Flex and fabric cannot be enabled simultaneously and a profile can only contain either flex SSIDs or fabric SSIDs and not both at the same time
`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
												},

												"local_to_vlan": &schema.Schema{
													Description: `Local to VLAN ID
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

									"policy_profile_name": &schema.Schema{
										Description: `Policy Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"ssid_name": &schema.Schema{
										Description: `SSID Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"vlan_group_name": &schema.Schema{
										Description: `VLAN Group Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"wlan_profile_name": &schema.Schema{
										Description: `WLAN Profile Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"wireless_profile_name": &schema.Schema{
							Description: `Wireless Profile Name
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

func dataSourceWirelessProfilesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vWirelessProfileName, okWirelessProfileName := d.GetOk("wireless_profile_name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okLimit, okOffset, okWirelessProfileName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetWirelessProfilesV1")
		queryParams1 := catalystcentersdkgo.GetWirelessProfilesV1QueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okWirelessProfileName {
			queryParams1.WirelessProfileName = vWirelessProfileName.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.GetWirelessProfilesV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetWirelessProfilesV1", err,
				"Failure at GetWirelessProfilesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetWirelessProfilesV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessProfilesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetWirelessProfileByIDV1")
		vvID := vID.(string)

		// has_unknown_response: None

		response2, restyResp2, err := client.Wireless.GetWirelessProfileByIDV1(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetWirelessProfileByIDV1", err,
				"Failure at GetWirelessProfileByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenWirelessGetWirelessProfileByIDV1Item(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWirelessProfileByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetWirelessProfilesV1Items(items *[]catalystcentersdkgo.ResponseWirelessGetWirelessProfilesV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["wireless_profile_name"] = item.WirelessProfileName
		respItem["ssid_details"] = flattenWirelessGetWirelessProfilesV1ItemsSSIDDetails(item.SSIDDetails)
		respItem["id"] = item.ID
		respItem["additional_interfaces"] = item.AdditionalInterfaces
		respItem["ap_zones"] = flattenWirelessGetWirelessProfilesV1ItemsApZones(item.ApZones)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfilesV1ItemsSSIDDetails(items *[]catalystcentersdkgo.ResponseWirelessGetWirelessProfilesV1ResponseSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ssid_name"] = item.SSIDName
		respItem["flex_connect"] = flattenWirelessGetWirelessProfilesV1ItemsSSIDDetailsFlexConnect(item.FlexConnect)
		respItem["enable_fabric"] = boolPtrToString(item.EnableFabric)
		respItem["wlan_profile_name"] = item.WLANProfileName
		respItem["interface_name"] = item.InterfaceName
		respItem["policy_profile_name"] = item.PolicyProfileName
		respItem["dot11be_profile_id"] = item.Dot11BeProfileID
		respItem["anchor_group_name"] = item.AnchorGroupName
		respItem["vlan_group_name"] = item.VLANGroupName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfilesV1ItemsSSIDDetailsFlexConnect(item *catalystcentersdkgo.ResponseWirelessGetWirelessProfilesV1ResponseSSIDDetailsFlexConnect) []map[string]interface{} {
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

func flattenWirelessGetWirelessProfilesV1ItemsApZones(items *[]catalystcentersdkgo.ResponseWirelessGetWirelessProfilesV1ResponseApZones) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ap_zone_name"] = item.ApZoneName
		respItem["rf_profile_name"] = item.RfProfileName
		respItem["ssids"] = item.SSIDs
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfileByIDV1Item(item *catalystcentersdkgo.ResponseWirelessGetWirelessProfileByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["wireless_profile_name"] = item.WirelessProfileName
	respItem["ssid_details"] = flattenWirelessGetWirelessProfileByIDV1ItemSSIDDetails(item.SSIDDetails)
	respItem["id"] = item.ID
	respItem["additional_interfaces"] = item.AdditionalInterfaces
	respItem["ap_zones"] = flattenWirelessGetWirelessProfileByIDV1ItemApZones(item.ApZones)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetWirelessProfileByIDV1ItemSSIDDetails(items *[]catalystcentersdkgo.ResponseWirelessGetWirelessProfileByIDV1ResponseSSIDDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ssid_name"] = item.SSIDName
		respItem["flex_connect"] = flattenWirelessGetWirelessProfileByIDV1ItemSSIDDetailsFlexConnect(item.FlexConnect)
		respItem["enable_fabric"] = boolPtrToString(item.EnableFabric)
		respItem["wlan_profile_name"] = item.WLANProfileName
		respItem["interface_name"] = item.InterfaceName
		respItem["policy_profile_name"] = item.PolicyProfileName
		respItem["dot11be_profile_id"] = item.Dot11BeProfileID
		respItem["anchor_group_name"] = item.AnchorGroupName
		respItem["vlan_group_name"] = item.VLANGroupName
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetWirelessProfileByIDV1ItemSSIDDetailsFlexConnect(item *catalystcentersdkgo.ResponseWirelessGetWirelessProfileByIDV1ResponseSSIDDetailsFlexConnect) []map[string]interface{} {
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

func flattenWirelessGetWirelessProfileByIDV1ItemApZones(items *[]catalystcentersdkgo.ResponseWirelessGetWirelessProfileByIDV1ResponseApZones) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ap_zone_name"] = item.ApZoneName
		respItem["rf_profile_name"] = item.RfProfileName
		respItem["ssids"] = item.SSIDs
		respItems = append(respItems, respItem)
	}
	return respItems
}
