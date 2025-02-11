package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

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

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"advisory_id": &schema.Schema{
							Description: `Id of the advisory
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"cves": &schema.Schema{
							Description: `CVE (Common Vulnerabilities and Exposures) IDs of the advisory
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"default_config_match_pattern": &schema.Schema{
							Description: `Regular expression used by the system to detect the advisory
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"default_detection_type": &schema.Schema{
							Description: `Original criteria for advisory detection
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"detection_type": &schema.Schema{
							Description: `Criteria for advisory detection
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_count": &schema.Schema{
							Description: `Number of devices vulnerable to the advisory
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"fixed_versions": &schema.Schema{
							Description: `Map where each key is a vulnerable version and the value is a list of versions in which the advisory has been fixed
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"hidden_device_count": &schema.Schema{
							Description: `Number of devices vulnerable to the advisory but were suppressed by the user
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"publication_url": &schema.Schema{
							Description: `CISCO publication URL for the advisory
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sir": &schema.Schema{
							Description: `Security Impact Rating of the advisory
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

func dataSourceSecurityAdvisoriesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAdvisoriesListV1")

		response1, restyResp1, err := client.SecurityAdvisories.GetAdvisoriesListV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAdvisoriesListV1", err,
				"Failure at GetAdvisoriesListV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSecurityAdvisoriesGetAdvisoriesListV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAdvisoriesListV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSecurityAdvisoriesGetAdvisoriesListV1Item(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesListV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
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
	respItem["fixed_versions"] = flattenSecurityAdvisoriesGetAdvisoriesListV1ItemFixedVersions(item.FixedVersions)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSecurityAdvisoriesGetAdvisoriesListV1ItemFixedVersions(item *catalystcentersdkgo.ResponseSecurityAdvisoriesGetAdvisoriesListV1ResponseFixedVersions) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
