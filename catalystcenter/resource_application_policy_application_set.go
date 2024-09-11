package catalystcenter

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplicationPolicyApplicationSet() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Application Policy.

- Create new custom application set/s

- Delete existing custom application set by id
`,

		CreateContext: resourceApplicationPolicyApplicationSetCreate,
		ReadContext:   resourceApplicationPolicyApplicationSetRead,
		UpdateContext: resourceApplicationPolicyApplicationSetUpdate,
		DeleteContext: resourceApplicationPolicyApplicationSetDelete,
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

						"default_business_relevance": &schema.Schema{
							Description: `Default business relevance
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": &schema.Schema{
							Description: `Display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Id of Application Set
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"identity_source": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": &schema.Schema{
										Description: `Type of identify source. NBAR: build in Application Set, APIC_EM: custom Application Set
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"instance_id": &schema.Schema{
							Description: `Instance id
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instance_version": &schema.Schema{
							Description: `Instance version
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Application Set name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": &schema.Schema{
							Description: `Namespace, valid value scalablegroup:application
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"scalable_group_external_handle": &schema.Schema{
							Description: `Scalable group external handle, should be equal to Application Set name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"scalable_group_type": &schema.Schema{
							Description: `Scalable group type, valid value APPLICATION_GROUP
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Type, valid value scalablegroup
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationSets`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"payload": &schema.Schema{
							Description: `Array of RequestApplicationPolicyCreateApplicationSets`,
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"default_business_relevance": &schema.Schema{
										Description: `Default business relevance
			`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `id path parameter. Id of custom application set to delete
			`,
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": &schema.Schema{
										Description: `Application Set name
			`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"namespace": &schema.Schema{
										Description: `Namespace, should be set to scalablegroup:application
			`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"qualifier": &schema.Schema{
										Description: `Qualifier, should be set to application
			`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"scalable_group_external_handle": &schema.Schema{
										Description: `Scalable group external handle, should be set to application set name
			`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"scalable_group_type": &schema.Schema{
										Description: `Scalable group type, should be set to APPLICATION_GROUP
			`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Description: `Type, should be set to scalablegroup
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
		},
	}
}

func resourceApplicationPolicyApplicationSetCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestApplicationPolicyApplicationSetCreateApplicationSets(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	queryParamImport := catalystcentersdkgo.GetApplicationSets2QueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchApplicationPolicyGetApplicationSets2(m, queryParamImport, "")
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceApplicationSetsRead(ctx, d, m)
	}

	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationSets(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationSets", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationSets", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationSets", err))
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
				"Failure when executing CreateApplicationSets", err1))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetApplicationSets2QueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchApplicationPolicyGetApplicationSets2(m, queryParamValidate, "")
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateApplicationSet", err,
			"Failure at CreateApplicationSet, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = vvName
	resourceMap["id"] = item3.ID

	d.SetId(joinResourceID(resourceMap))
	return resourceApplicationPolicyApplicationSetRead(ctx, d, m)
}

func resourceApplicationPolicyApplicationSetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvName := resourceMap["name"]
	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplicationSets")
		queryParams1 := catalystcentersdkgo.GetApplicationSets2QueryParams{}
		queryParams1.Name = vvName
		item1, err := searchApplicationPolicyGetApplicationSets2(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSets2Response{
			*item1,
		}
		vItem1 := flattenApplicationPolicyGetApplicationSets2Items(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationSets search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceApplicationPolicyApplicationSetUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceApplicationPolicyApplicationSetRead(ctx, d, m)
}

func resourceApplicationPolicyApplicationSetDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplicationSet(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteApplicationSet", err, restyResp1.String(),
				"Failure at DeleteApplicationSet, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteApplicationSet", err,
			"Failure at DeleteApplicationSet, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteApplicationSet", err))
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
				"Failure when executing DeleteApplicationSet", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestApplicationPolicyApplicationSetCreateApplicationSets(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyCreateApplicationSets {
	request := catalystcentersdkgo.RequestApplicationPolicyCreateApplicationSets{}
	if v := expandRequestApplicationPolicyApplicationSetCreateApplicationSetsItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationPolicyApplicationSetCreateApplicationSetsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationSets {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationSets{}
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
		i := expandRequestApplicationPolicyApplicationSetCreateApplicationSetsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationPolicyApplicationSetCreateApplicationSetsItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationSets {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationSets{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_type")))) {
		request.ScalableGroupType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".default_business_relevance")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".default_business_relevance")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".default_business_relevance")))) {
		request.DefaultBusinessRelevance = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".namespace")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".namespace")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".namespace")))) {
		request.Namespace = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qualifier")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qualifier")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qualifier")))) {
		request.Qualifier = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".scalable_group_external_handle")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".scalable_group_external_handle")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".scalable_group_external_handle")))) {
		request.ScalableGroupExternalHandle = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchApplicationPolicyGetApplicationSets2(m interface{}, queryParams catalystcentersdkgo.GetApplicationSets2QueryParams, vID string) (*catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSets2Response, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSets2Response
	var ite *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSets2
	ite, _, err = client.ApplicationPolicy.GetApplicationSets2(&queryParams)
	if err != nil || ite == nil {
		return foundItem, err

	}
	items := ite
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == queryParams.Name {
			var getItem *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSets2Response
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
