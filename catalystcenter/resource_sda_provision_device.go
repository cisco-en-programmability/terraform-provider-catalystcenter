package catalystcenter

import (
	"context"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaProvisionDevice() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Delete provisioned Wired Device

- Re-Provision Wired Device

- Provision Wired Device
`,

		CreateContext: resourceSdaProvisionDeviceCreate,
		ReadContext:   resourceSdaProvisionDeviceRead,
		UpdateContext: resourceSdaProvisionDeviceUpdate,
		DeleteContext: resourceSdaProvisionDeviceDelete,
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
							Description: `Wired Provisioned device detail retrieved successfully
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the device to be provisioned
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy for device location(only building / floor level)
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Description: `Status
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

						"device_management_ip_address": &schema.Schema{
							Description: `Management Ip Address of the device to be provisioned
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"site_name_hierarchy": &schema.Schema{
							Description: `Site Name Hierarchy for device location(only building / floor level)
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceSdaProvisionDeviceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestSdaProvisionDeviceProvisionWiredDevice(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vDeviceManagementIPAddress := resourceItem["device_management_ip_address"]
	vvDeviceManagementIPAddress := interfaceToString(vDeviceManagementIPAddress)
	queryParamImport := catalystcentersdkgo.GetProvisionedWiredDeviceQueryParams{}
	queryParamImport.DeviceManagementIPAddress = vvDeviceManagementIPAddress
	item2, _, err := client.Sda.GetProvisionedWiredDevice(&queryParamImport)
	if err != nil || item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["device_management_ip_address"] = item2.DeviceManagementIPAddress
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaProvisionDeviceRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.ProvisionWiredDevice(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing ProvisionWiredDevice", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing ProvisionWiredDevice", err))
		return diags
	}
	executionId := resp1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing ProvisionWiredDevice", err))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetProvisionedWiredDeviceQueryParams{}
	queryParamValidate.DeviceManagementIPAddress = vvDeviceManagementIPAddress
	item3, _, err := client.Sda.GetProvisionedWiredDevice(&queryParamValidate)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing ProvisionWiredDevice", err,
			"Failure at ProvisionWiredDevice, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["device_management_ip_address"] = item3.DeviceManagementIPAddress

	d.SetId(joinResourceID(resourceMap))
	return resourceSdaProvisionDeviceRead(ctx, d, m)
}

func resourceSdaProvisionDeviceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vDeviceManagementIPAddress := resourceMap["device_management_ip_address"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetProvisionedWiredDevice")
		queryParams1 := catalystcentersdkgo.GetProvisionedWiredDeviceQueryParams{}

		queryParams1.DeviceManagementIPAddress = vDeviceManagementIPAddress

		// has_unknown_response: None

		response1, restyResp1, err := client.Sda.GetProvisionedWiredDevice(&queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSdaGetProvisionedWiredDeviceItem(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetProvisionedWiredDevice response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceSdaProvisionDeviceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	if d.HasChange("parameters") {
		request1 := expandRequestSdaProvisionDeviceReProvisionWiredDevice(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.Sda.ReProvisionWiredDevice(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing ReProvisionWiredDevice", err, restyResp1.String(),
					"Failure at ReProvisionWiredDevice, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing ReProvisionWiredDevice", err,
				"Failure at ReProvisionWiredDevice, unexpected response", ""))
			return diags
		}

		executionId := response1.ExecutionID
		log.Printf("[DEBUG] ExecutionID => %s", executionId)
		if executionId != "" {
			time.Sleep(5 * time.Second)
			response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
			for response2.Status == "IN_PROGRESS" {
				time.Sleep(10 * time.Second)
				response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
				if err != nil || response2 == nil {
					if restyResp2 != nil {
						log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
					}
					diags = append(diags, diagErrorWithAlt(
						"Failure when executing GetExecutionByID", err,
						"Failure at GetExecutionByID, unexpected response", ""))
					return diags
				}
			}
			if response2.Status == "FAILURE" {
				log.Printf("[DEBUG] Error %s", response2.BapiError)
				diags = append(diags, diagError(
					"Failure when executing ReProvisionWiredDevice", err))
				return diags
			}
		}

	}

	return resourceSdaProvisionDeviceRead(ctx, d, m)
}

func resourceSdaProvisionDeviceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	queryParamDelete := catalystcentersdkgo.DeleteProvisionedWiredDeviceQueryParams{}

	vvDeviceManagementIPAddress := resourceMap["device_management_ip_address"]
	queryParamDelete.DeviceManagementIPAddress = vvDeviceManagementIPAddress

	response1, restyResp1, err := client.Sda.DeleteProvisionedWiredDevice(&queryParamDelete)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteProvisionedWiredDevice", err, restyResp1.String(),
				"Failure at DeleteProvisionedWiredDevice, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteProvisionedWiredDevice", err,
			"Failure at DeleteProvisionedWiredDevice, unexpected response", ""))
		return diags
	}

	executionId := response1.ExecutionID
	log.Printf("[DEBUG] ExecutionID => %s", executionId)
	if executionId != "" {
		time.Sleep(5 * time.Second)
		response2, restyResp2, err := client.Task.GetBusinessAPIExecutionDetails(executionId)
		if err != nil || response2 == nil {
			if restyResp2 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing GetExecutionByID", err,
				"Failure at GetExecutionByID, unexpected response", ""))
			return diags
		}
		for response2.Status == "IN_PROGRESS" {
			time.Sleep(10 * time.Second)
			response2, restyResp2, err = client.Task.GetBusinessAPIExecutionDetails(executionId)
			if err != nil || response2 == nil {
				if restyResp2 != nil {
					log.Printf("[DEBUG] Retrieved error response %s", restyResp2.String())
				}
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetExecutionByID", err,
					"Failure at GetExecutionByID, unexpected response", ""))
				return diags
			}
		}
		if response2.Status == "FAILURE" {
			log.Printf("[DEBUG] Error %s", response2.BapiError)
			diags = append(diags, diagError(
				"Failure when executing DeleteProvisionedWiredDevice", err))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestSdaProvisionDeviceProvisionWiredDevice(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaProvisionWiredDevice {
	request := catalystcentersdkgo.RequestSdaProvisionWiredDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaProvisionDeviceReProvisionWiredDevice(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaReProvisionWiredDevice {
	request := catalystcentersdkgo.RequestSdaReProvisionWiredDevice{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_management_ip_address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_management_ip_address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_management_ip_address")))) {
		request.DeviceManagementIPAddress = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_name_hierarchy")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_name_hierarchy")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_name_hierarchy")))) {
		request.SiteNameHierarchy = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
