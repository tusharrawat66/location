package main

import "user_proj/views"

func main() {
	r := views.SetupRouter()
	r.Run(":8081")
}
