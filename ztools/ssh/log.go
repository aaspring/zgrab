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

package ssh

// HandshakeLog contains detailed information about each step of the
// SSH handshake, and can be encoded to JSON.
type HandshakeLog struct {
	ClientProtocol *ProtocolAgreement `json:"client_protocol,omitempty"`
	ServerProtocol *ProtocolAgreement `json:"server_protocol,omitempty"`

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
	ServerKex KeyExchange        `json:"server_key_exchange_init,omitempty"`
	UserAuth  UserAuthentication `json:"userauth,omitempty"`
	Crypto    Crypto             `json:"crypto,omitempty"`
}

type UserAuthentication struct {
	MethodNames []string `json:"method_names,omitempty"`
}

type Crypto struct {
	SessionID []byte `json:"session_id,omitempty"`
}

// ProtocolAgreement represents the client and server protocol banners
//
// RFC specifies the format for the banner as specified in RFC 4523 Section 4.2
//       SSH-protoversion-softwareversion SP comments CR LF
//
// The server MAY send other lines of data before sending the version.
// The lines are terminated by CR LF, and SHOULD be encoded as UTF-8
//
// See http://tools.ietf.org/html/rfc4253 for details
type ProtocolAgreement struct {
	RawBanner string `json:"raw_banner,omitempty"`
	//ProtocolVersion string `json:"protocol_version,omitempty"`
	//SoftwareVersion string `json:"software_version,omitempty"`
	//Comments        string `json:"comments,omitempty"`
}

type KeyExchange struct {
	HostKeyAlgorithms []string `json:"host_key_algorithms"`
	/*
		Cookie                    Cookie   `json:"cookie"`
		KexAlgorithms             NameList `json:"key_exchange_algorithms"`
		EncryptionClientToServer  NameList `json:"encryption_client_to_server"`
		EncryptionServerToClient  NameList `json:"encryption_server_to_client"`
		MACClientToServer         NameList `json:"mac_client_to_server"`
		MACServerToClient         NameList `json:"mac_server_to_client"`
		CompressionClientToServer NameList `json:"compression_client_to_server"`
		CompressionServerToClient NameList `json:"compression_server_to_client"`
		LanguageClientToServer    NameList `json:"language_client_to_server"`
		LanguageServerToClient    NameList `json:"language_server_to_client"`
		FirstKexPacketFollows     bool     `json:"first_kex_packet_follows"`
		Zero                      uint32   `json:"zero"`
	*/
}

/*
type AlgorithmSelection struct {
	KexAlgorithm     string `json:"key_exchange_algorithm,omitempty"`
	HostKeyAlgorithm string `json:"host_key_algorithm,omitempty"`
}
*/
