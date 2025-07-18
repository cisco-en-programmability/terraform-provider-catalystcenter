package catalystcenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplications() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Create new Custom application

- Edit the attributes of an existing application

- Delete existing application by its id
`,

		CreateContext: resourceApplicationsCreate,
		ReadContext:   resourceApplicationsRead,
		UpdateContext: resourceApplicationsUpdate,
		DeleteContext: resourceApplicationsDelete,
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

						"application_set": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id_ref": &schema.Schema{
										Description: `Id Ref`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"indicative_network_identity": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `displayName`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"lower_port": &schema.Schema{
										Description: `lowerPort`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
									"ports": &schema.Schema{
										Description: `ports`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"protocol": &schema.Schema{
										Description: `protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"upper_port": &schema.Schema{
										Description: `upperPort`,
										Type:        schema.TypeInt,
										Computed:    true,
									},
								},
							},
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"network_applications": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"app_protocol": &schema.Schema{
										Description: `App Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"application_sub_type": &schema.Schema{
										Description: `Application Sub Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"application_type": &schema.Schema{
										Description: `Application Type`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"category_id": &schema.Schema{
										Description: `Category Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"dscp": &schema.Schema{
										Description: `Dscp`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"engine_id": &schema.Schema{
										Description: `Engine Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"help_string": &schema.Schema{
										Description: `Help String`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ignore_conflict": &schema.Schema{
										Description: `Ignore Conflict`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"long_description": &schema.Schema{
										Description: `Long Description`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"popularity": &schema.Schema{
										Description: `Popularity`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"rank": &schema.Schema{
										Description: `Rank`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"server_name": &schema.Schema{
										Description: `Server Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"traffic_class": &schema.Schema{
										Description: `Traffic Class`,
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
						"network_identity": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display Name`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"lower_port": &schema.Schema{
										Description: `Lower Port`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"ports": &schema.Schema{
										Description: `Ports`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol`,
										Type:        schema.TypeString,
										Computed:    true,
									},
									"upper_port": &schema.Schema{
										Description: `Upper Port`,
										Type:        schema.TypeString,
										Computed:    true,
									},
								},
							},
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplication`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplication`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"application_set": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id_ref": &schema.Schema{
													Description: `Id Ref`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "",
									},
									"indicative_network_identity": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `displayName`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"id": &schema.Schema{
													Description: `id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"lower_port": &schema.Schema{
													Description: `lowerPort`,
													Type:        schema.TypeInt,
													Optional:    true,
												},
												"ports": &schema.Schema{
													Description: `ports`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"protocol": &schema.Schema{
													Description: `protocol`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"upper_port": &schema.Schema{
													Description: `upperPort`,
													Type:        schema.TypeInt,
													Optional:    true,
												},
											},
										},
									},
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "",
									},
									"network_applications": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"app_protocol": &schema.Schema{
													Description: `App Protocol`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"application_sub_type": &schema.Schema{
													Description: `Application Sub Type`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"application_type": &schema.Schema{
													Description: `Application Type`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"category_id": &schema.Schema{
													Description: `Category Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"display_name": &schema.Schema{
													Description: `Display Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"dscp": &schema.Schema{
													Description: `Dscp`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"engine_id": &schema.Schema{
													Description: `Engine Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"help_string": &schema.Schema{
													Description: `Help String`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"ignore_conflict": &schema.Schema{
													Description: `Ignore Conflict`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"long_description": &schema.Schema{
													Description: `Long Description`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"name": &schema.Schema{
													Description: `Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"popularity": &schema.Schema{
													Description: `Popularity`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"rank": &schema.Schema{
													Description: `Rank`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"server_name": &schema.Schema{
													Description: `Server Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"traffic_class": &schema.Schema{
													Description: `Traffic Class`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"url": &schema.Schema{
													Description: `Url`,
													Type:        schema.TypeString,
													Optional:    true,
												},
											},
										},
									},
									"network_identity": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"display_name": &schema.Schema{
													Description: `Display Name`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"id": &schema.Schema{
													Description: `Id`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"lower_port": &schema.Schema{
													Description: `Lower Port`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"ports": &schema.Schema{
													Description: `Ports`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"protocol": &schema.Schema{
													Description: `Protocol`,
													Type:        schema.TypeString,
													Optional:    true,
												},
												"upper_port": &schema.Schema{
													Description: `Upper Port`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
								}}},
					},
				},
			},
		},
	}
}

func resourceApplicationsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestApplicationsCreateApplication(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vName := resourceItem["name"]
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vvName := interfaceToString(vName)
	queryParamImport := catalystcentersdkgo.GetApplicationsQueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchApplicationPolicyGetApplications(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceApplicationsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplication(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplication", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplication", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApplication", err))
		return diags
	}
	taskId := resp1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateApplication", err1))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetApplicationsQueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchApplicationPolicyGetApplications(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateApplication", err,
			"Failure at CreateApplication, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceApplicationsRead(ctx, d, m)
}

func resourceApplicationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvName := resourceMap["name"]
	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplications")
		queryParams1 := catalystcentersdkgo.GetApplicationsQueryParams{}
		queryParams1.Name = vvName
		item1, err := searchApplicationPolicyGetApplications(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []catalystcentersdkgo.ResponseItemApplicationPolicyGetApplications{
			*item1,
		}
		var items2 catalystcentersdkgo.ResponseApplicationPolicyGetApplications
		items2 = items
		vItem1 := flattenApplicationPolicyGetApplicationsItems(&items2)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplications search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceApplicationsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestApplicationsEditApplication(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.ApplicationPolicy.EditApplication(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing EditApplication", err, restyResp1.String(),
					"Failure at EditApplication, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing EditApplication", err,
				"Failure at EditApplication, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing EditApplication", err))
			return diags
		}
		taskId := response1.Response.TaskID
		log.Printf("[DEBUG] TASKID => %s", taskId)
		if taskId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetTaskByID(taskId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetTaskByID", err,
					"Failure at GetTaskByID, unexpected response", ""))
				return diags
			}
			if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
				log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
				errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing EditApplication", err1))
				return diags
			}
		}

	}

	return resourceApplicationsRead(ctx, d, m)
}

func resourceApplicationsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplication(&catalystcentersdkgo.DeleteApplicationQueryParams{
		ID: vvID,
	})
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteApplication", err, restyResp1.String(),
				"Failure at DeleteApplication, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteApplication2", err,
			"Failure at DeleteApplication2, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteApplication", err))
		return diags
	}
	taskId := response1.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response2.Response != nil && response2.Response.IsError != nil && *response2.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "Failure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing DeleteApplication", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestApplicationsCreateApplication(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyCreateApplication {
	request := catalystcentersdkgo.RequestApplicationPolicyCreateApplication{}
	if v := expandRequestApplicationsCreateApplicationItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplication {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplication{}
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
		i := expandRequestApplicationsCreateApplicationItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplication {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplication{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_applications")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_applications")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_applications")))) {
		request.NetworkApplications = expandRequestApplicationsCreateApplicationItemNetworkApplicationsArray(ctx, key+".network_applications", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_identity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_identity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_identity")))) {
		request.NetworkIDentity = expandRequestApplicationsCreateApplicationItemNetworkIDentityArray(ctx, key+".network_identity", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_set")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_set")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_set")))) {
		request.ApplicationSet = expandRequestApplicationsCreateApplicationItemApplicationSet(ctx, key+".application_set.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".indicative_network_identity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".indicative_network_identity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".indicative_network_identity")))) {
		request.IndicativeNetworkIDentity = expandRequestApplicationsCreateApplicationItemIndicativeNetworkIDentityArray(ctx, key+".indicative_network_identity", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItemNetworkApplicationsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationNetworkApplications {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationNetworkApplications{}
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
		i := expandRequestApplicationsCreateApplicationItemNetworkApplications(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItemNetworkApplications(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationNetworkApplications {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationNetworkApplications{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".app_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".app_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".app_protocol")))) {
		request.AppProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_sub_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_sub_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_sub_type")))) {
		request.ApplicationSubType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_type")))) {
		request.ApplicationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".category_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".category_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".category_id")))) {
		request.CategoryID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".engine_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".engine_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".engine_id")))) {
		request.EngineID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".help_string")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".help_string")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".help_string")))) {
		request.HelpString = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".long_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".long_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".long_description")))) {
		request.LongDescription = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".popularity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".popularity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".popularity")))) {
		request.Popularity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rank")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rank")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rank")))) {
		request.Rank = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_name")))) {
		request.ServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ignore_conflict")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ignore_conflict")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ignore_conflict")))) {
		request.IgnoreConflict = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItemNetworkIDentityArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationNetworkIDentity {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationNetworkIDentity{}
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
		i := expandRequestApplicationsCreateApplicationItemNetworkIDentity(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItemNetworkIDentity(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationNetworkIDentity {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationNetworkIDentity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lower_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lower_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lower_port")))) {
		request.LowerPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upper_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upper_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upper_port")))) {
		request.UpperPort = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItemApplicationSet(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationApplicationSet {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationApplicationSet{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItemIndicativeNetworkIDentityArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationIndicativeNetworkIDentity {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationIndicativeNetworkIDentity{}
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
		i := expandRequestApplicationsCreateApplicationItemIndicativeNetworkIDentity(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsCreateApplicationItemIndicativeNetworkIDentity(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationIndicativeNetworkIDentity {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationIndicativeNetworkIDentity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lower_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lower_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lower_port")))) {
		request.LowerPort = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upper_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upper_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upper_port")))) {
		request.UpperPort = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsEditApplication(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyEditApplication {
	request := catalystcentersdkgo.RequestApplicationPolicyEditApplication{}
	if v := expandRequestApplicationsEditApplicationItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsEditApplicationItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyEditApplication {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyEditApplication{}
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
		i := expandRequestApplicationsEditApplicationItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsEditApplicationItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyEditApplication {
	request := catalystcentersdkgo.RequestItemApplicationPolicyEditApplication{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_applications")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_applications")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_applications")))) {
		request.NetworkApplications = expandRequestApplicationsEditApplicationItemNetworkApplicationsArray(ctx, key+".network_applications", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_identity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_identity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_identity")))) {
		request.NetworkIDentity = expandRequestApplicationsEditApplicationItemNetworkIDentityArray(ctx, key+".network_identity", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_set")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_set")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_set")))) {
		request.ApplicationSet = expandRequestApplicationsEditApplicationItemApplicationSet(ctx, key+".application_set.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsEditApplicationItemNetworkApplicationsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationNetworkApplications {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationNetworkApplications{}
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
		i := expandRequestApplicationsEditApplicationItemNetworkApplications(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsEditApplicationItemNetworkApplications(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationNetworkApplications {
	request := catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationNetworkApplications{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".app_protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".app_protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".app_protocol")))) {
		request.AppProtocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_sub_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_sub_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_sub_type")))) {
		request.ApplicationSubType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".application_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".application_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".application_type")))) {
		request.ApplicationType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".category_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".category_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".category_id")))) {
		request.CategoryID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".engine_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".engine_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".engine_id")))) {
		request.EngineID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".help_string")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".help_string")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".help_string")))) {
		request.HelpString = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".long_description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".long_description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".long_description")))) {
		request.LongDescription = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".popularity")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".popularity")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".popularity")))) {
		request.Popularity = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rank")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rank")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rank")))) {
		request.Rank = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".traffic_class")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".traffic_class")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".traffic_class")))) {
		request.TrafficClass = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".server_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".server_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".server_name")))) {
		request.ServerName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dscp")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dscp")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dscp")))) {
		request.Dscp = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ignore_conflict")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ignore_conflict")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ignore_conflict")))) {
		request.IgnoreConflict = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsEditApplicationItemNetworkIDentityArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationNetworkIDentity {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationNetworkIDentity{}
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
		i := expandRequestApplicationsEditApplicationItemNetworkIDentity(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsEditApplicationItemNetworkIDentity(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationNetworkIDentity {
	request := catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationNetworkIDentity{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".display_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".display_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".display_name")))) {
		request.DisplayName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".lower_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".lower_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".lower_port")))) {
		request.LowerPort = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ports")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ports")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ports")))) {
		request.Ports = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upper_port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upper_port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upper_port")))) {
		request.UpperPort = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationsEditApplicationItemApplicationSet(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationApplicationSet {
	request := catalystcentersdkgo.RequestItemApplicationPolicyEditApplicationApplicationSet{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id_ref")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id_ref")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id_ref")))) {
		request.IDRef = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchApplicationPolicyGetApplications(m interface{}, queryParams catalystcentersdkgo.GetApplicationsQueryParams, vID string) (*catalystcentersdkgo.ResponseItemApplicationPolicyGetApplications, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error

	var ite *catalystcentersdkgo.ResponseApplicationPolicyGetApplications
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.ApplicationPolicy.GetApplications(nil)
		if err != nil {
			return nil, err
		}
		maxPageSize := len(*nResponse)
		for len(*nResponse) > 0 {

			for _, item := range *nResponse {
				if vID == item.ID {
					return &item, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.ApplicationPolicy.GetApplications(&queryParams)
			if nResponse == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.Name != "" {
		ite, _, err = client.ApplicationPolicy.GetApplications(&queryParams)
		if err != nil || ite == nil {
			return nil, err
		}
		itemsCopy := *ite
		if itemsCopy == nil {
			return nil, err
		}
		for _, item := range itemsCopy {
			if item.Name == queryParams.Name {
				return &item, err
			}
		}
		return nil, err
	}
	return nil, err
}
