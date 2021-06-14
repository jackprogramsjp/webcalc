package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	HTML_1 = `
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<title>Web Calculator</title>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
	
			<style>
				:root {
					--main-body-font-family: Verdana, Geneva, Tahoma, sans-serif;
					--main-button-font-family: arial, sans-serif;
					--main-font-size: xx-large;
				}
	
				body {
					font-family: var(--main-body-font-family);
				}
	
				h1 {
					padding: 1.5rem;
				}
	
				div.calc-container {
					display: grid;
					grid-template-columns: auto auto auto;
					gap: 10px;
				}
	
				button.calc-button {
					font-family: var(--main-button-font-family);
					font-size: var(--main-font-size);
					background-color: lightgray;
					border-radius: 10px;
					padding: 0.5rem;
				}
	
				input[type="submit"] {
					font-family: var(--main-button-font-family);
					font-size: var(--main-font-size);
					background-color: lightgray;
					border-radius: 10px;
				}
	
				p.calc-regular {
					font-size: var(--main-font-size);
				}
	
				input.input-text {
					font-size: var(--main-font-size);
					margin: 1rem;
				}
			</style>
	
			<script type="text/javascript">
				function appendInput(text) {
					var element = document.getElementById("calculator-value");
					element.value += text;
				}
			</script>
		</head>
		<body>
			<h1>Web Calculator</h1>
			<form method="POST" action="/">
				<input name="equation" id="calculator-value" class="input-text" type="text" />
				<input type="submit" value="Enter" />`
	HTML_2 = `
			</form>
			<div class="calc-container">
				<button onclick="appendInput('1')" class="calc-button">1</button>
				<button onclick="appendInput('2')" class="calc-button">2</button>
				<button onclick="appendInput('3')" class="calc-button">3</button>
				<button onclick="appendInput('4')" class="calc-button">4</button>
				<button onclick="appendInput('5')" class="calc-button">5</button>
				<button onclick="appendInput('6')" class="calc-button">6</button>
				<button onclick="appendInput('7')" class="calc-button">7</button>
				<button onclick="appendInput('8')" class="calc-button">8</button>
				<button onclick="appendInput('9')" class="calc-button">9</button>
				<button onclick="appendInput('0')" class="calc-button">0</button>
			</div>
			<br><br>
			<div class="calc-container">
				<button onclick="appendInput('-')" class="calc-button">-</button>
				<button onclick="appendInput('+')" class="calc-button">+</button>
				<button onclick="appendInput('*')" class="calc-button">*</button>
				<button onclick="appendInput('/')" class="calc-button">/</button>
				<button onclick="appendInput('.')" class="calc-button">.</button>
			</div>
		</body>
	</html>`

	HTML_FORMAT = "<p class=\"calc-regular\">%v</p>"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "index.html")
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}

		text := r.FormValue("equation")
		result, err := CalculatorRun(text)

		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		if text == "" {
			fmt.Fprintln(w, HTML_1+fmt.Sprintf(HTML_FORMAT, "0")+HTML_2)
		} else if err != nil {
			fmt.Fprintln(w, HTML_1+fmt.Sprintf(HTML_FORMAT, "Error: "+err.Error())+HTML_2)
		} else {
			fmt.Fprintln(w, HTML_1+fmt.Sprintf(HTML_FORMAT, result)+HTML_2)
		}
	default:
		fmt.Fprintln(w, "Only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/", index)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
