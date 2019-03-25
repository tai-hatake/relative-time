package related

import (
	"fmt"
	"strconv"
	"time"
)

var loc, _ = time.LoadLocation("Local")

func relatedTime(target time.Time) string {
	formatted := ""
	now := time.Now()
	fmt.Println(now.Sub(target))

	if isPast(now, target) {
		diff := now.Sub(target)
		hours0 := int(diff.Hours())
		days := hours0 / 24
		hours := hours0 % 24
		mins := int(diff.Minutes()) % 60
		secs := int(diff.Seconds()) % 60

		if days < 365 {
			// 記事が60秒前未満の場合は。「s秒前」形式
			if days == 0 && hours == 0 && mins == 0 {
				return strconv.Itoa(secs) + "秒前"
			}
			// 記事が60分前未満の場合は「m分前」形式
			if days == 0 && hours == 0 {
				return strconv.Itoa(mins) + "分前"
			}
			// 記事が7時間前未満の場合は1時間単位の切り捨てで「h時間前」形式。ただし現在時刻が午前6時台かつ記事の日付が1日前の場合を除く
			if days == 0 && hours != 0 {
				return strconv.Itoa(hours) + "時間前"
			}

			// 記事の日付が現在の1日前の場合は「きのう」
			if days <= 1 {
				return "昨日"
			}

			// 記事の日付が現在の3日前までの場合は「d日前」形式
			if days <= 3 {
				return strconv.Itoa(days) + "日前"
			}

			// 記事の年が現在と同じ場合は「M月D日」形式
			const layout = "01月02日"
			formatted = target.Format(layout)
			return formatted
		}

		// 記事の年が現在と異なる場合は「Y年M月D日」形式
		const layout = "2006年01月02日"
		formatted = target.Format(layout)
		return formatted
	}
	return "未来"
}

func isPast(now time.Time, target time.Time) bool {
	return target.Before(now)
}
