
data "catalystcenter_reserve_ip_subpool" "example" {
  provider                = catalystcenter
  group_name              = "string"
  ignore_inherited_groups = "false"
  limit                   = 1
  offset                  = 1
  pool_usage              = "string"
  site_id                 = "string"
}

output "catalystcenter_reserve_ip_subpool_example" {
  value = data.catalystcenter_reserve_ip_subpool.example.items
}
