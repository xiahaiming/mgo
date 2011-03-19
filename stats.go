/*
mgo - MongoDB driver for Go

Copyright (c) 2010-2011 - Gustavo Niemeyer <gustavo@niemeyer.net>

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

    * Redistributions of source code must retain the above copyright notice,
      this list of conditions and the following disclaimer.
    * Redistributions in binary form must reproduce the above copyright notice,
      this list of conditions and the following disclaimer in the documentation
      and/or other materials provided with the distribution.
    * Neither the name of the copyright holder nor the names of its
      contributors may be used to endorse or promote products derived from
      this software without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR
CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL,
EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO,
PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF
LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING
NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
*/

package mgo

import (
	"sync"
)


var stats *Stats
var statsMutex sync.Mutex

func CollectStats(enabled bool) {
	statsMutex.Lock()
	if enabled {
		if stats == nil {
			stats = &Stats{}
		}
	} else {
		stats = nil
	}
	statsMutex.Unlock()
}

func GetStats() (snapshot Stats) {
	statsMutex.Lock()
	snapshot = *stats
	statsMutex.Unlock()
	return
}

func ResetStats() {
	statsMutex.Lock()
	old := stats
	stats = &Stats{}
	// These are absolute values:
	stats.SocketsInUse = old.SocketsInUse
	stats.SocketsAlive = old.SocketsAlive
	stats.SocketRefs = old.SocketRefs
	statsMutex.Unlock()
	return
}

type Stats struct {
	MasterConns  int
	SlaveConns   int
	SentOps      int
	ReceivedOps  int
	ReceivedDocs int
	SocketsAlive int
	SocketsInUse int
	SocketRefs   int
}

func (stats *Stats) conn(delta int, master bool) {
	if stats != nil {
		statsMutex.Lock()
		if master {
			stats.MasterConns += delta
		} else {
			stats.SlaveConns += delta
		}
		statsMutex.Unlock()
	}
}

func (stats *Stats) sentOps(delta int) {
	if stats != nil {
		statsMutex.Lock()
		stats.SentOps += delta
		statsMutex.Unlock()
	}
}

func (stats *Stats) receivedOps(delta int) {
	if stats != nil {
		statsMutex.Lock()
		stats.ReceivedOps += delta
		statsMutex.Unlock()
	}
}

func (stats *Stats) receivedDocs(delta int) {
	if stats != nil {
		statsMutex.Lock()
		stats.ReceivedDocs += delta
		statsMutex.Unlock()
	}
}

func (stats *Stats) socketsInUse(delta int) {
	if stats != nil {
		statsMutex.Lock()
		stats.SocketsInUse += delta
		statsMutex.Unlock()
	}
}

func (stats *Stats) socketsAlive(delta int) {
	if stats != nil {
		statsMutex.Lock()
		stats.SocketsAlive += delta
		statsMutex.Unlock()
	}
}

func (stats *Stats) socketRefs(delta int) {
	if stats != nil {
		statsMutex.Lock()
		stats.SocketRefs += delta
		statsMutex.Unlock()
	}
}
