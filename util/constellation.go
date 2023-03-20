package util
import(
	// "fmt"
	"time"
	"github.com/astaxie/beego"
	"strings"
	"strconv"
	"math/rand"
	// "fmt"
	"encoding/json"
)

//y, m, d := utils.GetTimeFromStrDate("1983-10-07") 
//beego.Info("age:",utils.GetAge(y)," | 星座:",utils.GetConstellation(m,d)," | 生肖:",utils.GetZodiac(y))
func GetTimeFromStrDate(date string) (year, month, day int) { 
	// const shortForm = "2006-01-02" 
	d := strings.Split(date,"-")
	year,_ = strconv.Atoi(d[0])
	month,_ = strconv.Atoi(d[1])
	day,_ = strconv.Atoi(d[2])
	beego.Info("d:",d)
	beego.Info("year:",year," | month:",month," | day:",day," | date:",date)
	return 
} 

func GetZodiac(year int) (zodiac string) { 
	if year <= 0 { 
		zodiac = "-1" 
	} 
	start := 1901 
	x := (start - year) % 12 
	if x == 1 || x == -11 {
	 	zodiac = "鼠" 
	} 
	if x == 0 { 
		zodiac = "牛" 
	} 
	if x == 11 || x == -1 { 
		zodiac = "虎"
	} 
	if x == 10 || x == -2 { 
		zodiac = "兔" 	
	} 
	if x == 9 || x == -3 { 
		zodiac = "龙" 
	} 
	if x == 8 || x == -4 {
		zodiac = "蛇" 
	} 
	if x == 7 || x == -5 { 
		zodiac = "马" 
	} 
	if x == 6 || x == -6 { 
		zodiac = "羊" 
	} 
	if x == 5 || x == -7 { 
		zodiac = "猴" 
	} 
	if x == 4 || x == -8 { 
		zodiac = "鸡" 
	} 
	if x == 3 || x == -9 { 
		zodiac = "狗" 
	} 
	if x == 2 || x == -10 { 
		zodiac = "猪" 
	} 
	return 
} 

func GetAge(year int) (age int) { 
	if year <= 0 { 
		age = -1 
	} 
	nowyear := time.Now().Year() 
	age = nowyear - year 
	return 
}

func GetConstellation(month, day int) (star string) { 
	beego.Info("month:",month," | day:",day)
	if month <= 0 || month >= 13 { 
		beego.Info("month1:",month," | day:",day)
		star = "-1" 
	} 
	if day <= 0 || day >= 32 { 
		beego.Info("month2:",month," | day:",day)
		star = "-1" 
	} 
	if (month == 1 && day >= 20) || (month == 2 && day <= 18) { 
		star = "水瓶座" 
	} 
	if (month == 2 && day >= 19) || (month == 3 && day <= 20) { 
		star = "双鱼座" 
	} 
	if (month == 3 && day >= 21) || (month == 4 && day <= 19) { 
		star = "白羊座" 
	} 
	if (month == 4 && day >= 20) || (month == 5 && day <= 20) { 
		star = "金牛座" 
	} 
	if (month == 5 && day >= 21) || (month == 6 && day <= 21) {
	 	star = "双子座" 
	 } 
	 if (month == 6 && day >= 22) || (month == 7 && day <= 22) { 
	 	star = "巨蟹座" 
	 } 
	 if (month == 7 && day >= 23) || (month == 8 && day <= 22) { 
	 	star = "狮子座" 
	 } 
	 if (month == 8 && day >= 23) || (month == 9 && day <= 22) { 
	 	star = "处女座" 
	 } 
	 if (month == 9 && day >= 23) || (month == 10 && day <= 22) { 
	 	star = "天秤座" 
	 } 
	 if (month == 10 && day >= 23) || (month == 11 && day <= 21) { 
	 	star = "天蝎座" 
	 } 
	 if (month == 11 && day >= 22) || (month == 12 && day <= 21) { 
	 	star = "射手座" 
	 } 
	 if (month == 12 && day >= 22) || (month == 1 && day <= 19) { 
	 	star = "魔蝎座" 
	 }   
	 return star 
}

// 获取随机字符串
//    length：字符串长度
func GetRandomString(length int) string {
	str := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	var (
		result []byte
		b      []byte
		r      *rand.Rand
	)
	b = []byte(str)
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, b[r.Intn(len(b))])
	}
	return string(result)
}

func Strval(value interface{}) string {
	var key string
	if value == nil {
		return key
	}

	switch value.(type) {
	case float64:
		ft := value.(float64)
		key = strconv.FormatFloat(ft, 'f', -1, 64)
	case float32:
		ft := value.(float32)
		key = strconv.FormatFloat(float64(ft), 'f', -1, 64)
	case int:
		it := value.(int)
		key = strconv.Itoa(it)
	case uint:
		it := value.(uint)
		key = strconv.Itoa(int(it))
	case int8:
		it := value.(int8)
		key = strconv.Itoa(int(it))
	case uint8:
		it := value.(uint8)
		key = strconv.Itoa(int(it))
	case int16:
		it := value.(int16)
		key = strconv.Itoa(int(it))
	case uint16:
		it := value.(uint16)
		key = strconv.Itoa(int(it))
	case int32:
		it := value.(int32)
		key = strconv.Itoa(int(it))
	case uint32:
		it := value.(uint32)
		key = strconv.Itoa(int(it))
	case int64:
		it := value.(int64)
		key = strconv.FormatInt(it, 10)
	case uint64:
		it := value.(uint64)
		key = strconv.FormatUint(it, 10)
	case string:
		key = value.(string)
	case []byte:
		key = string(value.([]byte))
	default:
		newValue, _ := json.Marshal(value)
		key = string(newValue)
	}

	return key
}





































