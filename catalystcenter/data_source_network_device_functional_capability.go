package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceFunctionalCapability() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns the functional-capability for given devices

- Returns functional capability with given Id
`,

		ReadContext: dataSourceNetworkDeviceFunctionalCapabilityRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId query parameter. Accepts comma separated deviceid's and return list of functional-capabilities for the given id's. If invalid or not-found id's are provided, null entry will be returned in the list.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"function_name": &schema.Schema{
				Description: `functionName query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"id": &schema.Schema{
				Description: `id path parameter. Functional Capability UUID
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Description: `Deprecated
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"function_details": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_info": &schema.Schema{
										Description: `Deprecated
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Deprecated
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"property_name": &schema.Schema{
										Description: `Property Name of the function
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"string_value": &schema.Schema{
										Description: `Value for the property
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"function_name": &schema.Schema{
							Description: `Name of the function
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"function_op_state": &schema.Schema{
							Description: `Operational state of the function
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of the function
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attribute_info": &schema.Schema{
							Description: `Deprecated
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"device_id": &schema.Schema{
							Description: `Device Id of the device
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"functional_capability": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"attribute_info": &schema.Schema{
										Description: `Deprecated
`,
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"function_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"attribute_info": &schema.Schema{
													Description: `Deprecated
`,
													Type:     schema.TypeString, //TEST,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Deprecated
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"property_name": &schema.Schema{
													Description: `Property Name of the function
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"string_value": &schema.Schema{
													Description: `Value for the property
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"function_name": &schema.Schema{
										Description: `Name of the function
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"function_op_state": &schema.Schema{
										Description: `Operational state of the function
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Id of the function
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Deprecated
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

func dataSourceNetworkDeviceFunctionalCapabilityRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vFunctionName, okFunctionName := d.GetOk("function_name")
	vID, okID := d.GetOk("id")

	method1 := []bool{okDeviceID, okFunctionName}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFunctionalCapabilityForDevicesV1")
		queryParams1 := catalystcentersdkgo.GetFunctionalCapabilityForDevicesV1QueryParams{}

		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okFunctionName {
			queryParams1.FunctionName = interfaceToSliceString(vFunctionName)
		}

		response1, restyResp1, err := client.Devices.GetFunctionalCapabilityForDevicesV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFunctionalCapabilityForDevicesV1", err,
				"Failure at GetFunctionalCapabilityForDevicesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetFunctionalCapabilityForDevicesV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFunctionalCapabilityForDevicesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetFunctionalCapabilityByIDV1")
		vvID := vID.(string)

		response2, restyResp2, err := client.Devices.GetFunctionalCapabilityByIDV1(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetFunctionalCapabilityByIDV1", err,
				"Failure at GetFunctionalCapabilityByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetFunctionalCapabilityByIDV1Item(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFunctionalCapabilityByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetFunctionalCapabilityForDevicesV1Items(items *[]catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsAttributeInfo(item.AttributeInfo)
		respItem["device_id"] = item.DeviceID
		respItem["functional_capability"] = flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsFunctionalCapability(item.FunctionalCapability)
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsAttributeInfo(item *catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsFunctionalCapability(items *[]catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapability) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsFunctionalCapabilityAttributeInfo(item.AttributeInfo)
		respItem["function_details"] = flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsFunctionalCapabilityFunctionDetails(item.FunctionDetails)
		respItem["function_name"] = item.FunctionName
		respItem["function_op_state"] = item.FunctionOpState
		respItem["id"] = item.ID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsFunctionalCapabilityAttributeInfo(item *catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsFunctionalCapabilityFunctionDetails(items *[]catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityFunctionDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsFunctionalCapabilityFunctionDetailsAttributeInfo(item.AttributeInfo)
		respItem["id"] = item.ID
		respItem["property_name"] = item.PropertyName
		respItem["string_value"] = item.StringValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetFunctionalCapabilityForDevicesV1ItemsFunctionalCapabilityFunctionDetailsAttributeInfo(item *catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityForDevicesV1ResponseFunctionalCapabilityFunctionDetailsAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetFunctionalCapabilityByIDV1Item(item *catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityByIDV1ItemAttributeInfo(item.AttributeInfo)
	respItem["function_details"] = flattenDevicesGetFunctionalCapabilityByIDV1ItemFunctionDetails(item.FunctionDetails)
	respItem["function_name"] = item.FunctionName
	respItem["function_op_state"] = item.FunctionOpState
	respItem["id"] = item.ID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetFunctionalCapabilityByIDV1ItemAttributeInfo(item *catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityByIDV1ResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetFunctionalCapabilityByIDV1ItemFunctionDetails(items *[]catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityByIDV1ResponseFunctionDetails) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attribute_info"] = flattenDevicesGetFunctionalCapabilityByIDV1ItemFunctionDetailsAttributeInfo(item.AttributeInfo)
		respItem["id"] = item.ID
		respItem["property_name"] = item.PropertyName
		respItem["string_value"] = item.StringValue
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetFunctionalCapabilityByIDV1ItemFunctionDetailsAttributeInfo(item *catalystcentersdkgo.ResponseDevicesGetFunctionalCapabilityByIDV1ResponseFunctionDetailsAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
