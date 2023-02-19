package auth

import (
    "crypto/sha256"
    b64 "encoding/base64"
    "fmt"
    "encoding/hex"
    "github.com/golang-jwt/jwt"
    "math/rand"
    "time"
    "strconv"
    "moody.com/api/models"
)

var JWT_SECRET_KEY []byte = []byte("21656756270094038278781082827545")

func HashPwd(s string) string {
    hash_ := sha256.New()
    hash_.Write([]byte(s))
    return hex.EncodeToString(hash_.Sum(nil))
}

func CheckHash(Pwd string, Hash_ string) bool { return (HashPwd(Pwd) == Hash_) }

// Makes a jwt that stores data to be sent to the client.
func StoreUserIdentityInJWT(User_id int, User_pwd string, User_email string) (string, error) {
 
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "id": User_id, 
        "pwd": User_pwd,
        "email": User_email,
    })

    tokenString, err := token.SignedString(JWT_SECRET_KEY)
    return tokenString, err
}

// Extracts the identity of the logging admin. from the token string.
func GetUserIdentityFromJWT(TokenStr string) (models.User, bool) { 
	var User models.User;

    token, err := jwt.Parse(TokenStr, func(token *jwt.Token) (interface{}, error) {
        
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }

        return JWT_SECRET_KEY, nil
    })

    if err != nil {
        fmt.Println("ERR: ", err)
        return User, false
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        id, _ := claims["id"].(int)
        pwd, _ := claims["pwd"].(string) 
        email, _ := claims["email"].(string)
       
        User.Id = id
        User.Pwd = pwd
        User.Email = email
        
        return User, true
    } else {

        return User, false
    }
}

func GenerateAccessToken(salt string) string {
  
    rand.Seed(time.Now().UnixNano())

    var SaltAsBytes []byte = []byte(salt)
    var _IV string = ""
    var threash_hold int = 255;
    var n int
    var nstring string

    for i := 0; i < 32; i++ {
        n = rand.Intn(threash_hold)
        nstring = strconv.Itoa(n)
        _IV += nstring
    }
    
    var final string = _IV + b64.StdEncoding.EncodeToString(SaltAsBytes);
    return HashPwd(final);
}