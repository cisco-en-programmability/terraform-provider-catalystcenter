
data "catalystcenter_global_pool" "example" {
    provider = catalystcenter
    limit = 1
    offset = 1
}

output "catalystcenter_global_pool_example" {
    value = data.catalystcenter_global_pool.example.items
}
