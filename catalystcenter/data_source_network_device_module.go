package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceModule() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Returns modules by specified device id. The API returns a paginated response based on 'limit' and 'offset' parameters,
allowing up to 500 records per page. 'limit' specifies the number of records, and 'offset' sets the starting point using
1-based indexing. Use /dna/intent/api/v1/network-device/module/count API to get the total record count. For data sets
over 500 records, make multiple calls, adjusting 'limit' and 'offset' to retrieve all records incrementally.

- Returns Module info by 'module id'
`,

		ReadContext: dataSourceNetworkDeviceModuleRead,
		Schema: map[string]*schema.Schema{
			"device_id": &schema.Schema{
				Description: `deviceId query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Module id
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. Min: 1, Max: 500
`,
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name_list": &schema.Schema{
				Description: `nameList query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeInt,
				Optional:    true,
			},
			"operational_state_code_list": &schema.Schema{
				Description: `operationalStateCodeList query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"part_number_list": &schema.Schema{
				Description: `partNumberList query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"vendor_equipment_type_list": &schema.Schema{
				Description: `vendorEquipmentTypeList query parameter.`,
				Type:        schema.TypeList,
				Optional:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"assembly_number": &schema.Schema{
							Description: `Assembly number of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"assembly_revision": &schema.Schema{
							Description: `Assembly revision of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"attribute_info": &schema.Schema{
							Description: `Deprecated
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"containment_entity": &schema.Schema{
							Description: `Containment entity of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"entity_physical_index": &schema.Schema{
							Description: `Entity physical index of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Id of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_field_replaceable": &schema.Schema{
							Description: `To mention if field is replaceable
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_reporting_alarms_allowed": &schema.Schema{
							Description: `To mention if reporting alarms are allowed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"manufacturer": &schema.Schema{
							Description: `Manufacturer of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"module_index": &schema.Schema{
							Description: `Index of the module
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"operational_state_code": &schema.Schema{
							Description: `Operational state of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"part_number": &schema.Schema{
							Description: `Part number of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial number of the modules
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vendor_equipment_type": &schema.Schema{
							Description: `Vendor equipment type of the module
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

						"assembly_number": &schema.Schema{
							Description: `Assembly Number of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"assembly_revision": &schema.Schema{
							Description: `Assembly Revision of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"attribute_info": &schema.Schema{
							Description: `Deprecated
`,
							Type:     schema.TypeString, //TEST,
							Computed: true,
						},

						"containment_entity": &schema.Schema{
							Description: `Containment Entity of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `Description of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"entity_physical_index": &schema.Schema{
							Description: `Entity Physical Index of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `ID of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_field_replaceable": &schema.Schema{
							Description: `To mention if field is replaceable
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_reporting_alarms_allowed": &schema.Schema{
							Description: `To mention if reporting alarms are allowed
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"manufacturer": &schema.Schema{
							Description: `Manufacturer of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"module_index": &schema.Schema{
							Description: `Index of the module
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Name of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"operational_state_code": &schema.Schema{
							Description: `Operational state of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"part_number": &schema.Schema{
							Description: `Part number of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"serial_number": &schema.Schema{
							Description: `Serial number of the module
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"vendor_equipment_type": &schema.Schema{
							Description: `Vendor euipment type of the module
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

func dataSourceNetworkDeviceModuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceID, okDeviceID := d.GetOk("device_id")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vNameList, okNameList := d.GetOk("name_list")
	vVendorEquipmentTypeList, okVendorEquipmentTypeList := d.GetOk("vendor_equipment_type_list")
	vPartNumberList, okPartNumberList := d.GetOk("part_number_list")
	vOperationalStateCodeList, okOperationalStateCodeList := d.GetOk("operational_state_code_list")
	vID, okID := d.GetOk("id")

	method1 := []bool{okDeviceID, okLimit, okOffset, okNameList, okVendorEquipmentTypeList, okPartNumberList, okOperationalStateCodeList}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetModules")
		queryParams1 := catalystcentersdkgo.GetModulesQueryParams{}

		if okDeviceID {
			queryParams1.DeviceID = vDeviceID.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(int)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(int)
		}
		if okNameList {
			queryParams1.NameList = interfaceToSliceString(vNameList)
		}
		if okVendorEquipmentTypeList {
			queryParams1.VendorEquipmentTypeList = interfaceToSliceString(vVendorEquipmentTypeList)
		}
		if okPartNumberList {
			queryParams1.PartNumberList = interfaceToSliceString(vPartNumberList)
		}
		if okOperationalStateCodeList {
			queryParams1.OperationalStateCodeList = interfaceToSliceString(vOperationalStateCodeList)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Devices.GetModules(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetModules", err,
				"Failure at GetModules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetModules", err,
				"Failure at GetModules, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetModulesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetModules response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetModuleInfoByID")
		vvID := vID.(string)

		// has_unknown_response: None

		response2, restyResp2, err := client.Devices.GetModuleInfoByID(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetModuleInfoByID", err,
				"Failure at GetModuleInfoByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetModuleInfoByID", err,
				"Failure at GetModuleInfoByID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenDevicesGetModuleInfoByIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetModuleInfoByID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetModulesItems(items *[]catalystcentersdkgo.ResponseDevicesGetModulesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["assembly_number"] = item.AssemblyNumber
		respItem["assembly_revision"] = item.AssemblyRevision
		respItem["attribute_info"] = flattenDevicesGetModulesItemsAttributeInfo(item.AttributeInfo)
		respItem["containment_entity"] = item.ContainmentEntity
		respItem["description"] = item.Description
		respItem["entity_physical_index"] = item.EntityPhysicalIndex
		respItem["id"] = item.ID
		respItem["is_field_replaceable"] = item.IsFieldReplaceable
		respItem["is_reporting_alarms_allowed"] = item.IsReportingAlarmsAllowed
		respItem["manufacturer"] = item.Manufacturer
		respItem["module_index"] = item.ModuleIndex
		respItem["name"] = item.Name
		respItem["operational_state_code"] = item.OperationalStateCode
		respItem["part_number"] = item.PartNumber
		respItem["serial_number"] = item.SerialNumber
		respItem["vendor_equipment_type"] = item.VendorEquipmentType
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetModulesItemsAttributeInfo(item *catalystcentersdkgo.ResponseDevicesGetModulesResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}

func flattenDevicesGetModuleInfoByIDItem(item *catalystcentersdkgo.ResponseDevicesGetModuleInfoByIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["assembly_number"] = item.AssemblyNumber
	respItem["assembly_revision"] = item.AssemblyRevision
	respItem["attribute_info"] = flattenDevicesGetModuleInfoByIDItemAttributeInfo(item.AttributeInfo)
	respItem["containment_entity"] = item.ContainmentEntity
	respItem["description"] = item.Description
	respItem["entity_physical_index"] = item.EntityPhysicalIndex
	respItem["id"] = item.ID
	respItem["is_field_replaceable"] = item.IsFieldReplaceable
	respItem["is_reporting_alarms_allowed"] = item.IsReportingAlarmsAllowed
	respItem["manufacturer"] = item.Manufacturer
	respItem["module_index"] = item.ModuleIndex
	respItem["name"] = item.Name
	respItem["operational_state_code"] = item.OperationalStateCode
	respItem["part_number"] = item.PartNumber
	respItem["serial_number"] = item.SerialNumber
	respItem["vendor_equipment_type"] = item.VendorEquipmentType
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesGetModuleInfoByIDItemAttributeInfo(item *catalystcentersdkgo.ResponseDevicesGetModuleInfoByIDResponseAttributeInfo) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
