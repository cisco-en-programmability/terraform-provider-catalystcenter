package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegrationSettingsItsmInstances() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ITSM Integration.

- Fetches all ITSM Integration settings
`,

		ReadContext: dataSourceIntegrationSettingsItsmInstancesRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"created_by": &schema.Schema{
							Description: `Created By`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"created_date": &schema.Schema{
							Description: `Created Date`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"dyp_id": &schema.Schema{
							Description: `Dyp Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"dyp_major_version": &schema.Schema{
							Description: `Dyp Major Version`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"dyp_name": &schema.Schema{
							Description: `Dyp Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"schema_version": &schema.Schema{
							Description: `Schema Version`,
							Type:        schema.TypeFloat,
							Computed:    true,
						},

						"software_version_log": &schema.Schema{
							Description: `Software Version Log`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},

						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"unique_key": &schema.Schema{
							Description: `Unique Key`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"updated_by": &schema.Schema{
							Description: `Updated By`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIntegrationSettingsItsmInstancesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllItsmIntegrationSettingsV1")

		response1, restyResp1, err := client.ItsmIntegration.GetAllItsmIntegrationSettingsV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllItsmIntegrationSettingsV1", err,
				"Failure at GetAllItsmIntegrationSettingsV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenItsmIntegrationGetAllItsmIntegrationSettingsV1Items(response1)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllItsmIntegrationSettingsV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenItsmIntegrationGetAllItsmIntegrationSettingsV1Items(items *catalystcentersdkgo.ResponseItsmIntegrationGetAllItsmIntegrationSettingsV1) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["dyp_id"] = item.DypID
		respItem["dyp_name"] = item.DypName
		respItem["name"] = item.Name
		respItem["unique_key"] = item.UniqueKey
		respItem["dyp_major_version"] = item.DypMajorVersion
		respItem["description"] = item.Description
		respItem["created_date"] = item.CreatedDate
		respItem["created_by"] = item.CreatedBy
		respItem["updated_by"] = item.UpdatedBy
		respItem["software_version_log"] = flattenItsmIntegrationGetAllItsmIntegrationSettingsV1ItemsSoftwareVersionLog(item.SoftwareVersionLog)
		respItem["schema_version"] = item.SchemaVersion
		respItem["tenant_id"] = item.TenantID
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenItsmIntegrationGetAllItsmIntegrationSettingsV1ItemsSoftwareVersionLog(items *[]catalystcentersdkgo.ResponseItemItsmIntegrationGetAllItsmIntegrationSettingsV1SoftwareVersionLog) []interface{} {
	if items == nil {
		return nil
	}
	var respItems []interface{}
	for _, item := range *items {
		respItem := item
		respItems = append(respItems, responseInterfaceToString(respItem))
	}
	return respItems
}
