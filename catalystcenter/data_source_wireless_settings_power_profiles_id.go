package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessSettingsPowerProfilesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- This data source allows the user to get a Power Profile by Power Profile ID that captured in wireless settings design
`,

		ReadContext: dataSourceWirelessSettingsPowerProfilesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Power Profile ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `The description of the Power Profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Unique Identifier of the power profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"profile_name": &schema.Schema{
							Description: `The Name of the Power Profile.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"interface_id": &schema.Schema{
										Description: `Interface Id for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"interface_type": &schema.Schema{
										Description: `Interface Type for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"parameter_type": &schema.Schema{
										Description: `Parameter Type for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"parameter_value": &schema.Schema{
										Description: `Parameter Value for the rule.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"sequence": &schema.Schema{
										Description: `Sequential Ordered List of rules for Power Profile.
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessSettingsPowerProfilesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPowerProfileByIDV1")
		vvID := vID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Wireless.GetPowerProfileByIDV1(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPowerProfileByIDV1", err,
				"Failure at GetPowerProfileByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenWirelessGetPowerProfileByIDV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPowerProfileByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetPowerProfileByIDV1Item(item *catalystcentersdkgo.ResponseWirelessGetPowerProfileByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["profile_name"] = item.ProfileName
	respItem["description"] = item.Description
	respItem["rules"] = flattenWirelessGetPowerProfileByIDV1ItemRules(item.Rules)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenWirelessGetPowerProfileByIDV1ItemRules(items *[]catalystcentersdkgo.ResponseWirelessGetPowerProfileByIDV1ResponseRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["sequence"] = item.Sequence
		respItem["interface_type"] = item.InterfaceType
		respItem["interface_id"] = item.InterfaceID
		respItem["parameter_type"] = item.ParameterType
		respItem["parameter_value"] = item.ParameterValue
		respItems = append(respItems, respItem)
	}
	return respItems
}
