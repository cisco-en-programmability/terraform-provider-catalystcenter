package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocations() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Wireless.

- This data source action allows user to assign Anchor Managed AP Locations for WLC by device ID. The payload should
always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition
and deletion.


       When anchor managed location array present then it will add the anchor managed locations.
`,

		CreateContext: resourceWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocationsCreate,
		ReadContext:   resourceWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocationsRead,
		DeleteContext: resourceWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocationsDelete,
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
							Description: `Unique identifier for the task.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `URL for the task.
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
						"network_device_id": &schema.Schema{
							Description: `networkDeviceId path parameter. Network Device ID. This value can be obtained by using the API call GET: /dna/intent/api/v1/network-device/ip-address/${ipAddress}
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"anchor_managed_aplocations_site_ids": &schema.Schema{
							Description: `This API allows user to assign Anchor Managed AP Locations for WLC by device ID. The payload should always be a complete list. The Managed AP Locations included in the payload will be fully processed for both addition and deletion.               -  When anchor managed location array present then it will add the anchor managed locations.
`,
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
		},
	}
}

func resourceWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocationsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vNetworkDeviceID := resourceItem["network_device_id"]

	vvNetworkDeviceID := vNetworkDeviceID.(string)
	request1 := expandRequestWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocationsAssignAnchorManagedApLocationsForWLC(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Wireless.AssignAnchorManagedApLocationsForWLC(vvNetworkDeviceID, request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AssignAnchorManagedAPLocationsForWLC", err))
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
			if restyResp3 == nil || strings.Contains(restyResp3.String(), "<!doctype html>") {
				errorMsg = response2.Response.Progress + "\nFailure Reason: " + response2.Response.FailureReason
			} else {
				errorMsg = restyResp3.String()
			}
			err1 := errors.New(errorMsg)
			diags = append(diags, diagError(
				"Failure when executing AssignAnchorManagedAPLocationsForWLC", err1))
			return diags
		}
	}

	vItem1 := flattenWirelessAssignAnchorManagedApLocationsForWLCItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting AssignAnchorManagedApLocationsForWLC response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocationsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestWirelessSettingsNetworkDeviceIDAssignAnchorManagedApLocationsAssignAnchorManagedApLocationsForWLC(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestWirelessAssignAnchorManagedApLocationsForWLC {
	request := catalystcentersdkgo.RequestWirelessAssignAnchorManagedApLocationsForWLC{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".anchor_managed_aplocations_site_ids")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".anchor_managed_aplocations_site_ids")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".anchor_managed_aplocations_site_ids")))) {
		request.AnchorManagedApLocationsSiteIDs = interfaceToSliceString(v)
	}
	return &request
}

func flattenWirelessAssignAnchorManagedApLocationsForWLCItem(item *catalystcentersdkgo.ResponseWirelessAssignAnchorManagedApLocationsForWLCResponse) []map[string]interface{} {
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
