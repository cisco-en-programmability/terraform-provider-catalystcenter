package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

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
		log.Printf("[DEBUG] Selected method: GetsAnAreaV1")
		vvID := vID.(string)

		response1, restyResp1, err := client.SiteDesign.GetsAnAreaV1(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetsAnAreaV1", err,
				"Failure at GetsAnAreaV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetsAnAreaV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAnAreaV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSiteDesignGetsAnAreaV1Item(item *catalystcentersdkgo.ResponseSiteDesignGetsAnAreaV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["name_hierarchy"] = item.NameHierarchy
	respItem["parent_id"] = item.ParentID
	respItem["type"] = item.Type
	return []map[string]interface{}{
		respItem,
	}
}
