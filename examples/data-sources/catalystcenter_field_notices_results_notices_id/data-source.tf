
data "catalystcenter_field_notices_results_notices_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_field_notices_results_notices_id_example" {
  value = data.catalystcenter_field_notices_results_notices_id.example.item
}
