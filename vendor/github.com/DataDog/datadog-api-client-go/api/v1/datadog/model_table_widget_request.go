/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// TableWidgetRequest TODO.
type TableWidgetRequest struct {
	Aggregator *WidgetAggregator `json:"aggregator,omitempty"`
	// The column name (defaults to the metric name).
	Alias    *string             `json:"alias,omitempty"`
	ApmQuery *LogQueryDefinition `json:"apm_query,omitempty"`
	// TODO.
	ConditionalFormats *[]WidgetConditionalFormat `json:"conditional_formats,omitempty"`
	EventQuery         *EventQueryDefinition      `json:"event_query,omitempty"`
	// For metric queries, the number of lines to show in the table. Only one request should have this property.
	Limit        *int64                  `json:"limit,omitempty"`
	LogQuery     *LogQueryDefinition     `json:"log_query,omitempty"`
	NetworkQuery *LogQueryDefinition     `json:"network_query,omitempty"`
	Order        *WidgetSort             `json:"order,omitempty"`
	ProcessQuery *ProcessQueryDefinition `json:"process_query,omitempty"`
	// TODO.
	Q        *string             `json:"q,omitempty"`
	RumQuery *LogQueryDefinition `json:"rum_query,omitempty"`
}

// NewTableWidgetRequest instantiates a new TableWidgetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTableWidgetRequest() *TableWidgetRequest {
	this := TableWidgetRequest{}
	return &this
}

// NewTableWidgetRequestWithDefaults instantiates a new TableWidgetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTableWidgetRequestWithDefaults() *TableWidgetRequest {
	this := TableWidgetRequest{}
	return &this
}

// GetAggregator returns the Aggregator field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetAggregator() WidgetAggregator {
	if o == nil || o.Aggregator == nil {
		var ret WidgetAggregator
		return ret
	}
	return *o.Aggregator
}

// GetAggregatorOk returns a tuple with the Aggregator field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetAggregatorOk() (*WidgetAggregator, bool) {
	if o == nil || o.Aggregator == nil {
		return nil, false
	}
	return o.Aggregator, true
}

// HasAggregator returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasAggregator() bool {
	if o != nil && o.Aggregator != nil {
		return true
	}

	return false
}

// SetAggregator gets a reference to the given WidgetAggregator and assigns it to the Aggregator field.
func (o *TableWidgetRequest) SetAggregator(v WidgetAggregator) {
	o.Aggregator = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetAlias() string {
	if o == nil || o.Alias == nil {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetAliasOk() (*string, bool) {
	if o == nil || o.Alias == nil {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasAlias() bool {
	if o != nil && o.Alias != nil {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *TableWidgetRequest) SetAlias(v string) {
	o.Alias = &v
}

// GetApmQuery returns the ApmQuery field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetApmQuery() LogQueryDefinition {
	if o == nil || o.ApmQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.ApmQuery
}

// GetApmQueryOk returns a tuple with the ApmQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetApmQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.ApmQuery == nil {
		return nil, false
	}
	return o.ApmQuery, true
}

// HasApmQuery returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasApmQuery() bool {
	if o != nil && o.ApmQuery != nil {
		return true
	}

	return false
}

// SetApmQuery gets a reference to the given LogQueryDefinition and assigns it to the ApmQuery field.
func (o *TableWidgetRequest) SetApmQuery(v LogQueryDefinition) {
	o.ApmQuery = &v
}

// GetConditionalFormats returns the ConditionalFormats field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetConditionalFormats() []WidgetConditionalFormat {
	if o == nil || o.ConditionalFormats == nil {
		var ret []WidgetConditionalFormat
		return ret
	}
	return *o.ConditionalFormats
}

// GetConditionalFormatsOk returns a tuple with the ConditionalFormats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetConditionalFormatsOk() (*[]WidgetConditionalFormat, bool) {
	if o == nil || o.ConditionalFormats == nil {
		return nil, false
	}
	return o.ConditionalFormats, true
}

