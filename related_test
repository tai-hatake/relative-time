package related

import (
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFuture(t *testing.T) {
	// 明日
	target := time.Now().AddDate(0, 0, 1)
	result := relatedTime(target)
	assert.Equal(t, "未来", result, "ok")
}

func TestOverOneYearAgo(t *testing.T) {
	// 1年以上前
	target := time.Now().AddDate(-1, -1, -1)
	result := relatedTime(target)
	isTrue := check_regexp(`(19[0-9]{2}|20[0-9]{2})年(0[0-9]|1[0-2])月([1-9]|[12][0-9]|3[01])日`, result)
	assert.False(t, isTrue)
}

func TestWithinOneYear(t *testing.T) {
	// 1年以内
	target := time.Now().AddDate(0, -2, -1)
	result := relatedTime(target)
	isTrue := check_regexp(`([0-9]|1[0-2])月([1-9]|[12][0-9]|3[01])日`, result)
	assert.False(t, isTrue)
}

func TestThreeDaysAgo(t *testing.T) {
	// 3日前まで
	target := time.Now().AddDate(0, 0, -3)
	result := relatedTime(target)
	isTrue := check_regexp(`[2-3]日前`, result)
	assert.False(t, isTrue)
}

func TestYesterday(t *testing.T) {
	// 昨日まで
	target := time.Now().AddDate(0, 0, -1)
	result := relatedTime(target)
	assert.Equal(t, "昨日", result, "ok")
}

func TestToday(t *testing.T) {
	// 今日まで
	target := time.Now()
	diffHour := (target.Hour() - 1) * -1
	target = target.Add(time.Duration(diffHour) * time.Hour)
	result := relatedTime(target)
	isTrue := check_regexp(`[0-2][0-9]時間前`, result)
	assert.False(t, isTrue)
}

func TestMinuteAgo(t *testing.T) {
	// ○分前まで
	target := time.Now()
	diffMin := (target.Minute() - 1) * -1
	target = target.Add(time.Duration(diffMin) * time.Minute)
	result := relatedTime(target)
	fmt.Println("result", result)
	isTrue := check_regexp(`[0-5][0-9]分前`, result)
	assert.False(t, isTrue)
}

func TestSecondAgo(t *testing.T) {
	// ○秒前まで
	target := time.Now()
	diffSec := (target.Minute() - 1) * -1
	target = target.Add(time.Duration(diffSec) * time.Second)
	result := relatedTime(target)
	fmt.Println("result", result)
	isTrue := check_regexp(`[0-5][0-9]秒前`, result)
	assert.False(t, isTrue)
}

func check_regexp(reg, str string) bool {
	return !regexp.MustCompile(reg).Match([]byte(str))
}
