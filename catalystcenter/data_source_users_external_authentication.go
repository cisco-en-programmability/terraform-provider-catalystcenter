package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsersExternalAuthentication() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User and Roles.

- Get the External Authentication setting.
`,

		ReadContext: dataSourceUsersExternalAuthenticationRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"external_authentication_flag": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"enabled": &schema.Schema{
										Description: `External Authentication is enabled/disabled.
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
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

func dataSourceUsersExternalAuthenticationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetExternalAuthenticationSettingAPIV1")

		response1, restyResp1, err := client.UserandRoles.GetExternalAuthenticationSettingAPIV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetExternalAuthenticationSettingAPIV1", err,
				"Failure at GetExternalAuthenticationSettingAPIV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetExternalAuthenticationSettingAPIV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetExternalAuthenticationSettingAPIV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserandRolesGetExternalAuthenticationSettingAPIV1Item(item *catalystcentersdkgo.ResponseUserandRolesGetExternalAuthenticationSettingAPIV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["external_authentication_flag"] = flattenUserandRolesGetExternalAuthenticationSettingAPIV1ItemExternalAuthenticationFlag(item.ExternalAuthenticationFlag)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenUserandRolesGetExternalAuthenticationSettingAPIV1ItemExternalAuthenticationFlag(items *[]catalystcentersdkgo.ResponseUserandRolesGetExternalAuthenticationSettingAPIV1ResponseExternalAuthenticationFlag) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["enabled"] = boolPtrToString(item.Enabled)
		respItems = append(respItems, respItem)
	}
	return respItems
}
