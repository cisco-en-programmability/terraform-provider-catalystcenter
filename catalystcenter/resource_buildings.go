package catalystcenter

import (
	"context"
	"errors"
	"reflect"
	"time"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBuildings() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Site Design.

- Creates a building in the network hierarchy under area.

- Updates a building in the network hierarchy.

- Deletes building in the network hierarchy. This operations fails if there are any floors for this building, or if
there are any devices assigned to this building.
`,

		CreateContext: resourceBuildingsCreate,
		ReadContext:   resourceBuildingsRead,
		UpdateContext: resourceBuildingsUpdate,
		DeleteContext: resourceBuildingsDelete,
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

						"address": &schema.Schema{
							Description: `Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"country": &schema.Schema{
							Description: `Country name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"latitude": &schema.Schema{
							Description: `Building Latitude. Example: 37.403712
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"longitude": &schema.Schema{
							Description: `Building Longitude. Example: -121.971063
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Building name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_id": &schema.Schema{
							Description: `Parent Id
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Example: building
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

						"address": &schema.Schema{
							Description: `Building address. Example: 4900 Marie P. Debartolo Way, Santa Clara, California 95054, United States. Please note that if only the address is provided when creating a building, the UI will not display the geo-location on the map. To ensure the location is rendered, you must also provide the latitude and longitude. If a building has been created without these coordinates and you wish to display its geo-location on the map later, you can edit the building details via the UI to include the latitude and longitude. This limitation will be resolved in a future release.
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"country": &schema.Schema{
							Description: `Country name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. Building Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"latitude": &schema.Schema{
							Description: `Building Latitude. Example: 37.403712
`,
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"longitude": &schema.Schema{
							Description: `Building Longitude. Example: -121.971063
`,
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Building name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parent_id": &schema.Schema{
							Description: `Parent Id
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

func resourceBuildingsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestBuildingsCreatesABuildingV2(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if vvID != "" {
		getResponse1, _, err := client.SiteDesign.GetsABuildingV2(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceBuildingsRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.SiteDesign.CreatesABuildingV2(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreatesABuildingV2", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreatesABuildingV2", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = resp1.Response.TaskID
	d.SetId(joinResourceID(resourceMap))
	return resourceBuildingsRead(ctx, d, m)
}

func resourceBuildingsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsABuildingV2")
		vvID := vID

		// has_unknown_response: None

		response1, restyResp1, err := client.SiteDesign.GetsABuildingV2(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetsABuildingV2Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsABuildingV2 response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceBuildingsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestBuildingsUpdatesABuildingV2(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SiteDesign.UpdatesABuildingV2(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesABuildingV2", err, restyResp1.String(),
					"Failure at UpdatesABuildingV2, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesABuildingV2", err,
				"Failure at UpdatesABuildingV2, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesABuildingV2", err))
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
					"Failure when executing UpdatesABuildingV2", err1))
				return diags
			}
		}

	}

	return resourceBuildingsRead(ctx, d, m)
}

func resourceBuildingsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]

	response1, restyResp1, err := client.SiteDesign.DeletesABuildingV2(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesABuildingV2", err, restyResp1.String(),
				"Failure at DeletesABuildingV2, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesABuildingV2", err,
			"Failure at DeletesABuildingV2, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletesABuildingV2", err))
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
				"Failure when executing DeletesABuildingV2", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestBuildingsCreatesABuildingV2(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignCreatesABuildingV2 {
	request := catalystcentersdkgo.RequestSiteDesignCreatesABuildingV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_id")))) {
		request.ParentID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latitude")))) {
		request.Latitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".longitude")))) {
		request.Longitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".country")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".country")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".country")))) {
		request.Country = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestBuildingsUpdatesABuildingV2(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignUpdatesABuildingV2 {
	request := catalystcentersdkgo.RequestSiteDesignUpdatesABuildingV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_id")))) {
		request.ParentID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".latitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".latitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".latitude")))) {
		request.Latitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".longitude")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".longitude")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".longitude")))) {
		request.Longitude = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".address")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".address")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".address")))) {
		request.Address = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".country")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".country")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".country")))) {
		request.Country = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
