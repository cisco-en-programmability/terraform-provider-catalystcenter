package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceConfigFilesCount() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Archive.

- Retrieves count the details of the network device configuration files.
`,

		ReadContext: dataSourceNetworkDeviceConfigFilesCountRead,
		Schema: map[string]*schema.Schema{
			"file_type": &schema.Schema{
				Description: `fileType query parameter. Type of device configuration file. Available values : 'RUNNINGCONFIG', 'STARTUPCONFIG', 'VLAN'
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id query parameter. Unique identifier (UUID) of the configuration file.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_device_id": &schema.Schema{
				Description: `networkDeviceId query parameter. Unique identifier (UUID) of the network devices. The number of networkDeviceId(s) must not exceed 5.
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

func dataSourceNetworkDeviceConfigFilesCountRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vNetworkDeviceID, okNetworkDeviceID := d.GetOk("network_device_id")
	vFileType, okFileType := d.GetOk("file_type")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: CountOfNetworkDeviceConfigurationFilesV1")
		queryParams1 := catalystcentersdkgo.CountOfNetworkDeviceConfigurationFilesV1QueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okNetworkDeviceID {
			queryParams1.NetworkDeviceID = vNetworkDeviceID.(string)
		}
		if okFileType {
			queryParams1.FileType = vFileType.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.ConfigurationArchive.CountOfNetworkDeviceConfigurationFilesV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 CountOfNetworkDeviceConfigurationFilesV1", err,
				"Failure at CountOfNetworkDeviceConfigurationFilesV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting CountOfNetworkDeviceConfigurationFilesV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1Item(item *catalystcentersdkgo.ResponseConfigurationArchiveCountOfNetworkDeviceConfigurationFilesV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["count"] = item.Count
	return []map[string]interface{}{
		respItem,
	}
}
