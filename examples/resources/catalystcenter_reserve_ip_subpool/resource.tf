
resource "catalystcenter_reserve_ip_subpool" "example" {
  provider = catalystcenter

}

output "catalystcenter_reserve_ip_subpool_example" {
  value = catalystcenter_reserve_ip_subpool.example
}