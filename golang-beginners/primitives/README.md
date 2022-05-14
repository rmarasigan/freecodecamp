<style>
   /* Style the list */
   ul.breadcrumb {
      padding: 10px 16px;
      list-style: none;
      background-color: #FACC15;
   }

   /* Display list items side by side */
   ul.breadcrumb li {
      display: inline;
      font-size: 18px;
   }

   /* Add a slash symbol (/) before/behind each list item */
   ul.breadcrumb li+li:before {
      padding: 8px;
      color: black;
      content: "/\00a0";
   }

   /* Add a color to all links inside the list */
   ul.breadcrumb li a {
      color: #000000;
      font-weight: bold;
      text-decoration: none;
   }

   /* Add a color on mouse-over */
   ul.breadcrumb li a:hover {
      color: #01447E;
      text-decoration: none;
   }
</style>
<ul class="breadcrumb">
   <li>
      <a href="https://github.com/rmarasigan/freecodecamp#golang-tutorial-for-beginners">Golang Beginners</a>
   </li>
   <li>
      <a href="#">Primitives</a>
   </li>
</ul>
<br>

# Primitives

### Agenda
* Boolean type
* Numeric types
   * Integers
   * Floating point
   * Complex numbers
* Text types

## Boolean
It represents two states, you either have `true` or you have `false`.

```go
func main() {
   var n bool = true
   fmt.Println("%v, %T\n", n, n) // true, bool
}
```

There's a couple of uses for Boolean variables in our applications, perhaps one of the most common is to use them as state flags. The more common case for using boolean in Go is as a result of logical tests. Now, the `==` operatore is called *equals operator* and that's basically checking to see if the item on the left is equivalent to the item on the right.
```go
func main() {
   n := 1 == 1
   m := 1 == 2
   fmt.Println("%v, %T\n", n, n) // true, bool
   fmt.Println("%v, %T\n", m, m) // false, bool
}
```

The other thing that's important to know about the primitives is that every time you initialize a variable in Go, it actually has a zero value. The zero value for Boolean is false.
```go
func main() {
   var n bool
   fmt.Println("%v, %T\n", n, n) // false, bool
}
```

## Numeric types
The zero value for all numeric types is going to be zero (0) or the equivalent of zero for that numeric type.

### Integer

#### Signed integers
* `int`
   - integer of unspecified size
   - one thing that you're guaranteed is regardless of your environment, it will be at least 32 bits, but it could be 64 or even 128 bits depending on the system that you're running on
   - this will be the default integer type
   ```go
   func main() {
      n := 42
      fmt.Println("%v, %T\n", n, n) // 42, int
   }
   ```

   **Other types**:
   <table>
      <tbody>
         <tr>
            <td>int8</td>
            <td> -128 — 127 </td>
         </tr>
         <tr>
            <td>int16</td>
            <td> -32, 768 — 32, 767 </td>
         </tr>
         <tr>
            <td>int32</td>
            <td> -2, 147, 648 — 2, 147, 647 </td>
         </tr>
         <tr>
            <td>int64</td>
            <td> -9, 223, 372, 036, 854, 775, 808 — 9, 223, 372, 036, 854, 775, 807 </td>
         </tr>
      </tbody>
   </table>

#### Unsigned integers
* `uint`
   ```go
   func main() {
      var n uint16 = 42
      fmt.Println("%v, %T\n", n, n) // 42, uin16
   }
   ```

   **Other types**:
   <table>
      <tbody>
         <tr>
            <td>uint8</td>
            <td> 0 — 255 </td>
         </tr>
         <tr>
            <td>uint16</td>
            <td> 0 — 65, 535 </td>
         </tr>
         <tr>
            <td>uint32</td>
            <td> 0 — 4, 294, 295 </td>
         </tr>
      </tbody>
   </table>

#### Byte
* `byte`
   - an alias for an eight bit unsigned integer and the reason we have that is because the unsigned eight bit integer is very common, because that's what a lot of data streams are used to encode their data

#### Arithmetic Operations
```go
func main() {
   a := 10
   b := 3
   fmt.Println(a + b) // Output: 13
   fmt.Println(a - b) // Output: 7
   fmt.Println(a * b) // Output: 30
   fmt.Println(a / b) // Output: 3; This is called integer division
   fmt.Println(a % b) // Output: 1; This returns the remainder

   var x int = 10
   var y int8  = 3
   fmt.Println(x + y) // You will get an error of `mismatched type int and int8`. To make this work, you need to do a type conversion `int(b)`
}
```

