//go:build !custom || outputs || outputs.amqp_ext

package all

import _ "github.com/influxdata/telegraf/plugins/outputs/amqp_ext" // register plugin
