package main

import "github.com/aloneen/assign2-ex2-adv/initializers"

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDatabase()
}
func main() {

}
