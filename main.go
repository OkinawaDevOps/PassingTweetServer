package main

import (
	h "./httphandler"
)

func main() {
	h.StartServer("/", ":3000")
}
