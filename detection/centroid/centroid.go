// centroid is a simple package to contain a data structure across classes.
package centroid

// Centroid struct holds location data:
// Timestamp (millisecond precision), and X, Y pixel coordinates
type Centroid struct {
	Timestamp int64
	X, Y      int
}
