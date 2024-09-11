
data "catalystcenter_sda_extranet_policies_count" "example" {
  provider = catalystcenter
}

output "catalystcenter_sda_extranet_policies_count_example" {
  value = data.catalystcenter_sda_extranet_policies_count.example.item
}
