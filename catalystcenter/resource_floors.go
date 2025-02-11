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

func resourceFloors() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Site Design.

- Create a floor in the network hierarchy under building.

- Updates a floor in the network hierarchy.

- Deletes a floor from the network hierarchy. This operations fails if there are any devices assigned to this floor.
`,

		CreateContext: resourceFloorsCreate,
		ReadContext:   resourceFloorsRead,
		UpdateContext: resourceFloorsUpdate,
		DeleteContext: resourceFloorsDelete,
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

						"floor_number": &schema.Schema{
							Description: `Floor number
`,
							Type:     schema.TypeInt,
							Computed: true,
						},
						"height": &schema.Schema{
							Description: `Floor height. Example : 10.1
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `Floor Id. Read only.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"length": &schema.Schema{
							Description: `Floor length. Example : 110.3
`,
							Type:     schema.TypeFloat,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Floor name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name_hierarchy": &schema.Schema{
							Description: `Floor hierarchical name. Read only.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"parent_id": &schema.Schema{
							Description: `Parent Id.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"rf_model": &schema.Schema{
							Description: `RF Model
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Example : floor
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"units_of_measure": &schema.Schema{
							Description: `Units Of Measure`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"width": &schema.Schema{
							Description: `Floor width. Example : 100.5
`,
							Type:     schema.TypeFloat,
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

						"floor_number": &schema.Schema{
							Description: `Floor number
`,
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"height": &schema.Schema{
							Description: `Floor height. Example : 10.1
`,
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"id": &schema.Schema{
							Description: `id path parameter. Floor Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"length": &schema.Schema{
							Description: `Floor length. Example : 110.3
`,
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Floor name
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parent_id": &schema.Schema{
							Description: `Parent Id`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"rf_model": &schema.Schema{
							Description: `RF Model
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"units_of_measure": &schema.Schema{
							Description: `Units Of Measure`,
							Type:        schema.TypeString,
							Optional:    true,
							Computed:    true,
						},
						"width": &schema.Schema{
							Description: `Floor width. Example : 100.5
`,
							Type:     schema.TypeFloat,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceFloorsCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestFloorsCreatesAFloorV2(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if vvID != "" {
		getResponse1, _, err := client.SiteDesign.GetsAFloorV2(vvID, nil)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceFloorsRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.SiteDesign.CreatesAFloorV2(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreatesAFloorV2", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreatesAFloorV2", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceFloorsRead(ctx, d, m)
}

func resourceFloorsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsAFloorV2")
		vvID := vID
		queryParams1 := catalystcentersdkgo.GetsAFloorV2QueryParams{}

		// has_unknown_response: None

		response1, restyResp1, err := client.SiteDesign.GetsAFloorV2(vvID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetsAFloorV2Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAFloorV2 response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceFloorsUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestFloorsUpdatesAFloorV2(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SiteDesign.UpdatesAFloorV2(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesAFloorV2", err, restyResp1.String(),
					"Failure at UpdatesAFloorV2, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesAFloorV2", err,
				"Failure at UpdatesAFloorV2, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesAFloorV2", err))
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
					"Failure when executing UpdatesAFloorV2", err1))
				return diags
			}
		}

	}

	return resourceFloorsRead(ctx, d, m)
}

func resourceFloorsDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]

	response1, restyResp1, err := client.SiteDesign.DeletesAFloorV2(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesAFloorV2", err, restyResp1.String(),
				"Failure at DeletesAFloorV2, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesAFloorV2", err,
			"Failure at DeletesAFloorV2, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletesAFloorV2", err))
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
				"Failure when executing DeletesAFloorV2", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestFloorsCreatesAFloorV2(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignCreatesAFloorV2 {
	request := catalystcentersdkgo.RequestSiteDesignCreatesAFloorV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_id")))) {
		request.ParentID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".floor_number")))) {
		request.FloorNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_model")))) {
		request.RfModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".width")))) {
		request.Width = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".length")))) {
		request.Length = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".height")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".height")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".height")))) {
		request.Height = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".units_of_measure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".units_of_measure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".units_of_measure")))) {
		request.UnitsOfMeasure = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestFloorsUpdatesAFloorV2(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignUpdatesAFloorV2 {
	request := catalystcentersdkgo.RequestSiteDesignUpdatesAFloorV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_id")))) {
		request.ParentID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".floor_number")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".floor_number")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".floor_number")))) {
		request.FloorNumber = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_model")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_model")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_model")))) {
		request.RfModel = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".width")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".width")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".width")))) {
		request.Width = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".length")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".length")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".length")))) {
		request.Length = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".height")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".height")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".height")))) {
		request.Height = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".units_of_measure")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".units_of_measure")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".units_of_measure")))) {
		request.UnitsOfMeasure = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
