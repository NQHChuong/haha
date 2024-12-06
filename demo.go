// // What is Go Lang ?
// //  Go Lang is a strong, statically -typed, and comppiled language
// //  Go lang suooprts goroutines, provides strong security and has standard libraries
// // It was designed to reduce the complexity of managing the infrastructure and codebases within Google

// // -------------------------------------------------------------------------------------------------------

// //	Variables in Golang
// //
// // A variable is a storage area defined to store temporary data Varibale can further be changed according to program needs
// //
// //	Varibales declared without any corresponding valye are considered Zero-valued
// //	.. Using var keyword
// //	.. Using short variable declaration
// // -------------------------------------------------------------------------------------------------------

// // Data Types in Golang
// // Data types specify the type and space allotted to variables for functions
// // There are four categories for Data Types in Golang
// // .. Bool Type
// // .. Numeric Type
// // .. String Type
// // .. Derived Type

// // -------------------------------------------------------------------------------------------------------

// // Operators in Golang
// // AN operator is a specified operation that is performed on the operands
// // The following list describe the different oerators used in Golang
// // .. Arithmetic Operators
// // .. Relational Operators
// // .. Logical Operators
// // .. Bitwise Operators
// // .. Assignment Operators
// // .. Misc Operators

// // -------------------------------------------------------------------------------------------------------
// // Loops in GoLang
// //
// //	A loop is the repetition of a block of code that excutes a control statement specific number of times.

// // --------------------------------------------------------------------------------------------------------
// // Structure of  a Go File
// // 1. Package Declaration
// // 2. import packages
// // 3. Function Declaration
// // 4. Statements and Expressions

// // --------------------------------------------------------------------------------------------------------
// // Variables in Golang
// // In Go, there are different types of varuables that can be declared, for example:
// //  -- int - stores integers such as 122 or -123
// //  -- float32 - stores floating numbers , with decimals, such as  19.99 or -19.99
// //  -- string - stores text, such as " Hello World". string values ar esurrounded by double quotes
// //  -- bool stores values with two states: true or false

// // Creating Variables
// // In Go, thể are two ways to declare a variable:
// // 1. Using the "var" keyword
// // Use the var key word, followed by variable name and type:

// // Syntax
// // var variable_name = valued
// // .. Note: You always have to sepcify either type or value (or both)

// // 2. Ussing the := sign :
// // Syntax
// // variable_name := value
// // .. Note: In this casem you dont need specify the type of the variable becasue go will inferred from the values ( measn that the compiler decides the type of the variable, based on the value)
// // .. Note:It is not possible to declare a variable using := , without assiging a value to it

// // 3. Difference between var and :=
// // Var: - It can be useed inside and outside of functons
// //   - Variable declareation and value assignment can be done separately
// //
// // := : -  It can only be used inside functions
// //   - Varibale declaration and value assignment cannot be done sepately (must be done in the same line)

// //-----------------------------------------------------------------------------------------------------------------
// // Go Variables Naming Rules
// // A variable can have a short name ( like x and y) or more descriptive name ( age,price, carname, etc).
// // Go variable naming rules:
// // . A variable name must start with a letter or an underscore character (_)
// // . A variable name cannot start with a digit
// // . A variable name can only contain alpha-numberic characters and inderscores (a-z, A-Z, 0-9 and _)
// // . Variable names are case-sensitive (age, Age and AGE are three different variables)
// // . There is no linit on the length of the variable name
// // . A variable name cannot contain spaces

// // Multi-World Variable Names
// // Variable names with more than one word can be difficult to read.
// // There are serveral techniques you can use to make them more readable:
// // CAMEL  CASE :
// //  - Each word, except the first, starts with a capital letter
// //  Ex: myVariableName = "John"

// //  PASCAL CASE:
// //  - Each word stars with a captal letter
// //  Ex: MyVariableName ="John"

// //	SNAKE CASE:
// //
// // - Each word is seperated by an underscore character:
// // Ex: my_variable_name = "John"

// // --------------------------------------------------------------------------------------------------------
// // Go Constants
// // Go Constants
// // If a variable should have a fixed value that cannot be changes, you can use the "const" keyword.
// // The const keyword declares the variable as "constant", which means that it is unchanged and read-only.
// // Syntax
// // Ex: const CONSTNAMe type = value
// // Constants Rules
// // . Constants names follow the same naming rules as variables
// // Constant names are usuall written in upercase letters (for easy idenfitication and differentation from variables)
// // Constants can be declared both inside and outsidee of a function
// // Contaims Type
// // There are two types of Constants
// //  . Typed Constants
// //  . Untyped constants
// // Ex (Typed Constants) :
// //  - const A int = 1
// //  const  B string ="B"

