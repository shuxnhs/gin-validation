package formatter

var (
	ValidTypeInt     = "int"
	ValidTypeBool    = "bool"
	ValidTypeFloat   = "float"
	ValidTypeMap     = "map"
	ValidTypeString  = "string"
	ValidTypeStrings = "strings"
)

var (
	IntFormatter    intFormatter
	BoolFormatter   boolFormatter
	FloatFormatter  floatFormatter
	StringFormatter stringFormatter
)

type Rule map[string]interface{}
