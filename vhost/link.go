package vhost

import (
	"strings"
)

/*
要有二级

nc
可以没有二级
sdfds6f165-126546513.html
产品名-561461.html
*/

// buildProductUri1 链接1
func (s *Site) buildProductUri1(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + s.IdToStr(id) + s.Suffix
}

func (s *Site) ProductUri1(uri string) uint64 {
	return s.StrToId(uri)
}

// buildProductUri2 链接2
func (s *Site) buildProductUri2(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + s.Base32Id(id) + s.Suffix
}

func (s *Site) ProductUri2(uri string) uint64 {
	return s.Base32ToId(uri)
}

func (s *Site) buildProductUri3(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + s.IdToStr(s.Num2) + s.IdToStr(id) + s.Suffix
}

func (s *Site) ProductUri3(uri string) uint64 {
	uri = strings.TrimPrefix(uri, s.IdToStr(s.Num2))
	return s.StrToId(uri)
}

func (s *Site) buildProductUri4(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + UrlPathEscape(name) + "-" + s.IdToStr(id) + s.Suffix
}

func (s *Site) ProductUri4(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(uri[index+1:])
}

func (s *Site) buildProductUri5(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + s.IdToStr(id) + "-" + UrlPathEscape(name) + s.Suffix
}

func (s *Site) ProductUri5(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(uri[:index])
}

/**************/

func (s *Site) buildProductUri6(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + UrlPathEscape(name) + "-" + s.Base32Id(id) + s.Suffix
}

func (s *Site) ProductUri6(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.Base32ToId(uri[index+1:])
}

func (s *Site) buildProductUri7(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + s.Base32Id(id) + "-" + UrlPathEscape(name) + s.Suffix
}

func (s *Site) ProductUri7(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.Base32ToId(uri[:index])
}

func (s *Site) buildProductUri8(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + s.IdToStr(s.Num2) + s.IdToStr(id) + "-" + UrlPathEscape(name) + s.Suffix
}

func (s *Site) ProductUri8(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(strings.TrimPrefix(uri[:index], s.IdToStr(s.Num2)))
}

func (s *Site) buildProductUri9(host string, cid, id uint64, cname, name string, c, hc bool, option int32) string {
	cc := ""
	if c && hc {
		cc = s.BuildCategoryUriStr(host, cid, cname, option) + "/"
	} else {
		cc = host
		if s.IsParam {
			cc += "?"
		}
	}
	return cc + UrlPathEscape(name) + "-" + s.IdToStr(s.Num2) + s.IdToStr(id) + s.Suffix
}

func (s *Site) ProductUri9(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(strings.TrimPrefix(uri[index+1:], s.IdToStr(s.Num2)))
}