// // Ex (Untyed Constants):
// // const A = 1;
// // const B ="Hello B"

// // Note: In this case, the type of the constant is inferred from the value (means the compiler decides the type of the constant, based on the value)
// //       You can not change the value of a constant

// // Multiple Constants
// // You can declare multiple constants in a single line, using the following syntax:
// // const (
// //
// //	A int = 1
// //	C = "Hello"
// //	B = 3.14
// //
// // )

// // --------------------------------------------------------------------------------------------------------
// // Go Output Functions
// //  There are three dunctions to output text:
// //   .Print() - prints text to the console
// //   .Println() - prints text to the console and adds a new line
// //   .Printf() - prints formatted text to the console

// //  * Printf() Function
// //   The Printf() function first format its arguament based on the given formatting verb and then prints thems
// //   Here we wil use two formatting verbs:
// //   . %v is used to the value of arguments
// //   . %T is used to print the type of the arguments

// //   Ex:
// //   var i string = "Hello"
// //   var j int = 10
// //   fmt.Printf("i has value: %v and type: %T\n", i, i)
// //   fmt.Printf("j has value: %v and type: %T\n", j, j)
// //   The output will be:
// //   i has value: Hello and type: string
// //   j has value: 10 and type: int

// // Go Formatting Verbs
// // Formatting Verbs for Printf()
// // Go offers several formatting verbs that can be used with the Printf() fucntion.

// // General Formatting Verbs
// // The following verbs can be used with all data types
// // Verb:
// // %v: Prints the value in the default format
// // %#v : Prints the value in Go-syntax format
// // %T: Orints the type of the value
// // %%: The the % sign
// // Ex:
// // var i = 15.5
// // var txt = "Hello Word"

// // fmt.Printf ("%v\n", i) // 15.5
// // fmt.Printf ("%#v\n", txt) // "Hello World"
// // fmt.Printf ("%T\n", i) // float64
// //---------------------------------------------------------------------------------------------------------

// // Go Data Types:
// //  - Data type is an important concept in programming. Data type specifies the size and type of variabl values.
// //  - Go is statiscally typed, meaning that once a variable type is defined, it can only store data of that type.
// //  Go has three basic data types:
// //   . bool: represents a boolean value and is either true or false
// //   - Numeric: represents integer types, floating point values and complex types
// //   - string: represent a strign value

// //---------------------------------------------------------------------------------------------------------
// // Go Arrays

// // Arrays are used to store multiple values of the same type in a signle variable, instead of declearing seperat variables for each value
// // Declare  an Array
// // In Go, there are two ways to declare an array:
// // 1. With the var keyword:
// // Syntax
// // var array_name = [length]datatype{value1, value2, value3, ...}
// // or
// // var array_name = [...]datatype{value1, value2, value3, ...}

// // 2. With the := sign
// // array_name := [length]datatype{value1, value2, value3, ...}
// // or
// // array_name := [...]datatype{value1, value2, value3, ...}
// // .. *NOTE: The length speciify the number of elements to store in the array. In Go, arrays have a fixed length.
// // The length  ò the aray í either defined by a number or inferred (means that the compiler decides the length of the array, based on the number of values)
// // Ex:
// // varr arr1 = [3]int {1,2,3}
// // arr2 := [5]string {"Hello","World","Go","Programming","Language"}
// //----------------------------------------------------------------------------------------------------------------
// // Go Slieces
// // Slices are similar to arrays, but are more powerful and flexible.
// // Like arrays, slices are also used to store multiple values of the sme type in a single varaible.
// // However, unlike arrays, the length of a slice can grow and shrink as you see it
// // In Go, there are serveral ways to create a slice:
// //  .Using the []datatype{values} format
// //  . Creaet a slice from an array
// //  . Using the make()function

// // Create a Slice With []datatype{values}
// // Syntax
// // slice_name := []datatype{value1, value2, value3, ...}
// // In Go, there are two functions that can be used to return the length and capacity of a slice:
// //  . "Len()" function - return the length of the slice(the number of elements in the slice)
// //  . "Cap()" function - return the capacity of the slice(the number of elements the slice can hold)

