package g2d

import "math"

// 定义点结构体
type Point struct {
	X, Y float64
}

// 计算两点之间的 heading（与 X 轴的夹角，单位：弧度）
func CalculateHeading(p1, p2 Point) float64 {
	// 计算线段在 X 和 Y 方向的差值
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y

	// // 使用 atan2 计算角度，返回的角度单位为弧度
	// angle := math.Atan2(dy, dx) * 180 / math.Pi

	// // 将角度调整到 [0, 360) 范围内
	// if angle < 0 {
	// 	angle += 360
	// }

	// return angle

	// 使用 atan2 计算角度，返回的角度单位为弧度
	angle := math.Atan2(dy, dx)
	return angle
}

func CalculateDistance(p1, p2 Point) float64 {
	dx := p2.X - p1.X
	dy := p2.Y - p1.Y
	distance := math.Sqrt(dx*dx + dy*dy)
	return distance
}
