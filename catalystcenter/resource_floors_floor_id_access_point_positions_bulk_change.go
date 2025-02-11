package catalystcenter

import (
	"context"

	"errors"

	"time"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceFloorsFloorIDAccessPointPositionsBulkChange() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Site Design.

- Position or reposition the Access Points on the map.
`,

		CreateContext: resourceFloorsFloorIDAccessPointPositionsBulkChangeCreate,
		ReadContext:   resourceFloorsFloorIDAccessPointPositionsBulkChangeRead,
		DeleteContext: resourceFloorsFloorIDAccessPointPositionsBulkChangeDelete,
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
							Description: `Task Id in uuid format. e.g. : 3200a44a-9186-4caf-8c32-419cd1f3d3f5
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Description: `URL to get task details e.g. : /api/v1/task/3200a44a-9186-4caf-8c32-419cd1f3d3f5
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
						"floor_id": &schema.Schema{
							Description: `floorId path parameter. Floor Id
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"payload": &schema.Schema{
							Description: `Array of RequestSiteDesignEditTheAccessPointsPositionsV2`,
							Type:        schema.TypeList,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"id": &schema.Schema{
										Description: `Access Point Id
`,
										Type:     schema.TypeString,
										Optional: true,
										ForceNew: true,
										Computed: true,
									},
									"position": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"x": &schema.Schema{
													Description: `Access Point X coordinate in feet
`,
													Type:     schema.TypeFloat,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"y": &schema.Schema{
													Description: `Access Point Y coordinate in feet
`,
													Type:     schema.TypeFloat,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
												"z": &schema.Schema{
													Description: `Access Point Z coordinate in feet
`,
													Type:     schema.TypeFloat,
													Optional: true,
													ForceNew: true,
													Computed: true,
												},
											},
										},
									},
									"radios": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"antenna": &schema.Schema{
													Type:     schema.TypeList,
													Optional: true,
													ForceNew: true,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{

															"azimuth": &schema.Schema{
																Description: `Angle of the antenna, measured relative to the x axis, clockwise. The azimuth range is from 0 through 360
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"elevation": &schema.Schema{
																Description: `Elevation of the antenna. The elevation range is from -90 through 90
`,
																Type:     schema.TypeInt,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
															"name": &schema.Schema{
																Description: `Antenna type for this Access Point. Use /dna/intent/api/v1/maps/supported-access-points to find supported Antennas for a particualr Access Point model
`,
																Type:     schema.TypeString,
																Optional: true,
																ForceNew: true,
																Computed: true,
															},
														},
													},
												},
												"id": &schema.Schema{
													Description: `Radio Id
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

func resourceFloorsFloorIDAccessPointPositionsBulkChangeCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vFloorID := resourceItem["floor_id"]

	vvFloorID := vFloorID.(string)
	request1 := expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.SiteDesign.EditTheAccessPointsPositionsV2(vvFloorID, request1)

	if request1 != nil {
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
	}

	vItem1 := flattenSiteDesignEditTheAccessPointsPositionsV2Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting EditTheAccessPointsPositionsV2 response",
			err))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing EditTheAccessPointsPositionsV2", err))
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
				"Failure when executing EditTheAccessPointsPositionsV2", err1))
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
func resourceFloorsFloorIDAccessPointPositionsBulkChangeRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceFloorsFloorIDAccessPointPositionsBulkChangeDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*dnacentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignEditTheAccessPointsPositionsV2 {
	request := catalystcentersdkgo.RequestSiteDesignEditTheAccessPointsPositionsV2{}
	if v := expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemArray(ctx, key+".payload", d); v != nil {
		request = *v
	}
	return &request
}

func expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2 {
	request := []catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2{}
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
		i := expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2Item(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2Item(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2 {
	request := catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".position")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".position")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".position")))) {
		request.Position = expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemPosition(ctx, key+".position.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".radios")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".radios")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".radios")))) {
		request.Radios = expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemRadiosArray(ctx, key+".radios", d)
	}
	return &request
}

func expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemPosition(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2Position {
	request := catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2Position{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".x")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".x")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".x")))) {
		request.X = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".y")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".y")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".y")))) {
		request.Y = interfaceToFloat64Ptr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".z")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".z")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".z")))) {
		request.Z = interfaceToFloat64Ptr(v)
	}
	return &request
}

func expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemRadiosArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2Radios {
	request := []catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2Radios{}
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
		i := expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemRadios(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemRadios(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2Radios {
	request := catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2Radios{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".id")))) {
		request.ID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".antenna")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".antenna")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".antenna")))) {
		request.Antenna = expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemRadiosAntenna(ctx, key+".antenna.0", d)
	}
	return &request
}

func expandRequestFloorsFloorIDAccessPointPositionsBulkChangeEditTheAccessPointsPositionsV2ItemRadiosAntenna(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2RadiosAntenna {
	request := catalystcentersdkgo.RequestItemSiteDesignEditTheAccessPointsPositionsV2RadiosAntenna{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".azimuth")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".azimuth")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".azimuth")))) {
		request.Azimuth = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".elevation")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".elevation")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".elevation")))) {
		request.Elevation = interfaceToIntPtr(v)
	}
	return &request
}

func flattenSiteDesignEditTheAccessPointsPositionsV2Item(item *catalystcentersdkgo.ResponseSiteDesignEditTheAccessPointsPositionsV2Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["url"] = item.URL
	respItem["task_id"] = item.TaskID
	return []map[string]interface{}{
		respItem,
	}
}
