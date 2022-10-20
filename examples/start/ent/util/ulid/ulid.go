package ulid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"
)

var entropy *ulid.MonotonicEntropy
var nowFunc = time.Now

type ID string

func Parse(str string) (ID, error) {
	id, err := ulid.Parse(str)
	return ID(id.String()), err
}

func (i ID) String() string {
	return string(i)
}

func (i ID) Equals(id ID) bool {
	return i.String() == id.String()
}

func init() {
	// ulid 生成の度にエントロピーインスタンスを生成すると、同一ミリ秒に生成される ulid が同値になるので、
	// パッケージ初期化時に生成する
	entropy = ulid.Monotonic(rand.Reader, 0)
}

func New() ID {
	return ID(ulid.MustNew(ulid.Timestamp(nowFunc()), entropy).String())
}
