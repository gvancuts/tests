// Copyright (c) 2019 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package docker

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("sysctls", func() {
	var (
		args     []string
		id       string
		stdout   string
		exitCode int
	)

	BeforeEach(func() {
		id = randomDockerName()
	})

	AfterEach(func() {
		Expect(ExistDockerContainer(id)).NotTo(BeTrue())
	})

	Context("sysctls for fs", func() {
		It("should be applied", func() {
			fsValue := "512"
			args = []string{"--name", id, "--rm", "--sysctl", "fs.mqueue.queues_max=" + fsValue, Image, "cat", "/proc/sys/fs/mqueue/queues_max"}
			stdout, _, exitCode = dockerRun(args...)
			Expect(exitCode).To(Equal(0))
			Expect(stdout).To(ContainSubstring(fsValue))
		})
	})

	Context("sysctls for kernel", func() {
		It("should be applied", func() {
			kernelValue := "1024"
			args = []string{"--name", id, "--rm", "--sysctl", "kernel.shmmax=" + kernelValue, Image, "cat", "/proc/sys/kernel/shmmax"}
			stdout, _, exitCode = dockerRun(args...)
			Expect(exitCode).To(Equal(0))
			Expect(stdout).To(ContainSubstring(kernelValue))
		})
	})
})
