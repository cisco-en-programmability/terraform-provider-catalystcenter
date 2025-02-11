package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDnacPackages() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Platform.

- Provides information such as name, version of packages installed on the Catalyst center.
`,

		ReadContext: dataSourceDnacPackagesRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"name": &schema.Schema{
							Description: `Name of installed package
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Version of installed package
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

func dataSourceDnacPackagesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CiscoCatalystCenterPackagesSummaryV1")

		// has_unknown_response: None

		response1, restyResp1, err := client.Platform.CiscoCatalystCenterPackagesSummaryV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CiscoCatalystCenterPackagesSummaryV1", err,
				"Failure at CiscoCatalystCenterPackagesSummaryV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenPlatformCiscoCatalystCenterPackagesSummaryV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CiscoCatalystCenterPackagesSummaryV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenPlatformCiscoCatalystCenterPackagesSummaryV1Items(items *[]catalystcentersdkgo.ResponsePlatformCiscoCatalystCenterPackagesSummaryV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["version"] = item.Version
		respItems = append(respItems, respItem)
	}
	return respItems
}
