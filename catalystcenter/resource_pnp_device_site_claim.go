package catalystcenter

import (
	"context"

	"fmt"
	"reflect"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v3/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceAction
func resourcePnpDeviceSiteClaim() *schema.Resource {
	return &schema.Resource{
		Description: `It performs create operation on Device Onboarding (PnP).

- Claim a device based on Catalyst Center Site-based design process. Some required parameters differ based on device
platform:
Default/StackSwitch: imageInfo, configInfo.
AccessPoints: rfProfile.
Sensors: sensorProfile.
CatalystWLC/MobilityExpress/EWC: staticIP, subnetMask, gateway. vlanId and ipInterfaceName are also allowed for Catalyst
9800 WLCs.
`,

		CreateContext: resourcePnpDeviceSiteClaimCreate,
		ReadContext:   resourcePnpDeviceSiteClaimRead,
		DeleteContext: resourcePnpDeviceSiteClaimDelete,
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

						"response": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": &schema.Schema{
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
						"config_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"config_id": &schema.Schema{
										Description: `Config Id`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"config_parameters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										ForceNew: true,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"key": &schema.Schema{
													Description: `Key`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
												"value": &schema.Schema{
													Description: `Value`,
													Type:        schema.TypeString,
													Optional:    true,
													ForceNew:    true,
													Computed:    true,
												},
											},
										},
									},
								},
							},
						},
						"device_id": &schema.Schema{
							Description: `Device Id`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"gateway": &schema.Schema{
							Description: `for CatalystWLC/MobilityExpress
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"hostname": &schema.Schema{
							Description: `hostname to configure on Device.
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"image_info": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							ForceNew: true,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"image_id": &schema.Schema{
										Description: `Image Id`,
										Type:        schema.TypeString,
										Optional:    true,
										ForceNew:    true,
										Computed:    true,
									},
									"skip": &schema.Schema{
										Description: `Skip`,
										// Type:        schema.TypeBool,
										Type:         schema.TypeString,
										ValidateFunc: validateStringHasValueFunc([]string{"", "true", "false"}),
										Optional:     true,
										ForceNew:     true,
										Computed:     true,
									},
								},
							},
						},
						"ip_interface_name": &schema.Schema{
							Description: `for Catalyst 9800 WLC
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"rf_profile": &schema.Schema{
							Description: `for Access Points
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"sensor_profile": &schema.Schema{
							Description: `for Sensors
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"site_id": &schema.Schema{
							Description: `Site Id`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"static_ip": &schema.Schema{
							Description: `for CatalystWLC/MobilityExpress
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"subnet_mask": &schema.Schema{
							Description: `for CatalystWLC/MobilityExpress
`,
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
							Computed: true,
						},
						"type": &schema.Schema{
							Description: `Type`,
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Computed:    true,
						},
						"vlan_id": &schema.Schema{
							Description: `for Catalyst 9800 WLC
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
	}
}

func resourcePnpDeviceSiteClaimCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics

	request1 := expandRequestPnpDeviceSiteClaimClaimADeviceToASite(ctx, "parameters.0", d)

	// has_unknown_response: None

	response1, restyResp1, err := client.DeviceOnboardingPnp.ClaimADeviceToASite(request1)

	if err != nil || response1 == nil {
		if restyResp1 != nil {
			log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
		}
		diags = append(diags, diagError(
			"Failure when executing ClaimADeviceToASite", err))
		return diags
	}

	log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

	vItem1 := flattenDeviceOnboardingPnpClaimADeviceToASiteItem(response1)
	if err := d.Set("item", vItem1); err != nil {
		diags = append(diags, diagError(
			"Failure when setting ClaimADeviceToASite response",
			err))
		return diags
	}

	d.SetId(getUnixTimeString())
	return diags

}
func resourcePnpDeviceSiteClaimRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)
	var diags diag.Diagnostics
	return diags
}

func resourcePnpDeviceSiteClaimDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	//client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	return diags
}

func expandRequestPnpDeviceSiteClaimClaimADeviceToASite(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASite {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASite{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".device_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".device_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".device_id")))) {
		request.DeviceID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".site_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".site_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".site_id")))) {
		request.SiteID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".type")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".type")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".type")))) {
		request.Type = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_info")))) {
		request.ImageInfo = *expandRequestPnpDeviceSiteClaimClaimADeviceToASiteImageInfo(ctx, key+".image_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_info")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_info")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_info")))) {
		request.ConfigInfo = *expandRequestPnpDeviceSiteClaimClaimADeviceToASiteConfigInfo(ctx, key+".config_info.0", d)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".rf_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".rf_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".rf_profile")))) {
		request.RfProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".static_ip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".static_ip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".static_ip")))) {
		request.StaticIP = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".subnet_mask")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".subnet_mask")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".subnet_mask")))) {
		request.SubnetMask = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".gateway")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".gateway")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".gateway")))) {
		request.Gateway = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".vlan_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".vlan_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".vlan_id")))) {
		request.VLANID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".ip_interface_name")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".ip_interface_name")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".ip_interface_name")))) {
		request.IPInterfaceName = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".sensor_profile")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".sensor_profile")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".sensor_profile")))) {
		request.SensorProfile = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".hostname")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".hostname")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".hostname")))) {
		request.Hostname = interfaceToString(v)
	}
	return &request
}

func expandRequestPnpDeviceSiteClaimClaimADeviceToASiteImageInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteImageInfo {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteImageInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".image_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".image_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".image_id")))) {
		request.ImageID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".skip")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".skip")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".skip")))) {
		request.Skip = interfaceToBool(v)
	}
	return &request
}

func expandRequestPnpDeviceSiteClaimClaimADeviceToASiteConfigInfo(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfo{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_id")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_id")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_id")))) {
		request.ConfigID = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".config_parameters")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".config_parameters")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".config_parameters")))) {
		request.ConfigParameters = *expandRequestPnpDeviceSiteClaimClaimADeviceToASiteConfigInfoConfigParametersArray(ctx, key+".config_parameters", d)
	}
	return &request
}

func expandRequestPnpDeviceSiteClaimClaimADeviceToASiteConfigInfoConfigParametersArray(ctx context.Context, key string, d *schema.ResourceData) *[]catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters {
	request := []catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters{}
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
		i := expandRequestPnpDeviceSiteClaimClaimADeviceToASiteConfigInfoConfigParameters(ctx, fmt.Sprintf("%s.%d", key, item_no), d)
		if i != nil {
			request = append(request, *i)
		}
	}
	return &request
}

func expandRequestPnpDeviceSiteClaimClaimADeviceToASiteConfigInfoConfigParameters(ctx context.Context, key string, d *schema.ResourceData) *catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters {
	request := catalystcentersdkgo.RequestDeviceOnboardingPnpClaimADeviceToASiteConfigInfoConfigParameters{}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".key")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".key")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".key")))) {
		request.Key = interfaceToString(v)
	}
	if v, ok := d.GetOkExists(fixKeyAccess(key + ".value")); !isEmptyValue(reflect.ValueOf(d.Get(fixKeyAccess(key+".value")))) && (ok || !reflect.DeepEqual(v, d.Get(fixKeyAccess(key+".value")))) {
		request.Value = interfaceToString(v)
	}
	return &request
}

func flattenDeviceOnboardingPnpClaimADeviceToASiteItem(item *catalystcentersdkgo.ResponseDeviceOnboardingPnpClaimADeviceToASite) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["response"] = item.Response
	respItem["version"] = item.Version
	return []map[string]interface{}{
		respItem,
	}
}
