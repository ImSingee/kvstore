package kvstore

import (
	"fmt"
	"testing"

	"github.com/ImSingee/tt"
)

var testData = `{
    "_id":"611e0ba6551aff2f434298d1",
    "index":0,
    "guid":"8674f065-3ee8-4b40-a699-2397eddf0012",
    "isActive":false,
    "balance":"$1,844.03",
    "picture":"http://placehold.it/32x32",
    "age":37,
    "eyeColor":"blue",
    "name":"Patton Nieves",
    "gender":"male",
    "company":"XYQAG",
    "email":"pattonnieves@xyqag.com",
    "phone":"+1 (837) 456-3131",
    "address":"172 Claver Place, Juntura, Alaska, 5942",
    "about":"Anim do nulla Lorem velit. Dolore excepteur non dolor Lorem amet ut occaecat irure. Duis velit sunt in sunt ex ex ut cupidatat aliquip in eu cupidatat elit occaecat. Labore ad culpa do id irure duis esse cupidatat velit labore ullamco non consectetur. Ipsum nulla adipisicing cupidatat tempor nostrud non. Quis voluptate eu anim occaecat.\r\n",
    "registered":"2018-11-14T08:30:10 -08:00",
    "latitude":11.655197,
    "longitude":37.981309,
    "tags":[
        "ad",
        "anim",
        "quis",
        "ex",
        "velit",
        "voluptate",
        "et"
    ],
    "friends":[
        {
            "id":0,
            "name":"Molly Castaneda"
        },
        {
            "id":1,
            "name":"Colleen Jackson"
        },
        {
            "id":2,
            "name":"Dunlap Wood"
        }
    ],
    "greeting":"Hello, Patton Nieves! You have 5 unread messages.",
    "favoriteFruit":"banana",
    "children":{
        "a":{
            "_id":"611f1102339781c5e04394bc",
            "index":0,
            "guid":"15dd697b-ceee-4203-b443-b3a436a75ffb",
            "isActive":true,
            "balance":"$3,929.48",
            "picture":"http://placehold.it/32x32",
            "age":26,
            "eyeColor":"brown",
            "name":"Winifred Buck",
            "gender":"female",
            "company":"FRANSCENE",
            "email":"winifredbuck@franscene.com",
            "phone":"+1 (862) 480-3732",
            "address":"218 Flatlands Avenue, Caron, Minnesota, 7622",
            "about":"Reprehenderit dolore irure proident deserunt enim excepteur pariatur irure adipisicing in. Proident sit cupidatat deserunt excepteur eu non sunt mollit id mollit cupidatat id. Do duis cupidatat mollit commodo officia qui. Duis excepteur magna id consequat aute cillum sint sunt ipsum. Incididunt in amet non deserunt velit elit eu laboris est est ut qui sunt nisi.\r\n",
            "registered":"2018-12-23T05:43:00 -08:00",
            "latitude":84.325495,
            "longitude":35.96704,
            "tags":[
                "magna",
                "occaecat",
                "dolor",
                "veniam",
                "laboris",
                "nisi",
                "mollit"
            ],
            "friends":[
                {
                    "id":0,
                    "name":"Eula Rutledge"
                },
                {
                    "id":1,
                    "name":"Minnie Mcdonald"
                },
                {
                    "id":2,
                    "name":"Rosalinda Dejesus"
                }
            ],
            "greeting":"Hello, Winifred Buck! You have 9 unread messages.",
            "favoriteFruit":"strawberry"
        },
        "b":{
            "_id":"611f11df8253a09bf7d325fd",
            "index":0,
            "guid":"13e03951-42ea-431b-a115-e04f586bc455",
            "isActive":false,
            "balance":"$3,976.62",
            "picture":"http://placehold.it/32x32",
            "age":26,
            "eyeColor":"green",
            "name":"Kelly Mcguire",
            "gender":"male",
            "company":"EQUICOM",
            "email":"kellymcguire@equicom.com",
            "phone":"+1 (962) 402-2212",
            "address":"775 Reed Street, Bentonville, South Carolina, 4235",
            "about":"Mollit commodo duis consectetur enim mollit id nulla aliqua laborum aliquip labore pariatur excepteur voluptate. Officia aliqua laboris sit est officia aute deserunt veniam nostrud duis nostrud. Proident minim mollit ex consectetur nulla mollit ea sit ea ullamco et minim sint. Est magna qui commodo ex elit quis do ullamco aliqua dolor ex commodo.\r\n",
            "registered":"2020-09-22T06:27:22 -08:00",
            "latitude":55.82411,
            "longitude":-49.206329,
            "tags":[
                "deserunt",
                "nulla",
                "non",
                "officia",
                "voluptate",
                "veniam",
                "id"
            ],
            "friends":[
                {
                    "id":0,
                    "name":"Hogan Mcbride"
                },
                {
                    "id":1,
                    "name":"Dona Davenport"
                },
                {
                    "id":2,
                    "name":"Alfreda Harrington"
                }
            ],
            "greeting":"Hello, Kelly Mcguire! You have 9 unread messages.",
            "favoriteFruit":"banana"
        }
    },
    "simple":{
        "eyeColor":"blue",
        "name":"Aida Sanchez",
        "gender":"female",
        "company":"PORTALIS",
        "email":"aidasanchez@portalis.com",
        "phone":"+1 (819) 424-2032",
        "address":"885 Merit Court, Marbury, District Of Columbia, 9139",
        "about":"Aliqua consectetur incididunt minim Lorem cupidatat velit eiusmod do. Do tempor culpa veniam amet elit ipsum adipisicing nostrud eiusmod excepteur consequat. Magna cupidatat ipsum culpa consectetur ex deserunt incididunt sit cupidatat velit sint incididunt ad. Incididunt labore laborum laborum sint dolor est adipisicing est laboris eiusmod culpa. Aliqua nostrud ipsum ea officia esse qui. Lorem do laborum exercitation enim mollit elit proident nisi irure. Elit esse quis laborum consectetur non nulla in.\r\n",
        "registered":"2021-05-19T12:58:47 -08:00",
        "latitude":-37.629836,
        "longitude":-59.529829
    },
    "complicate":{
        "name":"Hannah Mack",
        "gender":"female",
        "company":"VERTON",
        "tags":[
            "duis",
            "irure",
            "non",
            {
                "id":0,
                "name":"Naomi Burris"
            },
            {
                "isActive":false,
                "balance":"$2,422.40",
                "oops":[
                    "duis",
                    "irure"
                ]
            },
            "id"
        ]
    }
}`

