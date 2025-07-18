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

func resourceGlobalCredentialHTTPRead() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read and update operations on Discovery.

- Adds HTTP read credentials

- Updates global HTTP Read credential

- Deletes global credential for the given ID
`,

		CreateContext: resourceGlobalCredentialHTTPReadCreate,
		ReadContext:   resourceGlobalCredentialHTTPReadRead,
		UpdateContext: resourceGlobalCredentialHTTPReadUpdate,
		DeleteContext: resourceGlobalCredentialHTTPReadDelete,
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

						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},

						"password": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},

						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},

						"secure": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"username": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestDiscoveryCreateHTTPReadCredentials`,
				Type:        schema.TypeList,
				MaxItems:    1,
				MinItems:    1,
				Required:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"comments": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"credential_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Default:  "",
						},
						"instance_tenant_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"instance_uuid": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"port": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"secure": &schema.Schema{
							// Type:     schema.TypeBool,
							Type:         schema.TypeString,
							ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
							Optional:     true,
						},
						"username": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
		},
	}
}

func resourceGlobalCredentialHTTPReadCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestGlobalCredentialHTTPReadCreateHTTPReadCredentials(ctx, "parameters", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vUsername := resourceItem["username"]
	vvUsername := interfaceToString(vUsername)
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	queryParams1 := catalystcentersdkgo.GetGlobalCredentialsQueryParams{}

	queryParams1.CredentialSubType = "HTTP_READ"
	item, err := searchDiscoveryGetGlobalCredentialsHttpRead(m, queryParams1, vvUsername, vvID)
	if item != nil || err != nil {
		resourceMap := make(map[string]string)
		resourceMap["username"] = vvUsername
		resourceMap["id"] = vvID
		d.SetId(joinResourceID(resourceMap))
		return resourceGlobalCredentialHTTPReadRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Discovery.CreateHTTPReadCredentials(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateHTTPReadCredentials", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateHTTPReadCredentials", err))
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
			errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateHTTPReadCredentials", err1))
			return diags
		}
	}
	resourceMap := make(map[string]string)
	resourceMap["username"] = vvUsername
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceGlobalCredentialHTTPReadRead(ctx, d, m)
}

func resourceGlobalCredentialHTTPReadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := "HTTP_READ"
	vUsername := resourceMap["username"]
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetGlobalCredentials")
		queryParams1 := catalystcentersdkgo.GetGlobalCredentialsQueryParams{}

		queryParams1.CredentialSubType = vCredentialSubType

		response1, err := searchDiscoveryGetGlobalCredentialsHttpRead(m, queryParams1, vUsername, vID)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalCredentials search response",
				err))
			return diags
		}
		if response1 == nil {
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR CATALYST
		items := []catalystcentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse{
			*response1,
		}
		vItem1 := flattenDiscoveryGetGlobalCredentialsItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetGlobalCredentials search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceGlobalCredentialHTTPReadUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vCredentialSubType := "HTTP_READ"
	vUsername := resourceMap["username"]
	vID := resourceMap["id"]

	queryParams1 := catalystcentersdkgo.GetGlobalCredentialsQueryParams{}
	queryParams1.CredentialSubType = vCredentialSubType
	response1, err := searchDiscoveryGetGlobalCredentialsHttpRead(m, queryParams1, vUsername, vID)
	if err != nil || response1 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalCredentials", err,
			"Failure at GetGlobalCredentials, unexpected response", ""))
		return diags
	}
	var vvName string
	// NOTE: Consider adding getAllItems and search function to get missing params
	// if selectedMethod == 1 { }
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] Name used for update operation %s", vvName)
		request1 := expandRequestGlobalCredentialHTTPReadUpdateHTTPReadCredential(ctx, "parameters.0", d)
		request1.ID = response1.ID
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Discovery.UpdateHTTPReadCredential(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateHTTPReadCredential", err, restyResp1.String(),
					"Failure at UpdateHTTPReadCredential, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateHTTPReadCredential", err,
				"Failure at UpdateHTTPReadCredential, unexpected response", ""))
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
				errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
				err1 := errors.New(errorMsg)
				diags = append(diags, diagError(
					"Failure when executing UpdateHTTPReadCredential", err1))
				return diags
			}
		}

	}

	return resourceGlobalCredentialHTTPReadRead(ctx, d, m)
}

func resourceGlobalCredentialHTTPReadDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// NOTE: Unable to delete GlobalCredentialSNMPv2ReadCommunity on Catalyst Center
	//       Returning empty diags to delete it on Terraform
	// DeleteGlobalCredentialsByID
	client := m.(*catalystcentersdkgo.Client)

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	vUsername := resourceMap["username"]

	queryParams1 := catalystcentersdkgo.GetGlobalCredentialsQueryParams{}

	queryParams1.CredentialSubType = "HTTP_READ"
	item, err := searchDiscoveryGetGlobalCredentialsHttpRead(m, queryParams1, vUsername, vID)
	if item == nil && err != nil {
		return resourceGlobalCredentialHTTPReadRead(ctx, d, m)
	}
	if vID != item.ID {
		vID = item.ID
	}
	if vID != "" {
		response1, restyResp1, err := client.Discovery.DeleteGlobalCredentialsByID(vID)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeleteGlobalCredentialsByID", err,
				"Failure at DeleteGlobalCredentialsByID, unexpected response", ""))
			return diags
		}
	}
	return diags
}
func expandRequestGlobalCredentialHTTPReadCreateHTTPReadCredentials(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDiscoveryCreateHTTPReadCredentials {
	request := catalystcentersdkgo.RequestDiscoveryCreateHTTPReadCredentials{}
	if v := expandRequestGlobalCredentialHTTPReadCreateHTTPReadCredentialsItemArray(ctx, key+".", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialHTTPReadCreateHTTPReadCredentialsItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemDiscoveryCreateHTTPReadCredentials {
	request := []catalystcentersdkgo.RequestItemDiscoveryCreateHTTPReadCredentials{}
	key = fixKeyAccess(key)
	o := d.Get(key)
	if o == nil {
		return nil
	}
	objs := o.([]interface{})
	if len(objs) == 0 {
		return nil
	}
	for item_no, _ := range objs {
		i := expandRequestGlobalCredentialHTTPReadCreateHTTPReadCredentialsItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialHTTPReadCreateHTTPReadCredentialsItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemDiscoveryCreateHTTPReadCredentials {
	request := catalystcentersdkgo.RequestItemDiscoveryCreateHTTPReadCredentials{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func expandRequestGlobalCredentialHTTPReadUpdateHTTPReadCredential(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDiscoveryUpdateHTTPReadCredential {
	request := catalystcentersdkgo.RequestDiscoveryUpdateHTTPReadCredential{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".comments")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".comments")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".comments")))) {
		request.Comments = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".credential_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".credential_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".credential_type")))) {
		request.CredentialType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".password")))) {
		request.Password = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port")))) {
		request.Port = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".secure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".secure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".secure")))) {
		request.Secure = interfaceToBoolPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".username")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".username")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".username")))) {
		request.Username = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchDiscoveryGetGlobalCredentialsHttpRead(m interface{}, queryParams catalystcentersdkgo.GetGlobalCredentialsQueryParams, vUsername string, vID string) (*catalystcentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse
	var ite *catalystcentersdkgo.ResponseDiscoveryGetGlobalCredentials
	queryParams.CredentialSubType = "HTTP_READ"
	ite, _, err = client.Discovery.GetGlobalCredentials(&queryParams)
	if err != nil {
		return foundItem, err
	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items
	for _, item := range *itemsCopy.Response {
		// Call get by _ method and set value to foundItem and return
		if item.ID == vID || item.Username == vUsername {
			var getItem *catalystcentersdkgo.ResponseDiscoveryGetGlobalCredentialsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
