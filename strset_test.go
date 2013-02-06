package villa

import(
    "testing"
    "fmt"
)

func TestStrSet(t *testing.T) {
    defer __(o_())

	var set StrSet
	
	set.Put("hello", "david")
	fmt.Println(set)
	AssertEquals(t, "set.In(hello)", set.In("hello"), true)
	AssertEquals(t, "set.In(david)", set.In("david"), true)
	AssertEquals(t, "set.In(villa)", set.In("villa"), false)
	
	set.Delete("david")
	AssertEquals(t, "set.In(david)", set.In("david"), false)
}