#### Bit Operations
* AND means 1 if the corresponding bits of two operands is 1. If eaither bit of an operand is 0, the result of the corresponding bit is evaluated to 0.
* OR means 1 if at least one corresponding bit of two operands is 1
* Exclusive OR means 1 if the corresponding bits of two operands are opposite
* AND NOT means 1 if neither one of the numbers have the bit set
   ```go
   func main(){
      a := 10 // 1010
      b := 3  // 0011

      fmt.Println(a & b) // Output: 2 -> 0010
      fmt.Println(a | b) // Output: 11 -> 1011
      fmt.Println(a ^ b) // Output: 9 -> 1001
      fmt.Println(a &^ b) // Output: 8 -> 0100
   }
   ```
   > Detailed bitwise operation
   >```
   >   a = 10 = 1010
   >   b =  3 = 0011
   >   
   >   Bit & (AND) operation
   >      1010
   >   &  0011
   >   ________
   >      0010 = 2 in decimal
   >   
   >   Bit | (OR) operation
   >      1010
   >   |  0011
   >   ________
   >      1011 = 11 in decimal
   >   
   >   Bit ^ (Exclusive OR) Operation
   >      1010
   >   ^  0011
   >   ________
   >      1001
   >
   >   Bit &^ (AND NOT) Operation
   >      1010
   >   &^ 0011
   >   ________
   >      0100
   >```

#### Bit shifting
* The first statement means, it is going to bit shift a left three places. The second statement means, going a bit shift a right three places. When we do bit shifting, we're basically adding to that exponent, as long as we're dealing with the power of two.
   ```go
   func main() {
      a := 8
      fmt.Println(a << 3) // Output: 63; (2^3) * (2^3) = 2^6 = 64
      fmt.Println(a >> 3) // Output: 1; (2^3) / (2^3) = 2^0 = 1
   }
   ```
   > Detailed bit shift operations
   > ```
   >    a = 8 = 1000
   >
   >    // Bit shift to the left
   >    a << 3 = 0100 0000
   >
   >    // Bit shift to the right
   >    a >> 3 = 0000 0001
   > ```


### Float
In order to store decimal numbers in Go, we're going to use the floating point numbers. The floating point numbers in Go follow IEEE 754 Stanndard. Two types under IEEE 754 standard:
* float32
   * You can store ± 1.18E-38 — ± 3.4E38
* float64
   * You can store ± 2.23E-308 — ± 1.8E308

#### Defining floating point literal
The `13.7e72` uses an exponential notation. So we can use 13.7 * 10^72 and that's going to use the short form IEEE 72nd to stand for that 10 to the 72nd. If you're going to use the initializer syntax on decimal, it's always going to be initialized to a `float64`. Keep in mind that you can't do arithmetic operations between float64 and float32.

```go
func main() {
   n := 3.14
   n = 13.7e72
   n = 2.1E14
   fmt.Printf("%v, %T", n, n) // 2.1e+14, float64
}
```

#### Arithmetic Operations
The modulo or remainder operator is only available on the integer types. Further, we don't have the bitwise operators or the bit shifting operators.

```go
func main() {
   a := 10.2
   b := 3.7
   fmt.Println(a + b) // Output: 13.899999999999999
   fmt.Println(a - b) // Output: 6.499999999999999
   fmt.Println(a * b) // Output: 37.74
   fmt.Println(a / b) // Output: 2.7567567567567566
}
```

### Complex
Complex numbers are actually treated as a first class citizen, and it opens up go to be used as a very powerful language for data science.

Two types of complex numbers:
* complex64
   * Taking a float32 + float32 for the real and imaginary part
* complex128
   * Taking a float64 + float64 for the real and imaginary part

The `i` literal means an imaginary number
```go
func main() {
  var n complex64 = 1 + 2i
  fmt.Printf("%v, %T\n", n, n) // Output: (1+2i), complex64
}
```

#### Arithmetic Operations
The modulo or remainder operator is only available on the integer types. Further, we don't have the bitwise operators or the bit shifting operators.

```go
func main() {
   a := 1 + 2i
   b := 2 + 5.2i
   fmt.Println(a + b) // Output: (3+7.2i)
   fmt.Println(a - b) // Output: (-1-3.2i)
   fmt.Println(a * b) // Output: (-8.4+9.2i)
   fmt.Println(a / b) // Output: (0.3994845360824742-0.038659793814433i)
}
```

What happens if I need to get the real part or the imaginary part? If you run this on the `complex64`, then the real and the imaginary function are going to give you `float32`. We see that we get `float32`.
```go
func main() {
  var n complex64 = 1 + 2i
  fmt.Printf("%v, %T\n", real(n), real(n)) // Output: 1, float32
  fmt.Printf("%v, %T\n", imag(n), imag(n)) // Output: 2, float32
}
```

