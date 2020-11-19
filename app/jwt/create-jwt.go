package jwt

import (
	"os"
	"time"

	"github.com/Tsuryu/arreglapp-user-profile/app/models"

	"github.com/dgrijalva/jwt-go"
)

// CreateJWT : creates a jwt to validate user's session
func CreateJWT(userProfile *models.UserProfile) (string, error) {
	jwtSecretKey := os.Getenv("JWT_SECRET_KEY")
	myKey := []byte(jwtSecretKey)

	payload := jwt.MapClaims{
		"username":   userProfile.Username,
		"first_name": userProfile.FirstName,
		"last_name":  userProfile.LastName,
		"email":      userProfile.Email,
		"phone":      userProfile.Phone,
		"exp":        time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	signedToken, err := token.SignedString(myKey)
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
