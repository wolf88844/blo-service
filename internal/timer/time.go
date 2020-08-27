package timer

import "time"
const(
	Nanosecond  = 1
	Microsecond =1000 * Nanosecond
	Millisecond = 1000 * Microsecond
	Second = 1000 * Millisecond
	Minute  = 60 *Second
	Hour = 60 * Minute
)
func GetNowTime() time.Time{
	location,_:=time.LoadLocation("Asia/Shanghai")
	return time.Now().In(location)
}

func GetCalculateTime(currentTimer time.Time,d string) (time.Time,error){
	duration,err:=time.ParseDuration(d)
	if err != nil{
		return time.Time{},err
	}
	return currentTimer.Add(duration),nil
}
