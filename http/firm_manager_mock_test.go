package http

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i ugc_test_task/http.FirmManager -o ./firm_manager_mock_test.go -n FirmManagerMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"
	"ugc_test_task/firmmng"
	"ugc_test_task/models"

	"github.com/gojuno/minimock/v3"
)

// FirmManagerMock implements FirmManager
type FirmManagerMock struct {
	t minimock.Tester

	funcGetFirms          func(query firmmng.GetQuery, clb func(firm models.Company) error)
	inspectFuncGetFirms   func(query firmmng.GetQuery, clb func(firm models.Company) error)
	afterGetFirmsCounter  uint64
	beforeGetFirmsCounter uint64
	GetFirmsMock          mFirmManagerMockGetFirms
}

// NewFirmManagerMock returns a mock for FirmManager
func NewFirmManagerMock(t minimock.Tester) *FirmManagerMock {
	m := &FirmManagerMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetFirmsMock = mFirmManagerMockGetFirms{mock: m}
	m.GetFirmsMock.callArgs = []*FirmManagerMockGetFirmsParams{}

	return m
}

type mFirmManagerMockGetFirms struct {
	mock               *FirmManagerMock
	defaultExpectation *FirmManagerMockGetFirmsExpectation
	expectations       []*FirmManagerMockGetFirmsExpectation

	callArgs []*FirmManagerMockGetFirmsParams
	mutex    sync.RWMutex
}

// FirmManagerMockGetFirmsExpectation specifies expectation struct of the FirmManager.GetFirms
type FirmManagerMockGetFirmsExpectation struct {
	mock   *FirmManagerMock
	params *FirmManagerMockGetFirmsParams

	Counter uint64
}

// FirmManagerMockGetFirmsParams contains parameters of the FirmManager.GetFirms
type FirmManagerMockGetFirmsParams struct {
	query firmmng.GetQuery
	clb   func(firm models.Company) error
}

// Expect sets up expected params for FirmManager.GetFirms
func (mmGetFirms *mFirmManagerMockGetFirms) Expect(query firmmng.GetQuery, clb func(firm models.Company) error) *mFirmManagerMockGetFirms {
	if mmGetFirms.mock.funcGetFirms != nil {
		mmGetFirms.mock.t.Fatalf("FirmManagerMock.GetFirms mock is already set by Set")
	}

	if mmGetFirms.defaultExpectation == nil {
		mmGetFirms.defaultExpectation = &FirmManagerMockGetFirmsExpectation{}
	}

	mmGetFirms.defaultExpectation.params = &FirmManagerMockGetFirmsParams{query, clb}
	for _, e := range mmGetFirms.expectations {
		if minimock.Equal(e.params, mmGetFirms.defaultExpectation.params) {
			mmGetFirms.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetFirms.defaultExpectation.params)
		}
	}

	return mmGetFirms
}

// Inspect accepts an inspector function that has same arguments as the FirmManager.GetFirms
func (mmGetFirms *mFirmManagerMockGetFirms) Inspect(f func(query firmmng.GetQuery, clb func(firm models.Company) error)) *mFirmManagerMockGetFirms {
	if mmGetFirms.mock.inspectFuncGetFirms != nil {
		mmGetFirms.mock.t.Fatalf("Inspect function is already set for FirmManagerMock.GetFirms")
	}

	mmGetFirms.mock.inspectFuncGetFirms = f

	return mmGetFirms
}

// Return sets up results that will be returned by FirmManager.GetFirms
func (mmGetFirms *mFirmManagerMockGetFirms) Return() *FirmManagerMock {
	if mmGetFirms.mock.funcGetFirms != nil {
		mmGetFirms.mock.t.Fatalf("FirmManagerMock.GetFirms mock is already set by Set")
	}

	if mmGetFirms.defaultExpectation == nil {
		mmGetFirms.defaultExpectation = &FirmManagerMockGetFirmsExpectation{mock: mmGetFirms.mock}
	}

	return mmGetFirms.mock
}

