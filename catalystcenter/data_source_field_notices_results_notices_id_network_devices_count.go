package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceFieldNoticesResultsNoticesIDNetworkDevicesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Compliance.

- Get count of field notice network devices for the notice
`,

		ReadContext: dataSourceFieldNoticesResultsNoticesIDNetworkDevicesCountRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. Id of the field notice
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. id of the network device
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"scan_status": &schema.Schema{
				Description: `scanStatus query parameter. status of the scan on the network device. Available values : NOT_SCANNED, IN_PROGRESS, SUCCESS, FAILED.
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

func dataSourceFieldNoticesResultsNoticesIDNetworkDevicesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vScanStatus, okScanStatus := d.GetOk("scan_status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCountOfFieldNoticeNetworkDevicesForTheNotice")
		vvID := vID.(string)
		queryParams1 := catalystcentersdkgo.GetCountOfFieldNoticeNetworkDevicesForTheNoticeQueryParams{}

		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okScanStatus {
			queryParams1.ScanStatus = vScanStatus.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Compliance.GetCountOfFieldNoticeNetworkDevicesForTheNotice(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfFieldNoticeNetworkDevicesForTheNotice", err,
				"Failure at GetCountOfFieldNoticeNetworkDevicesForTheNotice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCountOfFieldNoticeNetworkDevicesForTheNotice", err,
				"Failure at GetCountOfFieldNoticeNetworkDevicesForTheNotice, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeItem(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCountOfFieldNoticeNetworkDevicesForTheNotice response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeItem(item *catalystcentersdkgo.ResponseComplianceGetCountOfFieldNoticeNetworkDevicesForTheNoticeResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
