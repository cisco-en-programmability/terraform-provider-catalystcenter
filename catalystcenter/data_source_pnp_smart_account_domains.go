package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourcePnpSmartAccountDomains() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Device Onboarding (PnP).

- Returns the list of Smart Account domains
`,

		ReadContext: dataSourcePnpSmartAccountDomainsRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourcePnpSmartAccountDomainsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSmartAccountListV1")

		response1, restyResp1, err := client.DeviceOnboardingPnp.GetSmartAccountListV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSmartAccountListV1", err,
				"Failure at GetSmartAccountListV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDeviceOnboardingPnpGetSmartAccountListV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSmartAccountListV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDeviceOnboardingPnpGetSmartAccountListV1Items(items *catalystcentersdkgo.ResponseDeviceOnboardingPnpGetSmartAccountListV1) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, respItem)
	}
	return respItems
}
