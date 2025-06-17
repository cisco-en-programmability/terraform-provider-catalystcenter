
resource "catalystcenter_cisco_imcs_id" "example" {
  provider = catalystcenter
 
  parameters {

    id         = "string"
    ip_address = "string"
    password   = "******"
    username   = "string"
  }
}

output "catalystcenter_cisco_imcs_id_example" {
  value = catalystcenter_cisco_imcs_id.example
}