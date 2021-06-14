#!/usr/bin/bash
name=webcalc

aix_archs=(ppc64)
darwin_archs=(amd64 arm64)
dragonfly_archs=(amd64)
freebsd_archs=(386 amd64 arm arm64)
illumos_archs=(amd64)
linux_archs=(386 amd64 arm arm64 mips mips64 mips64le mipsle ppc64 ppc64le riscv64 s390x)
netbsd_archs=(386 amd64 arm arm64)
openbsd_archs=(386 amd64 arm arm64 mips64)
plan9_archs=(386 amd64 arm)
solaris_archs=(amd64)
windows_archs=(386 amd64 arm)

for arch in ${aix_archs[@]}
do
    build_build_os=aix
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${darwin_archs[@]}
do
    build_build_os=darwin
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${dragonfly_archs[@]}
do
    build_build_os=dragonfly
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${freebsd_archs[@]}
do
    build_build_os=freebsd
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${illumos_archs[@]}
do
    build_build_os=illumos
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${linux_archs[@]}
do
    build_build_os=linux
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${netbsd_archs[@]}
do
    build_build_os=netbsd
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${openbsd_archs[@]}
do
    build_build_os=openbsd
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${plan9_archs[@]}
do
    build_build_os=plan9
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${solaris_archs[@]}
do
    build_build_os=solaris
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done

for arch in ${windows_archs[@]}
do
    build_build_os=windows
    env GOOS=${build_build_os} GOARCH=${arch} go build -o ${name}_${build_build_os}_${arch}
done
