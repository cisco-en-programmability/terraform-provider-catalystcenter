package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSdaDeviceRole() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on SDA.

- Get device role in SDA Fabric
`,

		ReadContext: dataSourceSdaDeviceRoleRead,
		Schema: map[string]*schema.Schema{
			"device_management_ip_address": &schema.Schema{
				Description: `deviceManagementIpAddress query parameter. Device Management IP Address
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Device role successfully retrieved from sda fabric.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"roles": &schema.Schema{
							Description: `Assigned device roles. Possible roles are [Edge Node, Control Plane, Border Node, Extended Node, Wireless Controller, Transit Control Plane]
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"status": &schema.Schema{
							Description: `Status indicates if API failed or passed.
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

func dataSourceSdaDeviceRoleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vDeviceManagementIPAddress := d.Get("device_management_ip_address")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetDeviceRoleInSdaFabric")
		queryParams1 := catalystcentersdkgo.GetDeviceRoleInSdaFabricQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress.(string)

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.GetDeviceRoleInSdaFabric(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceRoleInSdaFabric", err,
				"Failure at GetDeviceRoleInSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetDeviceRoleInSdaFabric", err,
				"Failure at GetDeviceRoleInSdaFabric, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetDeviceRoleInSdaFabricItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetDeviceRoleInSdaFabric response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSdaGetDeviceRoleInSdaFabricItem(item *catalystcentersdkgo.ResponseSdaGetDeviceRoleInSdaFabric) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["roles"] = item.Roles
	respItem["status"] = item.Status
	respItem["description"] = item.Description
	return []map[string]interface{}{
		respItem,
	}
}
