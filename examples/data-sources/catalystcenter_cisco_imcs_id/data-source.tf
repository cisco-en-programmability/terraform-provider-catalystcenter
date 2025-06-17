
data "catalystcenter_cisco_imcs_id" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_cisco_imcs_id_example" {
  value = data.catalystcenter_cisco_imcs_id.example.item
}
