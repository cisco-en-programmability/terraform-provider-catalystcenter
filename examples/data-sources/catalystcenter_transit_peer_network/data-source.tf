
data "catalystcenter_transit_peer_network" "example" {
    provider = catalystcenter
    transit_peer_network_name = "string"
}

output "catalystcenter_transit_peer_network_example" {
    value = data.catalystcenter_transit_peer_network.example.item
}
