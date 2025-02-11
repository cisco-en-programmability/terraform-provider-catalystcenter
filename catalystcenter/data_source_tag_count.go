package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceTagCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Tag.

- Returns tag count
`,

		ReadContext: dataSourceTagCountRead,
		Schema: map[string]*schema.Schema{
			"attribute_name": &schema.Schema{
				Description: `attributeName query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name": &schema.Schema{
				Description: `name query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"name_space": &schema.Schema{
				Description: `nameSpace query parameter.`,
				Type:        schema.TypeString,
				Optional:    true,
			},
			"size": &schema.Schema{
				Description: `size query parameter. size in kilobytes(KB)
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

						"response": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"version": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceTagCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vName, okName := d.GetOk("name")
	vNameSpace, okNameSpace := d.GetOk("name_space")
	vAttributeName, okAttributeName := d.GetOk("attribute_name")
	vSize, okSize := d.GetOk("size")
	vSystemTag, okSystemTag := d.GetOk("system_tag")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetTagCountV1")
		queryParams1 := catalystcentersdkgo.GetTagCountV1QueryParams{}

		if okName {
			queryParams1.Name = vName.(string)
		}
		if okNameSpace {
			queryParams1.NameSpace = vNameSpace.(string)
		}
		if okAttributeName {
			queryParams1.AttributeName = vAttributeName.(string)
		}
		if okSize {
			queryParams1.Size = vSize.(string)
		}
		if okSystemTag {
			queryParams1.SystemTag = vSystemTag.(string)
		}

		response1, restyResp1, err := client.Tag.GetTagCountV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTagCountV1", err,
				"Failure at GetTagCountV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenTagGetTagCountV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTagCountV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenTagGetTagCountV1Item(item *catalystcentersdkgo.ResponseTagGetTagCountV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["version"] = item.Version
	respItem["response"] = item.Response
	return []map[string]interface{}{
		respItem,
	}
}
