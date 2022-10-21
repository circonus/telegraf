//go:build !custom || outputs || outputs.circonus

package all

import _ "github.com/influxdata/telegraf/plugins/outputs/circonus" // register plugin
