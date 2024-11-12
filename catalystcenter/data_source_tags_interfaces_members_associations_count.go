package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagsInterfacesMembersAssociationsCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Tag.

- Fetches the count of interfaces that are associated with at least one tag. A tag is a user-defined or system-defined
construct to group resources. When an interface is tagged, it is called a member of the tag.
`,

		ReadContext: dataSourceTagsInterfacesMembersAssociationsCountRead,
		Schema: map[string]*schema.Schema{

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `The reported count.
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTagsInterfacesMembersAssociationsCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1")

		response1, restyResp1, err := client.Tag.RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1", err,
				"Failure at RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1Item(item *catalystcentersdkgo.ResponseTagRetrieveTheCountOfInterfacesThatAreAssociatedWithAtLeastOneTagV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
