
data "catalystcenter_event_subscription_details_rest" "example" {
    provider = catalystcenter
    instance_id = "string"
    limit = 1
    name = "string"
    offset = 1
    order = "string"
    sort_by = "string"
}

output "catalystcenter_event_subscription_details_rest_example" {
    value = data.catalystcenter_event_subscription_details_rest.example.items
}
