package middlewares

var DefaultOptions = map[string]string{
	"BDNCaseProb":                "0.7",
	"BDNHexValueProb":            "0.7",
	"BDNSpacingMaxElems":         "4",
	"BDNOIDPrependZerosMaxElems": "4",

	"FiltSpacingMaxSpaces":            "4",
	"FiltTimestampGarbageMaxChars":    "6",
	"FiltAddBoolMaxDepth":             "4",
	"FiltDblNegBoolMaxDepth":          "1",
	"FiltANRSubstringMaxElems":        "4",
	"FiltAddWildcardProb":             "0.7",
	"FiltPrependZerosMaxElems":        "4",
	"FiltGarbageMaxElems":             "4",
	"FiltGarbageMaxSize":              "6",
	"FiltGarbageCharset":              "abcdefghijklmnopqrsutwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789",
	"FiltOIDAttributeMaxElems":        "4",
	"FiltCaseProb":                    "0.7",
	"FiltHexValueProb":                "0.7",
	"FiltOIDAttributePrependOID":      "true",
	"FiltEqExtensibleAppendDN":        "false",
	"FiltBitwiseDecompositionMaxBits": "32",

	"AttrsDuplicateProb":              "0.7",
	"AttrsGarbageExistingMaxElems":    "4",
	"AttrsGarbageNonExistingMaxElems": "4",
	"AttrsGarbageNonExistingMaxSize":  "6",
	"AttrsOIDSpacingMaxElems":         "4",
}

var DefaultOptionsKeys = []string{
	"BDNCaseProb",
	"BDNHexValueProb",
	"BDNSpacingMaxElems",
	"BDNOIDPrependZerosMaxElems",

	"FiltSpacingMaxSpaces",
	"FiltTimestampGarbageMaxChars",
	"FiltAddBoolMaxDepth",
	"FiltDblNegBoolMaxDepth",
	"FiltANRSubstringMaxElems",
	"FiltAddWildcardProb",
	"FiltPrependZerosMaxElems",
	"FiltGarbageMaxElems",
	"FiltGarbageMaxSize",
	"FiltGarbageCharset",
	"FiltOIDAttributeMaxElems",
	"FiltCaseProb",
	"FiltHexValueProb",
	"FiltOIDAttributePrependOID",
	"FiltEqExtensibleAppendDN",
	"FiltBitwiseDecompositionMaxBits",

	"AttrsDuplicateProb",
	"AttrsGarbageExistingMaxElems",
	"AttrsGarbageNonExistingMaxElems",
	"AttrsGarbageNonExistingMaxSize",
	"AttrsOIDSpacingMaxElems",
}