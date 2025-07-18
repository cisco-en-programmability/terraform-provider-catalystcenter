package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceCustomIssueDefinitions() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Issues.

- Retrieve the existing syslog-based custom issue definitions. The supported filters are id, name, profileId,
definition enable status, priority, severity, facility and mnemonic. The issue definition configurations may vary across
profiles, hence specifying the profile Id in the query parameter is important and the default profile is global.


  The assurance profile definitions can be obtain via the API endpoint: /api/v1/siteprofile?namespace=assurance. For
detailed information about the usage of the API, please refer to the Open API specification document
https://github.com/cisco-en-programmability/catalyst-center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-
AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml

- Get the custom issue definition for the given custom issue definition Id. For detailed information about the usage of
the API, please refer to the Open API specification document https://github.com/cisco-en-programmability/catalyst-
center-api-specs/blob/main/Assurance/CE_Cat_Center_Org-AssuranceUserDefinedIssueAPIs-1.0.0-resolved.yaml
`,

		ReadContext: dataSourceCustomIssueDefinitionsRead,
		Schema: map[string]*schema.Schema{
			"facility": &schema.Schema{
				Description: `facility query parameter. The syslog facility name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": &schema.Schema{
				Description: `id path parameter. Get the custom issue definition for the given custom issue definition Id.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"is_enabled": &schema.Schema{
				Description: `isEnabled query parameter. The enable status of the custom issue definition, either true or false.
`,
				Type:     schema.TypeBool,
				Optional: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The maximum number of records to return
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"mnemonic": &schema.Schema{
				Description: `mnemonic query parameter. The syslog mnemonic name
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Description: `name query parameter. The list of UDI issue names
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. Specifies the starting point within all records returned by the API. It's one based offset. The starting value is 1.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"order": &schema.Schema{
				Description: `order query parameter. The sort order of the field ascending or descending.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": &schema.Schema{
				Description: `priority query parameter. The Issue priority value, possible values are P1, P2, P3, P4. P1: A critical issue that needs immediate attention and can have a wide impact on network operations. P2: A major issue that can potentially impact multiple devices or clients. P3: A minor issue that has a localized or minimal impact. P4: A warning issue that may not be an immediate problem but addressing it can optimize the network performance
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"profile_id": &schema.Schema{
				Description: `profileId query parameter. The profile identifier to fetch the profile associated custom issue definitions. The default is global. For the custom profile, it is profile UUID. Example : 3fa85f64-5717-4562-b3fc-2c963f66afa6
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": &schema.Schema{
				Description: `severity query parameter. The syslog severity level. 0: Emergency 1: Alert, 2: Critical. 3: Error, 4: Warning, 5: Notice, 6: Info. Examples:severity=1&severity=2 (multi value support with & separator)
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"sort_by": &schema.Schema{
				Description: `sortBy query parameter. A field within the response to sort by.
`,
				Type:     schema.TypeString,
				Optional: true,
			},
			"xca_lle_rid": &schema.Schema{
				Description: `X-CALLER-ID header parameter. Caller ID can be used to trace the caller for queries executed on database. The caller id is like a optional attribute which can be added to API invocation like ui, python, postman, test-automation etc
