package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceFamilyIDentifiersDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Software Image Management (SWIM).

- API to get Device Family Identifiers for all Device Families that can be used for tagging an image golden.
`,

		ReadContext: dataSourceDeviceFamilyIDentifiersDetailsRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_family": &schema.Schema{
							Description: `Device Family e.g. : Cisco Catalyst 6503 Switch-Catalyst 6500 Series Supervisor Engine 2T
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_family_identifier": &schema.Schema{
							Description: `Device Family Identifier used for tagging an image golden for certain Device Family e.g. : 277696480-283933147
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

func dataSourceDeviceFamilyIDentifiersDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceFamilyIDentifiersV1")

		response1, restyResp1, err := client.SoftwareImageManagementSwim.GetDeviceFamilyIDentifiersV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceFamilyIDentifiersV1", err,
				"Failure at GetDeviceFamilyIDentifiersV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceFamilyIDentifiersV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1Items(items *[]catalystcentersdkgo.ResponseSoftwareImageManagementSwimGetDeviceFamilyIDentifiersV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_family"] = item.DeviceFamily
		respItem["device_family_identifier"] = item.DeviceFamilyIDentifier
		respItems = append(respItems, respItem)
	}
	return respItems
}
