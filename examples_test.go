package jin

import "fmt"

func ExampleGet() {
	path := []string{"following", "social"}
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	value, err := Get(json, path...)
	// or without path variable
	// value, err := Get(json, "following", "social")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(value))
	// Output: dev.to
}

func ExampleSet() {
	path := []string{"following"}
	newValue := []byte(`["computerphile","numberphile"]`)
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	json, err := Set(json, newValue, path...)
	// or without path variable
	// value, err := Set(json, newValue, "following")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"user":"eco","languages":["go","java","python","C","Cpp"],"following":["computerphile","numberphile"]}
}

func ExampleSetKey() {
	path := []string{"following"}
	newKey := "sites"
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	json, err := SetKey(json, newKey, path...)
	// or without path variable
	// json, err := Set(json, newKey, "following")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"user":"eco","languages":["go","java","python","C","Cpp"],"sites":{"social":"dev.to","code":"github"}}
}

func ExampleAdd() {
	newValue := []byte(`"godoc.org/github.com/ecoshub"`)
	json := []byte(`{"user":"eco","links":["github.com/ecoshub"]}`)

	json, err := Add(json, newValue, "links")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"user":"eco","links":["github.com/ecoshub","godoc.org/github.com/ecoshub"]}
}

func ExampleAddKeyValue() {
	newValue := []byte(`"go"`)
	newKey := "language"
	json := []byte(`{"user":"eco"}`)

	json, err := AddKeyValue(json, newKey, newValue)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"user":"eco","language":"go"}
}

func ExampleDelete() {
	json := []byte(`{"user":"eco","languages":["go","java","python","C", "Cpp"]}`)

	json, err := Delete(json, "languages", "1")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// After first deletion.
	// {"user":"eco","languages":["go","python","C", "Cpp"]}

	json, err = Delete(json, "user")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"languages":["go","python","C", "Cpp"]}
}

func ExampleInsert() {
	newValue := []byte(`"visual basic"`)
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"]}`)

	json, err := Insert(json, 2, newValue, "languages")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"user":"eco","languages":["go","java","visual basic","python","C","Cpp"]}
}

func ExampleIterateArray() {
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"]}`)

	err := IterateArray(json, func(value []byte) bool {
		fmt.Println(string(value))
		// this return is some kind control mechanism for escape the iteration any time you want.
		// true means keep iterate. false means stop iteration.
		return true
	}, "languages")

	if err != nil {
		fmt.Println(err)
		return
	}
	// Output: go
	//java
	//python
	//C
	//Cpp
}

func ExampleIterateKeyValue() {
	json := []byte(`{"index":42,"user":"eco","language":"go","uuid":"4a1531c25d5ef124295a","active":true}`)

	err := IterateKeyValue(json, func(key, value []byte) bool {
		fmt.Println("key  :", string(key))
		fmt.Println("value:", string(value))
		// this return is some kind control mechanism for escape the iteration any time you want.
		// true means keep iterate. false means stop iteration.
		return true
	})

	if err != nil {
		fmt.Println(err)
		return
	}
	// Output: key  : index
	//value: 42
	//key  : user
	//value: eco
	//key  : language
	//value: go
	//key  : uuid
	//value: 4a1531c25d5ef124295a
	//key  : active
	//value: true
}

func ExampleGetKeys() {
	json := []byte(`{"index":42,"user":"eco","language":"go","uuid":"4a1531c25d5ef124295a","active":true}`)

	keys, err := GetKeys(json)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(keys)
	// Output: [index user language uuid active]
}

func ExampleGetValues() {
	json := []byte(`{"index":42,"user":"eco","language":"go","uuid":"4a1531c25d5ef124295a","active":true}`)

	values, err := GetValues(json)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(values)
	// Output: [42 eco go 4a1531c25d5ef124295a true]
}

func ExampleGetKeysValues() {
	json := []byte(`{"index":42,"user":"eco","language":"go","uuid":"4a1531c25d5ef124295a","active":true}`)

	keys, values, err := GetKeysValues(json)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(keys)
	fmt.Println(values)
	// Output: [index user language uuid active]
	//[42 eco go 4a1531c25d5ef124295a true]
}

func ExampleGetMap() {
	json := []byte(`{"uuid":"4a1531c25d5ef124295a","personal":{"index":42,"user":"eco","language":"go"}}`)

	personal, err := GetMap(json, "personal")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(personal["index"])
	fmt.Println(personal["user"])
	fmt.Println(personal["language"])
	// Output: 42
	//eco
	//go
}

