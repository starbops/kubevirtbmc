// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

import (
	"errors"
)

// AccountServiceV1150AccountService - The AccountService schema defines an account service.  The properties are common to, and enable management of, all user accounts.  The properties include the password requirements and control features, such as account lockout.  Properties and actions in this service specify general behavior that should be followed for typical accounts, however implementations might override these behaviors for special accounts or situations to avoid denial of service or other deadlock situations.
type AccountServiceV1150AccountService struct {

	// The OData description of a payload.
	OdataContext string `json:"@odata.context,omitempty"`

	// The current ETag of the resource.
	OdataEtag string `json:"@odata.etag,omitempty"`

	// The unique identifier for a resource.
	OdataId string `json:"@odata.id"`

	// The type of a resource.
	OdataType string `json:"@odata.type"`

	// The period of time, in seconds, between the last failed login attempt and the reset of the lockout threshold counter.  This value must be less than or equal to the AccountLockoutDuration value.  A reset sets the counter to `0`.
	AccountLockoutCounterResetAfter int64 `json:"AccountLockoutCounterResetAfter,omitempty"`

	// An indication of whether the threshold counter is reset after AccountLockoutCounterResetAfter expires.  If `true`, it is reset.  If `false`, only a successful login resets the threshold counter and if the user reaches the AccountLockoutThreshold limit, the account will be locked out indefinitely and only an administrator-issued reset clears the threshold counter.  If this property is absent, the default is `true`.
	AccountLockoutCounterResetEnabled bool `json:"AccountLockoutCounterResetEnabled,omitempty"`

	// The period of time, in seconds, that an account is locked after the number of failed login attempts reaches the account lockout threshold, within the period between the last failed login attempt and the reset of the lockout threshold counter.  If this value is `0`, no lockout will occur.  If the AccountLockoutCounterResetEnabled value is `false`, this property is ignored.
	AccountLockoutDuration *int64 `json:"AccountLockoutDuration,omitempty"`

	// The number of allowed failed login attempts before a user account is locked for a specified duration.  If `0`, the account is never locked.
	AccountLockoutThreshold *int64 `json:"AccountLockoutThreshold,omitempty"`

	Accounts OdataV4IdRef `json:"Accounts,omitempty"`

	Actions AccountServiceV1150Actions `json:"Actions,omitempty"`

	ActiveDirectory AccountServiceV1150ExternalAccountProvider `json:"ActiveDirectory,omitempty"`

	AdditionalExternalAccountProviders OdataV4IdRef `json:"AdditionalExternalAccountProviders,omitempty"`

	// The number of authorization failures per account that are allowed before the failed attempt is logged to the manager log.
	AuthFailureLoggingThreshold int64 `json:"AuthFailureLoggingThreshold,omitempty"`

	// The description of this resource.  Used for commonality in the schema definitions.
	Description string `json:"Description,omitempty"`

	HTTPBasicAuth AccountServiceV1150BasicAuthState `json:"HTTPBasicAuth,omitempty"`

	// The unique identifier for this resource within the collection of similar resources.
	Id string `json:"Id"`

	LDAP AccountServiceV1150ExternalAccountProvider `json:"LDAP,omitempty"`

	LocalAccountAuth AccountServiceV1150LocalAccountAuth `json:"LocalAccountAuth,omitempty"`

	// The maximum password length for this account service.
	MaxPasswordLength int64 `json:"MaxPasswordLength,omitempty"`

	// The minimum password length for this account service.
	MinPasswordLength int64 `json:"MinPasswordLength,omitempty"`

	MultiFactorAuth AccountServiceV1150MultiFactorAuth `json:"MultiFactorAuth,omitempty"`

	// The name of the resource or array member.
	Name string `json:"Name"`

	OAuth2 AccountServiceV1150ExternalAccountProvider `json:"OAuth2,omitempty"`

	// The OEM extension.
	Oem map[string]interface{} `json:"Oem,omitempty"`

	OutboundConnections OdataV4IdRef `json:"OutboundConnections,omitempty"`

	// The number of days before account passwords in this account service will expire.
	PasswordExpirationDays *int64 `json:"PasswordExpirationDays,omitempty"`

	PrivilegeMap OdataV4IdRef `json:"PrivilegeMap,omitempty"`

	// An indication of whether clients are required to invoke the ChangePassword action to modify account passwords.
	RequireChangePasswordAction *bool `json:"RequireChangePasswordAction,omitempty"`

	// The set of restricted OEM privileges.
	RestrictedOemPrivileges []string `json:"RestrictedOemPrivileges,omitempty"`

	// The set of restricted Redfish privileges.
	RestrictedPrivileges []PrivilegesPrivilegeType `json:"RestrictedPrivileges,omitempty"`

	Roles OdataV4IdRef `json:"Roles,omitempty"`

	// An indication of whether the account service is enabled.  If `true`, it is enabled.  If `false`, it is disabled and users cannot be created, deleted, or modified, and new sessions cannot be started.  However, established sessions might still continue to run.  Any service, such as the session service, that attempts to access the disabled account service fails.  However, this does not affect HTTP Basic Authentication connections.
	ServiceEnabled *bool `json:"ServiceEnabled,omitempty"`

	Status ResourceStatus `json:"Status,omitempty"`

	// The account types supported by the service.
	SupportedAccountTypes []ManagerAccountAccountTypes `json:"SupportedAccountTypes,omitempty"`

	// The OEM account types supported by the service.
	SupportedOEMAccountTypes []string `json:"SupportedOEMAccountTypes,omitempty"`

	TACACSplus AccountServiceV1150ExternalAccountProvider `json:"TACACSplus,omitempty"`
}

