package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceRolePermissions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User and Roles.

- Get permissions for a role from Cisco DNA Center System.
`,

		ReadContext: dataSourceRolePermissionsRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"resource_types": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_permission": &schema.Schema{
										Description: `Default permission
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"description": &schema.Schema{
										Description: `Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
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

func dataSourceRolePermissionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPermissionsAPIV1")

		response1, restyResp1, err := client.UserandRoles.GetPermissionsAPIV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPermissionsAPIV1", err,
				"Failure at GetPermissionsAPIV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetPermissionsAPIV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPermissionsAPIV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserandRolesGetPermissionsAPIV1Item(item *catalystcentersdkgo.ResponseUserandRolesGetPermissionsAPIV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["resource_types"] = flattenUserandRolesGetPermissionsAPIV1ItemResourceTypes(item.ResourceTypes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenUserandRolesGetPermissionsAPIV1ItemResourceTypes(items *[]catalystcentersdkgo.ResponseUserandRolesGetPermissionsAPIV1ResponseResourceTypes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["display_name"] = item.DisplayName
		respItem["description"] = item.Description
		respItem["default_permission"] = item.DefaultPermission
		respItems = append(respItems, respItem)
	}
	return respItems
}
