package daos

import (
	"fmt"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

// NullType 空字节类型
type NullType byte

const (
	_ NullType = iota
	// IsNull the same as `is null`
	IsNull
	// IsNotNull the same as `is not null`
	IsNotNull
)

// WhereBuild sql build where
func WhereBuild(where map[string]interface{}) (whereSQL string, vals []interface{}, err error) {
	var wheres []string
	for k, v := range where {
		ks := strings.Split(k, " ")
		if len(ks) > 2 {
			return "", nil, fmt.Errorf("Error in query condition: %s. ", k)
		}
		switch len(ks) {
		case 1:
			switch v {
			case IsNull:
				wheres = append(wheres, fmt.Sprint(k, " IS NULL"))
			case IsNotNull, nil:
				wheres = append(wheres, fmt.Sprint(k, " IS NOT NULL"))
			default:
				switch reflect.TypeOf(v).Kind() {
				case reflect.Slice:
					lenArray := reflect.ValueOf(v).Len()
					if lenArray != 0 {
						wheres = append(wheres, k+" in ?")
						vals = append(vals, v)
					}
				default:
					wheres = append(wheres, k+" = ?")
					vals = append(vals, v)
				}
			}
		case 2:
			k = ks[0]
			switch ks[1] {
			case "=":
				wheres = append(wheres, fmt.Sprint(k, "=?"))
				vals = append(vals, v)
			case ">":
				wheres = append(wheres, fmt.Sprint(k, ">?"))
				vals = append(vals, v)
			case ">=":
				wheres = append(wheres, fmt.Sprint(k, ">=?"))
				vals = append(vals, v)
			case "<":
				wheres = append(wheres, fmt.Sprint(k, "<?"))
				vals = append(vals, v)
			case "<=":
				wheres = append(wheres, fmt.Sprint(k, "<=?"))
				vals = append(vals, v)
			case "!=":
				wheres = append(wheres, fmt.Sprint(k, "!=?"))
				vals = append(vals, v)
			case "<>":
				wheres = append(wheres, fmt.Sprint(k, "!=?"))
				vals = append(vals, v)
			case "in":
				wheres = append(wheres, fmt.Sprint(k, " in (?)"))
				vals = append(vals, v)
			case "like":
				wheres = append(wheres, fmt.Sprint(k, " like ?"))
				vals = append(vals, v)
			default:
			}
		default:
		}
	}

	whereSQL = strings.Join(wheres, " AND ")
	return
}

type Dao struct {
	DB *gorm.DB
}

func NewDao() *Dao {
	return &Dao{DB: GetDB()}
}
