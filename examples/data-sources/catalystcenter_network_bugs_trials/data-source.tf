
data "catalystcenter_network_bugs_trials" "example" {
  provider = catalystcenter
}

output "catalystcenter_network_bugs_trials_example" {
  value = data.catalystcenter_network_bugs_trials.example.item
}
