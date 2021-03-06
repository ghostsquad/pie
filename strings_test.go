package pie_test

import (
	"github.com/elliotchance/pie"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var stringsContainsTests = []struct {
	ss       pie.Strings
	contains string
	expected bool
}{
	{nil, "a", false},
	{nil, "", false},
	{pie.Strings{"a", "b", "c"}, "a", true},
	{pie.Strings{"a", "b", "c"}, "b", true},
	{pie.Strings{"a", "b", "c"}, "c", true},
	{pie.Strings{"a", "b", "c"}, "A", false},
	{pie.Strings{"a", "b", "c"}, "", false},
	{pie.Strings{"a", "b", "c"}, "d", false},
	{pie.Strings{"a", "", "c"}, "", true},
}

func TestStrings_Contains(t *testing.T) {
	for _, test := range stringsContainsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, test.ss.Contains(test.contains))
		})
	}
}

func TestStringsContains(t *testing.T) {
	for _, test := range stringsContainsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expected, pie.StringsContains(test.ss, test.contains))
		})
	}
}

var stringsOnlyAndWithoutTests = []struct {
	ss                pie.Strings
	condition         func(string) bool
	expectedOnly      pie.Strings
	expectedWithout   pie.Strings
	expectedTransform pie.Strings
}{
	{
		nil,
		func(s string) bool {
			return s == ""
		},
		nil,
		nil,
		nil,
	},
	{
		pie.Strings{"a", "b", "c"},
		func(s string) bool {
			return s != "b"
		},
		pie.Strings{"a", "c"},
		pie.Strings{"b"},
		pie.Strings{"A", "B", "C"},
	},
}

func TestStrings_Only(t *testing.T) {
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedOnly, test.ss.Only(test.condition))
		})
	}
}

func TestStringsOnly(t *testing.T) {
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, []string(test.expectedOnly), pie.StringsOnly(test.ss, test.condition))
		})
	}
}

func TestStrings_Without(t *testing.T) {
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedWithout, test.ss.Without(test.condition))
		})
	}
}

func TestStringsWithout(t *testing.T) {
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, []string(test.expectedWithout), pie.StringsWithout(test.ss, test.condition))
		})
	}
}

func TestStrings_Transform(t *testing.T) {
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.expectedTransform, test.ss.Transform(strings.ToUpper))
		})
	}
}

func TestStringsTransform(t *testing.T) {
	for _, test := range stringsOnlyAndWithoutTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, []string(test.expectedTransform), pie.StringsTransform(test.ss, strings.ToUpper))
		})
	}
}

var firstAndLastTests = []struct {
	ss             pie.Strings
	first, firstOr string
	last, lastOr   string
}{
	{
		nil,
		"",
		"default1",
		"",
		"default2",
	},
	{
		pie.Strings{"foo"},
		"foo",
		"foo",
		"foo",
		"foo",
	},
	{
		pie.Strings{"a", "b"},
		"a",
		"a",
		"b",
		"b",
	},
	{
		pie.Strings{"a", "b", "c"},
		"a",
		"a",
		"c",
		"c",
	},
}

func TestStrings_FirstOr(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.firstOr, test.ss.FirstOr("default1"))
		})
	}
}

func TestStringsFirstOr(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.firstOr, pie.StringsFirstOr(test.ss, "default1"))
		})
	}
}

func TestStrings_LastOr(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.lastOr, test.ss.LastOr("default2"))
		})
	}
}

func TestStringsLastOr(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.lastOr, pie.StringsLastOr(test.ss, "default2"))
		})
	}
}

func TestStrings_First(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.first, test.ss.First())
		})
	}
}

func TestStringsFirst(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.first, pie.StringsFirst(test.ss))
		})
	}
}

func TestStrings_Last(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.last, test.ss.Last())
		})
	}
}

func TestStringsLast(t *testing.T) {
	for _, test := range firstAndLastTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.last, pie.StringsLast(test.ss))
		})
	}
}

var stringsStatsTests = []struct {
	ss       []string
	min, max string
	len      int
}{
	{
		nil,
		"",
		"",
		0,
	},
	{
		[]string{},
		"",
		"",
		0,
	},
	{
		[]string{"foo"},
		"foo",
		"foo",
		1,
	},
	{
		[]string{"bar", "Baz", "qux", "foo"},
		"Baz",
		"qux",
		4,
	},
}

func TestStringsMin(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.min, pie.StringsMin(test.ss))
		})
	}
}

func TestStrings_Min(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.min, pie.Strings(test.ss).Min())
		})
	}
}

func TestStringsMax(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.max, pie.StringsMax(test.ss))
		})
	}
}

func TestStrings_Max(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.max, pie.Strings(test.ss).Max())
		})
	}
}


func TestStrings_Len(t *testing.T) {
	for _, test := range stringsStatsTests {
		t.Run("", func(t *testing.T) {
			assert.Equal(t, test.len, pie.Strings(test.ss).Len())
		})
	}
}
