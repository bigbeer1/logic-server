package number

import (
	"errors"
	"fmt"
	"github.com/w3liu/go-common/constant/timeformat"
	"math/rand"
	"strconv"
	"strings"
	"sync/atomic"
	"time"
)

var num int64

// 生成24位资产编号
// 前面17位代表时间精确到毫秒，中间3位代表进程id，最后4位代表序号
func GetAssetCode(t time.Time) string {
	s := t.Format(timeformat.Continuity)
	m := t.UnixNano()/1e6 - t.UnixNano()/1e9*1e3
	ms := sup(m, 3)
	//p := os.Getpid() % 1000
	//ps := sup(int64(p), 3)
	i := atomic.AddInt64(&num, 1)
	r := i % 10000
	rs := sup(r, 4)
	//n := fmt.Sprintf("%s%s%s%s", s, ms, ps, rs)
	n := fmt.Sprintf("%s%s%s", s, ms, rs)
	return n
}

// 对长度不足n的数字前面补0
func sup(i int64, n int) string {
	m := fmt.Sprintf("%d", i)
	for len(m) < n {
		m = fmt.Sprintf("0%s", m)
	}
	return m
}

//

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStr(n int) string {

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func NextSerialNumber(ticketPrefix string, year, month, day, serialNumberLen int64, upperSerialNumber string) (string, error) {
	var serialNumber string
	var serialNumberLenString string
	var serialNumberYear, serialNumberMonth, serialNumberDay string
	nowYear, nowMonth, nowDay := time.Now().Date()
	// 如果没有上一个编号规则 生成初始编号
	if upperSerialNumber == "" {
		serialNumber = CreateSerialNumber(ticketPrefix, year, month, day, serialNumberLen)
		return serialNumber, nil
	}

	// 归零标志
	zero := 0
	upperSerials := strings.Split(upperSerialNumber, "")
	index := len(ticketPrefix)
	if year == 1 {
		serialNumberYear = fmt.Sprintf("%v", nowYear)
		ticketYear := fmt.Sprintf("%v%v%v%v", upperSerials[index], upperSerials[index+1], upperSerials[index+2], upperSerials[index+3])
		// 判断是否超过本年是否归零
		if ticketYear != serialNumberYear {
			zero = 1
		}

		// 增加index计数
		index = index + 4
	}
	if month == 1 {
		if int(nowMonth) < 10 {
			serialNumberMonth = fmt.Sprintf("0%v", int(nowMonth))
		} else {
			serialNumberMonth = fmt.Sprintf("%v", int(nowMonth))
		}
		ticketMonth := fmt.Sprintf("%v%v", upperSerials[index], upperSerials[index+1])
		// 判断是否超过本年是否归零
		if ticketMonth != serialNumberMonth {
			zero = 1
		}

		// 增加index计数
		index = index + 2
	}

	if day == 1 {
		if nowDay < 10 {
			serialNumberDay = fmt.Sprintf("0%v", nowDay)
		} else {
			serialNumberDay = fmt.Sprintf("%v", nowDay)
		}
		ticketDay := fmt.Sprintf("%v%v", upperSerials[index], upperSerials[index+1])
		// 判断是否超过本年是否归零
		if ticketDay != serialNumberDay {
			zero = 1
		}

		// 增加index计数
		index = index + 2
	}

	// 如果zero==1 归零
	if zero == 1 {
		for i := 0; i < int(serialNumberLen); i++ {
			serialNumberLenString += "0"
		}
	} else {
		// 数据和前面数据符合
		if int(serialNumberLen)+index == len(upperSerials) {
			var ticketSerialNumberString string
			for i := 0; i < int(serialNumberLen); i++ {
				ticketSerialNumberString += upperSerials[index+i]
			}
			ticketSerialNumberInt, err := strconv.Atoi(ticketSerialNumberString)
			if err != nil {
				return "", err
			}
			ticketSerialNumberInt = ticketSerialNumberInt + 1

			differ := int(serialNumberLen) - len(fmt.Sprintf("%v", ticketSerialNumberInt))

			if differ < 0 {
				return "", errors.New("已超过流水号,无法创建")
			}
			for i := 0; i < differ; i++ {
				serialNumberLenString += "0"
			}
			serialNumberLenString += fmt.Sprintf("%v", ticketSerialNumberInt)

		} else {
			return "", errors.New("数据长度存在异常,无法生成,请联系管理员,或在管理员建议下删除此编号规则")
		}
	}

	serialNumber = fmt.Sprintf("%s%s%s%s%s", ticketPrefix, serialNumberYear, serialNumberMonth, serialNumberDay, serialNumberLenString)

	return serialNumber, nil

}

func CreateSerialNumber(ticketPrefix string, year, month, day, serialNumberLen int64) string {
	var serialNumber string
	var serialNumberLenString string
	var serialNumberYear, serialNumberMonth, serialNumberDay string
	nowYear, nowMonth, nowDay := time.Now().Date()
	for i := 0; i < int(serialNumberLen); i++ {
		serialNumberLenString += "0"
	}
	if year == 1 {
		serialNumberYear = fmt.Sprintf("%v", nowYear)
	}
	if month == 1 {
		if int(nowMonth) < 10 {
			serialNumberMonth = fmt.Sprintf("0%v", int(nowMonth))
		} else {
			serialNumberMonth = fmt.Sprintf("%v", int(nowMonth))
		}
	}
	if day == 1 {
		if nowDay < 10 {
			serialNumberDay = fmt.Sprintf("0%v", nowDay)
		} else {
			serialNumberDay = fmt.Sprintf("%v", nowDay)
		}
	}
	serialNumber = fmt.Sprintf("%s%s%s%s%s", ticketPrefix, serialNumberYear, serialNumberMonth, serialNumberDay, serialNumberLenString)

	return serialNumber
}

//func CreateSerialNumber() string {
//	rand.Seed(time.Now().UnixNano())
//	// 随机生成4位英文字母
//	cc := randStr(4)
//	// 加年
//	bb := time.Now().Year()
//	// 5位流水号
//	newSerialNumber := fmt.Sprintf("%v%v%v", cc, bb, "00000")
//	return newSerialNumber
//
//}
//
//func NextSerialNumber(oldSerialNumber string) (string, error) {
//
//	var newSerialNumber string
//
//	oldSerialNumbers := strings.Split(oldSerialNumber, "")
//	if oldSerialNumber == "" || len(oldSerialNumbers) != 13 {
//		newSerialNumber = CreateSerialNumber()
//		return newSerialNumber, nil
//	}
//
//	lastNumberString := oldSerialNumbers[4] + oldSerialNumbers[5] + oldSerialNumbers[6] + oldSerialNumbers[7] +
//		oldSerialNumbers[8] + oldSerialNumbers[9] + oldSerialNumbers[10] + oldSerialNumbers[11] + oldSerialNumbers[12]
//	lastNumber, err := strconv.ParseInt(lastNumberString, 10, 64)
//	if err != nil {
//		return "", err
//	}
//	for i := 0; i < len(oldSerialNumbers)-9; i++ {
//		newSerialNumber = newSerialNumber + oldSerialNumbers[i]
//	}
//	newSerialNumber = newSerialNumber + strconv.FormatInt(lastNumber+1, 10)
//
//	return newSerialNumber, nil
//
//}
