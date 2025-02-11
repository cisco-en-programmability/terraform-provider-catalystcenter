package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceNetworkDeviceConfigFilesID() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Configuration Archive.

- Retrieves the details of a specific network device configuration file using the id.
`,

		ReadContext: dataSourceNetworkDeviceConfigFilesIDRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Description: `id path parameter. The value of id can be obtained from the response of API /dna/intent/api/v1/networkDeviceConfigFiles
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"created_by": &schema.Schema{
							Description: `The entity responsible for creating the configuration changes.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"created_time": &schema.Schema{
							Description: `The UNIX epoch timestamp in milliseconds marking when the resource was created.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"file_type": &schema.Schema{
							Description: `Type of configuration file. Config File Type can be 'RUNNINGCONFIG' or 'STARTUPCONFIG' or 'VLAN'.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Description: `Unique identifier (UUID) of the configuration file.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"network_device_id": &schema.Schema{
							Description: `Unique identifier (UUID) of the network devices.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"version_id": &schema.Schema{
							Description: `The version unique identifier triggered after any config change.
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

func dataSourceNetworkDeviceConfigFilesIDRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID := d.Get("id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetConfigurationFileDetailsByIDV1")
		vvID := vID.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.ConfigurationArchive.GetConfigurationFileDetailsByIDV1(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetConfigurationFileDetailsByIDV1", err,
				"Failure at GetConfigurationFileDetailsByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenConfigurationArchiveGetConfigurationFileDetailsByIDV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetConfigurationFileDetailsByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenConfigurationArchiveGetConfigurationFileDetailsByIDV1Item(item *catalystcentersdkgo.ResponseConfigurationArchiveGetConfigurationFileDetailsByIDV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["network_device_id"] = item.NetworkDeviceID
	respItem["version_id"] = item.VersionID
	respItem["file_type"] = item.FileType
	respItem["created_by"] = item.CreatedBy
	respItem["created_time"] = item.CreatedTime
	return []map[string]interface{}{
		respItem,
	}
}
