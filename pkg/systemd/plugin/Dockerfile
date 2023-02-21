FROM fedora:36


RUN dnf install -y systemd
RUN dnf install -y golang
RUN mkdir /root/checks
ENTRYPOINT /usr/sbin/init