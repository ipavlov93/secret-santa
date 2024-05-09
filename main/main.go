package main

import "fmt"

const healthCheckPort = ":8080"

//func main() {
//	//LoadEnvVariableOrFatal("TELEGRAM_API_KEY")
//
//	//  open connection
//	//  close connection
//	//  gracefully
//
//
//	http.HandleFunc("/", HelloWorldHandler)
//	//http.hand("/", HelloWorldHandler)
//	//http.Handle("/", HelloWorldHandler)
//	http.ListenAndServe(healthCheckPort, nil)
//}
//
//func HelloWorldHandler(rw http.ResponseWriter, r *http.Request) {
//	fmt.Fprintln(rw, healthCheckPort)
//	fmt.Fprintln(os.Stdout, healthCheckPort)
//}

func main()  {
	var i, sum int
	fmt.Printf("i = %d\n", i)

	for i != 0 {
		fmt.Println("Enter a number: ")
		fmt.Scan(i)
		sum += i
	}
	fmt.Printf("sum = %d\n", sum)
}