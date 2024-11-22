package filtermid

import (
	"math/rand"
	"strings"

	"github.com/Macmod/ldapx/parser"
)

/*
	Value Middlewares

	References:
	- DEFCON32 - MaLDAPtive
	- Microsoft Open Specifications - MS-ADTS
	  https://learn.microsoft.com/en-us/openspecs/windows_protocols/ms-adts/d2435927-0999-4c62-8c6d-13ba31a52e1a)
*/

func ApproxMatchFilterObf() FilterMiddleware {
	return LeafApplierFilterMiddleware(
		func(filter parser.Filter) parser.Filter {
			switch f := filter.(type) {
			case *parser.FilterEqualityMatch:
				return &parser.FilterApproxMatch{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: f.AssertionValue,
				}
			default:
				return filter
			}
		},
	)
}

func RandHexValueFilterObf(prob float32) func(parser.Filter) parser.Filter {
	applyHexEncoding := func(attr string, value string) string {
		tokenFormat, err := parser.GetAttributeTokenFormat(attr)
		if err == nil && tokenFormat == parser.TokenStringUnicode {
			return RandomlyHexEncodeString(value, prob)
		}
		return value
	}

	applier := LeafApplierFilterMiddleware(
		func(filter parser.Filter) parser.Filter {
			switch f := filter.(type) {
			case *parser.FilterEqualityMatch:
				return &parser.FilterEqualityMatch{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: applyHexEncoding(f.AttributeDesc, f.AssertionValue),
				}

			case *parser.FilterSubstring:
				newSubstrings := make([]parser.SubstringFilter, len(f.Substrings))
				for i, sub := range f.Substrings {
					newSubstrings[i] = parser.SubstringFilter{
						Initial: applyHexEncoding("name", sub.Initial),
						Final:   applyHexEncoding("name", sub.Final),
					}

					newAny := make([]string, len(sub.Any))
					for j, _any := range sub.Any {
						newAny[j] = applyHexEncoding("name", _any)
					}
					newSubstrings[i].Any = newAny
				}
				return &parser.FilterSubstring{
					AttributeDesc: f.AttributeDesc,
					Substrings:    newSubstrings,
				}

			case *parser.FilterGreaterOrEqual:
				return &parser.FilterGreaterOrEqual{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: applyHexEncoding(f.AttributeDesc, f.AssertionValue),
				}

			case *parser.FilterLessOrEqual:
				return &parser.FilterLessOrEqual{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: applyHexEncoding(f.AttributeDesc, f.AssertionValue),
				}

			case *parser.FilterApproxMatch:
				return &parser.FilterApproxMatch{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: applyHexEncoding(f.AttributeDesc, f.AssertionValue),
				}

			case *parser.FilterExtensibleMatch:
				return &parser.FilterExtensibleMatch{
					MatchingRule:  f.MatchingRule,
					AttributeDesc: f.AttributeDesc,
					MatchValue:    applyHexEncoding(f.AttributeDesc, f.MatchValue),
					DNAttributes:  f.DNAttributes,
				}

			default:
				return filter
			}
		},
	)

	return applier
}

