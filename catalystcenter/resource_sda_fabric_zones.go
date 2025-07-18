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

func resourceSdaFabricZones() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on SDA.

- Updates a fabric zone based on user input.

- Adds a fabric zone based on user input.

- Deletes a fabric zone based on id.
`,

		CreateContext: resourceSdaFabricZonesCreate,
		ReadContext:   resourceSdaFabricZonesRead,
		UpdateContext: resourceSdaFabricZonesUpdate,
		DeleteContext: resourceSdaFabricZonesDelete,
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

						"authentication_profile_name": &schema.Schema{
							Description: `Authentication profile used for this fabric.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `ID of the fabric zone.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `ID of the network hierarchy.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"parameters": &schema.Schema{
				Description: `Array of RequestSdaAddFabricZone`,
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

									"authentication_profile_name": &schema.Schema{
										Description: `Authentication profile used for this fabric.
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"id": &schema.Schema{
										Description: `ID of the fabric zone (updating this field is not allowed).
`,
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"site_id": &schema.Schema{
										Description: `ID of the network hierarchy.
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

func resourceSdaFabricZonesCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters.0.payload"))
	request1 := expandRequestSdaFabricZonesAddFabricZone(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	// vID := resourceItem["id"]
	// vvID := interfaceToString(vID)
	vSiteID := resourceItem["site_id"]
	vvSiteID := interfaceToString(vSiteID)
	vName := resourceItem["authentication_profile_name"]
	vvName := interfaceToString(vName)
	queryParamImport := catalystcentersdkgo.GetFabricZonesQueryParams{}
	queryParamImport.SiteID = vvSiteID
	item2, err := searchSdaGetFabricZones(m, queryParamImport, vvName)
	if err == nil && item2 != nil {
		resourceMap := make(map[string]string)
		resourceMap["name"] = item2.AuthenticationProfileName
		resourceMap["site_id"] = item2.SiteID
		resourceMap["id"] = item2.ID
		d.SetId(joinResourceID(resourceMap))
		return resourceSdaFabricZonesRead(ctx, d, m)
	}
	resp1, restyResp1, err := client.Sda.AddFabricZone(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing AddFabricZone", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing AddFabricZone", err))
		return diags
	}
	if resp1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing AddFabricZone", err))
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
				"Failure when executing AddFabricZone", err1))
			return diags
		}
	}
	queryParamValidate := catalystcentersdkgo.GetFabricZonesQueryParams{}
	queryParamValidate.SiteID = vvSiteID
	item3, err := searchSdaGetFabricZones(m, queryParamValidate, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing AddFabricZone", err,
			"Failure at AddFabricZone, unexpected response", ""))
		return diags
	}

	resourceMap := make(map[string]string)
	resourceMap["name"] = item3.AuthenticationProfileName
	resourceMap["site_id"] = item3.SiteID
	resourceMap["id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceSdaFabricZonesRead(ctx, d, m)
}

func resourceSdaFabricZonesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vSiteID := resourceMap["site_id"]
	vName := resourceMap["name"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetFabricZones")
		queryParams1 := catalystcentersdkgo.GetFabricZonesQueryParams{}
		queryParams1.SiteID = vSiteID

		item1, err := searchSdaGetFabricZones(m, queryParams1, vName)
		if err != nil || item1 == nil {
			d.SetId("")
			return diags
		}
		// Review flatten function used
		items := []catalystcentersdkgo.ResponseSdaGetFabricZonesResponse{
			*item1,
		}
		vItem1 := flattenSdaGetFabricZonesItems(&items)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetFabricZones search response",
				err))
			return diags
		}

	}
	return diags
}

func resourceSdaFabricZonesUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]
	if d.HasChange("parameters") {
		request1 := expandRequestSdaFabricZonesUpdateFabricZone(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		if request1 != nil && len(*request1) > 0 {
			req := *request1
			req[0].ID = vID
			request1 = &req
		}
		response1, restyResp1, err := client.Sda.UpdateFabricZone(request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateFabricZone", err, restyResp1.String(),
					"Failure at UpdateFabricZone, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateFabricZone", err,
				"Failure at UpdateFabricZone, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdateFabricZone", err))
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
					"Failure when executing UpdateFabricZone", err1))
				return diags
			}
		}

	}

	return resourceSdaFabricZonesRead(ctx, d, m)
}

func resourceSdaFabricZonesDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.Sda.DeleteFabricZoneByID(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteFabricZoneByID", err, restyResp1.String(),
				"Failure at DeleteFabricZoneByID, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteFabricZoneByID", err,
			"Failure at DeleteFabricZoneByID, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeleteFabricZoneByID", err))
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
				"Failure when executing DeleteFabricZoneByID", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}

func expandRequestSdaFabricZonesAddFabricZone(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaAddFabricZone {
	request := catalystcentersdkgo.RequestSdaAddFabricZone{}
	if v := expandRequestSdaFabricZonesAddFabricZoneItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricZonesAddFabricZoneItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSdaAddFabricZone {
	request := []catalystcentersdkgo.RequestItemSdaAddFabricZone{}
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
		i := expandRequestSdaFabricZonesAddFabricZoneItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricZonesAddFabricZoneItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSdaAddFabricZone {
	request := catalystcentersdkgo.RequestItemSdaAddFabricZone{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_profile_name")))) {
		request.AuthenticationProfileName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricZonesUpdateFabricZone(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSdaUpdateFabricZone {
	request := catalystcentersdkgo.RequestSdaUpdateFabricZone{}
	if v := expandRequestSdaFabricZonesUpdateFabricZoneItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricZonesUpdateFabricZoneItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSdaUpdateFabricZone {
	request := []catalystcentersdkgo.RequestItemSdaUpdateFabricZone{}
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
		i := expandRequestSdaFabricZonesUpdateFabricZoneItem(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestSdaFabricZonesUpdateFabricZoneItem(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSdaUpdateFabricZone {
	request := catalystcentersdkgo.RequestItemSdaUpdateFabricZone{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".authentication_profile_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".authentication_profile_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".authentication_profile_name")))) {
		request.AuthenticationProfileName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchSdaGetFabricZones(m interface{}, queryParams catalystcentersdkgo.GetFabricZonesQueryParams, vName string) (*catalystcentersdkgo.ResponseSdaGetFabricZonesResponse, error) {
	client := m.(*catalystcentersdkgo.Client)
	var err error
	var foundItem *catalystcentersdkgo.ResponseSdaGetFabricZonesResponse
	var ite *catalystcentersdkgo.ResponseSdaGetFabricZones

	ite, _, err = client.Sda.GetFabricZones(&queryParams)
	if err != nil || ite == nil {
		return foundItem, err
	}
	itemsCopy := *ite.Response
	if itemsCopy == nil {
		return foundItem, err
	}
	for _, item := range itemsCopy {
		if item.AuthenticationProfileName == vName {
			foundItem = &item
			return foundItem, err
		}
	}
	return foundItem, err

}
