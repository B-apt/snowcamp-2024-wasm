package main

import "github.com/extism/go-pdk"

//export hello
func hello() int32 {
	// read the function argument from the memory
	input := pdk.Input()

	// display the argument
	pdk.Log(pdk.LogInfo, string(input))

	// read the config object
	url, _ := pdk.GetConfig("url")

	// display the value of url
	pdk.Log(pdk.LogInfo, string(url))

	// use the url to make an http request
	req := pdk.NewHTTPRequest(pdk.MethodGet, url)
	res := req.Send()

	// display the response
	pdk.Log(pdk.LogInfo, string(res.Body()))

	// return a message: output := "🎉 Extism is 💜"
	mem := pdk.AllocateString("🎉 Extism is 💜")
	pdk.OutputMemory(mem)
	return 0
}

func main() {}