// //----------------------------------------------------------------------------------------------------------------
// // Go Operators
// // Operators are used to perform operations on variables and values.
// // The "+" operator is used to add together two values, like in the example below:

// // Arithmetic Operators
// //  + Addition  Adds together two values x+y
// //  - Subtraction Subtracts one value from another
// //  * Multiplication Multiples two values
// //  / Divides one value by another
// //  % Moduls Return the division remainder
// //  ++ Increment Increases the value of a variable by 1
// //  -- Decrement Decreases the value of a variable by 1

// //  // List all assignment operators
// //   =  ,x = 5, x =5
// //   +=, x += 3 , x = x +3
// //   -=, x -= 3, x = x -3
// //   *=, x *= 3, x = x *3
// //   /=, x /= 3, x = x /3
// //   %=, x %= 3, x = x %3
// //   &=, x &= 3, x = x &3
// //   |=, x |= 3, x = x |3
// //   ^=, x ^= 3, x = x ^3
// //   >>=, x >>= 3, x = x >>3
// //   <<=, x <<= 3, x = x <<3

// // Comparison Operators
// // Comparison operators are used to compare two values
// // Note: The return value of a comparison is either true (1) or false (0)
// // In the following example, we use the greater than operator (>) to find out if 5 is greater than 3

// //----------------------------------------------------------------------------------------------------------------
// // // Go Conditional Statements
// // Conditional statements are used to perform different actions based on different conditions
// // In Go,
// // a consitin can be either true or false
// // Less than (<)
// // Less than or equal to (<=)
// // Greater than (>)
// // Greater than or equal to (>=)
// // Equal to (==)
// // Not equal to (!=)

// // Additionally, Go supports the usual logical operatiorS:
// // Logical AND (&&)
// // Logical OR (||)
// // Logical NOT (!)

// // // Go If Statement
// // Use the if statement to specify a block of Go code to be executed if a condition is true

// //----------------------------------------------------------------------------------------------------------------
// // Loop in Go Lang
// // The "for" loop loops through a block of code a specified number of times
// // The "for" loop is the only loop available in Go
// // Go for loop
// // Loops are handy if you want to run the same code over and over again, each time with a differet value
// // Each execution of a loop is called an interation
// // The for loop can take up to the three statements:
// // Syntax:
// // for statment1; statement2; statement3 {
// // 	// Code to be excuted for iteration
// // }
// // statement1 Initialize the loop counter value.
// // statement2: Evaluated for each loop iteration. If it evaluates to TRUEm, the loop continues. If it evaluates to FALSE, the loops ends.
// // statement3: Increases the loop counter value

// // Note: Theses statements

// //----------------------------------------------------------------------------------------------------------------
// // Go Functions
// // A function is a block of statements that can be used repeatedly in a program.
// // A function will not excute automatically when a page loads
// // A function will be executed by a call to the function
// // How to create a function ?
// // To create a function in Go, you need to:
// //  1. Use the "func" keyword
// //  2. Specify a name for the function, followed by parentheses()
// //  3. Finally, add code that defines what the function should do, inside curly braces {}

// // ----------------------------------------------------------------------------------------------------------------
// // Go Function Parameters and Arguments
// // Parameters and arguments
// // Information can be passed to functions as a parameter. Parameters act as variables inside the function.
// // Parameters and their types are specified after the function name, inside the parentheses.
// // You can add as many parameters as you want, just seperate them with a comma.
// // Syntax:
// //
// //	func FunctionName(parameter1 type, parameter2 type, parameter3 type) {
// //		// code
// //	}

// //----------------------------------------------------------------------------------------------------------------
// // Go Function Return Values
// // Return Values
// // If you want to the functuon to return a value, you need to define the data type of the return value (such as int, string, etc) and also use the
// // "return" keyword inside the fuction:
// // Syntax:
// // func FunctionName(param1 type, param2 type) type {
// // 	// code to be excuted
// // 	return output
// // }

