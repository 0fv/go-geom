package wkbhex

import (
	"reflect"
	"testing"

	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/wkb"
)

func Test(t *testing.T) {
	for _, tc := range []struct {
		g   geom.T
		ndr string
		xdr string
	}{
		{
			g:   geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{1, 2}),
			ndr: "0101000000000000000000f03f0000000000000040",
		},
		{
			g:   geom.NewPoint(geom.XY).MustSetCoords(geom.Coord{2, 4}),
			xdr: "000000000140000000000000004010000000000000",
		},
		{
			g: geom.NewLineString(geom.XY).MustSetCoords([]geom.Coord{
				{-71.160281, 42.258729},
				{-71.160837, 42.259113},
				{-71.161144, 42.25932},
			}),
			ndr: "010200000003000000e44a3d0b42ca51c06ec328081e21454027bf45274bca51c0f67b629d2a214540957cec2e50ca51c07099d36531214540",
		},
		{
			g: geom.NewMultiLineString(geom.XY).MustSetCoords([][]geom.Coord{
				{{-71.160281, 42.258729}, {-71.160837, 42.259113}, {-71.161144, 42.25932}},
			}),
			ndr: "010500000001000000010200000003000000e44a3d0b42ca51c06ec328081e21454027bf45274bca51c0f67b629d2a214540957cec2e50ca51c07099d36531214540",
		},
		{
			g: geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{
				{
					{-71.1776585052917, 42.3902909739571},
					{-71.1776820268866, 42.3903701743239},
					{-71.1776063012595, 42.3903825660754},
					{-71.1775826583081, 42.3903033653531},
					{-71.1776585052917, 42.3902909739571},
				},
			}),
			ndr: "010300000001000000050000006285c7c15ecb51c0ed88fc0df531454028a46f245fcb51c009075ea6f731454047ded1e65dcb51c0781c510ef83145404871a7835dcb51c0ebdaee75f53145406285c7c15ecb51c0ed88fc0df5314540",
		},
		{
			g: geom.NewPolygon(geom.XY).MustSetCoords([][]geom.Coord{
				{
					{-76.3249866420126, 40.0476632878855},
					{-76.3249504321909, 40.047748950936},
					{-76.3247989712005, 40.047709472172},
					{-76.3248351810222, 40.0476247311309},
					{-76.3249866420126, 40.0476632878855},
				},
			}),
			ndr: "010300000001000000050000002fc5c594cc1453c01258a3d4190644402fc5e5fccb1453c01e583ba31c06444028c59f81c91453c016580f581b06444028c57f19ca1453c011583391180644402fc5c594cc1453c01258a3d419064440",
		},
		{
			g: geom.NewMultiPolygon(geom.XY).MustSetCoords([][][]geom.Coord{
				{
					{
						{-71.1031880899493, 42.3152774590236},
						{-71.1031627617667, 42.3152960829043},
						{-71.102923838298, 42.3149156848307},
						{-71.1023097974109, 42.3151969047397},
						{-71.1019285062273, 42.3147384934248},
						{-71.102505233663, 42.3144722937587},
						{-71.10277487471, 42.3141658254797},
						{-71.103113945163, 42.3142739188902},
						{-71.10324876416, 42.31402489987},
						{-71.1033002961013, 42.3140393340215},
						{-71.1033488797549, 42.3139495090772},
						{-71.103396240451, 42.3138632439557},
						{-71.1041521907712, 42.3141153348029},
						{-71.1041411411543, 42.3141545014533},
						{-71.1041287795912, 42.3142114839058},
						{-71.1041188134329, 42.3142693656241},
						{-71.1041112482575, 42.3143272556118},
						{-71.1041072845732, 42.3143851580048},
						{-71.1041057218871, 42.3144430686681},
						{-71.1041065602059, 42.3145009876017},
						{-71.1041097995362, 42.3145589148055},
						{-71.1041166403905, 42.3146168544148},
						{-71.1041258822717, 42.3146748022936},
						{-71.1041375307579, 42.3147318674446},
						{-71.1041492906949, 42.3147711126569},
						{-71.1041598612795, 42.314808571739},
						{-71.1042515013869, 42.3151287620809},
						{-71.1041173835118, 42.3150739481917},
						{-71.1040809891419, 42.3151344119048},
						{-71.1040438678912, 42.3151191367447},
						{-71.1040194562988, 42.3151832057859},
						{-71.1038734225584, 42.3151140942995},
						{-71.1038446938243, 42.3151006300338},
						{-71.1038315271889, 42.315094347535},
						{-71.1037393329282, 42.315054824985},
						{-71.1035447555574, 42.3152608696313},
						{-71.1033436658644, 42.3151648370544},
						{-71.1032580383161, 42.3152269126061},
						{-71.103223066939, 42.3152517403219},
						{-71.1031880899493, 42.3152774590236},
					},
				},
				{
					{
						{-71.1043632495873, 42.315113108546},
						{-71.1043583974082, 42.3151211109857},
						{-71.1043443253471, 42.3150676015829},
						{-71.1043850704575, 42.3150793250568},
						{-71.1043632495873, 42.315113108546},
					},
				},
			}),
			ndr: "01060000000200000001030000000100000028000000d0ea37a29ac651c00fd603035b284540fefcfb379ac651c0c0503e9f5b284540ffdddd4d96c651c033ac3b284f2845402c7c643e8cc651c027d4465f58284540b03124ff85c651c0a206d8594928454017901c728fc651c08a98cca040284540b76e11dd93c651c063faf49536284540258f3b6b99c651c041cfb5203a284540db5ab4a09bc651c02189c9f731284540055bd8789cc651c0d6a3de703228454009a89e449dc651c0790a5d7f2f2845400dd1430b9ec651c0651ab8ab2c284540d24af36daac651c0451369ee3428454041db9a3faac651c00ad1f63636284540aabac10baac651c097b3f71438284540eea5f4e1a9c651c0bebe83fa39284540599a39c2a9c651c00c8c21e03b2845401ba199b1a9c651c03cfdd9c53d28454038b50baba9c651c01231a4ab3f2845408ad88faea9c651c08d27809141284540df0c26bca9c651c0a0e06d7743284540955cd7d8a9c651c0b13d765d452845408bc19affa9c651c0f75c9043472845400c397630aac651c07fd7422249284540eb5bc961aac651c0204b796b4a28454016701f8eaac651c04b0fb4a54b28454000417d0eacc651c04e95a723562845403648f5dba9c651c009a7d75754284540fd1f4f43a9c651c0d77c0c53562845404b7c9ca7a8c651c0c454e9d255284540debc3841a8c651c0a88c5cec57284540f633b6dca5c651c034c39ca855284540aaf53664a5c651c0da78aa37552845407c64fd2ca5c651c09de8f60255284540128f4caaa3c651c0f8e06cb753284540a5b22e7aa0c651c03263da775a284540ba48c02e9dc651c0b8ff45525728454079679ac79bc651c02b3b005b59284540882cec349bc651c04658452b5a284540d0ea37a29ac651c00fd603035b284540010300000001000000050000006ef831e3adc651c07cdf57a05528454099fbd7ceadc651c03cfb78e355284540e538d293adc651c0f3699a22542845408ff3b73eaec651c0e261f284542845406ef831e3adc651c07cdf57a055284540",
		},
	} {
		if tc.ndr != "" {
			if got, err := Encode(tc.g, wkb.NDR); err != nil || got != tc.ndr {
				t.Errorf("Encode(%#v, %#v) == %#v, %#v, want %#v, nil", tc.g, wkb.NDR, got, err, tc.ndr)
			}
			if got, err := Decode(tc.ndr); err != nil || !reflect.DeepEqual(got, tc.g) {
				t.Errorf("Decode(%#v) == %#v, %v, want %#v, nil", tc.ndr, got, err, tc.g)
			}
		}
		if tc.xdr != "" {
			if got, err := Encode(tc.g, wkb.XDR); err != nil || got != tc.xdr {
				t.Errorf("Encode(%#v, %#v) == %#v, %#v, want %#v, nil", tc.g, wkb.XDR, got, err, tc.xdr)
			}
			if got, err := Decode(tc.xdr); err != nil || !reflect.DeepEqual(got, tc.g) {
				t.Errorf("Decode(%#v) == %#v, %v, want %#v, nil", tc.xdr, got, err, tc.g)
			}
		}
	}
}
