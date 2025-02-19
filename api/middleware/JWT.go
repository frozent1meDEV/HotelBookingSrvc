package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func JWTAuth(c *fiber.Ctx) error {
	fmt.Println("-- JWT Auth --")

	tokenHeader := c.GetReqHeaders()["X-Api-Token"]
	if len(tokenHeader) == 0 {
		return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized: Missing X-Api-Token header")
	}

	// Extract the first token from the slice
	token := tokenHeader[0]

	if err := parseToken(token); err != nil {
		return err
	}

	fmt.Println("token: ", token)
	return nil
}

func parseToken(tokenStr string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			fmt.Println("Invalid signing method", token.Header["alg"])
			return nil, fmt.Errorf("unauthorized")
		}
		secret := os.Getenv("JWT_SECRET")
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(secret), nil
	})

	if err != nil {
		fmt.Println("failed to parse JWT Token:", err)
		return fmt.Errorf("unauthorized")
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)
	}
	return fmt.Errorf("unauthorized")
}
