package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Get list of all properties & operations valid for an interface.
`,

		ReadContext: dataSourceInterfaceRead,
		Schema: map[string]*schema.Schema{
			"interface_uuid": &schema.Schema{
				Description: `interfaceUuid path parameter. Interface ID
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"interface_uuid": &schema.Schema{
							Description: `Id of the Interface
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"operations": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"applicable": &schema.Schema{
										Description: `Checks if operation is applicable to interface
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_reason": &schema.Schema{
										Description: `Failure reason of the Operation
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the Operation
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"applicable": &schema.Schema{
										Description: `Checks if property is applicable to interface
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_reason": &schema.Schema{
										Description: `Failure reason of the Property
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Name of the Property
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vInterfaceUUID := d.Get("interface_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LegitOperationsForInterfaceV1")
		vvInterfaceUUID := vInterfaceUUID.(string)

		response1, restyResp1, err := client.Devices.LegitOperationsForInterfaceV1(vvInterfaceUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LegitOperationsForInterfaceV1", err,
				"Failure at LegitOperationsForInterfaceV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesLegitOperationsForInterfaceV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LegitOperationsForInterfaceV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesLegitOperationsForInterfaceV1Item(item *catalystcentersdkgo.ResponseDevicesLegitOperationsForInterfaceV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["interface_uuid"] = item.InterfaceUUID
	respItem["properties"] = flattenDevicesLegitOperationsForInterfaceV1ItemProperties(item.Properties)
	respItem["operations"] = flattenDevicesLegitOperationsForInterfaceV1ItemOperations(item.Operations)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesLegitOperationsForInterfaceV1ItemProperties(items *[]catalystcentersdkgo.ResponseDevicesLegitOperationsForInterfaceV1ResponseProperties) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["applicable"] = item.Applicable
		respItem["failure_reason"] = item.FailureReason
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesLegitOperationsForInterfaceV1ItemOperations(items *[]catalystcentersdkgo.ResponseDevicesLegitOperationsForInterfaceV1ResponseOperations) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["name"] = item.Name
		respItem["applicable"] = item.Applicable
		respItem["failure_reason"] = item.FailureReason
		respItems = append(respItems, respItem)
	}
	return respItems
}