//Set uses given function f to mock the FirmManager.GetFirms method
func (mmGetFirms *mFirmManagerMockGetFirms) Set(f func(query firmmng.GetQuery, clb func(firm models.Company) error)) *FirmManagerMock {
	if mmGetFirms.defaultExpectation != nil {
		mmGetFirms.mock.t.Fatalf("Default expectation is already set for the FirmManager.GetFirms method")
	}

	if len(mmGetFirms.expectations) > 0 {
		mmGetFirms.mock.t.Fatalf("Some expectations are already set for the FirmManager.GetFirms method")
	}

	mmGetFirms.mock.funcGetFirms = f
	return mmGetFirms.mock
}

// GetFirms implements FirmManager
func (mmGetFirms *FirmManagerMock) GetFirms(query firmmng.GetQuery, clb func(firm models.Company) error) {
	mm_atomic.AddUint64(&mmGetFirms.beforeGetFirmsCounter, 1)
	defer mm_atomic.AddUint64(&mmGetFirms.afterGetFirmsCounter, 1)

	if mmGetFirms.inspectFuncGetFirms != nil {
		mmGetFirms.inspectFuncGetFirms(query, clb)
	}

	mm_params := &FirmManagerMockGetFirmsParams{query, clb}

	// Record call args
	mmGetFirms.GetFirmsMock.mutex.Lock()
	mmGetFirms.GetFirmsMock.callArgs = append(mmGetFirms.GetFirmsMock.callArgs, mm_params)
	mmGetFirms.GetFirmsMock.mutex.Unlock()

	for _, e := range mmGetFirms.GetFirmsMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmGetFirms.GetFirmsMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetFirms.GetFirmsMock.defaultExpectation.Counter, 1)
		mm_want := mmGetFirms.GetFirmsMock.defaultExpectation.params
		mm_got := FirmManagerMockGetFirmsParams{query, clb}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetFirms.t.Errorf("FirmManagerMock.GetFirms got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmGetFirms.funcGetFirms != nil {
		mmGetFirms.funcGetFirms(query, clb)
		return
	}
	mmGetFirms.t.Fatalf("Unexpected call to FirmManagerMock.GetFirms. %v %v", query, clb)

}

// GetFirmsAfterCounter returns a count of finished FirmManagerMock.GetFirms invocations
func (mmGetFirms *FirmManagerMock) GetFirmsAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFirms.afterGetFirmsCounter)
}

// GetFirmsBeforeCounter returns a count of FirmManagerMock.GetFirms invocations
func (mmGetFirms *FirmManagerMock) GetFirmsBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetFirms.beforeGetFirmsCounter)
}

// Calls returns a list of arguments used in each call to FirmManagerMock.GetFirms.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetFirms *mFirmManagerMockGetFirms) Calls() []*FirmManagerMockGetFirmsParams {
	mmGetFirms.mutex.RLock()

	argCopy := make([]*FirmManagerMockGetFirmsParams, len(mmGetFirms.callArgs))
	copy(argCopy, mmGetFirms.callArgs)

	mmGetFirms.mutex.RUnlock()

	return argCopy
}

// MinimockGetFirmsDone returns true if the count of the GetFirms invocations corresponds
// the number of defined expectations
func (m *FirmManagerMock) MinimockGetFirmsDone() bool {
	for _, e := range m.GetFirmsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFirmsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFirmsCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFirms != nil && mm_atomic.LoadUint64(&m.afterGetFirmsCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetFirmsInspect logs each unmet expectation
func (m *FirmManagerMock) MinimockGetFirmsInspect() {
	for _, e := range m.GetFirmsMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to FirmManagerMock.GetFirms with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetFirmsMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetFirmsCounter) < 1 {
		if m.GetFirmsMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to FirmManagerMock.GetFirms")
		} else {
			m.t.Errorf("Expected call to FirmManagerMock.GetFirms with params: %#v", *m.GetFirmsMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetFirms != nil && mm_atomic.LoadUint64(&m.afterGetFirmsCounter) < 1 {
		m.t.Error("Expected call to FirmManagerMock.GetFirms")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *FirmManagerMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetFirmsInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *FirmManagerMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *FirmManagerMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetFirmsDone()
}
