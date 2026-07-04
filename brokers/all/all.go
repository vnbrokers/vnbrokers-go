// Package all imports every built-in broker adapter so their factories register.
package all

import (
	// Register DNSE broker factory via package init.
	_ "github.com/vnbrokers/vnbrokers-go/brokers/dnse"
	// Register Entrade broker factory via package init.
	_ "github.com/vnbrokers/vnbrokers-go/brokers/entrade"
	// Register FHSC broker factory via package init.
	_ "github.com/vnbrokers/vnbrokers-go/brokers/fhsc"
	// Register SSI broker factory via package init.
	_ "github.com/vnbrokers/vnbrokers-go/brokers/ssi"
	// Register TCBS broker factory via package init.
	_ "github.com/vnbrokers/vnbrokers-go/brokers/tcbs"
)
