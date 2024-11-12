package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceWirelessDynamicInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Get one or all dynamic interface(s)
`,

		ReadContext: dataSourceWirelessDynamicInterfaceRead,
		Schema: map[string]*schema.Schema{
			"interface_name": &schema.Schema{
				Description: `interface-name query parameter. dynamic-interface name, if not specified all the existing dynamic interfaces will be retrieved
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interface_name": &schema.Schema{
							Description: `dynamic interface name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vlan_id": &schema.Schema{
							Description: `Vlan id
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceWirelessDynamicInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vInterfaceName, okInterfaceName := d.GetOk("interface_name")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDynamicInterfaceV1")
		queryParams1 := catalystcentersdkgo.GetDynamicInterfaceV1QueryParams{}

		if okInterfaceName {
			queryParams1.InterfaceName = vInterfaceName.(string)
		}

		response1, restyResp1, err := client.Wireless.GetDynamicInterfaceV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDynamicInterfaceV1", err,
				"Failure at GetDynamicInterfaceV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetDynamicInterfaceV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDynamicInterfaceV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetDynamicInterfaceV1Items(items *catalystcentersdkgo.ResponseWirelessGetDynamicInterfaceV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["interface_name"] = item.InterfaceName
		respItem["vlan_id"] = item.VLANID
		respItems = append(respItems, respItem)
	}
	return respItems
}
