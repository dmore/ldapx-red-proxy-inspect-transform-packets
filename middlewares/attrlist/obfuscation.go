package attrlistmid

import (
	"math/rand"
	"strings"
	"unicode"

	"github.com/Macmod/ldapx/parser"
)

// RandCaseAttrListObf randomly changes case of attribute names
func RandCaseAttrListObf(prob float32) func([]string) []string {
	return func(attrs []string) []string {
		result := make([]string, len(attrs))
		for i, attr := range attrs {
			var builder strings.Builder
			for _, c := range attr {
				if rand.Float32() < prob {
					if unicode.IsUpper(c) {
						builder.WriteString(strings.ToLower(string(c)))
					} else {
						builder.WriteString(strings.ToUpper(string(c)))
					}
				} else {
					builder.WriteRune(c)
				}
			}
			result[i] = builder.String()
		}
		return result
	}
}

// OIDAttributeAttrListObf converts attributes to their OID form
func OIDAttributeAttrListObf() func([]string) []string {
	return func(attrs []string) []string {
		result := make([]string, len(attrs))
		for i, attr := range attrs {
			if oid, exists := parser.OidsMap[strings.ToLower(attr)]; exists {
				result[i] = oid
			} else {
				result[i] = attr
			}
		}
		return result
	}
}

// RandSpacingAttrListObf adds random spacing to attributes
func RandSpacingAttrListObf(maxSpaces int) func([]string) []string {
	return func(attrs []string) []string {
		result := make([]string, len(attrs))
		for i, attr := range attrs {
			/*
				var builder strings.Builder
				for _, c := range attr {
					spaces := strings.Repeat(" ", rand.Intn(maxSpaces))
					builder.WriteString(spaces)
					builder.WriteRune(c)
				}
			*/
			result[i] = attr + strings.Repeat(" ", rand.Intn(maxSpaces))
		}
		return result
	}
}

// DuplicateAttrListObf duplicates random attributes
func DuplicateAttrListObf(minDups int, maxDups int) func([]string) []string {
	return func(attrs []string) []string {
		result := make([]string, 0)
		for _, attr := range attrs {
			duplicates := 1 + minDups + rand.Intn(maxDups)
			for i := 0; i < duplicates; i++ {
				result = append(result, attr)
			}
		}
		return result
	}
}

// GarbageExistingAttrListObf adds garbage to existing attributes
func GarbageExistingAttrListObf(maxGarbage int) func([]string) []string {
	return func(attrs []string) []string {
		if len(attrs) == 0 {
			return attrs
		}

		result := make([]string, len(attrs))
		copy(result, attrs)

		// Get all attribute names from parser.AttrContexts
		existingAttrs := make([]string, 0, len(parser.AttrContexts))
		for attr := range parser.AttrContexts {
			existingAttrs = append(existingAttrs, attr)
		}

		garbageCount := 1 + rand.Intn(maxGarbage)
		for i := 0; i < garbageCount; i++ {
			randomAttr := existingAttrs[rand.Intn(len(existingAttrs))]
			result = append(result, randomAttr)
		}
		return result
	}
}

// GarbageNonExistingAttrListObf adds completely new garbage attributes
func GarbageNonExistingAttrListObf(maxGarbage int, garbageSize int, garbageCharset string) func([]string) []string {
	return func(attrs []string) []string {
		if len(attrs) == 0 {
			return attrs
		}

		result := make([]string, len(attrs))
		copy(result, attrs)

		garbageCount := 1 + rand.Intn(maxGarbage)
		for i := 0; i < garbageCount; i++ {
			var garbage string
			exists := true
			for exists {
				garbageBytes := make([]byte, garbageSize)
				for j := 0; j < garbageSize; j++ {
					garbageBytes[j] = garbageCharset[rand.Intn(len(garbageCharset))]
				}
				garbage = string(garbageBytes)
				_, exists = parser.OidsMap[strings.ToLower(garbage)]
			}
			result = append(result, garbage)
		}
		return result
	}
}

// AddWildcardAttrListObf adds wildcards to the attributes list
func AddWildcardAttrListObf() func([]string) []string {
	return func(attrs []string) []string {
		result := make([]string, len(attrs))
		copy(result, attrs)
		result = append(result, "*")
		return result
	}
}

func ReplaceWithWildcardAttrListObf() func([]string) []string {
	return func(attrs []string) []string {
		return []string{"*"}
	}
}

func ReplaceWithEmptyAttrListObf() func([]string) []string {
	return func(attrs []string) []string {
		return []string{}
	}
}

func ReorderListAttrListObf() func([]string) []string {
	return func(attrs []string) []string {
		result := make([]string, len(attrs))
		copy(result, attrs)
		rand.Shuffle(len(result), func(i, j int) {
			result[i], result[j] = result[j], result[i]
		})
		return result
	}
}