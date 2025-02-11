
resource "catalystcenter_field_notices_trials" "example" {
    provider = catalystcenter

}

output "catalystcenter_field_notices_trials_example" {
    value = catalystcenter_field_notices_trials.example
}
