package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMapsSupportedAccessPoints() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Gets the list of supported access point types as well as valid antenna pattern names that can be used for each.
`,

		ReadContext: dataSourceMapsSupportedAccessPointsRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"antenna_patterns": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"band": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"names": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"ap_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceMapsSupportedAccessPointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: MapsSupportedAccessPointsV1")

		response1, restyResp1, err := client.Sites.MapsSupportedAccessPointsV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 MapsSupportedAccessPointsV1", err,
				"Failure at MapsSupportedAccessPointsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSitesMapsSupportedAccessPointsV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting MapsSupportedAccessPointsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesMapsSupportedAccessPointsV1Items(items *catalystcentersdkgo.ResponseSitesMapsSupportedAccessPointsV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["antenna_patterns"] = flattenSitesMapsSupportedAccessPointsV1ItemsAntennaPatterns(item.AntennaPatterns)
		respItem["ap_type"] = item.ApType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesMapsSupportedAccessPointsV1ItemsAntennaPatterns(items *[]catalystcentersdkgo.ResponseItemSitesMapsSupportedAccessPointsV1AntennaPatterns) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["band"] = item.Band
		respItem["names"] = item.Names
		respItems = append(respItems, respItem)
	}
	return respItems
}
