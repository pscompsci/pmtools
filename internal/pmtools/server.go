package pmtools

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (p *pmtools) serve() error {
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 5050),
		Handler: p.routes(),
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		p.logger.PrintInfo("caught signal", map[string]string{
			"signal": s.String(),
		})

		fmt.Println("\ncaught signal", map[string]string{
			"signal": s.String(),
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}

		p.logger.PrintInfo("completing background tasks", map[string]string{
			"addr": srv.Addr,
		})

		p.wg.Wait()
		shutdownError <- nil
	}()

	p.logger.PrintInfo("starting server", map[string]string{
		"env":  p.config.Env,
		"addr": srv.Addr,
	})

	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	p.logger.PrintInfo("stopped server", map[string]string{
		"addr": srv.Addr,
	})

	return nil
}
