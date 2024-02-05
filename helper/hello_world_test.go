package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	benchmarks := []struct {
		Name    string
		Request string
	}{
		{
			Name:    "Vicry",
			Request: "Vicry",
		},
		{
			Name:    "Fahreza",
			Request: "Fahreza",
		},
		{
			Name:    "Thomy",
			Request: "Thomy",
		},
	}
	for _, benchmark := range benchmarks {
		b.Run(benchmark.Name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHelloWorld(benchmark.Request)
			}
		})
	}
	// helper
}

func BenchmarkSub(b *testing.B) {
	b.Run("vicry", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHelloWorld("vicry")
		}
	})
	b.Run("fahreza", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHelloWorld("fahreza")
		}
	})
	// go test -v -run=tidakada  -bench=BenchmarkSub/vicry
}

func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHelloWorld("vicry")
	}
	// go test -v -bench=.
	// run hanya benchmark saja
	// go test -v -run=testtidakada -bench=.
	// run seluruh
	// go test -v -bench=. ./...
}

func BenchmarkHelloWorldFahreza(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHelloWorld("fahreza")
	}
}

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		Name     string
		Request  string
		expected string
	}{
		{
			Name:     "Vicry",
			Request:  "Vicry",
			expected: "Hello Vicry",
		},
		{
			Name:     "Fahreza",
			Request:  "Fahreza",
			expected: "Hello Fahreza",
		},
		{
			Name:     "Thomy",
			Request:  "Thomy",
			expected: "Hello Thomy",
		},
	}

	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			result := SayHelloWorld(test.Request)
			require.Equal(t, test.expected, result)
		})
	}

}

func TestSubTest(t *testing.T) {
	t.Run("vicry", func(t *testing.T) {
		result := SayHelloWorld("vicry")
		require.Equal(t, "Hello vicry", result, "result must be 'Hello vicry'")
	})
	t.Run("fahreza", func(t *testing.T) {
		result := SayHelloWorld("fahreza")
		require.Equal(t, "hello fahreza", result, "result must be 'Hello fahreza'")
	})
	// Test salah satu sub test
	// go test -v -run=TestSubTest/vicry
}

func TestMain(m *testing.M) {
	fmt.Println("BEFORE UNIT TEST")
	m.Run()
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("Cannot run on windows")
	}
	result := SayHelloWorld("vicry")
	assert.Equal(t, "Hello vicry", result, "result must be 'Hello vicry'")
}

func TestSayHelloWorldAssert(t *testing.T) {
	result := SayHelloWorld("vicry")
	assert.Equal(t, "Hello vicry", result, "result must be 'Hello vicry'")
	fmt.Println("TestSayHelloWorldAssert done")
}

func TestSayHelloWorldRequire(t *testing.T) {
	result := SayHelloWorld("vicry")
	require.Equal(t, "Hello vicry", result, "result must be 'Hello vicry'")
	fmt.Println("TestSayHelloWorldRequire done")
}

func TestSayHelloWorld(t *testing.T) {
	result := SayHelloWorld("vicry")
	if result != "Hello vicry" {
		// panic("Error")
		t.Error("result must be 'Hello vicry'")
	}
	fmt.Println("TestSayHelloWorld done")
}

func TestSayHelloWorldFahreza(t *testing.T) {
	result := SayHelloWorld("fahreza")
	if result != "Hello fahreza" {
		// panic("Error")
		// t.FailNow()
		t.Error("result must be 'Hello fahreza'")
	}
	fmt.Println("TestSayHelloWorldFahreza done")
}
