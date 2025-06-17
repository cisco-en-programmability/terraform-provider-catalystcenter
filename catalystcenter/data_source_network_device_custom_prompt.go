package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceCustomPrompt() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on System Settings.

- Returns supported custom prompts by Catalyst Center
`,

		ReadContext: dataSourceNetworkDeviceCustomPromptRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom_password_prompt": &schema.Schema{
							Description: `Password for Custom Prompt
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"custom_username_prompt": &schema.Schema{
							Description: `Username for Custom Prompt
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"default_password_prompt": &schema.Schema{
							Description: `Default Password for Custom Prompt
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"default_username_prompt": &schema.Schema{
							Description: `Default Username for Custom Prompt
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

func dataSourceNetworkDeviceCustomPromptRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CustomPromptSupportGetAPI")

		// has_unknown_response: None

		response1, restyResp1, err := client.SystemSettings.CustomPromptSupportGetAPI()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CustomPromptSupportGetAPI", err,
				"Failure at CustomPromptSupportGetAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CustomPromptSupportGetAPI", err,
				"Failure at CustomPromptSupportGetAPI, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSystemSettingsCustomPromptSupportGetAPIItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CustomPromptSupportGetAPI response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSystemSettingsCustomPromptSupportGetAPIItem(item *catalystcentersdkgo.ResponseSystemSettingsCustomPromptSupportGetAPIResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["custom_username_prompt"] = item.CustomUsernamePrompt
	respItem["custom_password_prompt"] = item.CustomPasswordPrompt
	respItem["default_username_prompt"] = item.DefaultUsernamePrompt
	respItem["default_password_prompt"] = item.DefaultPasswordPrompt
	return []map[string]interface{}{
		respItem,
	}
}
