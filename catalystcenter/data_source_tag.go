package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTag() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Tag.

- Returns the tags for given filter criteria

- Returns tag specified by Id
`,

		ReadContext: dataSourceTagRead,
		Schema: map[string]*schema.Schema{
			"additional_info_attributes": &schema.Schema{
				Description: `additionalInfo.attributes query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"additional_info_name_space": &schema.Schema{
				Description: `additionalInfo.nameSpace query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"field": &schema.Schema{
				Description: `field query parameter. Available field names are :'name,id,parentId,type,additionalInfo.nameSpace,additionalInfo.attributes'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Tag ID
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"level": &schema.Schema{
				Description: `level query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. Tag name is mandatory when filter operation is used.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter.`,
				Type:        schema.TypeFloat,
				Optional:    true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. Available values are asc and des
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. size in kilobytes(KB)
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. Only supported attribute is name. SortyBy is mandatory when order is used.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"system_tag": &schema.Schema{
				Description: `systemTag query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"values": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"system_tag": &schema.Schema{
							// Type:     schema.TypeBool,
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

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"dynamic_rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"member_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"rules": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"items": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},

												"name": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"operation": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"value": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},

												"values": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"system_tag": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTagRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vAdditionalInfonameSpace, okAdditionalInfonameSpace := d.GetOk("additional_info_name_space")
	vAdditionalInfoattributes, okAdditionalInfoattributes := d.GetOk("additional_info_attributes")
	vLevel, okLevel := d.GetOk("level")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vSize, okSize := d.GetOk("size")
	vField, okField := d.GetOk("field")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vSystemTag, okSystemTag := d.GetOk("system_tag")
	vID, okID := d.GetOk("id")

	method1 := []bool{okName, okAdditionalInfonameSpace, okAdditionalInfoattributes, okLevel, okOffset, okLimit, okSize, okField, okSortBy, okOrder, okSystemTag}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTagV1")
		queryParams1 := catalystcentersdkgo.GetTagV1QueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okAdditionalInfonameSpace {
			queryParams1.AdditionalInfonameSpace = vAdditionalInfonameSpace.(string)
		}
		if okAdditionalInfoattributes {
			queryParams1.AdditionalInfoattributes = vAdditionalInfoattributes.(string)
		}
		if okLevel {
			queryParams1.Level = vLevel.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okSize {
			queryParams1.Size = vSize.(string)
		}
		if okField {
			queryParams1.Field = vField.(string)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}
		if okSystemTag {
			queryParams1.SystemTag = vSystemTag.(string)
		}

		response1, restyResp1, err := client.Tag.GetTagV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTagV1", err,
				"Failure at GetTagV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTagGetTagV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetTagByIDV1")
		vvID := vID.(string)

		response2, restyResp2, err := client.Tag.GetTagByIDV1(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTagByIDV1", err,
				"Failure at GetTagByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenTagGetTagByIDV1Item(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagGetTagV1Items(items *[]catalystcentersdkgo.ResponseTagGetTagV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["system_tag"] = boolPtrToString(item.SystemTag)
		respItem["description"] = item.Description
		respItem["dynamic_rules"] = flattenTagGetTagV1ItemsDynamicRules(item.DynamicRules)
		respItem["name"] = item.Name
		respItem["id"] = item.ID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTagGetTagV1ItemsDynamicRules(items *[]catalystcentersdkgo.ResponseTagGetTagV1ResponseDynamicRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["member_type"] = item.MemberType
		respItem["rules"] = flattenTagGetTagV1ItemsDynamicRulesRules(item.Rules)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTagGetTagV1ItemsDynamicRulesRules(item *catalystcentersdkgo.ResponseTagGetTagV1ResponseDynamicRulesRules) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["values"] = item.Values
	respItem["items"] = item.Items
	respItem["operation"] = item.Operation
	respItem["name"] = item.Name
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}

func flattenTagGetTagByIDV1Item(item *catalystcentersdkgo.ResponseTagGetTagByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["system_tag"] = boolPtrToString(item.SystemTag)
	respItem["description"] = item.Description
	respItem["dynamic_rules"] = flattenTagGetTagByIDV1ItemDynamicRules(item.DynamicRules)
	respItem["name"] = item.Name
	respItem["id"] = item.ID
	respItem["instance_tenant_id"] = item.InstanceTenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenTagGetTagByIDV1ItemDynamicRules(items *[]catalystcentersdkgo.ResponseTagGetTagByIDV1ResponseDynamicRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["member_type"] = item.MemberType
		respItem["rules"] = flattenTagGetTagByIDV1ItemDynamicRulesRules(item.Rules)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTagGetTagByIDV1ItemDynamicRulesRules(item *catalystcentersdkgo.ResponseTagGetTagByIDV1ResponseDynamicRulesRules) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["values"] = item.Values
	respItem["items"] = item.Items
	respItem["operation"] = item.Operation
	respItem["name"] = item.Name
	respItem["value"] = item.Value

	return []map[string]interface{}{
		respItem,
	}

}
