// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * API version: 2023.3
 */

package server

// AccountServiceV1150LdapSearchSettings - The settings to search a generic LDAP service.
type AccountServiceV1150LdapSearchSettings struct {

	// The base distinguished names to use to search an external LDAP service.
	BaseDistinguishedNames []*string `json:"BaseDistinguishedNames,omitempty"`

	// The attribute name that contains the LDAP user's email address.
	EmailAttribute *string `json:"EmailAttribute,omitempty"`

	// The attribute name that contains the LDAP group name entry.
	GroupNameAttribute *string `json:"GroupNameAttribute,omitempty"`

	// The attribute name that contains the groups for a user on the LDAP user entry.
	GroupsAttribute *string `json:"GroupsAttribute,omitempty"`

	// The attribute name that contains the LDAP user's SSH public key entry.
	SSHKeyAttribute *string `json:"SSHKeyAttribute,omitempty"`

	// The attribute name that contains the LDAP user name entry.
	UsernameAttribute *string `json:"UsernameAttribute,omitempty"`
}

// AssertAccountServiceV1150LdapSearchSettingsRequired checks if the required fields are not zero-ed
func AssertAccountServiceV1150LdapSearchSettingsRequired(obj AccountServiceV1150LdapSearchSettings) error {
	return nil
}

// AssertAccountServiceV1150LdapSearchSettingsConstraints checks if the values respects the defined constraints
func AssertAccountServiceV1150LdapSearchSettingsConstraints(obj AccountServiceV1150LdapSearchSettings) error {
	return nil
}