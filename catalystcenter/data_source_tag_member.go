package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagMember() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Tag.

- Returns tag members specified by id
`,

		ReadContext: dataSourceTagMemberRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Tag ID
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"level": &schema.Schema{
				Description: `level query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Used to Number of maximum members to return in the result
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"member_association_type": &schema.Schema{
				Description: `memberAssociationType query parameter. Indicates how the member is associated with the tag. Possible values and description. 1) DYNAMIC : The member is associated to the tag through rules. 2) STATIC – The member is associated to the tag manually. 3) MIXED – The member is associated manually and also satisfies the rule defined for the tag
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"member_type": &schema.Schema{
				Description: `memberType query parameter. Entity type of the member. Possible values can be retrieved by using /tag/member/type API
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Used for pagination. It indicates the starting row number out of available member records
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTagMemberRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vMemberType := d.Get("member_type")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vMemberAssociationType, okMemberAssociationType := d.GetOk("member_association_type")
	vLevel, okLevel := d.GetOk("level")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTagMembersByIDV1")
		vvID := vID.(string)
		queryParams1 := catalystcentersdkgo.GetTagMembersByIDV1QueryParams{}

		queryParams1.MemberType = vMemberType.(string)

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okMemberAssociationType {
			queryParams1.MemberAssociationType = vMemberAssociationType.(string)
		}
		if okLevel {
			queryParams1.Level = vLevel.(string)
		}

		response1, restyResp1, err := client.Tag.GetTagMembersByIDV1(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTagMembersByIDV1", err,
				"Failure at GetTagMembersByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenTagGetTagMembersByIDV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagMembersByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagGetTagMembersByIDV1Items(items *[]catalystcentersdkgo.ResponseTagGetTagMembersByIDV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["instance_uuid"] = item.InstanceUUID
		respItems = append(respItems, respItem)
	}
	return respItems
}
