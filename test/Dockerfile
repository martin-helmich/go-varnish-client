FROM centos:7

WORKDIR /

RUN yum install -y https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm && \
    yum install -y pygpgme yum-utils

COPY test/varnish.repo /etc/yum.repos.d/varnishcache_varnish60.repo

RUN yum -q makecache -y --disablerepo='*' --enablerepo='varnishcache_varnish60' && \
    yum install -y varnish

# this is just for testing purposes
RUN echo "72be6aba-00c4-4908-a99f-0e4eb7cc86ca" > /etc/varnish/secret

ENTRYPOINT [ "/usr/sbin/varnishd"]