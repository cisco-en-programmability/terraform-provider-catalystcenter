
data "catalystcenter_nfv_provision_detail" "example" {
  provider  = catalystcenter
  device_ip = "string"
}

output "catalystcenter_nfv_provision_detail_example" {
  value = data.catalystcenter_nfv_provision_detail.example.item
}
