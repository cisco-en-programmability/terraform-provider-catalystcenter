
resource "catalystcenter_sda_multicast_v1_update" "example" {
  provider = meraki
  parameters {

    fabric_id        = "string"
    replication_mode = "string"
  }
}

output "catalystcenter_sda_multicast_v1_update_example" {
  value = catalystcenter_sda_multicast_v1_update.example
}