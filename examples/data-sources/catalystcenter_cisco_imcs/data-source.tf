
data "catalystcenter_cisco_imcs" "example" {
  provider = catalystcenter
}

output "catalystcenter_cisco_imcs_example" {
  value = data.catalystcenter_cisco_imcs.example.items
}
