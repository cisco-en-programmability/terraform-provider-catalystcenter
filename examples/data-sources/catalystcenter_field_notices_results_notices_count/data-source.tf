
data "catalystcenter_field_notices_results_notices_count" "example" {
  provider     = catalystcenter
  device_count = 1.0
  id           = "string"
  type         = "string"
}

output "catalystcenter_field_notices_results_notices_count_example" {
  value = data.catalystcenter_field_notices_results_notices_count.example.item
}
