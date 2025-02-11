package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseSetting() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Retrieves license setting Default smart account id and virtual account id for auto registration of devices for smart
license flow. If default smart account is not configured, 'defaultSmartAccountId' is 'null'. Similarly, if auto
registration of devices for smart license flow is not enabled, 'autoRegistrationVirtualAccountId' is 'null'. For smart
proxy connection mode, 'autoRegistrationVirtualAccountId' is always 'null'.
`,

		ReadContext: dataSourceLicenseSettingRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"auto_registration_virtual_account_id": &schema.Schema{
							Description: `Virtual account id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"default_smart_account_id": &schema.Schema{
							Description: `Default smart account id
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

func dataSourceLicenseSettingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveLicenseSettingV1")

		response1, restyResp1, err := client.Licenses.RetrieveLicenseSettingV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveLicenseSettingV1", err,
				"Failure at RetrieveLicenseSettingV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesRetrieveLicenseSettingV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveLicenseSettingV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesRetrieveLicenseSettingV1Item(item *catalystcentersdkgo.ResponseLicensesRetrieveLicenseSettingV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["default_smart_account_id"] = item.DefaultSmartAccountID
	respItem["auto_registration_virtual_account_id"] = item.AutoRegistrationVirtualAccountID
	return []map[string]interface{}{
		respItem,
	}
}
