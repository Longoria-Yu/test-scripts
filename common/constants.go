package common

import "strconv"

const (
	Http           = "http"
	Timeout        = 500
	DefaultRetries = 1

	DefaultReturnSize = 40 // the size returned to the client, including search results, ads, etc
	MaxReturnSize     = 80 // the upper bound size returned to the client, including search results, ads, etc

	// endpoint names
	Search40                  = "search40"
	Search34                  = "search34"
	Profile33                 = "profile33"
	Spc33                     = "spc33"
	InternalSearch10          = "internal10"
	Following31               = "following31"
	MoreFromThisSeller10      = "mfts10"
	Hyperlocal10              = "hyperlocal10"
	Hyperlocal20              = "hyperlocal20"
	ProductSearch10           = "productsearch10"
	ProductHomeFeed10         = "producthomefeed10"
	CGProductVariantCounts10  = "cgproductvc10"
	CGProductSearchListings10 = "cgproductlistings10"

	YouMayAlsoLike11    = "ymal11"
	RecommendedForYou10 = "rfy10"
	ContinueShoppingFor = "csf10"

	// severity level reported to Sentry API
	Crit = "critical"
	Warn = "warning"

	PlatformAndroid = "android"
	PlatformIOS     = "ios"
	PlatformWeb     = "web"
)

// gRPC metadata keys
const (
	ImpressionEndpoint = "impression_endpoint"
)

// Cache Use Case
const (
	KeywordSearch  = "keyword"
	CategoryBrowse = "browse"
	ProfileOwn     = "profile:own"
	ProfileOthers  = "profile:others"
	Following      = "following"
	Spc            = "spc"
)

const (
	Singapore   = 1880251
	Philippines = 1694008
	Malaysia    = 1733045
	Indonesia   = 1643084
	HongKong    = 1819730
	Taiwan      = 1668284
	Australia   = 2077456
	Canada      = 6251999
	NewZealand  = 2186224
)

const (
	SingaporeCode   = "SG"
	PhilippinesCode = "PH"
	MalaysiaCode    = "MY"
	IndonesiaCode   = "ID"
	HongKongCode    = "HK"
	TaiwanCode      = "TW"
	AustraliaCode   = "AU"
	CanadaCode      = "CA"
	NewZealandCode  = "NZ"
)

var (
	CountryMapping = map[string]string{
		SingaporeCode:   strconv.Itoa(Singapore),
		PhilippinesCode: strconv.Itoa(Philippines),
		MalaysiaCode:    strconv.Itoa(Malaysia),
		IndonesiaCode:   strconv.Itoa(Indonesia),
		HongKongCode:    strconv.Itoa(HongKong),
		TaiwanCode:      strconv.Itoa(Taiwan),
		AustraliaCode:   strconv.Itoa(Australia),
		CanadaCode:      strconv.Itoa(Canada),
		NewZealandCode:  strconv.Itoa(NewZealand),
	}
	CountryMapIdToCode = map[int64]string{
		Singapore:   SingaporeCode,
		Philippines: PhilippinesCode,
		Malaysia:    MalaysiaCode,
		Indonesia:   IndonesiaCode,
		HongKong:    HongKongCode,
		Taiwan:      TaiwanCode,
		Australia:   AustraliaCode,
		Canada:      CanadaCode,
		NewZealand:  NewZealandCode,
	}
	CountryMapStringIdToCode = map[string]string{
		strconv.Itoa(Singapore):   SingaporeCode,
		strconv.Itoa(Philippines): PhilippinesCode,
		strconv.Itoa(Malaysia):    MalaysiaCode,
		strconv.Itoa(Indonesia):   IndonesiaCode,
		strconv.Itoa(HongKong):    HongKongCode,
		strconv.Itoa(Taiwan):      TaiwanCode,
		strconv.Itoa(Australia):   AustraliaCode,
		strconv.Itoa(Canada):      CanadaCode,
		strconv.Itoa(NewZealand):  NewZealandCode,
	}
	CountryIdCurrencyMapping = map[string]string{
		strconv.Itoa(Singapore):   "S$",
		strconv.Itoa(Philippines): "PHP",
		strconv.Itoa(Malaysia):    "RM",
		strconv.Itoa(Indonesia):   "Rp",
		strconv.Itoa(HongKong):    "HK$",
		strconv.Itoa(Taiwan):      "NT$",
		strconv.Itoa(Australia):   "A$",
		strconv.Itoa(Canada):      "C$",
		strconv.Itoa(NewZealand):  "NZ$",
	}
)

func IsValidCountry(countryId int64) bool {
	_, ok := CountryMapIdToCode[countryId]
	return ok
}