func RandTimestampSuffixFilterObf(prepend bool, append bool, maxChars int) func(parser.Filter) parser.Filter {
	replaceTimestampFixed := func(value string) string {
		return ReplaceTimestamp(value, prepend, append, maxChars)
	}

	applier := LeafApplierFilterMiddleware(
		func(filter parser.Filter) parser.Filter {
			switch f := filter.(type) {
			case *parser.FilterEqualityMatch:
				return &parser.FilterEqualityMatch{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: replaceTimestampFixed(f.AssertionValue),
				}
			case *parser.FilterSubstring:
				newSubstrings := make([]parser.SubstringFilter, len(f.Substrings))
				for i, sub := range f.Substrings {
					newSubstrings[i] = parser.SubstringFilter{
						Initial: replaceTimestampFixed(sub.Initial),
						Final:   replaceTimestampFixed(sub.Final),
					}
					newAny := make([]string, len(sub.Any))
					for j, _any := range sub.Any {
						newAny[j] = replaceTimestampFixed(_any)
					}
					newSubstrings[i].Any = newAny
				}
				return &parser.FilterSubstring{
					AttributeDesc: f.AttributeDesc,
					Substrings:    newSubstrings,
				}

			case *parser.FilterGreaterOrEqual:
				return &parser.FilterGreaterOrEqual{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: replaceTimestampFixed(f.AssertionValue),
				}

			case *parser.FilterLessOrEqual:
				return &parser.FilterLessOrEqual{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: replaceTimestampFixed(f.AssertionValue),
				}

			case *parser.FilterApproxMatch:
				return &parser.FilterApproxMatch{
					AttributeDesc:  f.AttributeDesc,
					AssertionValue: replaceTimestampFixed(f.AssertionValue),
				}

			case *parser.FilterExtensibleMatch:
				return &parser.FilterExtensibleMatch{
					MatchingRule:  f.MatchingRule,
					AttributeDesc: f.AttributeDesc,
					MatchValue:    replaceTimestampFixed(f.MatchValue),
					DNAttributes:  f.DNAttributes,
				}
			}
			return filter
		},
	)

	return applier
}

// Prepended 0's FilterObf
func RandPrependZerosFilterObf(maxZeros int) func(parser.Filter) parser.Filter {
	prependZerosFixed := func(attrName string, value string) string {
		tokenFormat, err := parser.GetAttributeTokenFormat(attrName)
		if err != nil || tokenFormat != parser.TokenIntEnumeration &&
			tokenFormat != parser.TokenIntTimeInterval &&
			tokenFormat != parser.TokenBitwise {
			return value
		}

		return PrependZeros(value, maxZeros)
	}

	return LeafApplierFilterMiddleware(func(filter parser.Filter) parser.Filter {
		switch f := filter.(type) {
		case *parser.FilterEqualityMatch:
			return &parser.FilterEqualityMatch{
				AttributeDesc:  f.AttributeDesc,
				AssertionValue: prependZerosFixed(f.AttributeDesc, f.AssertionValue),
			}
		case *parser.FilterSubstring:
			newSubstrings := make([]parser.SubstringFilter, len(f.Substrings))
			for i, sub := range f.Substrings {
				newSubstrings[i] = parser.SubstringFilter{
					Initial: prependZerosFixed(f.AttributeDesc, sub.Initial),
					Final:   prependZerosFixed(f.AttributeDesc, sub.Final),
				}
				newAny := make([]string, len(sub.Any))
				for j, _any := range sub.Any {
					newAny[j] = prependZerosFixed(f.AttributeDesc, _any)
				}
				newSubstrings[i].Any = newAny
			}
			return &parser.FilterSubstring{
				AttributeDesc: f.AttributeDesc,
				Substrings:    newSubstrings,
			}
		case *parser.FilterGreaterOrEqual:
			return &parser.FilterGreaterOrEqual{
				AttributeDesc:  f.AttributeDesc,
				AssertionValue: prependZerosFixed(f.AttributeDesc, f.AssertionValue),
			}
		case *parser.FilterLessOrEqual:
			return &parser.FilterLessOrEqual{
				AttributeDesc:  f.AttributeDesc,
				AssertionValue: prependZerosFixed(f.AttributeDesc, f.AssertionValue),
			}
		case *parser.FilterApproxMatch:
			return &parser.FilterApproxMatch{
				AttributeDesc:  f.AttributeDesc,
				AssertionValue: prependZerosFixed(f.AttributeDesc, f.AssertionValue),
			}
		case *parser.FilterExtensibleMatch:
			return &parser.FilterExtensibleMatch{
				MatchingRule:  f.MatchingRule,
				AttributeDesc: f.AttributeDesc,
				MatchValue:    prependZerosFixed(f.AttributeDesc, f.MatchValue),
				DNAttributes:  f.DNAttributes,
			}
		}
		return filter
	})
}

