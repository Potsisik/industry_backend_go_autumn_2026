package main

import (
	"context"
	"errors"
	"sync"
)

type Result[R any] struct {
	Value R
	Err   error
}

var ErrInvalidWorkers = errors.New("workers must be positive")

func ParallelMap[T any, R any](ctx context.Context, workers int, in []T, fn func(context.Context, T) (R, error)) ([]Result[R], error) {

	// проверка количества работников и контекста
	if workers <= 0 {
		return nil, ErrInvalidWorkers
	}
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	jobs := make(chan int) // создание канала для передачи индексов входного массива
	results := make([]Result[R], len(in))
	var wg sync.WaitGroup //тут я импортировала новый пакет для синхронизации завершения горутин

	// запуск работников
	for w := 0; w < min(workers, len(in)); w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//берем число из канала
			for index := range jobs {
				value, err := fn(ctx, in[index])
				if err != nil {
					results[index].Err = err
				} else {
					results[index].Value = value
				}
			}
		}()
	}

	// отправка индексов входного массива в канал
	for index := range in {
		select {
		case jobs <- index:

		case <-ctx.Done():
			// обработка отмены контекста
			close(jobs)
			wg.Wait()
			return nil, ctx.Err()
		}
	}
	close(jobs)
	wg.Wait()

	return results, nil
}
