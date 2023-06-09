// Package rest is port adapter via http/s protocol
// # This manifest was generated by ymir. DO NOT EDIT.
package rest

import "time"

// Option is server type return func.
type Option = func(s *Server) error

// WithHost will assign to host field server.
func WithHost(host string) Option {
	return func(s *Server) error {
		s.Host = host
		return nil
	}
}

// WithPort will assign to port field server.
func WithPort(port string) Option {
	return func(s *Server) error {
		s.Port = port
		return nil
	}
}

// WithReadTimeout will assign to read field server.
func WithReadTimeout(seconds int) Option {
	return func(s *Server) error {
		s.readTimeout = time.Duration(seconds) * time.Second
		return nil
	}
}

// WithWriteTimeout will assign to write field server.
func WithWriteTimeout(seconds int) Option {
	return func(s *Server) error {
		s.writeTimeout = time.Duration(seconds) * time.Second
		return nil
	}
}

// WithIdleTimeout will assign to idle timeout field server.
func WithIdleTimeout(seconds int) Option {
	return func(s *Server) error {
		s.idleTimeout = time.Duration(seconds) * time.Second
		return nil
	}
}

// WithShutdownTimeout will assign to shut down timeout field server.
func WithShutdownTimeout(seconds int) Option {
	return func(s *Server) error {
		s.shutdownTimeOut = time.Duration(seconds) * time.Second
		return nil
	}
}
