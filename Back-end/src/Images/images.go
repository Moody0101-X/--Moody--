package images;

import (
	// "fmt"
	"log"
    // "io/ioutil"
    "net/http"
    // "net/url"
    "encoding/json"
    "bytes"
    // "strings"
)

var AddImageRoute = "/products/img/add";
var Host = "http://localhost:8500";

/*
	Method: POST
	Expected data: { 
		"mime": ImageMime, "id": ProductID
	}
*/

func SaveImage(ProductId int, Mime string) (bool, string) {

    values := make(map[string]interface{})
    values["id"] = ProductId;
    values["mime"] = Mime;

    data, err := json.Marshal(values)

    resp, err := http.Post(Host + AddImageRoute, "application/json" , bytes.NewBuffer(data))

    if err != nil {
        log.Fatal(err)
    }

    var res map[string]interface{};

    json.NewDecoder(resp.Body).Decode(&res)
    
    if int(res["code"].(float64)) == 200 {
        return true, res["data"].(map[string]interface{})["url"].(string)
    }

    return false, res["data"].(string)
}


