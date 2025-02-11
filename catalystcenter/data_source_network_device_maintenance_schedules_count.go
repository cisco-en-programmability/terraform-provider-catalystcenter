package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceMaintenanceSchedulesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Retrieve the total count of all scheduled maintenance windows for network devices.
`,

		ReadContext: dataSourceNetworkDeviceMaintenanceSchedulesCountRead,
		Schema: map[string]*schema.Schema{
			"network_device_ids": &schema.Schema{
				Description: `networkDeviceIds query parameter. List of network device ids.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. The status of the maintenance schedule. Possible values are: UPCOMING, IN_PROGRESS, COMPLETED, FAILED. Refer features for more details.
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
							Description: `Count of scheduled maintenance windows
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

func dataSourceNetworkDeviceMaintenanceSchedulesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vNetworkDeviceIDs, okNetworkDeviceIDs := d.GetOk("network_device_ids")
	vStatus, okStatus := d.GetOk("status")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1")
		queryParams1 := catalystcentersdkgo.RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1QueryParams{}

		if okNetworkDeviceIDs {
			queryParams1.NetworkDeviceIDs = vNetworkDeviceIDs.(string)
		}
		if okStatus {
			queryParams1.Status = vStatus.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Devices.RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1", err,
				"Failure at RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting RetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1Item(item *catalystcentersdkgo.ResponseDevicesRetrieveTheTotalNumberOfScheduledMaintenanceWindowsV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
