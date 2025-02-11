package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceDeviceRebootApreboot() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Wireless.

- Users can query the access point reboot status using this intent API
`,

		ReadContext: dataSourceDeviceRebootAprebootRead,
		Schema: map[string]*schema.Schema{
			"parent_task_id": &schema.Schema{
				Description: `parentTaskId query parameter. task id of ap reboot request
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"ap_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"ap_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"failure_reason": &schema.Schema{
										Type:     schema.TypeString, //TEST,
										Computed: true,
									},

									"reboot_status": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"wlc_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceDeviceRebootAprebootRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vParentTaskID, okParentTaskID := d.GetOk("parent_task_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAccessPointRebootTaskResultV1")
		queryParams1 := catalystcentersdkgo.GetAccessPointRebootTaskResultV1QueryParams{}

		if okParentTaskID {
			queryParams1.ParentTaskID = vParentTaskID.(string)
		}

		response1, restyResp1, err := client.Wireless.GetAccessPointRebootTaskResultV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAccessPointRebootTaskResultV1", err,
				"Failure at GetAccessPointRebootTaskResultV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenWirelessGetAccessPointRebootTaskResultV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAccessPointRebootTaskResultV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenWirelessGetAccessPointRebootTaskResultV1Items(items *catalystcentersdkgo.ResponseWirelessGetAccessPointRebootTaskResultV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["wlc_ip"] = item.WlcIP
		respItem["ap_list"] = flattenWirelessGetAccessPointRebootTaskResultV1ItemsApList(item.ApList)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetAccessPointRebootTaskResultV1ItemsApList(items *[]catalystcentersdkgo.ResponseItemWirelessGetAccessPointRebootTaskResultV1ApList) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["ap_name"] = item.ApName
		respItem["reboot_status"] = item.RebootStatus
		respItem["failure_reason"] = flattenWirelessGetAccessPointRebootTaskResultV1ItemsApListFailureReason(item.FailureReason)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenWirelessGetAccessPointRebootTaskResultV1ItemsApListFailureReason(item *catalystcentersdkgo.ResponseItemWirelessGetAccessPointRebootTaskResultV1ApListFailureReason) interface{} {
	if item == nil {
		return nil
	}
	respItem := *item

	return responseInterfaceToString(respItem)

}
