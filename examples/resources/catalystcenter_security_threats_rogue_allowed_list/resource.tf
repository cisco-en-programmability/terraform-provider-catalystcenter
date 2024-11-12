
resource "catalystcenter_security_threats_rogue_allowed_list" "example" {
  provider = catalystcenter

  parameters {

    category    = 1
    mac_address = "string"
  }
}

output "catalystcenter_security_threats_rogue_allowed_list_example" {
  value = catalystcenter_security_threats_rogue_allowed_list.example
}