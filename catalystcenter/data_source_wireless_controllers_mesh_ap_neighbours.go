package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessControllersMeshApNeighbours() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Retrieves all Mesh accesspoint Neighbours details whether child, parent, etc.
`,

		ReadContext: dataSourceWirelessControllersMeshApNeighboursRead,
		Schema: map[string]*schema.Schema{
			"ap_name": &schema.Schema{
				Description: `apName query parameter. Employ this query parameter to obtain the details of the Access points corresponding to the provided ap name.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"ethernet_mac_address": &schema.Schema{
				Description: `ethernetMacAddress query parameter. Employ this query parameter to obtain the details of the Access points corresponding to the provided EthernetMacAddress.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"wlc_ip_address": &schema.Schema{
				Description: `wlcIpAddress query parameter. Employ this query parameter to obtain the details of the Access points corresponding to the provided WLC IP address.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_name": &schema.Schema{
							Description: `Name of the Wireless Access point
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ethernet_mac_address": &schema.Schema{
							Description: `AP Ethernet MacAddress mac
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"mesh_role": &schema.Schema{
							Description: `Mesh Role`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"neighbour_mac_address": &schema.Schema{
							Description: `AP Base Radio MacAddress mac.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"neighbour_type": &schema.Schema{
							Description: `Neighbour Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"wlc_ip_address": &schema.Schema{
							Description: `Device wireless Management IP
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

func dataSourceWirelessControllersMeshApNeighboursRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vWlcIPAddress, okWlcIPAddress := d.GetOk("wlc_ip_address")
	vApName, okApName := d.GetOk("ap_name")
	vEthernetMacAddress, okEthernetMacAddress := d.GetOk("ethernet_mac_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMeshApNeighboursV1")
		queryParams1 := catalystcentersdkgo.GetMeshApNeighboursV1QueryParams{}

		if okWlcIPAddress {
			queryParams1.WlcIPAddress = vWlcIPAddress.(string)
		}
		if okApName {
			queryParams1.ApName = vApName.(string)
		}
		if okEthernetMacAddress {
			queryParams1.EthernetMacAddress = vEthernetMacAddress.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.GetMeshApNeighboursV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetMeshApNeighboursV1", err,
				"Failure at GetMeshApNeighboursV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetMeshApNeighboursV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMeshApNeighboursV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetMeshApNeighboursV1Item(item *catalystcentersdkgo.ResponseWirelessGetMeshApNeighboursV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["ap_name"] = item.ApName
	respItem["ethernet_mac_address"] = item.EthernetMacAddress
	respItem["neighbour_mac_address"] = item.NeighbourMacAddress
	respItem["wlc_ip_address"] = item.WlcIPAddress
	respItem["neighbour_type"] = item.NeighbourType
	respItem["mesh_role"] = item.MeshRole
	return []map[string]interface{}{
		respItem,
	}
}
