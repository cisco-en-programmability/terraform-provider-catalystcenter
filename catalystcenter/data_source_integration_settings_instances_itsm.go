package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegrationSettingsInstancesItsm() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ITSM Integration.

- Fetches ITSM Integration setting by ID
`,

		ReadContext: dataSourceIntegrationSettingsInstancesItsmRead,
		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Description: `instanceId path parameter. Instance Id of the Integration setting instance
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"type_id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"data": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connection_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"auth_password": &schema.Schema{
													Description: `Auth Password`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"auth_user_name": &schema.Schema{
													Description: `Auth User Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},

												"url": &schema.Schema{
													Description: `Url`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
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

						"updated_date": &schema.Schema{
							Description: `Updated Date`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceIntegrationSettingsInstancesItsmRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vInstanceID := d.Get("instance_id")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetItsmIntegrationSettingByIDV1")
		vvInstanceID := vInstanceID.(string)

		response1, restyResp1, err := client.ItsmIntegration.GetItsmIntegrationSettingByIDV1(vvInstanceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetItsmIntegrationSettingByIDV1", err,
				"Failure at GetItsmIntegrationSettingByIDV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenItsmIntegrationGetItsmIntegrationSettingByIDV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetItsmIntegrationSettingByIDV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenItsmIntegrationGetItsmIntegrationSettingByIDV1Item(item *catalystcentersdkgo.ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type_id"] = item.TypeID
	respItem["id"] = item.ID
	respItem["dyp_id"] = item.DypID
	respItem["dyp_name"] = item.DypName
	respItem["dyp_major_version"] = item.DypMajorVersion
	respItem["name"] = item.Name
	respItem["unique_key"] = item.UniqueKey
	respItem["description"] = item.Description
	respItem["data"] = flattenItsmIntegrationGetItsmIntegrationSettingByIDV1ItemData(item.Data)
	respItem["updated_date"] = item.UpdatedDate
	respItem["updated_by"] = item.UpdatedBy
	respItem["tenant_id"] = item.TenantID
	return []map[string]interface{}{
		respItem,
	}
}

func flattenItsmIntegrationGetItsmIntegrationSettingByIDV1ItemData(item *catalystcentersdkgo.ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1Data) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["connection_settings"] = flattenItsmIntegrationGetItsmIntegrationSettingByIDV1ItemDataConnectionSettings(item.ConnectionSettings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenItsmIntegrationGetItsmIntegrationSettingByIDV1ItemDataConnectionSettings(item *catalystcentersdkgo.ResponseItsmIntegrationGetItsmIntegrationSettingByIDV1DataConnectionSettings) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	respItem["auth_user_name"] = item.AuthUserName
	respItem["auth_password"] = item.AuthPassword

	return []map[string]interface{}{
		respItem,
	}

}
