package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSensorTestResults() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Intent API to get SENSOR test result summary
`,

		ReadContext: dataSourceWirelessSensorTestResultsRead,
		Schema: map[string]*schema.Schema{
			"end_time": &schema.Schema{
				Description: `endTime query parameter. The epoch time in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"site_id": &schema.Schema{
				Description: `siteId query parameter. Assurance site UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. The epoch time in milliseconds
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"test_failure_by": &schema.Schema{
				Description: `testFailureBy query parameter. Obtain failure statistics group by "area", "building", or "floor" (case insensitive)
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"failure_stats": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"error_code": &schema.Schema{
										Description: `The error code
`,
										Type:     schema.TypeInt,
										Computed: true,
									},

									"error_title": &schema.Schema{
										Description: `The error title
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"test_category": &schema.Schema{
										Description: `The test category
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"test_type": &schema.Schema{
										Description: `The test type
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"summary": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_con_nec_tiv_ity": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"fil_etr_ans_fer": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},
														},
													},
												},

												"hos_tre_ach_abi_lit_y": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"web_ser_ver": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
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

									"ema_il": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"mai_lse_rve_r": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},

									"net_wor_kse_rvi_ces": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dns": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
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

									"onb_oar_din_g": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ass_oc": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"aut_h": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"dhc_p": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
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

									"per_for_man_ce": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"ips_las_end_er": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
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

									"rfa_sse_ssm_ent": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dat_ara_te": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
`,
																Type:     schema.TypeInt,
																Computed: true,
															},
														},
													},
												},

												"snr": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"fail_count": &schema.Schema{
																Description: `Total failed test count
`,
																Type:     schema.TypeFloat,
																Computed: true,
															},

															"pass_count": &schema.Schema{
																Description: `Total passed test count
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

									"total_test_count": &schema.Schema{
										Description: `Total test count
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

func dataSourceWirelessSensorTestResultsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vSiteID, okSiteID := d.GetOk("site_id")
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vTestFailureBy, okTestFailureBy := d.GetOk("test_failure_by")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SensorTestResultsV1")
		queryParams1 := catalystcentersdkgo.SensorTestResultsV1QueryParams{}

		if okSiteID {
			queryParams1.SiteID = vSiteID.(string)
		}
		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okTestFailureBy {
			queryParams1.TestFailureBy = vTestFailureBy.(string)
		}

		response1, restyResp1, err := client.Wireless.SensorTestResultsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 SensorTestResultsV1", err,
				"Failure at SensorTestResultsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessSensorTestResultsV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SensorTestResultsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessSensorTestResultsV1Item(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["summary"] = flattenWirelessSensorTestResultsV1ItemSummary(item.Summary)
	respItem["failure_stats"] = flattenWirelessSensorTestResultsV1ItemFailureStats(item.FailureStats)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessSensorTestResultsV1ItemSummary(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummary) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["total_test_count"] = item.TotalTestCount
	respItem["onb_oar_din_g"] = flattenWirelessSensorTestResultsV1ItemSummaryOnBoarding(item.OnBoarding)
	respItem["per_for_man_ce"] = flattenWirelessSensorTestResultsV1ItemSummaryPERfORMAncE(item.PERfORMAncE)
	respItem["net_wor_kse_rvi_ces"] = flattenWirelessSensorTestResultsV1ItemSummaryNETWORKSERVICES(item.NETWORKSERVICES)
	respItem["app_con_nec_tiv_ity"] = flattenWirelessSensorTestResultsV1ItemSummaryApPCONNECTIVITY(item.ApPCONNECTIVITY)
	respItem["rfa_sse_ssm_ent"] = flattenWirelessSensorTestResultsV1ItemSummaryRfASSESSMENT(item.RfASSESSMENT)
	respItem["ema_il"] = flattenWirelessSensorTestResultsV1ItemSummaryEmail(item.Email)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryOnBoarding(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoarding) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aut_h"] = flattenWirelessSensorTestResultsV1ItemSummaryOnBoardingAuth(item.Auth)
	respItem["dhc_p"] = flattenWirelessSensorTestResultsV1ItemSummaryOnBoardingDHCP(item.DHCP)
	respItem["ass_oc"] = flattenWirelessSensorTestResultsV1ItemSummaryOnBoardingAssoc(item.Assoc)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryOnBoardingAuth(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingAuth) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryOnBoardingDHCP(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingDHCP) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryOnBoardingAssoc(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryOnBoardingAssoc) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryPERfORMAncE(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryPERfORMAncE) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["ips_las_end_er"] = flattenWirelessSensorTestResultsV1ItemSummaryPERfORMAncEIPSLASENDER(item.IPSLASENDER)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryPERfORMAncEIPSLASENDER(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryPERfORMAncEIPSLASENDER) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryNETWORKSERVICES(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryNETWORKSERVICES) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dns"] = flattenWirelessSensorTestResultsV1ItemSummaryNETWORKSERVICESDNS(item.DNS)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryNETWORKSERVICESDNS(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryNETWORKSERVICESDNS) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryApPCONNECTIVITY(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITY) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["hos_tre_ach_abi_lit_y"] = flattenWirelessSensorTestResultsV1ItemSummaryApPCONNECTIVITYHOSTREACHABILITY(item.HOSTREACHABILITY)
	respItem["web_ser_ver"] = flattenWirelessSensorTestResultsV1ItemSummaryApPCONNECTIVITYWebServer(item.WebServer)
	respItem["fil_etr_ans_fer"] = flattenWirelessSensorTestResultsV1ItemSummaryApPCONNECTIVITYFileTransfer(item.FileTransfer)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryApPCONNECTIVITYHOSTREACHABILITY(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYHOSTREACHABILITY) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryApPCONNECTIVITYWebServer(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYWebServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryApPCONNECTIVITYFileTransfer(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryApPCONNECTIVITYFileTransfer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryRfASSESSMENT(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENT) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["dat_ara_te"] = flattenWirelessSensorTestResultsV1ItemSummaryRfASSESSMENTDATARATE(item.DATARATE)
	respItem["snr"] = flattenWirelessSensorTestResultsV1ItemSummaryRfASSESSMENTSNR(item.SNR)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryRfASSESSMENTDATARATE(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENTDATARATE) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryRfASSESSMENTSNR(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryRfASSESSMENTSNR) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryEmail(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryEmail) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["mai_lse_rve_r"] = flattenWirelessSensorTestResultsV1ItemSummaryEmailMailServer(item.MailServer)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemSummaryEmailMailServer(item *catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseSummaryEmailMailServer) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["pass_count"] = item.PassCount
	respItem["fail_count"] = item.FailCount

	return []map[string]interface{}{
		respItem,
	}

}

func flattenWirelessSensorTestResultsV1ItemFailureStats(items *[]catalystcentersdkgo.ResponseWirelessSensorTestResultsV1ResponseFailureStats) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["error_code"] = item.ErrorCode
		respItem["error_title"] = item.ErrorTitle
		respItem["test_type"] = item.TestType
		respItem["test_category"] = item.TestCategory
		respItems = append(respItems, respItem)
	}
	return respItems
}
