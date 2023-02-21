FROM fedora:36


RUN dnf install -y systemd
RUN dnf install -y golang
RUN dnf install -y nginx 
RUN mkdir /root/checks

ENTRYPOINT /usr/sbin/init
