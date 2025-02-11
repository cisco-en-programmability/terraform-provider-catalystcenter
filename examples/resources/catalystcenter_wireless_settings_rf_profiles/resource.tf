
resource "catalystcenter_wireless_settings_rf_profiles" "example" {
    provider = catalystcenter

    parameters {

      default_rf_profile = "false"
      enable_radio_type6_g_hz = "false"
      enable_radio_type_a = "false"
      enable_radio_type_b = "false"
      id = "string"
      radio_type6_g_hz_properties {

        broadcast_probe_response_interval = 1
        coverage_hole_detection_properties {

          chd_client_level = 1
          chd_data_rssi_threshold = 1
          chd_exception_level = 1
          chd_voice_rssi_threshold = 1
        }
        custom_rx_sop_threshold = 1
        data_rates = "string"
        discovery_frames6_g_hz = "string"
        enable_standard_power_service = "false"
        fra_properties {

          client_reset_count = 1
          client_utilization_threshold = 1
        }
        mandatory_data_rates = "string"
        max_dbs_width = 1
        max_power_level = 1
        max_radio_clients = 1
        min_dbs_width = 1
        min_power_level = 1
        multi_bssid_properties {

          dot11ax_parameters {

            mu_mimo_down_link = "false"
            mu_mimo_up_link = "false"
            ofdma_down_link = "false"
            ofdma_up_link = "false"
          }
          dot11be_parameters {

            mu_mimo_down_link = "false"
            mu_mimo_up_link = "false"
            ofdma_down_link = "false"
            ofdma_multi_ru = "false"
            ofdma_up_link = "false"
          }
          target_wake_time = "false"
          twt_broadcast_support = "false"
        }
        parent_profile = "string"
        power_threshold_v1 = 1
        preamble_puncture = "false"
        psc_enforcing_enabled = "false"
        radio_channels = "string"
        rx_sop_threshold = "string"
        spatial_reuse_properties {

          dot11ax_non_srg_obss_packet_detect = "false"
          dot11ax_non_srg_obss_packet_detect_max_threshold = 1
          dot11ax_srg_obss_packet_detect = "false"
          dot11ax_srg_obss_packet_detect_max_threshold = 1
          dot11ax_srg_obss_packet_detect_min_threshold = 1
        }
      }
      radio_type_a_properties {

        channel_width = "string"
        coverage_hole_detection_properties {

          chd_client_level = 1
          chd_data_rssi_threshold = 1
          chd_exception_level = 1
          chd_voice_rssi_threshold = 1
        }
        custom_rx_sop_threshold = 1
        data_rates = "string"
        fra_properties {

          client_aware = "false"
          client_reset = 1
          client_select = 1
        }
        mandatory_data_rates = "string"
        max_power_level = 1
        max_radio_clients = 1
        min_power_level = 1
        parent_profile = "string"
        power_threshold_v1 = 1
        preamble_puncture = "false"
        radio_channels = "string"
        rx_sop_threshold = "string"
        spatial_reuse_properties {

          dot11ax_non_srg_obss_packet_detect = "false"
          dot11ax_non_srg_obss_packet_detect_max_threshold = 1
          dot11ax_srg_obss_packet_detect = "false"
          dot11ax_srg_obss_packet_detect_max_threshold = 1
          dot11ax_srg_obss_packet_detect_min_threshold = 1
        }
        zero_wait_dfs_enable = "false"
      }
      radio_type_b_properties {

        coverage_hole_detection_properties {

          chd_client_level = 1
          chd_data_rssi_threshold = 1
          chd_exception_level = 1
          chd_voice_rssi_threshold = 1
        }
        custom_rx_sop_threshold = 1
        data_rates = "string"
        mandatory_data_rates = "string"
        max_power_level = 1
        max_radio_clients = 1
        min_power_level = 1
        parent_profile = "string"
        power_threshold_v1 = 1
        radio_channels = "string"
        rx_sop_threshold = "string"
        spatial_reuse_properties {

          dot11ax_non_srg_obss_packet_detect = "false"
          dot11ax_non_srg_obss_packet_detect_max_threshold = 1
          dot11ax_srg_obss_packet_detect = "false"
          dot11ax_srg_obss_packet_detect_max_threshold = 1
          dot11ax_srg_obss_packet_detect_min_threshold = 1
        }
      }
      rf_profile_name = "string"
    }
}

output "catalystcenter_wireless_settings_rf_profiles_example" {
    value = catalystcenter_wireless_settings_rf_profiles.example
}
