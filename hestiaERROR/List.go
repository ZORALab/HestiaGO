// Copyright 2022 "Holloway" Chew, Kean Ho <kean.ho.chew@zoralab.com>
// Copyright 2022 ZORALab Enterprise <tech@zoralab.com>
//
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not
// use this file except in compliance with the License. You may obtain a copy of
// the License at
//
//                  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations under
// the License.

package hestiaERROR

const (
	// Common errors and operations
	OK                     Error = 0
	BAD_EXEC               Error = 1
	BAD_DESCRIPTOR         Error = 2
	BAD_EXCHANGE           Error = 3
	BAD_MOUNT              Error = 4
	BAD_PIPE               Error = 5
	BAD_REQUEST            Error = 6
	BAD_STREAM_PIPE        Error = 7
	CANCELED               Error = 8
	CLEANING_REQUIRED      Error = 9
	DEADLOCK               Error = 10
	EXPIRED                Error = 11
	ILLEGAL_BYTE_SEQUENCE  Error = 12
	ILLEGAL_SEEK           Error = 13
	INVALID_ARGUMENT       Error = 14
	IS_EMPTY               Error = 15
	MAXED_EXCHANGE         Error = 16
	MAXED_QUOTA            Error = 17
	MISSING_LOCK           Error = 18
	NOT_EMPTY              Error = 19
	NOT_PERMITTED          Error = 20
	NOT_POSSIBLE           Error = 21
	NOT_POSSIBLE_BY_RFKILL Error = 22
	NOT_RECOVERABLE        Error = 23
	OUT_OF_RANGE           Error = 24
	PERMISSION_DENIED      Error = 25
	TIMEOUT                Error = 26
	TOO_MANY_READ          Error = 27
	TOO_MANY_LOOP          Error = 28
	TOO_MANY_REFERENCES    Error = 29
	TOO_MANY_LINK          Error = 30
	TRY_AGAIN              Error = 31
	UNSUPPORTED            Error = 32
	WOULD_BLOCK            Error = 33

	// lifecycle states
	RESTART  Error = 34
	RESUME   Error = 35
	SHUTDOWN Error = 36
	SLEEP    Error = 37
	STALLED  Error = 38
	STANDBY  Error = 39

	// progress
	PROGRESS_SCHEDULED         Error = 40
	PROGRESS_ALREADY_EXECUTING Error = 41
	PROGRESS_EXECUTING         Error = 42
	PROGRESS_COMPLETED         Error = 43

	// tri-tier inter-package communications
	LV1_NOT_SYNC Error = 44
	LV1_PAUSED   Error = 45
	LV1_RESET    Error = 46

	LV2_NOT_SYNC Error = 47
	LV2_PAUSED   Error = 48
	LV2_RESET    Error = 49

	LV3_NOT_SYNC Error = 50
	LV3_PAUSED   Error = 51
	LV3_RESET    Error = 52

	// data (input/output parameters, type, etc)
	DATA_BAD        Error = 53
	DATA_EMPTY      Error = 54
	DATA_INVALID    Error = 55
	DATA_IS_UNIQUE  Error = 56
	DATA_MISSING    Error = 57
	DATA_NOT_UNIQUE Error = 58
	DATA_OVERFLOW   Error = 59
	DATA_REMOVED    Error = 60
	DATA_TOO_LONG   Error = 61
	DATA_MISMATCHED Error = 62

	// entity (device, file, directory, etc)
	ENTITY_BAD             Error = 63
	ENTITY_BUSY            Error = 64
	ENTITY_DEAD            Error = 65
	ENTITY_EXISTS          Error = 66
	ENTITY_FAULTY          Error = 67
	ENTITY_MISSING         Error = 68
	ENTITY_MISSING_CHILD   Error = 69
	ENTITY_OUT_OF_BUFFER   Error = 70
	ENTITY_POISONED        Error = 71
	ENTITY_READ_ONLY       Error = 72
	ENTITY_TOO_BIG         Error = 73
	ENTITY_TOO_MANY_OPENED Error = 74
	ENTITY_UNATTACHED      Error = 75

	ENTITY_IS_DIRECTORY Error = 76
	ENTITY_IS_FILE      Error = 77
	ENTITY_IS_LINK      Error = 78
	ENTITY_IS_SOCKET    Error = 79

	ENTITY_IS_NOT_DIRECTORY Error = 80
	ENTITY_IS_NOT_FILE      Error = 81
	ENTITY_IS_NOT_LINK      Error = 82
	ENTITY_IS_NOT_SOCKET    Error = 83

	ENTITY_REMOTE_CHANGED Error = 84
	ENTITY_REMOTE_ERROR   Error = 85
	ENTITY_REMOTE_IO      Error = 86

	ENTITY_MISSING_STREAMABLE_RESOURCES Error = 87
	ENTITY_NOT_STREAMABLE               Error = 88
	ENTITY_STREAMABLE                   Error = 89

	ENTITY_A_TYPEWRITER     Error = 90
	ENTITY_NOT_A_TYPEWRITER Error = 91

	ENTITY_BAD_DESCRIPTOR     Error = 92
	ENTITY_FILETABLE_OVERFLOW Error = 93

	// key (cryptography)
	KEY_BAD       Error = 94
	KEY_DESTROYED Error = 95
	KEY_EXPIRED   Error = 96
	KEY_MISSING   Error = 87
	KEY_REJECTED  Error = 98
	KEY_REVOKED   Error = 99

	// library
	LIBRARY_BAD         Error = 100
	LIBRARY_CORRUPTED   Error = 101
	LIBRARY_EXEC_FAILED Error = 102
	LIBRARY_MAXED       Error = 103
	LIBRARY_MISSING     Error = 104

	// network
	NETWORK_BAD           Error = 105
	NETWORK_BAD_AD        Error = 106
	NETWORK_DOWN          Error = 107
	NETWORK_NOT_CONNECTED Error = 108
	NETWORK_RESET         Error = 109
	NETWORK_RFS           Error = 110
	NETWORK_UNREACHABLE   Error = 111

	NETWORK_HOST_DOWN          Error = 112
	NETWORK_HOST_UNREACHABLE   Error = 113
	NETWORK_SOCKET_UNSUPPORTED Error = 114

	NETWORK_ADDRESS_IN_USE      Error = 115
	NETWORK_ADDRESS_UNAVAILABLE Error = 116

	NETWORK_CONN_ABORTED              Error = 117
	NETWORK_CONN_IS_CONNECTED         Error = 118
	NETWORK_CONN_MISSING_DEST_ADDRESS Error = 119
	NETWORK_CONN_MULTIHOP             Error = 120
	NETWORK_CONN_NOT_CONNECTED        Error = 121
	NETWORK_CONN_REFUSED              Error = 122
	NETWORK_CONN_RESET                Error = 123

	NETWORK_PAYLOAD_BAD      Error = 124
	NETWORK_PAYLOAD_EMPTY    Error = 125
	NETWORK_PAYLOAD_MISSING  Error = 126
	NETWORK_PAYLOAD_TOO_LONG Error = 127

	// protocol
	PROTOCOL_ADDRESS_UNSUPPORTED Error = 128
	PROTOCOL_BAD                 Error = 129
	PROTOCOL_FAMILY_UNSUPPORTED  Error = 130
	PROTOCOL_MISSING             Error = 131
	PROTOCOL_UNSUPPORTED         Error = 132

	// system (e.g. os, interactable system)
	SYSTEM_BAD_IO               Error = 133
	SYSTEM_DEVICE_CROSS_LINK    Error = 134
	SYSTEM_INTERRUPT_CALL       Error = 135
	SYSTEM_INVALID              Error = 136
	SYSTEM_MISSING_BLOCK_DEVICE Error = 137
	SYSTEM_MISSING_DEVICE       Error = 138
	SYSTEM_MISSING_IO           Error = 139
	SYSTEM_MISSING_PROCESS      Error = 140
	SYSTEM_OUT_OF_DOMAIN        Error = 141
	SYSTEM_OUT_OF_MEMORY        Error = 142
	SYSTEM_OUT_OF_SPACE         Error = 143
	SYSTEM_READ_ONLY_FILESYSTEM Error = 144

	// user
	USER_ACCESS_BANNED       Error = 145
	USER_ACCESS_LOCKED       Error = 146
	USER_ACCESS_NOT_VERIFIED Error = 147
	USER_ACCESS_REJECTED     Error = 148
	USER_ACCESS_REVOKED      Error = 149

	USER_ID_BAD           Error = 150
	USER_ID_EXISTS        Error = 151
	USER_ID_MISSING       Error = 152
	USER_MFA_BAD          Error = 153
	USER_MFA_EXPIRED      Error = 154
	USER_MFA_MISSING      Error = 155
	USER_PASSWORD_BAD     Error = 156
	USER_PASSWORD_EXPIRED Error = 157
	USER_PASSWORD_MISSING Error = 158
	USER_KEYFILE_BAD      Error = 159
	USER_KEYFILE_EXPIRED  Error = 160
	USER_KEYFILE_MISSING  Error = 161

	USER_ACCESS_TOKEN_BAD     Error = 162
	USER_ACCESS_TOKEN_EXPIRED Error = 163
	USER_ACCESS_TOKEN_MISSING Error = 164
	USER_ACCESS_TOKEN_REVOKED Error = 165
)
