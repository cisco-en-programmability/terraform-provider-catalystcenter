
data "catalystcenter_global_credential_v2" "example" {
    provider = catalystcenter
}

output "catalystcenter_global_credential_v2_example" {
    value = data.catalystcenter_global_credential_v2.example.item
}
