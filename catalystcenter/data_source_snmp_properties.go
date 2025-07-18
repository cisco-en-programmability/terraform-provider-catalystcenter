package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceSNMPProperties() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Discovery.

- Returns SNMP properties
`,

		ReadContext: dataSourceSNMPPropertiesRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"id": &schema.Schema{
							Description: `Id of the SNMP Property
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Description: `[Deprecated] InstanceTenantId of the SNMP Property
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid of the SNMP Property. It is the same as the id. It will be deprecated in future version.
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"int_value": &schema.Schema{
							Description: `Integer Value of the SNMP 'Retry' or 'Timeout' property
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"system_property_name": &schema.Schema{
							Description: `Name of the SNMP Property as 'Retry' or 'Timeout'
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

func dataSourceSNMPPropertiesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSNMPProperties")

		// has_unknown_response: None

		response1, restyResp1, err := client.Discovery.GetSNMPProperties()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSNMPProperties", err,
				"Failure at GetSNMPProperties, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetSNMPProperties", err,
				"Failure at GetSNMPProperties, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDiscoveryGetSNMPPropertiesItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSNMPProperties response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDiscoveryGetSNMPPropertiesItems(items *[]catalystcentersdkgo.ResponseDiscoveryGetSNMPPropertiesResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["instance_tenant_id"] = item.InstanceTenantID
		respItem["instance_uuid"] = item.InstanceUUID
		respItem["int_value"] = item.IntValue
		respItem["system_property_name"] = item.SystemPropertyName
		respItems = append(respItems, respItem)
	}
	return respItems
}