func newStore() Store {
	s, err := LoadFromJSON([]byte(testData))
	if err != nil {
		panic(err)
	}

	return s
}

func TestStore_GetInt(t *testing.T) {
	cases := []struct {
		name      string
		in        string
		want      int64
		willError bool
	}{
		{
			name:      `not-exist`,
			in:        "xx",
			willError: true,
		},
		{
			name: `first-level`,
			in:   "age",
			want: 37,
		},
		{
			name: `nested`,
			in:   "friends.1.id",
			want: 1,
		},
		{
			name:      `array-index-invalid`,
			in:        "friends.-1.id",
			willError: true,
		},
		{
			name:      `array-index-overflow`,
			in:        "friends.99.id",
			willError: true,
		},
		{
			name:      `type-mismatch`,
			in:        "greeting",
			willError: true,
		},
		{
			name:      `middle-type-mismatch`,
			in:        "friends.a.xx",
			willError: true,
		},
		{
			name:      `attr-access-on-basic`,
			in:        "greeting.something",
			willError: true,
		},
	}

	store := newStore()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := store.GetInt(c.in)

			if !c.willError {
				tt.AssertIsNil(t, err)

				tt.AssertEqual(t, got, c.want)
			} else {
				tt.AssertIsNotNil(t, err)
				fmt.Println("Expected error:", err)
			}
		})
	}
}
