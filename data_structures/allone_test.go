package data_structures

import (
	"fmt"
	"testing"
)

func TestAllOne_Dec(t *testing.T) {

	//["AllOne","inc","inc","getMaxKey","getMinKey","inc","getMaxKey","getMinKey"]
	//[[],["hello"],["hello"],[],[],["leet"],[],[]]
	// [null,null,null,"hello","hello",null,"hello","leet"]
	ao := Constructor()
	fmt.Println(ao, ao.GetMaxKey(), ao.GetMinKey())

	ao.Inc("hello")
	ao.Inc("hello")
	fmt.Println(ao, ao.GetMaxKey(), ao.GetMinKey())
	ao.Inc("leet")
	fmt.Println(ao, ao.GetMaxKey(), ao.GetMinKey())
	//ao.Inc("hello")
	//fmt.Println(ao, ao.GetMinKey(), ao.GetMaxKey())
	//ao.Inc("hello")
	//fmt.Println(ao, ao.GetMinKey(), ao.GetMaxKey())
	//
	//ao.Dec("hello")
	//ao.Dec("hello")
	//fmt.Println(ao, ao.GetMinKey(), ao.GetMaxKey())
}
