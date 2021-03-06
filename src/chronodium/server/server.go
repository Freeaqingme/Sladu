// Chronodium - Keeping Time in Series
//
// Copyright 2016-2017 Dolf Schimmel
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package server

import (
	"chronodium/protocol/graphite"
	"chronodium/protocol/http"
	"chronodium/protocol/influxdb"
	"chronodium/storage"
	"chronodium/storage/redis"
	"chronodium/util/stop"
)

type Server struct {
	config  *Config
	stopper *stop.Stopper

	repo *redis.Redis
}

func NewServer(config *Config, stopper *stop.Stopper) *Server {
	return &Server{
		config:  config,
		stopper: stopper,
	}
}

func (s *Server) Start() error {
	s.repo = redis.NewRedis(&s.config.Redis, s.stopper, s.config.TierSets)

	if s.config.Graphite.Enable {
		graphite := graphite.NewServer(&s.config.Graphite, s.stopper)
		if err := graphite.Start(); err != nil {
			return err
		}
		s.repo.AddSource("graphite", graphite.Metrics())
	}

	if s.config.Influxdb.Enable {
		influxdb := influxdb.NewServer(&s.config.Influxdb, s.stopper)
		if err := influxdb.Start(); err != nil {
			return err
		}

		s.repo.AddSource("influxdb", influxdb.Metrics())
	}
	s.repo.Start()

	http.Start(s.Repo())

	return nil
}

func (s *Server) Repo() storage.Repo {
	return s.repo
}

func (s *Server) Stop() {
	// TODO: Abort all operations
}
