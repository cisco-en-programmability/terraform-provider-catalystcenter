
resource "catalystcenter_security_advisories_trials" "example" {
    provider = catalystcenter

}

output "catalystcenter_security_advisories_trials_example" {
    value = catalystcenter_security_advisories_trials.example
}
