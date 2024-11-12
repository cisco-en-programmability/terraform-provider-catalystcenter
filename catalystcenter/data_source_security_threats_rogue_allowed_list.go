package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSecurityThreatsRogueAllowedList() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Intent API to fetch all the allowed mac addresses in the system.
`,

		ReadContext: dataSourceSecurityThreatsRogueAllowedListRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The maximum number of entries to return. If the value exceeds the total count, then the maximum entries will be returned.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The offset of the first item in the collection to return.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"category": &schema.Schema{
							Description: `Category`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"last_modified": &schema.Schema{
							Description: `Last Modified`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"mac_address": &schema.Schema{
							Description: `Mac Address`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSecurityThreatsRogueAllowedListRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllowedMacAddressV1")
		queryParams1 := catalystcentersdkgo.GetAllowedMacAddressV1QueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Devices.GetAllowedMacAddressV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllowedMacAddressV1", err,
				"Failure at GetAllowedMacAddressV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetAllowedMacAddressV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllowedMacAddressV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetAllowedMacAddressV1Items(items *catalystcentersdkgo.ResponseDevicesGetAllowedMacAddressV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["mac_address"] = item.MacAddress
		respItem["category"] = item.Category
		respItem["last_modified"] = item.LastModified
		respItems = append(respItems, respItem)
	}
	return respItems
}