package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDeviceStatusDetails() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sensors.

- Get ICAP configuration status per network device using the activity ID, which was returned in property "taskId" of the
TaskResponse of the POST. For detailed information about the usage of the API, please refer to the Open API
specification document https://github.com/cisco-en-programmability/catalyst-center-api-
specs/blob/main/Assurance/CE_Cat_Center_Org-ICAP_APIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDeviceStatusDetailsRead,
		Schema: map[string]*schema.Schema{
			"preview_activity_id": &schema.Schema{
				Description: `previewActivityId path parameter. activity from the POST /deviceConfigugrationModels task response
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"network_device_id": &schema.Schema{
							Description: `Network Device Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"status": &schema.Schema{
							Description: `Status`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIcapSettingsConfigurationModelsPreviewActivityIDNetworkDeviceStatusDetailsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vPreviewActivityID := d.Get("preview_activity_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetICapConfigurationStatusPerNetworkDeviceV1")
		vvPreviewActivityID := vPreviewActivityID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Sensors.GetICapConfigurationStatusPerNetworkDeviceV1(vvPreviewActivityID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetICapConfigurationStatusPerNetworkDeviceV1", err,
				"Failure at GetICapConfigurationStatusPerNetworkDeviceV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenSensorsGetICapConfigurationStatusPerNetworkDeviceV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetICapConfigurationStatusPerNetworkDeviceV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSensorsGetICapConfigurationStatusPerNetworkDeviceV1Items(items *[]catalystcentersdkgo.ResponseSensorsGetICapConfigurationStatusPerNetworkDeviceV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["network_device_id"] = item.NetworkDeviceID
		respItem["status"] = item.Status
		respItems = append(respItems, respItem)
	}
	return respItems
}
