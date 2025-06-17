package catalystcenter

import (
	"context"
	"strings"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceTagsNetworkDevicesMembersAssociationsBulk() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Tag.

- Updates the tags associated with the devices. A tag is a user-defined or system-defined construct to group resources.
When a device is tagged, it is called a member of the tag. A tag can be created by using this POST
**/dna/intent/api/v1/tag** API.
`,

		CreateContext: resourceTagsNetworkDevicesMembersAssociationsBulkCreate,
		ReadContext:   resourceTagsNetworkDevicesMembersAssociationsBulkRead,
		DeleteContext: resourceTagsNetworkDevicesMembersAssociationsBulkDelete,
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
							Description: `The UUID of the task
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `The path to the API endpoint to GET for information on the task
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
						"payload": &schema.Schema{
							Description: `Array of RequestTagUpdateTagsAssociatedWithTheNetworkDevices`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Network device id
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"tags": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"id": &schema.Schema{
													Description: `Tag id
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
							},
						},
					},
				},
			},
		},
	}
}

func resourceTagsNetworkDevicesMembersAssociationsBulkCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevices(ctx, "parameters.0", d)

	response1, restyResp1, err := client.Tag.UpdateTagsAssociatedWithTheNetworkDevices(request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		d.SetId("")
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing UpdateTagsAssociatedWithTheNetworkDevices", err))
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
				"Failure when executing UpdateTagsAssociatedWithTheNetworkDevices", err1))
			return diags
		}
	}

	vItem1 := flattenTagUpdateTagsAssociatedWithTheNetworkDevicesItem(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateTagsAssociatedWithTheNetworkDevices response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags
}
func resourceTagsNetworkDevicesMembersAssociationsBulkRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceTagsNetworkDevicesMembersAssociationsBulkDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevices(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestTagUpdateTagsAssociatedWithTheNetworkDevices {
	request := catalystcentersdkgo.RequestTagUpdateTagsAssociatedWithTheNetworkDevices{}
	if v := expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevicesItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevicesItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemTagUpdateTagsAssociatedWithTheNetworkDevices {
	request := []catalystcentersdkgo.RequestItemTagUpdateTagsAssociatedWithTheNetworkDevices{}
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
		i := expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevicesItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevicesItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemTagUpdateTagsAssociatedWithTheNetworkDevices {
	request := catalystcentersdkgo.RequestItemTagUpdateTagsAssociatedWithTheNetworkDevices{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".tags")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".tags")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".tags")))) {
		request.Tags = expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevicesItemTagsArray(ctx, key+".tags", d)
	}
	return &request
}

func expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevicesItemTagsArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemTagUpdateTagsAssociatedWithTheNetworkDevicesTags {
	request := []catalystcentersdkgo.RequestItemTagUpdateTagsAssociatedWithTheNetworkDevicesTags{}
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
		i := expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevicesItemTags(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestTagsNetworkDevicesMembersAssociationsBulkUpdateTagsAssociatedWithTheNetworkDevicesItemTags(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemTagUpdateTagsAssociatedWithTheNetworkDevicesTags {
	request := catalystcentersdkgo.RequestItemTagUpdateTagsAssociatedWithTheNetworkDevicesTags{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	return &request
}

func flattenTagUpdateTagsAssociatedWithTheNetworkDevicesItem(item *catalystcentersdkgo.ResponseTagUpdateTagsAssociatedWithTheNetworkDevicesResponse) []map[string]interface{} {
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
