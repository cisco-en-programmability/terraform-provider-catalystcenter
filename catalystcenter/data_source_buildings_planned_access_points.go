package catalystcenter

import (
	"context"

	"log"

	catalystcentersdkgo "github.com/cisco-en-programmability/catalystcenter-go-sdk/v2/sdk"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBuildingsPlannedAccessPoints() *schema.Resource {
	return &schema.Resource{
		Description: `It performs read operation on Devices.

- Provides a list of Planned Access Points for the Building it is requested for
`,

		ReadContext: dataSourceBuildingsPlannedAccessPointsRead,
		Schema: map[string]*schema.Schema{
			"building_id": &schema.Schema{
				Description: `buildingId path parameter. The instance UUID of the building hierarchy element
`,
				Type:     schema.TypeString,
				Required: true,
			},
			"limit": &schema.Schema{
				Description: `limit query parameter. The page size limit for the response, e.g. limit=100 will return a maximum of 100 records
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"offset": &schema.Schema{
				Description: `offset query parameter. The page offset for the response. E.g. if limit=100, offset=0 will return first 100 records, offset=1 will return next 100 records, etc.
`,
				Type:     schema.TypeFloat,
				Optional: true,
			},
			"radios": &schema.Schema{
				Description: `radios query parameter. Whether to include the planned radio details of the planned access points
`,
				Type:     schema.TypeBool,
				Optional: true,
			},

			"items": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"attributes": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"create_date": &schema.Schema{
										Description: `Created date of the planned access point
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"domain": &schema.Schema{
										Description: `Service domain to which the planned access point belongs
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"heirarchy_name": &schema.Schema{
										Description: `Hierarchy name of the planned access point
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"id": &schema.Schema{
										Description: `Unique id of the planned access point
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"instance_uuid": &schema.Schema{
										Description: `Instance uuid of the planned access point
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"mac_address": &schema.Schema{
										Description: `MAC address of the planned access point
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"name": &schema.Schema{
										Description: `Display name of the planned access point
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"source": &schema.Schema{
										Description: `Source of the data used to create the planned access point
`,
										Type:     schema.TypeString,
										Computed: true,
									},

									"type_string": &schema.Schema{
										Description: `Type string representation of the planned access point
`,
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},

						"is_sensor": &schema.Schema{
							Description: `Determines if the planned access point is sensor or not
`,
							// Type:        schema.TypeBool,
							Type:     schema.TypeString,
							Computed: true,
						},

						"location": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"altitude": &schema.Schema{
										Description: `Altitude of the planned access point's location
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"lattitude": &schema.Schema{
										Description: `Latitude of the planned access point's location
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"longtitude": &schema.Schema{
										Description: `Longitude of the planned access point's location
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						"position": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"x": &schema.Schema{
										Description: `x-coordinate of the planned access point on the map, 0,0 point being the top-left corner
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"y": &schema.Schema{
										Description: `y-coordinate of the planned access point on the map, 0,0 point being the top-left corner
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},

									"z": &schema.Schema{
										Description: `z-coordinate, or height, of the planned access point on the map
`,
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},

						"radio_count": &schema.Schema{
							Description: `Number of radios of the planned access point
`,
							Type:     schema.TypeInt,
							Computed: true,
						},

						"radios": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"antenna": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"azimuth_angle": &schema.Schema{
													Description: `Azimuth angle of the antenna
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"elevation_angle": &schema.Schema{
													Description: `Elevation angle of the antenna
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"gain": &schema.Schema{
													Description: `Gain of the antenna
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"mode": &schema.Schema{
													Description: `Mode of the antenna associated with this radio
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"name": &schema.Schema{
													Description: `Name of the antenna
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"type": &schema.Schema{
													Description: `Type of the antenna associated with this radio
`,
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"attributes": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"channel": &schema.Schema{
													Description: `Channel in which the radio operates
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"channel_string": &schema.Schema{
													Description: `Channel string representation
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"id": &schema.Schema{
													Description: `Id of the radio
`,
													Type:     schema.TypeInt,
													Computed: true,
												},

												"if_mode": &schema.Schema{
													Description: `IF mode of the radio
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"if_type_string": &schema.Schema{
													Description: `String representation of native band
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"if_type_subband": &schema.Schema{
													Description: `Sub band type of the radio
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"instance_uuid": &schema.Schema{
													Description: `Instance Uuid of the radio
`,
													Type:     schema.TypeString,
													Computed: true,
												},

												"slot_id": &schema.Schema{
													Description: `Slot number in which the radio resides in the parent access point
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},

												"tx_power_level": &schema.Schema{
													Description: `Tx Power at which this radio operates (in dBm)
`,
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},

									"is_sensor": &schema.Schema{
										Description: `Determines if it is sensor or not
`,
										// Type:        schema.TypeBool,
										Type:     schema.TypeString,
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

func dataSourceBuildingsPlannedAccessPointsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*catalystcentersdkgo.Client)

	var diags diag.Diagnostics
	vBuildingID := d.Get("building_id")
	vLimit, okLimit := d.GetOk("limit")
	vOffset, okOffset := d.GetOk("offset")
	vRadios, okRadios := d.GetOk("radios")

	selectedMethod := 1
	if selectedMethod == 1 {
		log.Printf("[DEBUG] Selected method: GetPlannedAccessPointsForBuildingV1")
		vvBuildingID := vBuildingID.(string)
		queryParams1 := catalystcentersdkgo.GetPlannedAccessPointsForBuildingV1QueryParams{}

		if okLimit {
			queryParams1.Limit = vLimit.(float64)
		}
		if okOffset {
			queryParams1.Offset = vOffset.(float64)
		}
		if okRadios {
			queryParams1.Radios = vRadios.(bool)
		}

		response1, restyResp1, err := client.Devices.GetPlannedAccessPointsForBuildingV1(vvBuildingID, &queryParams1)

		if err != nil || response1 == nil {
			if restyResp1 != nil {
				log.Printf("[DEBUG] Retrieved error response %s", restyResp1.String())
			}
			diags = append(diags, diagErrorWithAlt(
				"Failure when executing 2 GetPlannedAccessPointsForBuildingV1", err,
				"Failure at GetPlannedAccessPointsForBuildingV1, unexpected response", ""))
			return diags
		}

		log.Printf("[DEBUG] Retrieved response %+v", responseInterfaceToString(*response1))

		vItems1 := flattenDevicesGetPlannedAccessPointsForBuildingV1Items(response1.Response)
		if err := d.Set("items", vItems1); err != nil {
			diags = append(diags, diagError(
				"Failure when setting GetPlannedAccessPointsForBuildingV1 response",
				err))
			return diags
		}

		d.SetId(getUnixTimeString())
		return diags

	}
	return diags
}

func flattenDevicesGetPlannedAccessPointsForBuildingV1Items(items *[]catalystcentersdkgo.ResponseDevicesGetPlannedAccessPointsForBuildingV1Response) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attributes"] = flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsAttributes(item.Attributes)
		respItem["location"] = flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsLocation(item.Location)
		respItem["position"] = flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsPosition(item.Position)
		respItem["radio_count"] = item.RadioCount
		respItem["radios"] = flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsRadios(item.Radios)
		respItem["is_sensor"] = boolPtrToString(item.IsSensor)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsAttributes(item *catalystcentersdkgo.ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["name"] = item.Name
	respItem["type_string"] = item.TypeString
	respItem["domain"] = item.Domain
	respItem["heirarchy_name"] = item.HeirarchyName
	respItem["source"] = item.Source
	respItem["create_date"] = item.CreateDate
	respItem["mac_address"] = item.MacAddress

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsLocation(item *catalystcentersdkgo.ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseLocation) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["altitude"] = item.Altitude
	respItem["lattitude"] = item.Lattitude
	respItem["longtitude"] = item.Longtitude

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsPosition(item *catalystcentersdkgo.ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponsePosition) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["x"] = item.X
	respItem["y"] = item.Y
	respItem["z"] = item.Z

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsRadios(items *[]catalystcentersdkgo.ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadios) []map[string]interface{} {
	if items == nil {
		return nil
	}
	var respItems []map[string]interface{}
	for _, item := range *items {
		respItem := make(map[string]interface{})
		respItem["attributes"] = flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsRadiosAttributes(item.Attributes)
		respItem["antenna"] = flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsRadiosAntenna(item.Antenna)
		respItem["is_sensor"] = boolPtrToString(item.IsSensor)
		respItems = append(respItems, respItem)
	}
	return respItems
}

func flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsRadiosAttributes(item *catalystcentersdkgo.ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadiosAttributes) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["id"] = item.ID
	respItem["instance_uuid"] = item.InstanceUUID
	respItem["slot_id"] = item.SlotID
	respItem["if_type_string"] = item.IfTypeString
	respItem["if_type_subband"] = item.IfTypeSubband
	respItem["channel"] = item.Channel
	respItem["channel_string"] = item.ChannelString
	respItem["if_mode"] = item.IfMode
	respItem["tx_power_level"] = item.TxPowerLevel

	return []map[string]interface{}{
		respItem,
	}

}

func flattenDevicesGetPlannedAccessPointsForBuildingV1ItemsRadiosAntenna(item *catalystcentersdkgo.ResponseDevicesGetPlannedAccessPointsForBuildingV1ResponseRadiosAntenna) []map[string]interface{} {
	if item == nil {
		return nil
	}
	respItem := make(map[string]interface{})
	respItem["name"] = item.Name
	respItem["type"] = item.Type
	respItem["mode"] = item.Mode
	respItem["azimuth_angle"] = item.AzimuthAngle
	respItem["elevation_angle"] = item.ElevationAngle
	respItem["gain"] = item.Gain

	return []map[string]interface{}{
		respItem,
	}

}
