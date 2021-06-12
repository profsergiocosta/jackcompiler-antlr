package main

func main() {
	// Setup the input

	compiler.compile(
		`
	class Main {

		function void f1() {
			return;
		}

		function void main(int x) {
			var Point p;
			var int a, b  ,c ;
			let b = Main.f1();
			let a = p.getX();
			return ;
		}

		
	}
	
	
	`)

}
