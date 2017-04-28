package xml

import (
	"fmt"
	"io/ioutil"
)

func ExampleParse() {
	failXML, err := ioutil.ReadFile("../fixuters/xml/response_success.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := Parse(failXML)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success:", resp.Success != nil, resp.Failure != nil)

	failXML, err = ioutil.ReadFile("../fixuters/xml/response_fail.xml")
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err = Parse(failXML)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("failure:", resp.Success != nil, resp.Failure != nil)

	// Output:
	// success: true false
	// failure: false true
}
