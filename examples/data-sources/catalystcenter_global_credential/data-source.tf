
data "catalystcenter_global_credential" "example" {
  provider            = catalystcenter
  credential_sub_type = "string"
  order               = "string"
  sort_by             = "string"
}

output "catalystcenter_global_credential_example" {
  value = data.catalystcenter_global_credential.example.items
}

data "catalystcenter_global_credential" "example" {
  provider = catalystcenter
  id       = "string"
}

output "catalystcenter_global_credential_example" {
  value = data.catalystcenter_global_credential.example.item
}
