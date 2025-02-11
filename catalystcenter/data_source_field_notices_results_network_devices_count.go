package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesResultsNetworkDevicesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of field notice network devices
`,

		ReadContext: dataSourceFieldNoticesResultsNetworkDevicesCountRead,
		Schema: map[string]*schema.Schema{
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Id of the network device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"notice_count": &schema.Schema{
				Description: `noticeCount query parameter. Return network devices with noticeCount greater than this noticeCount
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"scan_status": &schema.Schema{
				Description: `scanStatus query parameter. Status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"count": &schema.Schema{
							Description: `Count`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceFieldNoticesResultsNetworkDevicesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vScanStatus, okScanStatus := d.GetOk("scan_status")
	vNoticeCount, okNoticeCount := d.GetOk("notice_count")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfFieldNoticeNetworkDevicesV1")
		queryParams1 := catalystcentersdkgo.GetCountOfFieldNoticeNetworkDevicesV1QueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okScanStatus {
			queryParams1.ScanStatus = vScanStatus.(string)
		}
		if okNoticeCount {
			queryParams1.NoticeCount = vNoticeCount.(float64)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Compliance.GetCountOfFieldNoticeNetworkDevicesV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfFieldNoticeNetworkDevicesV1", err,
				"Failure at GetCountOfFieldNoticeNetworkDevicesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfFieldNoticeNetworkDevicesV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfFieldNoticeNetworkDevicesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfFieldNoticeNetworkDevicesV1Item(item *catalystcentersdkgo.ResponseComplianceGetCountOfFieldNoticeNetworkDevicesV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