// HasConditionalFormats returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasConditionalFormats() bool {
	if o != nil && o.ConditionalFormats != nil {
		return true
	}

	return false
}

// SetConditionalFormats gets a reference to the given []WidgetConditionalFormat and assigns it to the ConditionalFormats field.
func (o *TableWidgetRequest) SetConditionalFormats(v []WidgetConditionalFormat) {
	o.ConditionalFormats = &v
}

// GetEventQuery returns the EventQuery field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetEventQuery() EventQueryDefinition {
	if o == nil || o.EventQuery == nil {
		var ret EventQueryDefinition
		return ret
	}
	return *o.EventQuery
}

// GetEventQueryOk returns a tuple with the EventQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetEventQueryOk() (*EventQueryDefinition, bool) {
	if o == nil || o.EventQuery == nil {
		return nil, false
	}
	return o.EventQuery, true
}

// HasEventQuery returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasEventQuery() bool {
	if o != nil && o.EventQuery != nil {
		return true
	}

	return false
}

// SetEventQuery gets a reference to the given EventQueryDefinition and assigns it to the EventQuery field.
func (o *TableWidgetRequest) SetEventQuery(v EventQueryDefinition) {
	o.EventQuery = &v
}

// GetLimit returns the Limit field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetLimit() int64 {
	if o == nil || o.Limit == nil {
		var ret int64
		return ret
	}
	return *o.Limit
}

// GetLimitOk returns a tuple with the Limit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetLimitOk() (*int64, bool) {
	if o == nil || o.Limit == nil {
		return nil, false
	}
	return o.Limit, true
}

// HasLimit returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasLimit() bool {
	if o != nil && o.Limit != nil {
		return true
	}

	return false
}

// SetLimit gets a reference to the given int64 and assigns it to the Limit field.
func (o *TableWidgetRequest) SetLimit(v int64) {
	o.Limit = &v
}

// GetLogQuery returns the LogQuery field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetLogQuery() LogQueryDefinition {
	if o == nil || o.LogQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.LogQuery
}

// GetLogQueryOk returns a tuple with the LogQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetLogQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.LogQuery == nil {
		return nil, false
	}
	return o.LogQuery, true
}

// HasLogQuery returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasLogQuery() bool {
	if o != nil && o.LogQuery != nil {
		return true
	}

	return false
}

// SetLogQuery gets a reference to the given LogQueryDefinition and assigns it to the LogQuery field.
func (o *TableWidgetRequest) SetLogQuery(v LogQueryDefinition) {
	o.LogQuery = &v
}

// GetNetworkQuery returns the NetworkQuery field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetNetworkQuery() LogQueryDefinition {
	if o == nil || o.NetworkQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.NetworkQuery
}

// GetNetworkQueryOk returns a tuple with the NetworkQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetNetworkQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.NetworkQuery == nil {
		return nil, false
	}
	return o.NetworkQuery, true
}

// HasNetworkQuery returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasNetworkQuery() bool {
	if o != nil && o.NetworkQuery != nil {
		return true
	}

	return false
}

