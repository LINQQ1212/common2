package vhost

import (
	"github.com/LINQQ1212/common2/utils"
	"github.com/jxskiss/base62"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
)

type Site struct {
	Domain string
	Index  int
	Table  string // 字符串 表
	Key    string // 字符串 表
	//TableMap   map[string]string `json:"-" yaml:"-"`
	ListTemp   string
	TempsLen   int
	Temps      []string // 模板还是要一次性使用
	Num        int      // 随机数
	Num2       uint64   // 随机数2
	Path       string   // 伪装路径
	CPath      string   // 分类伪装路径
	Suffix     string   // 后缀
	SA         string   // html 前
	SE         string   // html 后
	SearchPath string   `json:"search-path" yaml:"search-path"`
	IsParam    bool
	Type       int              // 方式
	HashID     *base62.Encoding `json:"-" yaml:"-"`
	MinLength  int
}

func (s *Site) BuildImageLink(host, f, path string) string {
	if s.IsParam {
		return host + "?images=" + strings.TrimPrefix(path, "/images/")
	}
	return host + path
}
func (s *Site) BuildPath(host, f, path string) string {
	if s.IsParam {
		return host + "?" + path
	}
	return host + strings.TrimPrefix(path, "/")
}

func (s *Site) BuildPath2(host, f, path string) string {
	if s.IsParam {
		return host + "?" + path + s.Suffix
	}
	return host + path + s.Suffix
}

/*
	func (s *Site) BuildProUri(host, f string, id uint64) string {
		if s.IsParam {
			return host + "?" + s.Path + "/" + s.Base32Id(id) + s.Suffix
		}
		return host + s.Path + "/" + s.Base32Id(id) + s.Suffix
	}
*/
func (s *Site) BuildCategoryUri(host string, id uint64, name string, option int32) string {
	return s.BuildCategoryUriStr(host, id, name, option) + s.Suffix
}

func (s *Site) BuildCategoryUriStr(host string, id uint64, name string, option int32) string {
	///
	num := int(option)
	if num == 0 {
		num = s.Index % 9
	}
	switch num {
	case 1:
		return s.buildCategoryUri1(host, id, name)
	case 2:
		return s.buildCategoryUri2(host, id, name)
	case 3:
		return s.buildCategoryUri3(host, id, name)
	case 4:
		return s.buildCategoryUri4(host, id, name)
	case 5:
		return s.buildCategoryUri5(host, id, name)
	case 6:
		return s.buildCategoryUri6(host, id, name)
	case 7:
		return s.buildCategoryUri7(host, id, name)
	case 8:
		return s.buildCategoryUri8(host, id, name)
	case 9:
		return s.buildCategoryUri9(host, id, name)
	}
	return s.buildCategoryUri1(host, id, name)
}

func (s *Site) CategoryUriToId(uri string, option int32) uint64 {
	num := int(option)
	if num == 0 {
		num = s.Index % 10
	}
	switch num {
	case 1:
		return s.CategoryUri1(uri)
	case 2:
		return s.CategoryUri2(uri)
	case 3:
		return s.CategoryUri3(uri)
	case 4:
		return s.CategoryUri4(uri)
	case 5:
		return s.CategoryUri5(uri)
	case 6:
		return s.CategoryUri6(uri)
	case 7:
		return s.CategoryUri7(uri)
	case 8:
		return s.CategoryUri8(uri)
	case 9:
		return s.CategoryUri9(uri)
	}
	return s.CategoryUri1(uri)
}

func (s *Site) BuildProductUri(host string, cid, id uint64, cname, name string, coption, option int32, hc bool) string {
	num := int(option)
	if num == 0 {
		num = s.Index % 16
	}
	switch num {
	case 1:
		return s.buildProductUri1(host, cid, id, cname, name, true, hc, coption)
	case 2:
		return s.buildProductUri2(host, cid, id, cname, name, true, hc, coption)
	case 3:
		return s.buildProductUri3(host, cid, id, cname, name, true, hc, coption)
	case 4:
		return s.buildProductUri4(host, cid, id, cname, name, true, hc, coption)
	case 5:
		return s.buildProductUri5(host, cid, id, cname, name, true, hc, coption)
	case 6:
		return s.buildProductUri6(host, cid, id, cname, name, true, hc, coption)
	case 7:
		return s.buildProductUri7(host, cid, id, cname, name, true, hc, coption)
	case 8:
		return s.buildProductUri4(host, cid, id, cname, name, false, hc, coption)
	case 9:
		return s.buildProductUri5(host, cid, id, cname, name, false, hc, coption)
	case 10:
		return s.buildProductUri6(host, cid, id, cname, name, false, hc, coption)
	case 11:
		return s.buildProductUri7(host, cid, id, cname, name, false, hc, coption)
	case 12:
		return s.buildProductUri8(host, cid, id, cname, name, true, hc, option)
	case 13:
		return s.buildProductUri9(host, cid, id, cname, name, true, hc, option)
	case 14:
		return s.buildProductUri8(host, cid, id, cname, name, false, hc, coption)
	case 15:
		return s.buildProductUri9(host, cid, id, cname, name, false, hc, coption)
	}
	return s.buildProductUri1(host, cid, id, cname, name, true, hc, coption)
}

