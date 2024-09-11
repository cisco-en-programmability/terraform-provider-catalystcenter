
data "catalystcenter_dnac_packages" "example" {
  provider = catalystcenter
}

output "catalystcenter_dnac_packages_example" {
  value = data.catalystcenter_dnac_packages.example.items
}
