
data "catalystcenter_users_external_authentication" "example" {
  provider = catalystcenter
}

output "catalystcenter_users_external_authentication_example" {
  value = data.catalystcenter_users_external_authentication.example.item
}
