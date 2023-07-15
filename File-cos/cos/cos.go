package cos

import (
	"File-cos/config"
	sts "github.com/tencentyun/qcloud-cos-sts-sdk/go"
	"log"
	"strconv"
	"time"
)

type Credential interface {
	GetCredential() (*CredentialResult, error)
}
type GetCredential struct {
}
type GetRoleCredential struct {
}

func NewGetCredential() Credential {
	return &GetCredential{}
}
func NewGetRoleCredential() Credential {
	return &GetRoleCredential{}
}

type CredentialResult struct {
	Credentials *Credentials `json:"credentials,omitempty"`
	ExpiredTime int64        `json:"expiredTime,omitempty"`
	Expiration  string       `json:"expiration,omitempty"`
	Bucket      string       `json:"bucket"`
	Region      string       `json:"region"`
}
type Credentials struct {
	TmpSecretID  string `json:"tmpSecretId,omitempty"`
	TmpSecretKey string `json:"tmpSecretKey,omitempty"`
	SessionToken string `json:"sessionToken,omitempty"`
}
type Role struct {
	RoleArn         string
	RoleSessionName string
}

func getClient() *sts.Client {

	c := sts.NewClient(
		// 通过环境变量获取密钥, os.Getenv 方法表示获取环境变量
		//os.Getenv("SECRETID"),  // 用户的 SecretId，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考https://cloud.tencent.com/document/product/598/37140
		//os.Getenv("SECRETKEY"), // 用户的 SecretKey，建议使用子账号密钥，授权遵循最小权限指引，降低使用风险。子账号密钥获取可参考https://cloud.tencent.com/document/product/598/37140
		config.Secret.GetString("TencentCloudSecret.SecretId"),
		config.Secret.GetString("TencentCloudSecret.SecretKey"),
		nil,
		sts.Host(config.Confs.GetString("CosConf.Host")),
		//sts.Host("sts.tencentcloudapi.com"), // 设置域名, 默认域名sts.tencentcloudapi.com
		sts.Scheme("https"), // 设置协议, 默认为https，公有云sts获取临时密钥不允许走http，特殊场景才需要设置http
	)
	return c
}
func getOptions(role ...*Role) *sts.CredentialOptions {
	appid := config.Confs.GetString("CosConf.Appid")
	bucket := config.Confs.GetString("CosConf.Bucket")
	region := config.Confs.GetString("CosConf.Region")
	DurationSeconds, _ := strconv.ParseInt(config.Confs.GetString("CosConf.DurationSeconds"), 10, 64)
	// 策略概述 https://cloud.tencent.com/document/product/436/18023
	if DurationSeconds == 0 {
		DurationSeconds = int64(time.Hour.Seconds())
	}
	opt := &sts.CredentialOptions{
		DurationSeconds: DurationSeconds,
		Region:          region,
		Policy: &sts.CredentialPolicy{
			Statement: []sts.CredentialPolicyStatement{
				{
					// 密钥的权限列表。简单上传和分片需要以下的权限，其他权限列表请看 https://cloud.tencent.com/document/product/436/31923
					Action: []string{
						// 简单上传
						"name/cos:PostObject",
						"name/cos:PutObject",
						// 分片上传
						"name/cos:InitiateMultipartUpload",
						"name/cos:ListMultipartUploads",
						"name/cos:ListParts",
						"name/cos:UploadPart",
						"name/cos:CompleteMultipartUpload",
					},
					Effect: "allow",
					Resource: []string{
						// 这里改成允许的路径前缀，可以根据自己网站的用户登录态判断允许上传的具体路径，例子： a.jpg 或者 a/* 或者 * (使用通配符*存在重大安全风险, 请谨慎评估使用)
						// 存储桶的命名格式为 BucketName-APPID，此处填写的 bucket 必须为此格式
						"qcs::cos:" + region + ":uid/" + appid + ":" + bucket + "/*",
					},
					// 开始构建生效条件 condition
					// 关于 condition 的详细设置规则和COS支持的condition类型可以参考https://cloud.tencent.com/document/product/436/71306
					//Condition: map[string]map[string]interface{}{
					//	"ip_equal": map[string]interface{}{
					//		"qcs:ip": []string{
					//			"10.217.182.3/24",
					//			"111.21.33.72/24",
					//		},
					//	},
					//},
				},
			},
		},
	}
	if len(role) > 0 {
		opt.RoleSessionName = role[0].RoleSessionName
		opt.RoleArn = role[0].RoleArn
	}
	return opt
}

func (*GetCredential) GetCredential() (*CredentialResult, error) {
	c := getClient()
	opt := getOptions()
	res, err := c.GetCredential(opt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	r := &CredentialResult{
		Expiration:  res.Expiration,
		ExpiredTime: int64(res.ExpiredTime),
		Credentials: &Credentials{
			TmpSecretID:  res.Credentials.TmpSecretID,
			TmpSecretKey: res.Credentials.TmpSecretKey,
			SessionToken: res.Credentials.SessionToken,
		},
		Bucket: config.Confs.GetString("CosConf.Bucket"),
		Region: config.Confs.GetString("CosConf.Region"),
	}
	return r, nil

}

// case 3 发起角色授权临时密钥请求, policy选填
func (*GetRoleCredential) GetCredential() (*CredentialResult, error) {
	roleArn := config.Confs.GetString("CosConf.RoleArn")
	roleSessionName := config.Confs.GetString("CosConf.RoleSessionName")
	role := &Role{
		RoleArn:         roleArn,
		RoleSessionName: roleSessionName,
	}
	c := getClient()
	opt := getOptions(role)
	res, err := c.GetRoleCredential(opt)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	r := &CredentialResult{
		Expiration:  res.Expiration,
		ExpiredTime: int64(res.ExpiredTime),
		Credentials: &Credentials{
			TmpSecretID:  res.Credentials.TmpSecretID,
			TmpSecretKey: res.Credentials.TmpSecretKey,
			SessionToken: res.Credentials.SessionToken,
		},
		Bucket: config.Confs.GetString("CosConf.Bucket"),
		Region: config.Confs.GetString("CosConf.Region"),
	}
	return r, nil

}
