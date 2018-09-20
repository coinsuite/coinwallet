// Copyright (c) 2013-2015 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netparams

import "github.com/coinsuite/coind/chaincfg"

// Params is used to group parameters for various networks such as the main
// network and test networks.
type Params struct {
	*chaincfg.Params
	RPCClientPort string
	RPCServerPort string
}

// GetMainNet contains parameters specific running coinwallet and
// btcd on the main network (wire.MainNet).
func GetMainNet() *Params {
	return &Params{
		Params:        chaincfg.GetMainNet(),
		RPCClientPort: "8334",
		RPCServerPort: "8332",
	}
}

// GetTestNet contains parameters specific running coinwallet and
// btcd on the test network (version 3) (wire.TestNet3).
func GetTestNet() *Params {
	return &Params{
		Params:        chaincfg.GetTestNet(),
		RPCClientPort: "18334",
		RPCServerPort: "18332",
	}
}

// GetSimNet contains parameters specific to the simulation test network
// (wire.SimNet).
func GetSimNet() *Params {
	return &Params{
		Params:        chaincfg.GetSimNet(),
		RPCClientPort: "18556",
		RPCServerPort: "18554",
	}
}
