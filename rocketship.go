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
  fmt.Println("Welcome to %v !", planet)
}

// Create the function cantFly() here
func cantFly(){
  fmt.Println("We do not have the available fuel to fly there.")
}

// Create the function flyToPlanet() here
func flyToPlanet(planet string, fuel int) int{
  var fuelRemaining int = fuel
  var fuelCost int = calculateFuel(planet)
  if fuelRemaining > fuelCost{
    greetPlanet(planet)
    fuelRemaining = fuelRemaining - fuelCost
  } else{
    cantFly()
  }
  return fuelRemaining
}

func main() {
  // Test your functions!
  var fuel int = 69
  //fuelGauge(fuel)
  // Create `planetChoice` and `fuel`
  var planetChoice string = "Mars"
  //fuel = calculateFuel(planetChoice)
  //fmt.Println(fuel)
  // And then liftoff!
  fuel = 100000
  planetChoice = "Venus"
  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)

  //keep track of planet
  //function to return spaceship back to home planet
  //add more destinations and allow transportation between different planets
}