func RandSpacingFilterObf(maxSpaces int) func(f parser.Filter) parser.Filter {
	return LeafApplierFilterMiddleware(func(f parser.Filter) parser.Filter {
		switch v := f.(type) {
		case *parser.FilterEqualityMatch:
			tokenType, err := parser.GetAttributeTokenFormat(v.AttributeDesc)

			if strings.ToLower(v.AttributeDesc) == "anr" {
				v.AssertionValue = AddANRSpacing(v.AssertionValue, maxSpaces)
			} else if err == nil && tokenType == parser.TokenDNString {
				v.AssertionValue = AddDNSpacing(v.AssertionValue, maxSpaces)
			}
		case *parser.FilterSubstring:
			if v.AttributeDesc == "aNR" {
				for i := range v.Substrings {
					if v.Substrings[i].Initial != "" {
						v.Substrings[i].Initial = AddANRSpacing(v.Substrings[i].Initial, maxSpaces)
					}
					if v.Substrings[i].Final != "" {
						v.Substrings[i].Final = AddANRSpacing(v.Substrings[i].Final, maxSpaces)
					}
				}
			}
		case *parser.FilterGreaterOrEqual, *parser.FilterLessOrEqual:
			if v.(*parser.FilterGreaterOrEqual).AttributeDesc == "aNR" {
				v.(*parser.FilterGreaterOrEqual).AssertionValue = AddANRSpacing(v.(*parser.FilterGreaterOrEqual).AssertionValue, maxSpaces)
			}
		case *parser.FilterApproxMatch:
			tokenType, err := parser.GetAttributeTokenFormat(v.AttributeDesc)

			if strings.ToLower(v.AttributeDesc) == "anr" {
				v.AssertionValue = AddANRSpacing(v.AssertionValue, maxSpaces)
			} else if err == nil && tokenType == parser.TokenDNString {
				v.AssertionValue = AddDNSpacing(v.AssertionValue, maxSpaces)
			}
		}
		return f
	})
}

func RandAddWildcardFilterObf(prob float32) func(parser.Filter) parser.Filter {
	return LeafApplierFilterMiddleware(func(filter parser.Filter) parser.Filter {
		switch f := filter.(type) {
		case *parser.FilterEqualityMatch:
			if rand.Float32() < prob {
				// Only apply to string attributes
				tokenType, err := parser.GetAttributeTokenFormat(f.AttributeDesc)
				if err == nil && tokenType == parser.TokenStringUnicode {
					chars := []rune(f.AssertionValue)
					splitPoint := rand.Intn(len(chars))

					return &parser.FilterSubstring{
						AttributeDesc: f.AttributeDesc,
						Substrings: []parser.SubstringFilter{{
							Initial: string(chars[:splitPoint]),
							Final:   string(chars[splitPoint:]),
						}},
					}
				}
			}
			return f

		case *parser.FilterSubstring:
			if rand.Float32() < prob && len(f.Substrings) > 0 {
				// Pick a random substring and split it
				subIdx := rand.Intn(len(f.Substrings))
				sub := f.Substrings[subIdx]

				// Choose what to split (initial, any, or final)
				var target string
				if sub.Initial != "" {
					target = sub.Initial
					sub.Initial = ""
				} else if len(sub.Any) > 0 {
					anyIdx := rand.Intn(len(sub.Any))
					target = sub.Any[anyIdx]
					sub.Any = append(sub.Any[:anyIdx], sub.Any[anyIdx+1:]...)
				} else if sub.Final != "" {
					target = sub.Final
					sub.Final = ""
				}

				if target != "" {
					chars := []rune(target)
					splitPoint := rand.Intn(len(chars))
					part1 := string(chars[:splitPoint])
					part2 := string(chars[splitPoint:])

					if part1 != "" {
						sub.Any = append(sub.Any, part1)
					}
					if part2 != "" {
						sub.Any = append(sub.Any, part2)
					}
				}

				f.Substrings[subIdx] = sub
			}
			return f

		default:
			return filter
		}
	})
}
