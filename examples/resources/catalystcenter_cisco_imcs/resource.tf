
resource "catalystcenter_cisco_imcs" "example" {
  provider = catalystcenter
 
  parameters {

    ip_address = "string"
    node_id    = "string"
    password   = "******"
    username   = "string"
  }
}

output "catalystcenter_cisco_imcs_example" {
  value = catalystcenter_cisco_imcs.example
}