
data "catalystcenter_network_bugs_results_bugs_count" "example" {
  provider     = catalystcenter
  device_count = 1.0
  id           = "string"
  severity     = "string"
}

output "catalystcenter_network_bugs_results_bugs_count_example" {
  value = data.catalystcenter_network_bugs_results_bugs_count.example.item
}
