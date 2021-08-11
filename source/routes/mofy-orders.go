package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"yam-api/source/utils"
	"yam-api/source/utils/mongodb"

	"github.com/Jeffail/gabs"
	"github.com/go-chi/chi"
)

func SetMofyOrders(w http.ResponseWriter, r *http.Request) {
	nftid := chi.URLParam(r, "nftid")
	fmt.Println(nftid)

	order, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	orderObject, err := gabs.ParseJSON(order)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(orderObject)

	/// @dev Store order in db
	status := mongodb.InsertMofyOrder(strings.ToLower(nftid), orderObject.Data())
	if status == false {
		fmt.Errorf("Failed saving the order in db")
	} else {
		utils.ResJSON(http.StatusCreated, w,
			status,
		)
	}
}

func GetMofyOrders(w http.ResponseWriter, r *http.Request) {
	nftid := chi.URLParam(r, "nftid")
	fmt.Println(nftid)

	/// @dev Retrieve values from db
	values := mongodb.GetMofyOrders(strings.ToLower(nftid))
	if values == nil {
		values = map[string]interface{}{
			"nftid": strings.ToLower(nftid),
			"order": "{}",
		}
	}

	utils.ResJSON(http.StatusCreated, w,
		values,
	)

	if values == nil {
		fmt.Errorf("No db entry for the given nft id.")
	}
}
