package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceVirtualNetworkHealthSummaries() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get a paginated list of Virtual Networks with health summary. Layer 2 Virtual Networks are only included in health
reporting for EVPN protocol deployments. The special Layer 3 VN called ‘INFRA_VN’ is also not included for user access
through Assurance virtualNetworkHealthSummaries APIS. Please find INFRA_VN related health metrics under
/data/api/v1/fabricSiteHealthSummaries (Ex: attributes ‘pubsubInfraVnGoodHealthPercentage’ and
‘bgpPeerInfraVnScoreGoodHealthPercentage’).
This data source provides the latest health data until the given endTime. If data is not ready for the provided
endTime, the request will fail with error code 400 Bad Request, and the error message will indicate the recommended
endTime to use to retrieve a complete data set. This behavior may occur if the provided endTime=currentTime, since we
are not a real time system. When endTime is not provided, the API returns the latest data.
For detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
virtualNetworkHealthSummaries-1.0.1-resolved.yaml
`,

		ReadContext: dataSourceVirtualNetworkHealthSummariesRead,
		Schema: map[string]*schema.Schema{
			"attribute": &schema.Schema{
				Description: `attribute query parameter. The interested fields in the request. For valid attributes, verify the documentation.
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
				Description: `view query parameter. The specific summary view being requested. This is an optional parameter which can be passed to get one or more of the specific health data summaries associated with virtual networks.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"vn_layer": &schema.Schema{
				Description: `vnLayer query parameter. VN Layer information covering Layer 3 or Layer 2 VNs.
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

						"associated_l3_vn": &schema.Schema{
							Description: `Associated L3 Vn`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"bgp_peer_fair_health_device_count": &schema.Schema{
							Description: `Bgp Peer Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_peer_good_health_device_count": &schema.Schema{
							Description: `Bgp Peer Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_peer_good_health_percentage": &schema.Schema{
							Description: `Bgp Peer Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_peer_no_health_device_count": &schema.Schema{
							Description: `Bgp Peer No Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"bgp_peer_poor_health_device_count": &schema.Schema{
							Description: `Bgp Peer Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"bgp_peer_total_device_count": &schema.Schema{
							Description: `Bgp Peer Total Device Count`,
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

						"internet_avail_fair_health_device_count": &schema.Schema{
							Description: `Internet Avail Fair Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_good_health_device_count": &schema.Schema{
							Description: `Internet Avail Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_good_health_percentage": &schema.Schema{
							Description: `Internet Avail Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_no_health_device_count": &schema.Schema{
							Description: `Internet Avail No Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_poor_health_device_count": &schema.Schema{
							Description: `Internet Avail Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"internet_avail_total_device_count": &schema.Schema{
							Description: `Internet Avail Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"layer": &schema.Schema{
							Description: `Layer`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"multi_cast_fair_health_device_count": &schema.Schema{
							Description: `Multi Cast Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"multi_cast_good_health_device_count": &schema.Schema{
							Description: `Multi Cast Good Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"multi_cast_good_health_percentage": &schema.Schema{
							Description: `Multi Cast Good Health Percentage`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"multi_cast_no_health_device_count": &schema.Schema{
							Description: `Multi Cast No Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"multi_cast_poor_health_device_count": &schema.Schema{
							Description: `Multi Cast Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"multi_cast_total_device_count": &schema.Schema{
							Description: `Multi Cast Total Device Count`,
							Type:        schema.TypeFloat,
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

						"no_health_device_count": &schema.Schema{
							Description: `No Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"poor_health_device_count": &schema.Schema{
							Description: `Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_session_fair_health_device_count": &schema.Schema{
							Description: `Pubsub Session Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"pubsub_session_good_health_device_count": &schema.Schema{
							Description: `Pubsub Session Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_session_good_health_percentage": &schema.Schema{
							Description: `Pubsub Session Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_session_no_health_device_count": &schema.Schema{
							Description: `Pubsub Session No Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"pubsub_session_poor_health_device_count": &schema.Schema{
							Description: `Pubsub Session Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"pubsub_session_total_device_count": &schema.Schema{
							Description: `Pubsub Session Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"total_device_count": &schema.Schema{
							Description: `Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"total_endpoints": &schema.Schema{
							Description: `Total Endpoints`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"total_fabric_sites": &schema.Schema{
							Description: `Total Fabric Sites`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vlan": &schema.Schema{
							Description: `Vlan`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"vn_exit_fair_health_device_count": &schema.Schema{
							Description: `Vn Exit Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_exit_good_health_device_count": &schema.Schema{
							Description: `Vn Exit Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_exit_health_percentage": &schema.Schema{
							Description: `Vn Exit Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_exit_no_health_device_count": &schema.Schema{
							Description: `Vn Exit No Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_exit_poor_health_device_count": &schema.Schema{
							Description: `Vn Exit Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_exit_total_device_count": &schema.Schema{
							Description: `Vn Exit Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_fabric_control_plane_fair_health_device_count": &schema.Schema{
							Description: `Vn Fabric Control Plane Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_fabric_control_plane_good_health_device_count": &schema.Schema{
							Description: `Vn Fabric Control Plane Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_fabric_control_plane_good_health_percentage": &schema.Schema{
							Description: `Vn Fabric Control Plane Good Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_fabric_control_plane_no_health_device_count": &schema.Schema{
							Description: `Vn Fabric Control Plane No Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_fabric_control_plane_poor_health_device_count": &schema.Schema{
							Description: `Vn Fabric Control Plane Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_fabric_control_plane_total_device_count": &schema.Schema{
							Description: `Vn Fabric Control Plane Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_services_fair_health_device_count": &schema.Schema{
							Description: `Vn Services Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_services_good_health_device_count": &schema.Schema{
							Description: `Vn Services Good Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_services_health_percentage": &schema.Schema{
							Description: `Vn Services Health Percentage`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_services_no_health_device_count": &schema.Schema{
							Description: `Vn Services No Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_services_poor_health_device_count": &schema.Schema{
							Description: `Vn Services Poor Health Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_services_total_device_count": &schema.Schema{
							Description: `Vn Services Total Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"vn_status_fair_health_device_count": &schema.Schema{
							Description: `Vn Status Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_status_good_health_device_count": &schema.Schema{
							Description: `Vn Status Good Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_status_health_percentage": &schema.Schema{
							Description: `Vn Status Health Percentage`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"vn_status_no_health_device_count": &schema.Schema{
							Description: `Vn Status No Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_status_poor_health_device_count": &schema.Schema{
							Description: `Vn Status Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vn_status_total_device_count": &schema.Schema{
							Description: `Vn Status Total Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vni_fair_health_device_count": &schema.Schema{
							Description: `Vni Fair Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vni_good_health_device_count": &schema.Schema{
							Description: `Vni Good Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vni_good_health_percentage": &schema.Schema{
							Description: `Vni Good Health Percentage`,
							Type:        schema.TypeString, //TEST,
							Computed:    true,
						},

						"vni_no_health_device_count": &schema.Schema{
							Description: `Vni No Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vni_poor_health_device_count": &schema.Schema{
							Description: `Vni Poor Health Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vni_total_device_count": &schema.Schema{
							Description: `Vni Total Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"vnid": &schema.Schema{
							Description: `Vnid`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceVirtualNetworkHealthSummariesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vStartTime, okStartTime := d.GetOk("start_time")
	vEndTime, okEndTime := d.GetOk("end_time")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vID, okID := d.GetOk("id")
	vVnLayer, okVnLayer := d.GetOk("vn_layer")
	vAttribute, okAttribute := d.GetOk("attribute")
	vView, okView := d.GetOk("view")
	vXCaLLERID := d.Get("xca_lle_rid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ReadListOfVirtualNetworksWithTheirHealthSummaryV1")

		headerParams1 := catalystcentersdkgo.ReadListOfVirtualNetworksWithTheirHealthSummaryV1HeaderParams{}
		queryParams1 := catalystcentersdkgo.ReadListOfVirtualNetworksWithTheirHealthSummaryV1QueryParams{}

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
		if okVnLayer {
			queryParams1.VnLayer = vVnLayer.(string)
		}
		if okAttribute {
			queryParams1.Attribute = vAttribute.(string)
		}
		if okView {
			queryParams1.View = vView.(string)
		}
		headerParams1.XCaLLERID = vXCaLLERID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.ReadListOfVirtualNetworksWithTheirHealthSummaryV1(&headerParams1, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ReadListOfVirtualNetworksWithTheirHealthSummaryV1", err,
				"Failure at ReadListOfVirtualNetworksWithTheirHealthSummaryV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ReadListOfVirtualNetworksWithTheirHealthSummaryV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1Items(items *[]catalystcentersdkgo.ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["network_protocol"] = item.NetworkProtocol
		respItem["vlan"] = flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsVLAN(item.VLAN)
		respItem["vnid"] = item.Vnid
		respItem["layer"] = flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsLayer(item.Layer)
		respItem["total_fabric_sites"] = item.TotalFabricSites
		respItem["associated_l3_vn"] = flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsAssociatedL3Vn(item.AssociatedL3Vn)
		respItem["total_endpoints"] = item.TotalEndpoints
		respItem["good_health_percentage"] = item.GoodHealthPercentage
		respItem["total_device_count"] = item.TotalDeviceCount
		respItem["good_health_device_count"] = item.GoodHealthDeviceCount
		respItem["fair_health_device_count"] = item.FairHealthDeviceCount
		respItem["poor_health_device_count"] = item.PoorHealthDeviceCount
		respItem["no_health_device_count"] = item.NoHealthDeviceCount
		respItem["vn_fabric_control_plane_good_health_percentage"] = item.VnFabricControlPlaneGoodHealthPercentage
		respItem["vn_fabric_control_plane_total_device_count"] = item.VnFabricControlPlaneTotalDeviceCount
		respItem["vn_fabric_control_plane_good_health_device_count"] = item.VnFabricControlPlaneGoodHealthDeviceCount
		respItem["vn_fabric_control_plane_poor_health_device_count"] = item.VnFabricControlPlanePoorHealthDeviceCount
		respItem["vn_fabric_control_plane_fair_health_device_count"] = item.VnFabricControlPlaneFairHealthDeviceCount
		respItem["vn_fabric_control_plane_no_health_device_count"] = item.VnFabricControlPlaneNoHealthDeviceCount
		respItem["vn_services_health_percentage"] = item.VnServicesHealthPercentage
		respItem["vn_services_total_device_count"] = item.VnServicesTotalDeviceCount
		respItem["vn_services_good_health_device_count"] = item.VnServicesGoodHealthDeviceCount
		respItem["vn_services_poor_health_device_count"] = item.VnServicesPoorHealthDeviceCount
		respItem["vn_services_fair_health_device_count"] = item.VnServicesFairHealthDeviceCount
		respItem["vn_services_no_health_device_count"] = item.VnServicesNoHealthDeviceCount
		respItem["vn_exit_health_percentage"] = item.VnExitHealthPercentage
		respItem["vn_exit_total_device_count"] = item.VnExitTotalDeviceCount
		respItem["vn_exit_good_health_device_count"] = item.VnExitGoodHealthDeviceCount
		respItem["vn_exit_poor_health_device_count"] = item.VnExitPoorHealthDeviceCount
		respItem["vn_exit_fair_health_device_count"] = item.VnExitFairHealthDeviceCount
		respItem["vn_exit_no_health_device_count"] = item.VnExitNoHealthDeviceCount
		respItem["vn_status_health_percentage"] = flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsVnStatusHealthPercentage(item.VnStatusHealthPercentage)
		respItem["vn_status_total_device_count"] = item.VnStatusTotalDeviceCount
		respItem["vn_status_good_health_device_count"] = item.VnStatusGoodHealthDeviceCount
		respItem["vn_status_poor_health_device_count"] = item.VnStatusPoorHealthDeviceCount
		respItem["vn_status_fair_health_device_count"] = item.VnStatusFairHealthDeviceCount
		respItem["vn_status_no_health_device_count"] = item.VnStatusNoHealthDeviceCount
		respItem["pubsub_session_good_health_percentage"] = item.PubsubSessionGoodHealthPercentage
		respItem["pubsub_session_total_device_count"] = item.PubsubSessionTotalDeviceCount
		respItem["pubsub_session_good_health_device_count"] = item.PubsubSessionGoodHealthDeviceCount
		respItem["pubsub_session_poor_health_device_count"] = item.PubsubSessionPoorHealthDeviceCount
		respItem["pubsub_session_fair_health_device_count"] = item.PubsubSessionFairHealthDeviceCount
		respItem["pubsub_session_no_health_device_count"] = item.PubsubSessionNoHealthDeviceCount
		respItem["multi_cast_good_health_percentage"] = flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsMultiCastGoodHealthPercentage(item.MultiCastGoodHealthPercentage)
		respItem["multi_cast_total_device_count"] = item.MultiCastTotalDeviceCount
		respItem["multi_cast_good_health_device_count"] = item.MultiCastGoodHealthDeviceCount
		respItem["multi_cast_poor_health_device_count"] = item.MultiCastPoorHealthDeviceCount
		respItem["multi_cast_fair_health_device_count"] = item.MultiCastFairHealthDeviceCount
		respItem["multi_cast_no_health_device_count"] = item.MultiCastNoHealthDeviceCount
		respItem["internet_avail_good_health_percentage"] = item.InternetAvailGoodHealthPercentage
		respItem["internet_avail_total_device_count"] = item.InternetAvailTotalDeviceCount
		respItem["internet_avail_good_health_device_count"] = item.InternetAvailGoodHealthDeviceCount
		respItem["internet_avail_poor_health_device_count"] = item.InternetAvailPoorHealthDeviceCount
		respItem["internet_avail_fair_health_device_count"] = item.InternetAvailFairHealthDeviceCount
		respItem["internet_avail_no_health_device_count"] = item.InternetAvailNoHealthDeviceCount
		respItem["bgp_peer_good_health_percentage"] = item.BgpPeerGoodHealthPercentage
		respItem["bgp_peer_total_device_count"] = item.BgpPeerTotalDeviceCount
		respItem["bgp_peer_good_health_device_count"] = item.BgpPeerGoodHealthDeviceCount
		respItem["bgp_peer_poor_health_device_count"] = item.BgpPeerPoorHealthDeviceCount
		respItem["bgp_peer_fair_health_device_count"] = item.BgpPeerFairHealthDeviceCount
		respItem["bgp_peer_no_health_device_count"] = item.BgpPeerNoHealthDeviceCount
		respItem["vni_good_health_percentage"] = flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsVniGoodHealthPercentage(item.VniGoodHealthPercentage)
		respItem["vni_total_device_count"] = item.VniTotalDeviceCount
		respItem["vni_good_health_device_count"] = item.VniGoodHealthDeviceCount
		respItem["vni_poor_health_device_count"] = item.VniPoorHealthDeviceCount
		respItem["vni_fair_health_device_count"] = item.VniFairHealthDeviceCount
		respItem["vni_no_health_device_count"] = item.VniNoHealthDeviceCount
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsVLAN(item *catalystcentersdkgo.ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ResponseVLAN) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsLayer(item *catalystcentersdkgo.ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ResponseLayer) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsAssociatedL3Vn(item *catalystcentersdkgo.ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ResponseAssociatedL3Vn) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsVnStatusHealthPercentage(item *catalystcentersdkgo.ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ResponseVnStatusHealthPercentage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsMultiCastGoodHealthPercentage(item *catalystcentersdkgo.ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ResponseMultiCastGoodHealthPercentage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ItemsVniGoodHealthPercentage(item *catalystcentersdkgo.ResponseSdaReadListOfVirtualNetworksWithTheirHealthSummaryV1ResponseVniGoodHealthPercentage) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
