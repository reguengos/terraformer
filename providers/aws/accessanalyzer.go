// Copyright 2020 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package aws

import (
	"context"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
)

var accessanalyzerAllowEmptyValues = []string{"tags."}

type AccessAnalyzerGenerator struct {
	AWSService
}

func (g *AccessAnalyzerGenerator) InitResources() error {
	config, e := g.generateConfig()
	if e != nil {
		return e
	}
	svc := accessanalyzer.New(config)
	p := accessanalyzer.NewListAnalyzersPaginator(svc.ListAnalyzersRequest(&accessanalyzer.ListAnalyzersInput{}))
	var resources []terraform_utils.Resource
	for p.Next(context.Background()) {
		for _, analyzer := range p.CurrentPage().Analyzers {
			resourceName := aws.StringValue(analyzer.Name)
			resources = append(resources, terraform_utils.NewSimpleResource(
				resourceName,
				resourceName,
				"aws_accessanalyzer_analyzer",
				"aws",
				accessanalyzerAllowEmptyValues))
		}
	}
	g.Resources = resources
	return p.Err()
}
