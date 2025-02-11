package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaMulticastV1() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Returns a list of multicast configurations at a fabric site level that match the provided query parameters.
`,

		ReadContext: dataSourceSdaMulticastV1Read,
		Schema: map[string]*schema.Schema{
			"fabric_id": &schema.Schema{
				Description: `fabricId query parameter. ID of the fabric site where multicast is configured.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Maximum number of records to return.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting record for pagination.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"fabric_id": &schema.Schema{
							Description: `ID of the fabric site.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"replication_mode": &schema.Schema{
							Description: `Replication Mode deployed in the fabric site.
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

func dataSourceSdaMulticastV1Read(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vFabricID, okFabricID := d.GetOk("fabric_id")
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetMulticastV1")
		queryParams1 := catalystcentersdkgo.GetMulticastV1QueryParams{}

		if okFabricID {
			queryParams1.FabricID = vFabricID.(string)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.Sda.GetMulticastV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetMulticastV1", err,
				"Failure at GetMulticastV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSdaGetMulticastV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetMulticastV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetMulticastV1Items(items *[]catalystcentersdkgo.ResponseSdaGetMulticastV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["fabric_id"] = item.FabricID
		respItem["replication_mode"] = item.ReplicationMode
		respItems = append(respItems, respItem)
	}
	return respItems
}
