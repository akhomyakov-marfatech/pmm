// Copyright 2019 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//  http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package management

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/percona/pmm/api/managementpb/json/client/ha_proxy"
)

func TestAddHAProxy(t *testing.T) {
	t.Run("Basic", func(t *testing.T) {
		res := &addHAProxyResult{
			Service: &ha_proxy.AddHAProxyOKBodyService{
				ServiceID:   "/service_id/1",
				ServiceName: "myhost-redis",
			},
		}
		expected := strings.TrimSpace(`
HAProxy Service added.
Service ID  : /service_id/1
Service name: myhost-redis
`)
		assert.Equal(t, expected, strings.TrimSpace(res.String()))
	})
}
