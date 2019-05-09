package srv

import (
	"reflect"
	"testing"
)

func Test_parseAzureResourceID(t *testing.T) {
	testCases := []struct {
		id                 string
		expectedResourceID *resourceID
		expectError        bool
	}{
		{
			// Missing "resourceGroups".
			"/subscriptions/00000000-0000-0000-0000-000000000000//myResourceGroup/",
			nil,
			true,
		},
		{
			// Empty resource group ID.
			"/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups//",
			nil,
			true,
		},
		{
			"random",
			nil,
			true,
		},
		{
			"/subscriptions/6d74bdd2-9f84-11e5-9bd9-7831c1c4c038",
			nil,
			true,
		},
		{
			"subscriptions/6d74bdd2-9f84-11e5-9bd9-7831c1c4c038",
			nil,
			true,
		},
		{
			"/subscriptions/6d74bdd2-9f84-11e5-9bd9-7831c1c4c038/resourceGroups/testGroup1",
			&resourceID{
				SubscriptionID: "6d74bdd2-9f84-11e5-9bd9-7831c1c4c038",
				ResourceGroup:  "testGroup1",
				Provider:       "",
				Path:           map[string]string{},
			},
			false,
		},
		{
			"/subscriptions/6d74bdd2-9f84-11e5-9bd9-7831c1c4c038/resourceGroups/testGroup1/providers/Microsoft.Network",
			&resourceID{
				SubscriptionID: "6d74bdd2-9f84-11e5-9bd9-7831c1c4c038",
				ResourceGroup:  "testGroup1",
				Provider:       "Microsoft.Network",
				Path:           map[string]string{},
			},
			false,
		},
		{
			// Missing leading /
			"subscriptions/6d74bdd2-9f84-11e5-9bd9-7831c1c4c038/resourceGroups/testGroup1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1/",
			nil,
			true,
		},
		{
			"/subscriptions/6d74bdd2-9f84-11e5-9bd9-7831c1c4c038/resourceGroups/testGroup1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1",
			&resourceID{
				SubscriptionID: "6d74bdd2-9f84-11e5-9bd9-7831c1c4c038",
				ResourceGroup:  "testGroup1",
				Provider:       "Microsoft.Network",
				Path: map[string]string{
					"virtualNetworks": "virtualNetwork1",
				},
			},
			false,
		},
		{
			"/subscriptions/6d74bdd2-9f84-11e5-9bd9-7831c1c4c038/resourceGroups/testGroup1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1?api-version=2006-01-02-preview",
			&resourceID{
				SubscriptionID: "6d74bdd2-9f84-11e5-9bd9-7831c1c4c038",
				ResourceGroup:  "testGroup1",
				Provider:       "Microsoft.Network",
				Path: map[string]string{
					"virtualNetworks": "virtualNetwork1",
				},
			},
			false,
		},
		{
			"/subscriptions/6d74bdd2-9f84-11e5-9bd9-7831c1c4c038/resourceGroups/testGroup1/providers/Microsoft.Network/virtualNetworks/virtualNetwork1/subnets/publicInstances1?api-version=2006-01-02-preview",
			&resourceID{
				SubscriptionID: "6d74bdd2-9f84-11e5-9bd9-7831c1c4c038",
				ResourceGroup:  "testGroup1",
				Provider:       "Microsoft.Network",
				Path: map[string]string{
					"virtualNetworks": "virtualNetwork1",
					"subnets":         "publicInstances1",
				},
			},
			false,
		},
		{
			"/subscriptions/34ca515c-4629-458e-bf7c-738d77e0d0ea/resourcegroups/acceptanceTestResourceGroup1/providers/Microsoft.Cdn/profiles/acceptanceTestCdnProfile1",
			&resourceID{
				SubscriptionID: "34ca515c-4629-458e-bf7c-738d77e0d0ea",
				ResourceGroup:  "acceptanceTestResourceGroup1",
				Provider:       "Microsoft.Cdn",
				Path: map[string]string{
					"profiles": "acceptanceTestCdnProfile1",
				},
			},
			false,
		},
		{
			"/subscriptions/34ca515c-4629-458e-bf7c-738d77e0d0ea/resourceGroups/testGroup1/providers/Microsoft.ServiceBus/namespaces/testNamespace1/topics/testTopic1/subscriptions/testSubscription1",
			&resourceID{
				SubscriptionID: "34ca515c-4629-458e-bf7c-738d77e0d0ea",
				ResourceGroup:  "testGroup1",
				Provider:       "Microsoft.ServiceBus",
				Path: map[string]string{
					"namespaces":    "testNamespace1",
					"topics":        "testTopic1",
					"subscriptions": "testSubscription1",
				},
			},
			false,
		},
		{
			"/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/example-resources/providers/Microsoft.ApiManagement/service/service1/subscriptions/22222222-2222-2222-2222-222222222222",
			&resourceID{
				SubscriptionID: "11111111-1111-1111-1111-111111111111",
				ResourceGroup:  "example-resources",
				Provider:       "Microsoft.ApiManagement",
				Path: map[string]string{
					"service":       "service1",
					"subscriptions": "22222222-2222-2222-2222-222222222222",
				},
			},
			false,
		},
	}

	for _, test := range testCases {
		parsed, err := parseAzureResourceID(test.id)
		if test.expectError && err != nil {
			continue
		}
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if !reflect.DeepEqual(test.expectedResourceID, parsed) {
			t.Fatalf("Unexpected resource ID:\nExpected: %+v\nGot:      %+v\n", test.expectedResourceID, parsed)
		}
	}
}
