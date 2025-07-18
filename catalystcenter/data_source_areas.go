package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAreas() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Site Design.

- Gets an area in the network hierarchy.
`,

		ReadContext: dataSourceAreasRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Area Id
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Aread Id. Read only.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Area name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name_hierarchy": &schema.Schema{
							Description: `Area hierarchical name. Read only.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"parent_id": &schema.Schema{
							Description: `Parent Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_hierarchy_id": &schema.Schema{
							Description: `Site Hierarchical Id. Read Only. Can be used to add the access groups using the API POST:/dna/system/api/v1/accessGroups, this value should be used to populate the srcResourceId field of the request payload.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"type": &schema.Schema{
							Description: `Site Type.
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

func dataSourceAreasRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsAnArea")
		vvID := vID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.SiteDesign.GetsAnArea(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsAnArea", err,
				"Failure at GetsAnArea, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsAnArea", err,
				"Failure at GetsAnArea, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetsAnAreaItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAnArea response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetsAnAreaItem(item *catalystcentersdkgo.ResponseSiteDesignGetsAnAreaResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["site_hierarchy_id"] = item.SiteHierarchyID
	respItem["name"] = item.Name
	respItem["name_hierarchy"] = item.NameHierarchy
	respItem["parent_id"] = item.ParentID
	respItem["type"] = item.Type
	return []map[string]interface{}{
		respItem,
	}
}
