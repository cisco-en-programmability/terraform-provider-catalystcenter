package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceComplianceDeviceDetailsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Return  Compliance Count Detail
`,

		ReadContext: dataSourceComplianceDeviceDetailsCountRead,
		Schema: map[string]*schema.Schema{
			"compliance_status": &schema.Schema{
				Description: `complianceStatus query parameter. Specify "Compliance status(es)" separated by commas. The Compliance status can be 'COMPLIANT', 'NON_COMPLIANT', 'IN_PROGRESS', 'NOT_AVAILABLE', 'NOT_APPLICABLE', 'ERROR'.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"compliance_type": &schema.Schema{
				Description: `complianceType query parameter. Specify "Compliance type(s)" separated by commas. The Compliance type can be 'APPLICATION_VISIBILITY', 'EoX', 'FABRIC', 'IMAGE', 'NETWORK_PROFILE', 'NETWORK_SETTINGS', 'PSIRT', 'RUNNING_CONFIG', 'WORKFLOW'.
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
							Description: `Count of all devices or devices that match the query parameters.
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Version of API.
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

func dataSourceComplianceDeviceDetailsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vComplianceType, okComplianceType := d.GetOk("compliance_type")
	vComplianceStatus, okComplianceStatus := d.GetOk("compliance_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetComplianceDetailCountV1")
		queryParams1 := catalystcentersdkgo.GetComplianceDetailCountV1QueryParams{}

		if okComplianceType {
			queryParams1.ComplianceType = vComplianceType.(string)
		}
		if okComplianceStatus {
			queryParams1.ComplianceStatus = vComplianceStatus.(string)
		}

		response1, restyResp1, err := client.Compliance.GetComplianceDetailCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetComplianceDetailCountV1", err,
				"Failure at GetComplianceDetailCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetComplianceDetailCountV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetComplianceDetailCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetComplianceDetailCountV1Item(item *catalystcentersdkgo.ResponseComplianceGetComplianceDetailCountV1) []map[string]interface{} {
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
