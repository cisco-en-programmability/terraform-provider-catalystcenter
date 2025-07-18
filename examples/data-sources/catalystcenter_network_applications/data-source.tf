
data "catalystcenter_network_applications" "example" {
  provider           = catalystcenter
  application_name   = "string"
  attribute          = "string"
  business_relevance = "string"
  end_time           = 1609459200
  limit              = 1
  offset             = 1
  order              = "string"
  site_id            = "string"
  sort_by            = "string"
  ssid               = "string"
  start_time         = 1609459200
  xca_lle_rid        = "string"
}

output "catalystcenter_network_applications_example" {
  value = data.catalystcenter_network_applications.example.items
}
