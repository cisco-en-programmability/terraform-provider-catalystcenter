package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaLayer2VirtualNetworksCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns the count of layer 2 virtual networks that match the provided query parameters.
`,

		ReadContext: dataSourceSdaLayer2VirtualNetworksCountRead,
		Schema: map[string]*schema.Schema{
			"associated_layer3_virtual_network_name": &schema.Schema{
				Description: `associatedLayer3VirtualNetworkName query parameter. Name of the associated layer 3 virtual network.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the layer 2 virtual network is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"traffic_type": &schema.Schema{
				Description: `trafficType query parameter. The traffic type of the layer 2 virtual network.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"vlan_id": &schema.Schema{
				Description: `vlanId query parameter. The vlan ID of the layer 2 virtual network.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"vlan_name": &schema.Schema{
				Description: `vlanName query parameter. The vlan name of the layer 2 virtual network.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The number of layer 2 virtual networks
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSdaLayer2VirtualNetworksCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vVLANName, okVLANName := d.GetOk("vlan_name")
	vVLANID, okVLANID := d.GetOk("vlan_id")
	vTrafficType, okTrafficType := d.GetOk("traffic_type")
	vAssociatedLayer3VirtualNetworkName, okAssociatedLayer3VirtualNetworkName := d.GetOk("associated_layer3_virtual_network_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetLayer2VirtualNetworkCountV1")
		queryParams1 := catalystcentersdkgo.GetLayer2VirtualNetworkCountV1QueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okVLANName {
			queryParams1.VLANName = vVLANName.(string)
		}
		if okVLANID {
			queryParams1.VLANID = vVLANID.(float64)
		}
		if okTrafficType {
			queryParams1.TrafficType = vTrafficType.(string)
		}
		if okAssociatedLayer3VirtualNetworkName {
			queryParams1.AssociatedLayer3VirtualNetworkName = vAssociatedLayer3VirtualNetworkName.(string)
		}

		response1, restyResp1, err := client.Sda.GetLayer2VirtualNetworkCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetLayer2VirtualNetworkCountV1", err,
				"Failure at GetLayer2VirtualNetworkCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetLayer2VirtualNetworkCountV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetLayer2VirtualNetworkCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetLayer2VirtualNetworkCountV1Item(item *catalystcentersdkgo.ResponseSdaGetLayer2VirtualNetworkCountV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
