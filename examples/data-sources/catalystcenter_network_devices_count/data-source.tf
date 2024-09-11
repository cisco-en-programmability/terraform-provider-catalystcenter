
data "catalystcenter_network_devices_count" "example" {
  provider              = catalystcenter
  attribute             = "string"
  end_time              = 1609459200
  family                = "string"
  health_score          = "string"
  id                    = "string"
  mac_address           = "string"
  maintenance_mode      = "false"
  management_ip_address = "string"
  role                  = "string"
  serial_number         = "string"
  site_hierarchy        = "string"
  site_hierarchy_id     = "string"
  site_id               = "string"
  software_version      = "string"
  start_time            = 1609459200
  type                  = "string"
  view                  = "string"
}

output "catalystcenter_network_devices_count_example" {
  value = data.catalystcenter_network_devices_count.example.item
}
