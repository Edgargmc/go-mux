package main

func main() {
	a := App{}
	a.Initialize(
		"postgres",
		"edgargmc",
		"postgres")

	a.Run(":8010")
}