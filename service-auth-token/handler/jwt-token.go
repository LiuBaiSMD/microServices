/*
auth:   wuxun
date:   2020-01-15 11:20
mail:   lbwuxun@qq.com
desc:   how to use or use for what
*/

package handler

import (
	"JWTToken/proto"
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
)

var SecretKey = "abcdefg"
type JwtTokenCreator struct {
}

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid string `json:"uid"`
	Name string `json:"name"`
}

func (t *JwtTokenCreator) GetToken(ctx context.Context, req *jwtToken.TokenRequest, rsp *jwtToken.TokenResponse)error{
	log.Print("Received TokenCreator.TokenRequest request")
	fmt.Println(req)
	name := req.Name
	id := req.Uid
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
		},
		id,
		name,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	fmt.Println("tokenString: ", tokenString, err)
	fmt.Println("token:       ", token)
	// set response status
	//rsp.StatusCode = 200

	// respond with some json


	// set json body
	rsp.Token = string(tokenString)

	return nil
}