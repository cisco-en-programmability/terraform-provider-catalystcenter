package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEventCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Event Management.

- Get the count of registered events with provided eventIds or tags as mandatory
`,

		ReadContext: dataSourceEventCountRead,
		Schema: map[string]*schema.Schema{
			"event_id": &schema.Schema{
				Description: `eventId query parameter. The registered EventId should be provided
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": &schema.Schema{
				Description: `tags query parameter. The registered Tags should be provided
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"response": &schema.Schema{
							Description: `Response`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEventCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vEventID, okEventID := d.GetOk("event_id")
	vTags := d.Get("tags")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CountOfEventsV1")
		queryParams1 := catalystcentersdkgo.CountOfEventsV1QueryParams{}

		if okEventID {
			queryParams1.EventID = vEventID.(string)
		}
		queryParams1.Tags = vTags.(string)

		response1, restyResp1, err := client.EventManagement.CountOfEventsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CountOfEventsV1", err,
				"Failure at CountOfEventsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenEventManagementCountOfEventsV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountOfEventsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenEventManagementCountOfEventsV1Item(item *catalystcentersdkgo.ResponseEventManagementCountOfEventsV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	return []map[string]interface{}{
		respItem,
	}
}
