package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLicenseDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Licenses.

- Retrieve details of all smart accounts.
`,

		ReadContext: dataSourceLicenseDeviceRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"domain": &schema.Schema{
							Description: `Domain of smart account
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of smart account
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_active_smart_account": &schema.Schema{
							Description: `Is active smart account
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of smart account
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

func dataSourceLicenseDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: SmartAccountDetailsV1")

		response1, restyResp1, err := client.Licenses.SmartAccountDetailsV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 SmartAccountDetailsV1", err,
				"Failure at SmartAccountDetailsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLicensesSmartAccountDetailsV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SmartAccountDetailsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLicensesSmartAccountDetailsV1Items(items *[]catalystcentersdkgo.ResponseLicensesSmartAccountDetailsV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["domain"] = item.Domain
		respItem["is_active_smart_account"] = boolPtrToString(item.IsActiveSmartAccount)
		respItems = append(respItems, respItem)
	}
	return respItems
}
