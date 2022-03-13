package uuid

import (
	"github.com/GUAIK-ORG/go-snowflake/snowflake"
)

func Getuuid() (id int64, err error) {
	s, err := snowflake.NewSnowflake(int64(0), int64(0))
	id = s.NextVal()

	return id, err
}
