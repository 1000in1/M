package g2d

import (
	"fmt"
	"math"
	"runtime"
	"testing"
)

func Test_PointToSegmentDistance(t *testing.T) {
	// 定义点和线段
	p := Point{3, 4} // 点 P
	a := Point{0, 0} // 线段的起点 A
	b := Point{5, 0} // 线段的终点 B

	// 计算点到线段的距离
	distance := PointToSegmentDistance(p, a, b)
	fmt.Printf("点到线段的最短距离是：%.2f\n", distance)
}

func Test_IsPointInPolygon(t *testing.T) {
	// 定义一个点
	point := Point{3, 3}

	// 定义一个多边形 (五边形)
	polygon := []Point{
		{0, 0},
		{5, 0},
		{5, 5},
		{0, 5},
		{0, 0},
	}

	// 判断点是否在多边形内
	if IsPointInPolygon(point, polygon) {
		fmt.Println("点在多边形内")
	} else {
		fmt.Println("点不在多边形内")
	}
}

func TestSegmentInterpolation(t *testing.T) {
	// 定义测试用例
	testCases := []struct {
		p1      Point
		p2      Point
		dis     float64
		want    Point
		wantErr bool
	}{
		// 测试数据1：正常的插值
		{Point{0, 0}, Point{10, 10}, 5 * math.Sqrt2, Point{5, 5}, false},
		// 测试数据2：dis大于线段长度
		{Point{0, 0}, Point{10, 10}, 15 * math.Sqrt2, Point{}, true},
		// 测试数据3：两点重合
		{Point{5, 5}, Point{5, 5}, 5, Point{}, true},
	}

	for _, tc := range testCases {
		got, err := SegmentInterpolation(tc.p1, tc.p2, tc.dis)
		if (err != nil) != tc.wantErr {
			t.Errorf("SegmentInterpolation() error = %v, wantErr %v", err, tc.wantErr)
		}
		if !tc.wantErr && (got.X != tc.want.X || got.Y != tc.want.Y) {
			t.Errorf("SegmentInterpolation() got = %v, want %v", got, tc.want)
		}
	}
}

func TestSegmentInterpolationByDistance(t *testing.T) {
	// 定义测试用例
	testCases := []struct {
		p1      Point   // 起始点
		p2      Point   // 终止点
		d       float64 // 每个插值点之间的距离
		wantLen int     // 期望返回的插值点数量
		wantErr bool    // 期望是否出错
	}{
		// 测试用例数据
		{Point{0, 0}, Point{5, 5}, 1.0, 6, false},
		{Point{0, 0}, Point{0, 0}, 1.0, 0, true},  // 重合的点，应出错
		{Point{0, 0}, Point{5, 5}, 0.0, 0, true},  // 无效的距离，应出错
		{Point{0, 0}, Point{5, 5}, -1.0, 0, true}, // 负的距离，应出错
	}

	for _, tc := range testCases {
		got, err := SegmentInterpolationByDistance(tc.p1, tc.p2, tc.d)

		if tc.wantErr && err == nil {
			t.Errorf("SegmentInterpolationByDistance(%v, %v, %f) expected error but got none", tc.p1, tc.p2, tc.d)
		}

		if !tc.wantErr && err != nil {
			t.Errorf("SegmentInterpolationByDistance(%v, %v, %f) got unexpected error: %v", tc.p1, tc.p2, tc.d, err)
		}

		if len(got) != tc.wantLen {
			t.Errorf("SegmentInterpolationByDistance(%v, %v, %f) returned %d points, want %d", tc.p1, tc.p2, tc.d, len(got), tc.wantLen)
		}
	}
}

// TestCalculateHeading 测试 CalculateHeading 函数
func TestCalculateHeading(t *testing.T) {
	tests := []struct {
		p1, p2 Point
		want   float64
	}{
		// 添加测试用例
		{Point{0, 0}, Point{1, 0}, 0},             // 期望 0 弧度，即东方向
		{Point{0, 0}, Point{0, 1}, math.Pi / 2},   // 期望 π/2 弧度，即北方向
		{Point{0, 0}, Point{-1, 0}, math.Pi},      // 期望 π 弧度，即西方向
		{Point{0, 0}, Point{0, -1}, -math.Pi / 2}, // 期望 -π/2 弧度，即南方向
		// 可以考虑添加更多的测试用例来覆盖不同的方向和边界条件
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := CalculateHeading(tt.p1, tt.p2)
			if got != tt.want {
				t.Errorf("CalculateHeading(%v, %v) = %v, want %v", tt.p1, tt.p2, got, tt.want)
			}
		})
	}
}

// BenchmarkCalculateHeading 性能测试 CalculateHeading 函数
func BenchmarkCalculateHeading(b *testing.B) {
	// 定义一些测试点
	p1 := Point{0, 0}
	p2 := Point{1, 1}

	// 设置 GOMAXPROCS 为 1，确保测试在单核上运行
	runtime.GOMAXPROCS(1)

	// 使用 b.ResetTimer() 来确保计时器在实际测试开始前重置
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CalculateHeading(p1, p2)
	}

	b.Run()

}
