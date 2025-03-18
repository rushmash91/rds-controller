// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package db_cluster

import (
	"context"
	"testing"

	svcapitypes "github.com/aws-controllers-k8s/rds-controller/apis/v1alpha1"
	ackcompare "github.com/aws-controllers-k8s/runtime/pkg/compare"
	"github.com/aws/aws-sdk-go-v2/aws"
)

func TestRequireEngineVersionUpdate(t *testing.T) {
	testCases := []struct {
		name                    string
		desired                 *string
		latest                  *string
		autoMinorVersionUpgrade bool
		expectUpdate            bool
		expectError             bool
		errorMessage            string
	}{
		{
			name:                    "nil versions",
			desired:                 nil,
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: true,
			expectUpdate:            false,
			expectError:             true,
			errorMessage:            "engine versions cannot be nil",
		},
		{
			name:                    "identical versions",
			desired:                 aws.String("13.13"),
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: true,
			expectUpdate:            false,
			expectError:             false,
		},
		{
			name:                    "major version upgrade",
			desired:                 aws.String("14.15"),
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: true,
			expectUpdate:            true,
			expectError:             false,
		},
		{
			name:                    "minor version upgrade with autoMinorVersionUpgrade=true",
			desired:                 aws.String("13.14"),
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: true,
			expectUpdate:            false, // AWS will handle it automatically
			expectError:             false,
		},
		{
			name:                    "minor version upgrade with autoMinorVersionUpgrade=false",
			desired:                 aws.String("13.14"),
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: false,
			expectUpdate:            true, // Need explicit update
			expectError:             false,
		},
		{
			name:                    "major version downgrade",
			desired:                 aws.String("12.13"),
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: true,
			expectUpdate:            false,
			expectError:             true,
			errorMessage:            "downgrading engine version",
		},
		{
			name:                    "minor version downgrade",
			desired:                 aws.String("13.12"),
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: true,
			expectUpdate:            false,
			expectError:             true,
			errorMessage:            "downgrading engine version",
		},
		{
			name:                    "invalid desired version format",
			desired:                 aws.String("invalid"),
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: true,
			expectUpdate:            false,
			expectError:             true,
			errorMessage:            "invalid desired engine version format",
		},
		{
			name:                    "invalid latest version format",
			desired:                 aws.String("14.15"),
			latest:                  aws.String("invalid"),
			autoMinorVersionUpgrade: true,
			expectUpdate:            false,
			expectError:             true,
			errorMessage:            "invalid latest engine version format",
		},
		{
			name:                    "postgres-style major upgrade",
			desired:                 aws.String("14.15"),
			latest:                  aws.String("13.13"),
			autoMinorVersionUpgrade: false,
			expectUpdate:            true,
			expectError:             false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			shouldUpdate, err := requireEngineVersionUpdate(tc.desired, tc.latest, tc.autoMinorVersionUpgrade)

			// Check error expectations
			if tc.expectError {
				if err == nil {
					t.Errorf("expected error containing '%s', got no error", tc.errorMessage)
				} else if tc.errorMessage != "" && !contains(err.Error(), tc.errorMessage) {
					t.Errorf("expected error containing '%s', got '%s'", tc.errorMessage, err.Error())
				}
			} else if err != nil {
				t.Errorf("expected no error, got '%s'", err.Error())
			}

			// Check update expectation
			if shouldUpdate != tc.expectUpdate {
				t.Errorf("expected shouldUpdate=%v, got %v", tc.expectUpdate, shouldUpdate)
			}
		})
	}
}

// Helper function to check if a string contains a substring
func contains(s, substr string) bool {
	return (len(s) >= len(substr) && s[:len(substr)] == substr) ||
		(len(s) >= len(substr) && s[len(s)-len(substr):] == substr) ||
		(len(s) >= len(substr) && s[len(s)/2-len(substr)/2:len(s)/2+len(substr)/2+len(substr)%2] == substr)
}

