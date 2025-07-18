package catalystcenter

import (
	"context"
	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventSubscriptionRest() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Event Management.

- Create Rest/Webhook Subscription Endpoint for list of registered events

- Update Rest/Webhook Subscription Endpoint for list of registered events
`,

		CreateContext: resourceEventSubscriptionRestCreate,
		ReadContext:   resourceEventSubscriptionRestRead,
		UpdateContext: resourceEventSubscriptionRestUpdate,
		DeleteContext: resourceEventSubscriptionRestDelete,
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
							Type:        schema.TypeString,
							Computed:    true,
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

												"base_path": &schema.Schema{
													Description: `Base Path`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"body": &schema.Schema{
													Description: `Body`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"connect_timeout": &schema.Schema{
													Description: `Connect Timeout`,
													Type:        schema.TypeString,
													Computed:    true,
												},
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
												"headers": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"string": &schema.Schema{
																Description: `String`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"instance_id": &schema.Schema{
													Description: `Instance Id`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"method": &schema.Schema{
													Description: `Method`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"path_params": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"string": &schema.Schema{
																Description: `String`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"query_params": &schema.Schema{
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"string": &schema.Schema{
																Description: `String`,
																Type:        schema.TypeString,
																Computed:    true,
															},
														},
													},
												},
												"read_timeout": &schema.Schema{
													Description: `Read Timeout`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"resource": &schema.Schema{
													Description: `Resource`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"trust_cert": &schema.Schema{
													Description: `Trust Cert`,
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
				Description: `Array of RequestEventManagementCreateRestWebhookEventSubscription`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestEventManagementCreateRestWebhookEventSubscription`,
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
													Description: `(From 	Get Rest/Webhook Subscription Details --> pick instanceId)
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
																Description: `Connector Type (Must be REST)
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

func resourceEventSubscriptionRestCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscription(ctx, "parameters.0", d)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	vSubscriptionID := resourceItem["subscription_id"]
	vvSubscriptionID := interfaceToString(vSubscriptionID)

	queryParams1 := catalystcentersdkgo.GetRestWebhookEventSubscriptionsQueryParams{}
	item, err := searchEventManagementGetRestWebhookEventSubscriptions(m, queryParams1, vvName, vvSubscriptionID)
	if err == nil && (item != nil && len(*item) > 0) {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		resourceMap["subscription_id"] = vvSubscriptionID
		d.SetId(joinResourceID(resourceMap))
		return resourceEventSubscriptionRestRead(ctx, d, m)
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	resp1, restyResp1, err := client.EventManagement.CreateRestWebhookEventSubscription(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateRestWebhookEventSubscription", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateRestWebhookEventSubscription", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["subscription_id"] = vvSubscriptionID
	d.SetId(joinResourceID(resourceMap))
	return resourceEventSubscriptionRestRead(ctx, d, m)
}

func resourceEventSubscriptionRestRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, _ := resourceMap["name"]
	vSubscriptionID, _ := resourceMap["subscription_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetRestWebhookEventSubscriptions")
		queryParams1 := catalystcentersdkgo.GetRestWebhookEventSubscriptionsQueryParams{}
		item, err := searchEventManagementGetRestWebhookEventSubscriptions(m, queryParams1, vName, vSubscriptionID)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEventSubscriptions search response",
				err))
			return diags
		}
		if item == nil || len(*item) <= 0 {
			d.SetId("")
			return diags
		}

		if item != nil {
			log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*item))
		}

		vItem1 := flattenEventManagementGetRestWebhookEventSubscriptionsItems(item)
		if err := d.Set("item", vItem1); err != nil {

			diags = append(diags, diagError(
				"Failure when setting GetRestWebhookEventSubscriptions search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceEventSubscriptionRestUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, _ := resourceMap["name"]
	vSubscriptionID, _ := resourceMap["subscription_id"]

	queryParams1 := catalystcentersdkgo.GetRestWebhookEventSubscriptionsQueryParams{}
	item, err := searchEventManagementGetRestWebhookEventSubscriptions(m, queryParams1, vName, vSubscriptionID)
	if err != nil || item == nil || len(*item) <= 0 {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetRestWebhookEventSubscriptions", err,
			"Failure at GetRestWebhookEventSubscriptions, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	if d.HasChange("parameters") {
		request1 := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscription(ctx, "parameters.0", d)
		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}
		// Add SubscriptionID to update
		if request1 != nil && len(*request1) > 0 && item != nil && len(*item) > 0 {
			found := *item
			req := *request1
			req[0].SubscriptionID = found[0].SubscriptionID
			request1 = &req
		}
		response1, restyResp1, err := client.EventManagement.UpdateRestWebhookEventSubscription(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateRestWebhookEventSubscription", err, restyResp1.String(),
					"Failure at UpdateRestWebhookEventSubscription, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateRestWebhookEventSubscription", err,
				"Failure at UpdateRestWebhookEventSubscription, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceEventSubscriptionRestRead(ctx, d, m)
}

func resourceEventSubscriptionRestDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, _ := resourceMap["name"]
	vSubscriptionID, _ := resourceMap["subscription_id"]

	queryParams1 := catalystcentersdkgo.GetRestWebhookEventSubscriptionsQueryParams{}
	item, err := searchEventManagementGetRestWebhookEventSubscriptions(m, queryParams1, vName, vSubscriptionID)
	if err != nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetEventSubscriptions", err,
			"Failure at GetEventSubscriptions, unexpected response", ""))
		return diags
	}
	if item == nil || len(*item) == 0 {
		return diags
	}

	// REVIEW: Add getAllItems and search function to get missing params
	queryParams2 := catalystcentersdkgo.DeleteEventSubscriptionsQueryParams{}
	if len(*item) > 0 {
		itemCopy := *item
		queryParams2.Subscriptions = itemCopy[0].SubscriptionID
	}
	response1, restyResp1, err := client.EventManagement.DeleteEventSubscriptions(&queryParams2)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteEventSubscriptions", err, restyResp1.String(),
				"Failure at DeleteEventSubscriptions, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteEventSubscriptions", err,
			"Failure at DeleteEventSubscriptions, unexpected response", ""))
		return diags
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementCreateRestWebhookEventSubscription {
	request := catalystcentersdkgo.RequestEventManagementCreateRestWebhookEventSubscription{}
	if v := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscription {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscription{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscription {
	request := catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscription{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints {
	request := catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionFilter {
	request := catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains_subdomains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains_subdomains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains_subdomains")))) {
		request.DomainsSubdomains = expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemFilterDomainsSubdomainsArray(ctx, key+".domains_subdomains", d)
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

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemFilterDomainsSubdomainsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionFilterDomainsSubdomains {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionFilterDomainsSubdomains{}
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
		i := expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemFilterDomainsSubdomains(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestCreateRestWebhookEventSubscriptionItemFilterDomainsSubdomains(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionFilterDomainsSubdomains {
	request := catalystcentersdkgo.RequestItemEventManagementCreateRestWebhookEventSubscriptionFilterDomainsSubdomains{}
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

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscription(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscription {
	request := catalystcentersdkgo.RequestEventManagementUpdateRestWebhookEventSubscription{}
	if v := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscription {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscription{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscription {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscription{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionSubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemFilter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilter {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains_subdomains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains_subdomains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains_subdomains")))) {
		request.DomainsSubdomains = expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemFilterDomainsSubdomainsArray(ctx, key+".domains_subdomains", d)
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

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemFilterDomainsSubdomainsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilterDomainsSubdomains {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilterDomainsSubdomains{}
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
		i := expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemFilterDomainsSubdomains(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionRestUpdateRestWebhookEventSubscriptionItemFilterDomainsSubdomains(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilterDomainsSubdomains {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateRestWebhookEventSubscriptionFilterDomainsSubdomains{}
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

func searchEventManagementGetRestWebhookEventSubscriptions(m interface{}, queryParams catalystcentersdkgo.GetRestWebhookEventSubscriptionsQueryParams, name string, subscriptionID string) (*catalystcentersdkgo.ResponseEventManagementGetRestWebhookEventSubscriptions, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItems catalystcentersdkgo.ResponseEventManagementGetRestWebhookEventSubscriptions
	var items *catalystcentersdkgo.ResponseEventManagementGetRestWebhookEventSubscriptions
	items, _, err = client.EventManagement.GetRestWebhookEventSubscriptions(&queryParams)
	if err != nil {
		return nil, err
	}
	if items == nil {
		return nil, err
	}

	itemsCopy := *items
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.SubscriptionID == subscriptionID || item.Name == name {
			foundItems = append(foundItems, item)
			break
		}
	}
	return &foundItems, err
}
