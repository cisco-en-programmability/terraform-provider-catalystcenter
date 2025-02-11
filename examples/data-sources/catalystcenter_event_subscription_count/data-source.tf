
data "catalystcenter_event_subscription_count" "example" {
    provider = catalystcenter
    event_ids = "string"
}

output "catalystcenter_event_subscription_count_example" {
    value = data.catalystcenter_event_subscription_count.example.item
}
