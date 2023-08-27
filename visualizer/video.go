package visualizer

import (
	"fmt"
	"image/color"
	"maze-solver/io/writer"
	"maze-solver/maze"
	"os"
	"os/exec"
	"path"
	"sync"

	"github.com/mazznoer/colorgrad"
)

type VideoVisualizer struct {
	Filename   string
	Framerate  float64
	ffmpeg_cmd string
}

func (v *VideoVisualizer) Init(*maze.Maze) {
	path, err := exec.LookPath("ffmpeg")
	if err != nil {
		panic(err)
	}
	v.ffmpeg_cmd = path
}
func (v *VideoVisualizer) Run(lets_go chan<- bool) { lets_go <- true }

func (v *VideoVisualizer) Visualize(solved_chan <-chan *maze.SolvedMaze) {
	tmp_dir, err := os.MkdirTemp("", "maze-solver-go-")
	defer os.RemoveAll(tmp_dir)

	if err != nil {
		panic(err)
	}
	var wg sync.WaitGroup
	i := 0
	for solved := range solved_chan {
		wg.Add(1)
		go func() {
			img_writer := writer.ImageWriter{
				Filename:         path.Join(tmp_dir, fmt.Sprintf("%07v.png", i)),
				Maze:             solved,
				CellWidth:        5,
				CellHeight:       5,
				WallColor:        color.Black,
				PathColor:        color.White,
				SolutionGradient: colorgrad.Warm(),
			}
			img_writer.Write()
			wg.Done()
		}()
		i++
	}
	wg.Wait()
	cmd := exec.Command(
		v.ffmpeg_cmd,
		"-y",
		"-pattern_type", "glob",
		"-i", path.Join(tmp_dir, "*.png"),
		"-r", fmt.Sprint(int(v.Framerate)),
		v.Filename,
	)
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
