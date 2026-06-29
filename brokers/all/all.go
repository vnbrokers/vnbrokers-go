// Package all imports every built-in broker adapter so their factories register.
package all

import (
	_ "github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	_ "github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	_ "github.com/vnbrokers/vnbrokers-go/brokers/fhsc"
	_ "github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	_ "github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
)
