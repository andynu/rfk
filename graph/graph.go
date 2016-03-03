// Example of graph as grid
// with path based edges
package main

import (
  "strconv"
  "fmt"
  "strings"
  "sort"
  "os"
  "github.com/andynu/rfk/server/observer"
  "github.com/andynu/rfk/server/library"
)

func main(){
  observer.Observe("library.loaded", func(msg interface{}){
    fmt.Println("initializing graph building")
    // Load unique paths
    pathsm := make(map[string]bool)
    for _, song := range library.Songs {
      path_parts := strings.Split(song.Path, "/")
      plen := len(path_parts)
      partial_path := ""
      for p, path_part := range path_parts {
        if p > 0 {
          partial_path += "/"
        }
        if p == plen - 1 {
          break
        }
        partial_path += path_part
        pathsm[partial_path] = true
      }
    }

    var paths []string
    for path := range pathsm {
      paths = append(paths, path)
    }
    sort.Strings(paths);
    fmt.Printf("paths: %d\n", len(paths))

    graph := make([][]int, len(library.Songs))
    for i := range library.Songs {
      graph[i] = make([]int, len(paths))
    }

    fmt.Println("connecting edges")
    m := 0
    c := 0
    // Connect edges
    for _, song := range library.Songs {
      // find connected songs
      c += 1
      for k, path := range paths {
        if strings.HasPrefix(song.Path, path) {
          graph[songIdx(song)][k] += 1
          m += 1
        }
      }
    }
    fmt.Printf("consider: %d\n", c)
    fmt.Printf("matching: %d\n", m)

    // output graph
    filename := "graph.csv"
    f, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
      panic(err)
    }
    defer f.Close()
    for row := 0; row < len(library.Songs); row += 1 {
      write(f, strconv.Itoa(row))
      for col := 0; col < len(paths); col += 1 {
        //if graph[row][col] != 0 {
        //  fmt.Printf("%d - %d : %d", row, col, graph[row][col])
        //}
        write(f, ",")
        write(f,strconv.Itoa(graph[row][col]))
      }
      write(f, "\n")
    }
    fmt.Println("done")

  })

  library.Load()
}

func write(f *os.File, text string){
  if _, err := f.WriteString(text); err != nil {
    panic(err)
  }
}


func songIdx(song *library.Song) int {
  for i, libSong := range library.Songs {
    if libSong == song {
      return i
    }
  }
  return -1
}
