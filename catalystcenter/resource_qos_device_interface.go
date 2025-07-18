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

func resourceQosDeviceInterface() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Application Policy.

- Update existing qos device interface infos associate with network device id

- Create qos device interface infos associate with network device id to allow the user to mark specific interfaces as
WAN, to associate WAN interfaces with specific SP Profile and to be able to define a shaper on WAN interfaces

- Delete all qos device interface infos associate with network device id
`,

		CreateContext: resourceQosDeviceInterfaceCreate,
		ReadContext:   resourceQosDeviceInterfaceRead,
		UpdateContext: resourceQosDeviceInterfaceUpdate,
		DeleteContext: resourceQosDeviceInterfaceDelete,
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

						"cfs_change_info": &schema.Schema{
							Description: `Cfs change info
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"create_time": &schema.Schema{
							Description: `Create time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"custom_provisions": &schema.Schema{
							Description: `Custom provisions
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"deployed": &schema.Schema{
							Description: `Deployed
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"display_name": &schema.Schema{
							Description: `Display name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"excluded_interfaces": &schema.Schema{
							Description: `excluded interfaces ids
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"id": &schema.Schema{
							Description: `Id of Qos device info
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"instance_created_on": &schema.Schema{
							Description: `Instance created on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instance_id": &schema.Schema{
							Description: `Instance id
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instance_updated_on": &schema.Schema{
							Description: `Instance updated on
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"instance_version": &schema.Schema{
							Description: `Instance version
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"is_excluded": &schema.Schema{
							Description: `Is excluded
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_seeded": &schema.Schema{
							Description: `Is seeded
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_stale": &schema.Schema{
							Description: `Is stale
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"last_update_time": &schema.Schema{
							Description: `Last update time
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Device name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"namespace": &schema.Schema{
							Description: `Namespace
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"network_device_id": &schema.Schema{
							Description: `Network device id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"provisioning_state": &schema.Schema{
							Description: `Provisioning state
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"qos_device_interface_info": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"display_name": &schema.Schema{
										Description: `Display name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"dmvpn_remote_sites_bw": &schema.Schema{
										Description: `Dmvpn remote sites bandwidth
`,
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeInt,
										},
									},
									"download_bw": &schema.Schema{
										Description: `Download bandwidth
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `Id of Qos device interface info
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"instance_created_on": &schema.Schema{
										Description: `Instance created on
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"instance_id": &schema.Schema{
										Description: `Instance id
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
									"instance_updated_on": &schema.Schema{
										Description: `Instance updated on
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
									"interface_id": &schema.Schema{
										Description: `Interface id
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"interface_name": &schema.Schema{
										Description: `Interface name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"label": &schema.Schema{
										Description: `SP Profile name
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"role": &schema.Schema{
										Description: `Interface role
`,
										Type:     schema.TypeString,
										Computed: true,
									},
									"upload_bw": &schema.Schema{
										Description: `Upload bandwidth
`,
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"qualifier": &schema.Schema{
							Description: `Qualifier
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_version": &schema.Schema{
							Description: `Resource version
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"target_id_list": &schema.Schema{
							Description: `Target id list
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Description: `Type
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestApplicationPolicyCreateQosDeviceInterfaceInfo`,
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

									"excluded_interfaces": &schema.Schema{
										Description: `Excluded interfaces ids
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"id": &schema.Schema{
										Description: `Id of Qos device info
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"name": &schema.Schema{
										Description: `Device name
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"network_device_id": &schema.Schema{
										Description: `Network device id
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"qos_device_interface_info": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"dmvpn_remote_sites_bw": &schema.Schema{
													Description: `Dmvpn remote sites bandwidth
`,
													Type:     schema.TypeList,
													Optional: true,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeInt,
													},
												},
												"instance_id": &schema.Schema{
													Description: `Instance id
`,
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"interface_id": &schema.Schema{
													Description: `Interface id
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"interface_name": &schema.Schema{
													Description: `Interface name
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"label": &schema.Schema{
													Description: `SP Profile name
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"role": &schema.Schema{
													Description: `Interface role
`,
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"upload_bw": &schema.Schema{
													Description: `Upload bandwidth
`,
													Type:     schema.TypeInt,
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
				},
			},
		},
	}
}

func resourceQosDeviceInterfaceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfo(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vNetworkDeviceID := resourceItem["network_device_id"]
	vvNetworkDeviceID := interfaceToString(vNetworkDeviceID)
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	queryParamImport := catalystcentersdkgo.GetQosDeviceInterfaceInfoQueryParams{}
	queryParamImport.NetworkDeviceID = vvNetworkDeviceID
	item2, err := searchApplicationPolicyGetQosDeviceInterfaceInfo(m, queryParamImport, vvName, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["network_device_id"] = vvNetworkDeviceID
		resourceMap["id"] = item2.ID
		resourceMap["name"] = item2.Name
		d.SetId(joinResourceID(resourceMap))
		return resourceQosDeviceInterfaceRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.ApplicationPolicy.CreateQosDeviceInterfaceInfo(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateQosDeviceInterfaceInfo", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateQosDeviceInterfaceInfo", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing CreateQosDeviceInterfaceInfo", err))
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
				"Failure when executing CreateQosDeviceInterfaceInfo", err1))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetQosDeviceInterfaceInfoQueryParams{}
	queryParamValidate.NetworkDeviceID = vvNetworkDeviceID
	item3, err := searchApplicationPolicyGetQosDeviceInterfaceInfo(m, queryParamValidate, vvName, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateQosDeviceInterfaceInfo", err,
			"Failure at CreateQosDeviceInterfaceInfo, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["network_device_id"] = vvNetworkDeviceID
	resourceMap["id"] = item3.ID
	resourceMap["name"] = item3.Name

	d.SetId(joinResourceID(resourceMap))
	return resourceQosDeviceInterfaceRead(ctx, d, m)
}

func resourceQosDeviceInterfaceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	vvNetworkDeviceID := resourceMap["network_device_id"]
	vvName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetQosDeviceInterfaceInfo")
		queryParams1 := catalystcentersdkgo.GetQosDeviceInterfaceInfoQueryParams{}
		queryParams1.NetworkDeviceID = vvNetworkDeviceID
		item1, err := searchApplicationPolicyGetQosDeviceInterfaceInfo(m, queryParams1, vvName, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []catalystcentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponse{
			*item1,
		}
		vItem1 := flattenApplicationPolicyGetQosDeviceInterfaceInfoItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetQosDeviceInterfaceInfo search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceQosDeviceInterfaceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfo(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.ApplicationPolicy.UpdateQosDeviceInterfaceInfo(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateQosDeviceInterfaceInfo", err, restyResp1.String(),
					"Failure at UpdateQosDeviceInterfaceInfo, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateQosDeviceInterfaceInfo", err,
				"Failure at UpdateQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateQosDeviceInterfaceInfo", err))
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
					"Failure when executing UpdateQosDeviceInterfaceInfo", err1))
				return diags
			}
		}

	}

	return resourceQosDeviceInterfaceRead(ctx, d, m)
}

func resourceQosDeviceInterfaceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]

	response1, restyResp1, err := client.ApplicationPolicy.DeleteQosDeviceInterfaceInfo(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteQosDeviceInterfaceInfo", err, restyResp1.String(),
				"Failure at DeleteQosDeviceInterfaceInfo, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteQosDeviceInterfaceInfo", err,
			"Failure at DeleteQosDeviceInterfaceInfo, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteQosDeviceInterfaceInfo", err))
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
				"Failure when executing DeleteQosDeviceInterfaceInfo", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfo {
	request := catalystcentersdkgo.RequestApplicationPolicyCreateQosDeviceInterfaceInfo{}
	if v := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceCreateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo {
	request := catalystcentersdkgo.RequestItemApplicationPolicyCreateQosDeviceInterfaceInfoQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = interfaceToSliceInt(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfo {
	request := catalystcentersdkgo.RequestApplicationPolicyUpdateQosDeviceInterfaceInfo{}
	if v := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo {
	request := catalystcentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".excluded_interfaces")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".excluded_interfaces")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".excluded_interfaces")))) {
		request.ExcludedInterfaces = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".qos_device_interface_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".qos_device_interface_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".qos_device_interface_info")))) {
		request.QosDeviceInterfaceInfo = expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfoArray(ctx, key+".qos_device_interface_info", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfoArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo {
	request := []catalystcentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo{}
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
		i := expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfo(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestQosDeviceInterfaceUpdateQosDeviceInterfaceInfoItemQosDeviceInterfaceInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo {
	request := catalystcentersdkgo.RequestItemApplicationPolicyUpdateQosDeviceInterfaceInfoQosDeviceInterfaceInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".instance_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".instance_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".instance_id")))) {
		request.InstanceID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dmvpn_remote_sites_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dmvpn_remote_sites_bw")))) {
		request.DmvpnRemoteSitesBw = interfaceToSliceInt(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_id")))) {
		request.InterfaceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_name")))) {
		request.InterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".label")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".label")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".label")))) {
		request.Label = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".role")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".role")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".role")))) {
		request.Role = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".upload_bw")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".upload_bw")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".upload_bw")))) {
		request.UploadBW = interfaceToIntPtr(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchApplicationPolicyGetQosDeviceInterfaceInfo(m interface{}, queryParams catalystcentersdkgo.GetQosDeviceInterfaceInfoQueryParams, vName string, vID string) (*catalystcentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponse, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponse
	var ite *catalystcentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfo
	ite, _, err = client.ApplicationPolicy.GetQosDeviceInterfaceInfo(&queryParams)
	if err != nil || ite == nil {
		return foundItem, err

	}
	items := ite
	itemsCopy := *items.Response
	for _, item := range itemsCopy {
		// Call get by _ method and set value to foundItem and return
		if item.Name == vName || item.ID == vID {
			var getItem *catalystcentersdkgo.ResponseApplicationPolicyGetQosDeviceInterfaceInfoResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return foundItem, err
}
