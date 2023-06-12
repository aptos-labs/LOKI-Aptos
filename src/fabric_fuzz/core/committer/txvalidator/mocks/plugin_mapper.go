// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import (
	txvalidator "github.com/hyperledger/fabric/core/committer/txvalidator"
	validation "github.com/hyperledger/fabric/core/handlers/validation/api"
	mock "github.com/stretchr/testify/mock"
)

// PluginMapper is an autogenerated mock type for the PluginMapper type
type PluginMapper struct {
	mock.Mock
}

// PluginFactoryByName provides a mock function with given fields: name
func (_m *PluginMapper) PluginFactoryByName(name txvalidator.PluginName) validation.PluginFactory {
	ret := _m.Called(name)

	var r0 validation.PluginFactory
	if rf, ok := ret.Get(0).(func(txvalidator.PluginName) validation.PluginFactory); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(validation.PluginFactory)
		}
	}

	return r0
}