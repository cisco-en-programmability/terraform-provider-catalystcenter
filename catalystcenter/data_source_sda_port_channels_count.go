package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaPortChannelsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns the count of port channels that match the provided query parameters.
`,

		ReadContext: dataSourceSdaPortChannelsCountRead,
		Schema: map[string]*schema.Schema{
			"connected_device_type": &schema.Schema{
				Description: `connectedDeviceType query parameter. Connected device type of the port channel. The allowed values are [TRUNK, EXTENDED_NODE].
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric the device is assigned to.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. ID of the network device.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_channel_name": &schema.Schema{
				Description: `portChannelName query parameter. Name of the port channel.
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
							Description: `Number of port channels.
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

func dataSourceSdaPortChannelsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vPortChannelName, okPortChannelName := d.GetOk("port_channel_name")
	vConnectedDeviceType, okConnectedDeviceType := d.GetOk("connected_device_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPortChannelCountV1")
		queryParams1 := catalystcentersdkgo.GetPortChannelCountV1QueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okPortChannelName {
			queryParams1.PortChannelName = vPortChannelName.(string)
		}
		if okConnectedDeviceType {
			queryParams1.ConnectedDeviceType = vConnectedDeviceType.(string)
		}

		response1, restyResp1, err := client.Sda.GetPortChannelCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPortChannelCountV1", err,
				"Failure at GetPortChannelCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetPortChannelCountV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortChannelCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetPortChannelCountV1Item(item *catalystcentersdkgo.ResponseSdaGetPortChannelCountV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
