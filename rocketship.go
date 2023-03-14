package main

import "fmt"

// Create the function fuelGauge() here
func fuelGauge(fuel int) int{
  fmt.Println(fuel)
  return fuel
}

// Create the function calculateFuel() here
func calculateFuel(planet string) int{
  var fuel int = 0
  switch planet{
    case "Venus":
      fuel = 300000
    case "Mercury":
      fuel = 500000
    case "Mars":
      fuel = 700000
    default:
      fuel = 0
  }
  return fuel
}

// Create the function greetPlanet() here
func greetPlanet(planet string) {
  fmt.Printf("Welcome to %v!!", planet)
  fmt.Println(" ")
}

// Create the function cantFly() here
func cantFly(){
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(current string, planet string, fuel int) (string, int){
  var fuelRemaining int = fuel
  var fuelCost int = calculateFuel(planet)
  var currentPlanet string = current
  if fuelRemaining > fuelCost{
    greetPlanet(planet)
    fuelRemaining = fuelRemaining - fuelCost
    currentPlanet = planet
  } else{
    cantFly()
  }
  return currentPlanet, fuelRemaining
}
func options() {
  defer fmt.Println("Type one of the shown options to travel.")
  fmt.Println("Welcome to Planet Express. Your options are: ")
  fmt.Println("Venus")
  fmt.Println("Mercury")
  fmt.Println("Mars")
}

func main() {
  /*
  // Test your functions!
  var fuel int = 69
  fuelGauge(fuel)

  var planetChoice string = "Mars"
  fuel = calculateFuel(planetChoice)
  fmt.Println(fuel)

  fuel = 100000
  planetChoice = "Venus"
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
  */

  /*
  keep track of planet
  function to return spaceship back to home planet
  add more destinations and allow transportation between different planets
  add ToLower on function args?

  */

  var fuel int=1000000
  var planetChoice string
  options()
  fmt.Scan(&planetChoice)
  fuelGauge(fuel)
  var currentPlanet string = "Earth"
  currentPlanet, fuel = flyToPlanet(currentPlanet, planetChoice, fuel)
  fuelGauge(fuel)
  fmt.Println(currentPlanet)


}
