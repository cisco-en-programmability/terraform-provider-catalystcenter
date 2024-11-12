package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceLanAutomationLog() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on LAN Automation.

- Invoke this API to get the LAN Automation session logs.

- Invoke this API to get the LAN Automation session logs based on the given LAN Automation session id.
`,

		ReadContext: dataSourceLanAutomationLogRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. LAN Automation session identifier.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. Number of LAN Automation sessions to be retrieved. Limit value can range between 1 to 10.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Starting index of the LAN Automation session. Minimum value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"entry": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_id": &schema.Schema{
										Description: `Device serial number for which the log message is associated.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"log_level": &schema.Schema{
										Description: `Supported levels are ERROR, INFO, WARNING, TRACE and CONFIG. 
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"record": &schema.Schema{
										Description: `Detailed log message.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"time_stamp": &schema.Schema{
										Description: `Time at which the log message is created.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"nw_orch_id": &schema.Schema{
							Description: `LAN Automation session identifier.
`,
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

						"entry": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_id": &schema.Schema{
										Description: `Device serial number for which the log message is associated.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"log_level": &schema.Schema{
										Description: `Supported levels are ERROR, INFO, WARNING, TRACE and CONFIG. 
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"record": &schema.Schema{
										Description: `Detailed log message.
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"time_stamp": &schema.Schema{
										Description: `Time at which the log message is created.
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"nw_orch_id": &schema.Schema{
							Description: `LAN Automation session identifier.
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

func dataSourceLanAutomationLogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vOffset, okOffset := d.GetOk("offset")
	vLimit, okLimit := d.GetOk("limit")
	vID, okID := d.GetOk("id")

	method1 := []bool{okOffset, okLimit}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: LanAutomationLogV1")
		queryParams1 := catalystcentersdkgo.LanAutomationLogV1QueryParams{}

		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}

		response1, restyResp1, err := client.LanAutomation.LanAutomationLogV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LanAutomationLogV1", err,
				"Failure at LanAutomationLogV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenLanAutomationLanAutomationLogV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomationLogV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: LanAutomationLogByIDV1")
		vvID := vID.(string)

		response2, restyResp2, err := client.LanAutomation.LanAutomationLogByIDV1(vvID)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 LanAutomationLogByIDV1", err,
				"Failure at LanAutomationLogByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenLanAutomationLanAutomationLogByIDV1Item(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting LanAutomationLogByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenLanAutomationLanAutomationLogV1Items(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationLogV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["nw_orch_id"] = item.NwOrchID
		respItem["entry"] = flattenLanAutomationLanAutomationLogV1ItemsEntry(item.Entry)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationLogV1ItemsEntry(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationLogV1ResponseEntry) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["log_level"] = item.LogLevel
		respItem["time_stamp"] = item.TimeStamp
		respItem["record"] = item.Record
		respItem["device_id"] = item.DeviceID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationLogByIDV1Item(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationLogByIDV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["nw_orch_id"] = item.NwOrchID
		respItem["entry"] = flattenLanAutomationLanAutomationLogByIDV1ItemEntry(item.Entry)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenLanAutomationLanAutomationLogByIDV1ItemEntry(items *[]catalystcentersdkgo.ResponseLanAutomationLanAutomationLogByIDV1ResponseEntry) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["log_level"] = item.LogLevel
		respItem["time_stamp"] = item.TimeStamp
		respItem["record"] = item.Record
		respItem["device_id"] = item.DeviceID
		respItems = append(respItems, respItem)
	}
	return respItems
}
