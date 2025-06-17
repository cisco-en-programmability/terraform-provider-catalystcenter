
data "catalystcenter_network_devices" "example" {
  provider              = catalystcenter
  attribute             = "string"
  end_time              = 1609459200
  fabric_role           = "string"
  fabric_site_id        = "string"
  family                = "string"
  health_score          = "string"
  id                    = "string"
  l2_vn                 = "string"
  l3_vn                 = "string"
  limit                 = 1
  mac_address           = "string"
  maintenance_mode      = "false"
  management_ip_address = "string"
  offset                = 1
  order                 = "string"
  role                  = "string"
  serial_number         = "string"
  site_hierarchy        = "string"
  site_hierarchy_id     = "string"
  site_id               = "string"
  software_version      = "string"
  sort_by               = "string"
  start_time            = 1609459200
  transit_network_id    = "string"
  type                  = "string"
  view                  = "string"
}

output "catalystcenter_network_devices_example" {
  value = data.catalystcenter_network_devices.example.items
}

data "catalystcenter_network_devices" "example" {
  provider   = catalystcenter
  attribute  = "string"
  end_time   = 1609459200
  id         = "string"
  start_time = 1609459200
  view       = "string"
}

output "catalystcenter_network_devices_example" {
  value = data.catalystcenter_network_devices.example.item
}
