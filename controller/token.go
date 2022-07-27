package controller

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	// "io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func GetToken(c *gin.Context) {
	// empty interface for result
	var res map[string]interface{}

	// validate required parameters
	required := []string{"code", "shop_id"}

	// query parameters
	queries := c.Request.URL.Query()

	for _, v := range required {
		// check if param exist
		if _, ok := queries[v]; !ok {
			c.JSON(http.StatusBadRequest, gin.H{"message" : "code and shop_id is required"})

			return
		}
	}

	partnerId := os.Getenv("PARTNER_ID")
	path := fmt.Sprintf("%sauth/token/get", os.Getenv("VERSION_PATH"))
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	baseString := fmt.Sprintf("%s%s%s%s", partnerId, path, timestamp, c.Query("shop_id"))

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(os.Getenv("PARTNER_KEY")))

	// Write Data to it
	h.Write([]byte(baseString))
	
	// set query parameters
	params := url.Values{
		"timestamp": {timestamp},
		"shop_id": {c.Query("shop_id")},
		"partner_id": {partnerId},
		"sign": {hex.EncodeToString(h.Sum(nil))},
	}

	// Encode into URL encode form
	postUrl := fmt.Sprintf("%s%s?%s", os.Getenv("API_URL"), path, params.Encode())

	// prepare json data
	intPartnerId, _ := strconv.Atoi(partnerId)
	intShopId, _ := strconv.Atoi(c.Query("shop_id"))

	data, _ := json.Marshal(map[string]interface{}{
		"code": c.Query("code"),
		"partner_id": intPartnerId,
		"shop_id": intShopId,
	})

	// send request
	response, error := http.Post(postUrl, "application/json", bytes.NewBuffer(data))

	if error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": error,
		})
	}

	defer response.Body.Close()

	json.NewDecoder(response.Body).Decode(&res)

	c.JSON(http.StatusOK, res)
}