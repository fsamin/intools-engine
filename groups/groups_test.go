package groups_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/samalba/dockerclient/mockclient"
	"github.com/soprasteria/intools-engine/common/tests"

	"github.com/samalba/dockerclient"
	"github.com/soprasteria/intools-engine/groups"
	"github.com/soprasteria/intools-engine/intools"
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
