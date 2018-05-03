// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Ethereum network.
var MainnetBootnodes = []string{
	// Ethereum Foundation Go Bootnodes
	"enode://c99b7f8aaa0a50054bcecf3ba41c7d43228385761d3ff7de180b25cded0f989ceb845561ed3eca881ec83e4605884318e60e6d8357f6d4e88852e9c97972e163@167.99.199.110:30303", // IE
	"enode://3f2d12044546b76342d59d4a05532c14b85aa669704bfe1f864fe079415aa2c02d743e03218e57a33fb94523adb54032871a6c51b2cc5514cb7c7e35b3ed0a99@165.227.109.31:30303",  // US-WEST
	"enode://78de8a0916848094c73790ead81d1928bec737d565119932b98c6b100d944b7a95e94f847f689fc723399d2e31129d182f7ef3863f2b4c820abbf3ab2722344d@178.62.200.147:30303", // BR
	"enode://158f8aab45f6d19a6cbf4a089c2670541a8da11978a2f90dbf6a502a4a3bab80d288afdbeb7ec0ef6d92de563767f3b1ea9e8e334ca711e9f8e2df5a0385e8e6@46.101.238.104:30303", // AU


	// Ethereum Foundation C++ Bootnodes
	"enode://979b7fa28feeb35a4741660a16076f1943202cb72b6af70d327f053e248bab9ba81760f39d0701ef1d8f89cc1fbd2cacba0710a12cd5314d5e0c9021aa3637f9@5.1.83.226:30303", // DE
}

// TestnetBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var TestnetBootnodes = []string{
	"enode://30b7ab30a01c124a6cceca36863ece12c4f5fa68e3ba9b0b51407ccc002eeed3b3102d20a88f1c1d3c3154e2449317b8ef95090e77b312d5cc39354f86d5d606@52.176.7.10:30303",    // US-Azure geth
	"enode://865a63255b3bb68023b6bffd5095118fcc13e79dcf014fe4e47e065c350c7cc72af2e53eff895f11ba1bbb6a2b33271c1116ee870f266618eadfc2e78aa7349c@52.176.100.77:30303",  // US-Azure parity
	"enode://6332792c4a00e3e4ee0926ed89e0d27ef985424d97b6a45bf0f23e51f0dcb5e66b875777506458aea7af6f9e4ffb69f43f3778ee73c81ed9d34c51c4b16b0b0f@52.232.243.152:30303", // Parity
	"enode://94c15d1b9e2fe7ce56e458b9a3b672ef11894ddedd0c6f247e0f1d3487f52b66208fb4aeb8179fce6e3a749ea93ed147c37976d67af557508d199d9594c35f09@192.81.208.223:30303", // @gpip
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
	"enode://a24ac7c5484ef4ed0c5eb2d36620ba4e4aa13b8c84684e1b4aab0cebea2ae45cb4d375b77eab56516d34bfbd3c1a833fc51296ff084b770b94fb9028c4d25ccf@52.169.42.101:30303", // IE
	"enode://343149e4feefa15d882d9fe4ac7d88f885bd05ebb735e547f12e12080a9fa07c8014ca6fd7f373123488102fe5e34111f8509cf0b7de3f5b44339c9f25e87cb8@52.3.158.184:30303",  // INFURA
	"enode://b6b28890b006743680c52e64e0d16db57f28124885595fa03a562be1d2bf0f3a1da297d56b13da25fb992888fd556d4c1a27b1f39d531bde7de1921c90061cc6@159.89.28.211:30303", // AKASHA
}

// DiscoveryV5Bootnodes are the enode URLs of the P2P bootstrap nodes for the
// experimental RLPx v5 topic-discovery network.
var DiscoveryV5Bootnodes = []string{
//"enode://ecaebb39dafa9c1d811c4a2544b21a1bf604325f74c568f36ac7a9a5bb91c1b74d3378b059169fb56586735ff8b03533ac1748bec1edd07ea278bb2c40565e4f@52.161.13.240:30303",
}

var TouristCoinMainnetBootnodes = []string{
	"enode://75abd0dfe54f953b5bc75c304aaf1058e387df0bff91b80705540fb7917190c4025d13abee2f5b749b511702bcce9d33a319a4ed3b022862e90bd46f58bdd1af@188.138.1.43:30303",
	"enode://cbfed03a1389ff0d45c8da23f885856ce6c04220dfbebaa1d58f48daba6a0a1a706a11ec269ea18df255fa3e08ba832078d8577968c5e1320860a57cf38f266d@45.55.37.243:30303",
	"enode://2b8092251937e7f577c705877569030ee635da03db622c24b4bf45af29ffb46854cb87cf47d3fd3aa6497debad25777fd37af37034f14b855c000a0a825b3f9e@174.138.9.15:30303",
	"enode://ae1363a0f778a5f0b59a078a7d79abdec5dc15991c75722412def01fef166360baae2cde0a948b3cc7858e9829d59a65d6eeaca5d4f0b822d46e87bf6961a302@167.99.80.40:30303",
	"enode://45d0e42c0cbbf8643480bd5c7dfc242a1d0a6bfa181331ec6cf1b73e8d59c76a604def75883a4afc3e348a10656d531069c5913250885b0e53bdb446c28d2152@165.227.141.203:30303",
	"enode://b5ea277c0187f10f834e5509b9bcaf0398289ef3b4bcde40d02c53b6f85f0eafac1c1947d5795bce1c1ba81fe267fd1c740ba3d2a38f942cfec003ead40a5441@5.135.139.134:30311",
}


var TouristCoinTestnetBootnodes = []string{
	//"enode://1e1d25d9eafd92319ee85f502c2f5274a62029b3f9ec003cba259a600ae0def256a09ec505a569aabc4a447d2fd4df9c682243091730f83556693dab17206bd0@46.101.238.104:30303",
}
