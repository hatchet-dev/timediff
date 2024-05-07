package timediff_test

import (
	"fmt"
	"testing"

	"github.com/hatchet-dev/timediff"
)

var fixtures_ko_KR = map[string]string{
	"-10s":                            "방금",
	"-44s":                            "방금",
	"-45s":                            "약 1분 전",
	"-89s":                            "약 1분 전",
	"-90s":                            "2분 전",
	"-91s":                            "2분 전",
	"-2m":                             "2분 전",
	"-10m":                            "10분 전",
	"-44m":                            "44분 전",
	"-45m":                            "약 한 시간 전",
	"-60m":                            "약 한 시간 전",
	"-1h":                             "약 한 시간 전",
	"-80m":                            "약 한 시간 전",
	"-89m":                            "약 한 시간 전",
	"-90m":                            "2시간 전",
	"-2h":                             "2시간 전",
	"-20h":                            "20시간 전",
	"-21h":                            "21시간 전",
	"-21h30m":                         "어제",
	"-22h":                            "어제",
	"-24h":                            "어제",
	"-24h30m":                         "어제",
	"-34h59m":                         "어제",
	"-35h10m":                         "그제",
	"-36h":                            "그제",
	"-40h":                            "그제",
	"-58h59m":                         "그제",
	fmt.Sprintf("-%dh", 10*24):        "10일 전",
	fmt.Sprintf("-%dh", 25*24):        "25일 전",
	fmt.Sprintf("-%dh", 26*24):        "저번 달",
	fmt.Sprintf("-%dh", 45*24):        "저번 달",
	fmt.Sprintf("-%dh2m", 45*24):      "2개월 전",
	fmt.Sprintf("-%dh", 46*24+1):      "2개월 전",
	fmt.Sprintf("-%dh", 80*24):        "3개월 전",
	fmt.Sprintf("-%dh", 9*24*30):      "9개월 전",
	fmt.Sprintf("-%dh", 10*24*30):     "10개월 전",
	fmt.Sprintf("-%dh1m", 10*24*30):   "작년",
	fmt.Sprintf("-%dh", 12*24*30):     "작년",
	fmt.Sprintf("-%dh", 17*24*30+1):   "2년 전",
	fmt.Sprintf("-%dh", 24*24*30):     "2년 전",
	fmt.Sprintf("-%dh", 20*24*30*12):  "20년 전",
	fmt.Sprintf("-%dh", 100*24*30*12): "100년 전",

	"10s":                            "곧",
	"44s":                            "곧",
	"45s":                            "약 1분 뒤",
	"89s":                            "약 1분 뒤",
	"90s":                            "2분 뒤",
	"2m":                             "2분 뒤",
	"10m":                            "10분 뒤",
	"44m":                            "44분 뒤",
	"45m":                            "약 한 시간 뒤",
	"60m":                            "약 한 시간 뒤",
	"1h":                             "약 한 시간 뒤",
	"80m":                            "약 한 시간 뒤",
	"89m":                            "약 한 시간 뒤",
	"89m10s":                         "2시간 뒤",
	"90m":                            "2시간 뒤",
	"2h":                             "2시간 뒤",
	"20h":                            "20시간 뒤",
	"21h":                            "21시간 뒤",
	"21h30m":                         "내일",
	"22h":                            "내일",
	"24h":                            "내일",
	"35h10m":                         "모레",
	"36h":                            "모레",
	"40h":                            "모레",
	"58h59m":                         "모레",
	fmt.Sprintf("%dh", 10*24):        "10일 뒤",
	fmt.Sprintf("%dh", 25*24):        "25일 뒤",
	fmt.Sprintf("%dh", 26*24):        "다음 달",
	fmt.Sprintf("%dh", 45*24):        "다음 달",
	fmt.Sprintf("%dh1m", 45*24):      "2개월 뒤",
	fmt.Sprintf("%dh", 46*24):        "2개월 뒤",
	fmt.Sprintf("%dh", 80*24):        "3개월 뒤",
	fmt.Sprintf("%dh", 9*24*30):      "9개월 뒤",
	fmt.Sprintf("%dh", 10*24*30):     "10개월 뒤",
	fmt.Sprintf("%dh1m", 10*24*30):   "내년",
	fmt.Sprintf("%dh", 12*24*30):     "내년",
	fmt.Sprintf("%dh", 24*24*30):     "2년 뒤",
	fmt.Sprintf("%dh", 20*24*30*12):  "20년 뒤",
	fmt.Sprintf("%dh", 100*24*30*12): "100년 뒤",
}

func TestTimeDiffKoKR(t *testing.T) {
	execFixtures(t, fixtures_ko_KR, timediff.WithLocale("ko-KR"))
}
