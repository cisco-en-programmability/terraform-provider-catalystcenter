package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceMapsImportStatus() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Sites.

- Gets the status of a map archive import operation. For a map archive import that has just been initiated, will provide
the result of validation of the archive and a pre-import preview of what will be performed if the import is performed.
Once an import is requested to be performed, this API will give the status of the import and upon completion a post-
import summary of what was performed by the operation.
`,

		ReadContext: dataSourceMapsImportStatusRead,
		Schema: map[string]*schema.Schema{
			"import_context_uuid": &schema.Schema{
				Description: `importContextUuid path parameter. The unique import context UUID given by a previous and recent call to maps/import/start API
`,
				Type:     schema.TypeString,
				Required: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"audit_log": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"children": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"entities_count": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"entity_name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"entity_type": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"error_entities_count": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"errors": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"message": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"infos": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"message": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"matching_entities_count": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"sub_tasks_root_task_id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},

									"successfully_imported_floors": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},

									"warnings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"message": &schema.Schema{
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
								},
							},
						},

						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"uuid": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"least_significant_bits": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},

									"most_significant_bits": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceMapsImportStatusRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vImportContextUUID := d.Get("import_context_uuid")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: ImportMapArchiveImportStatusV1")
		vvImportContextUUID := vImportContextUUID.(string)

		response1, restyResp1, err := client.Sites.ImportMapArchiveImportStatusV1(vvImportContextUUID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 ImportMapArchiveImportStatusV1", err,
				"Failure at ImportMapArchiveImportStatusV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSitesImportMapArchiveImportStatusV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting ImportMapArchiveImportStatusV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenSitesImportMapArchiveImportStatusV1Item(item *catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["audit_log"] = flattenSitesImportMapArchiveImportStatusV1ItemAuditLog(item.AuditLog)
	respItem["status"] = item.Status
	respItem["uuid"] = flattenSitesImportMapArchiveImportStatusV1ItemUUID(item.UUID)
	return []map[string]interface{}{
		respItem,
	}
}

func flattenSitesImportMapArchiveImportStatusV1ItemAuditLog(item *catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1AuditLog) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["children"] = item.Children
	respItem["entities_count"] = flattenSitesImportMapArchiveImportStatusV1ItemAuditLogEntitiesCount(item.EntitiesCount)
	respItem["entity_name"] = item.EntityName
	respItem["entity_type"] = item.EntityType
	respItem["error_entities_count"] = flattenSitesImportMapArchiveImportStatusV1ItemAuditLogErrorEntitiesCount(item.ErrorEntitiesCount)
	respItem["errors"] = flattenSitesImportMapArchiveImportStatusV1ItemAuditLogErrors(item.Errors)
	respItem["infos"] = flattenSitesImportMapArchiveImportStatusV1ItemAuditLogInfos(item.Infos)
	respItem["matching_entities_count"] = flattenSitesImportMapArchiveImportStatusV1ItemAuditLogMatchingEntitiesCount(item.MatchingEntitiesCount)
	respItem["sub_tasks_root_task_id"] = item.SubTasksRootTaskID
	respItem["successfully_imported_floors"] = item.SuccessfullyImportedFloors
	respItem["warnings"] = flattenSitesImportMapArchiveImportStatusV1ItemAuditLogWarnings(item.Warnings)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenSitesImportMapArchiveImportStatusV1ItemAuditLogEntitiesCount(items *[]catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1AuditLogEntitiesCount) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesImportMapArchiveImportStatusV1ItemAuditLogErrorEntitiesCount(items *[]catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1AuditLogErrorEntitiesCount) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesImportMapArchiveImportStatusV1ItemAuditLogErrors(items *[]catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1AuditLogErrors) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesImportMapArchiveImportStatusV1ItemAuditLogInfos(items *[]catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1AuditLogInfos) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesImportMapArchiveImportStatusV1ItemAuditLogMatchingEntitiesCount(items *[]catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1AuditLogMatchingEntitiesCount) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["key"] = item.Key
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesImportMapArchiveImportStatusV1ItemAuditLogWarnings(items *[]catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1AuditLogWarnings) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["message"] = item.Message
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenSitesImportMapArchiveImportStatusV1ItemUUID(item *catalystcentersdkgo.ResponseSitesImportMapArchiveImportStatusV1UUID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["least_significant_bits"] = item.LeastSignificantBits
	respItem["most_significant_bits"] = item.MostSignificantBits

	return []map[string]interface{}{
		respItem,
	}

}
