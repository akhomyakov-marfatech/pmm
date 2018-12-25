// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: agent/agent.proto

package agent

import github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/golang/protobuf/ptypes/any"
import _ "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *QANDataRequest) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *QANDataResponse) Validate() error {
	return nil
}
func (this *StateChangedRequest) Validate() error {
	return nil
}
func (this *StateChangedResponse) Validate() error {
	return nil
}
func (this *PingRequest) Validate() error {
	return nil
}
func (this *PingResponse) Validate() error {
	if this.CurrentTime != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.CurrentTime); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("CurrentTime", err)
		}
	}
	return nil
}
func (this *SetStateRequest) Validate() error {
	for _, item := range this.AgentProcesses {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("AgentProcesses", err)
			}
		}
	}
	return nil
}
func (this *SetStateRequest_AgentProcess) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *SetStateResponse) Validate() error {
	return nil
}
func (this *AgentMessage) Validate() error {
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_QanData); ok {
		if oneOfNester.QanData != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.QanData); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("QanData", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_StateChanged); ok {
		if oneOfNester.StateChanged != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StateChanged); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StateChanged", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_Ping); ok {
		if oneOfNester.Ping != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ping); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ping", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*AgentMessage_State); ok {
		if oneOfNester.State != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.State); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("State", err)
			}
		}
	}
	return nil
}
func (this *ServerMessage) Validate() error {
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_QanData); ok {
		if oneOfNester.QanData != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.QanData); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("QanData", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_StateChanged); ok {
		if oneOfNester.StateChanged != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.StateChanged); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("StateChanged", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_Ping); ok {
		if oneOfNester.Ping != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Ping); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Ping", err)
			}
		}
	}
	if oneOfNester, ok := this.GetPayload().(*ServerMessage_State); ok {
		if oneOfNester.State != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.State); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("State", err)
			}
		}
	}
	return nil
}