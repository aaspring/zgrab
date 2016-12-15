/*
 * ZGrab Copyright 2015 Regents of the University of Michigan
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not
 * use this file except in compliance with the License. You may obtain a copy
 * of the License at http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License.
 */

package xssh

// HandshakeLog contains detailed information about each step of the
// SSH handshake, and can be encoded to JSON.
type HandshakeLog struct {
	ServerIDString     string       `json:"server_id_string,omitempty"`
	ClientIDString     string       `json:"client_id_string,omitempty"`
	ServerKex          *kexInitMsg  `json:"server_key_exchange,omitempty"`
	ClientKex          *kexInitMsg  `json:"client_key_exchange,omitempty"`
	AlgorithmSelection *algorithms  `json:"algorithm_selection,omitempty"`
	DHKeyExchange      kexAlgorithm `json:"dh_key_exchange,omitempty"`
	UserAuth           []string     `json:"userauth,omitempty"`
	Crypto             *kexResult   `json:"crypto,omitempty"`
}