// AssertAccountServiceV1150AccountServiceRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1150AccountServiceRequired(obj AccountServiceV1150AccountService) error {
	elements := map[string]interface{}{
		"@odata.id":   obj.OdataId,
		"@odata.type": obj.OdataType,
		"Id":          obj.Id,
		"Name":        obj.Name,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertOdataV4IdRefRequired(obj.Accounts); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ActionsRequired(obj.Actions); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ExternalAccountProviderRequired(obj.ActiveDirectory); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.AdditionalExternalAccountProviders); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ExternalAccountProviderRequired(obj.LDAP); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150MultiFactorAuthRequired(obj.MultiFactorAuth); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ExternalAccountProviderRequired(obj.OAuth2); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.OutboundConnections); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.PrivilegeMap); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefRequired(obj.Roles); err != nil {
		return err
	}
	if err := AssertResourceStatusRequired(obj.Status); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ExternalAccountProviderRequired(obj.TACACSplus); err != nil {
		return err
	}
	return nil
}

// AssertAccountServiceV1150AccountServiceConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1150AccountServiceConstraints(obj AccountServiceV1150AccountService) error {
	if obj.AccountLockoutCounterResetAfter < 0 {
		return &ParsingError{Param: "AccountLockoutCounterResetAfter", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AccountLockoutDuration != nil && *obj.AccountLockoutDuration < 0 {
		return &ParsingError{Param: "AccountLockoutDuration", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.AccountLockoutThreshold != nil && *obj.AccountLockoutThreshold < 0 {
		return &ParsingError{Param: "AccountLockoutThreshold", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertOdataV4IdRefConstraints(obj.Accounts); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ActionsConstraints(obj.Actions); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ExternalAccountProviderConstraints(obj.ActiveDirectory); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.AdditionalExternalAccountProviders); err != nil {
		return err
	}
	if obj.AuthFailureLoggingThreshold < 0 {
		return &ParsingError{Param: "AuthFailureLoggingThreshold", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertAccountServiceV1150ExternalAccountProviderConstraints(obj.LDAP); err != nil {
		return err
	}
	if obj.MaxPasswordLength < 0 {
		return &ParsingError{Param: "MaxPasswordLength", Err: errors.New(errMsgMinValueConstraint)}
	}
	if obj.MinPasswordLength < 0 {
		return &ParsingError{Param: "MinPasswordLength", Err: errors.New(errMsgMinValueConstraint)}
	}
	if err := AssertAccountServiceV1150MultiFactorAuthConstraints(obj.MultiFactorAuth); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ExternalAccountProviderConstraints(obj.OAuth2); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.OutboundConnections); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.PrivilegeMap); err != nil {
		return err
	}
	if err := AssertOdataV4IdRefConstraints(obj.Roles); err != nil {
		return err
	}
	if err := AssertResourceStatusConstraints(obj.Status); err != nil {
		return err
	}
	if err := AssertAccountServiceV1150ExternalAccountProviderConstraints(obj.TACACSplus); err != nil {
		return err
	}
	return nil
}
