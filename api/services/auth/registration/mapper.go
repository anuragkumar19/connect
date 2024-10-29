package registration

import (
	authv1 "github.com/anuragkumar19/connect/api/gen/auth/v1"
	"github.com/anuragkumar19/connect/database"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func mapRegistrationFlow(flowDB *database.RegistrationFlow) *authv1.RegistrationFlow {
	fields := &authv1.RegistrationFlowFields{}

	if flowDB.EmailValue.Valid {
		v := &authv1.ValueField{
			Value:   flowDB.EmailValue.String,
			IsValid: flowDB.EmailIsValid,
		}
		if flowDB.EmailErrorMessage.Valid {
			v.ErrorMessage = &flowDB.EmailErrorMessage.String
		}
		fields.Email = v
	}

	if flowDB.NameValue.Valid {
		v := &authv1.ValueField{
			Value:   flowDB.NameValue.String,
			IsValid: flowDB.NameIsValid,
		}
		if flowDB.NameErrorMessage.Valid {
			v.ErrorMessage = &flowDB.NameErrorMessage.String
		}
		fields.Name = v
	}
	if flowDB.UsernameValue.Valid {
		v := &authv1.ValueField{
			Value:   flowDB.UsernameValue.String,
			IsValid: flowDB.UsernameIsValid,
		}
		if flowDB.UsernameErrorMessage.Valid {
			v.ErrorMessage = &flowDB.UsernameErrorMessage.String
		}
		fields.Username = v
	}
	if flowDB.PhoneNumberValue.Valid {
		v := &authv1.ValueField{
			Value:   flowDB.PhoneNumberValue.String,
			IsValid: flowDB.PhoneNumberIsValid,
		}
		if flowDB.PhoneNumberErrorMessage.Valid {
			v.ErrorMessage = &flowDB.PhoneNumberErrorMessage.String
		}
		fields.PhoneNumber = v
	}
	if flowDB.PasswordValue.Valid {
		v := &authv1.PasswordField{
			IsValid: flowDB.PasswordIsValid,
		}
		if flowDB.PasswordErrorMessage.Valid {
			v.ErrorMessage = &flowDB.PasswordErrorMessage.String
		}
		fields.Password = v
	}
	if flowDB.EmailCode.Valid {
		v := &authv1.VerificationCodeField{
			IncorrectVerificationAttempts: int64(flowDB.EmailCodeIncorrectVerificationAttempts),
			RemainingVerificationAttempts: int64(flowDB.EmailCodeRemainingVerificationAttempts),
			//TODO: rate limit fields
			IsVerified: flowDB.EmailCodeIsVerified,
		}
		if flowDB.EmailCodeErrorMessage.Valid {
			v.ErrorMessage = &flowDB.EmailCodeErrorMessage.String
		}
		fields.EmailCode = v
	}
	if flowDB.PhoneNumberCode.Valid {
		v := &authv1.VerificationCodeField{
			IncorrectVerificationAttempts: int64(flowDB.PhoneNumberCodeIncorrectVerificationAttempts),
			RemainingVerificationAttempts: int64(flowDB.PhoneNumberCodeRemainingVerificationAttempts),
			//TODO: rate limit fields
			IsVerified: flowDB.PhoneNumberCodeIsVerified,
		}
		if flowDB.PhoneNumberCodeErrorMessage.Valid {
			v.ErrorMessage = &flowDB.PhoneNumberCodeErrorMessage.String
		}
		fields.PhoneNumberCode = v
	}
	return &authv1.RegistrationFlow{
		Id:        flowDB.ID,
		CreatedAt: timestamppb.New(flowDB.CreatedAt),
		UpdatedAt: timestamppb.New(flowDB.UpdatedAt),
		ExpireAt:  timestamppb.New(flowDB.ExpireAt),
		CsrfToken: flowDB.CsrfToken,
		Fields:    fields,
		State:     mapRegistrationFlowState(flowDB.State),
	}
}

func mapRegistrationFlowState(stateDB database.RegistrationFlowsState) authv1.RegistrationFlowState {
	switch stateDB {
	case database.RegistrationFlowsStateInvalidFields:
		return authv1.RegistrationFlowState_REGISTRATION_FLOW_STATE_INVALID_FIELDS
	case database.RegistrationFlowsStateVerificationCodeSent:
		return authv1.RegistrationFlowState_REGISTRATION_FLOW_STATE_VERIFICATION_CODE_SENT
	case database.RegistrationFlowsStateRegistrationCompleted:
		return authv1.RegistrationFlowState_REGISTRATION_FLOW_STATE_COMPLETED
	default:
		return authv1.RegistrationFlowState_REGISTRATION_FLOW_STATE_UNSPECIFIED
	}
}
