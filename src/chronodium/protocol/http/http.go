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
package http

import (
	"encoding/json"
	"net/http"

	"chronodium/storage"
)

type httpServer struct {
	repo storage.Repo
}

func Start(repo storage.Repo) {
	s := &httpServer{
		repo: repo,
	}

	http.HandleFunc("/metrics/index.json",
		func(w http.ResponseWriter, r *http.Request) { s.graphiteHandler(w, r) })
	http.HandleFunc("/chrono-ts/query",
		func(w http.ResponseWriter, r *http.Request) { s.queryHandler(w, r) })
	go http.ListenAndServe(":8080", nil)
}

func (s *httpServer) queryHandler(w http.ResponseWriter, r *http.Request) {
	query := &storage.Query{}
	query.ShardKey = r.URL.Query().Get("pk")
	if query.ShardKey == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No primary key specified"))
		return
	}

	s.repo.Query(query)
}

func (s *httpServer) graphiteHandler(w http.ResponseWriter, r *http.Request) {
	metrics, _ := s.repo.GetMetricNames()

	if jsonp := r.URL.Query().Get("jsonp"); jsonp != "" {
		w.Header().Set("Content-Type", "application/javascript")

		w.Write([]byte(jsonp + "("))
		json.NewEncoder(w).Encode(metrics)
		w.Write([]byte(")"))
	} else {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(metrics)
	}
}