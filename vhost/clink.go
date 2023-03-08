package vhost

import (
	"github.com/LINQQ1212/common2/utils"
	"strings"
)

func (s *Site) buildCategoryUri1(host string, id uint64, name string) string {
	if s.IsParam {
		return host + "?" + s.IdToStr(id)
	}
	return host + s.IdToStr(id)
}
func (s *Site) CategoryUri1(uri string) uint64 {
	return s.StrToId(uri)
}

func (s *Site) buildCategoryUri2(host string, id uint64, name string) string {
	if s.IsParam {
		return host + "?" + s.Base32Id(id)
	}
	return host + s.Base32Id(id)
}

func (s *Site) CategoryUri2(uri string) uint64 {
	return s.Base32ToId(uri)
}

func (s *Site) buildCategoryUri3(host string, id uint64, name string) string {
	if s.IsParam {
		return host + "?" + s.IdToStr(s.Num2) + s.IdToStr(id)
	}
	return host + s.IdToStr(s.Num2) + s.IdToStr(id)
}

func (s *Site) CategoryUri3(uri string) uint64 {
	uri = strings.TrimPrefix(uri, s.IdToStr(s.Num2))
	return s.StrToId(uri)
}

func (s *Site) buildCategoryUri4(host string, id uint64, name string) string {
	uu := strings.ReplaceAll(UrlPathEscape(name), "-", "")
	if s.IsParam {
		return host + "?" + uu + "-" + s.IdToStr(id)
	}
	return host + uu + "-" + s.IdToStr(id)
}

func (s *Site) CategoryUri4(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(uri[index+1:])
}

func (s *Site) buildCategoryUri5(host string, id uint64, name string) string {
	uu := strings.ReplaceAll(UrlPathEscape(name), "-", "")
	if s.IsParam {
		return host + "?" + s.IdToStr(id) + "-" + uu
	}
	return host + s.IdToStr(id) + "-" + uu
}

func (s *Site) CategoryUri5(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(uri[:index])
}

func (s *Site) buildCategoryUri6(host string, id uint64, name string) string {
	uu := strings.ReplaceAll(UrlPathEscape(name), "-", "")
	if s.IsParam {
		return host + "?" + uu + "-" + s.IdToStr(s.Num2) + s.IdToStr(id)
	}
	return host + uu + "-" + s.IdToStr(s.Num2) + s.IdToStr(id)
}

func (s *Site) CategoryUri6(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(uri[index+1:])
}

func (s *Site) buildCategoryUri7(host string, id uint64, name string) string {
	uu := strings.ReplaceAll(UrlPathEscape(name), "-", "")
	if s.IsParam {
		return host + "?" + s.IdToStr(s.Num2) + s.IdToStr(id) + "-" + uu
	}
	return host + s.IdToStr(s.Num2) + s.IdToStr(id) + "-" + uu
}

func (s *Site) CategoryUri7(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(uri[:index])
}

func (s *Site) buildCategoryUri8(host string, id uint64, name string) string {
	uu := strings.ReplaceAll(UrlPathEscape(name), "-", "")
	if s.IsParam {
		return host + "?" + s.Base32Id(id) + "-" + uu
	}
	return host + s.Base32Id(id) + "-" + uu
}

func (s *Site) CategoryUri8(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.Base32ToId(uri[:index])
}

func (s *Site) buildCategoryUri9(host string, id uint64, name string) string {
	uu := strings.ReplaceAll(UrlPathEscape(name), "-", "")
	if s.IsParam {
		return host + "?" + uu + "-" + s.Base32Id(id)
	}
	return host + uu + "-" + s.Base32Id(id)
}

func (s *Site) CategoryUri9(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.Base32ToId(uri[index+1:])
}

func (s *Site) buildCategoryUri10(host string, id uint64, name string) string {
	uu := utils.RandChar(6)
	if s.IsParam {
		return host + "?" + uu + "-" + s.IdToStr(id)
	}
	return host + uu + "-" + s.IdToStr(id)
}

func (s *Site) CategoryUri10(uri string) uint64 {
	index := strings.Index(uri, "-")
	if index < 0 {
		return 0
	}
	return s.StrToId(uri[index+1:])
}
