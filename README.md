# vCPU Calculator

vCPU plays an important role inside a virtual machine instance in today's digital era that is inseparable from cloud computing. You will easily find the number of cores or logical processing of your CPU in the information section of the operating system you are using, but what if you want to know the number of vCPUs that can be utilized by your bare metal.

The number of vCPUs that can be utilized in a bare metal can be calculated using the following calculation:
```
(Threads * Cores) * pCPU
```
- Threads = Total threads (logical processors) per physical cpu
- Cores = Total cores per physical cpu
- pCPUs = Total physical cpu

This software using the formula from above and written in Go version go1.23.0.

## Needed Package to Running Executable
- qemu-kvm
- libvirt-daemon-system

## Needed Package for Development
- libvirt-dev
- golang-go

#### References:
- https://forum.huawei.com/enterprise/en/what-is-machine-learning-explained-in-detail-part-1/thread/667281940820148224-667213860102352896
- https://www.ionos.com/digitalguide/server/know-how/cpu-vs-vcpu/
- https://www.datacenters.com/news/what-is-a-vcpu-and-how-do-you-calculate-vcpu-to-cpu