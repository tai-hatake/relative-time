# 相対日時表記
○時間前、○分前といった相対日時表記に変換します。

## 始め方

```
go get -u github.com/tai-hatake/relative-time
```

## 仕様
* stringで返却します。

| 条件（現在時刻と比較） | 形式 |
| ---- | ---- | 
| 60秒前未満 | s秒前 |
| 60分前未満 | m分前 |
| 24時間前未満 | h時間前 |
| 1日前 | 昨日 | |
| 3日前まで | d日前 |
| 今年 | M月D日 |
| 一昨年以降 | Y年M月D日 |
