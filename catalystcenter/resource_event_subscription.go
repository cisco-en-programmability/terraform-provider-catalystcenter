package catalystcenter

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEventSubscription() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Event Management.

- Delete EventSubscriptions

- Update SubscriptionEndpoint to list of registered events(Deprecated)

- Subscribe SubscriptionEndpoint to list of registered events (Deprecated)
`,

		CreateContext: resourceEventSubscriptionCreate,
		ReadContext:   resourceEventSubscriptionRead,
		UpdateContext: resourceEventSubscriptionUpdate,
		DeleteContext: resourceEventSubscriptionDelete,
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
													Type:        schema.TypeInt,
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
													Type:        schema.TypeInt,
													Computed:    true,
												},
												"resource": &schema.Schema{
													Description: `Resource`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"trust_cert": &schema.Schema{
													Description: `Trust Cert`,
													// Type:        schema.TypeBool,
													Type:     schema.TypeString,
													Computed: true,
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
				Description: `Array of RequestEventManagementCreateEventSubscriptionsV1`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestEventManagementCreateEventSubscriptions`,
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

func resourceEventSubscriptionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	type Error struct {
		APIStatus string `json:"apiStatus,omitempty"` // Error
	}

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestEventSubscriptionCreateEventSubscriptionsV1(ctx, "parameters.0", d)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	vSubscriptionID := resourceItem["subscription_id"]
	vvSubscriptionID := interfaceToString(vSubscriptionID)

	queryParams1 := catalystcentersdkgo.GetEventSubscriptionsV1QueryParams{}
	item, err := searchEventManagementGetEventSubscriptions(m, queryParams1, vvName, vvSubscriptionID)
	if err == nil && (item != nil && len(*item) > 0) {
		resourceMap := make(map[string]string)
		item2 := *item
		resourceMap["name"] = item2[0].Name
		resourceMap["subscription_id"] = item2[0].SubscriptionID
		d.SetId(joinResourceID(resourceMap))
		return resourceEventSubscriptionRead(ctx, d, m)
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	resp1, restyResp1, err := client.EventManagement.CreateEventSubscriptions(request1)

	if err != nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEventSubscriptions", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEventSubscriptions", err))
		return diags
	}

	restyResp3, err := client.CustomCall.GetCustomCall(resp1.StatusURI, nil)
	var errorValid Error
	err = json.Unmarshal(restyResp3.Body(), &errorValid)
	if err != nil {
		if restyResp3 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEventSubscriptions", err, restyResp3.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEventSubscriptions", err))
		return diags
	}
	if err != nil || strings.ToUpper(errorValid.APIStatus) == "FAILURE" {
		if restyResp3 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateEventSubscriptions", err, restyResp3.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateEventSubscriptions", err))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["subscription_id"] = vvSubscriptionID
	d.SetId(joinResourceID(resourceMap))
	return resourceEventSubscriptionRead(ctx, d, m)
}

func resourceEventSubscriptionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, _ := resourceMap["name"]
	vSubscriptionID, _ := resourceMap["subscription_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetEventSubscriptions")
		queryParams1 := catalystcentersdkgo.GetEventSubscriptionsV1QueryParams{}
		item, err := searchEventManagementGetEventSubscriptions(m, queryParams1, vName, vSubscriptionID)
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

		vItem1 := flattenEventManagementGetEventSubscriptionsV1Items(item)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetEventSubscriptions search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceEventSubscriptionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, _ := resourceMap["name"]
	vSubscriptionID, _ := resourceMap["subscription_id"]

	queryParams1 := catalystcentersdkgo.GetEventSubscriptionsV1QueryParams{}
	item, err := searchEventManagementGetEventSubscriptions(m, queryParams1, vName, vSubscriptionID)
	if err != nil || item == nil || len(*item) <= 0 {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetEventSubscriptions", err,
			"Failure at GetEventSubscriptions, unexpected response", ""))
		return diags
	}

	// NOTE: Consider adding getAllItems and search function to get missing params
	if d.HasChange("parameters") {
		request1 := expandRequestEventSubscriptionUpdateEventSubscriptionsV1(ctx, "parameters.0", d)
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
		response1, restyResp1, err := client.EventManagement.UpdateEventSubscriptions(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateEventSubscriptions", err, restyResp1.String(),
					"Failure at UpdateEventSubscriptions, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateEventSubscriptions", err,
				"Failure at UpdateEventSubscriptions, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceEventSubscriptionRead(ctx, d, m)
}

func resourceEventSubscriptionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vName, _ := resourceMap["name"]
	vSubscriptionID, _ := resourceMap["subscription_id"]

	queryParams1 := catalystcentersdkgo.GetEventSubscriptionsV1QueryParams{}
	item, err := searchEventManagementGetEventSubscriptions(m, queryParams1, vName, vSubscriptionID)
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
	queryParams2 := catalystcentersdkgo.DeleteEventSubscriptionsV1QueryParams{}
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

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestEventSubscriptionCreateEventSubscriptionsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementCreateEventSubscriptionsV1 {
	request := catalystcentersdkgo.RequestEventManagementCreateEventSubscriptionsV1{}
	if v := expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1 {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1 {
	request := catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpoints {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpoints {
	request := catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails {
	request := catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemFilter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1Filter {
	request := catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1Filter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains_subdomains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains_subdomains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains_subdomains")))) {
		request.DomainsSubdomains = expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemFilterDomainsSubdomainsArray(ctx, key+".domains_subdomains", d)
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

func expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemFilterDomainsSubdomainsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1FilterDomainsSubdomains {
	request := []catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1FilterDomainsSubdomains{}
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
		i := expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemFilterDomainsSubdomains(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionCreateEventSubscriptionsV1ItemFilterDomainsSubdomains(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1FilterDomainsSubdomains {
	request := catalystcentersdkgo.RequestItemEventManagementCreateEventSubscriptionsV1FilterDomainsSubdomains{}
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

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestEventManagementUpdateEventSubscriptionsV1 {
	request := catalystcentersdkgo.RequestEventManagementUpdateEventSubscriptionsV1{}
	if v := expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1 {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1 {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1{}
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
		request.SubscriptionEndpoints = expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemSubscriptionEndpointsArray(ctx, key+".subscription_endpoints", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".filter")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".filter")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".filter")))) {
		request.Filter = expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemFilter(ctx, key+".filter.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemSubscriptionEndpointsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpoints {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpoints{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemSubscriptionEndpoints(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemSubscriptionEndpoints(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpoints {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpoints{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subscription_details")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subscription_details")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subscription_details")))) {
		request.SubscriptionDetails = expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemSubscriptionEndpointsSubscriptionDetails(ctx, key+".subscription_details.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemSubscriptionEndpointsSubscriptionDetails(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1SubscriptionEndpointsSubscriptionDetails{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connector_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connector_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connector_type")))) {
		request.ConnectorType = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemFilter(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1Filter {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1Filter{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".event_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".event_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".event_ids")))) {
		request.EventIDs = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".domains_subdomains")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".domains_subdomains")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".domains_subdomains")))) {
		request.DomainsSubdomains = expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemFilterDomainsSubdomainsArray(ctx, key+".domains_subdomains", d)
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

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemFilterDomainsSubdomainsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1FilterDomainsSubdomains {
	request := []catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1FilterDomainsSubdomains{}
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
		i := expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemFilterDomainsSubdomains(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestEventSubscriptionUpdateEventSubscriptionsV1ItemFilterDomainsSubdomains(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1FilterDomainsSubdomains {
	request := catalystcentersdkgo.RequestItemEventManagementUpdateEventSubscriptionsV1FilterDomainsSubdomains{}
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

func searchEventManagementGetEventSubscriptions(m interface{}, queryParams catalystcentersdkgo.GetEventSubscriptionsV1QueryParams, name string, subscriptionID string) (*catalystcentersdkgo.ResponseEventManagementGetEventSubscriptionsV1, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItems catalystcentersdkgo.ResponseEventManagementGetEventSubscriptionsV1 = []catalystcentersdkgo.ResponseItemEventManagementGetEventSubscriptionsV1{}
	var items *catalystcentersdkgo.ResponseEventManagementGetEventSubscriptionsV1
	items, _, err = client.EventManagement.GetEventSubscriptions(&queryParams)
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
