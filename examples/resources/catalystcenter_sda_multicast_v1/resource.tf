
resource "catalystcenter_sda_multicast_v1" "example" {
  provider = catalystcenter

  parameters {

    fabric_id        = "string"
    replication_mode = "string"
  }
}

output "catalystcenter_sda_multicast_v1_example" {
  value = catalystcenter_sda_multicast_v1.example
}