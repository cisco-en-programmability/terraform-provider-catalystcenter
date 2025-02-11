
data "catalystcenter_disasterrecovery_system_operationstatus" "example" {
    provider = catalystcenter
}

output "catalystcenter_disasterrecovery_system_operationstatus_example" {
    value = data.catalystcenter_disasterrecovery_system_operationstatus.example.item
}
