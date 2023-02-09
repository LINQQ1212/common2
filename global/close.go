package global

import "github.com/LINQQ1212/common2/models"

func Close() {
	Versions.Range(func(k string, v *models.Version) bool {
		v.DB.Close()
		return true
	})
	//IMGDB.Close()
	//Review.Close()
}
