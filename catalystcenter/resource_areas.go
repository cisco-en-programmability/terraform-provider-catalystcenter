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

func resourceAreas() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on Site Design.

- Creates an area in the network hierarchy.

- Updates an area in the network hierarchy.

- Deletes an area in the network hierarchy. This operations fails if there are any child areas or buildings for this
area.
`,

		CreateContext: resourceAreasCreate,
		ReadContext:   resourceAreasRead,
		UpdateContext: resourceAreasUpdate,
		DeleteContext: resourceAreasDelete,
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

						"id": &schema.Schema{
							Description: `Aread Id. Read only.
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Description: `Area name
`,
							Type:     schema.TypeString,
							Computed: true,
						},
						"name_hierarchy": &schema.Schema{
							Description: `Area hierarchical name. Read only.
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
							Description: `Site Type.
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

						"id": &schema.Schema{
							Description: `id path parameter. Area Id
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Description: `Area name
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

func resourceAreasCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestAreasCreatesAnAreaV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vID := resourceItem["id"]
	vvID := interfaceToString(vID)
	if vvID != "" {
		getResponse1, _, err := client.SiteDesign.GetsAnAreaV1(vvID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["id"] = vvID
			d.SetId(joinResourceID(resourceMap))
			return resourceAreasRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.SiteDesign.CreatesAnAreaV1(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreatesAnAreaV1", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreatesAnAreaV1", err))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["id"] = vvID
	d.SetId(joinResourceID(resourceMap))
	return resourceAreasRead(ctx, d, m)
}

func resourceAreasRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vID := resourceMap["id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetsAnAreaV1")
		vvID := vID

		// has_unknown_response: None

		response1, restyResp1, err := client.SiteDesign.GetsAnAreaV1(vvID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenSiteDesignGetsAnAreaV1Item(response1.Response)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetsAnAreaV1 response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceAreasUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vvID := resourceMap["id"]

	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestAreasUpdatesAnAreaV1(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.SiteDesign.UpdatesAnAreaV1(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdatesAnAreaV1", err, restyResp1.String(),
					"Failure at UpdatesAnAreaV1, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdatesAnAreaV1", err,
				"Failure at UpdatesAnAreaV1, unexpected response", ""))
			return diags
		}

		if response1.Response == nil {
			diags = append(diags, diagError(
				"Failure when executing UpdatesAnAreaV1", err))
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
					"Failure when executing UpdatesAnAreaV1", err1))
				return diags
			}
		}

	}

	return resourceAreasRead(ctx, d, m)
}

func resourceAreasDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["id"]
	response1, restyResp1, err := client.SiteDesign.DeletesAnAreaV1(vvID)
	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeletesAnAreaV1", err, restyResp1.String(),
				"Failure at DeletesAnAreaV1, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeletesAnAreaV1", err,
			"Failure at DeletesAnAreaV1, unexpected response", ""))
		return diags
	}

	if response1.Response == nil {
		diags = append(diags, diagError(
			"Failure when executing DeletesAnAreaV1", err))
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
				"Failure when executing DeletesAnAreaV1", err1))
			return diags
		}
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestAreasCreatesAnAreaV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignCreatesAnAreaV1 {
	request := catalystcentersdkgo.RequestSiteDesignCreatesAnAreaV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_id")))) {
		request.ParentID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestAreasUpdatesAnAreaV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestSiteDesignUpdatesAnAreaV1 {
	request := catalystcentersdkgo.RequestSiteDesignUpdatesAnAreaV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".parent_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".parent_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".parent_id")))) {
		request.ParentID = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}
