package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceUsersExternalServersAAAAttribute() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on User and Roles.

- Get the current value of the custom AAA attribute.
`,

		ReadContext: dataSourceUsersExternalServersAAAAttributeRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"aaa_attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_name": &schema.Schema{
										Description: `Value of the custom AAA attribute name
`,
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

func dataSourceUsersExternalServersAAAAttributeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAAAAttributeAPIV1")

		response1, restyResp1, err := client.UserandRoles.GetAAAAttributeAPIV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAAAAttributeAPIV1", err,
				"Failure at GetAAAAttributeAPIV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenUserandRolesGetAAAAttributeAPIV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAAAAttributeAPIV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenUserandRolesGetAAAAttributeAPIV1Item(item *catalystcentersdkgo.ResponseUserandRolesGetAAAAttributeAPIV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["aaa_attributes"] = flattenUserandRolesGetAAAAttributeAPIV1ItemAAAAttributes(item.AAAAttributes)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenUserandRolesGetAAAAttributeAPIV1ItemAAAAttributes(items *[]catalystcentersdkgo.ResponseUserandRolesGetAAAAttributeAPIV1ResponseAAAAttributes) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_name"] = item.AttributeName
		respItems = append(respItems, respItem)
	}
	return respItems
}
