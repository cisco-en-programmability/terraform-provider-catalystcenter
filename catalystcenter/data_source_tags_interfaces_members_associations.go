package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagsInterfacesMembersAssociations() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Tag.

- Fetches the tags associated with the interfaces. Interfaces that don't have any tags associated will not be included
in the response. A tag is a user-defined or system-defined construct to group resources. When an interface is tagged, it
is called a member of the tag.
`,

		ReadContext: dataSourceTagsInterfacesMembersAssociationsRead,
		Schema: map[string]*schema.Schema{
			"limit": &schema.Schema{
				Description: `limit query parameter. The number of records to show for this page. minimum: 1, maximum: 500
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The first record to show for this page; the first record is numbered 1. minimum: 1
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id of the member (network device or interface)
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"tags": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Tag id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Tag name
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

func dataSourceTagsInterfacesMembersAssociationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTagsAssociatedWithTheInterfacesV1")
		queryParams1 := catalystcentersdkgo.RetrieveTagsAssociatedWithTheInterfacesV1QueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Tag.RetrieveTagsAssociatedWithTheInterfacesV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTagsAssociatedWithTheInterfacesV1", err,
				"Failure at RetrieveTagsAssociatedWithTheInterfacesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTagRetrieveTagsAssociatedWithTheInterfacesV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTagsAssociatedWithTheInterfacesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagRetrieveTagsAssociatedWithTheInterfacesV1Items(items *[]catalystcentersdkgo.ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["tags"] = flattenTagRetrieveTagsAssociatedWithTheInterfacesV1ItemsTags(item.Tags)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenTagRetrieveTagsAssociatedWithTheInterfacesV1ItemsTags(items *[]catalystcentersdkgo.ResponseTagRetrieveTagsAssociatedWithTheInterfacesV1ResponseTags) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItems = append(respItems, respItem)
	}
	return respItems
}
