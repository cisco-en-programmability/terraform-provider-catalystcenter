
data "catalystcenter_disasterrecovery_system_status" "example" {
  provider = catalystcenter
}

output "catalystcenter_disasterrecovery_system_status_example" {
  value = data.catalystcenter_disasterrecovery_system_status.example.item
}
