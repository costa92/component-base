package time

import (
	"fmt"
	"testing"
)

func Test_MarshalJSON(t *testing.T)  {
	time := Time{}
	result,_ := time.MarshalJSON()
	fmt.Println(string(result))
}


func Test_ToTime(t *testing.T)  {
	result,_ := ToTime("2022-04-03 00:00:00")

	fmt.Println(result.Time.Unix())
}
