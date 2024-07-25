package simple

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_OpenCSV(t *testing.T) {
	testCases := []struct {
		name               string
		inputPath          string
		expectedDataTables []string
		expectedData       [][]string
	}{
		{
			name:               "loads simple csv file",
			inputPath:          "../testdata/testing.csv",
			expectedDataTables: []string{"testing.csv"},
			expectedData: [][]string{
				{"title 1", "title 2", "title 3", "title 4"},
				{"c", "c", "c", "c"},
				{"b", "2", "3", "4"},
				{"b", "2", "j", "4"},
				{"b", "1", "2", "1"},
				{"b", "4", "3", "2"},
				{"1", "1", "1", "1"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s, err := OpenCSV(tc.inputPath)
			require.NoError(t, err)

			dataTables, err := s.List()
			require.NoError(t, err)
			require.Equal(t, tc.expectedDataTables, dataTables)

			coll, err := s.Get(dataTables[0])
			require.NoError(t, err)

			data := [][]string{}
			for coll.Next() {
				data = append(data, coll.Strings())
			}

			require.Equal(t, tc.expectedData, data)
		})
	}
}
