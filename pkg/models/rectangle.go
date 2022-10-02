package models

type Rectangle struct {
	Left_bottom Coordinates `json:"left_bottom"`
	Right_top   Coordinates `json:"right_top"`
}

func NewRectangle(left_bottom Coordinates, right_top Coordinates) Rectangle {
	return Rectangle{
		Left_bottom: left_bottom,
		Right_top:   right_top,
	}
}
