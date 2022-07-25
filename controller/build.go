package controller

import (
	"crypto/hmac"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Build(c *gin.Context) {
	partnerId := os.Getenv("PARTNER_ID")
	path := fmt.Sprintf("%sshop/auth_partner", os.Getenv("VERSION_PATH"))
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)

	baseString := fmt.Sprintf("%s%s%s", partnerId, path, timestamp)

	// Create a new HMAC by defining the hash type and the key (as byte array)
    h := hmac.New(sha256.New, []byte(os.Getenv("PARTNER_KEY")))

    // Write Data to it
    h.Write([]byte(baseString))
	
	params := url.Values{}
	params.Add("partner_id", partnerId)
	params.Add("redirect", fmt.Sprintf("http://%s/auth", os.Getenv("HOST")))
	params.Add("timestamp", timestamp)
	params.Add("sign", hex.EncodeToString(h.Sum(nil)))

	// Encode into URL encode form
	build := fmt.Sprintf("%s%s?%s", os.Getenv("API_URL"), path, params.Encode())

	c.Redirect(http.StatusFound, build)
}
