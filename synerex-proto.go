package proto // import "github.com/synerex/synerex_proto"
//

const ChannelTypeVersion = "0.1.1"  // string for pbase version

// Channel Types
const (
	RIDE_SHARE uint32 = 1  // Rideshare Service Information
	AD_SERVICE uint32 = 2  // Advertisement Service Information
	LIB_SERVICE uint32 = 3 // Public Library Service Information
	PT_SERVICE uint32 = 4  // Public Transit Information
	ROUTING_SERVICE uint32 = 5 // Routing Service
	MARKETING_SERVICE uint32 = 6 // Marketing (Ad/Enquate)
)

const ChannelTypeMax = 20  // may double default channel size
