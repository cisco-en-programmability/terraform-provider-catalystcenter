
resource "catalystcenter_network_bugs_trials_create" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_bugs_trials_create_example" {
  value = catalystcenter_network_bugs_trials_create.example
}