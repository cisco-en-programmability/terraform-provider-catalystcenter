package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceQosDeviceInterfaceInfoCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Application Policy.

- Get the number of all existing qos device interface infos group by network device id
`,

		ReadContext: dataSourceQosDeviceInterfaceInfoCountRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Total number of Qos Device Interface Info
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Description: `Version
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

func dataSourceQosDeviceInterfaceInfoCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetQosDeviceInterfaceInfoCountV1")

		response1, restyResp1, err := client.ApplicationPolicy.GetQosDeviceInterfaceInfoCountV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetQosDeviceInterfaceInfoCountV1", err,
				"Failure at GetQosDeviceInterfaceInfoCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenApplicationPolicyGetQosDeviceInterfaceInfoCountV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetQosDeviceInterfaceInfoCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenApplicationPolicyGetQosDeviceInterfaceInfoCountV1Item(item *catalystcentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoCountV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
