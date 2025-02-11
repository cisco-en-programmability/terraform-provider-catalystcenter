package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFabricSiteHealthSummaries() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get a paginated list of Fabric sites Networks with health summary.
This data source provides the latest health data until the given endTime. If data is not ready for the provided
endTime, the request will fail with error code 400 Bad Request, and the error message will indicate the recommended
endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
are not a real time system. When endTime is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
fabricSiteHealthSummaries-1.0.1-resolved.yaml
`,

		ReadContext: dataSourceFabricSiteHealthSummariesRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. The list of FabricSite health attributes. Please refer to fabricSiteAttributes section in the Open API specification document mentioned in the description.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_time": &schema.Schema{
				Description: `endTime query parameter. End time to which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. The list of entity Uuids. (Ex."6bef213c-19ca-4170-8375-b694e251101c") Examples: id=6bef213c-19ca-4170-8375-b694e251101c (single entity uuid requested) id=6bef213c-19ca-4170-8375-b694e251101c&id=32219612-819e-4b5e-a96b-cf22aca13dd9&id=2541e9a7-b80d-4955-8aa2-79b233318ba0 (multiple entity uuid with '&' separator)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. The sort order of the field ascending or descending.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"start_time": &schema.Schema{
				Description: `startTime query parameter. Start time from which API queries the data set related to the resource. It must be specified in UNIX epochtime in milliseconds. Value is inclusive.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"view": &schema.Schema{
				Description: `view query parameter. The specific summary view being requested. A maximum of 3 views can be queried at a time per request.  Please refer to fabricSiteViews section in the Open API specification document mentioned in the description.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID is used to trace the origin of API calls and their associated queries executed on the database. It's an optional header parameter that can be added to an API request.
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aaa_status_fair_health_device_count": &schema.Schema{
							Description: `Aaa Status Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"aaa_status_good_health_device_count": &schema.Schema{
							Description: `Aaa Status Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"aaa_status_good_health_percentage": &schema.Schema{
							Description: `Aaa Status Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"aaa_status_poor_health_device_count": &schema.Schema{
							Description: `Aaa Status Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"aaa_status_total_health_device_count": &schema.Schema{
							Description: `Aaa Status Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"associated_l2_vn_count": &schema.Schema{
							Description: `Associated L2 Vn Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"associated_l3_vn_count": &schema.Schema{
							Description: `Associated L3 Vn Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_bgp_site_fair_health_device_count": &schema.Schema{
							Description: `Bgp Bgp Site Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_bgp_site_good_health_device_count": &schema.Schema{
							Description: `Bgp Bgp Site Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_bgp_site_good_health_percentage": &schema.Schema{
							Description: `Bgp Bgp Site Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_bgp_site_poor_health_device_count": &schema.Schema{
							Description: `Bgp Bgp Site Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_bgp_site_total_health_device_count": &schema.Schema{
							Description: `Bgp Bgp Site Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_evpn_fair_health_device_count": &schema.Schema{
							Description: `Bgp Evpn Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_evpn_good_health_device_count": &schema.Schema{
							Description: `Bgp Evpn Good Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_evpn_good_health_percentage": &schema.Schema{
							Description: `Bgp Evpn Good Health Percentage`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"bgp_evpnpoor_health_device_count": &schema.Schema{
							Description: `Bgp Evpn Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_evpn_total_health_device_count": &schema.Schema{
							Description: `Bgp Evpn Total Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_peer_infra_vn_fair_health_device_count": &schema.Schema{
							Description: `Bgp Peer Infra Vn Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_peer_infra_vn_good_health_device_count": &schema.Schema{
							Description: `Bgp Peer Infra Vn Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_peer_infra_vn_poor_health_device_count": &schema.Schema{
							Description: `Bgp Peer Infra Vn Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_peer_infra_vn_score_good_health_percentage": &schema.Schema{
							Description: `Bgp Peer Infra Vn Score Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_peer_infra_vn_total_health_device_count": &schema.Schema{
							Description: `Bgp Peer Infra Vn Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_pubsub_site_fair_health_device_count": &schema.Schema{
							Description: `Bgp Pubsub Site Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_pubsub_site_good_health_device_count": &schema.Schema{
							Description: `Bgp Pubsub Site Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_pubsub_site_good_health_percentage": &schema.Schema{
							Description: `Bgp Pubsub Site Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_pubsub_site_poor_health_device_count": &schema.Schema{
							Description: `Bgp Pubsub Site Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_pubsub_site_total_health_device_count": &schema.Schema{
							Description: `Bgp Pubsub Site Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"border_to_control_plane_fair_health_device_count": &schema.Schema{
							Description: `Border To Control Plane Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"border_to_control_plane_good_health_device_count": &schema.Schema{
							Description: `Border To Control Plane Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"border_to_control_plane_good_health_percentage": &schema.Schema{
							Description: `Border To Control Plane Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"border_to_control_plane_poor_health_device_count": &schema.Schema{
							Description: `Border To Control Plane Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"border_to_control_plane_total_health_device_count": &schema.Schema{
							Description: `Border To Control Plane Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connectivity_fair_health_device_count": &schema.Schema{
							Description: `Connectivity Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connectivity_good_health_device_count": &schema.Schema{
							Description: `Connectivity Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connectivity_good_health_percentage": &schema.Schema{
							Description: `Connectivity Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connectivity_poor_health_device_count": &schema.Schema{
							Description: `Connectivity Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"connectivity_total_health_device_count": &schema.Schema{
							Description: `Connectivity Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"control_plane_fair_health_device_count": &schema.Schema{
							Description: `Control Plane Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"control_plane_good_health_device_count": &schema.Schema{
							Description: `Control Plane Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"control_plane_good_health_percentage": &schema.Schema{
							Description: `Control Plane Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"control_plane_poor_health_device_count": &schema.Schema{
							Description: `Control Plane Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"control_plane_total_health_device_count": &schema.Schema{
							Description: `Control Plane Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"cts_env_data_download_fair_health_device_count": &schema.Schema{
							Description: `Cts Env Data Download Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"cts_env_data_download_good_health_device_count": &schema.Schema{
							Description: `Cts Env Data Download Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"cts_env_data_download_good_health_percentage": &schema.Schema{
							Description: `Cts Env Data Download Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"cts_env_data_download_poor_health_device_count": &schema.Schema{
							Description: `Cts Env Data Download Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"cts_env_data_download_total_health_device_count": &schema.Schema{
							Description: `Cts Env Data Download Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"fair_health_device_count": &schema.Schema{
							Description: `Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"good_health_device_count": &schema.Schema{
							Description: `Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"good_health_percentage": &schema.Schema{
							Description: `Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"infra_fair_health_device_count": &schema.Schema{
							Description: `Infra Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"infra_good_health_device_count": &schema.Schema{
							Description: `Infra Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"infra_good_health_percentage": &schema.Schema{
							Description: `Infra Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"infra_poor_health_device_count": &schema.Schema{
							Description: `Infra Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"infra_total_health_device_count": &schema.Schema{
							Description: `Infra Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"lisp_session_fair_health_device_count": &schema.Schema{
							Description: `Lisp Session Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"lisp_session_good_health_device_count": &schema.Schema{
							Description: `Lisp Session Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"lisp_session_good_health_percentage": &schema.Schema{
							Description: `Lisp Session Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"lisp_session_poor_health_device_count": &schema.Schema{
							Description: `Lisp Session Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"lisp_session_total_health_device_count": &schema.Schema{
							Description: `Lisp Session Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"network_protocol": &schema.Schema{
							Description: `Network Protocol`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"peer_score_fair_health_device_count": &schema.Schema{
							Description: `Peer Score Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"peer_score_good_health_device_count": &schema.Schema{
							Description: `Peer Score Good Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"peer_score_good_health_percentage": &schema.Schema{
							Description: `Peer Score Good Health Percentage`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"peer_score_poor_health_device_count": &schema.Schema{
							Description: `Peer Score Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"peer_score_total_health_device_count": &schema.Schema{
							Description: `Peer Score Total Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"poor_health_device_count": &schema.Schema{
							Description: `Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"port_channel_fair_health_device_count": &schema.Schema{
							Description: `Port Channel Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"port_channel_good_health_device_count": &schema.Schema{
							Description: `Port Channel Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"port_channel_good_health_percentage": &schema.Schema{
							Description: `Port Channel Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"port_channel_poor_health_device_count": &schema.Schema{
							Description: `Port Channel Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"port_channel_total_health_device_count": &schema.Schema{
							Description: `Port Channel Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_infra_vn_fair_health_device_count": &schema.Schema{
							Description: `Pubsub Infra Vn Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_infra_vn_good_health_device_count": &schema.Schema{
							Description: `Pubsub Infra Vn Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_infra_vn_good_health_percentage": &schema.Schema{
							Description: `Pubsub Infra Vn Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_infra_vn_poor_health_device_count": &schema.Schema{
							Description: `Pubsub Infra Vn Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_infra_vn_total_health_device_count": &schema.Schema{
							Description: `Pubsub Infra Vn Total Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"total_device_count": &schema.Schema{
							Description: `Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFabricSiteHealthSummariesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vID, okID := d.GetOk("id")
	vAttribute, okAttribute := d.GetOk("attribute")
	vView, okView := d.GetOk("view")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReadListOfFabricSitesWithTheirHealthSummaryV1")

		headerParams1 := catalystcentersdkgo.ReadListOfFabricSitesWithTheirHealthSummaryV1HeaderParams{}
		queryParams1 := catalystcentersdkgo.ReadListOfFabricSitesWithTheirHealthSummaryV1QueryParams{}

		if okStartTime {
			queryParams1.StartTime = vStartTime.(float64)
		}
		if okEndTime {
			queryParams1.EndTime = vEndTime.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okID {
			queryParams1.ID = vID.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.ReadListOfFabricSitesWithTheirHealthSummaryV1(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReadListOfFabricSitesWithTheirHealthSummaryV1", err,
				"Failure at ReadListOfFabricSitesWithTheirHealthSummaryV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaReadListOfFabricSitesWithTheirHealthSummaryV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReadListOfFabricSitesWithTheirHealthSummaryV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaReadListOfFabricSitesWithTheirHealthSummaryV1Items(items *[]catalystcentersdkgo.ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["good_health_percentage"] = item.GoodHealthPercentage
		respItem["good_health_device_count"] = item.GoodHealthDeviceCount
		respItem["total_device_count"] = item.TotalDeviceCount
		respItem["poor_health_device_count"] = item.PoorHealthDeviceCount
		respItem["fair_health_device_count"] = item.FairHealthDeviceCount
		respItem["associated_l2_vn_count"] = item.AssociatedL2VnCount
		respItem["associated_l3_vn_count"] = item.AssociatedL3VnCount
		respItem["network_protocol"] = item.NetworkProtocol
		respItem["connectivity_good_health_percentage"] = item.ConnectivityGoodHealthPercentage
		respItem["connectivity_total_health_device_count"] = item.ConnectivityTotalHealthDeviceCount
		respItem["connectivity_good_health_device_count"] = item.ConnectivityGoodHealthDeviceCount
		respItem["connectivity_poor_health_device_count"] = item.ConnectivityPoorHealthDeviceCount
		respItem["connectivity_fair_health_device_count"] = item.ConnectivityFairHealthDeviceCount
		respItem["infra_good_health_percentage"] = item.InfraGoodHealthPercentage
		respItem["infra_total_health_device_count"] = item.InfraTotalHealthDeviceCount
		respItem["infra_good_health_device_count"] = item.InfraGoodHealthDeviceCount
		respItem["infra_fair_health_device_count"] = item.InfraFairHealthDeviceCount
		respItem["infra_poor_health_device_count"] = item.InfraPoorHealthDeviceCount
		respItem["control_plane_good_health_percentage"] = item.ControlPlaneGoodHealthPercentage
		respItem["control_plane_total_health_device_count"] = item.ControlPlaneTotalHealthDeviceCount
		respItem["control_plane_good_health_device_count"] = item.ControlPlaneGoodHealthDeviceCount
		respItem["control_plane_poor_health_device_count"] = item.ControlPlanePoorHealthDeviceCount
		respItem["control_plane_fair_health_device_count"] = item.ControlPlaneFairHealthDeviceCount
		respItem["pubsub_infra_vn_good_health_percentage"] = item.PubsubInfraVnGoodHealthPercentage
		respItem["pubsub_infra_vn_total_health_device_count"] = item.PubsubInfraVnTotalHealthDeviceCount
		respItem["pubsub_infra_vn_good_health_device_count"] = item.PubsubInfraVnGoodHealthDeviceCount
		respItem["pubsub_infra_vn_poor_health_device_count"] = item.PubsubInfraVnPoorHealthDeviceCount
		respItem["pubsub_infra_vn_fair_health_device_count"] = item.PubsubInfraVnFairHealthDeviceCount
		respItem["bgp_evpn_good_health_percentage"] = flattenSdaReadListOfFabricSitesWithTheirHealthSummaryV1ItemsBgpEvpnGoodHealthPercentage(item.BgpEvpnGoodHealthPercentage)
		respItem["bgp_evpn_total_health_device_count"] = item.BgpEvpnTotalHealthDeviceCount
		respItem["bgp_evpn_good_health_device_count"] = item.BgpEvpnGoodHealthDeviceCount
		respItem["bgp_evpnpoor_health_device_count"] = item.BgpEvpnpoorHealthDeviceCount
		respItem["bgp_evpn_fair_health_device_count"] = item.BgpEvpnFairHealthDeviceCount
		respItem["cts_env_data_download_good_health_percentage"] = item.CtsEnvDataDownloadGoodHealthPercentage
		respItem["cts_env_data_download_total_health_device_count"] = item.CtsEnvDataDownloadTotalHealthDeviceCount
		respItem["cts_env_data_download_good_health_device_count"] = item.CtsEnvDataDownloadGoodHealthDeviceCount
		respItem["cts_env_data_download_poor_health_device_count"] = item.CtsEnvDataDownloadPoorHealthDeviceCount
		respItem["cts_env_data_download_fair_health_device_count"] = item.CtsEnvDataDownloadFairHealthDeviceCount
		respItem["aaa_status_good_health_percentage"] = item.AAAStatusGoodHealthPercentage
		respItem["aaa_status_total_health_device_count"] = item.AAAStatusTotalHealthDeviceCount
		respItem["aaa_status_good_health_device_count"] = item.AAAStatusGoodHealthDeviceCount
		respItem["aaa_status_poor_health_device_count"] = item.AAAStatusPoorHealthDeviceCount
		respItem["aaa_status_fair_health_device_count"] = item.AAAStatusFairHealthDeviceCount
		respItem["port_channel_good_health_percentage"] = item.PortChannelGoodHealthPercentage
		respItem["port_channel_total_health_device_count"] = item.PortChannelTotalHealthDeviceCount
		respItem["port_channel_good_health_device_count"] = item.PortChannelGoodHealthDeviceCount
		respItem["port_channel_poor_health_device_count"] = item.PortChannelPoorHealthDeviceCount
		respItem["port_channel_fair_health_device_count"] = item.PortChannelFairHealthDeviceCount
		respItem["peer_score_good_health_percentage"] = flattenSdaReadListOfFabricSitesWithTheirHealthSummaryV1ItemsPeerScoreGoodHealthPercentage(item.PeerScoreGoodHealthPercentage)
		respItem["peer_score_total_health_device_count"] = item.PeerScoreTotalHealthDeviceCount
		respItem["peer_score_good_health_device_count"] = item.PeerScoreGoodHealthDeviceCount
		respItem["peer_score_poor_health_device_count"] = item.PeerScorePoorHealthDeviceCount
		respItem["peer_score_fair_health_device_count"] = item.PeerScoreFairHealthDeviceCount
		respItem["lisp_session_good_health_percentage"] = item.LispSessionGoodHealthPercentage
		respItem["lisp_session_total_health_device_count"] = item.LispSessionTotalHealthDeviceCount
		respItem["lisp_session_good_health_device_count"] = item.LispSessionGoodHealthDeviceCount
		respItem["lisp_session_poor_health_device_count"] = item.LispSessionPoorHealthDeviceCount
		respItem["lisp_session_fair_health_device_count"] = item.LispSessionFairHealthDeviceCount
		respItem["border_to_control_plane_good_health_percentage"] = item.BorderToControlPlaneGoodHealthPercentage
		respItem["border_to_control_plane_total_health_device_count"] = item.BorderToControlPlaneTotalHealthDeviceCount
		respItem["border_to_control_plane_good_health_device_count"] = item.BorderToControlPlaneGoodHealthDeviceCount
		respItem["border_to_control_plane_poor_health_device_count"] = item.BorderToControlPlanePoorHealthDeviceCount
		respItem["border_to_control_plane_fair_health_device_count"] = item.BorderToControlPlaneFairHealthDeviceCount
		respItem["bgp_bgp_site_good_health_percentage"] = item.BgpBgpSiteGoodHealthPercentage
		respItem["bgp_bgp_site_total_health_device_count"] = item.BgpBgpSiteTotalHealthDeviceCount
		respItem["bgp_bgp_site_good_health_device_count"] = item.BgpBgpSiteGoodHealthDeviceCount
		respItem["bgp_bgp_site_poor_health_device_count"] = item.BgpBgpSitePoorHealthDeviceCount
		respItem["bgp_bgp_site_fair_health_device_count"] = item.BgpBgpSiteFairHealthDeviceCount
		respItem["bgp_pubsub_site_good_health_percentage"] = item.BgpPubsubSiteGoodHealthPercentage
		respItem["bgp_pubsub_site_total_health_device_count"] = item.BgpPubsubSiteTotalHealthDeviceCount
		respItem["bgp_pubsub_site_good_health_device_count"] = item.BgpPubsubSiteGoodHealthDeviceCount
		respItem["bgp_pubsub_site_poor_health_device_count"] = item.BgpPubsubSitePoorHealthDeviceCount
		respItem["bgp_pubsub_site_fair_health_device_count"] = item.BgpPubsubSiteFairHealthDeviceCount
		respItem["bgp_peer_infra_vn_score_good_health_percentage"] = item.BgpPeerInfraVnScoreGoodHealthPercentage
		respItem["bgp_peer_infra_vn_total_health_device_count"] = item.BgpPeerInfraVnTotalHealthDeviceCount
		respItem["bgp_peer_infra_vn_good_health_device_count"] = item.BgpPeerInfraVnGoodHealthDeviceCount
		respItem["bgp_peer_infra_vn_poor_health_device_count"] = item.BgpPeerInfraVnPoorHealthDeviceCount
		respItem["bgp_peer_infra_vn_fair_health_device_count"] = item.BgpPeerInfraVnFairHealthDeviceCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaReadListOfFabricSitesWithTheirHealthSummaryV1ItemsBgpEvpnGoodHealthPercentage(item *catalystcentersdkgo.ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryV1ResponseBgpEvpnGoodHealthPercentage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSdaReadListOfFabricSitesWithTheirHealthSummaryV1ItemsPeerScoreGoodHealthPercentage(item *catalystcentersdkgo.ResponseSdaReadListOfFabricSitesWithTheirHealthSummaryV1ResponsePeerScoreGoodHealthPercentage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
