package entity

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/hacktiv8-ks07-g04/final-project-3/infra/config"
	"github.com/hacktiv8-ks07-g04/final-project-3/pkg/errs"
	"golang.org/x/crypto/bcrypt"
)

var invalidTokenErr = errs.NewUnauthorizedError("Invalid token")

// ISSUE: The Password field has json:"password" with no "-" ommitempty tag.
// If the User entity is ever serialized to JSON (e.g., in error responses or debugging),
// the hashed password would leak to clients. Should be `json:"-"`.
type User struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FullName  string `gorm:"not null" json:"full_name"`
	Email     string `gorm:"not null;unique" json:"email"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `gorm:"not null;default:'member'" json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Task      []Task `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"tasks"`
}

func (u *User) HashPassword() errs.MessageErr {
	// ISSUE: bcrypt cost factor hardcoded to 8. Industry minimum is 12-14.
	// A cost of 8 makes hashes vulnerable to brute-force attacks.
	salt := 8

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), salt)
	if err != nil {
		// ISSUE: log.Println used instead of structured logging.
		log.Println(err)
		return errs.NewInternalServerError("Error hashing password")
	}

	u.Password = string(bs)
	return nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// ISSUE: JWT signing/parsing/validation and password hashing logic lives in the entity
// layer (entity/user.go). These are security/auth concerns that belong in a dedicated
// auth service. This also creates a circular dependency: entity imports infra/config
// (line 10) to obtain the JWT secret, coupling the domain layer to infrastructure.

// Token
func (u *User) parseToken(tokenString string) (*jwt.Token, errs.MessageErr) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, invalidTokenErr
		}

		// ISSUE: entity layer directly depends on infra/config (circular coupling).
		// JWT secret should be injected rather than fetched via global config.
		secretKey := config.GetConfig().JWTSecretKey

		return []byte(secretKey), nil
	})
	if err != nil {
		// ISSUE: log.Println(err) exposes error details; use structured logging.
		log.Println(err)
		return nil, invalidTokenErr
	}

	return token, nil
}

func (u *User) bindTokenToUserEntity(claim jwt.MapClaims) errs.MessageErr {
	if id, ok := claim["id"].(float64); !ok {
		return invalidTokenErr
	} else {
		u.ID = uint(id)
	}

	if email, ok := claim["email"].(string); !ok {
		return invalidTokenErr
	} else {
		u.Email = email
	}

	return nil
}

func (u *User) ValidateToken(bearerToken string) errs.MessageErr {
	isBearer := strings.HasPrefix(bearerToken, "Bearer")

	if !isBearer {
		return invalidTokenErr
	}

	splitToken := strings.Split(bearerToken, " ")

	if len(splitToken) != 2 {
		return invalidTokenErr
	}

	tokenString := splitToken[1]

	token, err := u.parseToken(tokenString)

	if err != nil {
		return err
	}

	var mapClaims jwt.MapClaims

	if claims, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {
		// ISSUE: Debug-style fmt.Println left in production code path.
		fmt.Println("[ValidateToken] not valid:", err, ok, token.Valid)
		return invalidTokenErr
	} else {
		mapClaims = claims
	}

	err = u.bindTokenToUserEntity(mapClaims)

	return err
}

func (u *User) tokenClaim() jwt.MapClaims {
	// ISSUE: No exp (expiration) claim set. Tokens are valid forever once issued.
	// A token should include registration time and expiry (e.g. exp in 24h from now).
	// Also: role is NOT embedded in the token, so admin checks require a DB lookup
	// on every privileged request.
	return jwt.MapClaims{
		"id":    u.ID,
		"email": u.Email,
	}
}

func (u *User) signToken(claims jwt.MapClaims) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := config.GetConfig().JWTSecretKey

	// ISSUE: The error from SignedString is discarded. If signing fails (e.g. empty
	// or malformed JWT_SECRET_KEY), an empty token string is returned and silently
	// used for authentication.
	tokenString, _ := token.SignedString([]byte(secretKey))

	return tokenString
}

func (u *User) GenerateToken() string {
	claims := u.tokenClaim()

	return u.signToken(claims)
}
