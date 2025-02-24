package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceItsmCmdbSyncStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ITSM.

- This data source allows to retrieve the detail of CMDB sync status.It accepts two query parameter "status","date".The
supported values for status field are "Success","Failed","Unknown" and date field should be in "YYYY-MM-DD" format. By
default all the cmdb sync status will be send as response and based on the query parameter filtered detail will be send
as response.
`,

		ReadContext: dataSourceItsmCmdbSyncStatusRead,
		Schema: map[string]*schema.Schema{
			"date": &schema.Schema{
				Description: `date query parameter. Provide date in "YYYY-MM-DD" format
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Description: `status query parameter. Supported values are "Success","Failed" and "Unknown". Providing other values will result in all the available sync job status.
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"devices": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"device_id": &schema.Schema{
										Description: `Device Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"status": &schema.Schema{
										Description: `Status "Success" or "Failed"
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"failure_count": &schema.Schema{
							Description: `Failed device count
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"message": &schema.Schema{
							Description: `Message
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"success_count": &schema.Schema{
							Description: `Successfully synchronized device count
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"sync_time": &schema.Schema{
							Description: `Synchronization Time
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"unknown_error_count": &schema.Schema{
							Description: `Unknown error count
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

func dataSourceItsmCmdbSyncStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vStatus, okStatus := d.GetOk("status")
	vDate, okDate := d.GetOk("date")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetCmdbSyncStatusV1")
		queryParams1 := catalystcentersdkgo.GetCmdbSyncStatusV1QueryParams{}

		if okStatus {
			queryParams1.Status = vStatus.(string)
		}
		if okDate {
			queryParams1.Date = vDate.(string)
		}

		response1, restyResp1, err := client.Itsm.GetCmdbSyncStatusV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetCmdbSyncStatusV1", err,
				"Failure at GetCmdbSyncStatusV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenItsmGetCmdbSyncStatusV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetCmdbSyncStatusV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenItsmGetCmdbSyncStatusV1Items(items *catalystcentersdkgo.ResponseItsmGetCmdbSyncStatusV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["success_count"] = item.SuccessCount
		respItem["failure_count"] = item.FailureCount
		respItem["devices"] = flattenItsmGetCmdbSyncStatusV1ItemsDevices(item.Devices)
		respItem["unknown_error_count"] = item.UnknownErrorCount
		respItem["message"] = item.Message
		respItem["sync_time"] = item.SyncTime
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenItsmGetCmdbSyncStatusV1ItemsDevices(items *[]catalystcentersdkgo.ResponseItemItsmGetCmdbSyncStatusV1Devices) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["device_id"] = item.DeviceID
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
