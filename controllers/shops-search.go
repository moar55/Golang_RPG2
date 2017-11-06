package controllers

import (
	"Golang_RPG/errors"
	"Golang_RPG/models"
	"context"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"googlemaps.github.io/maps"
)

type ShopsSearchController struct {
	beego.Controller
}

type Response struct {
	Message string `json:"message"`
}

func ChatSearch(latitude float64, longitude float64, c *ChatController) {
	c2, err := maps.NewClient(maps.WithAPIKey(beego.AppConfig.String("googlePlacesKey")))
	r := &maps.NearbySearchRequest{Location: &maps.LatLng{Lat: latitude, Lng: longitude}, RankBy: "distance", Type: "stadium"}
	resp, err := c2.NearbySearch(context.Background(), r)
	if err != nil {
		c.Data["json"] = &errors.InvalidParameters.Message
		fmt.Println("the err is ", err)
		c.Ctx.ResponseWriter.WriteHeader(errors.InvalidParameters.HTTPStatus)
	} else {
		nearestLocation := resp.Results[0]
		o := orm.NewOrm()
		location := models.Locations{Name: nearestLocation.Name}
		err := o.Read(&location, "Name")
		if err == orm.ErrNoRows {
			c.Data["json"] = &Response{Message: "No nearby shops"}
		} else {
			var (
				latitudeOfNearest   string = strconv.FormatFloat(nearestLocation.Geometry.Location.Lat, 'f', -1, 64)
				longtitudeOfNearest string = strconv.FormatFloat(nearestLocation.Geometry.Location.Lng, 'f', -1, 64)
				destination         string = latitudeOfNearest + "," + longtitudeOfNearest
				origin              string = strconv.FormatFloat(latitude, 'f', 6, 64) + "," + strconv.FormatFloat(longitude, 'f', 6, 64)
			)

			c2, err := maps.NewClient(maps.WithAPIKey(beego.AppConfig.String("googleDistanceKey")))
			r := &maps.DistanceMatrixRequest{Origins: []string{origin}, Destinations: []string{destination}}
			resp, err := c2.DistanceMatrix(context.Background(), r)
			if err != nil {
				c.Data["json"] = &errors.InvalidParameters.Message
				fmt.Println(err.Error())
				c.Ctx.ResponseWriter.WriteHeader(errors.InvalidParameters.HTTPStatus)
			} else {
				var distance int = resp.Rows[0].Elements[0].Distance.Meters
				if distance <= 200 && distance > 5 {
					var message string = fmt.Sprintf("A nearby shop is located at %d meters away, get closer to access it :)", distance)
					c.Data["json"] = &Response{Message: message}
				} else if distance <= 5 {
					c.Data["json"] = &Response{Message: "A nearby shop is just beside you. Type access to access it!"}
					c.SetSession("nearShop", location.Id)
				} else {
					c.Data["json"] = &Response{Message: "No nearby shops!"}

				}
			}
		}
		c.ServeJSON(true)
	}
}
