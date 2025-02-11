
resource "catalystcenter_network_bugs_trials_create" "example" {
    provider = meraki
}

output "catalystcenter_network_bugs_trials_create_example" {
    value = catalystcenter_network_bugs_trials_create.example
}