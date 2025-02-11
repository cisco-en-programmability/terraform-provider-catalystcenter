package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Fetches registration status, authorization status and entitlements of the system with Cisco Smart Software Manage
(CSSM).
`,

		ReadContext: dataSourceLicenseStatusRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"authorization_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"evaluation_remainder_timestamp": &schema.Schema{
										Description: `A date and time represented as milliseconds since the Unix epoch.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"expiry_timestamp": &schema.Schema{
										Description: `A date and time represented as milliseconds since the Unix epoch.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"last_attempt_fail_reason": &schema.Schema{
										Description: `The reason for last authorization request failure
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_attempt_status": &schema.Schema{
										Description: `The last authorization request's status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_attempt_timestamp": &schema.Schema{
										Description: `A date and time represented as milliseconds since the Unix epoch.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"next_attempt_timestamp": &schema.Schema{
										Description: `A date and time represented as milliseconds since the Unix epoch.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `This denotes the authorization status of the system.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"entitlements": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Description: `Name or description of the license entitlement
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `This denotes the authorization status of the available licenses.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"tag": &schema.Schema{
										Description: `Entitlement tag associated with the available licenses
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"usage_count": &schema.Schema{
										Description: `Available license count
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},

						"export_control": &schema.Schema{
							Description: `Export-Controlled setting of Smart Account
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"registration_status": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"expiry_timestamp": &schema.Schema{
										Description: `A date and time represented as milliseconds since the Unix epoch.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"last_attempt_fail_reason": &schema.Schema{
										Description: `The reason for last registration request failure
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_attempt_status": &schema.Schema{
										Description: `The last registration request's status
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"last_attempt_timestamp": &schema.Schema{
										Description: `A date and time represented as milliseconds since the Unix epoch.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"next_attempt_timestamp": &schema.Schema{
										Description: `A date and time represented as milliseconds since the Unix epoch.
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `REGISTERED if the system is registered with CSSM, otherwise UNREGISTERED.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"smart_account_id": &schema.Schema{
							Description: `Smart Account id to which the system is registered
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"virtual_account_id": &schema.Schema{
							Description: `Virtual Account id to which the system is registered
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

func dataSourceLicenseStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SystemLicensingStatusV1")

		// has_unknown_response: None

		response1, restyResp1, err := client.Licenses.SystemLicensingStatusV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 SystemLicensingStatusV1", err,
				"Failure at SystemLicensingStatusV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenLicensesSystemLicensingStatusV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SystemLicensingStatusV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesSystemLicensingStatusV1Item(item *catalystcentersdkgo.ResponseLicensesSystemLicensingStatusV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["registration_status"] = flattenLicensesSystemLicensingStatusV1ItemRegistrationStatus(item.RegistrationStatus)
	respItem["authorization_status"] = flattenLicensesSystemLicensingStatusV1ItemAuthorizationStatus(item.AuthorizationStatus)
	respItem["entitlements"] = flattenLicensesSystemLicensingStatusV1ItemEntitlements(item.Entitlements)
	respItem["smart_account_id"] = item.SmartAccountID
	respItem["virtual_account_id"] = item.VirtualAccountID
	respItem["export_control"] = item.ExportControl
	return []map[string]interface{}{
		respItem,
	}
}

func flattenLicensesSystemLicensingStatusV1ItemRegistrationStatus(item *catalystcentersdkgo.ResponseLicensesSystemLicensingStatusV1ResponseRegistrationStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["last_attempt_timestamp"] = item.LastAttemptTimestamp
	respItem["expiry_timestamp"] = item.ExpiryTimestamp
	respItem["next_attempt_timestamp"] = item.NextAttemptTimestamp
	respItem["last_attempt_status"] = item.LastAttemptStatus
	respItem["last_attempt_fail_reason"] = item.LastAttemptFailReason

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesSystemLicensingStatusV1ItemAuthorizationStatus(item *catalystcentersdkgo.ResponseLicensesSystemLicensingStatusV1ResponseAuthorizationStatus) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["status"] = item.Status
	respItem["last_attempt_timestamp"] = item.LastAttemptTimestamp
	respItem["evaluation_remainder_timestamp"] = item.EvaluationRemainderTimestamp
	respItem["expiry_timestamp"] = item.ExpiryTimestamp
	respItem["next_attempt_timestamp"] = item.NextAttemptTimestamp
	respItem["last_attempt_status"] = item.LastAttemptStatus
	respItem["last_attempt_fail_reason"] = item.LastAttemptFailReason

	return []map[string]interface{}{
		respItem,
	}

}

func flattenLicensesSystemLicensingStatusV1ItemEntitlements(item *catalystcentersdkgo.ResponseLicensesSystemLicensingStatusV1ResponseEntitlements) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["tag"] = item.Tag
	respItem["description"] = item.Description
	respItem["usage_count"] = item.UsageCount
	respItem["status"] = item.Status

	return []map[string]interface{}{
		respItem,
	}

}