If you're working along in your program and all of a sudden you need to make a complex number. This is how you can do it.
```go
func main() {
  var n complex128 = complex(5, 12)
  fmt.Printf("%v, %T\n", n, n) // Output: (5+12i), complex128
}
```

## Text types

### String
A string in Go stands in any UTF-8 character so that makes it powerful but that means that strings cannot encode every type of character that's available.

```go
func main() {
   s := "this is a string"
   fmt.Printf("%v, %T\n", s, s) // Output: this is a string, string
}
```

I can actually treat it sort of like an `array`. I can actually treat the string of text as a collection of letters. So `s[2]` is asking for the third letter, because arrays in Go are zero-based so it starts with 0, 1, 2. For the output, what's happening is that  strings are actually aliases for bytes. To make it print as string, you can add `string` function.
```go
func main() {
   s := "this is a string"
   fmt.Printf("%v, %T\n", s[2], s[2]) // Output: 105, uint8
   fmt.Printf("%v, %T\n", string(s[2]), s[2]) // Output: i, uint8
}
```

Now, strings are generally immutable. So while I can inspect the second character, I can't do something like this. It will return an error of `cannot assign to s[2]` and `cannot use "u" (type string) as type byte in assignment`.
```go
func main() {
   s := "this is a string"
   s[2] = "u"
   fmt.Printf("%v, %T\n", s, s)
}
```

#### Arithmetic Operation
There is one arithmetic or pseudo arithmetic operation that we can do with strings and that is **string concatenation**. In simple terms, we can add strings together. If you run the program, it just merges all the strings together and it gives us the result.
```go
func main() {
   s := "this is a string"
   s2 := "this is also a string"
   fmt.Printf("%v, %T\n", s + s2, s + s2) // Output: this is a stringthis is also a string, string
}
```

Another thing with strings is that we can actually convert them into a collections of bytes, which is called a *slice of bytes*. After running the program, we actually get a string comes out as the ASCII values or the UTF-8 values for each character in that string.
```go
func main() {
   s := "this is a string"
   b := []byte(s) // Converting into a collection of bytes
   fmt.Printf("%v, %T\n", b, b) // Output: [116 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
```

Why would you use this one? A lot of the functions that we're going to use in Go actually work with byte slices and that makes them much more generic and much more flexible than if we work with hard coded string.

### Rune
A rune is a little bit different than a string type. A string type represent any UTF-8 character, a rune represents any UTF-32 character. UTF-32 is a little bit of a weird because while any character in UTF-32 can be up to 32 bits long, it doesn't have to be 32 bits long. For example, any UTF-8 character, which is 8 bits long, is a valid UTF-32 character.


If we're declaring a string we use a `"` double quotes. When we're declaring a single rune, we use `'` single quotes. But if you run this application, it will return `97, int32`. The reason for this is because runes are just type alias for int32. So strings can be converted back and forth between collection of bytes. With runes, they are a true type alias. When you are talking about rune, it is the same as talking about an int32.
```go
func main() {
   r := 'a'
   // This will still give you the same result
   // var r rune = 'a'
   fmt.Printf("%v, %T\n", r, r) // Output: 97, int32
```

## Summary
* **Boolean type**
   * Values are true or false
   * Not an alias for other types (e.g. int)
   * Zero value is false
* **Numeric types**
   * Integer
      * Signed integers
         * `int` type has varying size, but min 32 bits
         * 8 bits (int8) through 64 bit (int64)
      * Unsigned integers
         * 8 bit (byte and uint8) through 32 bit (uint32)
      * Arithmetic operations
         * Addition, subtraction, multiplication, division, remainder
      * Bitwise operations
         * AND &, OR |, XOR ^, AND NOT &^
      * Zero value is 0
      * Can't mix types in same family (uint16 + uint32 = error)
   * Floating point numbers
      * Follow IEEE 754 Standard
      * Zero value is 0
      * 32 and 64 bit versions
      * Literal styles
         * Decimal (3.14)
         * Exponential (13e18 or 2E10)
         * Mixed (13.7e12)
      * Arithmetic operations
         * Addition, subtraction, multiplication, division
   * **Complex numebers**
      * Zerp value is (0 + 0i)
      * 64 and 128 bit versions
      * Built-in functions
         * `complex` - make complex number from two floats
         * `real` - get real part as float
         * `imag` - get imaginary part as float
      * Arithmetic operations
         * Addition, subtraction, multiplication, division
   * **Text types**
      * Strings
         * UTF-8
         * Immutable
         * Can be concatenated with plus (`+`) operator
         * Can be converted to []byte
      * Rune
         * UTF-32
         * Special methods normally required to process
            * e.g. `strings.Reader#ReadRune`