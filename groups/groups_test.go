package groups_test

import (
	"github.com/fsamin/intools-engine/common/tests"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/samalba/dockerclient/mockclient"

	"github.com/fsamin/intools-engine/groups"
	"github.com/fsamin/intools-engine/intools"
	"github.com/samalba/dockerclient"
)

var _ = Describe("Groups", func() {

	var (
		engine *tests.IntoolsEngineMock
		cron   *tests.CronMock
		redis  *tests.RedisClientMock
		docker dockerclient.Client
	)

	BeforeEach(func() {
		cron = &tests.CronMock{}
		redis = &tests.RedisClientMock{}
		docker = &mockclient.MockClient{}
		engine = &tests.IntoolsEngineMock{docker, "mock.local:2576", redis, cron}

		intools.Engine = engine
	})

	Describe("Reloading Data from Redis Store", func() {
		Context("With no Redis Store", func() {
			It("Should do nothing", func() {
				groups.Reload()
				Expect(cron.AssertNumberOfCalls(GinkgoT(), "AddJob", 1))
			})
		})
	})
})
