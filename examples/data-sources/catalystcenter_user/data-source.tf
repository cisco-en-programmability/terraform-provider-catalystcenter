
data "catalystcenter_user" "example" {
    provider = catalystcenter
    auth_source = "string"
    invoke_source = "string"
}

output "catalystcenter_user_example" {
    value = data.catalystcenter_user.example.item
}
