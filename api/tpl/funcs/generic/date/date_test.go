package date

import (
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	ns = New(&deps.Deps{})
)

type noStringer struct{}

func TestNamespace_Date(t *testing.T) {

	tm, err := time.Parse("02 Jan 06 15:04:05 MST", "22 May 90 20:39:39 GMT")
	if err != nil {
		t.Error(err)
	}

	got, err := ns.Date("02/01/2006", tm)
	assert.Equal(t, "22/05/1990", got)
}

func TestNamespace_DateInZone(t *testing.T) {

	tt := map[string]struct {
		zone string
		time func(tm time.Time) interface{}
	}{
		"time.Time": {
			"UTC",
			func(tm time.Time) interface{} {
				return tm
			},
		},
		"*time.Time": {
			"UTC",
			func(tm time.Time) interface{} {
				return &tm
			},
		},
		"int64": {
			"UTC",
			func(tm time.Time) interface{} {
				return int64(643408779)
			},
		},
		"int32": {
			"UTC",
			func(tm time.Time) interface{} {
				return int32(643408779)
			},
		},
		"int": {
			"UTC",
			func(tm time.Time) interface{} {
				return 643408779
			},
		},
		"No Input": {
			"UTC",
			func(tm time.Time) interface{} {
				return 643408779
			},
		},
		"Invalid Timezone": {
			"wrongval",
			func(tm time.Time) interface{} {
				return 643408779
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			tm, err := time.Parse("02 Jan 06 15:04:05 MST", "22 May 90 20:39:39 GMT")
			if err != nil {
				t.Error(err)
			}
			got, _ := ns.DateInZone("02 Jan 06 15:04 -0700", test.time(tm), test.zone)
			assert.Equal(t, got, "22 May 90 20:39 +0000")
		})
	}

	t.Run("Error", func(t *testing.T) {
		_, err := ns.DateInZone("02 Jan 06 15:04 -0700", noStringer{}, "GMT")
		assert.Contains(t, err.Error(), "unable to cast date.noStringer{} of type date.noStringer to Time")
	})
}

func TestNamespace_Ago(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  string
	}{
		"Default": {
			"defaultval",
			"0s",
		},
		"Negative": {
			time.Now().Add(10 * time.Second),
			"-10s",
		},
		"time.Time": {
			time.Now().Add(-125 * time.Second),
			"2m5s",
		},
		"int64": {
			int64(1608235408),
			time.Since(time.Unix(1608235408, 0)).Round(time.Second).String(),
		},
		"int": {
			1608235408,
			time.Since(time.Unix(1608235408, 0)).Round(time.Second).String(),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Ago(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_Duration(t *testing.T) {

	tt := map[string]struct {
		input interface{}
		want  string
	}{
		"Minutes": {
			"90",
			"1m30s",
		},
		"Minutes 2": {
			"120",
			"2m0s",
		},
		"Hours": {
			"3600",
			"1h0m0s",
		},
		"Hours 2": {
			"9845",
			"2h44m5s",
		},
		"Days": {
			"103784",
			"28h49m44s",
		},
		"Days 2": {
			"117412",
			"32h36m52s",
		},
		"No Value": {
			nil,
			"",
		},
		"int64": {
			int64(10),
			"10s",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := ns.Duration(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_HTMLDate(t *testing.T) {
	got, _ := ns.HTMLDate(0)
	assert.Equal(t, "1970-01-01", got)
}

func TestNamespace_HTMLDateInZone(t *testing.T) {
	got, _ := ns.HTMLDateInZone(0, "GMT")
	assert.Equal(t, "1970-01-01", got)
}
