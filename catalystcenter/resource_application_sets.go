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

func resourceApplicationSets() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and delete operations on Application Policy.

- Delete existing application-set by it's id

- Create new custom application-set/s
`,

		CreateContext: resourceApplicationSetsCreate,
		ReadContext:   resourceApplicationSetsRead,
		UpdateContext: resourceApplicationSetsUpdate,
		DeleteContext: resourceApplicationSetsDelete,
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

						"id": &schema.Schema{
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"identity_source": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Id`,
										Type:        schema.TypeString,
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
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateApplicationSetV1`,
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"payload": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							MaxItems: 1,
							MinItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Description: `Name`,
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "",
									},
								}}},
					},
				},
			},
		},
	}
}

func resourceApplicationSetsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] Beginning ApplicationSets Create")
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestApplicationSetsCreateApplicationSetV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	queryParamImport := catalystcentersdkgo.GetApplicationSetsV1QueryParams{}
	queryParamImport.Name = vvName
	item2, err := searchApplicationPolicyGetApplicationSetsV1(m, queryParamImport, "")
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = vvName
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceApplicationSetsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.ApplicationPolicy.CreateApplicationSet(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateApplicationSet", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationSet", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateApplicationSet", err))
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
				"Failure when executing CreateApplicationSet", err1))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetApplicationSetsV1QueryParams{}
	queryParamValidate.Name = vvName
	item3, err := searchApplicationPolicyGetApplicationSetsV1(m, queryParamValidate, "")
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
	return resourceApplicationSetsRead(ctx, d, m)
}

func resourceApplicationSetsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvName := resourceMap["name"]
	vvID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetApplicationSets")
		queryParams1 := catalystcentersdkgo.GetApplicationSetsV1QueryParams{}
		queryParams1.Name = vvName
		item1, err := searchApplicationPolicyGetApplicationSetsV1(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSetsV1Response{
			*item1,
		}
		vItem1 := flattenApplicationPolicyGetApplicationSetsV1Items(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetApplicationSets search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceApplicationSetsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing ApplicationSetsUpdate", err, "Update method is not supported",
		"Failure at ApplicationSetsUpdate, unexpected response", ""))

	return diags
}

func resourceApplicationSetsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	response1, restyResp1, err := client.ApplicationPolicy.DeleteApplicationSet(&catalystcentersdkgo.DeleteApplicationSetV1QueryParams{
		ID: vvID,
	})
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
func expandRequestApplicationSetsCreateApplicationSetV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyCreateApplicationSetV1 {
	request := catalystcentersdkgo.RequestApplicationPolicyCreateApplicationSetV1{}
	if v := expandRequestApplicationSetsCreateApplicationSetV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationSetsCreateApplicationSetV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationSetV1 {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationSetV1{}
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
		i := expandRequestApplicationSetsCreateApplicationSetV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestApplicationSetsCreateApplicationSetV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationSetV1 {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateApplicationSetV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchApplicationPolicyGetApplicationSetsV1(m interface{}, queryParams catalystcentersdkgo.GetApplicationSetsV1QueryParams, vID string) (*catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSetsV1Response, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSetsV1Response
	var ite *catalystcentersdkgo.ResponseApplicationPolicyGetApplicationSetsV1
	if vID != "" {
		queryParams.Offset = 1
		queryParams.Name = ""
		nResponse, _, err := client.ApplicationPolicy.GetApplicationSets(&queryParams)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.ApplicationPolicy.GetApplicationSetsV1(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.Name != "" {
		ite, _, err = client.ApplicationPolicy.GetApplicationSetsV1(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
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
