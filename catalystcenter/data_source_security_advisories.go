package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityAdvisories() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Security Advisories.

- Retrieves list of advisories on the network
`,

		ReadContext: dataSourceSecurityAdvisoriesRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"advisory_id": &schema.Schema{
							Description: `Advisory Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"cves": &schema.Schema{
							Description: `Cves`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"default_config_match_pattern": &schema.Schema{
							Description: `Default Config Match Pattern`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"default_detection_type": &schema.Schema{
							Description: `Default Detection Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"detection_type": &schema.Schema{
							Description: `Detection Type`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"device_count": &schema.Schema{
							Description: `Device Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"fixed_versions": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"e1a_version": &schema.Schema{
										Description: `E1a Version`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"hidden_device_count": &schema.Schema{
							Description: `Hidden Device Count`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"publication_url": &schema.Schema{
							Description: `Publication Url`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"sir": &schema.Schema{
							Description: `Sir`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSecurityAdvisoriesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdvisoriesList")

		// has_unknown_response: None

		response1, restyResp1, err := client.SecurityAdvisories.GetAdvisoriesList()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAdvisoriesList", err,
				"Failure at GetAdvisoriesList, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAdvisoriesList", err,
				"Failure at GetAdvisoriesList, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSecurityAdvisoriesGetAdvisoriesListItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdvisoriesList response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSecurityAdvisoriesGetAdvisoriesListItems(items *[]catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesListResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["advisory_id"] = item.AdvisoryID
		respItem["device_count"] = item.DeviceCount
		respItem["hidden_device_count"] = item.HiddenDeviceCount
		respItem["cves"] = item.Cves
		respItem["publication_url"] = item.PublicationURL
		respItem["sir"] = item.Sir
		respItem["detection_type"] = item.DetectionType
		respItem["default_detection_type"] = item.DefaultDetectionType
		respItem["default_config_match_pattern"] = item.DefaultConfigMatchPattern
		respItem["fixed_versions"] = flattenSecurityAdvisoriesGetAdvisoriesListItemsFixedVersions(item.FixedVersions)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSecurityAdvisoriesGetAdvisoriesListItemsFixedVersions(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesListResponseFixedVersions) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["e1a_version"] = item.V27E1A

	return []map[string]interface{}{
		respItem,
	}

}
