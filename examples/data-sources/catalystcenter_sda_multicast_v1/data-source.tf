
data "catalystcenter_sda_multicast_v1" "example" {
  provider  = catalystcenter
  fabric_id = "string"
  limit     = 1
  offset    = 1
}

output "catalystcenter_sda_multicast_v1_example" {
  value = data.catalystcenter_sda_multicast_v1.example.items
}
