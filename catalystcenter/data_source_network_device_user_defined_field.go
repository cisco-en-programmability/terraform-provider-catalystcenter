package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceUserDefinedField() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Gets existing global User Defined Fields. If no input is given, it fetches ALL the Global UDFs. Filter/search is
supported by UDF Id(s) or UDF name(s) or both.
`,

		ReadContext: dataSourceNetworkDeviceUserDefinedFieldRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id query parameter. Comma-seperated id(s) used for search/filtering
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Comma-seperated name(s) used for search/filtering
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description for UDF
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `DeviceId of the Device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `UDF name
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

func dataSourceNetworkDeviceUserDefinedFieldRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vName, okName := d.GetOk("name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllUserDefinedFieldsV1")
		queryParams1 := catalystcentersdkgo.GetAllUserDefinedFieldsV1QueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}

		response1, restyResp1, err := client.Devices.GetAllUserDefinedFieldsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllUserDefinedFieldsV1", err,
				"Failure at GetAllUserDefinedFieldsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetAllUserDefinedFieldsV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllUserDefinedFieldsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetAllUserDefinedFieldsV1Items(items *[]catalystcentersdkgo.ResponseDevicesGetAllUserDefinedFieldsV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItems = append(respItems, respItem)
	}
	return respItems
}
