
data "catalystcenter_field_notices_results_notices" "example" {
  provider     = catalystcenter
  device_count = 1.0
  id           = "string"
  limit        = 1
  offset       = 1
  order        = "string"
  sort_by      = "string"
  type         = "string"
}

output "catalystcenter_field_notices_results_notices_example" {
  value = data.catalystcenter_field_notices_results_notices.example.items
}
