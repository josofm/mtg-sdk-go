package mtg

// import (
// 	"errors"
// 	"testing"

// 	. "github.com/smartystreets/goconvey/convey"
// 	"gopkg.in/jarcoal/httpmock.v1"
// )

// func Test_GenerateBooster(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()

// 	Convey("When generating a booster", t, func() {
// 		Convey("If the response is correct", func() {
// 			httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/PLS/booster",

// 			cards, err := SetCode("PLS").GenerateBooster()
// 			Convey("There should be no error", func() {
// 				So(err, ShouldBeNil)
// 			})
// 			Convey("The result should contain some cards", func() {
// 				So(cards, ShouldContainCard, "Planeswalker's Fury")
// 				So(cards, ShouldContainCard, "Strafe")
// 				So(cards, ShouldContainCard, "Crosis's Catacombs")
// 				So(cards, ShouldContainCard, "Warped Devotion")
// 				So(cards, ShouldContainCard, "Heroic Defiance")
// 				So(cards, ShouldContainCard, "Stormscape Familiar")
// 				So(cards, ShouldContainCard, "Honorable Scout")
// 			})
// 		})
// 	})
// }

// func Test_BoosterContentString(t *testing.T) {
// 	Convey("When converting a BoosterContent to a string", t, func() {
// 		Convey("A single type should be the type itself", func() {
// 			bc := BoosterContent{"Common"}
// 			So(bc.String(), ShouldEqual, "Common")
// 		})
// 		Convey("If there are multiple possible types they should be split by |", func() {
// 			bc := BoosterContent{"Common", "Rare"}
// 			So(bc.String(), ShouldEqual, "Common|Rare")
// 		})
// 	})
// }

// func Test_SetQuery(t *testing.T) {
// 	httpmock.Activate()
// 	defer httpmock.DeactivateAndReset()

// 	Convey("With a new SetQuery", t, func() {
// 		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=Planeshift&page=1&pageSize=500",
// 			NewStringResponderWithHeader(200, `{"sets":[{"code":"PLS","name":"Planeshift","type":"expansion","border":"black","booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"releaseDate":"2001-02-05","gathererCode":"PS","magicCardsInfoCode":"ps","block":"Invasion"}]}`,
// 				map[string]string{
// 					"Total-Count": "1337",
// 				}))
// 		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/PLS",
// 			httpmock.NewStringResponder(200, `{"set":{"code":"PLS","name":"Planeshift","type":"expansion","border":"black","booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"releaseDate":"2001-02-05","gathererCode":"PS","magicCardsInfoCode":"ps","block":"Invasion"}}`))
// 		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/FOO_BAR",
// 			httpmock.NewStringResponder(200, `{"sets":[]}`))
// 		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/network_issue",
// 			httpmock.NewErrorResponder(errors.New("Network Issue")))
// 		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/server_issue",
// 			httpmock.NewStringResponder(500, `{"status": "500", "error":"Internal server error"}`))
// 		httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets/invalid_json",
// 			httpmock.NewStringResponder(200, `{"sets":`))
// 		qry := NewSetQuery()

// 		Convey("When searching by name", func() {
// 			qry = qry.Where(SetName, "Planeshift")

// 			Convey("a copy should make no difference", func() {
// 				cpy := qry.Copy()
// 				So(cpy, ShouldResemble, qry)
// 				So(cpy, ShouldNotEqual, qry)
// 			})

// 			sets, totalCount, err := qry.Page(1)

// 			So(err, ShouldBeNil)
// 			So(totalCount, ShouldEqual, 1337)
// 			So(sets, ShouldHaveLength, 1)

// 			set := sets[0]
// 			So(set.Name, ShouldEqual, "Planeshift")
// 			So(set.String(), ShouldEqual, "Planeshift (PLS)")

// 			Convey("In case of errors", func() {
// 				Convey("Invalid json should be reported", func() {
// 					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=Planeshift&page=1&pageSize=500",
// 						httpmock.NewStringResponder(200, `{"sets":`))

// 					_, _, err := qry.Page(1)
// 					So(err, ShouldNotBeNil)
// 				})

// 				Convey("If Total-Count is not a number", func() {
// 					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=Planeshift&page=1&pageSize=500",
// 						NewStringResponderWithHeader(200, `{"sets":[{"code":"PLS","name":"Planeshift","type":"expansion","border":"black","booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"releaseDate":"2001-02-05","gathererCode":"PS","magicCardsInfoCode":"ps","block":"Invasion"}]}`,
// 							map[string]string{
// 								"Total-Count": "two",
// 							}))
// 					_, _, err := qry.Page(1)
// 					So(err, ShouldNotBeNil)
// 				})
// 			})

// 			Convey("fetching the same set by its id should result in the same values", func() {
// 				other, err := set.SetCode.Fetch()
// 				So(err, ShouldBeNil)
// 				So(other, ShouldResemble, set)
// 			})
// 			Convey("fetching an invalid setcode should return an error", func() {
// 				_, err := SetCode("FOO_BAR").Fetch()
// 				So(err, ShouldNotBeNil)
// 			})
// 			Convey("when we have network issues, there should also be an error", func() {
// 				_, err := SetCode("network_issue").Fetch()
// 				So(err, ShouldNotBeNil)
// 			})
// 			Convey("when the server reports an error we should get a ServerError", func() {
// 				_, err := SetCode("server_issue").Fetch()
// 				So(err, ShouldNotBeNil)
// 				_, isServerError := err.(ServerError)
// 				So(isServerError, ShouldBeTrue)
// 			})
// 			Convey("when the server sends invalid json there should be an error", func() {
// 				_, err := SetCode("invalid_json").Fetch()
// 				So(err, ShouldNotBeNil)
// 			})

// 			Convey("with paging", func() {
// 				qry = NewSetQuery().Where(SetName, "n")

// 				httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=n",
// 					NewStringResponderWithHeader(200, `{"sets":[{"code":"LEA","name":"Limited Edition Alpha","type":"core","border":"black","mkm_id":1,"booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"mkm_name":"Alpha","releaseDate":"1993-08-05","gathererCode":"1E","magicCardsInfoCode":"al"}]}`,
// 						map[string]string{
// 							"Link": `<https://api.magicthegathering.io/v1/sets?name=n&page=2>; rel="last", <https://api.magicthegathering.io/v1/sets?name=n&page=2>; rel="next"`,
// 						}))

// 				httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=n&page=2",
// 					httpmock.NewStringResponder(200, `{"sets":[{"code":"LEB","name":"Limited Edition Beta","type":"core","border":"black","mkm_id":2,"booster":["rare","uncommon","uncommon","uncommon","common","common","common","common","common","common","common","common","common","common","common"],"mkm_name":"Beta","releaseDate":"1993-10-01","gathererCode":"2E","magicCardsInfoCode":"be"}]}`))

// 				cards, err := qry.All()
// 				So(err, ShouldBeNil)
// 				So(cards, ShouldHaveLength, 2)

// 				Convey("If one of the following pages cause problems they should be reported", func() {
// 					httpmock.RegisterResponder("GET", "https://api.magicthegathering.io/v1/sets?name=n&page=2",
// 						httpmock.NewErrorResponder(errors.New("Network Issue")))
// 					_, err := qry.All()
// 					So(err, ShouldNotBeNil)
// 				})
// 			})
// 		})
// 	})
// }
