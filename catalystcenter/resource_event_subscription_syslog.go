package catalystcenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventSubscriptionSyslog() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Event Management.

- Update Syslog Subscription Endpoint for list of registered events

- Create Syslog Subscription Endpoint for list of registered events
`,

		CreateContext: resourceEventSubscriptionSyslogCreate,
		ReadContext:   resourceEventSubscriptionSyslogRead,
		UpdateContext: resourceEventSubscriptionSyslogUpdate,
		DeleteContext: resourceEventSubscriptionSyslogDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"item": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"filter": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"categories": &schema.Schema{
										Description: `Categories`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"domains_subdomains": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"domain": &schema.Schema{
													Description: `Domain`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"sub_domains": &schema.Schema{
													Description: `Sub Domains`,
													Type:        schema.TypeList,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"event_ids": &schema.Schema{
										Description: `Event Ids`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"others": &schema.Schema{
										Description: `Others`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"severities": &schema.Schema{
										Description: `Severities`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"site_ids": &schema.Schema{
										Description: `Site Ids`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"sources": &schema.Schema{
										Description: `Sources`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"types": &schema.Schema{
										Description: `Types`,
										Type:        schema.TypeList,
										Computed:    true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"is_private": &schema.Schema{
							Description: `Is Private`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"subscription_endpoints": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connector_type": &schema.Schema{
										Description: `Connector Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"instance_id": &schema.Schema{
										Description: `Instance Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"subscription_details": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"connector_type": &schema.Schema{
													Description: `Connector Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"description": &schema.Schema{
													Description: `Description`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"instance_id": &schema.Schema{
													Description: `Instance Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"syslog_config": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"config_id": &schema.Schema{
																Description: `Config Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"description": &schema.Schema{
																Description: `Description`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"host": &schema.Schema{
																Description: `Host`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"name": &schema.Schema{
																Description: `Name`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"port": &schema.Schema{
																Description: `Port`,
																Type:        schema.TypeInt,
																Computed:    true,
															},
															"tenant_id": &schema.Schema{
																Description: `Tenant Id`,
																Type:        schema.TypeString,
																Computed:    true,
															},
															"version": &schema.Schema{
																Description: `Version`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"subscription_id": &schema.Schema{
							Description: `Subscription Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"version": &schema.Schema{
							Description: `Version`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestEventManagementCreateSyslogEventSubscriptionV1`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestEventManagementCreateSyslogEventSubscription`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"description": &schema.Schema{
										Description: `Description
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"filter": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"categories": &schema.Schema{
													Description: `Categories`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"domains_subdomains": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"domain": &schema.Schema{
																Description: `Domain`,
																Type:        schema.TypeString,
																Optional:    true,
																Computed:    true,
															},
															"sub_domains": &schema.Schema{
																Description: `Sub Domains`,
																Type:        schema.TypeList,
																Optional:    true,
																Computed:    true,
																Elem: &schema.Schema{
																	Type: schema.TypeString,
																},
															},
														},
													},
												},
												"event_ids": &schema.Schema{
													Description: `Event Ids (Comma separated event ids)
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"severities": &schema.Schema{
													Description: `Severities`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"site_ids": &schema.Schema{
													Description: `Site Ids`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"sources": &schema.Schema{
													Description: `Sources`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"types": &schema.Schema{
													Description: `Types`,
													Type:        schema.TypeList,
													Optional:    true,
													Computed:    true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `Name
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"subscription_endpoints": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"instance_id": &schema.Schema{
													Description: `(From Get Syslog Subscription Details --> pick instanceId)
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"subscription_details": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"connector_type": &schema.Schema{
																Description: `Connector Type (Must be SYSLOG)
`,
																Type:     schema.TypeString,
																Optional: true,
																Computed: true,
															},
														},
													},
												},
											},
										},
									},
									"subscription_id": &schema.Schema{
										Description: `Subscription Id (Unique UUID)
`,
										Type:     schema.TypeString,
										Optional: true,
										Default:  "",
									},
									"version": &schema.Schema{
										Description: `Version
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								}}}},
				},
			},
		},
	}
}

func resourceEventSubscriptionSyslogCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	vSubscriptionID := resourceItem["subscription_id"]
	vvSubscriptionID := interfaceToString(vSubscriptionID)
	queryParamImport := catalystcentersdkgo.GetSyslogEventSubscriptionsV1QueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchEventManagementGetSyslogEventSubscriptions(m, queryParamImport, vvSubscriptionID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = item2.Name
		resourceMap["id"] = item2.SubscriptionID
		d.SetId(joinResourceID(resourceMap))
		return resourceEventSubscriptionSyslogRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.EventManagement.CreateSyslogEventSubscription(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateSyslogEventSubscription", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateSyslogEventSubscription", err))
		return diags
	}
	// TODO REVIEW
	queryParamValidate := catalystcentersdkgo.GetSyslogEventSubscriptionsV1QueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchEventManagementGetSyslogEventSubscriptions(m, queryParamValidate, "")
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateSyslogEventSubscription", err,
			"Failure at CreateSyslogEventSubscription, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["id"] = item3.SubscriptionID
	d.SetId(joinResourceID(resourceMap))
	return resourceEventSubscriptionSyslogRead(ctx, d, m)
}

func resourceEventSubscriptionSyslogRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vName := resourceMap["name"]
	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetSyslogEventSubscriptions")
		queryParams1 := catalystcentersdkgo.GetSyslogEventSubscriptionsV1QueryParams{}
		queryParams1.Name = vName
		item1, err := searchEventManagementGetSyslogEventSubscriptions(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		items := catalystcentersdkgo.ResponseEventManagementGetSyslogEventSubscriptionsV1{
			*item1,
		}
		// Review flatten function used
		vItem1 := flattenEventManagementGetSyslogEventSubscriptionsV1Items(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSyslogEventSubscriptions search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceEventSubscriptionSyslogUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.EventManagement.UpdateSyslogEventSubscription(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateSyslogEventSubscription", err, restyResp1.String(),
					"Failure at UpdateSyslogEventSubscription, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateSyslogEventSubscription", err,
				"Failure at UpdateSyslogEventSubscription, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceEventSubscriptionSyslogRead(ctx, d, m)
}

func resourceEventSubscriptionSyslogDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing EventSubscriptionSyslogDelete", err, "Delete method is not supported",
		"Failure at EventSubscriptionSyslogDelete, unexpected response", ""))

	return diags
}
func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementCreateSyslogEventSubscriptionV1 {
	request := catalystcentersdkgo.RequestEventManagementCreateSyslogEventSubscriptionV1{}
	if v := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1 {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1 {
	request := catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpoints {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpoints{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpoints {
	request := catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails {
	request := catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemFilter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1Filter {
	request := catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1Filter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains_subdomains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains_subdomains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains_subdomains")))) {
		request.DomainsSubdomains = expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemFilterDomainsSubdomainsArray(ctx, key+".domains_subdomains", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".types")))) {
		request.Types = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".categories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".categories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".categories")))) {
		request.Categories = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".severities")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".severities")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".severities")))) {
		request.Severities = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sources")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sources")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sources")))) {
		request.Sources = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemFilterDomainsSubdomainsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1FilterDomainsSubdomains {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1FilterDomainsSubdomains{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemFilterDomainsSubdomains(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogCreateSyslogEventSubscriptionV1ItemFilterDomainsSubdomains(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1FilterDomainsSubdomains {
	request := catalystcentersdkgo.RequestItemEventManagementCreateSyslogEventSubscriptionV1FilterDomainsSubdomains{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain")))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sub_domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sub_domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sub_domains")))) {
		request.SubDomains = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementUpdateSyslogEventSubscriptionV1 {
	request := catalystcentersdkgo.RequestEventManagementUpdateSyslogEventSubscriptionV1{}
	if v := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1 {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1 {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_id")))) {
		request.SubscriptionID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".version")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".version")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".version")))) {
		request.Version = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_endpoints")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_endpoints")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_endpoints")))) {
		request.SubscriptionEndpoints = expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpoints {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpoints{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpoints {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1SubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemFilter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1Filter {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1Filter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains_subdomains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains_subdomains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains_subdomains")))) {
		request.DomainsSubdomains = expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemFilterDomainsSubdomainsArray(ctx, key+".domains_subdomains", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".types")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".types")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".types")))) {
		request.Types = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".categories")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".categories")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".categories")))) {
		request.Categories = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".severities")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".severities")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".severities")))) {
		request.Severities = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sources")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sources")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sources")))) {
		request.Sources = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_ids")))) {
		request.SiteIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemFilterDomainsSubdomainsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1FilterDomainsSubdomains {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1FilterDomainsSubdomains{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no := range objs {
		i := expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemFilterDomainsSubdomains(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionSyslogUpdateSyslogEventSubscriptionV1ItemFilterDomainsSubdomains(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1FilterDomainsSubdomains {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateSyslogEventSubscriptionV1FilterDomainsSubdomains{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domain")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domain")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domain")))) {
		request.Domain = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sub_domains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sub_domains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sub_domains")))) {
		request.SubDomains = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchEventManagementGetSyslogEventSubscriptions(m interface{}, queryParams catalystcentersdkgo.GetSyslogEventSubscriptionsV1QueryParams, vID string) (*catalystcentersdkgo.ResponseItemEventManagementGetSyslogEventSubscriptionsV1, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseItemEventManagementGetSyslogEventSubscriptionsV1
	var ite *catalystcentersdkgo.ResponseEventManagementGetSyslogEventSubscriptionsV1
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.EventManagement.GetSyslogEventSubscriptions(nil)
		maxPageSize := len(*nResponse)
		for len(*nResponse) > 0 {
			for _, item := range *nResponse {
				if vID == item.SubscriptionID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.EventManagement.GetSyslogEventSubscriptions(&queryParams)
			if nResponse == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.Name != "" {
		ite, _, err = client.EventManagement.GetSyslogEventSubscriptions(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.Name == queryParams.Name {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
