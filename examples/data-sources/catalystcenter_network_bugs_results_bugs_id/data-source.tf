
data "catalystcenter_network_bugs_results_bugs_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_network_bugs_results_bugs_id_example" {
  value = data.catalystcenter_network_bugs_results_bugs_id.example.item
}