`,
				Type:     schema.TypeString,
				Optional: true,
			},

			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"created_time": &schema.Schema{
							Description: `Created Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_deletable": &schema.Schema{
							Description: `Is Deletable`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_enabled": &schema.Schema{
							Description: `Is Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_notification_enabled": &schema.Schema{
							Description: `Is Notification Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_time": &schema.Schema{
							Description: `Last Updated Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"profile_id": &schema.Schema{
							Description: `Profile Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"duration_in_minutes": &schema.Schema{
										Description: `Duration In Minutes`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"facility": &schema.Schema{
										Description: `Facility`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"mnemonic": &schema.Schema{
										Description: `Mnemonic`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"occurrences": &schema.Schema{
										Description: `Occurrences`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"pattern": &schema.Schema{
										Description: `Pattern`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"severity": &schema.Schema{
										Description: `Severity`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"trigger_id": &schema.Schema{
							Description: `Trigger Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"created_time": &schema.Schema{
							Description: `Created Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"is_deletable": &schema.Schema{
							Description: `Is Deletable`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_enabled": &schema.Schema{
							Description: `Is Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"is_notification_enabled": &schema.Schema{
							Description: `Is Notification Enabled`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"last_updated_time": &schema.Schema{
							Description: `Last Updated Time`,
							Type:        schema.TypeInt,
							Computed:    true,
						},

						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"priority": &schema.Schema{
							Description: `Priority`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"profile_id": &schema.Schema{
							Description: `Profile Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},

						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"duration_in_minutes": &schema.Schema{
										Description: `Duration In Minutes`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"facility": &schema.Schema{
										Description: `Facility`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"mnemonic": &schema.Schema{
										Description: `Mnemonic`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"occurrences": &schema.Schema{
										Description: `Occurrences`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"pattern": &schema.Schema{
										Description: `Pattern`,
										Type:        schema.TypeString,
										Computed:    true,
									},

									"severity": &schema.Schema{
										Description: `Severity`,
										Type:        schema.TypeInt,
										Computed:    true,
									},

									"type": &schema.Schema{
										Description: `Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},

						"trigger_id": &schema.Schema{
							Description: `Trigger Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCustomIssueDefinitionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vID, okID := d.GetOk("id")
	vProfileID, okProfileID := d.GetOk("profile_id")
	vName, okName := d.GetOk("name")
	vPriority, okPriority := d.GetOk("priority")
	vIsEnabled, okIsEnabled := d.GetOk("is_enabled")
	vSeverity, okSeverity := d.GetOk("severity")
	vFacility, okFacility := d.GetOk("facility")
	vMnemonic, okMnemonic := d.GetOk("mnemonic")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vSortBy, okSortBy := d.GetOk("sort_by")
	vOrder, okOrder := d.GetOk("order")
	vXCaLLERID, okXCaLLERID := d.GetOk("xca_lle_rid")

	method1 := []bool{okID, okProfileID, okName, okPriority, okIsEnabled, okSeverity, okFacility, okMnemonic, okLimit, okOffset, okSortBy, okOrder}
	log.Printf("[DEBUG] Selecting method. Method 1 %v", method1)
	method2 := []bool{okID, okXCaLLERID}
	log.Printf("[DEBUG] Selecting method. Method 2 %v", method2)

	selectedMethod := pickMethod([][]bool{method1, method2})
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters")
		queryParams1 := catalystcentersdkgo.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersQueryParams{}

		if okID {
			queryParams1.ID = vID.(string)
		}
		if okProfileID {
			queryParams1.ProfileID = vProfileID.(string)
		}
		if okName {
			queryParams1.Name = vName.(string)
		}
		if okPriority {
			queryParams1.Priority = vPriority.(string)
		}
		if okIsEnabled {
			queryParams1.IsEnabled = vIsEnabled.(bool)
		}
		if okSeverity {
			queryParams1.Severity = vSeverity.(float64)
		}
		if okFacility {
			queryParams1.Facility = vFacility.(string)
		}
		if okMnemonic {
			queryParams1.Mnemonic = vMnemonic.(string)
		}
		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okSortBy {
			queryParams1.SortBy = vSortBy.(string)
		}
		if okOrder {
			queryParams1.Order = vOrder.(string)
		}

		// has_unknown_response: None

		response1, restyResp1, err := client.Issues.GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters", err,
				"Failure at GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters", err,
				"Failure at GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItems(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetAllTheCustomIssueDefinitionsBasedOnTheGivenFilters response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	if selectedMethod == 2 {
		log.Printf("[DEBUG] Selected method: GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID")
		vvID := vID.(string)

		headerParams2 := catalystcentersdkgo.GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDHeaderParams{}

		if okXCaLLERID {
			headerParams2.XCaLLERID = vXCaLLERID.(string)
		}

		// has_unknown_response: None

		response2, restyResp2, err := client.Issues.GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID(vvID, &headerParams2)

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID", err,
				"Failure at GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID", err,
				"Failure at GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response2))

		vItem2 := flattenIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDItem(response2.Response)
		if err := d.Set("item", vItem2); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionID response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItems(items *[]catalystcentersdkgo.ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponse) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["id"] = item.ID
		respItem["name"] = item.Name
		respItem["description"] = item.Description
		respItem["profile_id"] = item.ProfileID
		respItem["trigger_id"] = item.TriggerID
		respItem["rules"] = flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItemsRules(item.Rules)
		respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
		respItem["priority"] = item.Priority
		respItem["is_deletable"] = boolPtrToString(item.IsDeletable)
		respItem["is_notification_enabled"] = boolPtrToString(item.IsNotificationEnabled)
		respItem["created_time"] = item.CreatedTime
		respItem["last_updated_time"] = item.LastUpdatedTime
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersItemsRules(items *[]catalystcentersdkgo.ResponseIssuesGetAllTheCustomIssueDefinitionsBasedOnTheGivenFiltersResponseRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["severity"] = item.Severity
		respItem["facility"] = item.Facility
		respItem["mnemonic"] = item.Mnemonic
		respItem["pattern"] = item.Pattern
		respItem["occurrences"] = item.Occurrences
		respItem["duration_in_minutes"] = item.DurationInMinutes
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDItem(item *catalystcentersdkgo.ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDResponse) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["name"] = item.Name
	respItem["description"] = item.Description
	respItem["profile_id"] = item.ProfileID
	respItem["trigger_id"] = item.TriggerID
	respItem["rules"] = flattenIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDItemRules(item.Rules)
	respItem["is_enabled"] = boolPtrToString(item.IsEnabled)
	respItem["priority"] = item.Priority
	respItem["is_deletable"] = boolPtrToString(item.IsDeletable)
	respItem["is_notification_enabled"] = boolPtrToString(item.IsNotificationEnabled)
	respItem["created_time"] = item.CreatedTime
	respItem["last_updated_time"] = item.LastUpdatedTime
	return []map[string]interface{}{
		respItem,
	}
}

func flattenIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDItemRules(items *[]catalystcentersdkgo.ResponseIssuesGetTheCustomIssueDefinitionForTheGivenCustomIssueDefinitionIDResponseRules) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["type"] = item.Type
		respItem["severity"] = item.Severity
		respItem["facility"] = item.Facility
		respItem["mnemonic"] = item.Mnemonic
		respItem["pattern"] = item.Pattern
		respItem["occurrences"] = item.Occurrences
		respItem["duration_in_minutes"] = item.DurationInMinutes
		respItems = append(respItems, respItem)
	}
	return respItems
}
