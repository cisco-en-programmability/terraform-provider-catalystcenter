
resource "catalystcenter_analytics_cmdb_endpoints" "example" {
  provider = catalystcenter
  parameters = [{

    asset_tag             = "string"
    department            = "string"
    display_name          = "string"
    last_update_timestamp = 1
    location              = "string"
    mac_address           = "string"
    managed_by            = "string"
    model                 = "string"
    model_category        = "string"
    serial_number         = "string"
  }]
}

output "catalystcenter_analytics_cmdb_endpoints_example" {
  value = catalystcenter_analytics_cmdb_endpoints.example
}