func ExampleGetAll() {
	json := []byte(`{"uuid":"4a1531c25d5ef124295a","personal":{"index":42,"user":"eco","language":"go"}}`)

	keys := []string{"index", "language"}
	allINeed, err := GetAll(json, keys, "personal")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allINeed)
	// Output: [42 go]
}

func ExampleGetAllMap() {
	json := []byte(`{"uuid":"4a1531c25d5ef124295a","personal":{"index":42,"user":"eco","language":"go"}}`)

	keys := []string{"index", "language"}
	allINeed, err := GetAllMap(json, keys, "personal")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(allINeed["index"])
	fmt.Println(allINeed["language"])
	// Output: 42
	//go
}

func ExampleLength() {
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"]}`)

	length, err := Length(json, "languages")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(length)
	// Output: 5
}

func ExampleMakeArray() {
	years := MakeArray(2005, 2009, 2013, 2019)
	// [2005,2009,2013,2019]
	active := MakeArray(false, true, true, true)
	// [false,true,true,true]
	languages := MakeArray("visual-basic", "java", "python", "go")
	// ["visual-basic","java","python","go"]

	all := MakeArrayBytes(years, active, languages)
	fmt.Println(string(all))
	// Output: [[2005,2009,2013,2019],[false,true,true,true],["visual-basic","java","python","go"]]
}

func ExampleMakeJson() {
	keys := []string{"username", "ip", "mac", "active"}
	values := []interface{}{"eco", "192.168.1.108", "bc:ae:c5:13:84:f9", true}

	user := MakeJson(keys, values)

	fmt.Println(string(user))
	// Output: {"username":"eco","ip":"192.168.1.108","mac":"bc:ae:c5:13:84:f9","active":true}
}

func ExampleFlatten() {
	json := []byte(`{
	"user": "eco",
	"languages": [
		"go",
		"java",
		"python",
		"C",
		"Cpp"
	],
	"following": {
		"social": "dev.to",
		"code": "github"
	}
}`)

	json = Flatten(json)
	fmt.Println(string(json))
	// Output: {"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}
}

func ExampleIndent() {
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	json = Indent(json)
	fmt.Println(string(json))
	// Output: {
	//	"user": "eco",
	//	"languages": [
	//		"go",
	//		"java",
	//		"python",
	//		"C",
	//		"Cpp"
	//	],
	//	"following": {
	//		"social": "dev.to",
	//		"code": "github"
	//	}
	//}
}

func ExampleParseArray() {
	arrayStr := `["eco", 1992, 28, false, "github", "jin"]`
	array := ParseArray(arrayStr)
	fmt.Printf("val:%v, type:%T\n", array, array)
	// Output: val:[eco 1992 28 false github jin], type:[]string
}

func ExampleParser_Tree() {
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)
	parse, _ := Parse(json)
	tree := parse.Tree()
	fmt.Println(tree)
	// Output: ─ user
	//─ languages
	//	└  0
	//	└  1
	//	└  2
	//	└  3
	//	└  4
	//─ following
	//	└  social
	//	└  code
}

func ExampleParser_TreeFull() {
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)
	parse, _ := Parse(json)
	tree := parse.TreeFull()
	fmt.Println(tree)
	// Output: ─ user   : "eco"
	//─ languages : ["go","java","python","C","Cpp"]
	//	└  0      : "go"
	//	└  1      : "java"
	//	└  2      : "python"
	//	└  3      : "C"
	//	└  4      : "Cpp"
	//─ following : {"social":"dev.to","code":"github"}
	//	└  social : "dev.to"
	//	└  code   : "github"
}

func ExampleScheme_MakeScheme() {
	// without Scheme
	json := MakeEmptyJson()
	json, _ = AddKeyValueString(json, "name", "eco")
	json, _ = AddKeyValueString(json, "lastname", "hub")
	json, _ = AddKeyValueString(json, "age", "28")
	// json = {"name":"eco","lastname":"hub","age":28}

	// with Scheme
	person := MakeScheme("name", "lastname", "age")
	eco := person.MakeJson("eco", "hub", "28")
	fmt.Println(string(eco))
	// {"name":eco,"lastname":hub,"age":28}

	// And it provides a limitless instantiation
	sheldon := person.MakeJson("Sheldon", "Bloom", "42")
	john := person.MakeJson("John", "Wiki", "28")
	fmt.Println(string(sheldon))
	fmt.Println(string(john))
	// {"name":"Sheldon","lastname":"Bloom","age":42}
	// {"name":"John","lastname":"Wiki","age":28}
	// Output: {"name":eco,"lastname":hub,"age":28}
	//{"name":"Sheldon","lastname":"Bloom","age":42}
	//{"name":"John","lastname":"Wiki","age":28}
}

