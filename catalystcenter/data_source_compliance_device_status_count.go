package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDeviceStatusCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return Compliance Status Count
`,

		ReadContext: dataSourceComplianceDeviceStatusCountRead,
		Schema: map[string]*schema.Schema{
			"compliance_status": &schema.Schema{
				Description: `complianceStatus query parameter. Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Returns count of compliant status
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Version of the API.
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

func dataSourceComplianceDeviceStatusCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vComplianceStatus, okComplianceStatus := d.GetOk("compliance_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetComplianceStatusCountV1")
		queryParams1 := catalystcentersdkgo.GetComplianceStatusCountV1QueryParams{}

		if okComplianceStatus {
			queryParams1.ComplianceStatus = vComplianceStatus.(string)
		}

		response1, restyResp1, err := client.Compliance.GetComplianceStatusCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetComplianceStatusCountV1", err,
				"Failure at GetComplianceStatusCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetComplianceStatusCountV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetComplianceStatusCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetComplianceStatusCountV1Item(item *catalystcentersdkgo.ResponseComplianceGetComplianceStatusCountV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["version"] = item.Version
	respItem["response"] = item.Response
	return []map[string]interface{}{
		respItem,
	}
}