// SetNetworkQuery gets a reference to the given LogQueryDefinition and assigns it to the NetworkQuery field.
func (o *TableWidgetRequest) SetNetworkQuery(v LogQueryDefinition) {
	o.NetworkQuery = &v
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetOrder() WidgetSort {
	if o == nil || o.Order == nil {
		var ret WidgetSort
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetOrderOk() (*WidgetSort, bool) {
	if o == nil || o.Order == nil {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasOrder() bool {
	if o != nil && o.Order != nil {
		return true
	}

	return false
}

// SetOrder gets a reference to the given WidgetSort and assigns it to the Order field.
func (o *TableWidgetRequest) SetOrder(v WidgetSort) {
	o.Order = &v
}

// GetProcessQuery returns the ProcessQuery field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetProcessQuery() ProcessQueryDefinition {
	if o == nil || o.ProcessQuery == nil {
		var ret ProcessQueryDefinition
		return ret
	}
	return *o.ProcessQuery
}

// GetProcessQueryOk returns a tuple with the ProcessQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetProcessQueryOk() (*ProcessQueryDefinition, bool) {
	if o == nil || o.ProcessQuery == nil {
		return nil, false
	}
	return o.ProcessQuery, true
}

// HasProcessQuery returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasProcessQuery() bool {
	if o != nil && o.ProcessQuery != nil {
		return true
	}

	return false
}

// SetProcessQuery gets a reference to the given ProcessQueryDefinition and assigns it to the ProcessQuery field.
func (o *TableWidgetRequest) SetProcessQuery(v ProcessQueryDefinition) {
	o.ProcessQuery = &v
}

// GetQ returns the Q field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetQ() string {
	if o == nil || o.Q == nil {
		var ret string
		return ret
	}
	return *o.Q
}

// GetQOk returns a tuple with the Q field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetQOk() (*string, bool) {
	if o == nil || o.Q == nil {
		return nil, false
	}
	return o.Q, true
}

// HasQ returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasQ() bool {
	if o != nil && o.Q != nil {
		return true
	}

	return false
}

// SetQ gets a reference to the given string and assigns it to the Q field.
func (o *TableWidgetRequest) SetQ(v string) {
	o.Q = &v
}

// GetRumQuery returns the RumQuery field value if set, zero value otherwise.
func (o *TableWidgetRequest) GetRumQuery() LogQueryDefinition {
	if o == nil || o.RumQuery == nil {
		var ret LogQueryDefinition
		return ret
	}
	return *o.RumQuery
}

// GetRumQueryOk returns a tuple with the RumQuery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TableWidgetRequest) GetRumQueryOk() (*LogQueryDefinition, bool) {
	if o == nil || o.RumQuery == nil {
		return nil, false
	}
	return o.RumQuery, true
}

// HasRumQuery returns a boolean if a field has been set.
func (o *TableWidgetRequest) HasRumQuery() bool {
	if o != nil && o.RumQuery != nil {
		return true
	}

	return false
}

// SetRumQuery gets a reference to the given LogQueryDefinition and assigns it to the RumQuery field.
func (o *TableWidgetRequest) SetRumQuery(v LogQueryDefinition) {
	o.RumQuery = &v
}

func (o TableWidgetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aggregator != nil {
		toSerialize["aggregator"] = o.Aggregator
	}
	if o.Alias != nil {
		toSerialize["alias"] = o.Alias
	}
	if o.ApmQuery != nil {
		toSerialize["apm_query"] = o.ApmQuery
	}
	if o.ConditionalFormats != nil {
		toSerialize["conditional_formats"] = o.ConditionalFormats
	}
	if o.EventQuery != nil {
		toSerialize["event_query"] = o.EventQuery
	}
	if o.Limit != nil {
		toSerialize["limit"] = o.Limit
	}
	if o.LogQuery != nil {
		toSerialize["log_query"] = o.LogQuery
	}
	if o.NetworkQuery != nil {
		toSerialize["network_query"] = o.NetworkQuery
	}
	if o.Order != nil {
		toSerialize["order"] = o.Order
	}
	if o.ProcessQuery != nil {
		toSerialize["process_query"] = o.ProcessQuery
	}
	if o.Q != nil {
		toSerialize["q"] = o.Q
	}
	if o.RumQuery != nil {
		toSerialize["rum_query"] = o.RumQuery
	}
	return json.Marshal(toSerialize)
}

type NullableTableWidgetRequest struct {
	value *TableWidgetRequest
	isSet bool
}

func (v NullableTableWidgetRequest) Get() *TableWidgetRequest {
	return v.value
}

func (v *NullableTableWidgetRequest) Set(val *TableWidgetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTableWidgetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTableWidgetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTableWidgetRequest(val *TableWidgetRequest) *NullableTableWidgetRequest {
	return &NullableTableWidgetRequest{value: val, isSet: true}
}

func (v NullableTableWidgetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTableWidgetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
