package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisoriesSummary() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Security Advisories.

- Retrieves summary of advisories on the network.
`,

		ReadContext: dataSourceSecurityAdvisoriesSummaryRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cri_tic_al": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"hig_h": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"inf_orm_ati_ona_l": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"low": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"med_ium": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"na": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"con_fig": &schema.Schema{
										Description: `Number of advisories matched using default config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"cus_tom_con_fig": &schema.Schema{
										Description: `Number of advisories matched using user provided config
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"tot_al": &schema.Schema{
										Description: `Sum of Config, Custom Config and Version
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"ver_sio_n": &schema.Schema{
										Description: `Number of advisories matched using software version
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

func dataSourceSecurityAdvisoriesSummaryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdvisoriesSummaryV1")

		response1, restyResp1, err := client.SecurityAdvisories.GetAdvisoriesSummaryV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAdvisoriesSummaryV1", err,
				"Failure at GetAdvisoriesSummaryV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSecurityAdvisoriesGetAdvisoriesSummaryV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdvisoriesSummaryV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryV1Item(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["inf_orm_ati_ona_l"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemINFORMATIONAL(item.INFORMATIONAL)
	respItem["low"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemLOW(item.LOW)
	respItem["med_ium"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemMEDIUM(item.MEDIUM)
	respItem["hig_h"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemHIGH(item.HIGH)
	respItem["cri_tic_al"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemCRITICaL(item.CRITICaL)
	respItem["na"] = flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemNA(item.NA)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemINFORMATIONAL(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseINFORMATIONAL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemLOW(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseLOW) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemMEDIUM(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseMEDIUM) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemHIGH(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseHIGH) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemCRITICaL(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseCRITICaL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSecurityAdvisoriesGetAdvisoriesSummaryV1ItemNA(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesSummaryV1ResponseNA) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["con_fig"] = item.CONFIG
	respItem["cus_tom_con_fig"] = item.CUSTOMCONFIG
	respItem["ver_sio_n"] = item.VERSION
	respItem["tot_al"] = item.TOTAL

	return []map[string]interface{}{
		respItem,
	}

}
