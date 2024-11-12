
resource "catalystcenter_integrate_ise" "example" {
  provider = meraki
  id       = "string"
  parameters {

    is_cert_accepted_by_user = "false"
  }
}

output "catalystcenter_integrate_ise_example" {
  value = catalystcenter_integrate_ise.example
}