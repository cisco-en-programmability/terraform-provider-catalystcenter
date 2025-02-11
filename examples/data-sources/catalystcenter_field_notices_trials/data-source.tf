
data "catalystcenter_field_notices_trials" "example" {
    provider = catalystcenter
}

output "catalystcenter_field_notices_trials_example" {
    value = data.catalystcenter_field_notices_trials.example.item
}
