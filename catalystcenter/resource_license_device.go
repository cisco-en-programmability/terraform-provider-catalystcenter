package catalystcenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceLicenseDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Licenses.

- Deregister device(s) from CSSM(Cisco Smart Software Manager).

- Register device(s) in CSSM(Cisco Smart Software Manager).

- Transfer device(s) from one virtual account to another within same smart account.
`,

		CreateContext: resourceLicenseDeviceCreate,
		ReadContext:   resourceLicenseDeviceRead,
		UpdateContext: resourceLicenseDeviceUpdate,
		DeleteContext: resourceLicenseDeviceDelete,
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

						"domain": &schema.Schema{
							Description: `Domain of smart account
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Id of smart account
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"is_active_smart_account": &schema.Schema{
							Description: `Is active smart account
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Name of smart account
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"device_uuids": &schema.Schema{
							Description: `Comma separated device ids
`,
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"smart_account_id": &schema.Schema{
							Description: `smart_account_id path parameter. Id of smart account
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"virtual_account_name": &schema.Schema{
							Description: `virtual_account_name path parameter. Name of target virtual account
`,
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceLicenseDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*catalystcentersdkgo.Client)
	resourceItem := *getResourceItem(d.Get("parameters"))
	vSmartAccountID := resourceItem["smart_account_id"]
	vvSmartAccountID := interfaceToString(vSmartAccountID)
	vVirtualAccountName := resourceItem["virtual_account_name"]
	vvVirtualAccountName := interfaceToString(vVirtualAccountName)
	request1 := expandRequestLicenseDeviceChangeVirtualAccount(ctx, "parameters.0", d)
	response1, err := searchLicensesSmartAccountDetails(m, vvVirtualAccountName, vvSmartAccountID)

	if err != nil || response1 != nil {
		resourceMap := make(map[string]string)
		resourceMap["smart_account_id"] = interfaceToString(resourceItem["smart_account_id"])
		resourceMap["virtual_account_name"] = interfaceToString(resourceItem["virtual_account_name"])
		d.SetId(joinResourceID(resourceMap))
		return resourceLicenseDeviceRead(ctx, d, m)
	}

	response2, restyResp1, err := client.Licenses.ChangeVirtualAccount(vvSmartAccountID, vvVirtualAccountName, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response2 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ChangeVirtualAccount2", err,
			"Failure at ChangeVirtualAccount2, unexpected response", ""))
		return diags
	}

	taskId := response2.Response.TaskID
	log.Printf("[DEBUG] TASKID => %s", taskId)
	if taskId != "" {
		time.Sleep(5 * time.Second)
		response3, restyResp2, err := client.Task.GetTaskByID(taskId)
		if err != nil || response3 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetTaskByID", err,
				"Failure at GetTaskByID, unexpected response", ""))
			return diags
		}
		if response3.Response != nil && response3.Response.IsError != nil && *response3.Response.IsError {
			log.Printf("[DEBUG] Error reason %s", response3.Response.FailureReason)
			errorMsg := response3.Response.Progress + "\nFailure Reason: " + response3.Response.FailureReason
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ChangeVirtualAccount2", err1))
			return diags
		}
	}

	resourceMap := make(map[string]string)
	resourceMap["smart_account_id"] = interfaceToString(resourceItem["smart_account_id"])
	resourceMap["virtual_account_name"] = interfaceToString(resourceItem["virtual_account_name"])
	d.SetId(joinResourceID(resourceMap))
	return resourceLicenseDeviceRead(ctx, d, m)
}

func resourceLicenseDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vSmartAccountID := resourceMap["smart_account_id"]
	vVirtualAccountName := resourceMap["virtual_account_name"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method 1: SmartAccountDetails")

		response1, err := searchLicensesSmartAccountDetails(m, vVirtualAccountName, vSmartAccountID)

		if err != nil {
			diags = append(diags, diagError(
				"Failure when setting searchLicensesSmartAccountDetails search response",
				err))
			return diags
		}
		if response1 == nil {
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		//TODO FOR CATALYST

		vItem1 := flattenLicensesSmartAccountDetailsItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting SmartAccountDetails search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceLicenseDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vSmartAccountID := resourceMap["smart_account_id"]
	vVirtualAccountName := resourceMap["virtual_account_name"]

	item, err := searchLicensesSmartAccountDetails(m, vVirtualAccountName, vSmartAccountID)

	if err != nil || item == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing GetGlobalCredentials", err,
			"Failure at GetGlobalCredentials, unexpected response", ""))
		return diags
	}

	if d.HasChange("parameters") {
		request1 := expandRequestLicenseDeviceDeviceRegistration(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Licenses.DeviceRegistration(vVirtualAccountName, request1)

		if request1 != nil {
			log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		}

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing DeviceRegistration2", err,
				"Failure at DeviceRegistration2, unexpected response", ""))
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
					"Failure when executing DeviceRegistration2", err1))
				return diags
			}
		}
	}

	return resourceLicenseDeviceRead(ctx, d, m)
}

func resourceLicenseDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	client := m.(*catalystcentersdkgo.Client)
	// NOTE: Unable to delete LicenseDevice on Catalyst Center
	//       Returning empty diags to delete it on Terraform
	log.Printf("[DEBUG] Selected method 1: DeviceDeregistration2")
	request1 := expandRequestLicenseDeviceDeregistrationDeviceDeregistration(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Licenses.DeviceDeregistration(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeviceDeregistration2", err,
			"Failure at DeviceDeregistration2, unexpected response", ""))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
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
				"Failure when executing DeviceDeregistration2", err1))
			return diags
		}
	}
	d.SetId("")
	return diags
}

func expandRequestLicenseDeviceChangeVirtualAccount(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestLicensesChangeVirtualAccount {
	request := catalystcentersdkgo.RequestLicensesChangeVirtualAccount{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuids")))) {
		request.DeviceUUIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLicenseDeviceDeviceRegistration(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestLicensesDeviceRegistration {
	request := catalystcentersdkgo.RequestLicensesDeviceRegistration{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuids")))) {
		request.DeviceUUIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestLicenseDeviceDeregistrationDeviceDeregistration(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestLicensesDeviceDeregistration {
	request := catalystcentersdkgo.RequestLicensesDeviceDeregistration{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_uuids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_uuids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_uuids")))) {
		request.DeviceUUIDs = interfaceToSliceString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}

	return &request
}

func searchLicensesSmartAccountDetails(m interface{}, vName string, vID string) (*catalystcentersdkgo.ResponseLicensesSmartAccountDetailsResponse, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseLicensesSmartAccountDetailsResponse
	var ite *catalystcentersdkgo.ResponseLicensesSmartAccountDetails
	ite, _, err = client.Licenses.SmartAccountDetails()
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
		if item.Name == vName && item.ID == vID {
			var getItem *catalystcentersdkgo.ResponseLicensesSmartAccountDetailsResponse
			getItem = &item
			foundItem = getItem
			return foundItem, err
		}
	}
	return nil, err
}
