package main

import (
	"errors"
	"fmt"
	"math"
	"io"
	"time"
	"os"
)

const enHello = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Home, "
const reapCount = 5
const countdownStart = 4
const finalWord = `3
2
1
Go!`

type Rectangle struct {
	Width float64 
	Height float64
}

type Circle struct {
    Radius float64
}

type Triangle struct {
    Base   float64
    Height float64
}

type Shape interface{
	Area() float64
}

type Bitcoin int 

type Wallet struct{
	balance Bitcoin
}

func main(){
	fmt.Println(Hello("Tau", "spanish"))
	fmt.Println(Add(10, 10))

	b := [...]string{"Penn", "Teller", "Bank", "Teller", "Penn"}
	fmt.Println(b)

	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}


func Hello(name string, language string) string{
	if name == ""{
		name = "World"
	}
	return prefixlanguage(language) + name
}


func prefixlanguage(language string) (prefix string){

	switch language {
	case "french":
		prefix = frenchHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = enHello
	}

	return
}


func Add(x, y int) int {
	// output: add X + y
	return x + y
}

// iteration
func Repeat(ch string) string {
	var reap string
	for i := 0; i < reapCount; i++{
		reap = reap + ch
	}
	return reap
}

// arrays & slices
func Sum(nums []int) int{
	sum := 0
	for _, i := range nums {
		sum += i
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
    var sums []int
    for _, numbers := range numbersToSum {
        sums = append(sums, Sum(numbers))
    }

    return sums
}

// structs, methods & interfaces
func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
    return (t.Base * t.Height) * 0.5
}


// pointers
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {

    if amount > w.balance {
        return errors.New("oh no")
    }

    w.balance -= amount
    return nil
}

type Dictionary map[string]string
var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string , error) {
	definition, ok := d[word]
	if !ok {
		return " ", ErrNotFound
	}
	return definition, nil
}
// yX9QPs&ngblMr

func (d Dictionary) Add(word, definition string) {}

// dependency injection
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

// mocking
type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct {}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	fmt.Fprint(out, finalWord)
}
