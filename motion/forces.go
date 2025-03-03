package motion

import (
	blueprintmotion "github.com/TheBitDrifter/blueprint/motion"
	"github.com/TheBitDrifter/blueprint/vector"
)

// Forces is the global forces handler instance
var Forces forcesHandler

// forcesHandler manages force and torque operations on dynamics objects
type forcesHandler struct {
	Generator forcesGenerator
}

// AddForce applies a force vector to the dynamics object
func (forcesHandler) AddForce(dyn *blueprintmotion.Dynamics, force vector.TwoReader) {
	forceConc := vector.Two{
		X: force.GetX(),
		Y: force.GetY(),
	}
	dyn.SumForces = dyn.SumForces.Add(forceConc)
}

// ClearForces resets the accumulated forces to zero
func (forcesHandler) ClearForces(dyn *blueprintmotion.Dynamics) {
	dyn.SumForces = vector.Two{}
}

// AddTorque applies rotational force to the dynamics object
func (forcesHandler) AddTorque(dyn *blueprintmotion.Dynamics, torque float64) {
	dyn.SumTorque = dyn.SumTorque + torque
}

// ClearTorque resets the accumulated torque to zero
func (forcesHandler) ClearTorque(dyn *blueprintmotion.Dynamics) {
	dyn.SumTorque = 0
}

// forcesGenerator creates common physics forces
type forcesGenerator struct{}

// NewGravityForce creates a gravity force vector based on mass and gravity constants
func (forcesGenerator) NewGravityForce(mass, gravity, pixelsPerMeter float64) vector.Two {
	return vector.Two{X: 0.0, Y: mass * gravity * pixelsPerMeter}
}

// NewFrictionForce creates a friction force opposing the direction of movement
func (forcesGenerator) NewFrictionForce(velocity vector.Two, frictionCoefficient float64) vector.Two {
	// Calculate the friction force direction (opposite of velocity vector)
	frictionDirection := velocity.Norm().Scale(-1)
	// Multiply the normalized direction vector by the friction coefficient
	frictionForce := frictionDirection.Scale(frictionCoefficient)
	return frictionForce
}
