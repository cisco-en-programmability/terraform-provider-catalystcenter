package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDiscoveryDeviceCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns the count of network devices discovered in the given discovery. Discovery ID can be obtained using the "Get
Discoveries by range" API.
`,

		ReadContext: dataSourceDiscoveryDeviceCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Discovery ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"task_id": &schema.Schema{
				Description: `taskId query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `The count of network devices discovered in the given discovery
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDiscoveryDeviceCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vTaskID, okTaskID := d.GetOk("task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDevicesDiscoveredByIDV1")
		vvID := vID.(string)
		queryParams1 := catalystcentersdkgo.GetDevicesDiscoveredByIDV1QueryParams{}

		if okTaskID {
			queryParams1.TaskID = vTaskID.(string)
		}

		response1, restyResp1, err := client.Discovery.GetDevicesDiscoveredByIDV1(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDevicesDiscoveredByIDV1", err,
				"Failure at GetDevicesDiscoveredByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDiscoveryGetDevicesDiscoveredByIDV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDevicesDiscoveredByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetDevicesDiscoveredByIDV1Item(item *catalystcentersdkgo.ResponseDiscoveryGetDevicesDiscoveredByIDV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
