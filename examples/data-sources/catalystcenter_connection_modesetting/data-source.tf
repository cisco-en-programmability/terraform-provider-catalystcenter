
data "catalystcenter_connection_modesetting" "example" {
  provider = catalystcenter
}

output "catalystcenter_connection_modesetting_example" {
  value = data.catalystcenter_connection_modesetting.example.item
}
