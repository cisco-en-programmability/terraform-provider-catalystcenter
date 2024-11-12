package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTransitPeerNetwork() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get Transit Peer Network Info from SD-Access
`,

		ReadContext: dataSourceTransitPeerNetworkRead,
		Schema: map[string]*schema.Schema{
			"transit_peer_network_name": &schema.Schema{
				Description: `transitPeerNetworkName query parameter. Transit or Peer Network Name
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Transit Peer network info retrieved successfully
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"autonomous_system_number": &schema.Schema{
										Description: `Autonomous System Number  
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"routing_protocol_name": &schema.Schema{
										Description: `Routing Protocol Name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"sda_transit_settings": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"transit_control_plane_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"device_management_ip_address": &schema.Schema{
													Description: `Device Management Ip Address 
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
											},
										},
									},
								},
							},
						},

						"status": &schema.Schema{
							Description: `status
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"transit_peer_network_id": &schema.Schema{
							Description: `Transit Peer Network Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"transit_peer_network_name": &schema.Schema{
							Description: `Transit Peer Network Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"transit_peer_network_type": &schema.Schema{
							Description: `Transit Peer Network Type
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

func dataSourceTransitPeerNetworkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vTransitPeerNetworkName := d.Get("transit_peer_network_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTransitPeerNetworkInfoV1")
		queryParams1 := catalystcentersdkgo.GetTransitPeerNetworkInfoV1QueryParams{}

		queryParams1.TransitPeerNetworkName = vTransitPeerNetworkName.(string)

		response1, restyResp1, err := client.Sda.GetTransitPeerNetworkInfoV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTransitPeerNetworkInfoV1", err,
				"Failure at GetTransitPeerNetworkInfoV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetTransitPeerNetworkInfoV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTransitPeerNetworkInfoV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetTransitPeerNetworkInfoV1Item(item *catalystcentersdkgo.ResponseSdaGetTransitPeerNetworkInfoV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["transit_peer_network_name"] = item.TransitPeerNetworkName
	respItem["transit_peer_network_type"] = item.TransitPeerNetworkType
	respItem["ip_transit_settings"] = flattenSdaGetTransitPeerNetworkInfoV1ItemIPTransitSettings(item.IPTransitSettings)
	respItem["sda_transit_settings"] = flattenSdaGetTransitPeerNetworkInfoV1ItemSdaTransitSettings(item.SdaTransitSettings)
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	respItem["transit_peer_network_id"] = item.TransitPeerNetworkID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSdaGetTransitPeerNetworkInfoV1ItemIPTransitSettings(item *catalystcentersdkgo.ResponseSdaGetTransitPeerNetworkInfoV1IPTransitSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["routing_protocol_name"] = item.RoutingProtocolName
	respItem["autonomous_system_number"] = item.AutonomousSystemNumber

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetTransitPeerNetworkInfoV1ItemSdaTransitSettings(item *catalystcentersdkgo.ResponseSdaGetTransitPeerNetworkInfoV1SdaTransitSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["transit_control_plane_settings"] = flattenSdaGetTransitPeerNetworkInfoV1ItemSdaTransitSettingsTransitControlPlaneSettings(item.TransitControlPlaneSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSdaGetTransitPeerNetworkInfoV1ItemSdaTransitSettingsTransitControlPlaneSettings(items *[]catalystcentersdkgo.ResponseSdaGetTransitPeerNetworkInfoV1SdaTransitSettingsTransitControlPlaneSettings) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["site_name_hierarchy"] = item.SiteNameHierarchy
		respItem["device_management_ip_address"] = item.DeviceManagementIPAddress
		respItems = append(respItems, respItem)
	}
	return respItems
}
