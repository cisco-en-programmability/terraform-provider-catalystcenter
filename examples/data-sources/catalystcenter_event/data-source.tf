
data "catalystcenter_event" "example" {
  provider = catalystcenter
  event_id = "string"
  limit    = 1
  offset   = 1
  order    = "string"
  sort_by  = "string"
  tags     = "string"
}

output "catalystcenter_event_example" {
  value = data.catalystcenter_event.example.items
}
