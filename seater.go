package seater

import (
	"errors"

	"github.com/attilasatan/seater/plane"
	"github.com/kataras/iris"
)

var Store map[string]*plane.Plane

func Init(planes ...plane.PlanePrototype) (err error) {

	Store = make(map[string]*plane.Plane)

	for _, pp := range planes {
		if _, exists := Store[pp.ID]; exists {
			return errors.New("plane id collision")
		}

		newPlane, err := plane.New(pp)
		if err != nil {
			return err
		}
		Store[pp.ID] = &newPlane
	}
	serve()
	return
}

func GetPlane(id string) *plane.Plane {
	if p, exists := Store[id]; exists {
		return p
	}
	return nil
}

func serve() {
	app := iris.Default()
	app.Get("/plane/{id:string}/{index:int}", func(ctx iris.Context) {
		id := ctx.Params().Get("id")
		index, _ := ctx.Params().GetInt("index")

		if p := GetPlane(id); p != nil {
			ctx.JSON(p.SeatByIndex(index))
		}
	})
	app.Run(iris.Addr(":8080"))
}
