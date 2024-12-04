package g2d

import (
	"errors"
	"math"
)

// 计算点到线段的距离
func PointToSegmentDistance(p, a, b Point) float64 {
	// 计算向量 AB 和 AP
	abx := b.X - a.X
	aby := b.Y - a.Y
	apx := p.X - a.X
	apy := p.Y - a.Y

	// 计算向量 AB 的长度的平方
	abSquared := abx*abx + aby*aby

	// 计算投影系数 t
	t := (apx*abx + apy*aby) / abSquared

	// 判断投影点的位置
	if t < 0 {
		// 投影点在 A 之前，距离为点到 A 的距离
		return math.Sqrt((p.X-a.X)*(p.X-a.X) + (p.Y-a.Y)*(p.Y-a.Y))
	} else if t > 1 {
		// 投影点在 B 之后，距离为点到 B 的距离
		return math.Sqrt((p.X-b.X)*(p.X-b.X) + (p.Y-b.Y)*(p.Y-b.Y))
	} else {
		// 投影点在线段上，计算点到投影点的距离
		projX := a.X + t*abx
		projY := a.Y + t*aby
		return math.Sqrt((p.X-projX)*(p.X-projX) + (p.Y-projY)*(p.Y-projY))
	}
}

// 计算两个向量的叉积
func crossProduct(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

// 判断点是否在线段上
func IsPointOnSegment(p, a, b Point) bool {
	// 计算叉积，判断是否共线
	if crossProduct(a, b, p) != 0 {
		return false
	}
	// 判断点是否在线段的范围内
	if p.X < math.Min(a.X, b.X) || p.X > math.Max(a.X, b.X) ||
		p.Y < math.Min(a.Y, b.Y) || p.Y > math.Max(a.Y, b.Y) {
		return false
	}
	return true
}

// 计算两条线段的交点
func GetIntersection(a, b, c, d Point) (Point, bool) {
	// 计算叉积
	c1 := crossProduct(a, b, c)
	c2 := crossProduct(a, b, d)
	c3 := crossProduct(c, d, a)
	c4 := crossProduct(c, d, b)

	// 判断两条线段是否相交
	if c1*c2 <= 0 && c3*c4 <= 0 {
		// 计算交点
		denominator := (b.X-a.X)*(d.Y-c.Y) - (b.Y-a.Y)*(d.X-c.X)
		if denominator == 0 {
			return Point{}, false // 平行或重合的情况
		}
		t := ((a.X-c.X)*(d.Y-c.Y) - (a.Y-c.Y)*(d.X-c.X)) / denominator
		intersectX := a.X + t*(b.X-a.X)
		intersectY := a.Y + t*(b.Y-a.Y)
		intersection := Point{X: intersectX, Y: intersectY}

		// 检查交点是否在线段上
		if IsPointOnSegment(intersection, a, b) && IsPointOnSegment(intersection, c, d) {
			return intersection, true
		}
	}
	// 如果没有交点
	return Point{}, false
}

func SegmentInterpolation(p1, p2 Point, dis float64) (Point, error) {
	// 计算两点之间的欧几里得距离
	distance := math.Sqrt((p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y))

	// 如果两点重合，返回错误
	if distance == 0 {
		return Point{}, errors.New("两点重合，无法进行插值")
	}

	// 计算比例 t
	t := dis / distance

	// 如果 dis 大于两点之间的距离，限制 t 为 1，防止超出线段
	if t > 1 {
		t = 1
		return Point{}, errors.New("距离大于线段长度，放弃插值")
	}

	// 计算插值点的坐标
	x := p1.X + (p2.X-p1.X)*t
	y := p1.Y + (p2.Y-p1.Y)*t

	return Point{X: x, Y: y}, nil
}

// 定义一个小的阈值用于浮点数比较
const epsilon = 1e-9

// 计算并返回两点间按固定间隔距离插值的点
func SegmentInterpolationByDistance(p1, p2 Point, d float64) ([]Point, error) {

	// 计算两点之间的欧几里得距离
	distance := math.Sqrt((p2.X-p1.X)*(p2.X-p1.X) + (p2.Y-p1.Y)*(p2.Y-p1.Y))

	// 如果两点重合，返回错误
	if distance < epsilon {
		return nil, errors.New("两点重合，无法进行插值")
	}

	// 检查 d 是否有效
	if d <= 0 {
		return nil, errors.New("d 必须是正数")
	}

	// 计算可以容纳多少个间隔点
	numIntervals := int(math.Ceil(distance / d))

	// 预先分配足够的空间
	points := make([]Point, 0, numIntervals+1)

	// 计算每个插值点
	for i := 0; i <= numIntervals; i++ {
		// 计算比例 t
		t := float64(i) * d / distance

		// 计算插值点的坐标
		x := p1.X + (p2.X-p1.X)*t
		y := p1.Y + (p2.Y-p1.Y)*t

		// 创建插值点并添加到切片中
		points = append(points, Point{X: x, Y: y})
	}

	// 如果最后一个插值点不是终点 p2，则手动将终点 p2 加入
	if len(points) == 0 || !pointsEqual(points[len(points)-1], p2) {
		points = append(points, p2)
	}

	// 返回所有插值点
	return points, nil
}

// 辅助函数：判断两个点是否相等
func pointsEqual(p1, p2 Point) bool {
	return math.Abs(p1.X-p2.X) < epsilon && math.Abs(p1.Y-p2.Y) < epsilon
}
