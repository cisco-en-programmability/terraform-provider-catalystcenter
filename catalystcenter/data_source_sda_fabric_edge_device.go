package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaFabricEdgeDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get edge device from SDA Fabric
`,

		ReadContext: dataSourceSdaFabricEdgeDeviceRead,
		Schema: map[string]*schema.Schema{
			"device_management_ip_address": &schema.Schema{
				Description: `deviceManagementIpAddress query parameter.`,
				Type:        schema.TypeString,
				Required:    true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Edge device info retrieved successfully in sda fabric
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the Device which is provisioned successfully
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"device_name": &schema.Schema{
							Description: `Device Name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"fabric_site_name_hierarchy": &schema.Schema{
							Description: `Fabric Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"roles": &schema.Schema{
							Description: `Assigned roles
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Status
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

func dataSourceSdaFabricEdgeDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceManagementIPAddress := d.Get("device_management_ip_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetEdgeDeviceFromSdaFabricV1")
		queryParams1 := catalystcentersdkgo.GetEdgeDeviceFromSdaFabricV1QueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		response1, restyResp1, err := client.Sda.GetEdgeDeviceFromSdaFabricV1(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetEdgeDeviceFromSdaFabricV1", err,
				"Failure at GetEdgeDeviceFromSdaFabricV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetEdgeDeviceFromSdaFabricV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEdgeDeviceFromSdaFabricV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetEdgeDeviceFromSdaFabricV1Item(item *catalystcentersdkgo.ResponseSdaGetEdgeDeviceFromSdaFabricV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["device_management_ip_address"] = item.DeviceManagementIPAddress
	respItem["device_name"] = item.DeviceName
	respItem["roles"] = item.Roles
	respItem["site_name_hierarchy"] = item.SiteNameHierarchy
	respItem["fabric_site_name_hierarchy"] = item.FabricSiteNameHierarchy
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	return []map[string]interface{}{
		respItem,
	}
}
