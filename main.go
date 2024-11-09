package main

import (
	"context"
	"jane/board"
	"jane/evaluator"
	"jane/paths"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const numberOfSolvers = 100

func pathsGenerator(ctx context.Context, board *board.ChessBoard, pathType paths.PathType, pathsC chan<- paths.Path) {
	generator, err := paths.NewPathGenerator(board, pathType)
	if err != nil {
		panic(err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		default:
			path, notFoundErr := generator.GeneratePaths(pathType)
			if notFoundErr == nil {
				pathsC <- path
			}
		}
	}
}

func solve(ctx context.Context, cb *board.ChessBoard, pathsC chan paths.Path) {
	var wg sync.WaitGroup
	wg.Add(numberOfSolvers)
	evaluatorInstance := evaluator.NewEvaluator(cb)

	for i := 0; i < numberOfSolvers; i++ {
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case path := <-pathsC:
					for _, values := range evaluatorInstance.GetValues() {
						evaluatorInstance.Evaluate(&values, &path, 2024)
					}
				}
			}
		}()
	}

	wg.Wait()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigC
		cancel()
	}()

	pathsC := make(chan paths.Path)

	// instantiate the chess board
	cb, err := board.NewChessBoard([][]string{
		{"A", "A", "A", "B", "B", "C"},
		{"A", "A", "A", "B", "B", "C"},
		{"A", "A", "B", "B", "C", "C"},
		{"A", "B", "B", "C", "C", "C"},
		{"A", "B", "B", "C", "C", "C"},
		{"A", "B", "B", "C", "C", "C"},
	})
	if err != nil {
		panic(err)
	}

	cb.Display()

	// generate moves with different corner-to-corner paths
	go pathsGenerator(ctx, cb, paths.PathA1ToF6, pathsC)
	go pathsGenerator(ctx, cb, paths.PathA6ToF1, pathsC)

	solve(ctx, cb, pathsC)
}
