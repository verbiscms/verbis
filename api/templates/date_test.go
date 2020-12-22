package templates

import (
	"testing"
	"time"
)

func Test_Date(t *testing.T) {
	f := newTestSuite()

	tm, err := time.Parse("02 Jan 06 15:04:05 MST", "22 May 90 20:39:39 GMT")
	if err != nil {
		t.Error(err)
	}

	tpl := `{{ date  "02/01/2006" .Time }}`
	runtv(t, f, tpl, "22/05/1990", map[string]interface{}{"Time": tm})
}

func Test_DateInZone(t *testing.T) {
	f := newTestSuite()

	tt := map[string]struct {
		zone string
		time func(tm time.Time) interface{}
	}{
		"time.Time": {
			zone: "UTC",
			time: func(tm time.Time) interface{} {
				return tm
			},
		},
		"*time.Time": {
			zone: "UTC",
			time: func(tm time.Time) interface{} {
				return &tm
			},
		},
		"int64": {
			zone: "UTC",
			time: func(tm time.Time) interface{} {
				return int64(643408779)
			},
		},
		"int32": {
			zone: "UTC",
			time: func(tm time.Time) interface{} {
				return int32(643408779)
			},
		},
		"int": {
			zone: "UTC",
			time: func(tm time.Time) interface{} {
				return 643408779
			},
		},
		"No Input": {
			zone: "UTC",
			time: func(tm time.Time) interface{} {
				return 643408779
			},
		},
		"Invalid Timezone": {
			zone: "wrongval",
			time: func(tm time.Time) interface{} {
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
			tpl := `{{ dateInZone  "02 Jan 06 15:04 -0700" .Time "` + test.zone + `" }}`
			runtv(t, f, tpl, "22 May 90 20:39 +0000", map[string]interface{}{"Time": test.time(tm)})
		})
	}

	t.Run("No Input", func(t *testing.T) {
		tpl := `{{ dateInZone  "02 Jan 06 15:04 -0700" .Time "UTC" }}`
		loc, _ := time.LoadLocation("UTC")
		runtv(t, f, tpl, time.Now().In(loc).Format("02 Jan 06 15:04 -0700"), map[string]interface{}{"Time": ""})
	})
}

func Test_Ago(t *testing.T) {
	f := newTestSuite()

	tt := map[string]struct {
		input interface{}
		want  string
	}{
		"Default": {
			input: "defaultval",
			want:  "0s",
		},
		"Negative": {
			input: time.Now().Add(10 * time.Second),
			want:  "-10s",
		},
		"time.Time": {
			input: time.Now().Add(-125 * time.Second),
			want:  "2m5s",
		},
		"int64": {
			input: int64(1608235408),
			want:  time.Since(time.Unix(1608235408, 0)).Round(time.Second).String(),
		},
		"int": {
			input: 1608235408,
			want:  time.Since(time.Unix(1608235408, 0)).Round(time.Second).String(),
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			tpl := `{{ ago .Time }}`
			runtv(t, f, tpl, test.want, map[string]interface{}{"Time": test.input})
		})
	}
}

func Test_HTMLDate(t *testing.T) {
	f := newTestSuite()
	tpl := `{{ htmlDate 0 }}`
	runt(t, f, tpl, "1970-01-01")
}

func Test_HTMLDateInZone(t *testing.T) {
	f := newTestSuite()
	tpl := `{{ htmlDateInZone 0 "GMT" }}`
	runt(t, f, tpl, "1970-01-01")
}

func Test_Duration(t *testing.T) {
	f := newTestSuite()

	tt := map[string]struct {
		input interface{}
		want  string
	}{
		"Minutes": {
			input: "90",
			want:  "1m30s",
		},
		"Minutes 2": {
			input: "120",
			want:  "2m0s",
		},
		"Hours": {
			input: "3600",
			want:  "1h0m0s",
		},
		"Hours 2": {
			input: "9845",
			want:  "2h44m5s",
		},
		"Days": {
			input: "103784",
			want:  "28h49m44s",
		},
		"Days 2": {
			input: "117412",
			want:  "32h36m52s",
		},
		"No Value": {
			input: nil,
			want:  "0s",
		},
		"int64": {
			input: int64(10),
			want:  "10s",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			tpl := "{{ duration .Secs }}"
			runtv(t, f, tpl, test.want, map[string]interface{}{"Secs": test.input})
		})
	}
}
