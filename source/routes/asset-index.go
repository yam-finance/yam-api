package routes

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"yam-api/source/utils"
	"yam-api/source/utils/mongodb"

	"github.com/go-chi/chi"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
)

func FetchAssetIndex(sheetTitle string, cycleName string) map[string]interface{} {
	eval_ts := int(time.Now().UTC().Unix())

	privateKeyId := os.Getenv("PRIVATE_KEY_ID")
	privateKey := os.Getenv("PRIVATE_KEY")
	clientId := os.Getenv("CLIENT_ID")
	data := fmt.Sprintf(`{
		"type": "service_account",
		"project_id": "index-ustonks",
		"private_key_id": "%s",
		"private_key": "%s",
		"client_email": "yam-api@index-ustonks.iam.gserviceaccount.com",
		"client_id": "%s",
		"auth_uri": "https://accounts.google.com/o/oauth2/auth",
		"token_uri": "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url": "https://www.googleapis.com/robot/v1/metadata/x509/yam-api%%40index-ustonks.iam.gserviceaccount.com"
	}`, privateKeyId, privateKey, clientId)

	conf, err := google.JWTConfigFromJSON([]byte(data), spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheetId := os.Getenv("SHEET_ID")
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetId)
	checkError(err)
	sheet, err := spreadsheet.SheetByTitle(sheetTitle)
	checkError(err)

	price := sheet.Columns[12][49].Value

	values := map[string]interface{}{"cycle": cycleName, "price": price, "timestamp": strconv.Itoa(eval_ts)}

	return values
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

// --------- uSTONKS Endpoints ---------

func GetUStonksIndex(w http.ResponseWriter, r *http.Request) {
	cycle := chi.URLParam(r, "cycle")
	fmt.Println(cycle)

	/// @dev Retrieve values from db
	values := mongodb.GetLatestAssetIndex(strings.ToUpper(cycle))
	if values == nil {
		values = map[string]interface{}{
			"cycle":     "0",
			"price":     "0",
			"timestamp": "0",
		}
	}

	utils.ResJSON(http.StatusCreated, w,
		values,
	)

	if values == nil {
		fmt.Errorf("No db entry for uSTONKS index.")
	}
}

func GetUStonksIndexHistory(w http.ResponseWriter, r *http.Request) {
	cycle := chi.URLParam(r, "cycle")
	fmt.Println(cycle)

	/// @dev Retrieve values from db
	history := mongodb.GetAssetIndexHistoryDaily(strings.ToUpper(cycle))
	if history == nil {
		history = []map[string]interface{}{
			{
				"cycle":     "0",
				"price":     "0",
				"timestamp": "0",
			},
		}
	}

	utils.ResJSON(http.StatusCreated, w,
		history,
	)

	if history == nil {
		fmt.Errorf("No db entry for uSTONKS index.")
	}
}