// // Function Return Example:
// // Example, myFunction() receives two integers (x and y) and returns their addition (x+y) as integer(int):
// // // You can name the return value:
// //
// //	In Go, you can name the return values of a function.
// //
// // in Go, you also can store return values in a variables
// // In Go, it support multiple returns value
// // Go  Maps
// // Map are used to store data values in key: value pairs
// // Each element in a map is a key: value pair
// // A map is an undored and changeable colletion that does not allow duplicates
// // The length of map is the number of its element. You can find it using the len() function
// // The default value of a map is "nil"
// // Maps hold references to an uderlying hash table
// // Go has multiple ways for creating maps.
// // Create map using "var" and ":="
// // Syntax:
// // var a = map[KeyType]ValueType{key1:value1, key2:value2, key3:value3, ...}
// // b:= map[KeyType]ValueType{key1:value1, key2:value2, key3:value3, ...}
// // Create Map Using "male()" function
// // syntax:
// // var a = make(map[KeyType]ValueType)
// // b:= make(map[KeyType]ValueType)

// package main

// import "fmt"

// //	func myMessage() {
// //		println("Hello World")
// //	}
// // func famiLyName(fName string) {
// // 	fmt.Println("Hello", fName)
// // }
// // func myFuntion(x int, y int) int {
// // 	return x + y
// // }
// // func myFuntion(x int, y int) (output1 int, output2 int, output3 int, output4 float32) {
// // 	output1 = x - y
// // 	output2 = x + y
// // 	output3 = x * y
// // 	output4 = float32(x) / float32(y)
// // 	return
// // }
// func main() {
// 	var a = make(map[string]string)
// 	a["Hello"] = "World"
// 	a["Go"] = "Programming"
// 	a["Language"] = "Go"

// 	b := make(map[int]string)
// 	b[1] = "Hello World"
// 	b[2] = "Go Programming"
// 	b[3] = "Language"

// 	fmt.Println(a)
// 	fmt.Println(b)
// 	// var a = map[string]bool{"1": true, "2": false, "3": true}
// 	// b := map[string]bool{"1": true, "2": false, "3": true}
// 	// fmt.Println(a)
// 	// fmt.Println(b)
// 	// output1, output2, output3, output4 := myFuntion(10, 5)
// 	// var arr []interface{}
// 	// arr = append(arr, output1, output2, output3, output4)
// 	// fmt.Println(arr)
// 	// famiLyName("Chuong")
// 	// famiLyName("Quynh Anh")
// 	// fmt.Println(quote.Go())
// 	// // Declaring a variable
// 	// // var a int = 7
// 	// // var b int = 10
// 	// // fmt.Println("Valuye of a is: ", a)
// 	// // fmt.Println("Value of b is: ", b)

// 	// // var d, e int
// 	// // d, e = a, b
// 	// // fmt.Println("Value of d is: ", d)
// 	// // fmt.Println("Value of e is: ", e)
// 	// // v := 76
// 	// // fmt.Println("Value of v is: ", v)
// 	// // a := 96.28
// 	// // b := 18.09
// 	// // c := a + b
// 	// // fmt.Print(c)

// 	// // str := " Nguyen Quoc Huy Chuong"
// 	// // str1 := " Nguyen Ngoc Quynh Anh"
// 	// // out := str == str1
// 	// // fmt.Println(out)

// 	// a := 96
// 	// b := 18
// 	// // out := a + b
// 	// // out1 := a - b
// 	// // out2 := a * b
// 	// // out3 := a / b
// 	// // fmt.Println(out, out1, out2, out3)

// 	// // out4 := a % b
// 	// // fmt.Println(out4)
// 	// if a != b && a >= b {
// 	// 	fmt.Println("True")
// 	// }

// 	// arro := a & b
// 	// fmt.Println(arro)

// 	// a = b
// 	// fmt.Println(a)
// 	// var arr1 = [3]int{1, 2, 3}
// 	// arr2 := [5]string{"Hello", "World", "Go", "Programming", "Language"}
// 	// fmt.Println(arr1)
// 	// fmt.Println(arr2)

// 	// Create Slice
// 	// mySlice := []int{}
// 	// fmt.Println(mySlice)
// 	// fmt.Println(len(mySlice))
// 	// fmt.Println(cap(mySlice))

// 	// mySlice2 := []string{"a", "b", "c", "d"}
// 	// fmt.Println(mySlice2)
// 	// fmt.Println(len(mySlice2))
// 	// fmt.Println(cap(mySlice2))
// 	// var x = 5
// 	// var y = 3
// 	// fmt.Println(x < y)
// 	// if 5 > 3 {
// 	// 	fmt.Println("Yes that right")
// 	// } else {
// 	// 	fmt.Print("Fuck wrong!!")
// 	// }
// 	// myMessage()

// }
