package proto // import "github.com/synerex/synerex_proto"
//

// ChannelTypeVersion is a common version number for Synerex Providers
const ChannelTypeVersion = "0.1.12" // string for pbase version

// if you change this number you should update "ChannelTypeVersion"
const ChannelTypeMax = 32 // Default Synerex Server channel size

// Channel Types
const (
	RIDE_SHARE         uint32 = 1  // Rideshare Service Information
	AD_SERVICE         uint32 = 2  // Advertisement Service Information
	LIB_SERVICE        uint32 = 3  // Public Library Service Information
	PT_SERVICE         uint32 = 4  // Public Transit Information
	ROUTING_SERVICE    uint32 = 5  // Routing Service
	MARKETING_SERVICE  uint32 = 6  // Marketing (Ad/Enquate)
	FLUENTD_SERVICE    uint32 = 7  // Fluentd Service (td-agent/fluetnd)
	MEETING_SERVICE    uint32 = 8  // RPA Meetinng Service (rpa provider)
	STORAGE_SERVICE    uint32 = 9  // Storage Service (storage providers)
	RETRIEVAL_SERVICE  uint32 = 10 // Retrieval Service (retrieval providers)
	PEOPLE_COUNTER_SVC uint32 = 11 // People Counter Service (Pflow providers)
	AREA_COUNTER_SVC   uint32 = 12 // Area counter service
	PEOPLE_AGENT_SVC   uint32 = 13 // people agent service
	GEOGRAPHIC_SVC     uint32 = 14 // Geographical mapping service
	JSON_DATA_SVC      uint32 = 15 // Json data service
	MQTT_GATEWAY_SVC   uint32 = 16 // MQTT Gateway service
	WAREHOUSE_SVC      uint32 = 17 // Warehouse Execution/Management service
	PEOPLE_FLOW_SVC    uint32 = 18 // People Flow service
	GRIDEYE_SVC        uint32 = 19 // Grid Eye service
	LATENT_DMD_SVC     uint32 = 20 // Latent Demand service
	LATENT_DMD_DSP_SVC uint32 = 21 // Latent Demand Display service
	ALT_PT_SVC         uint32 = 22 // Alternative Public Transit Service
	PEOPLE_WT_SVC      uint32 = 23 // People Weighted Tree service
)
