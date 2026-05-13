package speed

type Car struct {
	battery, batteryDrain, speed, distance int  
}
type Track struct {
	distance int 
}
func NewCar(speed, batteryDrain int) Car {
	return Car{
		speed:  speed,
		batteryDrain: batteryDrain,
		battery: 100,
		distance: 0,
	}
}
func NewTrack(distance int) Track {
	return Track{
		distance: distance,
	}
}
func Drive(car Car) Car {
	if car.battery < car.batteryDrain { 
		return car
	}
	return Car{
		speed: car.speed,
		batteryDrain: car.batteryDrain,
		battery: car.battery-car.batteryDrain,
		distance: car.distance+car.speed,
	}
}
func CanFinish(car Car, track Track) bool {
	remainingDistance := track.distance - car.distance
	lapsNeeded := (remainingDistance + car.speed - 1) / car.speed
	batteryNeeded := lapsNeeded * car.batteryDrain
    return car.battery >= batteryNeeded
}