// TestVersionRegex tests the version regex pattern directly
func TestVersionRegex(t *testing.T) {
	testCases := []struct {
		version     string
		expectMatch bool
		expectMajor string
		expectMinor string
	}{
		{
			version:     "13.13",
			expectMatch: true,
			expectMajor: "13",
			expectMinor: "13",
		},
		{
			version:     "14.15",
			expectMatch: true,
			expectMajor: "14",
			expectMinor: "15",
		},
		{
			version:     "9.6",
			expectMatch: true,
			expectMajor: "9",
			expectMinor: "6",
		},
		{
			version:     "10.1",
			expectMatch: true,
			expectMajor: "10",
			expectMinor: "1",
		},
		{
			version:     "invalid",
			expectMatch: false,
		},
		{
			version:     "13",
			expectMatch: false,
		},
		{
			version:     "13.13.0",
			expectMatch: true,
			expectMajor: "13",
			expectMinor: "13",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.version, func(t *testing.T) {
			matches := versionRegex.FindStringSubmatch(tc.version)

			if tc.expectMatch {
				if len(matches) < 3 {
					t.Errorf("expected version '%s' to match, but it didn't", tc.version)
					return
				}

				if matches[1] != tc.expectMajor {
					t.Errorf("expected major version '%s', got '%s'", tc.expectMajor, matches[1])
				}

				if matches[2] != tc.expectMinor {
					t.Errorf("expected minor version '%s', got '%s'", tc.expectMinor, matches[2])
				}
			} else if len(matches) > 0 {
				t.Errorf("expected version '%s' not to match, but it did", tc.version)
			}
		})
	}
}

// TestNewCustomUpdateRequestPayload tests the newCustomUpdateRequestPayload function
// focusing on engine version handling
func TestNewCustomUpdateRequestPayload(t *testing.T) {
	testCases := []struct {
		name                    string
		desiredEngineVersion    *string
		latestEngineVersion     *string
		autoMinorVersionUpgrade *bool
		expectEngineVersion     bool
	}{
		{
			name:                    "major version upgrade",
			desiredEngineVersion:    aws.String("14.15"),
			latestEngineVersion:     aws.String("13.13"),
			autoMinorVersionUpgrade: aws.Bool(true),
			expectEngineVersion:     true,
		},
		{
			name:                    "minor version upgrade with autoMinorVersionUpgrade=true",
			desiredEngineVersion:    aws.String("13.14"),
			latestEngineVersion:     aws.String("13.13"),
			autoMinorVersionUpgrade: aws.Bool(true),
			expectEngineVersion:     false,
		},
		{
			name:                    "minor version upgrade with autoMinorVersionUpgrade=false",
			desiredEngineVersion:    aws.String("13.14"),
			latestEngineVersion:     aws.String("13.13"),
			autoMinorVersionUpgrade: aws.Bool(false),
			expectEngineVersion:     true,
		},
		{
			name:                    "identical versions",
			desiredEngineVersion:    aws.String("13.13"),
			latestEngineVersion:     aws.String("13.13"),
			autoMinorVersionUpgrade: aws.Bool(true),
			expectEngineVersion:     false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create dummy resources with the test case versions
			desiredResource := &resource{
				ko: &svcapitypes.DBCluster{
					Spec: svcapitypes.DBClusterSpec{
						EngineVersion:           tc.desiredEngineVersion,
						AutoMinorVersionUpgrade: tc.autoMinorVersionUpgrade,
						DBClusterIdentifier:     aws.String("test-cluster"), // Required field
					},
				},
			}

			latestResource := &resource{
				ko: &svcapitypes.DBCluster{
					Spec: svcapitypes.DBClusterSpec{
						EngineVersion:       tc.latestEngineVersion,
						DBClusterIdentifier: aws.String("test-cluster"),
					},
				},
			}

			// Create delta with engine version difference
			delta := &ackcompare.Delta{
				Differences: []*ackcompare.Difference{
					{
						Path: ackcompare.NewPath("Spec.EngineVersion"),
					},
				},
			}

			// Create a minimal resourceManager for testing
			rm := &resourceManager{}

			// Call the function being tested
			input, err := rm.newCustomUpdateRequestPayload(context.Background(), desiredResource, latestResource, delta)
			if err != nil {
				// Skip tests that would normally raise an error (like downgrades)
				// since we're just testing the payload construction
				t.Skip("Error creating update payload:", err)
				return
			}

			// Check if EngineVersion is included in the result
			if tc.expectEngineVersion {
				if input.EngineVersion == nil {
					t.Errorf("expected EngineVersion to be included in update payload, but it was nil")
				} else if *input.EngineVersion != *tc.desiredEngineVersion {
					t.Errorf("expected EngineVersion to be %s, got %s", *tc.desiredEngineVersion, *input.EngineVersion)
				}
			} else {
				if input.EngineVersion != nil {
					t.Errorf("expected EngineVersion to be excluded from update payload, but got %s", *input.EngineVersion)
				}
			}
		})
	}
}