func ExampleScheme() {
	// This Section provides examples for;
	// Add(), Remove(), Save(), Restore(),
	// GetOriginalKeys(), GetCurrentKeys() functions.

	// MakeScheme need keys for construct a JSON scheme.
	person := MakeScheme("name", "lastname", "age")

	// now we can instantiate a JSON with values.
	eco := person.MakeJson("eco", "hub", "28")

	fmt.Println(string(eco))
	// {"name":"eco","lastname":"hub","age":28}

	// Add simply adds a new key to scheme.
	person.Add("ip")
	person.Add("location")

	// now with addition above we can assign two more value.
	sheldon := person.MakeJson("Sheldon", "Bloom", "42", "192.168.1.105", "USA")

	fmt.Println(string(sheldon))
	// {"name":"Sheldon","lastname":"Bloom","age":42,"ip":"192.168.1.105","location":"USA"}

	fmt.Println(person.GetCurrentKeys())
	// [name lastname age ip location]
	fmt.Println(person.GetOriginalKeys())
	// [name lastname age]

	// for remove a certain key.
	person.Remove("location")

	john := person.MakeJson("John", "Wiki", "28", "192.168.1.102")
	fmt.Println(string(john))
	// {"name":"John","lastname":"Wiki","age":28,"ip":"192.168.1.102"}

	// restores original form of scheme
	person.Restore()
	// and instantiate another
	ted := person.MakeJson("ted", "stinson", "38")
	fmt.Println(string(ted))

	//save saves current keys, now this is original scheme
	person.Save()
	fmt.Println(person.GetCurrentKeys())
	// [name lastname age ip location]
	fmt.Println(person.GetOriginalKeys())
	// [name lastname age]

	//Output: {"name":"eco","lastname":"hub","age":28}
	//{"name":"Sheldon","lastname":"Bloom","age":42,"ip":"192.168.1.105","location":"USA"}
	//[name lastname age ip location]
	//[name lastname age]
	//{"name":"John","lastname":"Wiki","age":28,"ip":"192.168.1.102"}
	//{"name":"ted","lastname":"stinson","age":38}
	//[name lastname age]
	//[name lastname age]
}

func ExampleParser_Get() {
	path := []string{"following", "social"}
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	pars, err := Parse(json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// get() returns whole JSON.
	json, err = pars.Get(path...)
	// or without path variable
	// json, err := Get("following", "social")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: dev.to
}

func ExampleParser_Set() {
	path := []string{"following"}
	newValue := []byte(`["computerphile","numberphile"]`)
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	pars, err := Parse(json)

	err = pars.Set(newValue, path...)
	// or without path variable
	// json, err := Set(newValue, "following")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// get() returns whole JSON.
	json, err = pars.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"user":"eco","languages":["go","java","python","C","Cpp"],"following":["computerphile","numberphile"]}
}

func ExampleParser_SetKey() {
	path := []string{"following"}
	newKey := "sites"
	json := []byte(`{"user":"eco","languages":["go","java","python","C","Cpp"],"following":{"social":"dev.to","code":"github"}}`)

	pars, err := Parse(json)

	err = pars.SetKey(newKey, path...)
	// or without path variable
	// json, err := Set(newKey, "following")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	json, err = pars.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"user":"eco","languages":["go","java","python","C","Cpp"],"sites":{"social":"dev.to","code":"github"}}
}

func ExampleParser_Add() {
	newValue := []byte(`"godoc.org/github.com/ecoshub"`)
	json := []byte(`{"user":"eco","links":["github.com/ecoshub"]}`)

	pars, err := Parse(json)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	err = pars.Add(newValue, "links")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// get() returns whole JSON.
	json, err = pars.Get()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(json))
	// Output: {"user":"eco","links":["github.com/ecoshub","godoc.org/github.com/ecoshub"]}
}
