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
	ServerIDString      string              `json:"server_id_string,omitempty"`
	ClientIDString      string              `json:"client_id_string,omitempty"`
	ServerKex           *KeyExchangeMsg     `json:"server_key_exchange,omitempty"`
	ClientKex           *KeyExchangeMsg     `json:"client_key_exchange,omitempty"`
	AlgorithmSelection  *AlgorithmSelection `json:"algorithm_selection,omitempty"`
	ServerDHKeyExchange kexAlgorithm        `json:"dh_key_exchange"`
	UserAuth            *UserAuthentication `json:"userauth,omitempty"`
	Crypto              *Crypto             `json:"crypto,omitempty"`

	/*
		ClientKexExchangeInit *KeyExchangeInit              `json:"client_key_exchange_init,omitempty"`
		Algorithms            *AlgorithmSelection           `json:"algorithms,omitempty"`
		KexDHGroupRequest     *KeyExchangeDHGroupRequest    `json:"key_exchange_dh_group_request,omitempty"`
		KexDHGroupParams      *KeyExchangeDHGroupParameters `json:"key_exchange_dh_group_params,omitempty"`
		KexDHGroupInit        *KeyExchangeDHGroupInit       `json:"key_exchange_dh_group_init,omitempty"`
		KexDHGroupReply       *KeyExchangeDHGroupReply      `json:"key_exchange_dh_group_reply,omitempty"`
		DHInit                *KeyExchangeDHInit            `json:"key_exchange_dh_init,omitempty"`
		DHReply               *KeyExchangeDHInitReply       `json:"key_exchange_dh_reply,omitempty"`
	*/
}

type UserAuthentication struct {
	MethodNames []string `json:"method_names,omitempty"`
}

type AlgorithmSelection struct {
	KexAlgorithm      string         `json:"key_exchange_algorithm,omitempty"`
	HostKeyAlgorithm  string         `json:"host_key_algorithm,omitempty"`
	ClientToServerAlg AlgorithmGroup `json:"client_to_server_algorithm_group"`
	ServerToClientAlg AlgorithmGroup `json:"server_to_client_algorithm_group"`
}

type Crypto struct {
	SessionID []byte `json:"session_id,omitempty"`
	H         []byte `json:"H"`
	K         []byte `json:"K"`
}

type AlgorithmGroup struct {
	Cipher      string `json:"cipher"`
	MAC         string `json:"mac"`
	Compression string `json:"compression"`
}

type KeyExchangeMsg struct {
	HostKeyAlgorithms       []string `json:"host_key_algorithms"`
	Cookie                  []byte   `json:"cookie"`
	KexAlgorithms           []string `json:"kex_algorithms"`
	CiphersClientServer     []string `json:"client_to_server_ciphers"`
	CiphersServerClient     []string `json:"server_to_client_ciphers"`
	MACsClientServer        []string `json:"client_to_server_macs"`
	MACsServerClient        []string `json:"server_to_client_macs"`
	CompressionClientServer []string `json:"client_to_server_compression"`
	CompressionServerClient []string `json:"server_to_client_compression"`
	LanguagesClientServer   []string `json:"client_to_server_languages"`
	LanguagesServerClient   []string `json:"server_to_client_languages"`
}
