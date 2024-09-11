
data "catalystcenter_client_detail" "example" {
  provider    = catalystcenter
  mac_address = "string"
  timestamp   = 1.0
}

output "catalystcenter_client_detail_example" {
  value = data.catalystcenter_client_detail.example.item
}