func (s *Site) ProductUriToId(uri string, option int32) uint64 {
	num := option
	if num == 0 {
		num = int32(s.Index) % 16
	}
	switch num {
	case 1:
		return s.ProductUri1(uri)
	case 2:
		return s.ProductUri2(uri)
	case 3:
		return s.ProductUri3(uri)
	case 4:
		return s.ProductUri4(uri)
	case 5:
		return s.ProductUri5(uri)
	case 6:
		return s.ProductUri6(uri)
	case 7:
		return s.ProductUri7(uri)
	case 8:
		return s.ProductUri4(uri)
	case 9:
		return s.ProductUri5(uri)
	case 10:
		return s.ProductUri6(uri)
	case 11:
		return s.ProductUri7(uri)
	case 12:
		return s.ProductUri8(uri)
	case 13:
		return s.ProductUri9(uri)
	case 14:
		return s.ProductUri8(uri)
	case 15:
		return s.ProductUri9(uri)
	}
	return s.ProductUri1(uri)
}

func (s *Site) BuildProUriCName(host, f, cname string, id uint64) string {
	if s.IsParam {
		return host + "?" + url.PathEscape(cname) + "/" + s.IdToStr(id) + s.Suffix
	}
	return host + url.PathEscape(cname) + "/" + s.IdToStr(id) + s.Suffix
}

func (s *Site) BuildProUriTPath(host, f, tpName string, id uint64) string {
	if s.IsParam {
		return host + "?" + tpName + "/" + s.IdToStr(id) + s.Suffix
	}
	return host + tpName + "/" + s.IdToStr(id) + s.Suffix
}

func GetSuffix() string {
	//arr := []string{"", ".html", ".shtml", ".jsp", ".asp", ".htm"}
	return utils.Suffixs[rand.Int()%len(utils.Suffixs)]
}
func GetSearch() string {
	arr := []string{"q", "query", "s", "keys", "search"}
	return arr[rand.Int()%5]
}

func (s *Site) InitTableMap() {
	/*i := 0
	s.TableMap = map[string]string{}
	for _, r := range s.Table {
		s.TableMap[string(r)] = strconv.Itoa(i)
		i++
		if i > 9 {
			i = 0
		}
	}*/
	s.HashID = base62.NewEncoding(s.Table)
}

func (s *Site) InitHtml() {
	switch rand.Int() % 8 {
	case 0:
		s.SA = "<p>"
		s.SE = "</p>"
	case 1:
		s.SA = "<span>"
		s.SE = "</span>"
	case 2:
		s.SA = ""
		s.SE = "</br>"
	case 3:
		s.SA = "<div><p>"
		s.SE = "</p></div>"
	case 4:
		s.SA = "<P>"
		s.SE = "</P>"
	case 5:
		s.SA = "<div><span>"
		s.SE = "</span></div>"
	case 6:
		s.SA = "<div>"
		s.SE = "</div>"
	case 7:
		s.SA = "<li>"
		s.SE = "</li>"
	}
}

func (s *Site) Str2Id(str string) uint64 {
	if str == "" {
		return 0
	}
	var id uint64
	for i, i2 := range str {
		id += uint64(i2 << i)
	}
	if id < 0 {
		id = ^id
	}
	return id * uint64(len(str)+s.Index)
}

func (s *Site) IdToStr(id uint64) string {
	return strconv.FormatUint(id+s.Num2, 10)
}
func (s *Site) StrToId(str string) uint64 {
	isNum := true
	for _, i2 := range str {
		if !unicode.IsNumber(i2) {
			isNum = false
			break
		}
	}
	if isNum {
		if num, err := strconv.ParseUint(str, 10, 64); err == nil {
			return num - s.Num2
		}
	}
	return 0
}

func (s *Site) HexId(str string) string {
	b := s.HashID.Encode([]byte(str))
	return *(*string)(unsafe.Pointer(&b))
}

func (s *Site) Base32Id(id uint64) string {
	b := s.HashID.Encode([]byte(strconv.FormatUint(id+s.Num2, 10)))
	return *(*string)(unsafe.Pointer(&b))
}
func (s *Site) Base32ToId(str string) uint64 {
	b, err := s.HashID.Decode([]byte(str))
	if err == nil {
		if i, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 64); err == nil {
			return i - s.Num2
		}
	}
	return 0
}

func (s *Site) Base32ToIdIsMy(str string) (uint64, bool) {
	b, err := s.HashID.Decode([]byte(str))
	if err == nil {
		if i, err := strconv.ParseUint(*(*string)(unsafe.Pointer(&b)), 10, 64); err == nil {
			return i, true
		}
		return s.Str2Id(str), true
	}
	return s.Str2Id(str), false
}

func (s *Site) IdToRandStr(id uint64) string {
	var b []byte
	cache := id
	v := false
	for i := 0; i < 10; i++ {
		if v {
			break
		}
		if idx := int(cache & 63); idx < len(s.Table) {
			if idx == 0 {
				v = true
			}
			b = append(b, s.Table[idx])
			cache >>= 6
		}
	}
	return *(*string)(unsafe.Pointer(&b))
}
