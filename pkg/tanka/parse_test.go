package tanka

import (
	"testing"

	"github.com/grafana/tanka/pkg/jsonnet"
	"github.com/grafana/tanka/pkg/process"
	"github.com/grafana/tanka/pkg/spec/v1alpha1"
	"github.com/stretchr/testify/assert"
)

func TestEvalJsonnet(t *testing.T) {
	cases := []struct {
		path string
		data interface{}
		envs []*v1alpha1.Config
	}{
		{
			path: "./testdata/cases/array/",
			data: []interface{}{
				[]interface{}{
					map[string]interface{}{"testCase": "nestedArray[0][0]"},
					map[string]interface{}{"testCase": "nestedArray[0][1]"},
				},
				[]interface{}{
					map[string]interface{}{"testCase": "nestedArray[1][0]"},
					map[string]interface{}{"testCase": "nestedArray[1][1]"},
				},
			},
			envs: nil,
		},
		{
			path: "./testdata/cases/object/",
			data: map[string]interface{}{
				"testCase": "object",
			},
			envs: nil,
		},
		{
			path: "./testdata/cases/withspecjson/",
			data: map[string]interface{}{
				"testCase": "object",
			},
			envs: []*v1alpha1.Config{
				{
					APIVersion: v1alpha1.New().APIVersion,
					Kind:       v1alpha1.New().Kind,
					Metadata: v1alpha1.Metadata{
						Name:   "cases/withspecjson",
						Labels: v1alpha1.New().Metadata.Labels,
					},
					Spec: v1alpha1.Spec{
						APIServer: "https://localhost",
						Namespace: "withspec",
					},
					Data: map[string]interface{}{
						"testCase": "object",
					},
				},
			},
		},
		{
			path: "./testdata/cases/withspecjson/main.jsonnet",
			data: map[string]interface{}{
				"testCase": "object",
			},
			envs: []*v1alpha1.Config{
				{
					APIVersion: v1alpha1.New().APIVersion,
					Kind:       v1alpha1.New().Kind,
					Metadata: v1alpha1.Metadata{
						Name:   "cases/withspecjson",
						Labels: v1alpha1.New().Metadata.Labels,
					},
					Spec: v1alpha1.Spec{
						APIServer: "https://localhost",
						Namespace: "withspec",
					},
					Data: map[string]interface{}{
						"testCase": "object",
					},
				},
			},
		},
		{
			path: "./testdata/cases/withenv/main.jsonnet",
			data: map[string]interface{}{
				"apiVersion": v1alpha1.New().APIVersion,
				"kind":       v1alpha1.New().Kind,
				"metadata": map[string]interface{}{
					"name": "withenv",
				},
				"spec": map[string]interface{}{
					"apiServer": "https://localhost",
					"namespace": "withenv",
				},
				"data": map[string]interface{}{
					"testCase": "object",
				},
			},
			envs: []*v1alpha1.Config{
				{
					APIVersion: v1alpha1.New().APIVersion,
					Kind:       v1alpha1.New().Kind,
					Metadata: v1alpha1.Metadata{
						Name: "withenv",
					},
					Spec: v1alpha1.Spec{
						APIServer: "https://localhost",
						Namespace: "withenv",
					},
					Data: map[string]interface{}{
						"testCase": "object",
					},
				},
			},
		},
		{
			path: "./testdata/cases/withenvs/main.jsonnet",
			data: map[string]interface{}{
				"envs": []interface{}{
					map[string]interface{}{
						"apiVersion": v1alpha1.New().APIVersion,
						"kind":       v1alpha1.New().Kind,
						"metadata": map[string]interface{}{
							"name": "withenv1",
						},
						"spec": map[string]interface{}{
							"apiServer": "https://localhost",
							"namespace": "withenv",
						},
						"data": map[string]interface{}{
							"testCase": "object",
						},
					},
					map[string]interface{}{
						"apiVersion": v1alpha1.New().APIVersion,
						"kind":       v1alpha1.New().Kind,
						"metadata": map[string]interface{}{
							"name": "withenv2",
						},
						"spec": map[string]interface{}{
							"apiServer": "https://localhost",
							"namespace": "withenv",
						},
						"data": map[string]interface{}{
							"testCase": "object",
						},
					},
				},
			},
			envs: []*v1alpha1.Config{
				{
					APIVersion: v1alpha1.New().APIVersion,
					Kind:       v1alpha1.New().Kind,
					Metadata: v1alpha1.Metadata{
						Name: "withenv1",
					},
					Spec: v1alpha1.Spec{
						APIServer: "https://localhost",
						Namespace: "withenv",
					},
					Data: map[string]interface{}{
						"testCase": "object",
					},
				},
				{
					APIVersion: v1alpha1.New().APIVersion,
					Kind:       v1alpha1.New().Kind,
					Metadata: v1alpha1.Metadata{
						Name: "withenv2",
					},
					Spec: v1alpha1.Spec{
						APIServer: "https://localhost",
						Namespace: "withenv",
					},
					Data: map[string]interface{}{
						"testCase": "object",
					},
				},
			},
		},
	}

	for _, test := range cases {
		data, envs, e := eval(test.path, jsonnet.Opts{})
		if data == nil {
			assert.NoError(t, e)
		} else if e != nil {
			assert.IsType(t, process.ErrorPrimitiveReached{}, e)
		}
		assert.Equal(t, test.data, data)
		assert.Equal(t, test.envs, envs)
	}
}
