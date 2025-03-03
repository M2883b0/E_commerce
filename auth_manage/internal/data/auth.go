package data

import (
	"auth_manage/internal/biz"
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/redis/go-redis/v9"
	"strconv"
	"time"
)

type MyClaims struct {
	Username string `json:"user_id"` //自定义Payload有效载荷字段
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
	redisKey := fmt.Sprintf("session_id:%d", a.User_id)
	timeKey := fmt.Sprintf("session_time:%d", a.User_id)
	redisValue, err := r.data.rdb.Get(ctx, redisKey).Result()
	if err != nil && err != redis.Nil {
		return "", errors.New("session auth error")
	}
	if redisValue != "" {
		return redisValue, nil
	}

	SetClaims := MyClaims{
		Username: strconv.Itoa(int(a.User_id)),
		//Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			//ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Hour)), //有效时间(过期时间)，持续5个小时
			IssuedAt:  jwt.NewNumericDate(time.Now()), //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), //生效时间，立即生效
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
	if err := r.data.rdb.Set(ctx, redisKey, token, 2*time.Hour).Err(); err != nil {
		return "", err
	}
	if err := r.data.rdb.Set(ctx, timeKey, 1, 2*time.Hour).Err(); err != nil {
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
	timekey := fmt.Sprintf("session_time:%d", claims.Username)
	timeValue, err := r.data.rdb.Get(ctx, timekey).Result()
	temp, _ := strconv.Atoi(timeValue)
	if time.Now().Unix() > claims.IssuedAt.Add(time.Duration(temp)*time.Hour).Unix() {
		r.data.rdb.Set(ctx, timekey, temp+1, 2*time.Hour)
	}

	num, err := strconv.Atoi(claims.Username)
	if err != nil {
		return false, "jwt 转型失败", 0, nil
	}
	return true, "jwt 鉴权成功", int64(num), nil
}

//实际操作（rdb，操作redis   |   jwt生成）
