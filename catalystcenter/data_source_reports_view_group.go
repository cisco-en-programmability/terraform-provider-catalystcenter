package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceReportsViewGroup() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Reports.

- Gives a list of summary of all view groups.

- Gives a list of summary of all views in a viewgroup. Use "Get all view groups" API to get the viewGroupIds (required
as a query param for this API) for available viewgroups.
`,

		ReadContext: dataSourceReportsViewGroupRead,
		Schema: map[string]*schema.Schema{
			"view_group_id": &schema.Schema{
				Description: `viewGroupId path parameter. viewGroupId of viewgroup.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"view_group_id": &schema.Schema{
							Description: `viewgroup Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"views": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"description": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"view_id": &schema.Schema{
										Description: `Unique id for a view within viewgroup
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"view_name": &schema.Schema{
										Description: `view name
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

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"category": &schema.Schema{
							Description: `category of the view group
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Description: `view group description
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `name of view group
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"view_group_id": &schema.Schema{
							Description: `id of viewgroup
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

func dataSourceReportsViewGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vViewGroupID, okViewGroupID := d.GetOk("view_group_id")

	method1 := []bool{}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okViewGroupID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllViewGroupsV1")

		response1, restyResp1, err := client.Reports.GetAllViewGroupsV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllViewGroupsV1", err,
				"Failure at GetAllViewGroupsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenReportsGetAllViewGroupsV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllViewGroupsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetViewsForAGivenViewGroupV1")
		vvViewGroupID := vViewGroupID.(string)

		response2, restyResp2, err := client.Reports.GetViewsForAGivenViewGroupV1(vvViewGroupID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetViewsForAGivenViewGroupV1", err,
				"Failure at GetViewsForAGivenViewGroupV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenReportsGetViewsForAGivenViewGroupV1Item(response2)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetViewsForAGivenViewGroupV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenReportsGetAllViewGroupsV1Items(items *catalystcentersdkgo.ResponseReportsGetAllViewGroupsV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["category"] = item.Category
		respItem["description"] = item.Description
		respItem["name"] = item.Name
		respItem["view_group_id"] = item.ViewGroupID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenReportsGetViewsForAGivenViewGroupV1Item(item *catalystcentersdkgo.ResponseReportsGetViewsForAGivenViewGroupV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["view_group_id"] = item.ViewGroupID
	respItem["views"] = flattenReportsGetViewsForAGivenViewGroupV1ItemViews(item.Views)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenReportsGetViewsForAGivenViewGroupV1ItemViews(items *[]catalystcentersdkgo.ResponseReportsGetViewsForAGivenViewGroupV1Views) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["description"] = item.Description
		respItem["view_id"] = item.ViewID
		respItem["view_name"] = item.ViewName
		respItems = append(respItems, respItem)
	}
	return respItems
}
