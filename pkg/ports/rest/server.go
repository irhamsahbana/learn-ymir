// Package rest is port adapter via http/s protocol
// # This manifest was generated by ymir. DO NOT EDIT.
package rest

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

var (
	// ErrServerNotStarted is define error when server not started.
	ErrServerNotStarted = errors.New("server not started")
	// ErrServerAlreadyStarted is define error when server already started.
	ErrServerAlreadyStarted = errors.New("server already started")
	// ErrServerHandlerNotProvided is define error when server handler not provided.
	ErrServerHandlerNotProvided = errors.New("server handler not provided")
)

// Server is the learnymir http server.
type Server struct {
	errCh chan error
	Host  string
	Port  string

	handler http.Handler
	server  *http.Server
	started bool

	idleTimeout     time.Duration
	readTimeout     time.Duration
	writeTimeout    time.Duration
	shutdownTimeOut time.Duration
}

// NewServer creates a server.
func NewServer(opts ...Option) *Server {
	s := &Server{
		Host:            "localhost",
		Port:            "8080",
		idleTimeout:     30 * time.Second,
		readTimeout:     10 * time.Second,
		writeTimeout:    10 * time.Second,
		shutdownTimeOut: 10 * time.Second,
	}

	for _, opt := range opts {
		err := opt(s)
		if err != nil {
			panic(err)
		}
	}

	return s
}

// Handler will assign http handler.
func (s *Server) Handler(h http.Handler) {
	s.handler = h
}

// ListenAndServe will run the server.
func (s *Server) ListenAndServe() error {
	if s.handler == nil {
		return ErrServerHandlerNotProvided
	}
	if s.started {
		return ErrServerAlreadyStarted
	}
	s.server = &http.Server{
		Addr:         fmt.Sprintf("%s:%s", s.Host, s.Port),
		Handler:      s.handler,
		ReadTimeout:  s.readTimeout,
		WriteTimeout: s.writeTimeout,
		IdleTimeout:  s.idleTimeout,
	}

	go func() {
		s.errCh <- s.server.ListenAndServe()
	}()
	s.started = true
	return nil
}

// Error is return channel for capture error.
func (s *Server) Error() chan error {
	return s.errCh
}

// Stop will close the server.
func (s *Server) Stop() {
	if err := s.server.Close(); err != nil {
		log.Error().Err(err).Msg("Server stopping")
	}
}

// Quite will shutdown the server.
func (s *Server) Quite(ctx context.Context) error {
	if !s.started {
		return ErrServerNotStarted
	}
	// Do not make the application hang when it is shutdown.
	ctxOut, cancel := context.WithTimeout(ctx, s.shutdownTimeOut)
	defer cancel()

	log.Info().Msg("I have to go...")
	log.Info().Msg("Stopping server gracefully")
	if err := s.server.Shutdown(ctxOut); err != nil {
		log.Error().Err(err).Msg("Wait is over due to error")
		if err = s.server.Close(); err != nil {
			log.Error().Err(err).Msg("closing failed")
			return err
		}
	}
	log.Info().Msgf("Stop server at %s", s.server.Addr)
	return nil
}
