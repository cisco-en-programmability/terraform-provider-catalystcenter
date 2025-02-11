
resource "catalystcenter_images_cco_sync" "example" {
    provider = meraki
}

output "catalystcenter_images_cco_sync_example" {
    value = catalystcenter_images_cco_sync.example
}