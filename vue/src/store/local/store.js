import { reactive } from "vue";

export const store = reactive({
	stakePopup: false,
	votePopup: false,
	currentVoteSelection: {
		proposal_id: null,
        current_vote: null
	},
	delegateInput: false,
	undelegateInput: false,
	valueDelegate: true,
	valueUndelegate: true,
	currentValidator: null,
	undelegate: false,
	redelegate: false,
	pools: [
		{
		  "id": "1",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/14F9BC3E44B8A9C1BE1FB08980FAB87034C9905EF17CF2F5008FC085218811CC",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos1m7uyxn26sz6w4755k6rch4dc2fj6cmzajkszvn",
		  "pool_coin_denom": "poolDFB8434D5A80B4EAFA94B6878BD5B85265AC6C5D37204AB899B1C3C52543DA7E"
		},
		{
		  "id": "2",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/68A333688E5B07451F95555F8FE510E43EF9D3D44DF0909964F92081EF9BE5A7",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos1uu07zd5p52pmwq27fexys54saknjejt6r0dqjx",
		  "pool_coin_denom": "poolE71FE13681A283B7015E4E4C4852B0EDA72CC97A5CDE2ECA2A6C8C06C86AC775"
		},
		{
		  "id": "3",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/12DA42304EE1CE96071F712AA4D58186AD11C3165C0DCDA71E017A54F3935E66",
			"ibc/42E47A5BA708EBE6E0C227006254F2784E209F4DBD3C6BB77EDC4B29EF875E8E"
		  ],
		  "reserve_account_address": "cosmos16cu6n9q5v3khzdkxt35yt58tx3twm57k2ukak3",
		  "pool_coin_denom": "poolD639A99414646D7136C65C6845D0EB3456EDD3D6C2C43050D3FA3A24995B0E75"
		},
		{
		  "id": "4",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/81D08BC39FB520EBD948CF017910DD69702D34BF5AC160F76D3B5CFC444EBCE0",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos19wwzfqeu4gngeyypasj3dya8yn5dxslu9yzyh3",
		  "pool_coin_denom": "pool2B9C24833CAA268C9081EC251693A724E8D343FC25A841FF00FD37B047BA4DEA"
		},
		{
		  "id": "5",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/2181AAB0218EAC24BC9F86BD1364FBBFA3E6E3FCC25E88E3E68C15DC6E752D86",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos1xtwsv6lff8jlmnracz0tke78xqwse22h4sez9u",
		  "pool_coin_denom": "pool32DD066BE949E5FDCC7DC09EBB67C7301D0CA957C2EF56A39B37430165447DAC"
		},
		{
		  "id": "6",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/1FBDD58D438B4D04D26CBFB2E722C18984A0F1A52468C4F42F37D102F3D3F399",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos1f0a0cjvhwmfs5naq6cpnzd0spnzwl3msjdn737",
		  "pool_coin_denom": "pool4BFAFC499776D30A4FA0D6033135F00CC4EFC770D19A74CAD37433B579F77FC0"
		},
		{
		  "id": "7",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/42E47A5BA708EBE6E0C227006254F2784E209F4DBD3C6BB77EDC4B29EF875E8E",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos10t3ersye68vgejz6n752phzk2zlcmhsdmzg40l",
		  "pool_coin_denom": "pool7AE391C099D1D88CC85A9FA8A0DC5650BF8DDE0DCE7D0824C073802C020A7747"
		},
		{
		  "id": "8",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/12DA42304EE1CE96071F712AA4D58186AD11C3165C0DCDA71E017A54F3935E66",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos172q9nqx9fc28f0wv7u80tl5gruacal8cuw0tzg",
		  "pool_coin_denom": "poolF2805980C54E1474BDCCF70EF5FE881F3B8EFCF8BA3198765C01D91904521788"
		},
		{
		  "id": "9",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/C932ADFE2B4216397A4F17458B6E4468499B86C3BC8116180F85D799D6F5CC1B",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos1h4034aagk8cx3stc78tr0hcjd95wcy9tjk9kpk",
		  "pool_coin_denom": "poolBD5F1AF7A8B1F068C178F1D637DF126968EC10AB204A10116E320B2B8AF4FAC2"
		},
		{
		  "id": "10",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/5BB694D466CCF099EF73F165F88472AF51D9C4991EAA42BD1168C5304712CC0D",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos1k3tuayjqcgsupamf2tldv5r0wsm4ahfclajeas",
		  "pool_coin_denom": "poolB457CE9240C221C0F76952FED6506F74375EDD38B32A6020B7DDDFD5A4867D5C"
		},
		{
		  "id": "11",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/B2B5AEE174062FA7804AC95223D8169852F8F58962C51C66391C272C838258B7",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos16yfpu788xxhaxhl2z08elgqyfg289ae64vvakd",
		  "pool_coin_denom": "poolD1121E78E731AFD35FEA13CF9FA0044A1472F73A0EE784160CCAAAAE5C7AAD7E"
		},
		{
		  "id": "12",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/E070CE91CC4BD15AEC9B5788C0826755AAD35052A3037E9AC62BE70B4C9A7DBB",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos1vrhmq7qh66ce8206mfs3gp9cuywc944h7pq2w3",
		  "pool_coin_denom": "pool60EFB07817D6B193A9FADA611404B8E11D82D6B7F0D10D57D3134C93E2BF7414"
		},
		{
		  "id": "13",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/DCD1849E20837BC8FB2C252A7AE1D8AA7A1876911EE669E6CE6FDF9FEA54083D",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos1e79cg7vh7h4e9wwgmwhyzet0v8ttuuyt3cvtzk",
		  "pool_coin_denom": "poolCF8B847997F5EB92B9C8DBAE41656F61D6BE708B1B42D31063291813014AD63F"
		},
		{
		  "id": "14",
		  "type_id": 1,
		  "reserve_coin_denoms": [
			"ibc/1D5826F7EDE6E3B13009FEF994DC9CAAF15CC24CA7A9FF436FFB2E56FD72F54F",
			"uatom"
		  ],
		  "reserve_account_address": "cosmos14jd00dywgjt6p2d0zz0y9pjxfg80qmnuc3hlhv",
		  "pool_coin_denom": "poolAC9AF7B48E4497A0A9AF109E4286464A0EF06E7C35AD79198F03AB17A6A4CCA7"
		}
	],
	IBCToTokenMap: {
		"uatom": "ATOM",
		"ibc/14F9BC3E44B8A9C1BE1FB08980FAB87034C9905EF17CF2F5008FC085218811CC": "OSMO",
		"ibc/68A333688E5B07451F95555F8FE510E43EF9D3D44DF0909964F92081EF9BE5A7": "IOV",
		"ibc/12DA42304EE1CE96071F712AA4D58186AD11C3165C0DCDA71E017A54F3935E66": "IRIS",
		"ibc/42E47A5BA708EBE6E0C227006254F2784E209F4DBD3C6BB77EDC4B29EF875E8E": "DVPN",
		"ibc/81D08BC39FB520EBD948CF017910DD69702D34BF5AC160F76D3B5CFC444EBCE0": "XPRT",
		"ibc/2181AAB0218EAC24BC9F86BD1364FBBFA3E6E3FCC25E88E3E68C15DC6E752D86": "AKT",
		"ibc/1FBDD58D438B4D04D26CBFB2E722C18984A0F1A52468C4F42F37D102F3D3F399": "REGEN",
		"ibc/C932ADFE2B4216397A4F17458B6E4468499B86C3BC8116180F85D799D6F5CC1B": "CRO",
		"ibc/5BB694D466CCF099EF73F165F88472AF51D9C4991EAA42BD1168C5304712CC0D": "ION",
		"ibc/B2B5AEE174062FA7804AC95223D8169852F8F58962C51C66391C272C838258B7": "IXO",
		"ibc/E070CE91CC4BD15AEC9B5788C0826755AAD35052A3037E9AC62BE70B4C9A7DBB": "NGM",
		"ibc/DCD1849E20837BC8FB2C252A7AE1D8AA7A1876911EE669E6CE6FDF9FEA54083D": "ROWAN",
		"ibc/1D5826F7EDE6E3B13009FEF994DC9CAAF15CC24CA7A9FF436FFB2E56FD72F54F": "LIKE"
	},
	TokenToIBCMap: {
		"ATOM" : "uatom",
		"OSMO": "ibc/14F9BC3E44B8A9C1BE1FB08980FAB87034C9905EF17CF2F5008FC085218811CC",
		"IOV": "ibc/68A333688E5B07451F95555F8FE510E43EF9D3D44DF0909964F92081EF9BE5A7",
		"IRIS": "ibc/12DA42304EE1CE96071F712AA4D58186AD11C3165C0DCDA71E017A54F3935E66",
		"DVPN": "ibc/42E47A5BA708EBE6E0C227006254F2784E209F4DBD3C6BB77EDC4B29EF875E8E",
		"XPRT": "ibc/81D08BC39FB520EBD948CF017910DD69702D34BF5AC160F76D3B5CFC444EBCE0",
		"AKT": "ibc/2181AAB0218EAC24BC9F86BD1364FBBFA3E6E3FCC25E88E3E68C15DC6E752D86",
		"REGEN": "ibc/1FBDD58D438B4D04D26CBFB2E722C18984A0F1A52468C4F42F37D102F3D3F399",
		"CRO": "ibc/C932ADFE2B4216397A4F17458B6E4468499B86C3BC8116180F85D799D6F5CC1B",
		"ION": "ibc/5BB694D466CCF099EF73F165F88472AF51D9C4991EAA42BD1168C5304712CC0D",
		"IXO": "ibc/B2B5AEE174062FA7804AC95223D8169852F8F58962C51C66391C272C838258B7",
		"NGM": "ibc/E070CE91CC4BD15AEC9B5788C0826755AAD35052A3037E9AC62BE70B4C9A7DBB",
		"ROWAN": "ibc/DCD1849E20837BC8FB2C252A7AE1D8AA7A1876911EE669E6CE6FDF9FEA54083D",
		"LIKE": "ibc/1D5826F7EDE6E3B13009FEF994DC9CAAF15CC24CA7A9FF436FFB2E56FD72F54F"
	},
	TokenToUMap: {
		"ATOM": "uatom",
		"OSMO": "uosmo",
		"IOV": "uiov",
		"IRIS": "uiris",
		"DVPN": "udvpn",
		"XPRT": "uxprt",
		"AKT": "uakt",
		"REGEN": "uregen",
		"CRO": "basecro",
		"ION": "uion",
		"IXO": "uixo",
		"NGM": "ungm",
		"ROWAN": "xrowan",
		"LIKE": "nanolike"
	},
	lastTxHash: null,
	lastTxLog: null,
	showTxSuccess: false,
	showTxFail: false,
	showTxStatus: false,
	govProps: [],
	sendingTx: false,
	twitter: null,
	updateStake: false,
	manageStake: false,
	validators: []
})