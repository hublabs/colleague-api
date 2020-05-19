package colleagues

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func Seed(xormEngine *xorm.Engine) error {

	var (
		colleagues = []Colleague{
			{Id: 1, Name: "xiao_ming", Email: "xiao_ming@email.com", Password: "1111", Enable: true},
			{Id: 2, Name: "xiao_zhang", Enable: true},
			{Id: 3, Name: "lao_li", Enable: true},
			{Id: 4, Name: "lao_wang", Enable: false},
			{Id: 5, Name: "lao_zhang", Enable: true},
		}

		stores = []Store{
			{Id: 1, Code: "C001", Name: "北京朝阳门店", Enable: true},
			{Id: 2, Code: "C002", Name: "北京新世界百货店", Enable: true},
			{Id: 3, Code: "C003", Name: "上海西单店", Enable: true},
		}

		storeColleagues = []StoreColleague{
			{Id: 1, StoreId: 1, ColleagueId: 1, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 2, StoreId: 2, ColleagueId: 1, StartDate: "", EndDate: "", Role: "member", Enable: true},
			{Id: 3, StoreId: 3, ColleagueId: 1, StartDate: "", EndDate: "", Role: "guest", Enable: true},
			{Id: 4, StoreId: 1, ColleagueId: 2, StartDate: "", EndDate: "", Role: "admin", Enable: true},
			{Id: 5, StoreId: 1, ColleagueId: 3, StartDate: "", EndDate: "", Role: "admin", Enable: true},
		}
	)

	for _, u := range colleagues {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	for _, u := range stores {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	for _, u := range storeColleagues {
		if _, err := xormEngine.Insert(&u); err != nil {
			return err
		}
	}

	return nil
}