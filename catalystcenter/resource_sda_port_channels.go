package catalystcenter

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"time"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSdaPortChannels() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Adds port channels based on user input.

- Updates port channels based on user input.

- Deletes port channels based on user input.

- Deletes a port channel based on id.
`,

		CreateContext: resourceSdaPortChannelsCreate,
		ReadContext:   resourceSdaPortChannelsRead,
		UpdateContext: resourceSdaPortChannelsUpdate,
		DeleteContext: resourceSdaPortChannelsDelete,
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

						"connected_device_type": &schema.Schema{
							Description: `Connected device type of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Description of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"fabric_id": &schema.Schema{
							Description: `ID of the fabric the device is assigned to.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"interface_names": &schema.Schema{
							Description: `Interface names of this port channel.
`,
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"network_device_id": &schema.Schema{
							Description: `ID of the network device.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"port_channel_name": &schema.Schema{
							Description: `Name of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Description: `Protocol of the port channel.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddPortChannels`,
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

									"connected_device_type": &schema.Schema{
										Description: `Connected device type of the port channel.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"description": &schema.Schema{
										Description: `Description of the port channel.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"fabric_id": &schema.Schema{
										Description: `ID of the fabric the device is assigned to.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the port channel (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"interface_names": &schema.Schema{
										Description: `Interface names for this port channel (Maximum 16 ports for LACP protocol, Maximum 8 ports for PAGP and ON protocol).
`,
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"network_device_id": &schema.Schema{
										Description: `ID of the network device.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"port_channel_name": &schema.Schema{
										Description: `Name of the port channel (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"protocol": &schema.Schema{
										Description: `Protocol of the port channel (only PAGP is allowed if connectedDeviceType is EXTENDED_NODE).
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

func resourceSdaPortChannelsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaPortChannelsAddPortChannelsV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	vName := resourceItem["port_channel_name"]
	vvName := interfaceToString(vName)
	queryParamImport := catalystcentersdkgo.GetPortChannelsV1QueryParams{}
	queryParamImport.PortChannelName = vvName
	item2, err := searchSdaGetPortChannels(m, queryParamImport, vvID)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["id"] = item2.ID
		resourceMap["port_channel_name"] = item2.PortChannelName
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaPortChannelsRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddPortChannels(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddPortChannels", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddPortChannels", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddPortChannels", err))
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
				"Failure when executing AddPortChannels", err1))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetPortChannelsV1QueryParams{}
	queryParamValidate.PortChannelName = vvName
	item3, err := searchSdaGetPortChannels(m, queryParamValidate, vvID)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddPortChannels", err,
			"Failure at AddPortChannels, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["id"] = item3.ID
	resourceMap["port_channel_name"] = item3.PortChannelName
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaPortChannelsRead(ctx, d, m)
}

func resourceSdaPortChannelsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	vvName := resourceMap["port_channel_name"]
	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPortChannels")
		queryParams1 := catalystcentersdkgo.GetPortChannelsV1QueryParams{}
		queryParams1.PortChannelName = vvName
		item1, err := searchSdaGetPortChannels(m, queryParams1, vvID)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		items := []catalystcentersdkgo.ResponseSdaGetPortChannelsV1Response{
			*item1,
		}
		// Review flatten function used
		vItem1 := flattenSdaGetPortChannelsV1Items(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPortChannels search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaPortChannelsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaPortChannelsUpdatePortChannelsV1(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdatePortChannels(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatePortChannels", err, restyResp1.String(),
					"Failure at UpdatePortChannels, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatePortChannels", err,
				"Failure at UpdatePortChannels, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatePortChannels", err))
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
					"Failure when executing UpdatePortChannels", err1))
				return diags
			}
		}

	}

	return resourceSdaPortChannelsRead(ctx, d, m)
}

func resourceSdaPortChannelsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]

	response1, restyResp1, err := client.Sda.DeletePortChannelByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletePortChannelByID", err, restyResp1.String(),
				"Failure at DeletePortChannelByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletePortChannelByID", err,
			"Failure at DeletePortChannelByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletePortChannelByID", err))
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
				"Failure when executing DeletePortChannelByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestSdaPortChannelsAddPortChannelsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaAddPortChannelsV1 {
	request := catalystcentersdkgo.RequestSdaAddPortChannelsV1{}
	if v := expandRequestSdaPortChannelsAddPortChannelsV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortChannelsAddPortChannelsV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSdaAddPortChannelsV1 {
	request := []catalystcentersdkgo.RequestItemSdaAddPortChannelsV1{}
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
		i := expandRequestSdaPortChannelsAddPortChannelsV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortChannelsAddPortChannelsV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSdaAddPortChannelsV1 {
	request := catalystcentersdkgo.RequestItemSdaAddPortChannelsV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_names")))) {
		request.InterfaceNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_device_type")))) {
		request.ConnectedDeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortChannelsUpdatePortChannelsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaUpdatePortChannelsV1 {
	request := catalystcentersdkgo.RequestSdaUpdatePortChannelsV1{}
	if v := expandRequestSdaPortChannelsUpdatePortChannelsV1ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortChannelsUpdatePortChannelsV1ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSdaUpdatePortChannelsV1 {
	request := []catalystcentersdkgo.RequestItemSdaUpdatePortChannelsV1{}
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
		i := expandRequestSdaPortChannelsUpdatePortChannelsV1Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaPortChannelsUpdatePortChannelsV1Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSdaUpdatePortChannelsV1 {
	request := catalystcentersdkgo.RequestItemSdaUpdatePortChannelsV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".fabric_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".fabric_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".fabric_id")))) {
		request.FabricID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".network_device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".network_device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".network_device_id")))) {
		request.NetworkDeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".port_channel_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".port_channel_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".port_channel_name")))) {
		request.PortChannelName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".interface_names")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".interface_names")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".interface_names")))) {
		request.InterfaceNames = interfaceToSliceString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connected_device_type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connected_device_type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connected_device_type")))) {
		request.ConnectedDeviceType = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".protocol")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".protocol")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".protocol")))) {
		request.Protocol = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetPortChannels(m interface{}, queryParams catalystcentersdkgo.GetPortChannelsV1QueryParams, vID string) (*catalystcentersdkgo.ResponseSdaGetPortChannelsV1Response, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseSdaGetPortChannelsV1Response
	var ite *catalystcentersdkgo.ResponseSdaGetPortChannelsV1
	if vID != "" {
		queryParams.Offset = 1
		nResponse, _, err := client.Sda.GetPortChannels(nil)
		maxPageSize := len(*nResponse.Response)
		for len(*nResponse.Response) > 0 {
			time.Sleep(15 * time.Second)
			for _, item := range *nResponse.Response {
				if vID == item.ID {
					foundItem = &item
					return foundItem, err
				}
			}
			queryParams.Limit = float64(maxPageSize)
			queryParams.Offset += float64(maxPageSize)
			nResponse, _, err = client.Sda.GetPortChannels(&queryParams)
			if nResponse == nil || nResponse.Response == nil {
				break
			}
		}
		return nil, err
	} else if queryParams.PortChannelName != "" {
		ite, _, err = client.Sda.GetPortChannels(&queryParams)
		if err != nil || ite == nil {
			return foundItem, err
		}
		itemsCopy := *ite.Response
		if itemsCopy == nil {
			return foundItem, err
		}
		for _, item := range itemsCopy {
			if item.PortChannelName == queryParams.PortChannelName {
				foundItem = &item
				return foundItem, err
			}
		}
		return foundItem, err
	}
	return foundItem, err
}
