// The objective of the Mediator is to have an object responsible for mediating what is happening
// The Mediator has methods that the components can call and vice versa
// In this example we have the stationManager that mediates whether if the train can arrive or not according with the availability
// of the station

package main

func main() {
	stationManager := newStationManger()

	passengerTrain := &passengerTrain{
		mediator: stationManager,
	}
	freightTrain := &freightTrain{
		mediator: stationManager,
	}

	passengerTrain.arrive()
	freightTrain.arrive()
	passengerTrain.depart()
}
