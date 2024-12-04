package g2d

import "math"

// 判断射线与线段是否相交
func doIntersect(p, a, b Point) bool {
	// 检查射线与线段是否相交
	if p.Y == a.Y || p.Y == b.Y {
		// 如果射线水平与线段重合，则避免射线与水平线重合
		p.Y += 0.0001 // 偏移
	}
	if p.Y < math.Min(a.Y, b.Y) || p.Y >= math.Max(a.Y, b.Y) {
		return false
	}
	if p.X > math.Max(a.X, b.X) {
		return false
	}
	if p.X < math.Min(a.X, b.X) {
		return true
	}

	// 计算交点
	intersectX := (p.Y-a.Y)*(b.X-a.X)/(b.Y-a.Y) + a.X
	if intersectX > p.X {
		return true
	}
	return false
}

// 判断点是否在多边形内
func IsPointInPolygon(p Point, polygon []Point) bool {
	// 边的数量
	n := len(polygon)
	// 计算交点数
	intersections := 0

	for i := 0; i < n; i++ {
		// 获取当前边的两个端点
		a := polygon[i]
		b := polygon[(i+1)%n]

		// 如果点在边上，直接返回 true
		if IsPointOnSegment(p, a, b) {
			return true
		}

		// 检查射线与边是否相交
		if doIntersect(p, a, b) {
			intersections++
		}
	}

	// 如果交点数是奇数，点在多边形内
	return intersections%2 == 1
}
