package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceIDBugsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of bugs affecting the network device
`,

		ReadContext: dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceIDBugsCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id query parameter. Id of the network bug
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId path parameter. Id of the network device
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter. Return network bugs with this severity. Available values : CATASTROPHIC, SEVERE, MODERATE
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
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceNetworkBugsResultsNetworkDevicesNetworkDeviceIDBugsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID := d.Get("network_device_id")
	vID, okID := d.GetOk("id")
	vSeverity, okSeverity := d.GetOk("severity")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfBugsAffectingTheNetworkDeviceV1")
		vvNetworkDeviceID := vNetworkDeviceID.(string)
		queryParams1 := catalystcentersdkgo.GetCountOfBugsAffectingTheNetworkDeviceV1QueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Compliance.GetCountOfBugsAffectingTheNetworkDeviceV1(vvNetworkDeviceID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfBugsAffectingTheNetworkDeviceV1", err,
				"Failure at GetCountOfBugsAffectingTheNetworkDeviceV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfBugsAffectingTheNetworkDeviceV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfBugsAffectingTheNetworkDeviceV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfBugsAffectingTheNetworkDeviceV1Item(item *catalystcentersdkgo.ResponseComplianceGetCountOfBugsAffectingTheNetworkDeviceV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
