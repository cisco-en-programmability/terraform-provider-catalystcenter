package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpWorkflowCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns the workflow count
`,

		ReadContext: dataSourcePnpWorkflowCountRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Description: `name query parameter. Workflow Name
`,
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Response`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourcePnpWorkflowCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetWorkflowCountV1")
		queryParams1 := catalystcentersdkgo.GetWorkflowCountV1QueryParams{}

		if okName {
			queryParams1.Name = interfaceToSliceString(vName)
		}

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetWorkflowCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetWorkflowCountV1", err,
				"Failure at GetWorkflowCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDeviceOnboardingPnpGetWorkflowCountV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetWorkflowCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetWorkflowCountV1Item(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpGetWorkflowCountV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	return []map[string]interface{}{
		respItem,
	}
}
