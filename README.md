# maze-solver-go

## Introduction

This is a simple little maze solver made for fun and practise writing code in
golang. This project is a complete re-write of another maze solver I've written
in Java back in 2018 after the first semester of uni at EPFL. Needless to say,
this version is _way_ better as I have written it now that I have my Bachelor's
degree in Computer Science.

### Goal of the project

The goal of this side project was to deepen my understanding of path-finding
algorithms, together with trying to create a good design, focusing on
dependency injection and unit testing. It was very instructive.

Not to brag or anything, but the design was quite good, because after
completing the following steps of the project:

- reading a maze;
- solving it;
- writing the solution to the file-system;

an idea came to my mind, add the feature of visualizing the progress of the
solving algorithm by opening a window that displays the current solution. Well,
each module (readers, writers and solvers) were design to be as decoupled as
possible, and it allowed me to implement the `visualizer` feature in only a
couple of hours (which I was honestly not expecting).

## Dependencies

- `go >= 1.21`
- `ffmpeg` (optional, for video visulazation)

## Usage

After downloading `maze-solver` from the
[assets of the latest release](releases/latest "Latest release"), you can use
it with the following arguments

| Short | Long              | Default          | Description                                                                                                                                                                         |
| ----- | ----------------- | ---------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| -h    | --help            |                  | Print help information                                                                                                                                                              |
| -v    | --verbose         | 0                | Verbose level of the solver                                                                                                                                                         |
| -i    | --input           | `maze.png`       | Input file                                                                                                                                                                          |
| -o    | --output          | `maze_sol.png`   | Output file                                                                                                                                                                         |
|       | --path-char-in    | `' '`            | Character to represent the path in an input text file.                                                                                                                              |
|       | --wall-char-in    | `'#'`            | Character to represent the wall in an input text file.                                                                                                                              |
|       | --path-char-out   | `' '`            | Character to represent the path in an output text file.                                                                                                                             |
|       | --wall-char-out   | `'#'`            | Character to represent the wall in an output text file.                                                                                                                             |
|       | --cell-size-in    | 3                | Size of a cell (in pixels) for input file of image type.                                                                                                                            |
|       | --cell-size-out   | 3                | Size of a cell (in pixels) for output file of image type.                                                                                                                           |
| -a    | --algo            | a-star           | Algorithm to solve the maze, available options: dfs, bfs, dijkstra, a-star.                                                                                                         |
|       | --visualize       |                  | Visualizer the progress of the solver, available options: video, window.<br> Window will give a live feed of the solver, whereas video creates a video --output with mp4 extension. |
|       | --video-name      | `'maze_sol.mp4'` | Name of the output file if --visualize is set to 'video'.                                                                                                                           |
|       | --video-framerate | 60               | Framerate of the video if --visualize is set to 'video'.                                                                                                                            |

For the verbose level of the solver

| Level | Description                                                               |
| ----- | ------------------------------------------------------------------------- |
| 0     | nothing printed to stdout                                                 |
| 1     | print the total time taken by the solver (time of the main() function)    |
| 2     | prints the time the solving algorithm took to run                         |
| 3     | prints the time taken by each section (reader, solving algorithm, writer) |

<!-- ## Examples -->
