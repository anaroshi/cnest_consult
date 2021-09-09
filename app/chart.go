package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// 2021 대학 내신평균
// diamond
func generateScatterItems(data []float32) []opts.ScatterData {
	items := make([]opts.ScatterData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.ScatterData{
			Value:        data[i],
			Symbol:       "diamond",
			SymbolSize:   10,
			SymbolRotate: 0,
		})
	}
	return items
}

func scatterShowLabel(SubjNm []string, AvgValue, StdValue []float32) *charts.Scatter {
	scatter := charts.NewScatter()
	scatter.SetGlobalOptions(charts.WithTitleOpts(
		opts.Title{
			Title: "지원 가능 대학 분포도",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: "대학학과",
			SplitLine: &opts.SplitLine{
				Show: true,
			},
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: "석차등급",
			SplitLine: &opts.SplitLine{
				Show: true,
			}}),
			charts.WithDataZoomOpts(opts.DataZoom{
				Type:       "inside",
				Start:      50,
				End:        100,
				XAxisIndex: []int{0},
			}),
			charts.WithDataZoomOpts(opts.DataZoom{
				Type:       "slider",
				Start:      50,
				End:        100,
				XAxisIndex: []int{0},
			}),
	)

	scatter.SetXAxis(SubjNm).
		AddSeries("Category A", generateScatterItems(AvgValue)).		
		SetSeriesOptions(charts.WithLabelOpts(
			opts.Label{
				Show:     true,
				Position: "top",
				Color: "black",
			}),
		)
	
	scatter.Overlap(line(StdValue)) // 학생점수
	return scatter
}

// 학생 점수
// line
func generateLineData(data []float32) []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < len(data); i++ {
		items = append(items, opts.LineData{Value: data[i]})
	}
	return items
}

func line(StdValue []float32) *charts.Line {
	line := charts.NewLine()

	line.AddSeries("slice", generateLineData(StdValue),
			charts.WithLabelOpts(opts.Label{Show: false, Position: "bottom"})).
		SetSeriesOptions(
			charts.WithMarkLineNameTypeItemOpts(opts.MarkLineNameTypeItem{
				Name: "Average",
				Type: "average",
			}),
		)
	return line
}



// create a chart 
func chart(w http.ResponseWriter, No []string, SubjNm []string, AvgValue[]float32, StdValue []float32) {

	page := components.NewPage()

	// SubjNm   = []string{"Swimming", "Surfing", "Shooting ", "Skating", "Wrestling", "Diving", "Wrestling", "Diving", "Skating"}
	// AvgValue = []float32{2.9, 3.28, 2.73, 2.63, 3.39, 4.94, 4.27, 3.6, 1.85}
	// StdValue  = []float32{3, 3, 3, 3, 3, 3, 3, 3, 3}
	fmt.Println("No:", No)
	fmt.Println("SubjNm:", SubjNm)
	fmt.Println("AvgValue:", AvgValue)
	fmt.Println("StdValue:", StdValue)

	page.AddCharts(
		scatterShowLabel(No, AvgValue, StdValue),		
	)
	// f, err := os.Create("./templates/chart.gohtml")
	// if err != nil {
	// 	panic(err)
	// }
	page.Render(io.MultiWriter(w))
}
