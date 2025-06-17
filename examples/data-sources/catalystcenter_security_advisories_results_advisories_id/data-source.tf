
data "catalystcenter_security_advisories_results_advisories_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_security_advisories_results_advisories_id_example" {
  value = data.catalystcenter_security_advisories_results_advisories_id.example.item
}
