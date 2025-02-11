package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIntegrationSettingsStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on ITSM Integration.

- Fetches ITSM Integration status
`,

		ReadContext: dataSourceIntegrationSettingsStatusRead,
		Schema: map[string]*schema.Schema{

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"configurations": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"dyp_instance_id": &schema.Schema{
										Description: `DYP instance Id of the configuration
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"dyp_schema_name": &schema.Schema{
										Description: `DYP name of the configuration
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"id": &schema.Schema{
							Description: `Bundle Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"name": &schema.Schema{
							Description: `Bundle name
`,
							Type:     schema.TypeString,
							Computed: true,
						},

						"status": &schema.Schema{
							Description: `Bundle Status
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

func dataSourceIntegrationSettingsStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetItsmIntegrationStatusV1")

		response1, restyResp1, err := client.ItsmIntegration.GetItsmIntegrationStatusV1()

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetItsmIntegrationStatusV1", err,
				"Failure at GetItsmIntegrationStatusV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenItsmIntegrationGetItsmIntegrationStatusV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetItsmIntegrationStatusV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenItsmIntegrationGetItsmIntegrationStatusV1Items(items *[]catalystcentersdkgo.ResponseItsmIntegrationGetItsmIntegrationStatusV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["status"] = item.Status
		respItem["configurations"] = flattenItsmIntegrationGetItsmIntegrationStatusV1ItemsConfigurations(item.Configurations)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenItsmIntegrationGetItsmIntegrationStatusV1ItemsConfigurations(items *[]catalystcentersdkgo.ResponseItsmIntegrationGetItsmIntegrationStatusV1ResponseConfigurations) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["dyp_schema_name"] = item.DypSchemaName
		respItem["dyp_instance_id"] = item.DypInstanceID
		respItems = append(respItems, respItem)
	}
	return respItems
}
