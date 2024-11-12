package catalystcenter

import (
	"context"

	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourceInterfaceUpdate() *schema.Resource {
	return &schema.Resource{
		Description: `It performs update operation on Devices.

- Add/Update Interface description, VLAN membership, Voice VLAN and change Interface admin status ('UP'/'DOWN') from
Request body.
`,

		CreateContext: resourceInterfaceUpdateCreate,
		ReadContext:   resourceInterfaceUpdateRead,
		DeleteContext: resourceInterfaceUpdateDelete,
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

						"properties": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"task_id": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
									"url": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"type": &schema.Schema{
													Description: `Type`,
													Type:        schema.TypeString,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"required": &schema.Schema{
							Description: `Required`,
							Type:        schema.TypeList,
							Computed:    true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Computed:    true,
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
						"deployment_mode": &schema.Schema{
							Description: `deploymentMode query parameter. Preview/Deploy ['Preview' means the configuration is not pushed to the device. 'Deploy' makes the configuration pushed to the device]
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"interface_uuid": &schema.Schema{
							Description: `interfaceUuid path parameter. Interface ID
`,
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
						"admin_status": &schema.Schema{
							Description: `Admin status as ('UP'/'DOWN')
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"description": &schema.Schema{
							Description: `Description for the Interface
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"vlan_id": &schema.Schema{
							Description: `VLAN Id to be Updated
`,
							Type:     schema.TypeInt,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"voice_vlan_id": &schema.Schema{
							Description: `Voice Vlan Id to be Updated
`,
							Type:     schema.TypeInt,
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

func resourceInterfaceUpdateCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	resourceItem := *getResourceItem(d.Get("parameters"))

	vInterfaceUUID := resourceItem["interface_uuid"]

	vvInterfaceUUID := vInterfaceUUID.(string)
	request1 := expandRequestInterfaceUpdateUpdateInterfaceDetailsV1(ctx, "parameters.0", d)
	queryParams1 := catalystcentersdkgo.UpdateInterfaceDetailsV1QueryParams{}

	// has_unknown_response: None

	response1, restyResp1, err := client.Devices.UpdateInterfaceDetailsV1(vvInterfaceUUID, request1, &queryParams1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing UpdateInterfaceDetailsV1", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDevicesUpdateInterfaceDetailsV1Item(response1.Response)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting UpdateInterfaceDetailsV1 response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourceInterfaceUpdateRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourceInterfaceUpdateDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestInterfaceUpdateUpdateInterfaceDetailsV1(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDevicesUpdateInterfaceDetailsV1 {
	request := catalystcentersdkgo.RequestDevicesUpdateInterfaceDetailsV1{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".description")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".description")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".description")))) {
		request.Description = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".admin_status")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".admin_status")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".admin_status")))) {
		request.AdminStatus = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToIntPtr(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".voice_vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".voice_vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".voice_vlan_id")))) {
		request.VoiceVLANID = interfaceToIntPtr(v)
	}
	return &request
}

func flattenDevicesUpdateInterfaceDetailsV1Item(item *catalystcentersdkgo.ResponseDevicesUpdateInterfaceDetailsV1Response) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type
	respItem["properties"] = flattenDevicesUpdateInterfaceDetailsV1ItemProperties(item.Properties)
	respItem["required"] = item.Required
	return []map[string]interface{}{
		respItem,
	}
}

func flattenDevicesUpdateInterfaceDetailsV1ItemProperties(item *catalystcentersdkgo.ResponseDevicesUpdateInterfaceDetailsV1ResponseProperties) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["task_id"] = flattenDevicesUpdateInterfaceDetailsV1ItemPropertiesTaskID(item.TaskID)
	respItem["url"] = flattenDevicesUpdateInterfaceDetailsV1ItemPropertiesURL(item.URL)

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesUpdateInterfaceDetailsV1ItemPropertiesTaskID(item *catalystcentersdkgo.ResponseDevicesUpdateInterfaceDetailsV1ResponsePropertiesTaskID) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesUpdateInterfaceDetailsV1ItemPropertiesURL(item *catalystcentersdkgo.ResponseDevicesUpdateInterfaceDetailsV1ResponsePropertiesURL) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["type"] = item.Type

	return []map[string]interface{}{
		respItem,
	}

}
