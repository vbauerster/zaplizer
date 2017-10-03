// Code generated by "stringer -type=LogLevel,LogMode"; DO NOT EDIT.

package zaplizer

import "fmt"

const _LogLevel_name = "LogLevelUnknownDEBUGINFOWARNERRORDPANICPANICFATAL"

var _LogLevel_index = [...]uint8{0, 15, 20, 24, 28, 33, 39, 44, 49}

func (i LogLevel) String() string {
	if i < 0 || i >= LogLevel(len(_LogLevel_index)-1) {
		return fmt.Sprintf("LogLevel(%d)", i)
	}
	return _LogLevel_name[_LogLevel_index[i]:_LogLevel_index[i+1]]
}

const _LogMode_name = "LogModeUnknownDEVELOPMENTPRODUCTION"

var _LogMode_index = [...]uint8{0, 14, 25, 35}

func (i LogMode) String() string {
	if i < 0 || i >= LogMode(len(_LogMode_index)-1) {
		return fmt.Sprintf("LogMode(%d)", i)
	}
	return _LogMode_name[_LogMode_index[i]:_LogMode_index[i+1]]
}
