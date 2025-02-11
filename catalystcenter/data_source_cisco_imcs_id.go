package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCiscoImcsID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Cisco IMC.

- This data source retrieves the Cisco Integrated Management Controller (IMC) configuration for a Catalyst Center node,
identified by the specified ID.
The Cisco IMC configuration APIs enable the management of connections between Cisco IMC and Cisco Catalyst Center. By
providing the Cisco IMC IP address and credentials to Catalyst Center, Catalyst Center can access and report the health
status of hardware components within the Cisco appliance.
More data about the Cisco IMC can be retrieved using the APIs exposed directly by Cisco IMC. Details are available in
the Cisco IMC documentation https://www.cisco.com/c/en/us/support/servers-unified-computing/ucs-c-series-integrated-
management-controller/series.html#~tab-documents
The Cisco IMC configuration is relevant only for Catalyst Center deployments based on UCS appliances. In cases where
Cisco IMC configuration is not supported by the deployment, these APIs will respond with a 404 Not Found status code.
`,

		ReadContext: dataSourceCiscoImcsIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The unique identifier for this Cisco IMC configuration
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `The unique identifier for this Cisco IMC configuration
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": &schema.Schema{
							Description: `IP address of the Cisco IMC
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"node_id": &schema.Schema{
							Description: `The UUID that represents the Catalyst Center node. Its value can be obtained from the id attribute of the response of the /dna/intent/api/v1/nodes-config API.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"username": &schema.Schema{
							Description: `Username of the Cisco IMC
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

func dataSourceCiscoImcsIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1")
		vvID := vID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.CiscoIMC.RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1", err,
				"Failure at RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1Item(item *catalystcentersdkgo.ResponseCiscoIMCRetrievesTheCiscoIMCConfigurationForACatalystCenterNodeV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["node_id"] = item.NodeID
	respItem["ip_address"] = item.IPAddress
	respItem["username"] = item.Username
	return []map[string]interface{}{
		respItem,
	}
}
