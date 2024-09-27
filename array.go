package main

import (
	"fmt"
	"log"
	"os/exec"
	"os"
	"math/rand/v2"
	"github.com/eiannone/keyboard"
)
//main function is basically a functioning level. I could copy and paste it and make more levels. that'd be pretty cool, and it'd make it easier for people to make their own levels. although I might be the only person who would.
func main() {

	var chr = "\033[41mX"

	var fuck [20][20]string

	fuck = newMap()
	
	printMap(fuck)
	//var choice string
	var x int
	var y int
	x = 3
	y = 5
	//cX and cY are the customers pos
	cX := 5
	cY := 4
	//drugs sold
	drugSold := 0
	//player money
	money := 0
	//stash is the amount of drugs the player has, the player starts with a q
	stash:= 7

	// Start listening to the keyboard
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer keyboard.Close()

	for {
		
		//fmt.Println("type command up")
		//fmt.Scanln(&choice)
		// Non-blocking listen for keypress
		char, key, err := keyboard.GetKey()
		if err != nil {
			log.Fatal(err)
		} 

		fmt.Println(key)
		switch char {
			

			case 'w':
				x--
				if boundries(x, y) == true {
					x++ 
					continue
				}
				if x == 0 {
					x++
					continue
				}
			case  's':
				x++
				if boundries(x, y) == true {
                                        x--
                                        continue
                                }
				if x == 19 {
                                        x--
                                        continue
                                }

			case 'd':
				y++
				if boundries(x, y) == true {
                                        y--
                                        continue
                                }
				if y == 19 {
                                        y--
                                        continue
                                }

			case 'a':
				y--
				if boundries(x, y) == true {
                                        y++
                                        continue
                                }
				if y == 0 {
                                        y++
                                        continue
                                }

			case 'q': // Quit the game with 'q' or ESC
			return

		}


		

	clear()
		
	fuck = newMap()
	//reminder to make collision work	
	fuck[3][7] = "\033[45m="
	fuck[3][8] = "\033[45m="
        fuck[3][9] = "\033[45m="
	fuck[5][7] = "\033[44mH"
	fuck[5][8] = "\033[49mn" 
	fuck[5][9] = "\033[44mH"
	fuck[4][7] = "\033[44mH"
        fuck[4][8] = "\033[44mH"
        fuck[4][9] = "\033[44mH"

	fuck[cX][cY] = "\033[41mE"
	
	fuck[x][y] = chr
	//print map here
	printMap(fuck)

	//if customer is touching player they'll do a drug deal
	if fuck[cX][cY] == chr {
                //log.Fatal()
		if stash >= 1{

			stash--
			//10 bucks a gram firms
			money += 10
			drugSold++
			cX = rand.IntN(18) + 1
			cY = rand.IntN(18) + 1
			if boundries(cX, cY) == true {
				cX = rand.IntN(18) + 1
      			        cY = rand.IntN(18) + 1
	
			}
		} else {
			fmt.Println("get more dope!")
		}
		
        }
	//printMap(fuck)
	// this is what happens when you enter the house
	if fuck[5][8] == chr {
		//levelTwo()
		if money >= 60 {
			money -= 60
			stash += 7
			
		} else {
			fmt.Println("get more money")
		}
	}

	//printMap(fuck)
	//I fucked up and var x measures the vertical position. var y measures the horizontal position. I switched it in the print statement so it's technically right.
	fmt.Println("Y:", x)
	fmt.Println("X:", y)
	fmt.Println("money:", money)
	fmt.Println("Stash", stash, "grams")
	fmt.Println("Drugs Sold:", drugSold)
	fmt.Println("press Q to quit")
	continue
	}
	


}
//boundries func is called everytime the player moves. if it returns true the player shouldnt move into that space cause its prolly a house.
func boundries(x, y int) bool {
	var penis bool
	switch {
	//here you can add switch case statements to make the player collidee with parts of the grid

	//basically if x is one coordinate of the house and y is also a coordinate of the house, player no move
	case x == 3 && y == 7:
		penis = true
	case x == 3 && y == 8:
		penis = true
	case x == 3 && y == 9:
		penis = true
	case x == 4 && y == 7:
		penis = true
	case x == 4 && y == 8:
		penis = true
	case x == 4 && y == 9:
		penis = true
	case x == 5 && y == 7:
		penis = true
		//x5 y8 at the house is the door
	case x == 5 && y == 8:
                penis = false
	case x == 5 && y == 9:
                penis = true

	}
	return penis 
}

func levelTwo() {
var chr = "8"
        var fuck [20][20]string

        fuck = newMap()

        printMap(fuck)

        var choice string
        var x int
        var y int

        for {
		
                fmt.Println("type command up")
                fmt.Scanln(&choice)

                switch choice {

                        case "up", "w":
                                x--
                        case "down", "s":
                                x++
                        case "left", "a":
                                y--
                        case "right", "d":
                                y++

                        }


        if fuck[5][4] == chr {
                log.Fatal()
        }

        if fuck[5][8] == chr {

        }
	
	clear()
        fuck = newMap()

        fuck[1][5] = "%"
        fuck[1][6] = "%"
        fuck[1][7] = "%"
        fuck[3][5] = "H"
        fuck[3][6] = "n"
        fuck[3][7] = "H"
        fuck[2][5] = "H"
        fuck[2][6] = "H"
        fuck[2][7] = "H"


        fuck[5][4] = "E"

        fuck[x][y] = chr


        printMap(fuck)

        }


}

func clear() {

        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
}


//prints 2d array
func printMap(fuck [20][20]string) {

/*	for d := 0; d <= 9; d++ {

                fmt.Println("", fuck[d], " ", "\n")

        }
	fmt.Println()*/

	for i := 0; i < len(fuck); i++ {    // Loop through each row
    for j := 0; j < len(fuck[i]); j++ { // Loop through each element in the row
        fmt.Print(fuck[i][j], " ")  // Print element with a space
    }
    fmt.Println() // Move to the next line after printing a row
}

}
//THIS IS WHERE YOU CAN CHANGE THE BACKGROUND SHIT
//returns 2d array of all Xs
func newMap() [20][20]string  {
	
	borderColor := "\033[40mX"
	backGroundColor := "\033[42m0"

	var fuck [20][20]string

        for j := 0; j <= 19; j++ {
                for i := 0; i <= 19; i++ {

                        fuck[j][i] = borderColor

                }
        }
	//this part fills the inside with 0s
	for d := 1; d <= 18; d++ {
		for f := 1; f <= 18; f++ {
			fuck[d][f] = backGroundColor
		}
        }



return fuck

}


