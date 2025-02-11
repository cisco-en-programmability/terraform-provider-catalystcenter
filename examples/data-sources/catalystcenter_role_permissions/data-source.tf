
data "catalystcenter_role_permissions" "example" {
    provider = catalystcenter
}

output "catalystcenter_role_permissions_example" {
    value = data.catalystcenter_role_permissions.example.item
}
