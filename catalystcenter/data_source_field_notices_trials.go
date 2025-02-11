package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesTrials() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get trial details for field notices detection on network devices
`,

		ReadContext: dataSourceFieldNoticesTrialsRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"active": &schema.Schema{
							Description: `Indicates if the trial is active
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"contract_level": &schema.Schema{
							Description: `Contract level for which trial was created. this was used in older versions and exists only for compatibility.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"end_time": &schema.Schema{
							Description: `Trial end time; as measured in Unix epoch time in milliseconds
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"feature": &schema.Schema{
							Description: `Name of the feature for which trial was created. for older versions that created contract type trials, this field will be absent.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"seconds_remaining_to_expiry": &schema.Schema{
							Description: `Seconds remaining in the trial before it expires. for expired trials this will be 0.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"seconds_since_expired": &schema.Schema{
							Description: `Seconds elapsed after the trial has expired. for active trials this will be 0.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"start_time": &schema.Schema{
							Description: `Trial start time; as measured in Unix epoch time in milliseconds
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Type of trial: * 'feature - the trial is of type feature. this is the currently supported type. * 'contract' - the trial is of type contract. this was used in older versions and exists only for compatibility.
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

func dataSourceFieldNoticesTrialsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1")

		// has_unknown_response: None

		response1, restyResp1, err := client.Compliance.GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1", err,
				"Failure at GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1Item(item *catalystcentersdkgo.ResponseComplianceGetTrialDetailsForFieldNoticesDetectionOnNetworkDevicesV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["feature"] = item.Feature
	respItem["contract_level"] = item.ContractLevel
	respItem["active"] = boolPtrToString(item.Active)
	respItem["start_time"] = item.StartTime
	respItem["end_time"] = item.EndTime
	respItem["seconds_remaining_to_expiry"] = item.SecondsRemainingToExpiry
	respItem["seconds_since_expired"] = item.SecondsSinceExpired
	return []map[string]interface{}{
		respItem,
	}
}
