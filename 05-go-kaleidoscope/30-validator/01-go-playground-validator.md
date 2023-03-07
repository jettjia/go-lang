# 一、简介

在上一篇文章“[关于go-playground/validator源码分析](http://www.findme.wang/blog/detail/id/592.html)”，大致过了一遍validator源码，了解了基于tag中规则验证数据的原理，本文我们将从使用角度说说validator，如何使用tag规则来验证数据。首先，我们来看看validator提供哪些规则，如下，针对常用的，做了一些备注说明

```
"required":             hasValue, // 是否必填，即不能为零值
"required_with":        requiredWith, // 关联矫正，如required_with=Name 若Name字段不为零值，则当前字段不能为零值
"required_with_all":    requiredWithAll,// 多个字段都不为空的时候，才会矫正当前字段
"required_without":     requiredWithout, // 当某个字段为零值的时候，才会矫正当前字段
"required_without_all": requiredWithoutAll,// 当多个字段都为零值的时候，才会矫正当前字段
"isdefault":            isDefault, // 当前字段是否为零值
"len":                  hasLengthOf, // 针对数组、切片、map、string，判断长度是否为len对应的值，针对int、float、uint等，判断当前字段值是否为len对应的值
"min":                  hasMinOf, //用于验证当前字段的值是否大于或等于参数的值，若当前字段为数组、切片、map，则将它们的长度用于比较
"max":                  hasMaxOf,//用于验证当前字段的值是否大于或等于参数的值，若当前字段为数组、切片、map，则将它们的长度用于比较
"eq":                   isEq,// 判断是否相当,若当前字段为数组、切片、map，则将它们的长度用于比较
"ne":                   isNe,// 判断是否不等于
"lt":                   isLt, // 判断是否小于
"lte":                  isLte,// 判断是否小于等于
"gt":                   isGt,// 判断是否大于
"gte":                  isGte,// 判断是否大于等于
"eqfield":              isEqField, // 和另外一个字段进行比较，判断是否相当，可用于比较注册时，密码和确认密码是否相等
"eqcsfield":            isEqCrossStructField,
"necsfield":            isNeCrossStructField,
"gtcsfield":            isGtCrossStructField,
"gtecsfield":           isGteCrossStructField,
"ltcsfield":            isLtCrossStructField,
"ltecsfield":           isLteCrossStructField,
"nefield":              isNeField,
"gtefield":             isGteField,
"gtfield":              isGtField,
"ltefield":             isLteField,
"ltfield":              isLtField,
"fieldcontains":        fieldContains,
"fieldexcludes":        fieldExcludes,
"alpha":                isAlpha,
"alphanum":             isAlphanum,
"alphaunicode":         isAlphaUnicode,
"alphanumunicode":      isAlphanumUnicode,
"numeric":              isNumeric,
"number":               isNumber,
"hexadecimal":          isHexadecimal,
"hexcolor":             isHEXColor,
"rgb":                  isRGB,
"rgba":                 isRGBA,
"hsl":                  isHSL,
"hsla":                 isHSLA,
"e164":                 isE164,
"email":                isEmail,
"url":                  isURL,
"uri":                  isURI,
"urn_rfc2141":          isUrnRFC2141, // RFC 2141
"file":                 isFile,
"base64":               isBase64,
"base64url":            isBase64URL,
"contains":             contains,
"containsany":          containsAny,
"containsrune":         containsRune,
"excludes":             excludes,
"excludesall":          excludesAll,
"excludesrune":         excludesRune,
"startswith":           startsWith,
"endswith":             endsWith,
"startsnotwith":        startsNotWith,
"endsnotwith":          endsNotWith,
"isbn":                 isISBN,
"isbn10":               isISBN10,
"isbn13":               isISBN13,
"eth_addr":             isEthereumAddress,
"btc_addr":             isBitcoinAddress,
"btc_addr_bech32":      isBitcoinBech32Address,
"uuid":                 isUUID,
"uuid3":                isUUID3,
"uuid4":                isUUID4,
"uuid5":                isUUID5,
"uuid_rfc4122":         isUUIDRFC4122,
"uuid3_rfc4122":        isUUID3RFC4122,
"uuid4_rfc4122":        isUUID4RFC4122,
"uuid5_rfc4122":        isUUID5RFC4122,
"ascii":                isASCII,
"printascii":           isPrintableASCII,
"multibyte":            hasMultiByteCharacter,
"datauri":              isDataURI,
"latitude":             isLatitude,
"longitude":            isLongitude,
"ssn":                  isSSN,
"ipv4":                 isIPv4,
"ipv6":                 isIPv6,
"ip":                   isIP,
"cidrv4":               isCIDRv4,
"cidrv6":               isCIDRv6,
"cidr":                 isCIDR,
"tcp4_addr":            isTCP4AddrResolvable,
"tcp6_addr":            isTCP6AddrResolvable,
"tcp_addr":             isTCPAddrResolvable,
"udp4_addr":            isUDP4AddrResolvable,
"udp6_addr":            isUDP6AddrResolvable,
"udp_addr":             isUDPAddrResolvable,
"ip4_addr":             isIP4AddrResolvable,
"ip6_addr":             isIP6AddrResolvable,
"ip_addr":              isIPAddrResolvable,
"unix_addr":            isUnixAddrResolvable,
"mac":                  isMAC,
"hostname":             isHostnameRFC952,  // RFC 952
"hostname_rfc1123":     isHostnameRFC1123, // RFC 1123
"fqdn":                 isFQDN,
"unique":               isUnique, 用来矫正切片、数组元素是否全局唯一
"oneof":                isOneOf,// 如`validate:"oneof=2 3 4"`，限制当前字段值只能为2或3或4
"html":                 isHTML,
"html_encoded":         isHTMLEncoded,
"url_encoded":          isURLEncoded,
"dir":                  isDir,
"json":                 isJSON, // 判断当前值是否为有效的json字符串
"hostname_port":        isHostnamePort,
"lowercase":            isLowercase,
"uppercase":            isUppercase,
"datetime":             isDatetime, // 时间矫正,如 `validate:"datetime=2006-01-02"`
```

# 二、简单的结构体验证

```go
type Student struct {
   Name string `validate:"required,min=2,max=10"` // 必填，最大长度 10 最小长度10
   Age  int64  `validate:"min=6,max=15"` // 最大值为15，最小值为6
   Num  string `validate:"required,len=6"` // 必填,长度必须为6
}
 
s := Student{
   Age:3,
   Num:"no9527",
   Name:"zhangsan",
}
 
validate := validator.New() // 创建验证器
err := validate.Struct(s) // 执行验证
 
if err != nil {
   fmt.Println(err) // 执行结果 Key: 'Student.Age' Error:Field validation for 'Age' failed on the 'min' tag
}
```

# 三、自定义tag验证

有的时候，我们需要做一些自定义的业务验证，此时，我们就需要扩展tag验证函数。如下：

```go
// CreateSysMenuReq 创建SysMenu 请求对象
type CreateSysMenuReq struct {
		MenuName    string `validate:"required,min=1,max=32" err_info:"长度在1-32个字符" json:"menu_name"` // menu名称
		Desc        string `validate:"required,min=1,max=32" json:"desc"`                            // 描述
		Route       string ` json:"route"`                                                           // 菜单路由
		State       uint   ` json:"state"`                                                           // 1显示,2否
		Pid         uint64 ` json:"pid"`                                                             // 父id
	}
```

```go
package validate

import (
	"reflect"
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/gogf/gf/v2/text/gstr"
)

var Validate *validator.Validate

// 初始化Validate/v10国际化
func init() {
	Validate = validator.New()
    // 自定义校验规则
	if err := Validate.RegisterValidation("specialChar", CheckSpecialChar); err != nil {
		panic(err)
	}
}

// ValidatedProcessErr 处理参数错误的显示
func ValidatedProcessErr(u interface{}, err error) string {
	if err == nil {
		return ""
	}
	invalid, ok := err.(*validator.InvalidValidationError)
	if ok {
		return "您输入的参数错误有误：" + invalid.Error()
	}
	validationErrs := err.(validator.ValidationErrors) //断言是ValidationErrors
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field()                    //获取是哪个字段不符合格式
		field, ok := reflect.TypeOf(u).FieldByName(fieldName) //通过反射获取filed
		if ok {
			errorInfo := field.Tag.Get("err_info") //获取field对应的reg_error_info tag值
			if errorInfo == "" {
				return err.Error()
			}
			return gstr.CaseSnake(fieldName) + ":" + errorInfo //返回错误
		} else {
			return "缺失err_info"
		}
	}
	return ""
}

// CheckSpecialChar 校验特殊字符
func CheckSpecialChar(f validator.FieldLevel) bool {
	value := f.Field().String()
	if value == "" {
		return true
	}

	flag, err := regexp.MatchString("^([A-Za-z0-9)$", value)

	if err != nil {
		return false
	}

	return flag
}
```

```go
	// 参数过滤
	err = validate.Validate.Struct(&dtoReq)
	if err != nil {
		err = gerror.NewCode(responseutil.CommBadRequest, validate.ValidatedProcessErr(dtoReq, err))
		_ = c.Error(err)
		return
	}
```



# 四、国际化支持 

在上面的两个案例中，验证失败的时候，报错为“Key: 'Student.Age' Error:Field validation for 'Age' failed on the 'min' tag”和“Key: 'Student.Name' Error:Field validation for 'Name' failed on the 'unique_name' tag”，这种报错是很不友好的，我们需要中文。如下：

```go
package main
 
import (
   "fmt"
   "github.com/go-playground/locales/en"
   "github.com/go-playground/locales/zh"
   "github.com/go-playground/universal-translator"
   "github.com/go-playground/validator/v10"
   en_trans "github.com/go-playground/validator/v10/translations/en"
   zh_trans "github.com/go-playground/validator/v10/translations/zh"
)
 
type Student struct {
   Name string `validate:"required,min=2,max=10,unique_name"` // 必填，最大长度 10 最小长度10
   Age  int64  `validate:"min=6,max=15"`                      // 最大值为15，最小值为6
   Num  string `validate:"required,len=6"`
}
 
func main() {
   s := Student{
      Age:  2,
      Num:  "no9527",
      Name: "zhangsan",
   }
 
   // 创建翻译器
   zhTrans := zh.New() // 中文转换器
   enTrans := en.New() // 因为转换器
 
   uni := ut.New(zhTrans, zhTrans, enTrans) // 创建一个通用转换器
 
   curLocales := "zh"                        // 设置当前语言类型
   trans, _ := uni.GetTranslator(curLocales) // 获取对应语言的转换器
 
   validate := validator.New()                                     // 创建验证器
   _ = validate.RegisterValidation("unique_name", CheckUniqueName) // 注册自定义tag回调函数
 
   switch curLocales {
   case "zh":
      // 内置tag注册 中文翻译器
      _ = zh_trans.RegisterDefaultTranslations(validate, trans)
   case "en":
     // 内置tag注册 英文翻译器
      _ = en_trans.RegisterDefaultTranslations(validate, trans)
   }
 
   err := validate.Struct(s)
 
   if err != nil {
      errs := err.(validator.ValidationErrors)
      for _, e := range errs {
         // can translate each error one at a time.
         fmt.Println(e.Translate(trans))
      }
   }
}
 
// 矫正名字的唯一性
// fl 包含了字段相关所有信息
func CheckUniqueName(fl validator.FieldLevel) bool {
   // 获取字段当前值 fl.Field()
   // 获取tag 对应的参数 fl.Param() ，针对unique_name标签 ，不需要参数
   // 获取字段名称 fl.FieldName()
 
   // balabala处理一波，比如查库比较
 
   return false
}
```

执行结果如下：

![Xnip2020-07-19_22-55-48.jpg](http://www.findme.wang/Uploads/Editor/2020-07-19/5f145f020304b.jpg)

可以看出针对validator默认的tag验证规则，已翻译为中文了，但是针对自定义tag验证，仍然是英文的。怎么解决呢？

这个时候，就需要注册validator翻译器了，如下：

```go
switch curLocales {
case "zh":
   // 内置tag注册 中文翻译器
   _ = zh_trans.RegisterDefaultTranslations(validate, trans)
 
   // 自定义tag注册 中文翻译器
   _ = validate.RegisterTranslation("unique_name", trans, func(ut ut.Translator) error {
      if err := ut.Add("unique_name", "{0}已被占用", false); err != nil {
         return err
      }
 
      return nil
   }, func(ut ut.Translator, fe validator.FieldError) string {
      t, err := ut.T(fe.Tag(), fe.Field())
      if err != nil {
         log.Printf("警告: 翻译字段错误: %#v", fe)
         return fe.(error).Error()
      }
 
      return t
   })
case "en":
   // 内置tag注册 英文翻译器
   _ = en_trans.RegisterDefaultTranslations(validate, trans)
}
```

执行结果如下：

![Xnip2020-07-20_11-16-34.jpg](http://www.findme.wang/Uploads/Editor/2020-07-20/5f150fbc849fb.jpg)

看着还不错，但是“Name已被占用”和“Age最小只能为6”中，“Name”和“Age”还没翻译呢？

这个时候，我们可以注册RegisterTagNameFunc，如下：

```go
// 在tag中设置 label，代表中文翻译
type Student struct {
   Name string `validate:"required,min=2,max=10,unique_name" label:"姓名"` // 必填，最大长度 10 最小长度10
   Age  int64  `validate:"min=6,max=15" label:"年龄"`                      // 最大值为15，最小值为6
   Num  string `validate:"required,len=6" label:"序号"`
}
 
// 注册 RegisterTagNameFunc
validate.RegisterTagNameFunc(func(field reflect.StructField) string {
   name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
   if name == "-" {
      return ""
   }
 
   return name
})
```

执行结果如下：

![Xnip2020-07-20_14-48-25.jpg](http://www.findme.wang/Uploads/Editor/2020-07-20/5f153e4477a42.jpg)





# 常用校验

```
-  忽略
|  或
omitempty 有则验证，空值则不验证
dive  潜入到切片、数组、映射中，例如 NumList []int `validate:"len=2,dive,gt=18"` //切片长度为2，潜入切片后，里面的成员必须大于18
required 、 required_with[_all]、 required_without[_all]   //required_with表示指定字段有值，则本字段必须有值；required_without表示指定字段没有值时，本字段必须有值；指定字段有值，则本字段有值无值都可以。
len  数字时等效于eq, 字符串时等效字符串长度(是rune长度，比如"世界"或"sj"都满足len=2的约束)，切片或映射的话是元素的个数
max、min
eq、ne、{l|g}t[e]
oneof  例如 oneof=male female
alpha alphanum numeric hexadecimal

file  文件是否存在
contains(包含某个字符) containsany(包含任意一个给定字符串中的字符)  containsrune  excludes excludesall(不包含给定字符串中的所有字符) excludesrune
ip ipv4 ipv6 cidr cidrv{4|6} {tc|ud}p[{4|6}]_addr(就是多了端口号验证) fqdn eth_addr   可解析的IP地址(测试发现没觉得和ipv4、ipv6、ip标签有什么不同)：ip[{4|6}]_addr mac 
email url uri base64 base64url(因为base64的+和/在url中有特殊意义)  uuid  uuid3 uuid4 uuid5

startswith( 以什么开始)  endswith(以什么结束)   v9版本及其之后才能支持
ascii asciiprint multibyte(多字节，比如汉字，注意：如果是空，校验也能通过)
```



# 结构体嵌套校验

required,dive

```go
// 请求对象
type (
	// CreateCappMetaSysServiceReq 创建CappMetaSysService 请求对象
	CreateCappMetaSysServiceReq struct {
		Code         string          `validate:"specialCharacter"  json:"code"`         // 编号
		Name         string          `validate:"required,min=1,max=128" json:"name"`    // 描述
		State        uint            `validate:"required,min=1" json:"state"`           // 状态
		TagId        uint64          `validate:"required,min=1" json:"tag_id"`          // 标签id
		ItemId       uint64          `validate:"required,min=1" json:"item_id"`         // 编目id
		BizId        uint64          `validate:"required,min=1" json:"biz_id"`          // 业务组id
		AsMServiceId uint64          `validate:"required,min=1" json:"as_m_service_id"` // model服务service_id
		Filter       []ServiceFilter `validate:"required,dive" json:"filter"`           // filter
    }
    
    ServiceFilter struct {
		AsModelId   int64  `validate:"required,min=1" json:"as_model_id"`     // model服务中的字段id
		Code        string `validate:"required,min=1,max=128" json:"code"`    // 字段编号
		Name        string `validate:"required,min=1,max=128" json:"name"`    // 字段描述
		IsItem      int    `validate:"required,min=1" json:"is_item"`         // 关联编目字段:1是;2否
		IsSearch    int    `validate:"required,min=1" json:"is_search"`       // 搜索字段:1是;2否
		CodeType    int    `validate:"required,min=1,max=3" json:"code_type"` // 搜索类型:1文本框;2区间;3时间
		AsItemField string `validate:"required,min=1" json:"as_item_field"`   // as模板field
	}
 }
```

