package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFileNamespaces() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on File.

- Returns list of available namespaces
`,

		ReadContext: dataSourceFileNamespacesRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFileNamespacesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetListOfAvailableNamespacesV1")

		response1, restyResp1, err := client.File.GetListOfAvailableNamespacesV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetListOfAvailableNamespacesV1", err,
				"Failure at GetListOfAvailableNamespacesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenFileGetListOfAvailableNamespacesV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetListOfAvailableNamespacesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenFileGetListOfAvailableNamespacesV1Items(items *catalystcentersdkgo.ResponseFileGetListOfAvailableNamespacesV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = items.Response
	respItem["version"] = items.Version
	return []map[string]interface{}{
		respItem,
	}
}
