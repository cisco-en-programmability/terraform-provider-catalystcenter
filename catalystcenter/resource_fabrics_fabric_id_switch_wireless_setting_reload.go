package catalystcenter

import (
	"context"

	"errors"

	"time"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceFabricsFabricIDSwitchWirelessSettingReload() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Fabric Wireless.

- This data source action is used to reload switches after disabling wireless to remove the wireless-controller
configuration on the device. When wireless is disabled on a switch, all wireless configurations are removed except for
the wireless-controller configuration. To completely remove the wireless-controller configuration, you can use this API.
Please note that using this API will cause a reload of the device(s). This data source action should only be used for
devices that have wireless disabled but still have the 'wireless-controller' configuration present. The reload payload
can have a maximum of two switches as only two switches can have a wireless role in a fabric site.
`,

		CreateContext: resourceFabricsFabricIDSwitchWirelessSettingReloadCreate,
		ReadContext:   resourceFabricsFabricIDSwitchWirelessSettingReloadRead,
		DeleteContext: resourceFabricsFabricIDSwitchWirelessSettingReloadDelete,
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

						"task_id": &schema.Schema{
							Description: `Task ID
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `Task URL
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				MinItems: 1,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"fabric_id": &schema.Schema{
							Description: `fabricId path parameter. The 'fabricId' represents the Fabric ID of a particular Fabric Site. The 'fabricId' can be obtained using the api /dna/intent/api/v1/sda/fabricSites.  Example : e290f1ee-6c54-4b01-90e6-d701748f0851
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"device_id": &schema.Schema{
							Description: `Network Device ID
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFabricsFabricIDSwitchWirelessSettingReloadCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vFabricID := resourceItem["fabric_id"]
	vvFabricID := vFabricID.(string)
	request1 := expandRequestFabricsFabricIDSwitchWirelessSettingReloadReloadSwitchForWirelessControllerCleanupV1(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.FabricWireless.ReloadSwitchForWirelessControllerCleanupV1(vvFabricID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vItem1 := flattenFabricWirelessReloadSwitchForWirelessControllerCleanupV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ReloadSwitchForWirelessControllerCleanupV1 response",
			err))
		return diags
	}
	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing ReloadSwitchForWirelessControllerCleanupV1", err))
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
			restyResp3, err := client.CustomCall.GetCustomCall(response2.Response.AdditionalStatusURL, nil)
			if err != nil {
				diags = append(diags, diagErrorWithAlt(
					"Failure when executing GetCustomCall", err,
					"Failure at GetCustomCall, unexpected response", ""))
				return diags
			}
			var errorMsg string
			if restyResp3 == nil {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing ReloadSwitchForWirelessControllerCleanupV1", err1))
			return diags
		}
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))
	d.SetId(getUnixTimeString())
	return diags
}
func resourceFabricsFabricIDSwitchWirelessSettingReloadRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceFabricsFabricIDSwitchWirelessSettingReloadDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestFabricsFabricIDSwitchWirelessSettingReloadReloadSwitchForWirelessControllerCleanupV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestFabricWirelessReloadSwitchForWirelessControllerCleanupV1 {
	request := catalystcentersdkgo.RequestFabricWirelessReloadSwitchForWirelessControllerCleanupV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	return &request
}

func flattenFabricWirelessReloadSwitchForWirelessControllerCleanupV1Item(item *catalystcentersdkgo.ResponseFabricWirelessReloadSwitchForWirelessControllerCleanupV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = item.TaskID
	respItem["url"] = item.URL
	return []map[string]interface{}{
		respItem,
	}
}
