package catalystcenter

import (
	"context"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIntegrationSettingsInstancesItsm() *schema.Resource {
	return &schema.Resource{
		Description: `It manages create, read, update and delete operations on ITSM Integration.

- Creates ITSM Integration setting

- Updates the ITSM Integration setting

- Deletes the ITSM Integration setting
`,

		CreateContext: resourceIntegrationSettingsInstancesItsmCreate,
		ReadContext:   resourceIntegrationSettingsInstancesItsmRead,
		UpdateContext: resourceIntegrationSettingsInstancesItsmUpdate,
		DeleteContext: resourceIntegrationSettingsInstancesItsmDelete,
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
							Description: `Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"data": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connection_settings": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"auth_password": &schema.Schema{
													Description: `Auth Password`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"auth_user_name": &schema.Schema{
													Description: `Auth User Name`,
													Type:        schema.TypeString,
													Computed:    true,
												},
												"url": &schema.Schema{
													Description: `Url`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `Description`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"dyp_id": &schema.Schema{
							Description: `Dyp Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"dyp_major_version": &schema.Schema{
							Description: `Dyp Major Version`,
							Type:        schema.TypeInt,
							Computed:    true,
						},
						"dyp_name": &schema.Schema{
							Description: `Dyp Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"name": &schema.Schema{
							Description: `Name`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"tenant_id": &schema.Schema{
							Description: `Tenant Id`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"unique_key": &schema.Schema{
							Description: `Unique Key`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"updated_by": &schema.Schema{
							Description: `Updated By`,
							Type:        schema.TypeString,
							Computed:    true,
						},
						"updated_date": &schema.Schema{
							Description: `Updated Date`,
							Type:        schema.TypeInt,
							Computed:    true,
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

						"data": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"connection_settings": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"auth_password": &schema.Schema{
													Description: `Auth Password`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"auth_user_name": &schema.Schema{
													Description: `Auth User Name`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
												"url": &schema.Schema{
													Description: `Url`,
													Type:        schema.TypeString,
													Optional:    true,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"description": &schema.Schema{
							Description: `Description of the setting instance
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"dyp_name": &schema.Schema{
							Description: `It can be ServiceNowConnection
`,
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"instance_id": &schema.Schema{
							Description: `instanceId path parameter. Instance Id of the Integration setting instance
`,
							Type:     schema.TypeString,
							Required: true,
						},
						"name": &schema.Schema{
							Description: `Name of the setting instance
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

func resourceIntegrationSettingsInstancesItsmCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))
	request1 := expandRequestIntegrationSettingsInstancesItsmCreateItsmIntegrationSettingV1(ctx, "parameters.0", d)
	log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))

	vInstanceID, okInstanceID := resourceItem["instance_id"]
	var vvInstanceID string
	if okInstanceID {
		vvInstanceID = interfaceToString(vInstanceID)
	} else {
		vvInstanceID = ""
	}
	vName := resourceItem["name"]
	vvName := interfaceToString(vName)
	if okInstanceID && vvInstanceID != "" {
		getResponse1, _, err := client.ItsmIntegration.GetItsmIntegrationSettingByID(vvInstanceID)
		if err == nil && getResponse1 != nil {
			resourceMap := make(map[string]string)
			resourceMap["instance_id"] = vvInstanceID
			d.SetId(joinResourceID(resourceMap))
			return resourceIntegrationSettingsInstancesItsmRead(ctx, d, m)
		}
	}
	resp1, restyResp1, err := client.ItsmIntegration.CreateItsmIntegrationSetting(request1)
	if err != nil || resp1 == nil {
		if restyResp1 != nil {
			diags = append(diags, diagErrorWithResponse(
				"Failure when executing CreateItsmIntegrationSetting", err, restyResp1.String()))
			return diags
		}
		diags = append(diags, diagError(
			"Failure when executing CreateItsmIntegrationSetting", err))
		return diags
	}

	item3, err := searchITSM(m, vvName)
	if err != nil || item3 == nil {
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing CreateItsmIntegrationSetting", err,
			"Failure at CreateItsmIntegrationSetting, unexpected response", ""))
		return diags
	}
	resourceMap := make(map[string]string)
	resourceMap["instance_id"] = item3.ID
	d.SetId(joinResourceID(resourceMap))
	return resourceIntegrationSettingsInstancesItsmRead(ctx, d, m)
}

func resourceIntegrationSettingsInstancesItsmRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)

	vInstanceID := resourceMap["instance_id"]

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetItsmIntegrationSettingByID")
		vvInstanceID := vInstanceID

		response1, restyResp1, err := client.ItsmIntegration.GetItsmIntegrationSettingByID(vvInstanceID)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			d.SetId("")
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItem1 := flattenItsmIntegrationGetItsmIntegrationSettingByIDV1Item(response1)
		if err := d.Set("item", vItem1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetItsmIntegrationSettingByID response",
				err))
			return diags
		}

		return diags

	}
	return diags
}

func resourceIntegrationSettingsInstancesItsmUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["instance_id"]
	if d.HasChange("parameters") {
		log.Printf("[DEBUG] ID used for update operation %s", vvID)
		request1 := expandRequestIntegrationSettingsInstancesItsmUpdateItsmIntegrationSettingV1(ctx, "parameters.0", d)
		log.Printf("[DEBUG] request sent => %v", responseInterfaceToString(*request1))
		response1, restyResp1, err := client.ItsmIntegration.UpdateItsmIntegrationSetting(vvID, request1)
		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] resty response for update operation => %v", restyResp1.String())
				diags = append(diags, diagErrorWithAltAndResponse(
					"Failure when executing UpdateItsmIntegrationSetting", err, restyResp1.String(),
					"Failure at UpdateItsmIntegrationSetting, unexpected response", ""))
				return diags
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing UpdateItsmIntegrationSetting", err,
				"Failure at UpdateItsmIntegrationSetting, unexpected response", ""))
			return diags
		}

		//TODO REVIEW

	}

	return resourceIntegrationSettingsInstancesItsmRead(ctx, d, m)
}

func resourceIntegrationSettingsInstancesItsmDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics

	resourceID := d.Id()
	resourceMap := separateResourceID(resourceID)
	vvID := resourceMap["instance_id"]

	restyResp1, err := client.ItsmIntegration.DeleteItsmIntegrationSetting(vvID)
	if err != nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] resty response for delete operation => %v", restyResp1.String())
			diags = append(diags, diagErrorWithAltAndResponse(
				"Failure when executing DeleteItsmIntegrationSetting", err, restyResp1.String(),
				"Failure at DeleteItsmIntegrationSetting, unexpected response", ""))
			return diags
		}
		diags = append(diags, diagErrorWithAlt(
			"Failure when executing DeleteItsmIntegrationSetting", err,
			"Failure at DeleteItsmIntegrationSetting, unexpected response", ""))
		return diags
	}

	//TODO REVIEW

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
func expandRequestIntegrationSettingsInstancesItsmCreateItsmIntegrationSettingV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItsmIntegrationCreateItsmIntegrationSettingV1 {
	request := catalystcentersdkgo.RequestItsmIntegrationCreateItsmIntegrationSettingV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestIntegrationSettingsInstancesItsmCreateItsmIntegrationSettingV1Data(ctx, key+".data.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dyp_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dyp_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dyp_name")))) {
		request.DypName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIntegrationSettingsInstancesItsmCreateItsmIntegrationSettingV1Data(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItsmIntegrationCreateItsmIntegrationSettingV1Data {
	request := catalystcentersdkgo.RequestItsmIntegrationCreateItsmIntegrationSettingV1Data{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_settings")))) {
		request.ConnectionSettings = expandRequestIntegrationSettingsInstancesItsmCreateItsmIntegrationSettingV1DataConnectionSettings(ctx, key+".connection_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIntegrationSettingsInstancesItsmCreateItsmIntegrationSettingV1DataConnectionSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItsmIntegrationCreateItsmIntegrationSettingV1DataConnectionSettings {
	request := catalystcentersdkgo.RequestItsmIntegrationCreateItsmIntegrationSettingV1DataConnectionSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_user_name")))) {
		request.AuthUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_password")))) {
		request.AuthPassword = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIntegrationSettingsInstancesItsmUpdateItsmIntegrationSettingV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItsmIntegrationUpdateItsmIntegrationSettingV1 {
	request := catalystcentersdkgo.RequestItsmIntegrationUpdateItsmIntegrationSettingV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".name")))) {
		request.Name = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".data")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".data")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".data")))) {
		request.Data = expandRequestIntegrationSettingsInstancesItsmUpdateItsmIntegrationSettingV1Data(ctx, key+".data.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".dyp_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".dyp_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".dyp_name")))) {
		request.DypName = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIntegrationSettingsInstancesItsmUpdateItsmIntegrationSettingV1Data(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItsmIntegrationUpdateItsmIntegrationSettingV1Data {
	request := catalystcentersdkgo.RequestItsmIntegrationUpdateItsmIntegrationSettingV1Data{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".connection_settings")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".connection_settings")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".connection_settings")))) {
		request.ConnectionSettings = expandRequestIntegrationSettingsInstancesItsmUpdateItsmIntegrationSettingV1DataConnectionSettings(ctx, key+".connection_settings.0", d)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func expandRequestIntegrationSettingsInstancesItsmUpdateItsmIntegrationSettingV1DataConnectionSettings(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestItsmIntegrationUpdateItsmIntegrationSettingV1DataConnectionSettings {
	request := catalystcentersdkgo.RequestItsmIntegrationUpdateItsmIntegrationSettingV1DataConnectionSettings{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".url")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".url")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".url")))) {
		request.URL = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_user_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_user_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_user_name")))) {
		request.AuthUserName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".auth_password")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".auth_password")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".auth_password")))) {
		request.AuthPassword = interfaceToString(v)
	}
	if isEmptyValue(reflect.ValueOf(request)) {
		return nil
	}
	return &request
}

func searchITSM(m interface{}, vName string) (foundItem *catalystcentersdkgo.ResponseItemItsmIntegrationGetAllItsmIntegrationSettingsV1, err error) {
	client := m.(*catalystcentersdkgo.Client)

	nResponse, _, err := client.ItsmIntegration.GetAllItsmIntegrationSettings()
	if err != nil {
		return foundItem, err
	}

	for _, item := range *nResponse {
		if item.Name == vName {
			return &item, err
		}
	}

	return foundItem, err
}
