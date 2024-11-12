package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseUsageDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Get count of purchased and in use Cisco DNA and Network licenses.
`,

		ReadContext: dataSourceLicenseUsageDetailsRead,
		Schema: map[string]*schema.Schema{
			"device_type": &schema.Schema{
				Description: `device_type query parameter. Type of device like router, switch, wireless or ise
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"smart_account_id": &schema.Schema{
				Description: `smart_account_id path parameter. Id of smart account
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"virtual_account_name": &schema.Schema{
				Description: `virtual_account_name path parameter. Name of virtual account. Putting "All" will give license term detail for all virtual accounts.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"purchased_dna_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"purchased_ise_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"purchased_network_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"used_dna_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"used_ise_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"used_network_license": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"license_count_by_type": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"license_count": &schema.Schema{
													Description: `Number of licenses
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"license_type": &schema.Schema{
													Description: `Type of license
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"total_license_count": &schema.Schema{
										Description: `Total number of licenses
`,
										Type:     schema.TypeInt,
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

func dataSourceLicenseUsageDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSmartAccountID := d.Get("smart_account_id")
	vVirtualAccountName := d.Get("virtual_account_name")
	vDeviceType := d.Get("device_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LicenseUsageDetailsV1")
		vvSmartAccountID := vSmartAccountID.(string)
		vvVirtualAccountName := vVirtualAccountName.(string)
		queryParams1 := catalystcentersdkgo.LicenseUsageDetailsV1QueryParams{}

		queryParams1.DeviceType = vDeviceType.(string)

		response1, restyResp1, err := client.Licenses.LicenseUsageDetailsV1(vvSmartAccountID, vvVirtualAccountName, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LicenseUsageDetailsV1", err,
				"Failure at LicenseUsageDetailsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesLicenseUsageDetailsV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LicenseUsageDetailsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesLicenseUsageDetailsV1Item(item *catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["purchased_dna_license"] = flattenLicensesLicenseUsageDetailsV1ItemPurchasedDnaLicense(item.PurchasedDnaLicense)
	respItem["purchased_network_license"] = flattenLicensesLicenseUsageDetailsV1ItemPurchasedNetworkLicense(item.PurchasedNetworkLicense)
	respItem["used_dna_license"] = flattenLicensesLicenseUsageDetailsV1ItemUsedDnaLicense(item.UsedDnaLicense)
	respItem["used_network_license"] = flattenLicensesLicenseUsageDetailsV1ItemUsedNetworkLicense(item.UsedNetworkLicense)
	respItem["purchased_ise_license"] = flattenLicensesLicenseUsageDetailsV1ItemPurchasedIseLicense(item.PurchasedIseLicense)
	respItem["used_ise_license"] = flattenLicensesLicenseUsageDetailsV1ItemUsedIseLicense(item.UsedIseLicense)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenLicensesLicenseUsageDetailsV1ItemPurchasedDnaLicense(item *catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1PurchasedDnaLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsV1ItemPurchasedDnaLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsV1ItemPurchasedDnaLicenseLicenseCountByType(items *[]catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1PurchasedDnaLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesLicenseUsageDetailsV1ItemPurchasedNetworkLicense(item *catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1PurchasedNetworkLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsV1ItemPurchasedNetworkLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsV1ItemPurchasedNetworkLicenseLicenseCountByType(items *[]catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1PurchasedNetworkLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesLicenseUsageDetailsV1ItemUsedDnaLicense(item *catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1UsedDnaLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsV1ItemUsedDnaLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsV1ItemUsedDnaLicenseLicenseCountByType(items *[]catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1UsedDnaLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesLicenseUsageDetailsV1ItemUsedNetworkLicense(item *catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1UsedNetworkLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsV1ItemUsedNetworkLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsV1ItemUsedNetworkLicenseLicenseCountByType(items *[]catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1UsedNetworkLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesLicenseUsageDetailsV1ItemPurchasedIseLicense(item *catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1PurchasedIseLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsV1ItemPurchasedIseLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsV1ItemPurchasedIseLicenseLicenseCountByType(items *[]catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1PurchasedIseLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLicensesLicenseUsageDetailsV1ItemUsedIseLicense(item *catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1UsedIseLicense) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_license_count"] = item.TotalLicenseCount
	respItem["license_count_by_type"] = flattenLicensesLicenseUsageDetailsV1ItemUsedIseLicenseLicenseCountByType(item.LicenseCountByType)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesLicenseUsageDetailsV1ItemUsedIseLicenseLicenseCountByType(items *[]catalystcentersdkgo.ResponseLicensesLicenseUsageDetailsV1UsedIseLicenseLicenseCountByType) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["license_type"] = item.LicenseType
		respItem["license_count"] = item.LicenseCount
		respItems = append(respItems, respItem)
	}
	return respItems
}
