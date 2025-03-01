package data

import (
	"auth_manage/internal/biz"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type MyClaims struct {
	Username int64 `json:"user_id"` //自定义Payload有效载荷字段
	//Age      string `json:"age"`
	jwt.RegisteredClaims //提供标准验证功能,固定写法
}

// jwt加密的密钥
var key = "abcdefg123456"

type authRepo struct {
	data *Data
	log  *log.Helper
}

// NewAuthRepo .
func NewAuthRepo(data *Data, logger log.Logger) biz.AuthRepo {
	return &authRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *authRepo) SetToken(ctx context.Context, a *biz.Auth) (string, error) {
	if a.User_id == 0 {
		return "", errors.New("user_id不能为空")
	}
	SetClaims := MyClaims{
		Username: a.User_id,
		//Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Hour)), //有效时间(过期时间)，持续5个小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                    //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                    //生效时间，立即生效
			//Issuer:    os.Getenv("JWT_ISSUER"),                            //签发人
			//Subject:   "somebody",                                         //主题
			//ID:        "1",                                                //JWT ID用于标识该JWT
			//Audience:  []string{"somebody_else"},                          //用户
		},
	}
	//使用指定的加密方式(HS256签名算法)，和声明类型创建新令牌
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, SetClaims)
	//获得完整的、签名的令牌
	token, err := tokenStruct.SignedString([]byte(key)) //加密的密钥
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *authRepo) CheckToken(ctx context.Context, a *biz.Verfy) (bool, string, int64, error) {
	if a.Token == "" {
		return false, "jwt 内容为空", 0, nil
	}
	//解析、验证并返回token。
	tokenObj, err := jwt.ParseWithClaims(a.Token, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil //key 为自定义的密钥
	})
	if err != nil {
		return false, "jwt 解析错误", 0, nil
	}
	//类型断言获取Claims
	claims, ok := tokenObj.Claims.(*MyClaims)
	if !(ok && tokenObj.Valid) {
		return false, "jwt 鉴权失败", 0, nil
	}
	//信息鉴权成功
	//还需要判断token是否过期
	if time.Now().Unix() > claims.ExpiresAt.Unix() {
		return false, "jwt 时间过期", 0, nil
	}
	return true, "jwt 鉴权成功", claims.Username, nil
}

//实际操作（rdb，操作redis   |   jwt生成）
