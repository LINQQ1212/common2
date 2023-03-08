package models

import (
	"encoding/json"
	"github.com/LINQQ1212/common2/config/core"
	"go.etcd.io/bbolt"
	"golang.org/x/sync/singleflight"
	"google.golang.org/protobuf/proto"
	"path"
)

var (
	BCate         = []byte("Cate")
	BCateProducts = []byte("CateProducts")
	BDomain       = []byte("Domain")
	BProduct      = []byte("Product")
	BProductInfo  = []byte("ProductInfo")
	BInfo         = []byte("Info")
)

type Cate struct {
	ID         uint64   `json:"ID,omitempty" storm:"id,increment"` // @gotags: storm:"id,increment"
	Name       string   `json:"Name,omitempty"`
	CategoryId string   `json:"CategoryId,omitempty"`
	DomainID   uint64   `json:"DomainID,omitempty"`
	ParentId   uint64   `json:"ParentId,omitempty"`
	Count      uint64   `json:"Count,omitempty"`
	Products   []uint64 `json:"Products"`
	Children   []*Cate  `json:"-"`
}

type NewVersionReq struct {
	Domain         string `json:"domain"`
	ProductTarLink string `json:"product_tar_link"`
	GoogleImg      string `json:"google_img"`
	DownMainPic    bool   `json:"down_main_pic"`
	YahooDsc       string `json:"yahoo_dsc"`
	BingDsc        string `json:"bing_dsc"`
	YoutubeDsc     string `json:"youtube_dsc"`
}

type NewVersionReqV2 struct {
	Domain         string `json:"domain"`
	ProductTarLink string `json:"product_tar_link"`
	core.NewVersionOption
}

type Version struct {
	DB            *bbolt.DB
	Info          *VersionInfo
	Cates         map[uint64]*Cate
	Categories    []*Cate
	CategoriesLen int
	Domains       map[uint64]string
	ReviewNum     uint64
	SingleFlight  *singleflight.Group
}

func (v *Version) GetCate(id uint64) *Cate {
	if c, ok := v.Cates[id]; ok {
		return c
	}
	return v.Categories[id%uint64(v.CategoriesLen)]
}

func (v *Version) GetCateList(id uint64, l uint64) []*Cate {
	var list []*Cate
	index := id - l/2
	var i uint64 = 0
	for ; i < l; i++ {
		list = append(list, v.GetCate(index+i))
	}
	return list
}

func (v *Version) GetCates(id uint64) []*Cate {
	var list []*Cate
	tid := id
	for tid != 0 {
		if c, ok := v.Cates[tid]; ok {
			list = append(list, c)
			tid = c.ParentId
		} else {
			tid = 0
		}
	}
	return list
}

func (v *Version) EndCategories2(c *Cate, list *[]*Cate) {
	if c.Count > 0 {
		*list = append(*list, c)
	} else {
		for _, child := range c.Children {
			v.EndCategories2(child, list)
		}
	}
}

func NewVersion(dir string, fname string) (*Version, error) {
	db, err := bbolt.Open(path.Join(dir, fname+".db"), 0600, &bbolt.Options{Timeout: 0, FreelistType: bbolt.FreelistArrayType})
	//db, err := storm.Open(path.Join(dir, fname+".db"), storm.Codec(protobuf.Codec), storm.BoltOptions(0644, &bbolt.Options{ReadOnly: true}))
	if err != nil {
		return nil, err
	}

	//v := &Version{DB: db, Domains: map[uint64]string{}, Info: &VersionInfo{}}
	v := &Version{DB: db, Cates: map[uint64]*Cate{}, Domains: map[uint64]string{}, Info: &VersionInfo{}}
	v.SingleFlight = &singleflight.Group{}

	err = db.View(func(tx *bbolt.Tx) error {
		err = proto.Unmarshal(tx.Bucket(BInfo).Get(BInfo), v.Info)
		if err != nil {
			return err
		}

		err = json.Unmarshal(tx.Bucket(BCate).Get(BCate), &v.Categories)
		if err != nil {
			return err
		}

		err = json.Unmarshal(tx.Bucket(BDomain).Get(BDomain), &v.Domains)
		if err != nil {
			return err
		}
		return nil
	})

	for _, c := range v.Categories {
		v.Cates[c.ID] = c
	}
	v.CategoriesLen = len(v.Categories)
	for _, cate := range v.Cates {
		if cate.ParentId != 0 {
			if v.Cates[cate.ParentId] == nil {
				continue
			}
			if v.Cates[cate.ParentId].Children == nil {
				v.Cates[cate.ParentId].Children = []*Cate{}
			}
			v.Cates[cate.ParentId].Children = append(v.Cates[cate.ParentId].Children, cate)
		}
	}

	for _, cate := range v.Cates {
		cate.Count = GetChildrenCount(cate)
	}

	if err != nil {
		return nil, err
	}
	/*if num > v.Info.Count {
		v.ReviewNum = num / v.Info.Count

	}*/
	return v, nil
}

func GetChildrenCount(c *Cate) uint64 {
	if len(c.Children) > 0 {
		for _, child := range c.Children {
			c.Count += GetChildrenCount(child)
		}
	}
	return c.Count
}
