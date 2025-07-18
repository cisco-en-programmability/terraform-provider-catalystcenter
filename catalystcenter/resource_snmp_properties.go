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

func resourceSNMPProperties() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create and read operations on Discovery.

- Adds SNMP properties
`,

		CreateContext: resourceSNMPPropertiesCreate,
		ReadContext:   resourceSNMPPropertiesRead,
		UpdateContext: resourceSNMPPropertiesUpdate,
		DeleteContext: resourceSNMPPropertiesDelete,
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
							Description: `Id of the SNMP Property
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_tenant_id": &schema.Schema{
							Description: `[Deprecated] InstanceTenantId of the SNMP Property
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_uuid": &schema.Schema{
							Description: `Instance Uuid of the SNMP Property. It is the same as the id. It will be deprecated in future version.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"int_value": &schema.Schema{
							Description: `Integer Value of the SNMP 'Retry' or 'Timeout' property
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"system_property_name": &schema.Schema{
							Description: `Name of the SNMP Property as 'Retry' or 'Timeout'
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestDiscoveryCreateUpdateSNMPProperties`,
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
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"instance_tenant_id": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"instance_uuid": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"int_value": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"system_property_name": &schema.Schema{
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

func resourceSNMPPropertiesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSNMPPropertiesCreateUpdateSNMPProperties(ctx, "parameters.0", d)
	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}
	vInstanceTenantId := resourceItem["instance_tenant_id"]
	vSystemPropertyName := resourceItem["system_property_name"]
	vvInstanceTenantId := interfaceToString(vInstanceTenantId)
	vvSystemPropertyName := interfaceToString(vSystemPropertyName)

	item, err := searchDiscoveryGetSNMPProperties(m, vvInstanceTenantId, vvSystemPropertyName)

	if item != nil || err != nil {
		resourceMap := make(map[string]string)
		resourceMap["instance_tenant_id"] = vvInstanceTenantId
		resourceMap["system_property_name"] = vvSystemPropertyName
		d.SetId(joinResourceID(resourceMap))
		return resourceSNMPPropertiesRead(ctx, d, m)
	}

	resp1, restyResp1, err := client.Discovery.CreateUpdateSNMPProperties(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateUpdateSNMPProperties", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateUpdateSNMPProperties", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateUpdateSNMPProperties", err))
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
			log.Printf("[DEBUG] Error => %v", response2.Response.FailureReason)
			errorMsg := response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing CreateUpdateSNMPProperties", err1))
			return diags
		}
	}

	resourceMap := make(map[string]string)
	resourceMap["instance_tenant_id"] = vvInstanceTenantId
	resourceMap["system_property_name"] = vvSystemPropertyName
	d.SetId(joinResourceID(resourceMap))
	return resourceSNMPPropertiesRead(ctx, d, m)
}

func resourceSNMPPropertiesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vInstanceTenantId := resourceMap["instance_tenant_id"]
	vSystemPropertyName := resourceMap["system_property_name"]
	var diags diag.Diagnostics

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: GetSNMPProperties")

		response1, err := searchDiscoveryGetSNMPProperties(m, vInstanceTenantId, vSystemPropertyName)
		if err != nil {
			diags = append(diags, diagError(
				"Failure when setting searchDiscoveryGetSNMPProperties search response",
				err))
			return diags
		}
		if response1 == nil {
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
		items := []catalystcentersdkgo.ResponseDiscoveryGetSNMPPropertiesResponse{
			*response1,
		}
		vItem1 := flattenDiscoveryGetSNMPPropertiesItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetSNMPProperties search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSNMPPropertiesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Update not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SNMPPropertiesUpdate", err, "Update method is not supported",
		"Failure at SNMPPropertiesUpdate, unexpected response", ""))

	return diags
}

func resourceSNMPPropertiesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	err := errors.New("Delete not possible in this resource")
	diags = append(diags, diagErrorWithAltAndResponse(
		"Failure when executing SNMPPropertiesDelete", err, "Delete method is not supported",
		"Failure at SNMPPropertiesDelete, unexpected response", ""))

	return diags
}

func expandRequestSNMPPropertiesCreateUpdateSNMPProperties(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDiscoveryCreateUpdateSNMPProperties {
	request := catalystcentersdkgo.RequestDiscoveryCreateUpdateSNMPProperties{}
	if v := expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemDiscoveryCreateUpdateSNMPProperties {
	request := []catalystcentersdkgo.RequestItemDiscoveryCreateUpdateSNMPProperties{}
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
		i := expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSNMPPropertiesCreateUpdateSNMPPropertiesItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemDiscoveryCreateUpdateSNMPProperties {
	request := catalystcentersdkgo.RequestItemDiscoveryCreateUpdateSNMPProperties{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_tenant_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_tenant_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_tenant_id")))) {
		request.InstanceTenantID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_uuid")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_uuid")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_uuid")))) {
		request.InstanceUUID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".int_value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".int_value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".int_value")))) {
		request.IntValue = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".system_property_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".system_property_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".system_property_name")))) {
		request.SystemPropertyName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchDiscoveryGetSNMPProperties(m interface{}, vID string, vName string) (*catalystcentersdkgo.ResponseDiscoveryGetSNMPPropertiesResponse, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseDiscoveryGetSNMPPropertiesResponse
	var ite *catalystcentersdkgo.ResponseDiscoveryGetSNMPProperties
	ite, _, err = client.Discovery.GetSNMPProperties()
	if err != nil {
		return foundItem, err

	}
	items := ite
	if items == nil {
		return foundItem, err
	}
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.InstanceTenantID == vID || item.SystemPropertyName == vName {
			var getItem *catalystcentersdkgo.ResponseDiscoveryGetSNMPPropertiesResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
